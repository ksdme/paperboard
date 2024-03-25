package widgets

import (
	"bytes"
	"html/template"
)

type Clock struct{}

var templ = template.Must(template.New("clock").Parse(`
	<div>
		<h2>
			12:30 PM
		</h2>
	</div>
`))

func (clock *Clock) Render() template.HTML {
	var buffer bytes.Buffer
	templ.Execute(&buffer, nil)
	return template.HTML(buffer.String())
}
