package templates

import (
	"embed"
)

//go:embed *.tmpl cmd/*.tmpl events/*.tmpl infrastructures/repositories/*.tmpl interfaces/handlers/*.tmpl interfaces/middleware/*.tmpl models/*.tmpl services/*.tmpl

var FS embed.FS
