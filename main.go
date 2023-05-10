package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()

	// Add a cron job to run every 5 seconds
	c.AddFunc("*/5 * * * * *", func() {
		fmt.Println("Running job at:", time.Now().String())
	})

	// Add another cron job to run every minute
	c.AddFunc("0 * * * * *", func() {
		fmt.Println("Running job at:", time.Now().String())
	})

	c.Start()

	// Wait for the cron job to run a few times
	time.Sleep(time.Minute)

	// Stop the cron job
	c.Stop()
}
