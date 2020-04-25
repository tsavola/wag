// Copyright (c) 2018 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build wagarm64,!waginterp arm64,!wagamd64,!waginterp

package reglayout

const (
	AllocIntFirst = 2
	AllocIntLast  = 25

	AllocFloatFirst = 2
	AllocFloatLast  = 31
)
