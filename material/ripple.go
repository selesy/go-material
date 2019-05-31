package material

type rippleColor string

const (
	RippleDefaultColor rippleColor = ""
	RipplePrimaryColor rippleColor = "mdc-ripple-surface--primary"
	RippleAccentColor  rippleColor = "mdc-ripple-surface--accent"
)

type RippleSpec struct {
	ComponentSpec
	Unbounded bool
	Color     rippleColor
}

func DefaultRippleSpec() RippleSpec {
	return RippleSpec{
		ComponentSpec: ComponentSpec{
			Package:   "ripple",
			Component: "MDCRipple",
			Template:  "",
		},
	}
}

func NewRippleSpec(opts ...RippleOption) RippleSpec {
	rs := DefaultRippleSpec()
	for _, opt := range opts {
		opt(&rs)
	}
	return rs
}

type RippleOption func(rs *RippleSpec)

func RippleColor(rc rippleColor) RippleOption {
	return func(rs *RippleSpec) {
		rs.Color = rc
	}
}

func UnboundedRipple(unbounded bool) RippleOption {
	return func(rs *RippleSpec) {
		rs.Unbounded = unbounded
	}
}

type Ripple struct {
	Component
}

func AsRipple(c Component) Ripple {
	return Ripple{
		Component: c,
	}
}

func NewRipple(opts ...RippleOption) Ripple {
	cmp, err := NewComponent(NewRippleSpec(opts...))
	if err != nil {

	}
	return AsRipple(cmp)
}

func (r Ripple) Unbounded() bool {
	return r.Get("unbounded").Bool()
}

func (r Ripple) SetUnbounded(unbounded bool) {
	r.Set("unbounded", unbounded)
}
