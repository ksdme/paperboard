package widgets

import (
	"bytes"
	"html/template"
)

type Calendar struct{}

var calendarTempl = template.Must(template.New("calendar").Parse(`
	<div
		style="
			display: flex;
			flex-direction: column;
			font-size: 1.15rem;
			margin-top: 1.25rem;
		"
	>
		{{ range $i := . }}
		<div
			style="
				display: flex;
				flex-direction: column;
				border-left: 6px solid black;
				margin-top: 1rem;
				padding: 8px;
				background: black;
				color: white;
				border-radius: 6px;
			"
		>
			<div
				style="
					display: flex;
				"
			>
				<div
					style="
						flex-grow: 1;
						font-weight: 800;
					"
				>
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
