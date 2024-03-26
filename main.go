package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ksdme/paperboard/pages"
	"github.com/ksdme/paperboard/widgets"
)

func main() {
	clock := &widgets.ClockWidget{}
	clock.Init()

	calendar := &widgets.CalendarWidget{
		Calendars: []widgets.Calendar{
			{
				Name: "Personal",
				ICS:  "",
			},
		},
	}
	calendar.Init()

	dashboard := pages.Dashboard{
		Widgets: []widgets.Widget{
			clock,
			calendar,
		},
	}

	router := httprouter.New()
	router.GET("/", dashboard.Handler)

	bind := ":8080"
	log.Println("Starting paperboard on", bind)
	log.Fatal(http.ListenAndServe(bind, router))
}
