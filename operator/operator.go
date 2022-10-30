// License: GNU General Public License v3.0
// Author: Kamran Valijonov
// Version: 1.0.0
// Date: 2022-10-29
// Description: Mongo Filter Builder
// This tool is used to build bson filter for mongodb based on provided struct.
// Motivation: I was tired of writing bson.M{} for every query and wanted
// something more elegant and easy to use like django-filter.

package operator

import "reflect"

type IOperator interface {
	// IsCompatible returns true if the operator is compatible with the field type
	// e.g. EQOperator.IsCompatible(reflect.String) returns true because
	// you can compare two strings with == operator.
	IsCompatible(fieldType reflect.Kind) bool
	ExternalName() string
}

// EQOperator is the equal operator (==)
// Compatible types: string, int, int16, int32, int64, float32, float64, bool
type EQOperator struct{}

func (o EQOperator) IsCompatible(fieldType reflect.Kind) bool {
	return fieldType == reflect.String ||
		fieldType == reflect.Int ||
		fieldType == reflect.Int16 ||
		fieldType == reflect.Int32 ||
		fieldType == reflect.Int64 ||
		fieldType == reflect.Uint ||
		fieldType == reflect.Uint16 ||
		fieldType == reflect.Uint32 ||
		fieldType == reflect.Uint64 ||
		fieldType == reflect.Float32 ||
		fieldType == reflect.Float64 ||
		fieldType == reflect.Bool ||
		fieldType == reflect.Slice ||
		fieldType == reflect.Array
}

func (o EQOperator) ExternalName() string {
	return "eq"
}

// RegexOperator is the regex operator
// Compatible types: string
type RegexOperator struct{}

func (o RegexOperator) IsCompatible(fieldType reflect.Kind) bool {
	return fieldType == reflect.String
}

func (o RegexOperator) ExternalName() string {
	return "regex"
}

// LTOperator is the less than operator (<)
// Compatible types: int, int16, int32, int64, float32, float64
type LTOperator struct{}

func (o LTOperator) ExternalName() string {
	return "lt"
}

func (o LTOperator) IsCompatible(fieldType reflect.Kind) bool {
	return fieldType == reflect.Int ||
		fieldType == reflect.Int16 ||
		fieldType == reflect.Int32 ||
		fieldType == reflect.Int64 ||
		fieldType == reflect.Float64 ||
		fieldType == reflect.Float32
}

func (o LTEOperator) ExternalName() string {
	return "lte"
}

// LTEOperator is the less than or equal operator (<=)
// Compatible types: int, int16, int32, int64, float32, float64
type LTEOperator struct{}

func (o LTEOperator) IsCompatible(fieldType reflect.Kind) bool {
	return fieldType == reflect.Int ||
		fieldType == reflect.Int16 ||
		fieldType == reflect.Int32 ||
		fieldType == reflect.Int64 ||
		fieldType == reflect.Float64 ||
		fieldType == reflect.Float32
}

func (o GTOperator) ExternalName() string {
	return "gt"
}

// GTOperator is the greater than operator (>)
// Compatible types: int, int16, int32, int64, float32, float64
type GTOperator struct{}

func (o GTOperator) IsCompatible(fieldType reflect.Kind) bool {
	return fieldType == reflect.Int ||
		fieldType == reflect.Int16 ||
		fieldType == reflect.Int32 ||
		fieldType == reflect.Int64 ||
		fieldType == reflect.Float64 ||
		fieldType == reflect.Float32
}

func (o GTEOperator) ExternalName() string {
	return "gt"
}

// GTEOperator is the greater than or equal operator (>=)
// Compatible types: int, int16, int32, int64, float32, float64
type GTEOperator struct{}

func (o GTEOperator) IsCompatible(fieldType reflect.Kind) bool {
	return fieldType == reflect.Int ||
		fieldType == reflect.Int16 ||
		fieldType == reflect.Int32 ||
		fieldType == reflect.Int64 ||
		fieldType == reflect.Float64 ||
		fieldType == reflect.Float32
}

// NEOperator is the not equal operator (!=)
// Compatible types: string, int, int16, int32, int64, float32, float64, bool
type NEOperator struct{}

func (o NEOperator) ExternalName() string {
	return "ne"
}

func (o NEOperator) IsCompatible(fieldType reflect.Kind) bool {
	return fieldType == reflect.String ||
		fieldType == reflect.Int ||
		fieldType == reflect.Int16 ||
		fieldType == reflect.Int32 ||
		fieldType == reflect.Int64 ||
		fieldType == reflect.Float32 ||
		fieldType == reflect.Float64 ||
		fieldType == reflect.Bool
}

// INOperator is the in operator
// Compatible types: slice
type INOperator struct{}

func (o INOperator) ExternalName() string {
	return "in"
}

func (o INOperator) IsCompatible(fieldType reflect.Kind) bool {
	return fieldType == reflect.Slice || fieldType == reflect.Array || fieldType == reflect.String
}

// NINOperator is the not in operator
// Compatible types: slice
type NINOperator struct{}

func (o NINOperator) ExternalName() string {
	return "nin"
}

func (o NINOperator) IsCompatible(fieldType reflect.Kind) bool {
	return fieldType == reflect.Slice || fieldType == reflect.Array || fieldType == reflect.String
}
