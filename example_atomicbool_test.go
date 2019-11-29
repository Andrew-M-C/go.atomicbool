package atomicbool_test

import (
	"fmt"

	atomicbool "github.com/Andrew-M-C/go.atomicbool"
)

func ExampleB() {
	f := fmt.Println

	b := atomicbool.New(true)
	f(b.Load())

	b.Store(false)
	f(b.Load())

	swapped := b.CompareAndSwap(false, true)
	f(swapped, b.Load())
	// Output:
	// true
	// false
	// true true
}
