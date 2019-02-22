package material

import "honnef.co/go/js/dom"

type TopAppBar struct {
	*dom.HTMLHeadingElement
}

func NewTopAppBar() *TopAppBar {
	e := dom.GetWindow().Document().CreateElement("header").(TopAppBar)
	return &e
}

type TopAppBarRow dom.HTMLDivElement

func (t *TopAppBar) NewTopAppBarRow() *TopAppBarRow {
	r := dom.GetWindow().Document().CreateElement("div").(TopAppBarRow)
	t.AppendChild(r)
	return &r
}
