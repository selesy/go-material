//+build wasm,js

package material

import (
	"github.com/dennwc/dom"
	"github.com/dennwc/dom/js"
)

type ComponentOption func(c *Component)

type ComponentSpec struct {
	Package   string
	Component string
	Class     string
}

type Component struct {
	dom.HTMLElement
	js.Value
}

type Variant string

func (c Component) AttachTo(dom.HTMLElement) {

}

type Rippler interface {
	GetRipple() Ripple
}

type Selector interface {
	Selected() bool
	SetSelected(selected bool)
}
