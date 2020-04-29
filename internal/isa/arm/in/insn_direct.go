// Copyright (c) 2018 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !waginterp

package in

// Add/subtract instructions

func (op Addsub) OpcodeImm() RegRegImm12ShiftSf {
	return RegRegImm12ShiftSf(uint32(op)<<30 | 0<<29 | 0x11<<24)
}

func (op Addsub) OpcodeRegExt() RegRegImm3ExtRegSf {
	return RegRegImm3ExtRegSf(uint32(op)<<30 | 0<<29 | 0x0b<<24 | 0<<22 | 1<<21)
}

// Logical instructions

func (op Logic) OpcodeImm() RegRegImm6Imm6NSf {
	return RegRegImm6Imm6NSf(uint32(op)<<29 | 0x24<<23)
}

func (op Logic) OpcodeReg() RegRegImm6RegShiftSf {
	return RegRegImm6RegShiftSf(uint32(op)<<29 | 0x0a<<24 | 0<<21)
}

// Bitfield instruction's "opc" field
const (
	ExtendS = Bitfield(0) // SBFM
	ExtendU = Bitfield(2) // UBFM
)

func (op Bitfield) Opcode() RegRegImm6Imm6NSf {
	return RegRegImm6Imm6NSf(uint32(op)<<29 | 0x26<<23 | 0<<22)
}

// Data-processing (2 source) instruction's "opcode" field
const (
	DivisionUnsigned = DataProcessing2(0x2) // UDIV
	DivisionSigned   = DataProcessing2(0x3) // SDIV
	VariableShiftL   = DataProcessing2(0x8) // LSLV
	VariableShiftLR  = DataProcessing2(0x9) // LSRV
	VariableShiftAR  = DataProcessing2(0xa) // ASRV
	VariableShiftRR  = DataProcessing2(0xb) // RORV
)

func (op DataProcessing2) OpcodeReg() RegRegRegSf {
	return RegRegRegSf(0<<30 | 0<<29 | 0xd6<<21 | uint32(op)<<10)
}

// Floating-point (1 source) instruction's bits 15-20
const (
	//                            17     15
	UnaryFloatAbs     = UnaryFloat(0<<2 | 1<<0) // FABS
	UnaryFloatNeg     = UnaryFloat(0<<2 | 2<<0) // FNEG
	UnaryFloatSqrt    = UnaryFloat(0<<2 | 3<<0) // FSQRT
	UnaryFloatCvtTo32 = UnaryFloat(1<<2 | 0<<0) // FCVTto32
	UnaryFloatCvtTo64 = UnaryFloat(1<<2 | 1<<0) // FCVTto64

	//                             18     15
	UnaryFloatRIntN = UnaryFloat(1<<3 | 0<<0) // FRINTN
	UnaryFloatRIntP = UnaryFloat(1<<3 | 1<<0) // FRINTP
	UnaryFloatRIntM = UnaryFloat(1<<3 | 2<<0) // FRINTM
	UnaryFloatRIntZ = UnaryFloat(1<<3 | 3<<0) // FRINTZ
)

func (op UnaryFloat) Opcode() RegRegType {
	return RegRegType(0<<31 | 0<<30 | 0<<29 | 0x1e<<24 | 1<<21 | uint32(op)<<15 | 0x10<<10)
}

// Floating-point (2 source) instruction's bits 10-15
const (
	//                             13     12     10
	BinaryFloatAdd = BinaryFloat(1<<3 | 0<<2 | 2<<0) // FADD
	BinaryFloatSub = BinaryFloat(1<<3 | 1<<2 | 2<<0) // FSUB

	//                             15     12     10
	BinaryFloatMul = BinaryFloat(0<<5 | 0<<2 | 2<<0) // FMUL
	BinaryFloatDiv = BinaryFloat( /***/ 1<<2 | 2<<0) // FDIV

	//                             14     12     10
	BinaryFloatMax = BinaryFloat(1<<4 | 0<<2 | 2<<0) // FMAX
	BinaryFloatMin = BinaryFloat(1<<4 | 1<<2 | 2<<0) // FMIN
)

