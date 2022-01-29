package main

import (
	g "github.com/AllenDang/giu"
)

func loop(w *g.MasterWindow) func() {
	return func() {
		g.SingleWindow().Layout(
			g.Align(g.AlignCenter).To(g.Label("Hello")),
			g.Button("Quit").OnClick(w.Close),
		)
	}
}

func gui() {
	w := g.NewMasterWindow("Engarde Importer", 400, 200, 0)
	w.Run(loop(w))
}
