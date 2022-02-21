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
	cfg          *EngardeConfig
}

func (entry *entryCfg) buildTarget() {
	if entry.manualTarget {
		return
	}

	if header.name == "" {
		entry.target = ""
	} else {
		entry.target = fmt.Sprintf("%s_%s%s",
			header.name, entry.gender.ShortString(), entry.weapon.ShortString())
	}
}

func (entry *entryCfg) targetDone() bool {
	return entry.manualTarget || filepath.IsAbs(entry.target)
}

func (entry *entryCfg) fullTarget() string {
	if entry.targetDone() {
		return entry.target
	}
	return filepath.Join(header.targetDir, entry.target)
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

func Line(label string, widget ...g.Widget) GridLine {
	return GridLine{label, g.Row(widget...)}
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
				Size(comboSize),
				g.Labelf("DA auf %d Punkte", entry.ageGroup.KOPoints())),
			Line("Unterverzeichnis", g.InputText(&entry.target).OnChange(func() {
				if entry.target == "" {
					entry.manualTarget = false
					entry.buildTarget()
				} else {
					entry.manualTarget = true
				}
			}),
				g.Tooltip(`Wenn nur ein Name: Unterverzeichnis unter Zielverzeichnis.
Wenn ein Pfad: Zielverzeichnis wird ignoriert, vollständiger Pfad wird genommen.

Name wird autogeneriert. Autogenerierung ist inaktiv, sobald manuelle Änderungen vorgenommen werden.
Um die Autogenierung wieder zu aktivieren, einmal den Inhalt löschen.`),
				g.Button(chooseStr).OnClick(func() {
					dir, err := zenity.SelectFile(
						zenity.Directory(),
						zenity.Filename(filepath.Join(header.targetDir, entry.target)+string(filepath.Separator)))
					if err == nil && dir != "" {
						entry.target = dir
						entry.manualTarget = true
					}
				})),
			Line("Ophardt-Export",
				g.InputText(&entry.inputFile),
				g.Tooltip("Pfad zur CSV aus Ophardt."),
				g.Button(chooseStr).OnClick(func() {
					file, err := zenity.SelectFile(
						zenity.Filename(entry.inputFile),
						zenity.FileFilters{
							{Name: "CSV Files", Patterns: []string{"*.csv"}},
						})
					if err == nil && file != "" {
						entry.inputFile = file
					}
				})),
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
			}), g.Tooltip("Kurzname des Turniers (z.B. WH2022)")),
			Line("Beschreibung", g.InputText(&header.description),
				g.Tooltip("Langname / Beschreibung des Turniers (z.B. Weißherbst 2022)")),
			Line("Wettkampftag", g.DatePicker("##date", &header.date).
				Format("02.01.2006").StartOfWeek(time.Monday).
				Size(comboSize)),
			Line("Zielverzeichnis",
				g.InputText(&header.targetDir),
				g.Tooltip("Oberverzeichnis in dem für alle Konfigurationen ein Unterverzeichnis angelegt wird"),
				g.Button(chooseStr).OnClick(func() {
					dir, err := zenity.SelectFile(zenity.Directory(), zenity.Filename(header.targetDir+"/"))
					if err == nil && dir != "" {
						header.targetDir = dir
					}
				})),
		),
		entryBuilder(),
		g.Style().SetFont(icomoonFI).To(g.Row(
			g.Button("\ue900").OnClick(func() {
				entries = append(entries, newEntry())
			}),
			g.Button("\ue903").Disabled(len(entries) == 0).OnClick(func() {
				entries = entries[:len(entries)-1]
			}))),
		g.Align(g.AlignCenter).To(g.Row(
			g.Button("Import").OnClick(doImport).Disabled(dataMissing()),
			g.Button("Quit").OnClick(shouldQuit),
		)),
		g.PrepareMsgbox(),
	)
}

func dataMissing() bool {
	if header.name == "" {
		return true
	}
	for _, e := range entries {
		if e.inputFile == "" {
			return true
		}
		if !e.targetDone() && header.targetDir == "" {
			return true
		}
	}

	return false
}

func doImport() {
	for i, e := range entries {
		var err error
		cfg := EngardeConfig{
			InputFile:   e.inputFile,
			OutputDir:   e.fullTarget(),
			Name:        header.name,
			Description: header.description,
			Gender:      e.gender,
			AgeGroup:    e.ageGroup,
			Weapon:      e.weapon,
			Date:        header.date,
		}
		entries[i].cfg = &cfg

		if cfg.Participants, cfg.Clubs, err = parseOphardtInput(cfg.InputFile); err != nil {
			g.Msgbox("Fehler", fmt.Sprintf("Beim Import von %s ist ein Fehler aufgetreten:\n%v",
				cfg.InputFile, err)).Buttons(g.MsgboxButtonsOk)

			return
		}
	}

	for _, e := range entries {
		if err := Write(*e.cfg); err != nil {
			g.Msgbox("Fehler",
				fmt.Sprintf("Beim Schreiben der Daten für Verzeichnis %s ist ein Fehler aufgetreten\n%v",
					e.target, err))
		}
	}

	g.Msgbox("Import erfolgreich", "Programm schließen?").Buttons(g.MsgboxButtonsYesNo).
		ResultCallback(func(result g.DialogResult) {
			if result {
				shouldQuit()
			}
		})
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
