// USB armory Mk II support for tamago/arm
// https://github.com/inversepath/tamago
//
// Copyright (c) F-Secure Corporation
// https://foundry.f-secure.com
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.
//
// +build tamago,arm

// Go applications meant for tamago/arm on the USB armory Mk II simply need to
// import this package for all necessary hardware initialization.

package usbarmory

import (
	_ "unsafe"

	"github.com/inversepath/tamago/imx6"
	_ "github.com/inversepath/tamago/imx6/imx6ul"
)

//go:linkname ramSize runtime.ramSize
var ramSize uint32 = 0x20000000 // 512 MB

//go:linkname printk runtime.printk
func printk(c byte) {
	imx6.UART2.Write(c)
}
