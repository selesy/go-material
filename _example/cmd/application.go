package main

import (
	"github.com/selesy/go-material-components-web/material"
	"honnef.co/go/js/dom"
)

func main() {
	d := dom.GetWindow().Document()
	t := material.NewTopAppBar()
	d.AppendChild(t)
}
