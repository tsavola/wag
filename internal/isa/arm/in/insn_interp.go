// Copyright (c) 2020 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build waginterp

package in

const (
	prefixB = 1 << 4
)

const (
	// x1
	BRK = Imm16(0)

	// x1
	B  = Imm26(prefixB | 0)
	BL = Imm26(prefixB | 1)

	// x16
	Bc = CondImm19(16)

	// x1
	BR  = Reg(1)
	BLR = Reg(2)
	RET = Reg(3)

	// x1
	TBZ  = RegImm14Bit(prefixB | 2)
	TBNZ = RegImm14Bit(prefixB | 3)

	// x2
	MOVN = RegImm16HwSf(8)
	MOVZ = RegImm16HwSf(10)
	MOVK = RegImm16HwSf(12)

	// x1
	ADR = RegImm19Imm2(prefixB | 8)

	// x2
	CBZ  = RegImm19Size(prefixB | 4)
	CBNZ = RegImm19Size(prefixB | 6)

	// x4
	ADDe  = RegRegImm3ExtRegSf(312)
	SUBe  = RegRegImm3ExtRegSf(316)
	SUBSe = RegRegImm3ExtRegSf(320)

	// x2
	ANDi = RegRegImm6Imm6NSf(374)
	ORRi = RegRegImm6Imm6NSf(376)
	EORi = RegRegImm6Imm6NSf(378)
	SBFM = RegRegImm6Imm6NSf(380)
	UBFM = RegRegImm6Imm6NSf(382)

	// x8
	ADDs  = RegRegImm6RegShiftSf(232)
	SUBs  = RegRegImm6RegShiftSf(240)
	SUBSs = RegRegImm6RegShiftSf(248)
	ANDs  = RegRegImm6RegShiftSf(256)
	ORRs  = RegRegImm6RegShiftSf(264)
	EORs  = RegRegImm6RegShiftSf(272)
	ANDSs = RegRegImm6RegShiftSf(280)

	// x2
	LDRpost  = RegRegImm9Size(324)
	STRpre   = RegRegImm9Size(326)
	STURB    = RegRegImm9Size(328)
	LDURB    = RegRegImm9Size(330)
	LDURSB64 = RegRegImm9Size(332)
	LDURSB32 = RegRegImm9Size(334)
	STURH    = RegRegImm9Size(336)
	LDURH    = RegRegImm9Size(338)
	LDURSH64 = RegRegImm9Size(340)
	LDURSH32 = RegRegImm9Size(342)
	STUR     = RegRegImm9Size(344)
	LDUR     = RegRegImm9Size(346)
	LDURSW   = RegRegImm9Size(348)
	STURf    = RegRegImm9Size(350)
	LDURf    = RegRegImm9Size(352)

	// x2
	ADDi  = RegRegImm12ShiftSf(408)
	ADDSi = RegRegImm12ShiftSf(410)
	SUBi  = RegRegImm12ShiftSf(412)
	SUBSi = RegRegImm12ShiftSf(414)

	// x2
	STR = RegRegImm12Size(420)
	LDR = RegRegImm12Size(422)

	// x2
	RBIT = RegRegSf(424)
	CLZ  = RegRegSf(426)

	// x8
	STRBr    = RegRegSOptionRegSize(128)
	LDRBr    = RegRegSOptionRegSize(136)
	LDRSBr64 = RegRegSOptionRegSize(144)
	LDRSBr32 = RegRegSOptionRegSize(152)
	STRHr    = RegRegSOptionRegSize(160)
	LDRHr    = RegRegSOptionRegSize(168)
	LDRSHr64 = RegRegSOptionRegSize(176)
	LDRSHr32 = RegRegSOptionRegSize(184)
	STRr     = RegRegSOptionRegSize(192)
	LDRr     = RegRegSOptionRegSize(200)
	LDRSWr   = RegRegSOptionRegSize(208)
	STRrf    = RegRegSOptionRegSize(216)
	LDRrf    = RegRegSOptionRegSize(224)

	// x32
	CSEL  = RegRegCondRegSf(64)
	CSINC = RegRegCondRegSf(96)

	// x32
	FCSEL = RegRegCondRegType(32)

	// x2
	FMOV     = RegRegType(354)
	FABS     = RegRegType(356)
	FNEG     = RegRegType(358)
	FSQRT    = RegRegType(360)
	FCVTto32 = RegRegType(362)
	FCVTto64 = RegRegType(364)
	FRINTN   = RegRegType(366)
	FRINTP   = RegRegType(368)
	FRINTM   = RegRegType(370)
	FRINTZ   = RegRegType(372)

	// x4
	SCVTF     = RegRegTypeSf(288)
	UCVTF     = RegRegTypeSf(292)
	FMOVtog   = RegRegTypeSf(296)
	FMOVfromg = RegRegTypeSf(300)
	FCVTZS    = RegRegTypeSf(304)
	FCVTZU    = RegRegTypeSf(308)

	// x2
	FADD = RegRegRegType(384)
	FSUB = RegRegRegType(386)
	FMUL = RegRegRegType(388)
	FDIV = RegRegRegType(390)
	FMAX = RegRegRegType(392)
	FMIN = RegRegRegType(394)

	// x2
	LSLV = RegRegRegSf(396)
	LSRV = RegRegRegSf(398)
	ASRV = RegRegRegSf(400)
	RORV = RegRegRegSf(402)
	UDIV = RegRegRegSf(404)
	SDIV = RegRegRegSf(406)

	// x2
	MADD = RegRegRegRegSf(428)
	MSUB = RegRegRegRegSf(430)

	// x2
	FCMP = DiscardRegRegType(14)
)

