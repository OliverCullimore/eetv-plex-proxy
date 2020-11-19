package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	ssdp "github.com/koron/go-ssdp"
	"github.com/olivercullimore/eetv-plex-proxy/config"
	"github.com/olivercullimore/eetv-plex-proxy/eetv"
	"github.com/olivercullimore/eetv-plex-proxy/utils"
)

// LineupItem struct
type LineupItem struct {
	GuideName   string
	GuideNumber string
	URL         string
}

// LineupStatus struct
type LineupStatus struct {
	ScanInProgress int64
	ScanPossible   int64
	Source         string
	SourceList     []string
}

// DeviceDataSpecVersion struct
type DeviceDataSpecVersion struct {
	Major int64 `xml:"major"`
	Minor int64 `xml:"minor"`
}

// DeviceDataDevice struct
type DeviceDataDevice struct {
	DeviceType   string `xml:"deviceType"`
	FriendlyName string `xml:"friendlyName"`
	Manufacturer string `xml:"manufacturer"`
	ModelName    string `xml:"modelName"`
	ModelNumber  string `xml:"modelNumber"`
	SerialNumber string `xml:"serialNumber"`
	UDN          string `xml:"UDN"`
}

// DeviceData struct
type DeviceData struct {
	XMLName     xml.Name              `xml:"root"`
	XMLNS       string                `xml:"xmlns,attr"`
	SpecVersion DeviceDataSpecVersion `xml:"specVersion"`
	URLBase     string                `xml:"URLBase"`
	Device      DeviceDataDevice      `xml:"device"`
}

// DiscoverData struct
type DiscoverData struct {
	FriendlyName    string
	Manufacturer    string
	ModelNumber     string
	FirmwareName    string
	TunerCount      int64
	FirmwareVersion string
	DeviceID        string
	DeviceAuth      string
	BaseURL         string
	LineupURL       string
}

var proxyHost = "localhost"
var proxyPort = "5004"
var proxyBaseURL = "http://" + proxyHost + ":" + proxyPort + "/"
var configFilepath = "/config/config.json"
var configuration = config.New()
var eetvBaseURL = ""
var eetvAppKey = ""
var eetvTuners int64 = 1
var eetvFriendlyName = "PlexProxy"
var eetvAPI = eetv.New("", "")
var discoverData = new(DiscoverData)

func ssdpAdvertise(quit chan bool) {
	ad, err := ssdp.Advertise(
		"urn:schemas-upnp-org:device:MediaServer:1", // send as "ST"
		"uuid:"+discoverData.DeviceID,               // send as "USN"
		proxyBaseURL+"device.xml",                   // send as "LOCATION"
		"ssdp for EETV Plex Proxy",                  // send as "SERVER"
		3600)                                        // send as "maxAge" in "CACHE-CONTROL"
	if err != nil {
		log.Fatal("Error advertising ssdp: ", err)
	}

	aliveTick := time.Tick(5 * time.Second)

	// run Advertiser infinitely.
	for {
		select {
		case <-aliveTick:
			ad.Alive()
		case <-quit:
			fmt.Println("Closing ssdp service")
			// send/multicast "byebye" message.
			ad.Bye()
			// teminate Advertiser.
			ad.Close()
			return
		}
	}
}

func device(w http.ResponseWriter, r *http.Request) {
	specVersion := DeviceDataSpecVersion{
		Major: 1,
		Minor: 0,
	}
	device := DeviceDataDevice{
		DeviceType:   "urn:schemas-upnp-org:device:MediaServer:1",
		FriendlyName: discoverData.FriendlyName,
		Manufacturer: discoverData.Manufacturer,
		ModelName:    discoverData.ModelNumber,
		ModelNumber:  discoverData.ModelNumber,
		SerialNumber: "",
		UDN:          "uuid:" + discoverData.DeviceID,
	}
	d := &DeviceData{
		XMLNS:       "urn:schemas-upnp-org:device-1-0",
		SpecVersion: specVersion,
		Device:      device,
		URLBase:     discoverData.BaseURL,
	}
	deviceXML, err := xml.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/xml")
	fmt.Fprintf(w, string(deviceXML))
}

func discover(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(discoverData)
}

