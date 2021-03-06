package material

import (
	"github.com/dennwc/dom"
	"github.com/dennwc/dom/js"
	log "github.com/sirupsen/logrus"
)

const ChipTemplate = `
{{- define "ComponentTemplate" -}}
<div class="mdc-chip">
{{ if .LeadingIcon }}
<i class="material-icons mdc-chip__icon mdc-chip__icon--leading">{{ .LeadingIcon }}</i>
{{ end }}
<div class="mdc-chip__checkmark">
	<svg class="mdc-chip__checkmark-svg" viewBox="-2 -3 30 30">
		<path class="mdc-chip__checkmark-path" fill="none" stroke="black"
			d="M1.73,12.91 8.1,19.28 22.79,4.59" />
	</svg>
</div>
<div class="mdc-chip__text">{{.Text}}</div>
</div>
{{- end -}}
`

type ChipSpec struct {
	ComponentSpec
	LeadingIcon  string
	Text         string
	TrailingIcon string
}

func DefaultChipSpec() ChipSpec {
	return ChipSpec{
		ComponentSpec: ComponentSpec{
			Package:   "chips",
			Component: "MDCChip",
			Template:  ChipTemplate,
		},
	}
}

func NewChipSpec(opts ...ChipOption) ChipSpec {
	cs := DefaultChipSpec()
	for _, opt := range opts {
		opt(&cs)
	}
	return cs
}

type ChipOption func(cs *ChipSpec)

func FromChipSpec(ics *ChipSpec) ChipOption {
	return func(ocs *ChipSpec) {
		ocs.LeadingIcon = ics.LeadingIcon
	}
}

func ChipLeadingIcon(icon string) ChipOption {
	return func(cs *ChipSpec) {
		cs.LeadingIcon = icon
	}
}

func ChipText(text string) ChipOption {
	return func(cs *ChipSpec) {
		cs.Text = text
	}
}

// func Wrap(el dom.Element) ChipOption {
// 	return func(c *ChipSpec) {
// 		c.Element = el
// 	}
// }

type Chip struct {
	Component
}

func AsChip(c Component) (Chip, error) {
	return Chip{
		Component: c,
	}, nil
}

func NewChip(opts ...ChipOption) (Chip, error) {
	cmp, err := NewComponent(NewChipSpec(opts...))
	if err != nil {
		return Chip{}, err
	}
	return AsChip(cmp)
}

func DefaultChip(text string) Chip {
	cmp, _ := NewChip()
	return cmp
}

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

type ChipSet struct {
	Component
}

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
	// cso.BlockElement.SetAttribute("data-mdc-auto-init", "MDCChipSet")

	mdc := js.Get("mdc")
	pack := mdc.Get("chips")
	comp := pack.Get("MDCChipSet")
	// cs := comp.Call("attachTo", cso.BlockElement.JSValue())
	// cs.Call("initialize")
	cs := comp.New(cso.BlockElement.JSValue())
	//mdc.Call("autoInit", cso.BlockElement.JSValue())

	return ChipSet{
		Component: Component{
			Value: cs,
		},
	}
}

func (cs ChipSet) AddChip(c Chip) {
	if !c.Value.Valid() {
		return
	}

	c.Value.Call("destroy")
	//c.Value = *new(js.Value)
	cs.Root().AppendChild(c.Root())
	cs.Value.Call("addChip", c.Root().JSValue())
}

func (cs ChipSet) Chips() []Chip {
	var chips []Chip
	for _, v := range cs.Get("chips").Slice() {
		log.Info("Chip: ", v)
		log.Info("Chip selected: ", v.Get("selected").Bool())
		c := Chip{
			Component: Component{
				Value: v,
			},
		}
		el := c.Root()
		log.Info("Chip id: ", el.GetAttribute("id"))
		log.Info("Chip classes: ", el.GetAttribute("class"))
		chips = append(chips, c)
	}
	return chips //TODO
}

func (cs ChipSet) SelectedChipIds() []string {
	var ids []string
	for _, v := range cs.Get("selectedChipIds").Slice() {
		ids = append(ids, v.String())
	}
	return ids
}
