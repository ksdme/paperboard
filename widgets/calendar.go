package widgets

import (
	"bytes"
	"html/template"
)

type Calendar struct{}

var calendarTempl = template.Must(template.New("calendar").Parse(`
	<div class="flex-col text-md mt-1">
		{{ range $i := . }}
		<div class="flex-col mt-1 p-1 left-accent">
			<div class="flex">
				<div class="text-semibold" style="flex-grow: 1;">
					Meeting
				</div>

				<div>
					2:30 PM - 4:30 PM
				</div>
			</div>

			<div>
				Google Calendar
			</div>
		</div>
		{{ end }}
	</div>
`))

func (calendar *Calendar) Render() template.HTML {
	var buffer bytes.Buffer
	calendarTempl.Execute(&buffer, [8]int{})
	return template.HTML(buffer.String())
}
