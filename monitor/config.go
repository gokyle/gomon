package monitor

import (
        "encoding/json"
        "github.com/gokyle/gopush/pushover"
        "io/ioutil"
)

// ConfigFile points to the file containing the configuration data for the
// notifications
var ConfigFile string = "monitor.json"
var notifications = make(map[string]bool)

type jsonConfig struct {
	Mail     mailConfig `json:"mail"`
	Pushover pushover.Identity `json:"pushover"`
}

type mailConfig struct {
	Server  string  `json:"server"`
	User    string  `json:"user"`
	Pass    string  `json:"pass"`
	Address string  `json:"address"`
	Port    string     `json:"port"`
        To      []string `json:"to"`
}

func loadConfig(filename string) (*jsonConfig, error) {
        var cfg *jsonConfig

        jsonByte, err := ioutil.ReadFile(filename)
        if err != nil {
                return cfg, err
        }
        
        err = json.Unmarshal(jsonByte, &cfg)
        return cfg, err
}

// ConfigFromJson reads the file specified by ConfigFile
func ConfigFromJson() error {
        DisableEmail()
        DisablePushover()

        cfg, err := loadConfig(ConfigFile)
        if cfg == nil {
                return err
        }
        
        if validMailConfig(&cfg.Mail) {
                EnableEmail()
        }

        if validPushoverConfig(&cfg.Pushover) {
                EnablePushover()
        }

        return err
}
