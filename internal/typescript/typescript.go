package typescript

import (
	"strings"
)

type AccessModifier string

const (
	PrivateAccess   AccessModifier = "private"
	ProtectedAccess AccessModifier = "protected"
	PublicAccess    AccessModifier = "public"
	ReadOnlyAccess  AccessModifier = "readonly"
)

func IsAccessModifier(token *Token) (AccessModifier, bool) {
	key := strings.ToLower(string(token.Bytes))

	accessModifier, ok := map[string]AccessModifier{
		"private":   PrivateAccess,
		"protected": ProtectedAccess,
		"public":    PublicAccess,
		"readonly":  ReadOnlyAccess,
	}[key]

	return accessModifier, ok
}

type AccessorModifier string

const (
	AccessorAccessorModifier AccessorModifier = "get"
	MutatorAccessorModifier  AccessorModifier = "set"
)

func IsAccessorModifier(token *Token) (AccessorModifier, bool) {
	key := strings.ToLower(string(token.Bytes))

	accessorModifier, ok := map[string]AccessorModifier{
		"get": AccessorAccessorModifier,
		"set": MutatorAccessorModifier,
	}[key]

	return accessorModifier, ok
}

type MutabilityModifier string

const (
	ReadOnlyMutability  MutabilityModifier = "readonly"
	ReadWriteMutability MutabilityModifier = "readwrite"
)

func IsMutabilityModifier(token *Token) (MutabilityModifier, bool) {
	key := strings.ToLower(string(token.Bytes))

	mutabilityModifier, ok := map[string]MutabilityModifier{
		"readonly":  ReadOnlyMutability,
		"readwrite": ReadWriteMutability,
	}[key]

	return mutabilityModifier, ok
}

type ClassObjectModifier string

const (
	StaticModifier ClassObjectModifier = "static"
)

func IsClassObjectModifier(token *Token) (ClassObjectModifier, bool) {
	key := strings.ToLower(string(token.Bytes))

	classObjectModifier, ok := map[string]ClassObjectModifier{
		"static": StaticModifier,
	}[key]

	return classObjectModifier, ok
}

func IsModifier(token *Token) bool {
	_, aok := IsAccessModifier(token)
	_, aaok := IsAccessorModifier(token)
	_, mok := IsMutabilityModifier(token)
	_, sok := IsClassObjectModifier(token)

	return aok || aaok || mok || sok
}

type Argument struct {
	Name string
	Type string
}

type Class struct {
	Name    string
	Extends string
	Fields  []Field
	Methods []Method
}

func (c *Class) String() string {
	sb := &strings.Builder{}
	sb.WriteRune('\n')
	c.writeValue(sb, "", "Name", c.Name)
	c.writeValue(sb, "", "Extends", c.Extends)

	c.writeValue(sb, "", "Fields", "")

	for _, field := range c.Fields {
		c.writeValue(sb, "\t", "Name", field.Name)
	}

	c.writeValue(sb, "", "Methods", "")

	for _, method := range c.Methods {
		c.writeValue(sb, "\t", "Name", method.Name)
		c.writeValue(sb, "\t\t", "Modifiers", "")

		for _, modifier := range method.Modifiers {
			c.writeValue(sb, "\t\t\t", "Modifier", string(modifier.Bytes))
		}

		c.writeValue(sb, "\t\t", "Arguments", "")

		for _, argument := range method.Arguments {
			c.writeValue(sb, "\t\t\t", "Argument", "")
			c.writeValue(sb, "\t\t\t\t", "Name", argument.Name)
			c.writeValue(sb, "\t\t\t\t", "Type", argument.Type)
		}

		c.writeValue(sb, "\t\t", "Return type", method.Type)
	}

	return sb.String()
}

func (*Class) writeValue(sb *strings.Builder, prefix, name, value string) {
	sb.WriteString(prefix)
	sb.WriteString(name)
	sb.WriteString(": ")
	sb.WriteString(value)
	sb.WriteRune('\n')
}

