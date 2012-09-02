package monitor

import (
	"errors"
	"github.com/gokyle/gopush/pushover"
	"log"
)

var poConfig *pushover.Identity

// EnablePushover enables email notifications
func EnablePushover() {
	log.Println("[+] monitor enabling pushover notifications")
	notifications["pushover"] = true
}

// DisablePushover disables email notifications
func DisablePushover() {
	log.Println("[+] monitor disabling pushover notifications")
	notifications["pushover"] = false
}

// PushoverEnabled returns true if pushover notifications are enabled.
func PushoverEnabled() bool {
        return notifications["pushover"]
}

func validPushoverConfig(poCfg *pushover.Identity) bool {
	valid := false
	if poCfg != nil && poCfg.Token != "" && poCfg.User != "" {
		valid = true
		cfg := pushover.Authenticate(poCfg.Token, poCfg.User)
		poConfig = &cfg
	}
	return valid
}

func pushoverNotify(err error) error {
	sent := pushover.Notify_titled(*poConfig, err.Error(), "monitor alert")

	if !sent {
		return errors.New("failed to send pushover notification")
	}
	return nil
}
