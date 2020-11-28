package generator

type Block struct {
	Class string
	Tag   string
}

type Element struct {
	Name      string
	Component Component
}

type Component struct {
	Name  string
	Block Block
	// Elements  []Element
	// Modifiers []string
}
