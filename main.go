package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	c.AddFunc("5 0 * * *", callDailyWeather) //every 0.05 AM
	c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}

func callDailyWeather() {
	godotenv.Load(".env")
	client := &http.Client{}
	url := "http://localhost:" + os.Getenv("SERVICE_PORT") + "/weather/daily"
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Println(err.Error())
	}

	request.Header.Set("API_KEY", os.Getenv("APP_API_KEY"))
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
	}
	if resp.StatusCode != http.StatusAccepted {
		err := errors.New("Requesting to daily weather failed")
		log.Println(err.Error())
		return
	}
}
