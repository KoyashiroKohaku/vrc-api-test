package vrcapi

import (
	"testing"
)

func TestGetConfig(t *testing.T) {
	config, err := GetConfig()

	if err != nil {
		t.Errorf("err: %s\n", err)
	}

	if len(config.APIKey) == 0 {
		t.Error("apiKey is blank")
	}
}
