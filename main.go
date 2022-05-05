package main

import (
	"fmt"
	"time"
	"os"
	"net/http"

	"github.com/gin-gonic/gin"
	ical "github.com/arran4/golang-ical"
	"github.com/joho/godotenv"
)

type Calendar struct {
	Title string
	StartAt time.Time
	EndAt time.Time
	ZoomURL string
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("can not load .env!")
	}

	r := gin.Default()
	r.GET("/icals", func(c *gin.Context) {
		ical := generateIcal()
		c.String(http.StatusOK, ical)
	})
	r.Run()
}

func generateIcal() string {
	cal := ical.NewCalendar()
	cal.SetMethod(ical.MethodRequest)

	var calendars []Calendar

	for i := 0; i < 10; i++ {
		calendar := Calendar{
			Title: fmt.Sprintf("HALO %d", i),
			StartAt: time.Now().Add(time.Duration(i) * time.Hour),
			EndAt: time.Now().Add(time.Duration(i + 1) * time.Hour),
			ZoomURL: os.Getenv("ZOOM_URL"),
		}

		calendars = append(calendars, calendar)
	}

	for _, calendar := range calendars {
		event := cal.AddEvent(calendar.Title)
		event.SetStartAt(calendar.StartAt)
		event.SetEndAt(calendar.EndAt)
		event.SetSummary(calendar.Title)
		event.SetDescription(calendar.Title)
		event.SetDescription(calendar.ZoomURL)
		event.SetURL(calendar.ZoomURL)
		event.SetLocation(calendar.ZoomURL)
	}

	return cal.Serialize()
}