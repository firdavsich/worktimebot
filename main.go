package main

import (
	"log"
	"net"
	"net/http"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	err := getvars()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting bot...")

	// heroku hack
	go func() {
		err = http.ListenAndServe(net.JoinHostPort("", port), nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Println(botToken)
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	if err != nil {
		log.Fatal(err)
	}

	for range time.Tick(time.Minute) {

		if time.Now().UTC().Hour() == startTime-timezone && !startedWork {
			startedWork = true
			msg := tgbotapi.NewMessage(userID, "Start workday")
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}

		} else if time.Now().UTC().Hour() == endTime-timezone && startedWork {
			startedWork = false
			msg := tgbotapi.NewMessage(userID, "End workday")
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}

		}

		// heroku hack
		_, err = http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

	}
}