func (op BinaryFloat) OpcodeReg() RegRegRegType {
	return RegRegRegType(0<<31 | 0<<30 | 0<<29 | 0x1e<<24 | 1<<21 | uint32(op)<<10)
}

// Floating-point/integer instruction's "rmode" and "opcode" fields
const (
	//                                   19     16
	ConvertIntS      = ConvertCategory(0<<3 | 2<<0) // SCVTF
	ConvertIntU      = ConvertCategory(0<<3 | 3<<0) // UCVTF
	ReinterpretFloat = ConvertCategory(0<<3 | 6<<0) // FMOVtog
	ReinterpretInt   = ConvertCategory(0<<3 | 7<<0) // FMOVfromg
	TruncFloatS      = ConvertCategory(3<<3 | 0<<0) // FCVTZS
	TruncFloatU      = ConvertCategory(3<<3 | 1<<0) // FCVTZU
)

func (op ConvertCategory) Opcode() RegRegTypeSf {
	return RegRegTypeSf(0<<30 | 0<<29 | 0x1e<<24 | 1<<21 | uint32(op)<<16 | 0<<10)
}

// Load/store instruction's most significant half-word excluding bit 24 (and 21)
const (
	//                   30      27      26     22
	StoreB   = Memory(0<<14 | 7<<11 | 0<<10 | 0<<6) // STRB, STURB
	LoadB    = Memory(0<<14 | 7<<11 | 0<<10 | 1<<6) // LDRB, LDURB
	LoadSB64 = Memory(0<<14 | 7<<11 | 0<<10 | 2<<6) // LDRSB64, LDURSB64
	LoadSB32 = Memory(0<<14 | 7<<11 | 0<<10 | 3<<6) // LDRSB32, LDURSB32
	StoreH   = Memory(1<<14 | 7<<11 | 0<<10 | 0<<6) // STRH, STURH
	LoadH    = Memory(1<<14 | 7<<11 | 0<<10 | 1<<6) // LDRH, LDURH
	LoadSH64 = Memory(1<<14 | 7<<11 | 0<<10 | 2<<6) // LDRSH64, LDURSH64
	LoadSH32 = Memory(1<<14 | 7<<11 | 0<<10 | 3<<6) // LDRSH32, LDURSH32
	StoreW   = Memory(2<<14 | 7<<11 | 0<<10 | 0<<6) // STR, STUR
	LoadW    = Memory(2<<14 | 7<<11 | 0<<10 | 1<<6) // LDR, LDUR
	LoadSW64 = Memory(2<<14 | 7<<11 | 0<<10 | 2<<6) // LDRSW, LDURSW
	StoreF32 = Memory(2<<14 | 7<<11 | 1<<10 | 0<<6) // STRf, STURf
	LoadF32  = Memory(2<<14 | 7<<11 | 1<<10 | 1<<6) // LDRf, LDURf
	StoreD   = Memory(3<<14 | 7<<11 | 0<<10 | 0<<6) // STR, STUR
	LoadD    = Memory(3<<14 | 7<<11 | 0<<10 | 1<<6) // LDR, LDUR
	StoreF64 = Memory(3<<14 | 7<<11 | 1<<10 | 0<<6) // STRf, STURf
	LoadF64  = Memory(3<<14 | 7<<11 | 1<<10 | 1<<6) // LDRf, LDURf
)

func (op Memory) OpcodeUnscaled() RegRegImm9 {
	return RegRegImm9(uint32(op)<<16 | 0<<24 | 0<<21 | 0<<10)
}

func (op Memory) OpcodeReg() RegRegSOptionReg {
	return RegRegSOptionReg(uint32(op)<<16 | 0<<24 | 1<<21 | 2<<10)
}
