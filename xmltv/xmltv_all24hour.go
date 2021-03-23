package xmltv

import (
	"encoding/json"
	"log"
)

// DeviceInfoOOH struct
type DeviceInfoOOH struct {
	ApplicationID string `json:"applicationId"`
	SubscriberID  string `json:"subscriberId"`
}

// DeviceInfo struct
type DeviceInfo struct {
	OOH           DeviceInfoOOH `json:"ooh"`
	ApplicationID string        `json:"applicationId"`
	SubscriberID  string        `json:"subscriberId"`
}

// GetInfo function
func (api xmltvapi) GetInfo() (*DeviceInfo, error) {
	body, err := api.MakeRequest("feed/6743", "", nil)
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
