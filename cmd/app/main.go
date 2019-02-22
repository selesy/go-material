// build wasm

package main

import (
	"github.com/dennwc/dom"
	"github.com/selesy/go-material-components-web/material"
)

func main() {
	dom.Require("https://unpkg.com/material-components-web@latest/dist/material-components-web.min.css")
	dom.Require("https://fonts.googleapis.com/icon?family=Material+Icons")
	dom.Require("https://unpkg.com/material-components-web@latest/dist/material-components-web.min.js")

	t := material.NewTopAppBar()
	dom.Body.AppendChild(t)
}
