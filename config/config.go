package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Frameless         bool     `json:"frameless"`
	X                 int      `json:"x"`
	Y                 int      `json:"y"`
	Width             int      `json:"width"`
	Height            int      `json:"height"`
	Link              string   `json:"link"`
	HideBarAndSaveKey []string `json:"hidebar_save_key"`
	ShowHideWindowKey []string `json:"show_hide_key"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			conf := &Config{
				Frameless:         false,
				X:                 0,
				Y:                 0,
				Width:             350,
				Height:            800,
				Link:              "https://github.com/bruxaodev/overlay-for-chat/blob/main/README.md",
				HideBarAndSaveKey: []string{"ctrl", "shift", "x"},
				ShowHideWindowKey: []string{"ctrl", "shift", "z"},
			}
			SaveConfig(path, conf)
			return conf, nil
		}
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func SaveConfig(path string, config *Config) error {
	bytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, bytes, 0644)
}
