package catalog

import (
	"io/fs"
	"sort"

	"github.com/cosmos-toolkit/cli/internal/loader"
)

// templatesFS is set by embed.go init function
var templatesFS fs.FS

type Catalog struct {
	templates map[string]fs.FS
}

func New() *Catalog {
	c := &Catalog{
		templates: make(map[string]fs.FS),
	}
	c.loadEmbeddedTemplates()
	return c
}

func (c *Catalog) loadEmbeddedTemplates() {
	if templatesFS == nil {
		return
	}

	// templatesFS root is the "templates" dir (api, worker, cli)
	entries, err := fs.ReadDir(templatesFS, ".")
	if err != nil {
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			templateType := entry.Name()
			subFS, err := fs.Sub(templatesFS, templateType)
			if err == nil {
				c.templates[templateType] = subFS
			}
		}
	}
}

func (c *Catalog) GetEmbeddedTemplate(templateType string) (fs.FS, bool) {
	fs, ok := c.templates[templateType]
	return fs, ok
}

func (c *Catalog) ListEmbeddedTypes() []string {
	types := make([]string, 0, len(c.templates))
	for t := range c.templates {
		types = append(types, t)
	}
	return types
}

// TemplateInfo holds metadata for display in lists.
type TemplateInfo struct {
	Type     string
	Name     string
	Version  string
	Features []string
}

// ListTemplates returns all embedded templates with metadata for listing/discovery.
func (c *Catalog) ListTemplates() []TemplateInfo {
	infos := make([]TemplateInfo, 0, len(c.templates))
	for templateType, templateFS := range c.templates {
		tmpl, err := loader.LoadFromFS(templateFS)
		if err != nil {
			infos = append(infos, TemplateInfo{Type: templateType, Name: templateType})
			continue
		}
		infos = append(infos, TemplateInfo{
			Type:     templateType,
			Name:     tmpl.Name,
			Version:  tmpl.Version,
			Features: tmpl.Features,
		})
	}
	sort.Slice(infos, func(i, j int) bool { return infos[i].Type < infos[j].Type })
	return infos
}
