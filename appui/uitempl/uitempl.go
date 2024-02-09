package uitempl

import (
	"html/template"

	"worble.ow6.foo/appui/uimodels"
	"worble.ow6.foo/com/templatefuncs"
)

func InitTemplates() (*template.Template, error) {
	funcMap := template.FuncMap{
		"sub":                 templatefuncs.Sub,
		"span":                templatefuncs.Span,
		"mapGuessCodeToClass": uimodels.MapGuessCodeToClass,
	}
	return template.New("").Funcs(funcMap).ParseGlob("./ui/html/*.html")
}
