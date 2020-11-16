package config

import (
	"encoding/json"
	"io/ioutil"
)

// configuration struct
type configuration struct {
	UUID string `json:"uuid"`
}

// New function
func New() configuration {
	config := configuration{}
	return config
}

// Save function
func (c configuration) Save(filename string) error {
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, bytes, 0644)
}

// Load function
func (c configuration) Load(filename string) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return err
	}

	return nil
}
