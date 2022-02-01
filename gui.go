package main

import (
	"time"

	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
)

var guiCfg EngardeConfig

type gridLayout struct {
	size    float32
	widgets []g.Widget
	labels  []*g.LabelWidget
}

type gridLine struct {
	label  string
	widget g.Widget
}

func Line(label string, widget g.Widget) gridLine {
	return gridLine{label, widget}
}

func Grid(lines ...gridLine) *gridLayout {
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

	return &gridLayout{size, widgets, labels}
}

func (grid *gridLayout) Build() {
	for i := range grid.labels {
		g.AlignTextToFramePadding()
		grid.labels[i].Build()
		imgui.SameLineV(grid.size, 0)
		grid.widgets[i].Build()
	}
}

func loop(w *g.MasterWindow) func() {
	return func() {
		g.SingleWindow().Layout(
			g.Align(g.AlignCenter).To(g.Label("Engarde Importer")),
			g.Spacing(),
			Grid(
				Line("Name", g.InputText(&guiCfg.Name)),
				Line("Beschreibung", g.InputText(&guiCfg.Description)),
				Line("Wettkampftag", g.DatePicker("##date", &guiCfg.Date).
					Format("02.01.2006").StartOfWeek(time.Monday).
					Size(120)),
			),
			g.Spacing(),
			g.Align(g.AlignCenter).To(g.Button("Quit").OnClick(w.Close)),
		)
	}
}

func gui() {
	guiCfg.Date = time.Now()
	w := g.NewMasterWindow("Engarde Importer", 400, 200, 0)
	w.Run(loop(w))
}
