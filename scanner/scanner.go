// License: GNU General Public License v3.0
// Author: Kamran Valijonov
// Version: 1.0.0
// Date: 2022-10-29
// Description: Mongo Filter Builder
// This tool is used to build bson filter for mongodb based on provided struct.
// Motivation: I was tired of writing bson.M{} for every query and wanted
// something more elegant and easy to use like django-filter.

package scanner

import (
	"github.com/jobsearch-demos/mongo-filter-struct/field"
	"github.com/jobsearch-demos/mongo-filter-struct/operator"
	"github.com/pkg/errors"
	"reflect"
)

// IScanner is used to scan struct and find fields with tags
// Also it is used to check against several rules e.g. is field compatible with operator
// or is field a relation field
// It is used to find field name, operator and value
type IScanner interface {
	// Validate runs the validations against the provided struct fields
	validate(reflectionValue reflect.Value, reflectionType reflect.StructField) error

	// MakeField creates a new filter field from provided struct field
	makeField(reflectionValue reflect.Value,
		reflectionType reflect.StructField) (field.IFilterField, error)

	// Scan scans the provided field and returns IFilterField after validation
	Scan(filterStruct interface{}) ([]field.IFilterField, error)
}

type scanner struct {
	operatorMap     operator.IOperatorMap
	lookupTagName   string
	operatorTagName string
	relationTagName string
}

// Scan scans the provided field and returns a list of IFilterField after running the validations
func (s *scanner) Scan(filterStruct interface{}) ([]field.IFilterField, error) {
	// prepare the list of fields to return
	var filterFields []field.IFilterField

	// get the reflection value and type of the provided struct
	rv, rt := reflect.ValueOf(filterStruct), reflect.TypeOf(filterStruct)

	// if the provided struct is a pointer,
	// get the value and type of the struct
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
		rt = rt.Elem()
	}

	// if the provided struct is not a struct, return error
	if rv.Kind() != reflect.Struct {
		return nil, errors.Errorf("filterStruct has to be a struct")
	}

	// iterate over the fields of the provided struct
	for i := 0; i < rv.NumField(); i++ {
		// get the reflection value and type of the field
		fieldValue := rv.Field(i)
		fieldType := rt.Field(i)

		// if the field is a pointer, get the value and type of the field
		if fieldValue.Kind() == reflect.Ptr {
			fieldValue = fieldValue.Elem()
		}

		// if the field is a struct, recursively call Scan
		if fieldValue.Kind() == reflect.Struct {
			fields, err := s.Scan(fieldValue.Interface())
			if err != nil {
				return nil, err
			}
			filterFields = append(filterFields, fields...)
		}

		// create a new filter field
		fields, err := s.makeField(fieldValue, fieldType)

		// if field could not be created, return error (validation error or unsupported field type)
		if err != nil {
			return nil, err
		}

		// append the field to the list of fields
		filterFields = append(filterFields, fields)
	}
	return filterFields, nil
}

func (s *scanner) validate(reflectionValue reflect.Value, reflectionType reflect.StructField) error {
	// get operator tag value
	operatorTagValue := reflectionType.Tag.Get(s.operatorTagName)

	// if operator tag value is empty, return error
	if operatorTagValue == "" {
		return errors.Errorf("operator tag value is empty")
	}

	// get operator from operator map
	op := s.operatorMap.Get(operatorTagValue)

	// if operator tag value is not in the operator map, return error
	if op == nil {
		return errors.Errorf("operator %s is not supported", operatorTagValue)
	}

	// if operator is not compatible with the field, return error
	if !op.IsCompatible(reflectionValue.Kind()) {
		return errors.Errorf("operator %s is not compatible with field %s", operatorTagValue, reflectionType.Name)
	}
	return nil
}

func (s *scanner) makeField(reflectionValue reflect.Value, reflectionType reflect.StructField) (field.IFilterField, error) {
	// validate the field
	if err := s.validate(reflectionValue, reflectionType); err != nil {
		return nil, errors.Errorf("field %s is not valid", reflectionType.Name)
	}

	// get the tag value of the field
	lookupTagValue := reflectionType.Tag.Get(s.lookupTagName)
	collectionTagValue := reflectionType.Tag.Get(s.relationTagName)

	// if lookup tag value is empty, get the field name
	if lookupTagValue == "" {
		lookupTagValue = reflectionType.Name
	}

	// TODO: fix the logic here
	filterField := field.NewFilterField(
		collectionTagValue,
		reflectionValue.Kind().String(),
		reflectionType.Name,
		reflectionValue.Interface(),
		nil,
		0,
		0,
	)
	return filterField, nil
}

// NewScanner creates new scanner instance with provided options. Factory method.
func NewScanner(operatorMap operator.IOperatorMap, lookupTagName string,
	operatorTagName string, relationTagName string) IScanner {
	return &scanner{
		operatorMap:     operatorMap,
		lookupTagName:   lookupTagName,
		operatorTagName: operatorTagName,
		relationTagName: relationTagName,
	}
}
