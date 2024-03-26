package pages

import (
	"bytes"
	"html/template"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ksdme/paperboard/widgets"
)

type Dashboard struct {
	Widgets []widgets.Widget
}

var templ = template.Must(template.New("dashboard").Parse(`
	<html>
		<head>
			<title>Paperboard • Dashboard</title>
			<style>
				html, body {
					margin: 0;
					padding: 1rem;
					color: black;
					font-family: Ubuntu;
				}

				h1, h2, h3, h4, h5, h6 {
					margin: 0;
					padding: 0;
				}

				.flex, .flex-col { display: flex; }
				.flex-col { flex-direction: column; }
				.centered { justify-content: center; align-items: center; }

				.text-semibold { font-weight: 600; }
				.text-bold { font-weight: 800; }
				.text-normal { font-weight: 400; }

				.text-2xl { font-size: 4.25rem; }
				.text-lg { font-size: 1.25rem; }
				.text-md { font-size: 1.15rem; }
				.text-sm { font-size: 1rem; }

				.pt-half { padding-top: 0.5rem; }
				.pb-half { padding-bottom: 0.5rem; }

				.p-1 { padding: 1rem; }
				.pt-1 { padding-top: 1rem; }
				.pb-1 { padding-bottom: 1rem; }
				.pl-1 { padding-top: 1rem; }

				.m-1 { margin: 1rem; }
				.mt-1 { margin-top: 1rem; }
				.mb-1 { margin-bottom: 1rem; }

				.rounded { border-radius: 6px; }
				.left-accent { border-left: 4px solid black;  }

				.inverted { color: white; background: black; }
			</style>
			<script>
				setInterval(function() {window.location.reload();}, 60 * 1000)
			</script>
		</head>
		<body>
			<div class="flex-col">
				{{ range $widget := . }}
				{{ $widget.Render }}
				{{ end }}
			</div>
		</body>
	</html>
`))

func (dashboard *Dashboard) Handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var buffer bytes.Buffer
	templ.Execute(&buffer, dashboard.Widgets)

	w.Header().Add("Content-Type", "text/html")
	io.WriteString(w, buffer.String())
}
