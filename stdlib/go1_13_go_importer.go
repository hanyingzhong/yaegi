// Code generated by 'github.com/containous/yaegi/extract go/importer'. DO NOT EDIT.

// +build go1.13,!go1.14

package stdlib

import (
	"go/importer"
	"reflect"
)

func init() {
	Symbols["go/importer"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Default":     reflect.ValueOf(importer.Default),
		"For":         reflect.ValueOf(importer.For),
		"ForCompiler": reflect.ValueOf(importer.ForCompiler),

		// type definitions
		"Lookup": reflect.ValueOf((*importer.Lookup)(nil)),
	}
}
