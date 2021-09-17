package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	botToken    string
	chatID      int64
	startedWork bool
	startTime   int
	endTime     int
	port        string
	url         string
	err         error
)

func getvars() error {
	port = os.Getenv("PORT")
	if len(port) == 0 {
		return fmt.Errorf("PORT is not set")
	}

	url = os.Getenv("PING_URL")
	if len(url) == 0 {
		return fmt.Errorf("PING_URL is not set")
	}

	botToken = os.Getenv("BOT_TOKEN")

	if len(botToken) == 0 {
		return fmt.Errorf("BOT_TOKEN is missing")
	}

	chatID, err = strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 64)

	if err != nil || chatID == 0 {
		return fmt.Errorf("CHAT_ID is missing")
	}
	startTime, err = strconv.Atoi(os.Getenv("TIME_START"))

	if err != nil || len(os.Getenv("TIME_START")) == 0 {
		return fmt.Errorf("TIME_START is missing")
	}

	endTime, err = strconv.Atoi(os.Getenv("TIME_END"))

	if err != nil || len(os.Getenv("TIME_END")) == 0 {
		return fmt.Errorf("TIME_END is missing")
	}
	return nil

}
