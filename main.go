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

		if time.Minute != 0 {
			continue
		}
		// Skip weekend
		if now.Weekday() == time.Saturday || now.Weekday() == time.Sunday {
			continue
		}

		if now.Hour() == startTime {
			if err := bot.Send(chatID, "Start workday"); err != nil {
				log.Println(err)
			}
			time.Sleep(time.Minute)

		} else if now.Hour() == endTime {
			if err := bot.Send(chatID, "End workday"); err != nil {
				log.Println(err)
			}
			time.Sleep(time.Minute)
		}

	}

}
