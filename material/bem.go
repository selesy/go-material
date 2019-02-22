package material

type Block struct {
	Tag       string
	Name      string
	Elements  []Element
	Modifiers []Modifier
}

type Element struct {
	Tag       string
	Name      string
	Elements  []Element
	Modifiers []Modifier
}

type Modifier struct {
	Name      string
	Modifiers []Modifier
}
