package atomicbool

import (
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func getCallerInfo(invokeLevel int) (fileName string, line int, funcName string) {
	funcName = "FILE"
	line = -1
	fileName = "FUNC"

	if invokeLevel <= 0 {
		invokeLevel = 2
	} else {
		invokeLevel++
	}

	pc, fileName, line, ok := runtime.Caller(invokeLevel)
	if ok {
		fileName = filepath.Base(fileName)
		funcName = runtime.FuncForPC(pc).Name()
		funcName = filepath.Ext(funcName)
		funcName = strings.TrimPrefix(funcName, ".")
		// funcName, _ = url.QueryUnescape(funcName)
	}
	// fmt.Println(reflect.TypeOf(pc), reflect.ValueOf(pc))
	return
}

func check(t *testing.T, actual, expected bool) {
	_, li, fu := getCallerInfo(1)
	if actual == expected {
		t.Logf("%s(), Line %d - expected %v and got %v", fu, li, expected, actual)
	} else {
		t.Errorf("%s(), Line %d - expected %v but got %v, error!!!", fu, li, expected, actual)
	}
}

func TestBasic(t *testing.T) {
	var swapped bool
	b := &B{}
	check(t, b.Load(), false)

	b = New(true)
	check(t, b.Load(), true)

	b = New(false)
	check(t, b.Load(), false)

	b.Store(true)
	check(t, b.Load(), true)

	b.Store(false)
	check(t, b.Load(), false)

	swapped = b.CompareAndSwap(false, true)
	check(t, swapped, true)
	check(t, b.Load(), true)

	swapped = b.CompareAndSwap(true, true)
	check(t, swapped, true)
	check(t, b.Load(), true)

	swapped = b.CompareAndSwap(false, false)
	check(t, swapped, false)
	check(t, b.Load(), true)

	swapped = b.CompareAndSwap(true, true)
	check(t, swapped, true)
	check(t, b.Load(), true)

	swapped = b.CompareAndSwap(false, true)
	check(t, swapped, false)
	check(t, b.Load(), true)
}
