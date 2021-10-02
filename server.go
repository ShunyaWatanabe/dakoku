package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/carlescere/scheduler"
	"github.com/slack-go/slack"
)

func main() {
	SLACK_TOKEN := os.Getenv("SLACK_TOKEN")
	SLACK_CHANNEL := os.Getenv("SLACK_CHANNEL")

	api := slack.New(SLACK_TOKEN, slack.OptionDebug(true))
	job := func() {
		t := time.Now()
		fmt.Println("Sent at:", t.UTC())
		_, _, err := api.PostMessage(SLACK_CHANNEL, slack.MsgOptionText("Hello, world!", false))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	// Run every 10 seconds but not now.
	scheduler.Every(10).Seconds().Run(job)
	// Run every day at 8:30 am
	// scheduler.Every().Day().At("08:30").Run(job)

	// Keep the program running until an interrupt signal is received.
	runtime.Goexit()
}
