package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"
	"text/template"
)

const templatePath = "templates/"
const verbatimPath = templatePath + "verbatim"
const templateGlob = templatePath + "*.tpl"

//go:embed templates
var templates embed.FS

func writeVerbatimFile(outputDir string, entry fs.DirEntry) error {
	inName := path.Join(verbatimPath, entry.Name())

	input, err := templates.Open(inName)
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
	entries, err := templates.ReadDir(verbatimPath)
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

	if err := writeTemplates(config.outputDir, config); err != nil {
		return fmt.Errorf("writing template files: %w", err)
	}

	return nil
}

var funcMap = template.FuncMap{
	"upper": strings.ToUpper,
}

func writeTemplates(outputDir string, config EngardeConfig) error {
	tpls, err := template.New("root").Funcs(funcMap).ParseFS(templates, templateGlob)
	if err != nil {
		return fmt.Errorf("parsing templates: %w", err)
	}

	for _, tpl := range tpls.Templates() {
		if err = writeTemplate(tpl, outputDir, config); err != nil {
			return err
		}
	}

	return nil
}

func writeTemplate(tpl *template.Template, outputDir string, config EngardeConfig) error {
	resultName := strings.TrimSuffix(tpl.Name(), ".tpl")
	outName := path.Join(outputDir, resultName)
	output, err := os.Create(outName)
	if err != nil {
		return fmt.Errorf("creating '%s': %w", outName, err)
	}
	defer output.Close()

	encodedOutput := encodedWriter(output)

	if err = tpl.Execute(encodedOutput, config); err != nil {
		return fmt.Errorf("executing template '%s': %w", tpl.Name(), err)
	}
	return nil
}
