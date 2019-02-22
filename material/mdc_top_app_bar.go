package material

import (
	"github.com/dennwc/dom"
	"github.com/selesy/go-bem/bem"
)

var row = `
<div class="mdc-top-app-bar__row">
<section class="mdc-top-app-bar__section mdc-top-app-bar__section--align-start">
  <a href="#" class="material-icons mdc-top-app-bar__navigation-icon" style="font-family: "Material Icons";">menu</a>
  <span class="mdc-top-app-bar__title">Material Components for the Web</span>
</section>
</div>
`

var topAppBar = bem.Block{
	Tag:  "header",
	Name: "mdc-top-app-bar",
	Elements: []bem.Element{
		bem.Element{
			Tag:  "div",
			Name: "row",
			Elements: []bem.Element{
				bem.Element{
					Tag:  "section",
					Name: "section",
					Modifiers: []bem.Modifier{
						bem.Modifier{
							Name: "align-start",
						},
						bem.Modifier{
							Name: "align-end",
						},
					},
				},
			},
		},
	},
	Modifiers: []bem.Modifier{
		bem.Modifier{
			Name: "short",
			Modifiers: []bem.Modifier{
				bem.Modifier{
					Name: "collapsed",
				},
			},
		},
		bem.Modifier{
			Name: "Fixed",
		},
		bem.Modifier{
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
