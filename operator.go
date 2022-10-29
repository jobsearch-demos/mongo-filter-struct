package filterbuilder

import "reflect"

type IOperator interface {
	// IsCompatible returns true if the operator is compatible with the field type
	// e.g. EQOperator.IsCompatible(reflect.String) returns true because
	// you can compare two strings with == operator.
	IsCompatible(fieldType reflect.Kind) bool
}

type EQOperator struct{}

func (o EQOperator) IsCompatible(fieldType reflect.Kind) bool {
	return fieldType == reflect.String ||
		fieldType == reflect.Int ||
		fieldType == reflect.Int16 ||
		fieldType == reflect.Int32 ||
		fieldType == reflect.Int64 ||
		fieldType == reflect.Float32 ||
		fieldType == reflect.Float64 ||
		fieldType == reflect.Bool
}

type RegexOperator struct{}

func (o RegexOperator) IsCompatible(fieldType reflect.Kind) bool {
	return fieldType == reflect.String
}

type LTOperator struct{}

func (o LTOperator) IsCompatible(fieldType reflect.Kind) bool {
	return fieldType == reflect.Int ||
		fieldType == reflect.Int16 ||
		fieldType == reflect.Int32 ||
		fieldType == reflect.Int64 ||
		fieldType == reflect.Float64 ||
		fieldType == reflect.Float32
}

type LTEOperator struct{}

func (o LTEOperator) IsCompatible(fieldType reflect.Kind) bool {
	return fieldType == reflect.Int ||
		fieldType == reflect.Int16 ||
		fieldType == reflect.Int32 ||
		fieldType == reflect.Int64 ||
		fieldType == reflect.Float64 ||
		fieldType == reflect.Float32
}

type GTOperator struct{}

func (o GTOperator) IsCompatible(fieldType reflect.Kind) bool {
	return fieldType == reflect.Int ||
		fieldType == reflect.Int16 ||
		fieldType == reflect.Int32 ||
		fieldType == reflect.Int64 ||
		fieldType == reflect.Float64 ||
		fieldType == reflect.Float32
}

type GTEOperator struct{}

func (o GTEOperator) IsCompatible(fieldType reflect.Kind) bool {
	return fieldType == reflect.Int ||
		fieldType == reflect.Int16 ||
		fieldType == reflect.Int32 ||
		fieldType == reflect.Int64 ||
		fieldType == reflect.Float64 ||
		fieldType == reflect.Float32
}

type NEOperator struct{}

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

type INOperator struct{}

func (o INOperator) IsCompatible(fieldType reflect.Kind) bool {
	return fieldType == reflect.Slice
}

type NINOperator struct{}

func (o NINOperator) IsCompatible(fieldType reflect.Kind) bool {
	return fieldType == reflect.Slice
}
