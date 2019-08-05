//+build wasm,js

package material

import (
	"errors"
	"reflect"
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

func NewComponent(spec interface{}) (Component, error) {
	cs, ok := reflect.ValueOf(spec).FieldByName("ComponentSpec").Interface().(ComponentSpec)
	if !ok {
		return Component{}, errors.New("blah")
	}

	comp := js.Class("mdc", cs.Package, cs.Component)

	if cs.Template != "" {
		el := dom.Doc.CreateElement("div")

		//TODO: Replace foo with a real name
		t, err := template.New("foo").Parse(cs.Template)
		if err != nil {
			return Component{}, err
		}

		var b strings.Builder
		log.Info("Spec: ", spec)
		err = t.ExecuteTemplate(&b, "ComponentTemplate", spec)
		if err != nil {
			return Component{}, err
		}

		el.SetInnerHTML(b.String())
		if len(el.ChildNodes()) != 1 {
			return Component{}, errors.New("Only a single element node should be returned by a component template")
		}
		el = el.ChildNodes()[0]
		comp = comp.Call("attachTo", el)
	}

	return AsComponent(comp), nil
}

type Variant string

func (c *Component) AttachTo(root *dom.Element) {
	*c = AsComponent(c.Call("attachTo", root))
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
