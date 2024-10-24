package templates

import (
	"embed"
)

//go:embed *.tmpl cmd/*.tmpl events/*.tmpl handlers/*.tmpl models/*.tmpl repositories/*.tmpl services/*.tmpl

var FS embed.FS
