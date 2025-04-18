package tools

import (
	"crypto/sha256"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/michielhemme/gotth/pkg/lib"
	"github.com/michielhemme/gotth/pkg/logger"
)

type Executable string

type Tool struct {
	Name     string
	URL      string
	Filename string
	Archive  bool
}

type Binary struct {
	Name string
	Data []byte
}

var Tools = map[string]Binary{
	"air":      {Name: airBinaryName, Data: airBinaryData},
	"tailwind": {Name: tailwindBinaryName, Data: tailwindBinaryData},
	"templ":    {Name: templBinaryName, Data: templBinaryData},
}

func GetExecutable(option string) Executable {
	binary, ok := Tools[option]
	if !ok {
		logger.Log(1, fmt.Sprintf("GetExecutable selected option does not exist: %v", option))
	}
	return Executable(path.Join(lib.GetCacheDir(), binary.Name))
}

func ExecuteCommand(executable Executable, args ...string) error {
	cmd := exec.Command(string(executable), args...)
	cmd.Dir, _ = os.Getwd()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func InitializeTools() {
	logger.Log(5, "initializing tools")

	cacheDir := lib.GetCacheDir()
	lib.EnsureCacheDir(cacheDir)

	for _, data := range Tools {
		binaryPath := filepath.Join(cacheDir, data.Name)
		needsWrite := true

		if _, err := os.Stat(binaryPath); err == nil {
			existingHash, err := lib.FileChecksum(binaryPath)
			if err == nil {
				expectedHash := sha256.Sum256(data.Data)
				if string(existingHash) == string(expectedHash[:]) {
					needsWrite = false
					logger.Log(5, fmt.Sprintf("binary already exists, skipping: %s", data.Name))
				}
			}
		}

		if needsWrite {
			err := os.WriteFile(binaryPath, data.Data, 0755)
			if err != nil {
				logger.Log(1, "failed to write binary file: %s", err)
			}
			logger.Log(5, fmt.Sprintf("written to disk: %s", data.Name))
		}

	}
	logger.Log(5, "tools initialized")
}
