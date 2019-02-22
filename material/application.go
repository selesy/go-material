package material

import "github.com/dennwc/dom"

func Title(title string) {
	titles := dom.Doc.GetElementsByTagName("title")

	if len(titles) < 1 {
		dom.Head.AppendChild(dom.Doc.CreateElement("title"))
		titles = dom.Doc.GetElementsByTagName("title")
	}

	e := titles[0]
	e.SetInnerHTML(title)
}
