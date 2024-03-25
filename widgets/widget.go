package widgets

import "html/template"

type Widget interface {
	Render() template.HTML
}
