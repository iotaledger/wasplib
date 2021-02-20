// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// +build feature_types

package contracts

import (
	"github.com/iotaledger/wasplib/tools/schema/generator"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRustToGo(t *testing.T) {
	//t.SkipNow()
	err := generator.RustConvertor(generator.RustToGoLine, "converted/$c/$c.go")
	require.NoError(t, err)
}

func TestRustToJava(t *testing.T) {
	//t.SkipNow()
	err := generator.RustConvertor(generator.RustToJavaLine, "converted/$c/$C.java")
	require.NoError(t, err)
}
