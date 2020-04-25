// Copyright (c) 2019 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !cgo waginterp

package dump

import (
	"errors"
	"io"

	"github.com/tsavola/wag/section"
)

func Text(io.Writer, []byte, uintptr, []uint32, *section.NameSection) error {
	return errors.New("object/debug/dump.Text is incompatible with cgo and waginterp")
}
