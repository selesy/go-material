//+build wasm

package main

import (
	"github.com/dennwc/dom"
	"github.com/dennwc/dom/js"
	"github.com/dennwc/dom/require"
	"github.com/selesy/go-material/material"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Trace("-> main()")
	log.Info("Starting Go Material Catalog")

	require.Stylesheet("https://unpkg.com/material-components-web@latest/dist/material-components-web.min.css")
	require.Stylesheet("https://fonts.googleapis.com/icon?family=Material+Icons")
	require.Stylesheet("https://material-components.github.io/material-components-web-catalog/static/css/main.0729fb5b.css")
	require.Script("https://unpkg.com/material-components-web@latest/dist/material-components-web.min.js")

	theme := dom.Doc.CreateElement("meta")
	theme.SetAttribute("name", "theme-color")
	theme.SetAttribute("content", "#000000")
	dom.Head.AppendChild(theme)

	material.Title("Material Components Web - Catalog")

	t := material.NewTopAppBar()
	dom.Body.AppendChild(t)

	s := dom.Doc.CreateElement("section")
	s.SetAttribute("class", "mdc-top-app-bar--fixed-adjust")
	dom.Body.AppendChild(s)

	chipSet := material.NewChipSet(material.ChipSetVariant(material.FilterChipSet))
	s.AppendChild(&chipSet.Element)

	chip1 := material.DefaultChip("Chip 1")
	chipSet.AddChip(chip1)
	chip2 := material.DefaultChip("Chip 2")
	chipSet.AddChip(chip2)
	//	chip2.SetSelected(true)

	log.Info("Selected chips: ", chipSet.SelectedChips())

	chip3 := material.DefaultChip(("Chip 3"))
	s.AppendChild(&chip3.Element)

	mdc := js.Get("mdc")
	auto := mdc.Call("autoInit")
	log.Info("Auto valid: ", auto.Valid())

	log.Info("Exiting Go Material Catalog")
	log.Trace("main() ->")
}
