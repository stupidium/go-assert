package assert_test

import (
	"assert"
	"fmt"
	"runtime"
	"testing"
)

//internal function used to test whether a function panics or not
//returns true if a panic occurs. Could also return true if fn is nil
func testPanic(fn func()) (ret bool) {
	defer func() {
		if err:=recover(); err!=nil {
			ret=true;
		}
	}();
	fn();
	return;
}

//logs an error for testing, prints the %v value for name. name is intended to be a function/string but can be anything
func logError(t *testing.T, name interface{}, msg ...interface{}) {
	_, file, no, ok := runtime.Caller(1)
    if ok {
        t.Errorf("error with %v at %s:%d:%v", name, file, no, msg)
    } else {
		panic(fmt.Errorf("error with %v", name));
	}
}

func TestAssertT(t *testing.T) {
	if testPanic(func() {assert.T(true)}) {
		logError(t, assert.T, "panics when delivered a true value");
	}
	if !testPanic(func() {assert.T(false)}) {
		logError(t, assert.T, "panics does not panic when delivered a true value");
	}
}

func TestAssertEq(t *testing.T) {
	if testPanic(func() {assert.Eq(1,1)}) {
		logError(t, assert.Eq[int], "panics when delivered equal value");
	}
	if !testPanic(func() {assert.Eq(1,2)}) {
		logError(t, assert.Eq[int], "does not panic when delivered inequal value");
	}
}

func TestAsertUeq(t *testing.T) {
	if testPanic(func() {assert.Ueq(1,2)}) {
		logError(t, assert.Ueq[int], "panics when delivered unequal value");
	}
	if !testPanic(func() {assert.Ueq(1,1)}) {
		logError(t, assert.Ueq[int], "does not panic when delivered equal value");
	}
}