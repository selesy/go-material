// build wasm

package main

import (
	"github.com/dennwc/dom"
	"github.com/selesy/go-material-components-web/material"
)

func main() {
	dom.Require("https://unpkg.com/material-components-web@latest/dist/material-components-web.min.css")
	dom.Require("https://fonts.googleapis.com/icon?family=Material+Icons#.css")
	dom.Require("https://material-components.github.io/material-components-web-catalog/static/css/main.0729fb5b.css")
	dom.Require("https://unpkg.com/material-components-web@latest/dist/material-components-web.min.js")

	theme := dom.Doc.CreateElement("meta")
	theme.SetAttribute("name", "theme-color")
	theme.SetAttribute("content", "#000000")
	dom.Head.AppendChild(theme)

	title := dom.Doc.GetElementsByTagName("title")[0]
	title.SetInnerHTML("Material Components Web - Catalog")

	t := material.NewTopAppBar()
	dom.Body.AppendChild(t)
}
