package widgets

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"time"

	ics "github.com/arran4/golang-ical"
)

type Calendar struct {
	Name     string
	ICS      string
	Timezone string
}

type calendarEntry struct {
	event    *ics.VEvent
	calendar *Calendar
}

type CalendarWidget struct {
	Calendars []Calendar
	entries   []calendarEntry
}

var calendarTempl = template.Must(template.New("calendar").Parse(`
	<div class="flex-col text-md mt-1">
		{{ range . }}
		<div class="flex-col mt-1 p-1 {{ if .InProgress }} rounded inverted {{ else }} left-accent {{ end }}">
			<div class="flex">
				<div class="text-semibold" style="flex-grow: 1;">
					{{ .Name }}
				</div>

				<div>
					{{ .Starts }} &bull; {{ .Ends }}
				</div>
			</div>

			<div class="pt-half">
				{{ .Source }}
			</div>
		</div>
		{{ end }}
	</div>
`))

// Fetch and load all the events into memory.
func (widget *CalendarWidget) LoadTodayEvents() {
	log.Println("widget.calendar", "pulling calendars")
	entries := make([]calendarEntry, 0)

	for _, calendar := range widget.Calendars {
		log.Println("widget.calendar", "pulling calendar,", calendar.Name)
		content, err := http.Get(calendar.ICS)
		if err != nil {
			log.Println("widget.calendar", "failed to load ics feed,", err)
			continue
		}

		cal, err := ics.ParseCalendar(content.Body)
		if err != nil {
			log.Println("widget.calendar", "failed to parse ics feed,", err)
			continue
		}

		yesterday := date(time.Now()).AddDate(0, 0, -1)
		for _, event := range cal.Events() {
			starts, err := event.GetStartAt()
			if err != nil {
				log.Println("widget.calendar", "ignoring event, failed to parse starts,", err)
				continue
			}

			_, err = event.GetEndAt()
			if err != nil {
				log.Println("widget.calendar", "ignoring event, failed to parse ends,", err)
				continue
			}

			// We only care about events that are scheduled for today.
			if starts.After(yesterday) {
				entries = append(entries, calendarEntry{
					event:    event,
					calendar: &calendar,
				})
			}
		}
	}

	widget.entries = entries
}

type item struct {
	Source     string
	Name       string
	Starts     string
	Ends       string
	InProgress bool
}

func (widget *CalendarWidget) Render() template.HTML {
	items := make([]item, 0)

	now := time.Now()
	for _, entry := range widget.entries {
		starts, _ := entry.event.GetStartAt()
		ends, _ := entry.event.GetEndAt()

		// We only care about events for today.
		if !date(starts).Equal(date(starts)) {
			continue
		}

		// Event has ended.
		if ends.Before(now) {
			continue
		}

		name := "Calendar Entry"
		if summary := entry.event.GetProperty(ics.ComponentPropertySummary); summary != nil {
			name = summary.Value
		}

		tz := time.Local
		if entry.calendar.Timezone != "" {
			loc, err := time.LoadLocation(entry.calendar.Timezone)
			if err == nil {
				tz = loc
			} else {
				log.Println(
					"widget.calendar",
					"failed to resolve timezone,",
					entry.calendar.Timezone,
					"for calendar",
					entry.calendar.Name,
					",",
					err,
				)
			}
		}

		items = append(items, item{
			Source:     entry.calendar.Name,
			Name:       name,
			Starts:     starts.In(tz).Format(time.Kitchen),
			Ends:       ends.In(tz).Format(time.Kitchen),
			InProgress: starts.Before(now),
		})
	}

	var buffer bytes.Buffer
	calendarTempl.Execute(&buffer, items)
	return template.HTML(buffer.String())
}

func date(u time.Time) time.Time {
	return u.Truncate(24 * time.Hour)
}
