// Copyright (c) 2018 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package in

type Addsub uint8
type Logic uint8
type Bitfield uint8
type DataProcessing2 uint8
type UnaryFloat uint8
type BinaryFloat uint8
type ConvertCategory uint8
type Memory uint16

// Add/subtract instruction's "op" field
const (
	AddsubAdd = Addsub(0) // ADDi, ADDe
	AddsubSub = Addsub(1) // SUBi, SUBe
)

// Logical instruction's "opc" field
const (
	LogicAnd = Logic(0) // ANDi, ANDs
	LogicOrr = Logic(1) // ORRi, ORRs
	LogicEor = Logic(2) // EORi, EORs
)
