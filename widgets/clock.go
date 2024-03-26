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
	<div class="flex-col centered">
		<h1 class="text-2xl bold">
			{{ .Time }}
		</h1>

		<h3 class="pt-1 text-semibold">
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
