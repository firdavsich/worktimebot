package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	botToken    string
	userID      int64
	startedWork bool
	startTime   int
	endTime     int
	timezone    int
)

func getvars() error {
	botToken := os.Getenv("BOT_TOKEN")
	timezone = 3

	if len(botToken) == 0 {
		return fmt.Errorf("BOT_TOKEN is missing")
	}

	userID, err := strconv.ParseInt(os.Getenv("USER_ID"), 10, 64)

	if err != nil || userID == 0 {
		return fmt.Errorf("USER_ID is missing")
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
