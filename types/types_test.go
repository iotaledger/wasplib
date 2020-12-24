package types

import (
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestGoTypes(t *testing.T) {
	//t.SkipNow()
	err := filepath.Walk("../contracts",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, "\\types.json") {
				return GenerateGoTypes(path)
			}
			return nil
		})
	require.NoError(t, err)
}

func TestJavaTypes(t *testing.T) {
	//t.SkipNow()
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
	//t.SkipNow()
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
	err := filepath.Walk("../rust/contracts",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, "\\lib.rs") {
				var matchContract = regexp.MustCompile(".+\\W(\\w+)\\Wsrc\\W.+")
				contract := matchContract.ReplaceAllString(path, "$1")
				return RustToGo(path, contract)
			}
			return nil
		})
	require.NoError(t, err)
}

func TestRustToJava(t *testing.T) {
	t.SkipNow()
	err := filepath.Walk("../rust/contracts",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, "\\lib.rs") {
				var matchContract = regexp.MustCompile(".+\\W(\\w+)\\Wsrc\\W.+")
				contract := matchContract.ReplaceAllString(path, "$1")
				return RustToJava(path, contract)
			}
			return nil
		})
	require.NoError(t, err)
}
