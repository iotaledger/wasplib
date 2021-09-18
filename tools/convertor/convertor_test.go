// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// +build feature_types

package convertor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRustToGo(t *testing.T) {
	t.SkipNow()
	err := RustConvertor(RustToGoLine, "converted/$c/$c.go")
	require.NoError(t, err)
}

func TestRustToJava(t *testing.T) {
	t.SkipNow()
	err := RustConvertor(RustToJavaLine, "converted/$c/$C.java")
	require.NoError(t, err)
}
