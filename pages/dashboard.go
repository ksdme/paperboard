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

				.text-2xl { font-size: 4rem; }
				.text-lg { font-size: 1.25rem; }
				.text-md { font-size: 1.15rem; }
				.text-sm { font-size: 1rem; }

				.p-1 { padding: 1rem; }
				.pt-1 { padding-top: 1rem; }
				.pb-1 { padding-bottom: 1rem; }

				.m-1 { margin: 1rem; }
				.mt-1 { margin-top: 1rem; }
				.mb-1 { margin-bottom: 1rem; }

				.rounded { border-radius: 6px; }
				.left-accent { border-left: 6px solid black;  }

				.inverted { color: white; background: black; }
			</style>
		</head>
		<body>
			<div
				style="
					display: flex;
					flex-direction: column;
					gap: 4rem;
				"
			>
				{{ range $widget := . }}
				{{ $widget.Render }}
				{{ end }}
			</div>
		</body>
	</html>
`))

func Dashboard(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var buffer bytes.Buffer
	templ.Execute(&buffer, []widgets.Widget{&widgets.Clock{}, &widgets.Calendar{}})

	w.Header().Add("Content-Type", "text/html")
	io.WriteString(w, buffer.String())
}
