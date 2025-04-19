package boilerplate

import (
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
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

//go:embed templates/custom.tmpl
var customTmpl []byte

//go:embed templates/gitignore.tmpl
var gitignoreTmpl []byte

//go:embed templates/tailwind.config.tmpl
var tailwindConfigJsTmpl []byte

var fileMapping = []File{
	{File: "main.go", Data: mainTmpl},
	{File: "internal/server/server.go", Data: serverTmpl},
	{File: "internal/logger/logger.go", Data: loggerTmpl},
	{File: "internal/templates/index.templ", Data: indexTmpl},
	{File: "static/css/custom.css", Data: customTmpl},
	{File: ".gitignore", Data: gitignoreTmpl},
	{File: "tailwind.config.js", Data: tailwindConfigJsTmpl},
}

func alreadyInitialized(path string) (bool, error) {
	for _, file := range fileMapping {
		path := filepath.Join(path, file.File)
		_, err := os.Stat(path)
		if err == nil {
			return true, nil
		}
		if !os.IsNotExist(err) {
			return true, err
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

func hasGo(dir string) (bool, error) {
	hasGo := false
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && strings.HasSuffix(d.Name(), ".go") {
			hasGo = true
			return fs.SkipDir // we found a .go file, no need to keep walking
		}
		return nil
	})

	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return hasGo, nil
}

func InitializeProject(modulePath string, childDirectory bool) error {
	moduleSplit := strings.Split(modulePath, "/")
	projectName := moduleSplit[len(moduleSplit)-1]

	templateData := TemplateData{
		ProjectName: projectName,
		ModulePath:  modulePath,
	}

	goProject, err := hasGo(filepath.Join(parseChild(projectName, "", childDirectory)))

	if err != nil {
		return err
	}

	if goProject {
		return fmt.Errorf("directory already containing go files")
	}

	for _, file := range fileMapping {
		dir := filepath.Join(parseChild(projectName, filepath.Dir(file.File), childDirectory))
		filePath := filepath.Join(parseChild(projectName, file.File, childDirectory))

		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
		if err := renderTemplateToFile(file.Data, templateData, filePath); err != nil {
			return err
		}
	}
	return nil
	// if runGoModInit()
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
	cmd := exec.Command("go", "mod", "init", modulePath)
	cmd.Dir = projectDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