func lineup(w http.ResponseWriter, r *http.Request) {
	// Set params
	fields := "name,id,zap,isDVB,hidden,rank,isHD,logo,rec"
	allowHidden := "0"
	details := ""
	avoidHD := "0"
	tvOnly := "0"

	// Get channels list from EETV Box
	liveChannels, err := eetvAPI.GetLiveChannels(fields, allowHidden, details, avoidHD, tvOnly)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("Live Channels: %+v", liveChannels)

	// Convert channels list into line-up
	var lineup []LineupItem
	for _, channel := range liveChannels {
		lineup = append(lineup, LineupItem{
			GuideName:   channel.Name,
			GuideNumber: strconv.Itoa(int(channel.Zap)),
			URL:         proxyBaseURL + "Live/Channels/" + channel.ID,
		})
		//fmt.Printf("Live Channels: %+v", channel)
	}
	//fmt.Printf("Lineup: %+v", lineup)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lineup)
}

func liveChannel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	http.Redirect(w, r, eetvBaseURL+"Live/Channels/"+vars["channel"], 301)
}

func lineupPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "")
}

func lineupStatus(w http.ResponseWriter, r *http.Request) {
	lineupStatusData := LineupStatus{
		ScanInProgress: 0,
		ScanPossible:   1,
		Source:         "Antenna",
		SourceList:     []string{"Antenna"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lineupStatusData)
}

func handleRequests() {
	// Create a new instance of a mux router
	r := mux.NewRouter().StrictSlash(true)

	// Handle routes
	r.HandleFunc("/device.xml", device)
	r.HandleFunc("/discover.json", discover)
	r.HandleFunc("/lineup.json", lineup)
	r.HandleFunc("/lineup.post", lineupPost).Methods("GET", "POST")
	r.HandleFunc("/lineup_status.json", lineupStatus)
	r.HandleFunc("/Live/Channels/{channel}", liveChannel)

	fmt.Println(proxyBaseURL)
	log.Fatal(http.ListenAndServe(":"+proxyPort, r))
}

func main() {
	// Check for enviroment variables
	envval, envpresent := os.LookupEnv("PROXY_HOST")
	if envpresent && envval != "" {
		proxyHost = envval
	}
	envval, envpresent = os.LookupEnv("PROXY_PORT")
	if envpresent && envval != "" {
		proxyPort = envval
	}
	envval, envpresent = os.LookupEnv("EETV_IP")
	if envpresent && envval != "" {
		eetvBaseURL = "http://" + envval + "/"
	}
	envval, envpresent = os.LookupEnv("EETV_APP_KEY")
	if envpresent && envval != "" {
		eetvAppKey = envval
	}
	// Set new proxy base URL using the enviroment variables
	proxyBaseURL = "http://" + proxyHost + ":" + proxyPort + "/"
	// Set UUID
	configuration.UUID = uuid.New().String()

	// Check for config file
	if utils.FileExists(configFilepath) {
		err := configuration.Load(configFilepath)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := configuration.Save(configFilepath)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Init EETV API with config
	eetvAPI = eetv.New(eetvBaseURL, eetvAppKey)

	// Get EETV Box Info (for Tuners & Friendly Name)
	eetvInfo, err := eetvAPI.GetInfo()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("Info: %+v", eetvInfo)
	if eetvInfo.System.Tuners > 0 {
		eetvTuners = eetvInfo.System.Tuners
	}
	if eetvInfo.System.FriendlyName != "" {
		reg, err := regexp.Compile("[^A-Za-z0-9\\[\\]]+")
		if err != nil {
			log.Fatal(err)
		}
		eetvFriendlyName = eetvFriendlyName + "-" + reg.ReplaceAllString(eetvInfo.System.FriendlyName, "-")
		fmt.Println("Connected to: " + eetvInfo.System.FriendlyName)
	}

	// Set discover data
	discoverData.FriendlyName = eetvFriendlyName
	discoverData.Manufacturer = "Silicondust"
	discoverData.ModelNumber = "HDTC-2US"
	discoverData.FirmwareName = "hdhomeruntc_atsc"
	discoverData.TunerCount = eetvTuners
	discoverData.FirmwareVersion = "20150826"
	discoverData.DeviceID = configuration.UUID
	discoverData.DeviceAuth = "test1234"
	discoverData.BaseURL = proxyBaseURL
	discoverData.LineupURL = proxyBaseURL + "lineup.json"

	// Advertise server
	var sigTerm = make(chan os.Signal)
	quit := make(chan bool)
	signal.Notify(sigTerm, syscall.SIGTERM)
	signal.Notify(sigTerm, syscall.SIGINT)
	go func() {
		sig := <-sigTerm
		fmt.Printf("caught sig: %+v\n", sig)
		fmt.Println("Waiting for a second to finish processing")
		quit <- true
		time.Sleep(1 * time.Second)
		os.Exit(0)
	}()

	go ssdpAdvertise(quit)

	fmt.Println("Starting EETV Plex Proxy...")
	handleRequests()
}
