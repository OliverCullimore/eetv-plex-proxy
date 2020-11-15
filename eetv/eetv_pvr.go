package eetv

import (
	"encoding/json"
	"log"
)

// RecordingOriginalEvent struct
type RecordingOriginalEvent struct {
	ID                  int64  `json:"id"`
	Name                string `json:"name"`
	Text                string `json:"text"`
	Image               string `json:"image"`
	StartTime           int64  `json:"startTime"`
	Duration            int64  `json:"duration"`
	EndTime             int64  `json:"endTime"`
	Description         string `json:"description"`
	Category            string `json:"category"`
	ParentalRating      string `json:"parentalRating"`
	ProgramCRID         string `json:"programCRID"`
	RecommendationCRID  string `json:"recommendationCRID"`
	Audio               string `json:"audio"`
	Video               string `json:"video"`
	Rank                int64  `json:"rank"`
	Icon                string `json:"icon"`
	ChannelID           string `json:"channelId"`
	ChannelName         string `json:"channelName"`
	ProgramName         string `json:"programName"`
	SerieID             string `json:"serieId"`
	EpisodeInfo         string `json:"episodeInfo"`
	Year                int64  `json:"year"`
	CategoryID          string `json:"categoryID"`
	EpisodeNum          string `json:"episodeNum"`
	UseCategoryForImage bool   `json:"useCategoryForImage"`
	SeasonDescription   string `json:"seasonDescription"`
	ChannelZap          int64  `json:"channelZap"`
	SeriesGroupURL      string `json:"seriesGroupUrl"`
	Updating            bool   `json:"updating"`
	ImageType           string `json:"imageType"`
}

// RecordingEvent struct
type RecordingEvent struct {
	ID                  int64                  `json:"id"`
	Name                string                 `json:"name"`
	Text                string                 `json:"text"`
	StartTime           int64                  `json:"startTime"`
	Duration            int64                  `json:"duration"`
	EndTime             int64                  `json:"endTime"`
	Description         string                 `json:"description"`
	Category            string                 `json:"category"`
	ParentalRating      string                 `json:"parentalRating"`
	ProgramCRID         string                 `json:"programCRID"`
	RecommendationCRID  string                 `json:"recommendationCRID"`
	Audio               string                 `json:"audio"`
	Video               string                 `json:"video"`
	Rank                int64                  `json:"rank"`
	Icon                string                 `json:"icon"`
	ChannelID           string                 `json:"channelId"`
	ChannelName         string                 `json:"channelName"`
	ProgramName         string                 `json:"programName"`
	SerieID             string                 `json:"serieId"`
	OriginalEvent       RecordingOriginalEvent `json:"originalEvent"`
	EpisodeInfo         string                 `json:"episodeInfo"`
	Year                int64                  `json:"year"`
	CategoryID          string                 `json:"categoryID"`
	EpisodeNum          string                 `json:"episodeNum"`
	UseCategoryForImage bool                   `json:"useCategoryForImage"`
	SeasonDescription   string                 `json:"seasonDescription"`
	ChannelZap          int64                  `json:"channelZap"`
	SeriesGroupURL      string                 `json:"seriesGroupUrl"`
	Updating            bool                   `json:"updating"`
}

// Recording struct
type Recording struct {
	ID         string         `json:"id"`
	Name       string         `json:"name"`
	Size       int64          `json:"size"`
	URL        string         `json:"url"`
	StartTime  int64          `json:"startTime"`
	Duration   int64          `json:"duration"`
	EndTime    int64          `json:"endTime"`
	State      string         `json:"state"`
	Event      RecordingEvent `json:"event"`
	Repeat     string         `json:"repeat"`
	IsHD       bool           `json:"isHD"`
	Prio       int64          `json:"prio"`
	IsCatchup  bool           `json:"isCatchup"`
	DeviceName string         `json:"deviceName"`
	NetworkID  string         `json:"networkId"`
	GroupID    string         `json:"groupId"`
	IsDeepLink bool           `json:"isDeepLink"`
}

// GetRecordings function
func (api eetvapi) GetRecordings(recordingtype string, avoidHD string, tvOnly string) ([]Recording, error) {
	if recordingtype == "" {
		recordingtype = "regular"
	}
	if avoidHD == "" {
		avoidHD = "0"
	}
	if tvOnly == "" {
		tvOnly = "0"
	}

	queryParams := map[string]string{
		"type":    recordingtype,
		"avoidHD": avoidHD,
		"tvOnly":  tvOnly,
	}

	body, err := api.MakeRequest("PVR/Records/getList", "", queryParams)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//fmt.Println("Recordings: " + string(body))

	var recordings []Recording
	err = json.Unmarshal(body, &recordings)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return recordings, err
}
