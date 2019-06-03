//+build wasm,js

package material

import "github.com/dennwc/dom"

func titleEl() *dom.HTMLElement {
	titles := dom.Doc.GetElementsByTagName("title")
	if len(titles) > 0 {
		return titles[0].AsHTMLElement()
	}

	titleEl := dom.Doc.CreateElement("title")
	dom.Head.AppendChild(titleEl)
	return titleEl.AsHTMLElement()
}

func Title() string {
	return titleEl().InnerText()
}

func SetTitle(title string) {
	titleEl().SetInnerText(title)
}
