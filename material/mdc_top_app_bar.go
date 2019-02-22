package material

import (
	"github.com/dennwc/dom"
)

var row = `
<div class="mdc-top-app-bar__row">
<section class="mdc-top-app-bar__section mdc-top-app-bar__section--align-start">
  <a href="#" class="material-icons mdc-top-app-bar__navigation-icon" style="font-family: "Material Icons";">menu</a>
  <span class="mdc-top-app-bar__title">Material Components for the Web</span>
</section>
</div>
`

var topAppBar = Block{
	Tag:  "header",
	Name: "mdc-top-app-bar",
	Elements: []Element{
		Element{
			Tag:  "div",
			Name: "row",
			Elements: []Element{
				Element{
					Tag:  "section",
					Name: "section",
					Modifiers: []Modifier{
						Modifier{
							Name: "align-start",
						},
						Modifier{
							Name: "align-end",
						},
					},
				},
			},
		},
	},
	Modifiers: []Modifier{
		Modifier{
			Name: "short",
			Modifiers: []Modifier{
				Modifier{
					Name: "collapsed",
				},
			},
		},
		Modifier{
			Name: "Fixed",
		},
		Modifier{
			Name: "Prominent",
		},
	},
}

type topAppBarVariant string

const (
	SHORT topAppBarVariant = "short"
)

type TopAppBar struct {
	dom.Element
}

func NewTopAppBar() *TopAppBar {
	e := dom.Doc.CreateElement("header")
	e.SetAttribute("class", "mdc-top-app-bar")
	e.SetInnerHTML(row)
	t := &TopAppBar{*e}
	return t
}

type TopAppBarRow struct {
}

func (t *TopAppBar) NewTopAppBarRow() *TopAppBarRow {
	return nil
}
