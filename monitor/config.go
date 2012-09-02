package monitor

import (
	"encoding/json"
	"github.com/gokyle/gopush/pushover"
	"io/ioutil"
	"os"
	"strings"
)

// ConfigFile points to the file containing the configuration data for the
// notifications
var ConfigFile string = "monitor.json"
var notifications = make(map[string]bool)

type jsonConfig struct {
	Mail     mailConfig        `json:"mail"`
	Pushover pushover.Identity `json:"pushover"`
}

type mailConfig struct {
	Server  string   `json:"server"`
	User    string   `json:"user"`
	Pass    string   `json:"pass"`
	Address string   `json:"address"`
	Port    string   `json:"port"`
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

// ConfigFromEnv reads the configuration from the environment
func ConfigFromEnv() {
	DisableEmail()
	DisablePushover()
	m := mailConfig{os.Getenv("MAIL_SERVER"),
		os.Getenv("MAIL_USER"),
		os.Getenv("MAIL_PASS"),
		os.Getenv("MAIL_ADDRESS"),
		os.Getenv("MAIL_PORT"),
		strings.Split(os.Getenv("MAIL_TO"), ","),
	}
	if validMailConfig(&m) {
		EnableEmail()
	}

	p := pushover.Identity{os.Getenv("PO_APIKEY"),
		os.Getenv("PO_USER"),
	}
	if validPushoverConfig(&p) {
		EnablePushover()
	}
}
