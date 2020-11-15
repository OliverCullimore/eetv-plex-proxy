package eetv

import (
	"io/ioutil"
	"log"
	"net/http"
)

// eetvapi struct
type eetvapi struct {
	BaseURL string
	AppKey  string
}

// New function
func New(baseURL string, appKey string) eetvapi {
	if appKey == "" {
		appKey = "9CCEE7773DE99C8E687AE3AB6009156B8B2A5309" // SHA-1 of [account.service.moonstone]appKey= in netgem-op.ini inside the APK
	}
	api := eetvapi{baseURL, appKey}
	return api
}

// MakeRequest function
func (api eetvapi) MakeRequest(url string, method string, params map[string]string) ([]byte, error) {
	if method == "" {
		method = "GET"
	}

	req, err := http.NewRequest(method, api.BaseURL+url, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Define the query
	q := req.URL.Query()
	// Add query parameters
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

	// Check reponse status
	if resp.Status == "200 OK" {
		// Read reponse body
		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		return responseData, nil
	}
	return nil, err
}
