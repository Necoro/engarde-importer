package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"

	"golang.org/x/text/encoding/charmap"
)

const verbatimPath = "templates/verbatim"

//go:embed templates/verbatim
var verbatim embed.FS

//goland:noinspection GoUnhandledErrorResult
func writeVerbatimFile(outputDir string, entry fs.DirEntry) error {
	inName := path.Join(verbatimPath, entry.Name())

	input, err := verbatim.Open(inName)
	if err != nil {
		return fmt.Errorf("reading bundled file '%s': %w", entry.Name(), err)
	}
	defer input.Close()

	outName := path.Join(outputDir, entry.Name())
	output, err := os.Create(outName)
	if err != nil {
		return fmt.Errorf("creating '%s': %w", outName, err)
	}
	defer output.Close()

	encodedOutput := charmap.ISO8859_1.NewEncoder().Writer(output)

	if _, err = io.Copy(encodedOutput, input); err != nil {
		return fmt.Errorf("writing to '%s': %w", outName, err)
	}

	return nil
}

func writeVerbatim(outputDir string) error {
	entries, err := verbatim.ReadDir(verbatimPath)
	if err != nil {
		return err
	}

	for _, e := range entries {
		if err = writeVerbatimFile(outputDir, e); err != nil {
			return err
		}
	}

	return nil
}

func write(config EngardeConfig) error {
	if err := os.MkdirAll(config.outputDir, 0755); err != nil {
		return fmt.Errorf("creating output directory '%s': %w", config.outputDir, err)
	}

	if err := writeVerbatim(config.outputDir); err != nil {
		return fmt.Errorf("copying default files: %w", err)
	}

	return nil
}
