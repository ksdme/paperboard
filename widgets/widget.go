package widgets

import "html/template"

type Widget interface {
	Init()
	Render() template.HTML
}
