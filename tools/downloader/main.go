package main

import (
	"fmt"
	"os"
	"path"
)

type FileDownload struct {
	URL          string `yaml:"url"`
	Name         string `yaml:"name"`
	ExtractTarGz bool   `yaml:"extractTarGz"`
	TargetFile   string `yaml:"targetFile"`
}

type GOOS struct {
	Name     string         `yaml:"name"`
	Download []FileDownload `yaml:"download"`
}

type DownloadConfig []GOOS

func main() {
	config, err := loadConfig("downloads.yaml")
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}

	for _, goos := range config {
		basepath := path.Join("pkg", "tools", goos.Name)
		os.MkdirAll(basepath, 0755)
		for _, file := range goos.Download {
			fmt.Printf("Downloading %s...\n", file.Name)
			filepath := path.Join(basepath, file.Name)
			err := downloadFile(file.URL, filepath)
			if err != nil {
				fmt.Printf("Error downloading %s: %v\n", file.Name, err)
				continue
			}
			if file.ExtractTarGz {
				fmt.Println("Extracting", file.Name)
				err := extractSingleFileFromTarGz(filepath, file.TargetFile, path.Join("pkg", "tools", goos.Name, file.TargetFile))
				if err != nil {
					fmt.Printf("Extraction error: %v\n", err)
				}
			}
		}
	}
}
