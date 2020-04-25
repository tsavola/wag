// Copyright (c) 2020 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build waginterp

package in

import (
	"github.com/tsavola/wag/internal/gen/reg"
	"github.com/tsavola/wag/wa"
)

const (
	shiftA = 22
	shiftB = 27
)

func sizeBit(t wa.Size) uint32 {
	return uint32(t >> 3)
}

func typeBit(t wa.Type) uint32 {
	return uint32((t & 8) >> 3)
}

func R(r reg.R) uint32 {
	if r < 16 {
		return uint32(r)
	} else {
		return uint32(r) - 16
	}
}

const (
	Unscaled = S(0)
	Scaled   = S(1)
)

const (
	LSL = Shift(0)
	LSR = Shift(1)
	ASR = Shift(2)
)

const (
	UXTB = Ext(0<<1 | Unscaled)
	UXTH = Ext(1<<1 | Unscaled)
	UXTW = Ext(2<<1 | Unscaled)
	UXTX = Ext(3<<1 | Unscaled)
	SXTB = Ext(0<<1 | Scaled)
	SXTH = Ext(1<<1 | Scaled)
	SXTW = Ext(2<<1 | Scaled)
	SXTX = Ext(3<<1 | Scaled)
)

func scaleBit(ext Ext) uint32 {
	return uint32(ext & 1)
}

func order(ext Ext) uint32 {
	return uint32(ext >> 1)
}

func SizeZeroExt(t wa.Size) Ext {
	bit3 := uint32(t & 8)
	return UXTW | Ext(bit3>>2)
}

func SizeSignExt(t wa.Size) Ext {
	bit3 := uint32(t & 8)
	return SXTW | Ext(bit3>>2)
}

func (op Imm16) I16(imm uint32) uint32 {
	return uint32(op)<<shiftA | imm
}

func (op Imm26) I26(imm uint32) uint32 {
	return uint32(op)<<shiftB | imm
}

func (op Reg) Rn(rn reg.R) uint32 {
	return uint32(op)<<shiftA | R(rn)
}

func (op CondImm19) CondI19(cond Cond, imm uint32) uint32 {
	return (uint32(op)+uint32(cond))<<shiftA | imm
}

func (op RegImm14Bit) RtI14Bit(rt reg.R, imm, bit uint32) uint32 {
	return uint32(op)<<shiftB | bit<<18 | imm<<4 | R(rt)
}

func (op RegImm16HwSf) RdI16Hw(rd reg.R, imm, hw uint32, t wa.Size) uint32 {
	return (uint32(op)+sizeBit(t))<<shiftA | hw<<20 | imm<<4 | R(rd)
}

func (op RegImm19Imm2) RdI19hiI2lo(r reg.R, hi, lo uint32) uint32 {
	return uint32(op)<<shiftB | lo<<23 | hi<<4 | R(r)
}

func (op RegImm19Size) RtI19(r reg.R, imm uint32, t wa.Size) uint32 {
	return (uint32(op)+sizeBit(t))<<shiftB | imm<<4 | R(r)
}

func (op RegRegImm3ExtRegSf) RdRnI3ExtRm(rd, rn reg.R, imm uint32, option Ext, rm reg.R, t wa.Size) uint32 {
	return (uint32(op)+sizeBit(t)<<1+scaleBit(option))<<shiftA | R(rm)<<13 | order(option)<<11 | imm<<8 | R(rn)<<4 | R(rd)
}

func (op RegRegImm6Imm6NSf) RdRnI6sI6r(rd, rn reg.R, imms, immr uint32, t wa.Size) uint32 {
	return (uint32(op)+sizeBit(t))<<shiftA | immr<<14 | imms<<8 | R(rn)<<4 | R(rd)
}

func (op RegRegImm6RegShiftSf) RdRnI6RmS2(rd, rn reg.R, imm uint32, rm reg.R, shift Shift, t wa.Size) uint32 {
	return (uint32(op)+sizeBit(t)<<1+uint32(shift))<<shiftA | R(rm)<<14 | imm<<8 | R(rn)<<4 | R(rd)
}

