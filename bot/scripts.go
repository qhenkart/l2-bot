package bot

import (
	"context"
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
	quests  = sikuli + " -r ./l2bot/l2quests.sikuli"
)

var scripts = map[string]string{
	"grind":   grind,
	"dungeon": dungeon,
	"weeklys": weeklys,
	"sub":     sub,
	"quests":  quests,
}

var todos = map[string]bool{
	"weeklys":  false,
	"dungeons": false,
}

// runScript runs a bash
func runScript(ctx context.Context, command string, retry int) bool {
	n := notifications.Context(ctx)
	n.Send("starting " + command)
	err := bash.RunProgram(command)
	if err != nil {
		log.Println("an error occured, retrying ", retry, err)
		retry++
		if retry > 3 {
			n.Send(err.Error())
			return false
		}
		runScript(ctx, command, retry)
	}
	n.Send(command + " successful")
	return true
}

// Run runs all of the available scripts
func Run(ctx context.Context) {
	startup(ctx)
	n := notifications.Context(ctx)

	if !todos["weekly"] {
		todos["weeklys"] = runScript(ctx, weeklys, 0)
	}
	if !todos["dungeon"] {
		todos["dungeon"] = runScript(ctx, dungeon, 0)
	}

	runScript(ctx, grind, 0)

	n.Send(fmt.Sprintf("something happened. Rerunning scripts, todos: %+v", todos))
	if err := closeNox(); err != nil {
		log.Println("unable to close nox ", err)
	}
	Run(ctx)
}

// Script runs an individual script
func Script(ctx context.Context, script string) {
	n := notifications.Context(ctx)
	path, ok := scripts[script]
	if !ok {
		log.Printf("\n\nerror, unable to run script. \nOptions: %+v\n\n", scripts)
	}
	startup(ctx)
	if success := runScript(ctx, path, 0); !success {
		n.Send(fmt.Sprintf("something happened. Rerunning script: %s", path))
		if err := closeNox(); err != nil {
			log.Println("unable to close nox ", err)
		}
	}
}
