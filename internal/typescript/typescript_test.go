package typescript_test

import (
	"flag"
	"io/fs"
	"io/ioutil"
	"testing"

	"github.com/selesy/go-material/internal/typescript"
	"github.com/stretchr/testify/require"
)

var update = flag.Bool("update", false, "update golden files")

func TestParser(t *testing.T) {
	t.Parallel()

	testcases := []string{
		"chip_master",
		"chip_v10.0.0",
	}

	const (
		// inputFile  = "testdata/chip-master.ts"
		parsedFile          = "testdata/chip_parsed.golden"
		prefix              = "testdata/"
		typescriptExtension = ".ts"
		goldenExtension     = ".golden"
	)

	for _, filename := range testcases {
		filename := filename

		t.Run(filename, func(t *testing.T) {
			t.Parallel()

			inputFile := prefix + filename + typescriptExtension
			parsedFile := prefix + filename + goldenExtension

			bytes, err := ioutil.ReadFile(inputFile)
			require.NoError(t, err)
			require.NotEmpty(t, bytes)

			lexer := typescript.NewLexer(bytes)
			require.NotNil(t, lexer)

			classes := typescript.Parse(lexer)
			require.NotNil(t, classes)
			require.Len(t, classes, 1)

			if *update {
				err := ioutil.WriteFile(parsedFile, []byte(classes[0].String()), fs.ModePerm)
				require.NoError(t, err)

				return
			}

			exp, err := ioutil.ReadFile(parsedFile)
			require.NoError(t, err)

			require.Equal(t, string(exp), classes[0].String())
		})
	}
}
