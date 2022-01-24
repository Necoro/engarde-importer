package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
)

const verbatimPath = "templates/verbatim"
const participantsName = "tireur.txt"
const clubsName = "club.txt"

//go:embed templates/verbatim
var verbatim embed.FS

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

	if _, err = io.Copy(encodedWriter(output), input); err != nil {
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

func write(config EngardeConfig, participants []participant, clubs []club) error {
	if err := os.MkdirAll(config.outputDir, 0755); err != nil {
		return fmt.Errorf("creating output directory '%s': %w", config.outputDir, err)
	}

	if err := writeVerbatim(config.outputDir); err != nil {
		return fmt.Errorf("copying default files: %w", err)
	}

	if err := writeParticipants(config.outputDir, participants); err != nil {
		return fmt.Errorf("writing participant data: %w", err)
	}

	if err := writeClubs(config.outputDir, clubs); err != nil {
		return fmt.Errorf("writing club data: %w", err)
	}

	return nil
}

func writeClubs(outputDir string, clubs []club) error {
	entries := make([]engarde, len(clubs))
	for i := range clubs {
		entries[i] = clubs[i]
	}

	return writeEngarde(outputDir, clubsName, entries)
}

func writeParticipants(outputDir string, participants []participant) error {
	entries := make([]engarde, len(participants))
	for i := range participants {
		entries[i] = participants[i]
	}

	return writeEngarde(outputDir, participantsName, entries)
}

func writeEngarde(outputDir, fileName string, entries []engarde) error {
	outName := path.Join(outputDir, fileName)
	output, err := os.Create(outName)
	if err != nil {
		return fmt.Errorf("creating '%s': %w", outName, err)
	}
	defer output.Close()

	encodedOutput := encodedWriter(output)

	for _, entry := range entries {
		str, err := entry.Engarde()
		if err != nil {
			return err
		}
		if _, err = encodedOutput.Write([]byte(str)); err != nil {
			return fmt.Errorf("writing to '%s': %w", outName, err)
		}
	}

	return nil
}
