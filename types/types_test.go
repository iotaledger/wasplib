// +build feature_types

package types

import (
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGoCoreSchema(t *testing.T){
	t.SkipNow()
	err := GenerateGoCoreSchema()
	require.NoError(t, err)
}

func TestRustCoreSchema(t *testing.T){
	t.SkipNow()
	err := GenerateRustCoreSchema()
	require.NoError(t, err)
}

func TestGoTypes(t *testing.T) {
	t.SkipNow()
	err := filepath.Walk("../contracts",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, "\\schema.json") {
				return GenerateGoTypes(path)
			}
			return nil
		})
	require.NoError(t, err)
}

func TestJavaTypes(t *testing.T) {
	t.SkipNow()
	err := filepath.Walk("../contracts",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, "\\types.json") {
				return GenerateJavaTypes(path)
			}
			return nil
		})
	require.NoError(t, err)
}

func TestRustTypes(t *testing.T) {
	t.SkipNow()
	err := filepath.Walk("../rust/contracts",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, "\\types.json") {
				return GenerateRustTypes(path)
			}
			return nil
		})
	require.NoError(t, err)
}

func TestRustToGo(t *testing.T) {
	t.SkipNow()
	err := RustConvertor(RustToGoLine, "../contracts/$1/lib.go")
	require.NoError(t, err)
}

func TestRustToJava(t *testing.T) {
	t.SkipNow()
	err := RustConvertor(RustToJavaLine, "../java/src/org/iota/wasplib/contracts/$1/lib.java")
	require.NoError(t, err)
}
