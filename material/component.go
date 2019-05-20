//+build wasm,js

package material

import (
	"github.com/dennwc/dom"
	"github.com/dennwc/dom/js"
)

type Attacher interface {
	AttachTo(root dom.HTMLElement)
}

type HTMLElementer interface {
	HTMLElement() dom.HTMLElement
}

type Component struct {
	Element dom.HTMLElement
	Value   js.Value
}

type Variant string

func (c Component) AttachTo(e dom.HTMLElement) {

}

type Rippler interface {
	GetRipple() Ripple
}

type Selector interface {
	Selected() bool
	SetSelected(selected bool)
}
