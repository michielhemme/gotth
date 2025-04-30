package boilerplate

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/michielhemme/gotth/pkg/lib"
)

type TemplateData struct {
	ProjectName string
	ModulePath  string
}

type File struct {
	File string
	Data []byte
}

//go:embed templates/main.tmpl
var mainTmpl []byte

//go:embed templates/server.tmpl
var serverTmpl []byte

//go:embed templates/logger.tmpl
var loggerTmpl []byte

//go:embed templates/base.tmpl
var baseTmpl []byte

//go:embed templates/basego.tmpl
var baseGoTmpl []byte

//go:embed templates/custom.tmpl
var customTmpl []byte

//go:embed templates/gitignore.tmpl
var gitignoreTmpl []byte

//go:embed templates/tailwind.config.tmpl
var tailwindConfigJsTmpl []byte

//go:embed templates/htmx-min-js.tmpl
var htmlMinJsTmpl []byte

//go:embed templates/json-enc-js.tmpl
var jsonEncJsTmpl []byte

//go:embed templates/home.tmpl
var homeTmpl []byte

//go:embed templates/post.tmpl
var postTmpl []byte

//go:embed templates/get.tmpl
var getTmpl []byte

var fileMapping = []File{
	{File: "go.mod", Data: nil},
	{File: "go.sum", Data: nil},
	{File: "main.go", Data: mainTmpl},
	{File: ".gitignore", Data: gitignoreTmpl},
	{File: "tailwind.config.js", Data: tailwindConfigJsTmpl},
	{File: "internal/server/server.go", Data: serverTmpl},
	{File: "internal/server/get/routes.go", Data: getTmpl},
	{File: "internal/server/post/routes.go", Data: postTmpl},
	{File: "internal/logger/logger.go", Data: loggerTmpl},
	{File: "internal/templates/base.templ", Data: baseTmpl},
	{File: "internal/templates/base_templ.go", Data: baseGoTmpl},
	{File: "internal/templates/home.templ", Data: homeTmpl},
	{File: "static/css/custom.css", Data: customTmpl},
	{File: "static/js/htmx.min.js", Data: htmlMinJsTmpl},
	{File: "static/js/json-enc.js", Data: jsonEncJsTmpl},
}

func alreadyInitialized(path string) (bool, error) {
	for _, file := range fileMapping {
		path := filepath.Join(path, file.File)
		if _, err := os.Stat(path); err == nil {
			return true, fmt.Errorf("project already initialized, found: %v", path)
		}
	}
	return false, nil
}

func parseChild(projectName string, path string, isChild bool) string {
	if isChild {
		return filepath.Join(projectName, path)
	}
	return path
}

func InitializeProject(modulePath string, childDirectory bool) error {
	moduleSplit := strings.Split(modulePath, "/")
	projectName := moduleSplit[len(moduleSplit)-1]
	projectDir := parseChild(projectName, "", childDirectory)

	templateData := TemplateData{
		ProjectName: projectName,
		ModulePath:  modulePath,
	}

	if _, err := alreadyInitialized(projectDir); err != nil {
		return err
	}

	for _, file := range fileMapping {
		if file.Data == nil {
			continue
		}
		dir := filepath.Join(projectDir, filepath.Dir(file.File))
		filePath := filepath.Join(projectDir, file.File)

		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
		if err := renderTemplateToFile(file.Data, templateData, filePath); err != nil {
			return err
		}

	}
	commands := []lib.Command{
		{
			WorkingDir: projectDir,
			Program:    "go",
			Args:       []string{"mod", "init", modulePath},
		},
		{
			WorkingDir: projectDir,
			Program:    "go",
			Args:       []string{"mod", "tidy"},
		},
		{
			Program: "go",
			Args:    []string{"install", "github.com/a-h/templ/cmd/templ@latest"},
		},
		{
			WorkingDir: projectDir,
			Program:    "templ",
			Args:       []string{"generate"},
		},
	}

	for _, cmd := range commands {
		if err := lib.RunCommand(cmd); err != nil {
			return err
		}
	}
	return nil
}

func renderTemplateToFile(tmplByte []byte, templateData TemplateData, targetFile string) error {
	tmpl, err := template.New("template").Parse(string(tmplByte))
	if err != nil {
		return err
	}
	file, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer file.Close()
	return tmpl.Execute(file, templateData)
}
