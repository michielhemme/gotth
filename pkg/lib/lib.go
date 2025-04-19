package lib

import (
	"crypto/sha256"
	"io"
	"os"
	"path/filepath"
	"runtime"
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

func EnsureCacheDir(cacheDir string) error {
	err := os.MkdirAll(cacheDir, 0755)
	if err != nil {
		return err
	}
	return nil
}

func GetCacheDir() (string, error) {
	base, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, ".gotth"), nil
}
