package tools

import (
	"crypto/sha256"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/michielhemme/gotth/pkg/lib"
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

func GetExecutable(option string) (Executable, error) {
	binary, ok := Tools[option]
	if !ok {
		return Executable(""), fmt.Errorf("GetExecutable selected option does not exist: %v", option)
	}
	cacheDir, err := lib.GetCacheDir()
	if err != nil {
		return Executable(""), err
	}
	return Executable(path.Join(cacheDir, binary.Name)), nil
}

func ExecuteCommand(executable Executable, args ...string) error {
	cmd := exec.Command(string(executable), args...)
	cmd.Dir, _ = os.Getwd()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func verifyHashMatch(path string, data []byte) (bool, error) {
	existingHash, err := lib.FileChecksum(path)
	if err != nil {
		return false, err
	}
	expectedHash := sha256.Sum256(data)
	if string(existingHash) == string(expectedHash[:]) {
		return false, nil
	}
	return true, nil
}

func hasTool(tool string) (bool, error) {
	cacheDir, err := lib.GetCacheDir()
	if err != nil {
		return false, err
	}
	binaryPath := filepath.Join(cacheDir, tool)

	_, err = os.Stat(binaryPath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func InitializeTools() error {
	cacheDir, err := lib.GetCacheDir()
	if err != nil {
		return err
	}
	lib.EnsureCacheDir(cacheDir)

	for _, data := range Tools {
		needsWrite := true
		binaryPath := filepath.Join(cacheDir, data.Name)
		// Check if tool is already present on the system
		exists, err := hasTool(data.Name)
		if err != nil {
			return err
		}
		// Check if tool has right checksum
		correctHash, err := verifyHashMatch(binaryPath, data.Data)
		if err != nil {
			return err
		}
		if exists && correctHash {
			needsWrite = false
		}
		// Write file to disk if it is not present or has invalid checksum
		if needsWrite {
			err := os.WriteFile(binaryPath, data.Data, 0755)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
