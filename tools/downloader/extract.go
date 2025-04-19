package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func extractSingleFileFromTarGz(tarGzPath, targetFileName, outputPath string) error {
	f, err := os.Open(tarGzPath)
	if err != nil {
		return err
	}
	defer f.Close()

	gzr, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			return fmt.Errorf("file %s not found in archive", targetFileName)
		}
		if err != nil {
			return err
		}

		// Match only the filename (strip any path from the archive)
		_, file := filepath.Split(header.Name)
		if file == targetFileName {
			outFile, err := os.Create(outputPath)
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, tr)
			if err != nil {
				return err
			}
			fmt.Printf("Extracted %s to %s\n", targetFileName, outputPath)
			return nil
		}
	}
}
