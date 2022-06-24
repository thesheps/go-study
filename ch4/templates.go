package main

import (
	"html/template"
	"os"
)

type Data struct {
	A string
	B template.HTML
}

func main() {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))

	d := Data{"Test1", "Test2"}
	t.Execute(os.Stdout, d)
}
