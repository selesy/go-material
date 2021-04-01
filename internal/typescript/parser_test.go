package typescript_test

import (
	"io/fs"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/selesy/go-material/internal/typescript"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMethodArguments(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name      string
		snippet   string
		arguments []typescript.Argument
	}{
		{"no arguments", ")", []typescript.Argument{}},
		{"untyped argument", "root)", []typescript.Argument{{Name: "root"}}},
		{"typed argment", "root: Element)", []typescript.Argument{{Name: "root", Type: "Element"}}},
		{"lamba argument", "actionFactory: MDCChipActionFactory = (el: Element) => new MDCChipAction(el))", []typescript.Argument{{Name: "actionFactory", Type: "MDCChipActionFactory"}}},
		{"two arguments", "first: Element, second: Element)", []typescript.Argument{{Name: "first", Type: "Element"}, {Name: "second", Type: "Element"}}},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			parser := LexerParser(t, testcase.snippet)
			arguments := parser.MethodArguments()
			assert.Equal(t, testcase.arguments, arguments)
		})
	}
}

func TestSkipMethodLambdaType(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name    string
		snippet string
		next    string
	}{
		{"commas means next argument", ", next", "next"},
		{"right parenthesis means end of arguments", ")", ")"},
		{"lambda should be skipped as last argument", "= (el: Element) => new MDCChipAction(el))", ")"},
		{"lambda should be skipped for next argument", "= (el: Element) => new MDCChipAction(el), next", "next"},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			parser := LexerParser(t, testcase.snippet)
			token := parser.SkipMethodLambdaType()
			assert.Equal(t, testcase.next, string(token.Bytes))
		})
	}
}

func TestToken(t *testing.T) {
	t.Parallel()

	const (
		inputFile     = "testdata/chip.ts"
		tokenizedFile = "testdata/chip_tokenized.golden"
	)

	bytes, err := ioutil.ReadFile(inputFile)
	require.NoError(t, err)
	require.NotEmpty(t, bytes)

	parser := LexerParser(t, string(bytes))

	sb, token := strings.Builder{}, parser.Token()
	for !token.EOF {
		sb.WriteString(token.String())
		sb.WriteByte('\n')

		token = parser.Token()
	}

	if *update {
		err := ioutil.WriteFile(tokenizedFile, []byte(sb.String()), fs.ModePerm)
		require.NoError(t, err)

		return
	}

	exp, err := ioutil.ReadFile(tokenizedFile)
	require.NoError(t, err)

	require.Equal(t, string(exp), sb.String())
}

func LexerParser(t *testing.T, snippet string) *typescript.Parser {
	t.Helper()

	lexer := typescript.NewLexer([]byte(snippet))
	parser := &typescript.Parser{
		Lexer:   lexer,
		Classes: []*typescript.Class{},
	}

	return parser
}
