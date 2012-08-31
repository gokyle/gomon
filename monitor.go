package monitor

import (
        "github.com/kisom/gopush_git/pushover"
        //"log"
        //"net/smtp"
)


var identity pushover.Identity
var notifications = make(map[string]bool)

func notify(err error) {
        log.Printf("[!] MONITOR critical failure: %s\n", err)
}

func monitorTarget(target (func () error)) error {
        var err error = nil

        defer func() {
                if err = recover(); err != nil {
                        notify(err) 
                }
        }

        err = target()
        return err
}

func Monitor(target (func () error)) {
        for {
                err := monitorTarget(target)
                if err == nil {
                        break
                }
        }

        fmt.Println("[+] MONITOR nominal exit")
}