// Add/subtract instructions

func (op Addsub) OpcodeImm() RegRegImm12ShiftSf {
	return RegRegImm12ShiftSf(0)
}

func (op Addsub) OpcodeRegExt() RegRegImm3ExtRegSf {
	return RegRegImm3ExtRegSf(0)
}

// Logical instructions

func (op Logic) OpcodeImm() RegRegImm6Imm6NSf {
	return RegRegImm6Imm6NSf(0)
}

func (op Logic) OpcodeReg() RegRegImm6RegShiftSf {
	return RegRegImm6RegShiftSf(0)
}

// Bitfield instructions
const (
	ExtendS = Bitfield(0) // SBFM
	ExtendU = Bitfield(0) // UBFM
)

func (op Bitfield) Opcode() RegRegImm6Imm6NSf {
	return RegRegImm6Imm6NSf(0)
}

// Data-processing (2 source) instructions
const (
	DivisionUnsigned = DataProcessing2(0) // UDIV
	DivisionSigned   = DataProcessing2(0) // SDIV
	VariableShiftL   = DataProcessing2(0) // LSLV
	VariableShiftLR  = DataProcessing2(0) // LSRV
	VariableShiftAR  = DataProcessing2(0) // ASRV
	VariableShiftRR  = DataProcessing2(0) // RORV
)

func (op DataProcessing2) OpcodeReg() RegRegRegSf {
	return RegRegRegSf(0)
}

// Floating-point (1 source) instructions
const (
	UnaryFloatAbs     = UnaryFloat(0) // FABS
	UnaryFloatNeg     = UnaryFloat(0) // FNEG
	UnaryFloatSqrt    = UnaryFloat(0) // FSQRT
	UnaryFloatCvtTo32 = UnaryFloat(0) // FCVTto32
	UnaryFloatCvtTo64 = UnaryFloat(0) // FCVTto64
	UnaryFloatRIntN   = UnaryFloat(0) // FRINTN
	UnaryFloatRIntP   = UnaryFloat(0) // FRINTP
	UnaryFloatRIntM   = UnaryFloat(0) // FRINTM
	UnaryFloatRIntZ   = UnaryFloat(0) // FRINTZ
)

func (op UnaryFloat) Opcode() RegRegType {
	return RegRegType(0)
}

// Floating-point (2 source) instructions
const (
	BinaryFloatAdd = BinaryFloat(0) // FADD
	BinaryFloatSub = BinaryFloat(0) // FSUB
	BinaryFloatMul = BinaryFloat(0) // FMUL
	BinaryFloatDiv = BinaryFloat(0) // FDIV
	BinaryFloatMax = BinaryFloat(0) // FMAX
	BinaryFloatMin = BinaryFloat(0) // FMIN
)

func (op BinaryFloat) OpcodeReg() RegRegRegType {
	return RegRegRegType(0)
}

// Floating-point/integer instructions
const (
	ConvertIntS      = ConvertCategory(0) // SCVTF
	ConvertIntU      = ConvertCategory(0) // UCVTF
	ReinterpretFloat = ConvertCategory(0) // FMOVtog
	ReinterpretInt   = ConvertCategory(0) // FMOVfromg
	TruncFloatS      = ConvertCategory(0) // FCVTZS
	TruncFloatU      = ConvertCategory(0) // FCVTZU
)

func (op ConvertCategory) Opcode() RegRegTypeSf {
	return RegRegTypeSf(0)
}

// Load/store instructions
const (
	StoreB   = Memory(0) // STRBr, STURB
	LoadB    = Memory(0) // LDRBr, LDURB
	LoadSB64 = Memory(0) // LDRSBr64, LDURSB64
	LoadSB32 = Memory(0) // LDRSBr32, LDURSB32
	StoreH   = Memory(0) // STRHr, STURH
	LoadH    = Memory(0) // LDRHr, LDURH
	LoadSH64 = Memory(0) // LDRSHr64, LDURSH64
	LoadSH32 = Memory(0) // LDRSHr32, LDURSH32
	StoreW   = Memory(0) // STRr, STUR
	LoadW    = Memory(0) // LDRr, LDUR
	LoadSW64 = Memory(0) // LDRSWr, LDURSW
	StoreF32 = Memory(0) // STRrf, STURf
	LoadF32  = Memory(0) // LDRrf, LDURf
	StoreD   = Memory(0) // STRr, STUR
	LoadD    = Memory(0) // LDRr, LDUR
	StoreF64 = Memory(0) // STRrf, STURf
	LoadF64  = Memory(0) // LDRrf, LDURf
)

func (op Memory) OpcodeUnscaled() RegRegImm9 {
	return RegRegImm9(0)
}

func (op Memory) OpcodeReg() RegRegSOptionReg {
	return RegRegSOptionReg(0)
}
