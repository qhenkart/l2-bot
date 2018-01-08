package main

import (
	"context"
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/qhenkart/l2bot/bot"
	"github.com/qhenkart/l2bot/notifications"
	"github.com/robfig/cron"
)

// Config general configuration for the CLI
type Config struct {
	Token    string
	User     string
	Midnight string
}

type key string

var midnight = key("midnight")

func loadConfig() context.Context {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	conf := map[string]string{}
	err := decoder.Decode(&conf)
	if err != nil {
		log.Println("error:", err)
	}

	ctx := notifications.Init(context.Background(), conf)
	ctx = context.WithValue(ctx, midnight, conf["midnight"])

	return ctx

}

func initCron(ctx context.Context) {
	schedule := ctx.Value(midnight).(string)
	if _, err := cron.Parse(schedule); err != nil {
		schedule = "@midnight"
	}
	log.Println("initializing cron")
	c := cron.New()
	//runs everyday at midnight. This might have to be adjusted depending on your server or timezone
	//for example to run at 3:00am (for EST), you would have "0 0 3 * * *" see https://godoc.org/github.com/robfig/cron
	//for more info
	c.AddFunc(schedule, func() {
		log.Println("starting cron job")
		bot.Run(ctx)
	})
	c.Start()
}

// wait keeps the app from shutting down after starting the cron job
func wait() os.Signal {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	return <-ch
}

//initializes logging
func logInit() {
	//initializes logging output
	logFile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	//prints to the terminal as well as the log.txt
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
}

//recovery returns the script in the event of a runtime error
func recovery(ctx context.Context, script string) {
	if rec := recover(); rec != nil {
		log.Printf("Panic-Recovery recovery %s", rec)
		if script != "" {
			bot.Script(ctx, script)
			return
		}
		bot.Run(ctx)
	}
}

func main() {
	//initializes logging
	logInit()

	ctx := loadConfig()
	//a flag to manage the CLI script
	script := flag.String("script", "", "run an individual script")
	flag.Parse()
	//restarts the script in the event of a runtime error
	defer recovery(ctx, *script)
	if *script != "" {
		log.Println("script selected: ", *script)
		bot.Script(ctx, *script)
		os.Exit(0)
	}

	initCron(ctx)
	wait()
}
