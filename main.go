package main

import (
	"flag"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/qhenkart/l2bot/bot"
	"github.com/robfig/cron"
)

var notificationLimit = 0

func initCron() {
	log.Println("initializing cron")
	c := cron.New()
	//runs everyday at midnight
	c.AddFunc("0 0 15 * * *", func() {
		log.Println("starting cron job")
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
	logFile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	script := flag.String("script", "", "run an individual script")
	flag.Parse()
	log.Println("script selected: ", *script)
	if *script == "" {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("Panic-Recovery recovery %s", rec)
				bot.Run()
			}
		}()
		initCron()
	} else {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("Panic-Recovery recovery %s", rec)
				bot.Script(*script)
			}
		}()
		bot.Script(*script)
		os.Exit(0)
	}
	Wait()
}
