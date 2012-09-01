package main

import (
	"fmt"
	"github.com/gokyle/gomon/monitor"
        "math/rand"
        "os"
	"time"
)

func good_job() error {
	fmt.Println("[+] entering nominal code")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("[+] finished")
	return nil
}

func bad_job() error {
	fmt.Println("[+] entering error code")
	time.Sleep(500 * time.Millisecond)
	return fmt.Errorf("something bad has happened!")
}

func panic_job() error {
	fmt.Println("[+] preparing to panic")
	time.Sleep(500 * time.Millisecond)
	panic("this is really happening!")
}

func run() error {
	var err error
	fmt.Println("[+] entering run loop")
	for {
                rand.Seed(time.Now().UnixNano())
		n := rand.Intn(4)
                should_break := false
                fmt.Println("[+] start run loop")

		switch n {
		case 0:
			err = good_job()
		case 1:
			err = bad_job()
		case 2:
			err = panic_job()
                case 3:
                        err = nil
                        fmt.Println("[+] returning")
                        should_break = true
		default:
			return nil
		}

                if err != nil {
                        fmt.Println("[!] error detected")
                } else {
                        fmt.Println("[+] no errors")
                }

                if should_break {
                        fmt.Println("[+] breaking")
                        break
            }
	}

	fmt.Println("[+] finished.")
	return err
}

func main() {
        if len(os.Args) > 1 && os.Args[1] == "persist" {
                monitor.CanExit = false
        }
        monitor.ConfigFile = "monitor.json"
        err := monitor.ConfigFromJson()
        if err != nil {
                fmt.Println("[!] error configuring monitor: ", err)
                os.Exit(1)
        }
	monitor.Monitor(run)
}
