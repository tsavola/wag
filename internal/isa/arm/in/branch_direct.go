// Copyright (c) 2020 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !waginterp

package in

import (
	"fmt"
)

func UpdateBranchOffset(insn uint32, offset int32) uint32 {
	switch {
	case insn>>25 == 0x2a: // Conditional branch.
		return insn&^(0x7ffff<<5) | Int19(offset)<<5

	case (insn>>26)&0x1f == 0x05: // Unconditional branch.
		return insn&^0x3ffffff | Int26(offset)

	default:
		panic(fmt.Sprintf("unknown branch instruction encoding: %#v", insn))
	}
}
