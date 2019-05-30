//+build wasm

package main

import (
	"time"

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
	s.AppendChild(chipSet.Root())

	chipSpec := material.NewChipSpec(material.ChipLeadingIcon("golf_course"))

	chip1, err := material.NewChip(material.FromChipSpec(&chipSpec), material.ChipText("Chip 1"))
	if err == nil {
		chipSet.AddChip(chip1)
	}
	// chip1 := material.DefaultChip("Chip 1")
	// chipSet.AddChip(chip1)
	chip2 := material.DefaultChip("Chip 2")
	chipSet.AddChip(chip2)
	chip3 := material.DefaultChip("Chip 3")
	chipSet.AddChip(chip3)

	chip4, err := material.NewChip(material.ChipLeadingIcon("face"), material.ChipText("Chip 4"))
	if err == nil {
		chipSet.AddChip(chip4)
		for i, e := range dom.Doc.GetElementById("mdc-chip-4").ChildNodes() {
			log.Info("Child index: ", i, ", name: ", e.NodeName(), ", type: ", e.JSValue().Get("nodeType"))
		}
	}

	chipSet.Chips()

	chipSet.Root().OnClick(func(_ *dom.MouseEvent) {
		log.Info("Selected chips: ", chipSet.SelectedChipIds())
	})

	rippleFactory := js.Get("mdc", "ripple", "MDCRipple")

	boundedRippleEl := dom.Doc.CreateElement("div")
	boundedRippleEl.SetAttribute("class", "mdc-ripple-surface")
	boundedRippleEl.SetAttribute("style", "height: 200px;width: 200px;border: black solid 1px")
	s.AppendChild(boundedRippleEl)
	rippleFactory.Call("attachTo", boundedRippleEl)

	unboundedRippleEl := dom.Doc.CreateElement("div")
	unboundedRippleEl.SetAttribute("class", "mdc-ripple-surface")
	unboundedRippleEl.SetAttribute("style", "height: 200px;width: 200px;border: black solid 1px")
	s.AppendChild(unboundedRippleEl)
	unboundedRipple := rippleFactory.New(unboundedRippleEl)
	unboundedRipple.Set("unbounded", true)

	rippleEl := dom.Doc.CreateElement("div")
	rippleEl.SetAttribute("class", "mdc-ripple-surface")
	rippleEl.SetAttribute("style", "height: 200px;width: 200px;border: black solid 1px")
	s.AppendChild(rippleEl)
	ripple := material.NewRipple()
	ripple.AttachTo(rippleEl)

	// chip3 := material.DefaultChip(("Chip 3"))
	// s.AppendChild(&chip3.Element)
	// chip3.Element.SetAttribute("data-mdc-auto-init", "MDCChip")

	// mdc := js.Get("mdc")
	// //auto := mdc.Call("autoInit", chip3.Element.JSValue())
	// auto := mdc.Call("autoInit")
	// log.Info("Auto valid: ", auto.Valid())

	const interval = time.Millisecond * 1000
	for {
		time.Sleep(interval)
		// for _, c := range chipSet.Chips() {
		// 	c.SetSelected(!c.Selected())
		// }
	}

	log.Info("Exiting Go Material Catalog")
	log.Trace("main() ->")
}
