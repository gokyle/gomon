package monitor

import (
        "log"
        "net/smtp"
)

var mail_config *mailConfig

// EnableEmail enables email notifications
func EnableEmail() {
        log.Println("[+] monitor enabling email notifications")
	notifications["email"] = true
}

// DisableEmail disables email notifications
func DisableEmail() {
        log.Println("[+] monitor disabling email notifications")
	notifications["email"] = false
}

func validMailConfig(mail *mailConfig) bool {
        valid := true
        if mail.Server == "" || mail.User == ""     ||
           mail.Pass   == "" || mail.Address == ""  ||
           mail.Port   == "" {
                valid = false
        }
        
        if valid {
                mail_config = mail
        }
        return valid
}

func mailNotify(err error) error {
        auth := smtp.PlainAuth(mail_config.Address, 
                               mail_config.User,
                               mail_config.Pass,
                               mail_config.Server)        
        //subject := "monitor alert"
        body := err.Error()
        err = smtp.SendMail(mail_config.Server + ":" + mail_config.Port,
                            auth,
                            mail_config.User,
                            mail_config.To,
                            []byte(body))
        return err
}