type Field struct {
	Modifiers []*Token
	Name      string
	Type      string
}

type Method struct {
	Modifiers []*Token
	Name      string
	Arguments []Argument
	Type      string
}

type Parser struct {
	Lexer   *Lexer
	Classes []*Class
}

func Parse(lexer *Lexer) []*Class {
	parser := &Parser{
		Lexer:   lexer,
		Classes: []*Class{},
	}

	parser.ClassSearch()

	return parser.Classes
}

func (p *Parser) ClassBlock(class *Class) {
	token := p.Token()
	for string(token.Bytes) != "}" {
		modifiers := []*Token{}
		for IsModifier(token) {
			modifiers = append(modifiers, token)
			token = p.Token()
		}

		name := string(token.Bytes)

		// TODO: should get the field type here?
		token = p.Token()
		for string(token.Bytes) != "(" && string(token.Bytes) != ";" {
			token = p.Token()
		}

		if string(token.Bytes) == ";" {
			class.Fields = append(class.Fields, Field{
				Modifiers: modifiers,
				Name:      name,
				Type:      "",
			})

			token = p.Token()

			continue
		}

		arguments := []Argument{}
		if string(token.Bytes) == "(" {
			arguments = p.MethodArguments()
			token = p.Token()
		}

		returnType := ""

		if string(token.Bytes) == ":" {
			token = p.Token()
			returnType = string(token.Bytes)
			token = p.Token()
		}

		class.Methods = append(class.Methods, Method{
			Modifiers: modifiers,
			Name:      name,
			Arguments: arguments,
			Type:      returnType,
		})

		// This is a bit brittle (is there stuff after the return type?)
		if string(token.Bytes) == "{" {
			p.SkipBlock()
		}

		token = p.Token()
	}
}

func (p *Parser) ClassDefinition(token *Token) *Token {
	name := p.Token()
	next := p.Token()

	var extends *Token = nil
	if string(next.Bytes) == "extends" {
		extends = p.Token()
		next = p.Token()
	}

	if string(next.Bytes) == "<" {
		next = p.SkipGeneric(next)
	}

	if string(next.Bytes) == "{" {
		class := &Class{
			Name:    string(name.Bytes),
			Extends: string(extends.Bytes),
			Fields:  []Field{},
			Methods: []Method{},
		}
		p.ClassBlock(class)
		p.Classes = append(p.Classes, class)
	}

	return token
}

func (p *Parser) ClassSearch() {
	token := &Token{
		Row:    0,
		Column: 0,
		Bytes:  []byte{},
		EOL:    false,
		EOF:    false,
	}

	for !token.EOF {
		token = p.Token()
		token = p.SkipComment(token)

		if string(token.Bytes) == "class" {
			token = p.ClassDefinition(token)
		}
	}
}

func (p *Parser) SkipBlock() {
	nest := 1

	token := &Token{}
	for string(token.Bytes) != "}" || nest != 0 {
		token = p.Token()

		if string(token.Bytes) == "}" {
			nest--
		}

		if string(token.Bytes) == "{" {
			nest++
		}
	}
}

func (p *Parser) SkipComment(token *Token) *Token {
	if strings.HasPrefix(string(token.Bytes), "/*") {
		token = p.SkipBlockComment(token)
	}

	if strings.HasPrefix(string(token.Bytes), "//") {
		token = p.SkipLineComment(token)
	}

	return token
}

func (p *Parser) SkipBlockComment(token *Token) *Token {
	for string(token.Bytes) != "*/" {
		token = p.Token()
	}

	return p.SkipComment(p.Token())
}

func (p *Parser) SkipGeneric(token *Token) *Token {
	for string(token.Bytes) != ">" {
		token = p.Token()
	}

	return p.Token()
}

func (p *Parser) SkipLineComment(token *Token) *Token {
	for !token.EOL {
		token = p.Token()
	}

	return p.SkipComment(p.Token())
}
