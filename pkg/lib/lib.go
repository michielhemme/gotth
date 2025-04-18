package lib

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"

	"github.com/michielhemme/gotth/pkg/logger"
)

func AppendIfExe(input string) (output string) {
	if runtime.GOOS == "windows" {
		output = input + ".exe"
	}
	return
}

func FileChecksum(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}

func EnsureCacheDir(cacheDir string) {
	err := os.MkdirAll(cacheDir, 0755)
	if err != nil {
		logger.Log(1, fmt.Sprintf("failed to create cache dir: %v", err))
	}
}

func GetCacheDir() string {
	base, err := os.UserCacheDir()
	if err != nil {
		logger.Log(1, fmt.Sprintf("error retrieving cache dir: %v", err))
	}
	logger.Log(5, fmt.Sprintf("application cache dir set to: %v", base))
	return filepath.Join(base, ".gotth")
}
