package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	_ "image/png"
	"os"
	"path/filepath"
	"strconv"
	"time"

	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
	"github.com/ncruces/zenity"
)

//go:generate goversioninfo -64 -o resource_amd64.syso res/versioninfo.json

/*
 * Embedded Resources
 */

//go:embed res/monitor_48.png
var icon []byte

//go:embed res/icomoon.ttf
var icomoon []byte
var icomoonFI *g.FontInfo

/*
 * Data / State
 */

type entryCfg struct {
	inputFile    string
	target       string
	manualTarget bool
	gender       Gender
	ageGroup     AgeGroup
	weapon       Weapon
}

func (entry *entryCfg) buildTarget() {
	if entry.manualTarget {
		return
	}

	entry.target = fmt.Sprintf("%s_%s%s",
		header.name, entry.gender.ShortString(), entry.weapon.ShortString())
}

func newEntry() entryCfg {
	e := entryCfg{}
	e.buildTarget()
	return e
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

/*
 * Grid Layout
 */
// GridLayout is used to calculate the maximum width for labels,
// so that the following widgets all start at the same offset.
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

/*
 * Further additions to Giu
 */

func Layout(widgets ...g.Widget) g.Layout {
	return widgets
}

func PushID(id string) g.Widget {
	return g.Custom(func() { imgui.PushID(id) })
}

func PopID() g.Widget {
	return g.Custom(imgui.PopID)
}

func shouldQuit() {
	g.Context.GetPlatform().SetShouldStop(true)
}

/*
 * Putting together the GUI
 */
const (
	comboSize = 120
	chooseStr = "Wähle..."
)

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
				Size(comboSize).OnChange(entry.buildTarget)),
			Line("Geschlecht", g.Combo("",
				entry.gender.String(), GenderStrings, (*int32)(&entry.gender)).
				Size(comboSize).OnChange(entry.buildTarget)),
			Line("Altersklasse", g.Combo("",
				entry.ageGroup.String(), AgeGroupStrings, (*int32)(&entry.ageGroup)).
				Size(comboSize)),
			Line("Unterverzeichnis", g.Row(
				g.InputText(&entry.target).OnChange(func() {
					if entry.target == "" {
						entry.manualTarget = false
						entry.buildTarget()
					} else {
						entry.manualTarget = true
					}
				}),
				g.Button(chooseStr).OnClick(func() {
					dir, err := zenity.SelectFile(
						zenity.Directory(),
						zenity.Filename(filepath.Join(header.targetDir, entry.target)+string(filepath.Separator)))
					if err == nil && dir != "" {
						entry.target = dir
						entry.manualTarget = true
					}
				}))),
			Line("Ophardt-Export", g.Row(
				g.InputText(&entry.inputFile),
				g.Button(chooseStr).OnClick(func() {
					file, err := zenity.SelectFile(
						zenity.Filename(entry.inputFile),
						zenity.FileFilters{
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

func loop() {
	g.SingleWindow().Layout(
		g.Align(g.AlignCenter).To(g.Label("Engarde Importer")),
		g.Spacing(),
		Grid(
			Line("Name", g.InputText(&header.name).OnChange(func() {
				for i := range entries {
					entries[i].buildTarget()
				}
			})),
			Line("Beschreibung", g.InputText(&header.description)),
			Line("Wettkampftag", g.DatePicker("##date", &header.date).
				Format("02.01.2006").StartOfWeek(time.Monday).
				Size(comboSize)),
			Line("Zielverzeichnis", g.Row(
				g.InputText(&header.targetDir),
				g.Button(chooseStr).OnClick(func() {
					dir, err := zenity.SelectFile(zenity.Directory(), zenity.Filename(header.targetDir+"/"))
					if err == nil && dir != "" {
						header.targetDir = dir
					}
				}))),
		),
		entryBuilder(),
		g.Style().SetFont(icomoonFI).To(
			g.Button("\ue900").OnClick(func() {
				entries = append(entries, newEntry())
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
