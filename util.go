package main

import (
	"html/template"
	"strings"
)

func getTemplate(ts string) *template.Template {
	t := template.New("base")
	template.Must(t.Parse(LAYOUT_TEMPLATE))
	template.Must(t.Parse(ts))
	return t
}

func titleFromName(name string) string {
	return strings.Title(strings.Replace(name, "-", " ", -1))
}

func nameFromTitle(title string) string {
	return strings.Replace(strings.ToLower(title), " ", "-", -1)
}
