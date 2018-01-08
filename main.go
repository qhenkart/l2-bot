package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/qhenkart/l2bot/bot"
	"github.com/robfig/cron"
)

var notificationLimit = 0

func initCron() {
	fmt.Println("initializing cron")
	c := cron.New()
	//runs everyday at midnight
	c.AddFunc("5 15 * * * *", func() {
		fmt.Println("starting cron job")
		bot.Run()
	})
	c.Start()
}

// Wait wait for the interruption signals
func Wait() os.Signal {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	return <-ch
}

func main() {
	script := flag.String("script", "", "run an individual script")
	flag.Parse()
	fmt.Println("script selected: ", *script)
	if *script == "" {
		defer func() {
			if rec := recover(); rec != nil {
				fmt.Printf("Panic-Recovery recovery %s", rec)
				bot.Run()
			}
		}()
		initCron()
	} else {
		defer func() {
			if rec := recover(); rec != nil {
				fmt.Printf("Panic-Recovery recovery %s", rec)
				bot.Script(*script)
			}
		}()
		bot.Script(*script)
		os.Exit(0)
	}
	Wait()
}
