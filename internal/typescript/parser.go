package typescript

func (p *Parser) MethodArguments() []Argument {
	arguments := []Argument{}

	token := p.Token()
	for string(token.Bytes) != ")" && !token.EOF {
		argument := Argument{
			Name: string(token.Bytes),
			Type: "",
		}

		token = p.Token()
		if string(token.Bytes) == "," || string(token.Bytes) == ")" {
			arguments = append(arguments, argument)

			continue
		}

		// if the current token isn't !: or : we should return an error
		token = p.Token()
		argument.Type = string(token.Bytes)
		arguments = append(arguments, argument)

		token = p.SkipMethodLambdaType()
	}

	return arguments
}

func (p *Parser) SkipMethodLambdaType() *Token {
	nest := 1

	token := &Token{}
	for string(token.Bytes) != ")" || nest != 0 { // exit ) && nest == 0 ---> continue !) || nest != 0
		token = p.Token()

		if string(token.Bytes) == "," && nest == 1 {
			token = p.Token()

			return token
		}

		if string(token.Bytes) == "(" {
			nest++
		}

		if string(token.Bytes) == ")" {
			nest--
		}
	}

	return token
}

func (p *Parser) Token() *Token {
	return p.SkipComment(p.Lexer.Token())
}
