package material

import (
	"github.com/dennwc/dom"
	"github.com/dennwc/dom/js"
	log "github.com/sirupsen/logrus"
)

type ChipOption func(c *Chip)

func Wrap(el dom.Element) ChipOption {
	return func(c *Chip) {
		c.Element.Element = el
	}
}

type Chip Component

func DefaultChip(text string) Chip {
	el := dom.Doc.CreateElement("div")
	el.SetAttribute("class", "mdc-chip")
	el.SetInnerHTML(text)
	mdc := js.Get("mdc")
	pkg := mdc.Get("chips")
	comp := pkg.Get("MDCChip")
	c := comp.Call("attachTo", el)
	return Chip{
		Element: dom.HTMLElement{
			Element: *el,
		},
		Value: c,
	}
}

// func NewChip(opts ...ChipOption) Chip {
// 	var c Chip
// 	for _, opt := range opts {
// 		opt(&c)
// 	}
// 	e := dom.Doc.CreateElement("div")
// 	log.Info("Got here 1")
// 	c := js.Class("mdc.chips.MDCChip")
// 	log.Info("Instance of class", c.InstanceOfClass("mdc.chips.MDCChip"))
// 	log.Info("Got here 2")
// 	log.Info("Class is null: ", c.IsNull())
// 	log.Info("Class is undefined: ", c.IsUndefined())
// 	log.Info("Class: ", c.String())
// 	cl := js.Get("mdc.chips.MDCChip")
// 	log.Info("Class is null: ", cl.IsNull())
// 	log.Info("Class is undefined: ", cl.IsUndefined())
// 	v1 := js.Call("mdc.chips.MDCChip.attachTo", e)
// 	log.Info("v1", v1)
// 	v := c.Call("attachTo", e)
// 	return Chip{
// 		Element: dom.HTMLElement{
// 			Element: *e,
// 		},
// 		Value: v,
// 	}
// }

func (c Chip) Selected() bool {
	return c.Value.Get("selected").Bool()
}

func (c Chip) SetSelected(selected bool) {
	c.Value.Set("selected", selected)
}

//
//ChipSet
//

const (
	chipSetDefaultElement string = "div"
	chipSetBlockClass     string = "mdc-chip-set"
)

type chipSetConfig struct {
	BlockElement  *dom.Element
	BlockClass    string
	ModifierClass string
}

func defaultChipSetConfig() chipSetConfig {
	return chipSetConfig{
		BlockClass: chipSetBlockClass,
	}
}

type chipSetVariant Variant

const (
	ChoiceChipSet chipSetVariant = "mdc-chip-set--choice"
	FilterChipSet chipSetVariant = "mdc-chip-set--filter"
	InputChipSet  chipSetVariant = "mdc-chip-set--input"
)

type ChipSetOption func(cso *chipSetConfig)

func ChipSetElement(el *dom.Element) ChipSetOption {
	return func(cso *chipSetConfig) {
		cso.BlockElement = el
	}
}

func ChipSetVariant(v chipSetVariant) ChipSetOption {
	return func(cso *chipSetConfig) {
		cso.ModifierClass = string(v)
	}
}

type ChipSet Component

func NewChipSet(opts ...ChipSetOption) ChipSet {
	log.Trace("NewChipSet")
	cso := defaultChipSetConfig()
	for _, opt := range opts {
		opt(&cso)
	}

	if cso.BlockElement == nil {
		cso.BlockElement = dom.Doc.CreateElement(chipSetDefaultElement)
	}

	cl := cso.BlockClass
	if cso.ModifierClass != "" {
		cl = cl + " " + cso.ModifierClass
	}
	cso.BlockElement.SetAttribute("class", cl)

	mdc := js.Get("mdc")
	pack := mdc.Get("chips")
	comp := pack.Get("MDCChipSet")
	cs := comp.Call("attachTo", cso.BlockElement)

	return ChipSet{
		Element: dom.HTMLElement{
			Element: *cso.BlockElement,
		},
		Value: cs,
	}
}

func (cs ChipSet) AddChip(c Chip) {
	log.Info("AddChip")
	log.Info("ChipSet is valid: ", cs.Value.Valid())
	log.Info("Chip is valid: ", c.Value.Valid())
	v := js.ValueOf(c.Element.Element)
	cs.Value.Call("addChip", v)
}

func (cs ChipSet) Chips() []Chip {
	return nil //TODO
}

func (cs ChipSet) SelectedChips() []Chip {
	return nil //TODO
}
