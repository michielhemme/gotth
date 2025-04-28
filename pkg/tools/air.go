package tools

import (
	_ "embed"
	"html/template"
	"os"
	"path"
	"strings"

	"github.com/michielhemme/gotth/pkg/lib"
)

//go:embed air.tmpl
var airTmplData []byte
var airTmplName = "air.toml"

func InitializeConfiguration() error {
	cacheDir, err := lib.GetCacheDir()
	if err != nil {
		return err
	}

	outputFilePath := path.Join(cacheDir, airTmplName)
	_, err = os.Stat(outputFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	tailwindCSS, err := GetExecutable("tailwind")
	if err != nil {
		return err
	}
	data := struct {
		TailwindCSS string
		OutputFile  string
	}{
		TailwindCSS: strings.ReplaceAll(string(tailwindCSS), "\\", "\\\\"),
		OutputFile:  lib.AppendIfExe("./tmp/main"),
	}

	tmpl, err := template.New("airConfig").Parse(string(airTmplData))
	if err != nil {
		return err
	}

	f, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tmpl.Execute(f, data)
	if err != nil {
		return err
	}
	return nil
}
