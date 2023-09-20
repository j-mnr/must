// Package must provides utility functions for testing. The standard library's
// [testing] package is enough, but if you find yourself writing the same lines
// several times it's better to put them in their own function and/or properly
// name module.
package must

import (
	"reflect"
	"testing"
)

// NoErr fails the test immediately if the passed in error is not nil.
func NoErr(tb testing.TB, err error) {
	tb.Helper()
	if err != nil {
		tb.Fatal("Error is not nil:\n", err.Error())
	}
}

// Equal checks if want and got have all the same values besides the pointers in
// both values. If a pointer is encountered the value will be checked for nil.
// If not nil, it is dereferenced to check equality instead of comparing pointer
// addresses, as this is usually never what is wanted.
func Equal[C comparable](tb testing.TB, want, got C) {
	tb.Helper()
	if isNil(want) && isNil(got) {
		return
	}
	if want != got {
		tb.Fatalf("Values differ:\nwant: %+v\n got: %+v\n", want, got)
	}
}

func isNil(x any) bool {
	if x == nil {
		return true
	}
	v := reflect.ValueOf(x)
	kind := v.Kind()
	if reflect.Chan <= kind && kind <= reflect.Slice || 
		kind == reflect.UnsafePointer {
		return v.IsNil()
	}
	return false
}

func printNil(want, got any) string {
	if isNil(want) {
		return ""
	}
	return ""
}
