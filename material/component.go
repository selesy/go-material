//+build wasm,js

package material

import (
	"strings"
	"sync"
	"text/template"

	"github.com/dennwc/dom"
	"github.com/dennwc/dom/js"
	log "github.com/sirupsen/logrus"
)

//TODO: Remove when https://github.com/dennwc/dom/issues/40 is fixed
var (
	mu    sync.RWMutex
	comps = make(map[string]js.Value)
)

func comp(pkg string, cmp string) js.Value {
	name := pkg + "." + cmp
	mu.RLock()
	v, ok := comps[name]
	mu.RUnlock()
	if !ok {
		v = js.Get(cmp, "mdc", pkg)
		mu.Lock()
		comps[name] = v
		mu.Unlock()
	}
	return v
}

type ComponentSpec struct {
	Package   string
	Component string
	Classes   []string
	Template  string
}

type Component struct {
	js.Value
}

func AsComponent(v js.Value) Component {
	return Component{
		Value: v,
	}
}

func NewComponent(spec ChipSpec) (Component, error) {
	cs := spec.ComponentSpec
	el := dom.Doc.CreateElement("div")
	t, err := template.New("foo").Parse(cs.Template)
	if err != nil {
		return Component{}, err
	}

	var b strings.Builder
	err = t.ExecuteTemplate(&b, "ChipTemplate", spec)
	if err != nil {
		return Component{}, err
	}
	log.Info("Outer HTML: ", b.String())
	log.Info("el outer HTML before: ", el.OuterHTML())
	el.SetInnerHTML(b.String())
	log.Info("el outer HTML after: ", el.OuterHTML())
	// log.Info("Return string: ", s)
	// el.SetAttribute("class", "mdc-chip")
	// el.SetInnerHTML(text)
	mdc := js.Get("mdc")
	pkg := mdc.Get("chips")
	comp := pkg.Get("MDCChip")
	el = el.ChildNodes()[0]
	// c := comp.Call("attachTo", el)
	c := comp.Call("attachTo", el)
	return AsComponent(c), nil
}

type Variant string

func (c Component) AttachTo(root dom.HTMLElement) {

}

func (c *Component) Root() *dom.HTMLElement {
	v := c.Get("root_")
	if !v.Valid() {
		return nil
	}
	return &dom.HTMLElement{
		Element: *dom.AsElement(v),
	}
}

//
//
//
//
//
//
//
//
//
//
//

type Rippler interface {
	GetRipple() Ripple
}

type Selector interface {
	Selected() bool
	SetSelected(selected bool)
}
