package monitor

import (
	"errors"
	"fmt"
	"log"
)

// Panicked is used to determine whether a panic has occurred; normally
// the function will exit with a nil error, but we also want to know if
// it exited due to a panic.
type Panicked struct {
	panicked bool
}

var CanExit = true

func notify(err error) {
	shouldMail, present := notifications["mail"]
	if present && shouldMail {
		mailErr := mailNotify(err)
		if mailErr != nil {
			log.Println("[!] MONITOR mail notify failed: ",
				mailErr)
		} else {
			log.Println("[+] monitor mail notification sent")
		}
	}

	shouldPush, present := notifications["pushover"]
	if present && shouldPush {
		poErr := pushoverNotify(err)
		if poErr != nil {
			log.Println("[!] MONITOR pushover notify failed: ",
				poErr)
		} else {
			log.Println("[+] monitor pushover notification sent")
		}
	}
	log.Println("[!] MONITOR critical failure: ", err.Error())
}

func monitorTarget(target (func() error), panicked *Panicked) error {
	var err error = nil

	defer func() {
		if rec := recover(); rec != nil {
			err = errors.New(fmt.Sprintf("panic recovery: %s", rec))
			panicked.panicked = true
			notify(err)
		} else {
			return
		}
	}()

	err = target()
	if err != nil {
		fmt.Println("[+] send notification")
		notify(err)
	}
	fmt.Println("[+] returning from target")
	return err
}

func Monitor(target (func() error)) {
	var panicked = new(Panicked)
	panicked.panicked = false

	for {
		err := monitorTarget(target, panicked)
		if panicked.panicked {
			fmt.Println("****** PANIC DETECTED")
		}
		if err == nil && !panicked.panicked {
			log.Println("[+] nominal exit")
			if CanExit {
				break
			}
		}
		panicked.panicked = false
	}

	log.Println("[+] MONITOR nominal exit")
}
