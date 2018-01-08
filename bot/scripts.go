package bot

import (
	"fmt"
	"log"

	"github.com/qhenkart/l2bot/bash"
	"github.com/qhenkart/l2bot/notifications"
)

const (
	sikuli  = "/Applications/SikuliX.app/run"
	nox     = "/Applications/Nox App Player"
	grind   = sikuli + " -r ./l2bot/l2grind.sikuli"
	dungeon = sikuli + " -r ./l2bot/l2dungeon.sikuli"
	weeklys = sikuli + " -r ./l2bot/l2dailys.sikuli"
	sub     = sikuli + " -r ./l2bot/l2sub.sikuli"
)

var scripts = map[string]string{
	"grind":   grind,
	"dungeon": dungeon,
	"weeklys": weeklys,
	"sub":     sub,
}

var todos = map[string]bool{
	"weeklys":  false,
	"dungeons": false,
}

func runScript(command string, retry int) bool {
	notifications.Send("starting " + command)
	err := bash.RunProgram(command)
	if err != nil {
		log.Println("an error occured, retrying ", retry, err)
		retry++
		if retry > 3 {
			notifications.Send(err.Error())
			return false
		}
		runScript(command, retry)
	}
	notifications.Send(command + " successful")
	return true
}

// Run runs all of the available scripts
func Run() {
	startup()

	if !todos["weekly"] {
		todos["weeklys"] = runScript(weeklys, 0)
	}
	if !todos["dungeon"] {
		todos["dungeon"] = runScript(dungeon, 0)
	}

	runScript(grind, 0)

	notifications.Send(fmt.Sprintf("something happened. Rerunning scripts, todos: %+v", todos))
	if err := closeNox(); err != nil {
		log.Println("unable to close nox ", err)
	}
	Run()
}

// Script runs an individual script
func Script(script string) {
	path, ok := scripts[script]
	if !ok {
		log.Printf("\n\nerror, unable to run script. \nOptions: %+v\n\n", scripts)
	}
	startup()
	if success := runScript(path, 0); !success {
		notifications.Send(fmt.Sprintf("something happened. Rerunning script: %s", path))
		if err := closeNox(); err != nil {
			log.Println("unable to close nox ", err)
		}
	}
}
