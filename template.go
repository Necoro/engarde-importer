package main

import (
	"embed"
)

//go:embed templates/verbatim
var verbatim embed.FS

func writeVerbatim() error {
	entries, err := verbatim.ReadDir("templates/verbatim")
	if err != nil {
		return err
	}

	for _, e := range entries {
		print(e.Name())
	}

	return nil
}
