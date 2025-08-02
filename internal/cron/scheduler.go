package cron

import (
	"log"

	"docs-notify/internal/jobs"

	"github.com/robfig/cron/v3"
)

func StartScheduler(notifyJob *jobs.NotifyJob) {
	c := cron.New()

	// Запускать каждую минуту
	_, err := c.AddJob("@every 1m", notifyJob)
	if err != nil {
		log.Fatalf("Could not add cron job: %v", err)
	}

	c.Start()
	log.Println("Cron scheduler started")
}
