package renderer

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/cosmos-toolkit/cosmos-cli/internal/writer"
)

type Context struct {
	ProjectName       string
	Module            string
	GoVersion         string
	ModulePlaceholder string // e.g. "github.com/your-org/your-app" - replaced in non-.tmpl text files (external templates)
}

func Render(fsys fs.FS, ctx Context, outputDir string) error {
	return fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip template.yaml
		if path == "template.yaml" {
			return nil
		}

		if d.IsDir() {
			return nil
		}

		// Determine output path (support dynamic paths)
		resolvedPath := resolvePath(path, ctx)
		outputPath := filepath.Join(outputDir, resolvedPath)

		// Read file content
		data, err := fs.ReadFile(fsys, path)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", path, err)
		}

		// Check if it's a template file
		if strings.HasSuffix(path, ".tmpl") {
			return renderTemplate(data, outputPath, ctx)
		}

		// For external templates: replace module placeholder in text files
		if ctx.ModulePlaceholder != "" && isTextFile(path) {
			data = []byte(strings.ReplaceAll(string(data), ctx.ModulePlaceholder, ctx.Module))
		}

		return writer.WriteFile(outputPath, data)
	})
}

func resolvePath(path string, ctx Context) string {
	// Replace template variables in path
	tmpl, err := template.New("path").Parse(path)
	if err != nil {
		return path
	}

	var buf strings.Builder
	if err := tmpl.Execute(&buf, ctx); err != nil {
		return path
	}

	// Remove .tmpl extension
	result := strings.TrimSuffix(buf.String(), ".tmpl")

	return result
}

func isTextFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".go", ".mod", ".yaml", ".yml", ".json", ".toml", ".md", ".txt", ".env", ".gitignore":
		return true
	}
	return false
}

func renderTemplate(data []byte, outputPath string, ctx Context) error {
	tmpl, err := template.New("file").Parse(string(data))
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Execute template to buffer
	var buf strings.Builder
	if err := tmpl.Execute(&buf, ctx); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	// Write to file using writer package
	return writer.WriteFile(outputPath, []byte(buf.String()))
}
