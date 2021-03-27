package typescript_test

import (
	"flag"
	"io/fs"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/selesy/go-material/internal/typescript"
	"github.com/stretchr/testify/require"
)

func TestLexer(t *testing.T) {
	t.Parallel()

	update := flag.Bool("update", false, "update golden files")

	if !flag.Parsed() {
		flag.Parse()
	}

	bytes, err := ioutil.ReadFile("testdata/chip.ts")
	require.NoError(t, err)
	require.NotEmpty(t, bytes)

	lexer := typescript.NewLexer(bytes)
	require.NotNil(t, lexer)

	sb, token := strings.Builder{}, lexer.Token()
	for !token.EOF {
		sb.WriteString(token.String())
		sb.WriteByte('\n')

		token = lexer.Token()
	}

	if *update {
		err := ioutil.WriteFile("testdata/chip.golden", []byte(sb.String()), fs.ModePerm)
		require.NoError(t, err)

		return
	}

	exp, err := ioutil.ReadFile("testdata/chip.golden")
	require.NoError(t, err)

	require.Equal(t, exp, []byte(sb.String()))
}

func TestTypescriptLexingParsing(t *testing.T) {
	t.Parallel()
}
