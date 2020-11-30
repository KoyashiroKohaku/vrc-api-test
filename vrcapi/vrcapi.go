package vrcapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var (
	BASE_URL = "https://vrchat.com/api/1"
)

func GetConfig() (*Config, error) {
	resp, err := http.Get(BASE_URL + "/config")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// TODO: HTTP statuc code checking

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(body, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
