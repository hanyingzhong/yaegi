// Code generated by 'github.com/containous/yaegi/extract io/ioutil'. DO NOT EDIT.

// +build go1.13,!go1.14

package stdlib

import (
	"io/ioutil"
	"reflect"
)

func init() {
	Symbols["io/ioutil"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Discard":   reflect.ValueOf(&ioutil.Discard).Elem(),
		"NopCloser": reflect.ValueOf(ioutil.NopCloser),
		"ReadAll":   reflect.ValueOf(ioutil.ReadAll),
		"ReadDir":   reflect.ValueOf(ioutil.ReadDir),
		"ReadFile":  reflect.ValueOf(ioutil.ReadFile),
		"TempDir":   reflect.ValueOf(ioutil.TempDir),
		"TempFile":  reflect.ValueOf(ioutil.TempFile),
		"WriteFile": reflect.ValueOf(ioutil.WriteFile),
	}
}
