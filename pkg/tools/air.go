package tools

import (
	_ "embed"
	"fmt"
	"html/template"
	"os"
	"path"
	"strings"

	"github.com/michielhemme/gotth/pkg/lib"
)

//go:embed air.tmpl
var airTmplData []byte
var airTmplName = "air.toml"

func InitializeConfiguration() {
	data := struct {
		TailwindCSS string
		Templ       string
		OutputFile  string
	}{
		TailwindCSS: strings.ReplaceAll(string(GetExecutable("tailwind")), "\\", "\\\\"),
		Templ:       strings.ReplaceAll(string(GetExecutable("templ")), "\\", "\\\\"),
		OutputFile:  lib.AppendIfExe("./tmp/main"),
	}

	tmpl, err := template.New("airConfig").Parse(string(airTmplData))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	outputFilePath := path.Join(lib.GetCacheDir(), airTmplName)
	f, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer f.Close()

	err = tmpl.Execute(f, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Println("✅ air.toml written successfully.")
}
