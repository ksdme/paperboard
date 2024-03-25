package pages

import (
	"bytes"
	"html/template"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ksdme/paperboard/widgets"
)

var templ = template.Must(template.New("dashboard").Parse(`
	<html>
		<head>
			<title>Paperboard • Dashboard</title>
		</head>
		<body>
			{{ range $widget := . }}
			{{ $widget.Render }}
			{{ end }}
		</body>
	</html>
`))

func Dashboard(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var buffer bytes.Buffer
	templ.Execute(&buffer, []widgets.Widget{&widgets.Clock{}})

	w.Header().Add("Content-Type", "text/html")
	io.WriteString(w, buffer.String())
}
