package x86

import (
	"github.com/tsavola/wag/internal/gen"
	"github.com/tsavola/wag/internal/regs"
	"github.com/tsavola/wag/internal/types"
	"github.com/tsavola/wag/internal/values"
)

type rexSizePrefix struct{}

func (rexSizePrefix) writeTo(code gen.Coder, t types.T, ro, index, rmOrBase regs.R) {
	var rex byte

	switch t.Size() {
	case types.Size32:

	case types.Size64:
		rex |= rexW

	default:
		panic(t)
	}

	writeRexTo(code, rex, ro, index, rmOrBase)
}

var (
	rexSize rexSizePrefix
)

var (
	Movzx8 = insnModRegFromReg{0x0f, 0xb6}

	MovImm = insnPrefixRegImm{rexSize, 0xb8}

	Mul = insnPrefixModOpReg{rexSize, []byte{0xf7}, 4}
	Div = insnPrefixModOpReg{rexSize, []byte{0xf7}, 6}
	Inc = insnPrefixModOpReg{rexSize, []byte{0xff}, 0}
	Dec = insnPrefixModOpReg{rexSize, []byte{0xff}, 1}

	Push = insnReg_sizeless_PrefixModOpReg{insnReg{0x50}, insnPrefixModOpReg{rexSize, []byte{0xff}, 6}}
	Pop  = insnReg_sizeless_PrefixModOpReg{insnReg{0x58}, insnPrefixModOpReg{rexSize, []byte{0x8f}, 0}}

	Add    = insnPrefixModRegFromReg{rexSize, []byte{0x03}}
	Or     = insnPrefixModRegFromReg{rexSize, []byte{0x0b}}
	And    = insnPrefixModRegFromReg{rexSize, []byte{0x23}}
	Sub    = insnPrefixModRegFromReg{rexSize, []byte{0x2b}}
	Xor    = insnPrefixModRegFromReg{rexSize, []byte{0x33}}
	Cmp    = insnPrefixModRegFromReg{rexSize, []byte{0x3b}}
	Movsxd = insnPrefixModRegFromReg{rexSize, []byte{0x63}}
	Cmovl  = insnPrefixModRegFromReg{rexSize, []byte{0x0f, 0x4c}}
	Bsf    = insnPrefixModRegFromReg{rexSize, []byte{0x0f, 0xbc}}

	Test = insnPrefixModRegToReg{rexSize, []byte{0x85}, ModReg}
	Mov  = insnPrefixModRegToReg{rexSize, []byte{0x89}, ModReg}

	CmpImm32 = insnPrefixModOpRegImm{rexSize, []byte{0x81}, 7}
	ShlImm   = insnPrefixModOpRegImm{rexSize, []byte{0xc1}, 0}
	ShrImm   = insnPrefixModOpRegImm{rexSize, []byte{0xc1}, 5}
	MovImm32 = insnPrefixModOpRegImm{rexSize, []byte{0xc7}, 0}

	CmpFromStack    = insnPrefixModRegSibImm{rexSize, []byte{0x3b}, sib{0, regStackPtr, regStackPtr}}
	MovsxdFromStack = insnPrefixModRegSibImm{rexSize, []byte{0x63}, sib{0, regStackPtr, regStackPtr}}
	MovToStack      = insnPrefixModRegSibImm{rexSize, []byte{0x89}, sib{0, regStackPtr, regStackPtr}}
	MovFromStack    = insnPrefixModRegSibImm{rexSize, []byte{0x8b}, sib{0, regStackPtr, regStackPtr}}

	MovsxdFromIndirectScaleIndex = insnPrefixModRegCustomSibImm{rexSize, []byte{0x63}}
	MovFromIndirectScaleIndex    = insnPrefixModRegCustomSibImm{rexSize, []byte{0x8b}}

	AddImm = insnPrefixArithmeticModOpRegImm{rexSize, 0}
	SubImm = insnPrefixArithmeticModOpRegImm{rexSize, 5}
	CmpImm = insnPrefixArithmeticModOpRegImm{rexSize, 7}
)

var binaryIntInsns = map[string]binaryInsn{
	"add": Add,
	"and": And,
	"or":  Or,
	"sub": Sub,
	"xor": Xor,
}

var binaryIntConditions = map[string]values.Condition{
	"eq":   values.EQ,
	"gt_s": values.GT_S,
	"gt_u": values.GT_U,
	"lt_s": values.LT_S,
	"ne":   values.NE,
}

func (x86 X86) unaryIntOp(code gen.RegCoder, name string, t types.T, x values.Operand) values.Operand {
	switch name {
	case "ctz":
		var targetReg regs.R

		sourceReg, own := x86.opBorrowReg(code, t, x)
		if own {
			targetReg = sourceReg
		} else {
			targetReg = code.OpAllocReg(t)
		}

		Bsf.opReg(code, t, targetReg, sourceReg)
		return values.RegTempOperand(targetReg)

	case "eqz":
		reg, own := x86.opBorrowReg(code, t, x)
		if own {
			defer code.FreeReg(t, reg)
		}

		Test.op(code, t, reg, reg)
		return values.ConditionFlagsOperand(values.EQ)
	}

	panic(name)
}

