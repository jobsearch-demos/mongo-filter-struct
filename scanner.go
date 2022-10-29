// License: MIT
// Author: Kamran Valijonov
// Version: 1.0.0
// Date: 2022-10-29
// Description: Mongo Filter Builder
// This tool is used to build bson filter for mongodb based on provided struct.
// Motivation: I was tired of writing bson.M{} for every query and wanted
// something more elegant and easy to use like django-filter.

package filterbuilder

import "reflect"

// IScanner is used to scan struct and find fields with tags
// Also it is used to check against several rules e.g. is field compatible with operator
// or is field a relation field
// It is used to find field name, operator and value
type IScanner interface {
	// Validate runs the validations against the provided struct fields
	validate(reflectionValue reflect.Value, reflectionType reflect.Type) bool

	// MakeField creates a new filter field from provided struct field
	makeField() IFilterField

	// Scan scans the provided field and returns IFilterField after validation
	Scan(filterStruct interface{}) []IFilterField
}

type scanner struct {
	operatorMap     IOperatorMap
	lookupTagName   string
	operatorTagName string
	relationTagName string
}

func (s *scanner) Scan(fileterStruct interface{}) []IFilterField {
	panic("implement me")
}

func (s *scanner) validate(reflectionValue reflect.Value, reflectionType reflect.Type) bool {
	//TODO implement me
	panic("implement me")
}

func (s *scanner) makeField() IFilterField {
	//TODO implement me
	panic("implement me")
}

// NewScanner creates new scanner instance with provided options. Factory method.
func NewScanner(operatorMap IOperatorMap, lookupTagName string,
	operatorTagName string, relationTagName string) IScanner {
	return &scanner{
		operatorMap:     operatorMap,
		lookupTagName:   lookupTagName,
		operatorTagName: operatorTagName,
		relationTagName: relationTagName,
	}
}
