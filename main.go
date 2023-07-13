package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"photovoltaic-system-cron/db"
	"photovoltaic-system-cron/repositories"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	c.AddFunc("5 0 * * *", callDailyWeather)       //every 0.05 AM
	c.AddFunc("15 0 * * *", generateProjectReport) //every 0.15 AM
	c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}

func generateProjectReport() {
	defer db.Disconnect()
	db.Connect()
	projects, err := repositories.GetProjects()
	if err != nil {
		log.Println(err.Error())
	}
	for _, project := range projects {
		callProjectGenerateReport(project.UserId, project.Id)
	}
}

func callProjectGenerateReport(userId int, projectId int) {
	godotenv.Load(".env")
	client := &http.Client{}
	url := "http://localhost:" + os.Getenv("SERVICE_PORT") + "/api/v1/project/generate-report/" + fmt.Sprint(projectId)
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Println(err.Error())
	}

	request.Header.Set("api-key", os.Getenv("APP_API_KEY"))

	q := request.URL.Query()
	q.Add("user-id", fmt.Sprint(userId))
	request.URL.RawQuery = q.Encode()
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
	}
	if resp.StatusCode != http.StatusAccepted {
		err := errors.New("Requesting to project generate-report failed")
		log.Println(err.Error())
		return
	}
}

func callDailyWeather() {
	godotenv.Load(".env")
	client := &http.Client{}
	url := "http://localhost:" + os.Getenv("SERVICE_PORT") + "/api/v1/weather/daily"
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Println(err.Error())
	}

	request.Header.Set("api-key", os.Getenv("APP_API_KEY"))
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
