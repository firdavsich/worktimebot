package main

import (
	"log"
	"time"

	"github.com/firdavsich/worktimebot/pkg/telegram"
)

func main() {
	var now time.Time

	err := getvars()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting bot...")

	// heroku hack
	go herokuHack()

	bot := telegram.Bot{
		Token:  botToken,
		Silent: false,
	}

	for range time.Tick(time.Minute) {
		now = time.Now()
		// if not weekend
		if now.Weekday() != time.Saturday && now.Weekday() != time.Sunday {
			if now.Hour() == startTime && !startedWork {
				startedWork = true
				if err := bot.Send(chatID, "Start workday"); err != nil {
					log.Println(err)
				}

			} else if now.Hour() == endTime && startedWork {
				startedWork = false
				if err := bot.Send(chatID, "End workday"); err != nil {
					log.Println(err)
				}

			}
		}

	}
}
