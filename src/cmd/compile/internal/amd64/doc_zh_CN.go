// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

package amd64 // import "cmd/compile/internal/amd64"

import (
    "cmd/compile/internal/big"
    "cmd/compile/internal/gc"
    "cmd/internal/obj"
    "cmd/internal/obj/x86"
    "fmt"
)

// For ProgInfo.
const (
    AX  = 1 << (x86.REG_AX - x86.REG_AX)
    BX  = 1 << (x86.REG_BX - x86.REG_AX)
    CX  = 1 << (x86.REG_CX - x86.REG_AX)
    DX  = 1 << (x86.REG_DX - x86.REG_AX)
    DI  = 1 << (x86.REG_DI - x86.REG_AX)
    SI  = 1 << (x86.REG_SI - x86.REG_AX)
    R15 = 1 << (x86.REG_R15 - x86.REG_AX)
    X0  = 1 << 16
)

const (
    LeftRdwr  uint32 = gc.LeftRead | gc.LeftWrite
    RightRdwr uint32 = gc.RightRead | gc.RightWrite
)

const (
    NREGVAR = 32
)

const (
    ODynam   = 1 << 0
    OAddable = 1 << 1
)

var MAXWIDTH int64 = 1 << 50

func BtoF(b uint64) int

func BtoR(b uint64) int

// *    bit    reg
//  		X0
//  	...
//  		X15

// reg
//  		X0
//  	...
//  		X15
func FtoB(f int) uint64

func Main()

func RtoB(r int) uint64
