package syncthing

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/syncthing/syncthing/lib/config"
)

func initConfig() *config.Configuration {
	return &config.Configuration{
		Version: 0,
		GUI: config.GUIConfiguration{
			Enabled:    true,
			RawAddress: "0.0.0.0:8384",
			APIKey:     "envd",
			Theme:      "default",
		},
		Options: config.OptionsConfiguration{
			GlobalAnnEnabled:     false,
			LocalAnnEnabled:      false,
			ReconnectIntervalS:   1,
			StartBrowser:         true, // TODO: disable later
			NATEnabled:           false,
			URAccepted:           1,
			URPostInsecurely:     false,
			URInitialDelayS:      1800,
			AutoUpgradeIntervalH: 0, // Disable auto upgrade
			StunKeepaliveStartS:  0, // Disable STUN keepalive
		},
	}
}

func WriteConfig(path string, cfg *config.Configuration) error {
	file, err := xml.MarshalIndent(cfg, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("notes1.xml", file, 0644)
	if err != nil {
		return err
	}

	return nil
}
