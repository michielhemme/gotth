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

//go:embed templates/index.tmpl
var indexTmpl []byte

//go:embed templates/indexgo.tmpl
var indexGoTmpl []byte

//go:embed templates/custom.tmpl
var customTmpl []byte

//go:embed templates/gitignore.tmpl
var gitignoreTmpl []byte

//go:embed templates/tailwind.config.tmpl
var tailwindConfigJsTmpl []byte

var fileMapping = []File{
	{File: "go.mod", Data: nil},
	{File: "go.sum", Data: nil},
	{File: "main.go", Data: mainTmpl},
	{File: "internal/server/server.go", Data: serverTmpl},
	{File: "internal/logger/logger.go", Data: loggerTmpl},
	{File: "internal/templates/index.templ", Data: indexTmpl},
	{File: "internal/templates/index_templ.go", Data: indexGoTmpl},
	{File: "static/css/custom.css", Data: customTmpl},
	{File: ".gitignore", Data: gitignoreTmpl},
	{File: "tailwind.config.js", Data: tailwindConfigJsTmpl},
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
	for _, step := range []func() error{
		func() error { return runGoModInit(projectDir, modulePath) },
		func() error { return runGoModTidy(projectDir, modulePath) },
		func() error { return runTempl(projectDir) },
	} {
		if err := step(); err != nil {
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

func runGoModInit(projectDir, modulePath string) error {
	if err := lib.RunCommand(lib.Command{
		WorkingDir: projectDir,
		Program:    "go",
		Args:       []string{"mod", "init", modulePath},
	}); err != nil {
		return err
	}
	return nil
}

func runGoModTidy(projectDir, modulePath string) error {
	if err := lib.RunCommand(lib.Command{
		WorkingDir: projectDir,
		Program:    "go",
		Args:       []string{"mod", "tidy"},
	}); err != nil {
		return err
	}
	return nil
}

func runTempl(projectDir string) error {
	if err := lib.RunCommand(lib.Command{
		WorkingDir: projectDir,
		Program:    "templ",
		Args:       []string{"generate"},
	}); err != nil {
		return err
	}
	return nil
}
