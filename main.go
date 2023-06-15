package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	c.AddFunc("* * * * * *", func() { callDailyWeather() })
	go c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}

func callDailyWeather() {
	request, err := http.Get("http://localhost:" + os.Getenv("SERVICE_PORT") + "/weather/daily")
	if err != nil {
		log.Fatal(err)
		return
	}
	if request.StatusCode != 200 {
		err := errors.New("Requesting to daily weather failed")
		log.Fatal(err)
		return
	}
}
