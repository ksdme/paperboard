package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ksdme/paperboard/pages"
	"github.com/ksdme/paperboard/widgets"
)

func main() {
	calendar := &widgets.CalendarWidget{
		Calendars: []widgets.Calendar{
			{
				Name: "Personal",
				ICS:  "",
			},
		},
	}
	calendar.LoadTodayEvents()

	dashboard := pages.Dashboard{
		Widgets: []widgets.Widget{
			&widgets.ClockWidget{},
			calendar,
		},
	}

	router := httprouter.New()
	router.GET("/", dashboard.Handler)

	bind := ":8080"
	log.Println("Starting paperboard on", bind)
	log.Fatal(http.ListenAndServe(bind, router))
}
