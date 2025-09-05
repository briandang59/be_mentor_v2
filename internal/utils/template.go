package utils

import (
	"bytes"
	"html/template"
	"os"
)

func RenderTemplate(path string, data interface{}) (string, error) {
	tmplBytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	tmpl, err := template.New("email").Parse(string(tmplBytes))
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
