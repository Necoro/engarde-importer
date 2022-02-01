package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"os"
	"strconv"
	"time"

	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
	"github.com/ncruces/zenity"
)

//go:generate goversioninfo -64 -o resource_amd64.syso res/versioninfo.json

//go:embed res/monitor_48.png
var icon []byte

//go:embed res/icomoon.ttf
var icomoon []byte
var icomoonFI *g.FontInfo

const comboSize = 120

type entryCfg struct {
	inputFile string
	outputDir string
	gender    Gender
	ageGroup  AgeGroup
	weapon    Weapon
}

var (
	header struct {
		name        string
		description string
		date        time.Time
		targetDir   string
	}
	entries []entryCfg
)

type GridLayout struct {
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

const gridPadding = 10

var gridSize float32 = gridPadding

func Grid(lines ...GridLine) *GridLayout {
	widgets := make([]g.Widget, len(lines))
	labels := make([]*g.LabelWidget, len(lines))

	for i, line := range lines {
		labelSize, _ := g.CalcTextSize(line.label)
		if labelSize+gridPadding > gridSize {
			gridSize = labelSize + gridPadding
		}

		labels[i] = g.Label(line.label)
		widgets[i] = line.widget
	}

	return &GridLayout{widgets, labels}
}

func (grid *GridLayout) Build() {
	for i := range grid.labels {
		g.AlignTextToFramePadding()
		grid.labels[i].Build()
		imgui.SameLineV(gridSize, 0)
		grid.widgets[i].Build()
	}
}

func Layout(widgets ...g.Widget) g.Layout {
	return widgets
}

func PushID(id string) g.Widget {
	return g.Custom(func() { imgui.PushID(id) })
}

func PopID() g.Widget {
	return g.Custom(imgui.PopID)
}

func buildEntry(idx int) g.Widget {
	entry := &entries[idx]

	return Layout(
		PushID(strconv.Itoa(idx)),
		g.Spacing(),
		g.Separator(),
		g.Spacing(),
		Grid(
			Line("Waffe", g.Combo("",
				entry.weapon.String(), WeaponStrings, (*int32)(&entry.weapon)).
				Size(comboSize)),
			Line("Geschlecht", g.Combo("",
				entry.gender.String(), GenderStrings, (*int32)(&entry.gender)).
				Size(comboSize)),
			Line("Altersklasse", g.Combo("",
				entry.ageGroup.String(), AgeGroupStrings, (*int32)(&entry.ageGroup)).
				Size(comboSize)),
			Line("Ophardt-Export", g.Row(
				g.InputText(&entry.inputFile),
				g.Button("Wähle...").OnClick(func() {
					file, err := zenity.SelectFile(zenity.FileFilters{
						{Name: "CSV Files", Patterns: []string{"*.csv"}},
					})
					if err == nil && file != "" {
						entry.inputFile = file
					}
				}))),
		),
		PopID())
}

func entryBuilder() g.Widget {
	const id = "entries"

	return g.Custom(func() {
		imgui.PushID(id)
		defer imgui.PopID()

		for i := range entries {
			buildEntry(i).Build()
		}
	})
}

func shouldQuit() {
	g.Context.GetPlatform().SetShouldStop(true)
}

func loop() {
	g.SingleWindow().Layout(
		g.Align(g.AlignCenter).To(g.Label("Engarde Importer")),
		g.Spacing(),
		Grid(
			Line("Name", g.InputText(&header.name)),
			Line("Beschreibung", g.InputText(&header.description)),
			Line("Wettkampftag", g.DatePicker("##date", &header.date).
				Format("02.01.2006").StartOfWeek(time.Monday).
				Size(comboSize)),
			Line("Zielverzeichnis", g.Row(
				g.InputText(&header.targetDir),
				g.Button("Wähle...").OnClick(func() {
					dir, err := zenity.SelectFile(zenity.Directory(), zenity.Filename(header.targetDir+"/"))
					if err == nil && dir != "" {
						header.targetDir = dir
					}
				}))),
		),
		entryBuilder(),
		g.Style().SetFont(icomoonFI).To(
			g.Button("\ue900").OnClick(func() {
				entries = append(entries, entryCfg{})
			})),
		g.Align(g.AlignCenter).To(g.Button("Quit").OnClick(shouldQuit)),
	)
}

func gui() {
	header.date = time.Now()
	entries = make([]entryCfg, 1)
	header.targetDir, _ = os.UserHomeDir()

	icomoonFI = g.AddFontFromBytes("icomoon", icomoon, 16)
	w := g.NewMasterWindow("Engarde Importer", 500, 400, 0)

	if img, _, err := image.Decode(bytes.NewReader(icon)); err == nil {
		w.SetIcon([]image.Image{img})
	}

	w.Run(loop)
}
