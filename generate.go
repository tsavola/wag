// Copyright (c) 2018 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wag

//go:generate go run internal/cmd/arm-in/generate.go
//go:generate go run internal/cmd/opcodes/generate.go
//go:generate go run internal/cmd/syscalls/generate.go
//go:generate make -C internal/test/runner
