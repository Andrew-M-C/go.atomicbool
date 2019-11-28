// Package atomicbool provides a type with function interfaces like official sync/atomic.
package atomicbool

import (
	"sync/atomic"
)

// B is the replacement of bool in atomic scenarios. It is OK to initialize an atomicbool.B with b := atomicbool.B{}
type B struct {
	i int32
}

// New returns an atomicbool value with specified boolean value. It is useful when you want to initialize it as true or values by other channel. Otherwise, just use b := atomicbool.B{}.
func New(b bool) *B {
	if b {
		return &B{1}
	}
	return &B{0}
}

// Store set the boolean value thread-safelly.
func (a *B) Store(b bool) {
	if b {
		atomic.StoreInt32(&a.i, 1)
	} else {
		atomic.StoreInt32(&a.i, 0)
	}
}

// Load read the boolean value thread-safelly.
func (a *B) Load() bool {
	i := atomic.LoadInt32(&a.i)
	return i != 0
}

// CompareAndSwap executes the compare-and-swap operation for an atomicbool.
func (a *B) CompareAndSwap(old, new bool) (swapped bool) {
	var oldI, newI int32
	if old {
		oldI = 1
	} else {
		oldI = 0
	}

	if new {
		newI = 1
	} else {
		newI = 0
	}

	return atomic.CompareAndSwapInt32(&a.i, oldI, newI)
}
