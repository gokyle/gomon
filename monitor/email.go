package monitor

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

var mail_config *mailConfig

// EnableEmail enables email notifications
func EnableEmail() {
	log.Println("[+] monitor enabling email notifications")
	notifications["mail"] = true
}

// DisableEmail disables email notifications
func DisableEmail() {
	log.Println("[+] monitor disabling email notifications")
	notifications["mail"] = false
}

// Returns true if email notifications are enabled.
func EmailEnabled() bool {
	return notifications["mail"]
}

func validMailConfig(mail *mailConfig) bool {
	valid := true
	if mail.Server == "" || mail.User == "" ||
		mail.Pass == "" || mail.Address == "" ||
		mail.Port == "" {
		valid = false
	}

	if valid {
		mail_config = mail
	}
	return valid
}

func stringTo(to []string) string {
	address := strings.Join(to, ";")
	return address
}

func buildMessageBody(err error) string {
	header := fmt.Sprintf("From: %s\n", mail_config.Address)
	header += fmt.Sprintf("To: %s\n", strings.Join(mail_config.To, ";"))
	header += fmt.Sprintf("Subject: monitor alert\n\n")
	header += err.Error()
	return header
}

func mailNotify(err error) error {
	auth := smtp.PlainAuth(mail_config.Address,
		mail_config.User,
		mail_config.Pass,
		mail_config.Server)
	err = smtp.SendMail(mail_config.Server+":"+mail_config.Port,
		auth,
		mail_config.User,
		mail_config.To,
		[]byte(buildMessageBody(err)))
	return err
}
