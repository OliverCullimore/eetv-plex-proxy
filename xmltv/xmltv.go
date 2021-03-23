package xmltv

import (
	"io/ioutil"
	"log"
	"net/http"
)

// xmltvapi struct
type xmltvapi struct {
	BaseURL string
}

// New function
func New() xmltvapi {
	api := xmltvapi{"http://www.xmltv.co.uk/"}
	return api
}

// MakeRequest function
func (api xmltvapi) MakeRequest(url string, method string, params map[string]string) ([]byte, error) {
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