func (x86 X86) binaryIntOp(code gen.RegCoder, name string, t types.T, a, b values.Operand) values.Operand {
	value, immediate := b.CheckImmValue(t)

	if immediate && (value < -0x80000000 || value >= 0x80000000) {
		reg := code.OpAllocReg(t)
		b = x86.OpMove(code, t, reg, b)
		immediate = false
	}

	switch name {
	case "add":
		if immediate {
			switch value {
			case 1:
				reg := x86.opOwnReg(code, t, a)

				Inc.op(code, t, reg)
				return values.RegTempOperand(reg)

			case -1:
				reg := x86.opOwnReg(code, t, a)

				Dec.op(code, t, reg)
				return values.RegTempOperand(reg)
			}
		}

	case "sub":
		if immediate {
			switch value {
			case 1:
				reg := x86.opOwnReg(code, t, a)

				Dec.op(code, t, reg)
				return values.RegTempOperand(reg)

			case -1:
				reg := x86.opOwnReg(code, t, a)

				Inc.op(code, t, reg)
				return values.RegTempOperand(reg)
			}
		}

	case "div_u":
		if immediate && value > 0 && isPowerOfTwo(uint64(value)) {
			reg := x86.opOwnReg(code, t, a)

			ShrImm.op(code, t, reg, uimm8(log2(uint64(value))))
			return values.RegTempOperand(reg)
		} else {
			reg, own := x86.opPrepareDivMul(code, t, a, b)
			if own {
				defer code.FreeReg(t, reg)
			}

			Test.op(code, t, reg, reg)
			Je.op(code)
			code.TrapLinks().DivideByZero.AddSite(code.Len())
			Xor.opReg(code, t, regDividendHi, regDividendHi)
			Div.op(code, t, reg)
			return values.RegTempOperand(regDividendLo)
		}

	case "mul":
		if immediate && value > 0 && isPowerOfTwo(uint64(value)) {
			reg := x86.opOwnReg(code, t, a)

			ShlImm.op(code, t, reg, uimm8(log2(uint64(value))))
			return values.RegTempOperand(reg)
		} else {
			reg, own := x86.opPrepareDivMul(code, t, a, b)
			if own {
				defer code.FreeReg(t, reg)
			}

			Mul.op(code, t, reg)
			return values.RegTempOperand(regDividendLo)
		}
	}

	if insn, found := binaryIntInsns[name]; found {
		targetReg := x86.opOwnReg(code, t, a)

		sourceReg, own := x86.opBorrowReg(code, t, b)
		if own {
			defer code.FreeReg(t, sourceReg)
		}

		insn.opReg(code, t, targetReg, sourceReg)
		return values.RegTempOperand(targetReg)
	}

	if cond, found := binaryIntConditions[name]; found {
		aReg, own := x86.opBorrowReg(code, t, a)
		if own {
			defer code.FreeReg(t, aReg)
		}

		switch b.Storage {
		case values.Imm:
			CmpImm32.op(code, t, aReg, imm32(int(value)))

		case values.ROData:
			Cmp.opIndirect(code, t, aReg, regRODataPtr, b.Addr())

		case values.RegVar:
			Cmp.opReg(code, t, aReg, b.Reg())

		case values.RegTemp:
			Cmp.opReg(code, t, aReg, b.Reg())
			code.FreeReg(t, b.Reg())

		case values.StackVar:
			CmpFromStack.op(code, t, aReg, b.Offset())

		case values.StackPop:
			// TODO: try to allocate register for popping
			AddImm.op(code, types.I64, regStackPtr, wordSize)
			CmpFromStack.op(code, t, aReg, -wordSize)

		default:
			panic(b)
		}

		return values.ConditionFlagsOperand(cond)
	}

	panic(name)
}

func (x86 X86) opPrepareDivMul(code gen.RegCoder, t types.T, a, b values.Operand) (bReg regs.R, own bool) {
	bReg, ok := b.CheckRegVar()
	if !ok {
		bReg, ok = b.CheckRegTemp()
		if !ok {
			bReg, own = x86.opBorrowReg(code, t, b)
		} else if bReg == regDividendLo {
			bReg = code.OpAllocReg(t)
			own = true

			x86.OpMove(code, t, bReg, b)
		}
	}

	aReg, aRegOk := a.CheckAnyReg()
	if !aRegOk || aReg != regDividendLo {
		x86.OpMove(code, t, regDividendLo, a)
	}

	return
}

func isPowerOfTwo(value uint64) bool {
	return (value & (value - 1)) == 0
}

// log2 assumes that value isPowerOfTwo.
func log2(value uint64) (count int) {
	for {
		value >>= 1
		if value == 0 {
			return
		}
		count++
	}
}