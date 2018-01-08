package bot

import (
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/qhenkart/l2bot/bash"
	"github.com/qhenkart/l2bot/notifications"
)

func startup() {
	log.Println("opening nox")
	var openErr error
	for i := 0; i < 4; i++ {
		if openErr = bash.RunCommand("open", "/Applications/Nox App Player.app"); openErr != nil {
			log.Println("unable to open ", openErr)
		}
	}

	if !isOpen() {
		log.Println("failed to open nox")
		notifications.Send(openErr.Error())
		os.Exit(1)
	}

	log.Println("nox started successfully")

	time.Sleep(30 * time.Second)
}

func isOpen() bool {
	if _, err := exec.Command("bash", "-c", "ps cax | grep Nox").Output(); err != nil {
		log.Println("not open")
		return false
	}
	log.Println("open")
	return true
}

func closeNox() error {
	if !isOpen() {
		return nil
	}

	log.Println("closing nox")
	if _, err := exec.Command("bash", "-c", "killall -v 'Nox App Player'").Output(); err != nil {
		log.Println(err.Error())
		return err
	}

	time.Sleep(20 * time.Second)
	return nil
}
