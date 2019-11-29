# go.atomicbool

[![Build Status](https://travis-ci.org/Andrew-M-C/go.atomicbool.svg?branch=master)](https://travis-ci.org/Andrew-M-C/go.atomicbool)
[![GoDoc](https://godoc.org/github.com/Andrew-M-C/go.atomicbool?status.svg)](https://godoc.org/github.com/Andrew-M-C/go.atomicbool)
[![Coverage Status](https://coveralls.io/repos/github/Andrew-M-C/go.atomicbool/badge.svg?branch=master)](https://coveralls.io/github/Andrew-M-C/go.atomicbool?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/Andrew-M-C/go.atomicbool)](https://goreportcard.com/report/github.com/Andrew-M-C/go.atomicbool)

An atomic Go boolean type with function interfaces like official sync/atomic.

## Example:

[Playground](https://play.golang.org/p/qt5O3_504DQ)

```go
package main

import (
	"fmt"
	"github.com/Andrew-M-C/go.atomicbool"
)

func main() {
	f := fmt.Println

	b := atomicbool.New(true)
	f(b.Load())

	b.Store(false)
	f(b.Load())

	swapped := b.CompareAndSwap(false, true)
	f(swapped, b.Load())
}

```
