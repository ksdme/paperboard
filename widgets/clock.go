package widgets

import (
	"bytes"
	"html/template"
	"time"
)

type Clock struct{}

type clockTemplateContext struct {
	Time string
	Date string
}

var templ = template.Must(template.New("clock").Parse(`
	<div
		style="
			display: flex;
			flex-direction: column;
			justify-content: center;
			align-items: center;
			gap: 2.5rem;
		"
	>
		<h1
			style="
				margin: 0;
				font-size: 4rem;
				font-weight: 800;
			"
		>
			{{ .Time }}
		</h1>

		<h3
			style="
				margin: 0;
				font-size: 1.25rem;
				font-weight: 600;
			"
		>
			{{ .Date }}
		</h3>
	</div>
`))

func (clock *Clock) Render() template.HTML {
	now := time.Now()

	var buffer bytes.Buffer
	templ.Execute(&buffer, clockTemplateContext{
		Time: now.Format("3:04"),
		Date: now.Format("Monday, _2 January"),
	})

	return template.HTML(buffer.String())
}
