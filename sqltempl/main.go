package main

import (
	"strings"
	"text/template"
)

func main() {
	qtmpl := `
SELECT name, email
FROM users
WHERE 1=1
{{if .name}}AND name like {{.name}}{{end}}
{{if .email}}AND email like {{.email}}{{end}}
`
	t, err := template.New("query").Parse(qtmpl)
	if err != nil {
		panic(err)
	}

	options := map[string]any{
		"name":  "?",
		"email": "?",
	}

	sql := new(strings.Builder)
	err = t.Execute(sql, options)
	if err != nil {
		panic(err)
	}
	print(sql.String())
}