func (op RegRegImm9) RtRnI9(rt, rn reg.R, imm uint32) uint32 {
	return uint32(op)<<shiftA | imm<<8 | R(rn)<<4 | R(rt)
}

func (op RegRegImm9Size) RtRnI9(rt, rn reg.R, imm uint32, t wa.Type) uint32 {
	return (uint32(op)+typeBit(t))<<shiftA | imm<<8 | R(rn)<<4 | R(rt)
}

func (op RegRegImm12ShiftSf) RdRnI12S2(rd, rn reg.R, imm, shift uint32, t wa.Size) uint32 {
	return (uint32(op)+sizeBit(t))<<shiftA | shift<<20 | imm<<8 | R(rn)<<4 | R(rd)
}

func (op RegRegImm12Size) RdRnI12(rt, rn reg.R, imm uint32, t wa.Type) uint32 {
	return (uint32(op)+typeBit(t))<<shiftA | imm<<8 | R(rn)<<4 | R(rt)
}

func (op RegRegSf) RdRn(rd, rn reg.R, t wa.Size) uint32 {
	return (uint32(op)+sizeBit(t))<<shiftA | R(rn)<<4 | R(rd)
}

func (op RegRegSOptionReg) RtRnSOptionRm(rt, rn reg.R, s S, option Ext, rm reg.R) uint32 {
	return (uint32(op)+scaleBit(option)<<1+uint32(s))<<shiftA | R(rm)<<10 | order(option)<<8 | R(rn)<<4 | R(rt)
}

func (op RegRegSOptionRegSize) RtRnSOptionRm(rt, rn reg.R, s S, option Ext, rm reg.R, t wa.Type) uint32 {
	return (uint32(op)+typeBit(t)<<2+scaleBit(option)<<1+uint32(s))<<shiftA | R(rm)<<10 | order(option)<<8 | R(rn)<<4 | R(rt)
}

func (op RegRegCondRegSf) RdRnCondRm(rd, rn reg.R, cond Cond, rm reg.R, t wa.Size) uint32 {
	return (uint32(op)+sizeBit(t)<<1+uint32(cond))<<shiftA | R(rm)<<8 | R(rn)<<4 | R(rd)
}

func (op RegRegCondRegType) RdRnCondRm(rd, rn reg.R, cond Cond, rm reg.R, t wa.Size) uint32 {
	return (uint32(op)+sizeBit(t)<<4+uint32(cond))<<shiftA | R(rm)<<8 | R(rn)<<4 | R(rd)
}

func (op RegRegType) RdRn(rd, rn reg.R, t wa.Size) uint32 {
	return (uint32(op)+sizeBit(t))<<shiftA | R(rn)<<4 | R(rd)
}

func (op RegRegTypeSf) RdRn(rd, rn reg.R, floatType, intType wa.Size) uint32 {
	return (uint32(op)+sizeBit(intType)<<1+sizeBit(floatType))<<shiftA | R(rn)<<4 | R(rd)
}

func (op RegRegRegSf) RdRnRm(rd, rn, rm reg.R, t wa.Size) uint32 {
	return (uint32(op)+sizeBit(t))<<shiftA | R(rm)<<8 | R(rn)<<4 | R(rd)
}

func (op RegRegRegType) RdRnRm(rd, rn, rm reg.R, t wa.Size) uint32 {
	return (uint32(op)+sizeBit(t))<<shiftA | R(rm)<<8 | R(rn)<<4 | R(rd)
}

func (op RegRegRegRegSf) RdRnRaRm(rd, rn, ra, rm reg.R, t wa.Size) uint32 {
	return (uint32(op)+sizeBit(t))<<shiftA | R(rm)<<12 | R(ra)<<8 | R(rn)<<4 | R(rd)
}

func (op DiscardRegRegType) RnRm(rn, rm reg.R, t wa.Size) uint32 {
	return (uint32(op)+sizeBit(t))<<shiftA | R(rm)<<4 | R(rn)
}
