// +build feature_types

package generator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

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
