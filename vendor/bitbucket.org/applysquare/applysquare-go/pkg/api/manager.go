package api

import (
	"html/template"

	"github.com/eknkc/amber"
)

type Manager struct {
	templates map[string]*template.Template
}

func NewManager() *Manager {
	return &Manager{
		templates: make(map[string]*template.Template),
	}
}

func (m *Manager) ComplieTemplateDir(dirname string) {
	result := amber.MustCompileDir("tpl/"+dirname, amber.DefaultDirOptions, amber.DefaultOptions)
	for k, v := range result {
		m.templates[dirname+"/"+k] = v
	}
}
