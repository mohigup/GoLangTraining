package main

import (
	"html/template"
	"os"
)

func main() {

	l := Language{
		Name: "Java",
		Rank: 5.0,
	}
	fm := template.FuncMap{}
	fm["comp"] = func(rank float32) float32 {
		return rank * 50
	}
	template := template.Must(template.New("").Funcs(fm).Parse(templateString))
	template.Execute(os.Stdout, l)
}

type Language struct {
	Name string
	Rank float32
}

const templateString = `
{{- "Language Information" }}
Name {{ .Name}}
Rank {{ printf "@%.2f" .Rank}}
Complexity {{ comp .Rank | printf "@%.2f"}}
`
