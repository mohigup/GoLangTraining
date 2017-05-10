package main

import (
	"html/template"
	"os"
)

func main() {

	l := Language{
		Name:   "Java",
		Rank: 5.0,
	}

	template := template.Must(template.New("").Parse(templateString))
	template.Execute(os.Stdout, l)
}

type Language struct {
	Name string
	Rank float32
}

func (l Language) LanguageWithComplexity() float32 {
	return l.Rank * 50
}

const templateString = `
{{- "Language Information" }}
Name {{ .Name}}
Rank {{ printf "@%.2f" .Rank}}
Complexity {{ .LanguageWithComplexity | printf "@%.2f"}}
`
