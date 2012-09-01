package monitor

import (
        "fmt"
        "github.com/kisom/gopush_git/pushover"
        "log"
)

type Panicked struct {
        panicked bool
}

var identity pushover.Identity
var notifications = make(map[string]bool)

func notify(err error) {
        log.Printf("[!] MONITOR critical failure: %s\n", err)
}

func monitorTarget(target (func () error), panicked *Panicked) error {
        var err error = nil
        defer func() {
                if rec := recover(); rec != nil {
                        err = fmt.Errorf("panic recovery: ", rec)
                        panicked.panicked = true
                        notify(err) 
                } else {
                        return
                }
        }()

        err = target()
        if err != nil {
                notify(err)
        }
        fmt.Println("[+] returning from target")
        return err
}

func Monitor(target (func () error)) {
        var panicked = new(Panicked)
        panicked.panicked = false

        for {
                err := monitorTarget(target, panicked)
                if panicked.panicked {
                        fmt.Println("****** PANIC DETECTED")
                }
                if err == nil && !panicked.panicked {
                        log.Println("[+] nominal exit")
                        break
                } 
                panicked.panicked = false
        }

        log.Println("[+] MONITOR nominal exit")
}
