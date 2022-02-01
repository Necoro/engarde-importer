package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"time"

	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
	"github.com/ncruces/zenity"
)

//go:generate goversioninfo -64 -o resource_amd64.syso res/versioninfo.json

//go:embed res/monitor_48.png
var icon []byte

var (
	guiCfg EngardeConfig
)

type GridLayout struct {
	size    float32
	widgets []g.Widget
	labels  []*g.LabelWidget
}

type GridLine struct {
	label  string
	widget g.Widget
}

func Line(label string, widget g.Widget) GridLine {
	return GridLine{label, widget}
}

func Grid(lines ...GridLine) *GridLayout {
	var size float32
	widgets := make([]g.Widget, len(lines))
	labels := make([]*g.LabelWidget, len(lines))

	for i, line := range lines {
		labelSize, _ := g.CalcTextSize(line.label)
		if labelSize > size {
			size = labelSize
		}

		labels[i] = g.Label(line.label)
		widgets[i] = line.widget
	}

	// add a default padding
	size = size + 10

	return &GridLayout{size, widgets, labels}
}

func (grid *GridLayout) Build() {
	for i := range grid.labels {
		g.AlignTextToFramePadding()
		grid.labels[i].Build()
		imgui.SameLineV(grid.size, 0)
		grid.widgets[i].Build()
	}
}

func loop(w *g.MasterWindow) func() {
	const comboSize = 120

	return func() {
		g.SingleWindow().Layout(
			g.Align(g.AlignCenter).To(g.Label("Engarde Importer")),
			g.Spacing(),
			Grid(
				Line("Name", g.InputText(&guiCfg.Name)),
				Line("Beschreibung", g.InputText(&guiCfg.Description)),
				Line("Wettkampftag", g.DatePicker("##date", &guiCfg.Date).
					Format("02.01.2006").StartOfWeek(time.Monday).
					Size(comboSize)),
				Line("Altersklasse", g.Combo(
					"", guiCfg.AgeGroup.String(), AgeGroupStrings,
					(*int32)(&guiCfg.AgeGroup)).
					Size(comboSize)),
				Line("Waffe", g.Combo(
					"", guiCfg.Weapon.String(), WeaponStrings,
					(*int32)(&guiCfg.Weapon)).
					Size(comboSize)),
				Line("Ophardt-Export", g.Row(g.InputText(&guiCfg.inputFile), g.SmallButton("Wähle...").OnClick(func() {
					file, err := zenity.SelectFile(zenity.FileFilters{
						{Name: "CSV Files", Patterns: []string{"*.csv"}},
					})

					if err == nil && file != "" {
						guiCfg.inputFile = file
					}
				}))),
			),
			g.Spacing(),
			g.Align(g.AlignCenter).To(g.Button("Quit").OnClick(w.Close)),
		)
	}
}

func gui() {
	guiCfg.Date = time.Now()
	w := g.NewMasterWindow("Engarde Importer", 500, 200, 0)
	if img, _, err := image.Decode(bytes.NewReader(icon)); err == nil {
		w.SetIcon([]image.Image{img})
	}

	w.Run(loop(w))
}
