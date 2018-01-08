package bot

import (
	"context"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/qhenkart/l2bot/bash"
	"github.com/qhenkart/l2bot/notifications"
)

// startup manages starting the Nox App Player via bash
func startup(ctx context.Context) {
	n := notifications.Context(ctx)

	log.Println("opening nox")
	var openErr error
	for i := 0; i < 4; i++ {
		if openErr = bash.RunCommand("open", "/Applications/Nox App Player.app"); openErr != nil {
			log.Println("unable to open ", openErr)
		}
	}

	if !isOpen() {
		log.Println("failed to open nox")
		n.Send(openErr.Error())
		os.Exit(1)
	}

	log.Println("nox started successfully")

	time.Sleep(30 * time.Second)
}

// isOpen returns true if nox player is currently open and false if it is not
func isOpen() bool {
	if _, err := exec.Command("bash", "-c", "ps cax | grep Nox").Output(); err != nil {
		log.Println("not open")
		return false
	}
	log.Println("open")
	return true
}

// closeNox terminates the application and can be used for a fresh restart of the program
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
