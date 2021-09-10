package main

import (
	"log"
	"time"

	"github.com/firdavsich/worktimebot/pkg/telegram"
)

func main() {
	err := getvars()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting bot...")

	// heroku hack
	go herokuHack(url)

	bot := telegram.Bot{
		Token:  botToken,
		Silent: false,
	}

	for range time.Tick(time.Minute) {

		if time.Now().Hour() == startTime && !startedWork {
			startedWork = true
			if err := bot.Send(chatID, "Start workday"); err != nil {
				log.Println(err)
			}

		} else if time.Now().Hour() == endTime && startedWork {
			startedWork = false
			if err := bot.Send(chatID, "End workday"); err != nil {
				log.Println(err)
			}

		}

	}
}
