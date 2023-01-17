// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

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
			RawAddress: "127.0.0.1:8384",
			APIKey:     "envd",
			Theme:      "default",
		},
		Options: config.OptionsConfiguration{
			GlobalAnnEnabled:     false,
			LocalAnnEnabled:      false,
			ReconnectIntervalS:   1,
			StartBrowser:         true, // TODO: disable later
			NATEnabled:           false,
			URAccepted:           -1, // Disallow telemetry
			URPostInsecurely:     false,
			URInitialDelayS:      1800,
			AutoUpgradeIntervalH: 0, // Disable auto upgrade
			StunKeepaliveStartS:  0, // Disable STUN keepalive
		},
	}
}

func GetConfigByte(cfg *config.Configuration) ([]byte, error) {
	tmp := struct {
		XMLName xml.Name `xml:"configuration"`
		*config.Configuration
	}{
		Configuration: cfg,
	}

	configByte, err := xml.MarshalIndent(tmp, "", "  ")
	if err != nil {
		return []byte{}, err
	}

	return configByte, nil
}
