package typescript_test

import (
	"io/fs"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/selesy/go-material/internal/typescript"
	"github.com/stretchr/testify/require"
)

func TestLexer(t *testing.T) {
	t.Parallel()

	const (
		inputFile = "testdata/chip_master.ts"
		lexedFile = "testdata/chip_lexed.golden"
	)

	bytes, err := ioutil.ReadFile(inputFile)
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
		err := ioutil.WriteFile(lexedFile, []byte(sb.String()), fs.ModePerm)
		require.NoError(t, err)

		return
	}

	exp, err := ioutil.ReadFile(lexedFile)
	require.NoError(t, err)

	require.Equal(t, exp, []byte(sb.String()))
}
