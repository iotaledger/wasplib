// +build feature_types

package types

import (
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateSchemas(t *testing.T) {
	t.SkipNow()
	err := filepath.Walk("../contracts",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, "\\schema.json") {
				return GenerateSchema(path)
			}
			return nil
		})
	require.NoError(t, err)
}

func TestGenerateCoreContractsSchema(t *testing.T) {
	t.SkipNow()
	err := GenerateGoCoreContractsSchema()
	require.NoError(t, err)
	err = GenerateRustCoreContractsSchema()
	require.NoError(t, err)
}

func TestRustToGo(t *testing.T) {
	t.SkipNow()
	err := RustConvertor(RustToGoLine, "../../contracts/$1/$1.go")
	require.NoError(t, err)
}

func TestRustToJava(t *testing.T) {
	t.SkipNow()
	err := RustConvertor(RustToJavaLine, "../../contracts/$1/$1.java")
	require.NoError(t, err)
}
