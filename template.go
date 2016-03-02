package goweb

import (
	"html/template"
	"io"
)

var MyTemplates Template

func init() {
	MyTemplates.Load()
}

type Template struct {
	Templates *template.Template
}

func (this *Template) Load() {
	this.Templates = template.Must(template.New("").Delims("[[", "]]").ParseGlob("views/*.*"))
}

func (this *Template) Execute(wr io.Writer, name string, data interface{}) error {
	return this.Templates.ExecuteTemplate(wr, name+".tpl", data)
}
