// Copyright (c) 2018 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package in

func Int9(i int32) uint32    { return uint32(i) & 0x1ff }
func Uint12(i uint64) uint32 { return uint32(i) & 0xfff }
func Int14(i int32) uint32   { return uint32(i) & 0x3fff }
func Uint16(i uint64) uint32 { return uint32(i) & 0xffff }
func Int19(i int32) uint32   { return uint32(i) & 0x7ffff }
func Int26(i int32) uint32   { return uint32(i) & 0x3ffffff }

type Cond uint32

const (
	EQ = Cond(0x0) // equal to
	NE = Cond(0x1) // not equal to
	CS = Cond(0x2) // carry set
	CC = Cond(0x3) // carry clear
	MI = Cond(0x4) // minus, negative
	PL = Cond(0x5) // positive or zero
	VS = Cond(0x6) // signed overflow
	VC = Cond(0x7) // no signed overflow
	HI = Cond(0x8) // greater than (unsigned)
	LS = Cond(0x9) // less than or equal to (unsigned)
	GE = Cond(0xa) // greater than or equal to (signed)
	LT = Cond(0xb) // less than (signed)
	GT = Cond(0xc) // greater than (signed)
	LE = Cond(0xd) // less than or equal to (signed)

	HS = CS // greater than or equal to (unsigned)
	LO = CC // less than (unsigned)
)

type S uint32

type Shift uint32

type Ext uint32

type Imm16 uint32
type Imm26 uint32
type CondImm19 uint32
type Reg uint32
type RegImm14Bit uint32
type RegImm16HwSf uint32
type RegImm19Imm2 uint32
type RegImm19Size uint32
type RegRegImm3ExtRegSf uint32
type RegRegImm6Imm6NSf uint32
type RegRegImm6RegShiftSf uint32
type RegRegImm9 uint32
type RegRegImm9Size uint32
type RegRegImm12ShiftSf uint32
type RegRegImm12Size uint32
type RegRegSf uint32
type RegRegSOptionReg uint32
type RegRegSOptionRegSize uint32
type RegRegType uint32
type RegRegTypeSf uint32
type RegRegCondRegSf uint32
type RegRegCondRegType uint32
type RegRegRegSf uint32
type RegRegRegType uint32
type RegRegRegRegSf uint32
type DiscardRegRegType uint32
