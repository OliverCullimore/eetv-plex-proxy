package eetv

import (
	"encoding/json"
	"log"
)

// LiveChannel struct
type LiveChannel struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Zap    int64   `json:"zap"`
	Hidden bool    `json:"hidden"`
	Rec    bool    `json:"rec"`
	Logo   string  `json:"logo"`
	Rank   float64 `json:"rank"`
	IsDVB  bool    `json:"isDVB"`
}

// GetLiveChannels function
func (api eetvapi) GetLiveChannels(fields string, allowHidden string, details string, avoidHD string, tvOnly string) ([]LiveChannel, error) {
	if fields == "" {
		fields = "name,id,zap,isDVB,hidden,rank,isHD,logo,rec"
	}
	if allowHidden == "" {
		allowHidden = "0"
	}
	if details == "" {
		details = ""
	}
	if avoidHD == "" {
		avoidHD = "0"
	}
	if tvOnly == "" {
		tvOnly = "0"
	}

	queryParams := map[string]string{
		"fields":      fields,
		"allowHidden": allowHidden,
		"details":     details,
		"avoidHD":     avoidHD,
		"tvOnly":      tvOnly,
	}

	body, err := api.MakeRequest("Live/Channels/getList", "", queryParams)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//fmt.Println("Live Channels: " + string(body))

	var liveChannels []LiveChannel
	err = json.Unmarshal(body, &liveChannels)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return liveChannels, err
}
