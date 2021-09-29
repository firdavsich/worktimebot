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

	for range time.Tick(time.Second * 20) {
		now = time.Now()

		if now.Minute() != 0 {
			continue
		}
		// Skip weekend
		if now.Weekday() == time.Saturday || now.Weekday() == time.Sunday {
			continue
		}

		var msg string

		if now.Hour() == startTime {
			msg = "Start workday"

		} else if now.Hour() == endTime {
			msg = "End workday"

		}

		if len(msg) != 0 {
			if err := bot.Send(chatID, msg); err != nil {
				log.Println(err)
			}
			time.Sleep(time.Minute)
		}
	}

}
