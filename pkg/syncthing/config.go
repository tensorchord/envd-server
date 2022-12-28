package syncthing

import (
	"encoding/xml"

	"github.com/syncthing/syncthing/lib/config"
)

// @source: https://docs.syncthing.net/users/config.html
func InitConfig() *config.Configuration {
	return &config.Configuration{
		Version: 37,
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

func GetConfigString(cfg *config.Configuration) (string, error) {
	tmp := struct {
		XMLName xml.Name `xml:"configuration"`
		*config.Configuration
	}{
		Configuration: cfg,
	}

	configByte, err := xml.MarshalIndent(tmp, "", "  ")
	if err != nil {
		return "", err
	}

	return string(configByte), nil
}
