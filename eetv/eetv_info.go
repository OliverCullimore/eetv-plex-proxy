package eetv

import (
	"encoding/json"
	"log"
)

// DeviceInfoSystem struct
type DeviceInfoSystem struct {
	Time         int64  `json:"time"`
	Tuners       int64  `json:"tuners"`
	FriendlyName string `json:"friendlyName"`
}

// DeviceInfoCatchup struct
type DeviceInfoCatchup struct {
	Enabled   bool `json:"enabled"`
	Available bool `json:"available"`
}

// DeviceInfoFVP struct
type DeviceInfoFVP struct {
	Enabled   bool `json:"enabled"`
	Available bool `json:"available"`
}

// DeviceInfoPVRStatusCatchup struct
type DeviceInfoPVRStatusCatchup struct {
	IsRemovedBeforeExpiration bool  `json:"isRemovedBeforeExpiration"`
	RealValidity              int64 `json:"realValidity"`
	DisplayedValidity         int64 `json:"displayedValidity"`
	AppliedValidity           int64 `json:"appliedValidity"`
	LastChanceThreshold       int64 `json:"lastChanceThreshold"`
}

// DeviceInfoPVRStatusDisk struct
type DeviceInfoPVRStatusDisk struct {
	Percent      float64 `json:"percent"`
	UsedSpace    int64   `json:"usedSpace"`
	TotalSpace   int64   `json:"totalSpace"`
	WarningLevel int64   `json:"warningLevel"`
}

// DeviceInfoPVRStatus struct
type DeviceInfoPVRStatus struct {
	Catchup      DeviceInfoPVRStatusCatchup `json:"catchup"`
	DownloadToGo bool                       `json:"downloadToGo"`
	Disk         DeviceInfoPVRStatusDisk    `json:"disk"`
}

// DeviceInfoPVR struct
type DeviceInfoPVR struct {
	Enabled   bool                `json:"enabled"`
	Status    DeviceInfoPVRStatus `json:"status"`
	Available bool                `json:"available"`
}

// DeviceInfoLive struct
type DeviceInfoLive struct {
	Enabled   bool `json:"enabled"`
	Available bool `json:"available"`
}

// DeviceInfoEPG struct
type DeviceInfoEPG struct {
	End          int64  `json:"end"`
	Timer        bool   `json:"timer"`
	WarningWords string `json:"warningWords"`
	Begin        int64  `json:"begin"`
}

// DeviceInfoAMS struct
type DeviceInfoAMS struct {
	Enabled bool `json:"enabled"`
}

// DeviceInfoUserRights struct
type DeviceInfoUserRights struct {
	Available []string `json:"available"`
	Expired   []string `json:"expired"`
}

// DeviceInfoOOHTimers struct
type DeviceInfoOOHTimers struct {
	Enabled bool `json:"enabled"`
}

// DeviceInfoOOHChannels struct
type DeviceInfoOOHChannels struct {
	Enabled bool `json:"enabled"`
}

// DeviceInfoOOHRegister struct
type DeviceInfoOOHRegister struct {
	Enabled bool `json:"enabled"`
}

// DeviceInfoOOHPreRegister struct
type DeviceInfoOOHPreRegister struct {
	Enabled bool `json:"enabled"`
}

// DeviceInfoOOHCMD struct
type DeviceInfoOOHCMD struct {
	Enabled bool `json:"enabled"`
}

// DeviceInfoOOHConfigurationHosts struct
type DeviceInfoOOHConfigurationHosts struct {
	PVR           string `json:"pvr"`
	EPG           string `json:"epg"`
	CMD           string `json:"cmd"`
	CRM           string `json:"crm"`
	AuthDevice    string `json:"authdevice"`
	UpgradeDevice string `json:"upgradedevice"`
}

// DeviceInfoOOHConfiguration struct
type DeviceInfoOOHConfiguration struct {
	Hosts                DeviceInfoOOHConfigurationHosts `json:"hosts"`
	MaxPairedDevices     int64                           `json:"maxPairedDevices"`
	ExpiracyDelay        int64                           `json:"expiracyDelay"`
	StblistPollingDelay  string                          `json:"stblistpollingdelay"`
	CommandsPollingDelay string                          `json:"commandspollingdelay"`
	NpvrPollingDelay     string                          `json:"npvrpollingdelay"`
}

// DeviceInfoOOH struct
type DeviceInfoOOH struct {
	Timers        DeviceInfoOOHTimers        `json:"timers"`
	Channels      DeviceInfoOOHChannels      `json:"channels"`
	Register      DeviceInfoOOHRegister      `json:"register"`
	PreRegister   DeviceInfoOOHPreRegister   `json:"preregister"`
	CMD           DeviceInfoOOHCMD           `json:"cmd"`
	Configuration DeviceInfoOOHConfiguration `json:"configuration"`
	ApplicationID string                     `json:"applicationId"`
	SubscriberID  string                     `json:"subscriberId"`
}

// DeviceInfo struct
type DeviceInfo struct {
	System     DeviceInfoSystem     `json:"system"`
	Catchup    DeviceInfoCatchup    `json:"catchup"`
	FVP        DeviceInfoFVP        `json:"fvp"`
	PVR        DeviceInfoPVR        `json:"pvr"`
	Live       DeviceInfoLive       `json:"live"`
	EPG        DeviceInfoEPG        `json:"epg"`
	AMS        DeviceInfoAMS        `json:"ams"`
	UserRights DeviceInfoUserRights `json:"userrights"`
	OOH        DeviceInfoOOH        `json:"ooh"`
}

// GetInfo function
func (api eetvapi) GetInfo() (*DeviceInfo, error) {
	queryParams := map[string]string{
		"appKey": api.AppKey,
	}

	body, err := api.MakeRequest("UPnP/Device/getInfo", "", queryParams)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//fmt.Println("Device Info: " + string(body))

	var info = new(DeviceInfo)
	err = json.Unmarshal(body, &info)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return info, err
}
