package helper

import (
	"os"
	"text/template"
)

func ParseTemplate(templateAsset []byte, data interface{}, destFile string) error {

	funcMap := template.FuncMap{
		"contains": contains,
	}

	tmpl, err := template.New("Template").Funcs(funcMap).Delims("{%", "%}").Parse(string(templateAsset[:]))
	if err != nil {
		return err
	}
	file, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := tmpl.Execute(file, data); err != nil {
		return err
	}
	return nil
}

func contains(list []string, value string) bool {
	for _, curr := range list {
		if curr == value {
			return true
		}
	}
	return false
}
