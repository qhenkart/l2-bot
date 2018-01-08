package bot

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/qhenkart/l2bot/bash"
	"github.com/qhenkart/l2bot/notifications"
)

func startup() {
	fmt.Println("opening nox")
	var openErr error
	for i := 0; i < 4; i++ {
		if openErr = bash.RunCommand("open", "/Applications/Nox App Player.app"); openErr != nil {
			fmt.Println("unable to open ", openErr)
		}
	}

	if !isOpen() {
		fmt.Println("failed to open nox")
		notifications.Send(openErr.Error())
		os.Exit(1)
	}

	fmt.Println("nox started successfully")

	time.Sleep(30 * time.Second)
}

func isOpen() bool {
	if _, err := exec.Command("bash", "-c", "ps cax | grep Nox").Output(); err != nil {
		fmt.Println("not open")
		return false
	}
	fmt.Println("open")
	return true
}

func closeNox() error {
	if !isOpen() {
		return nil
	}

	fmt.Println("closing nox")
	if _, err := exec.Command("bash", "-c", "killall -v 'Nox App Player'").Output(); err != nil {
		fmt.Println(err.Error())
		return err
	}

	time.Sleep(20 * time.Second)
	return nil
}
