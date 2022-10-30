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
	"github.com/jobsearch-demos/mongo-filter-struct/validator"
	"github.com/pkg/errors"
	"reflect"
)

// IScanner is used to scan struct and find fields with tags
// Also it is used to check against several rules e.g. is field compatible with operator
// or is field a relation field
// It is used to find field name, operator and value
type IScanner interface {
	// makeField creates a new filter field from provided struct field
	makeField(reflectionValue reflect.Value,
		reflectionType reflect.StructField,
		parentField *reflect.StructField, index int) (field.IFilterField, error)

	// Scan scans the provided field and returns a list of IFilterField
	Scan(filterStruct interface{}, collection string,
		parentField *reflect.StructField, index int) ([]field.IFilterField, error)
}

type scanner struct {
	operatorMap     operator.IOperatorMap
	validators      []validator.IValidator
	lookupTagName   string
	operatorTagName string
	relationTagName string
}

// Scan scans the provided field and returns a list of IFilterField
// It does not do anything other than scanning the struct and creating a list of IFilterField
// It is responsible for checking the type of the fields and creating respective IFilterField.
func (s *scanner) Scan(filterStruct interface{},
	collection string, parentField *reflect.StructField, index int) ([]field.IFilterField, error) {
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
			fields, err := s.Scan(fieldValue.Interface(), collection, &fieldType, index)
			if err != nil {
				return nil, err
			}
			filterFields = append(filterFields, fields...)

			// increment the index by the number of nested fields
			index += len(fields)
			continue
		}

		// create a new filter field
		fields, err := s.makeField(fieldValue, fieldType, parentField, index)

		// if field could not be created, return error (validation error or unsupported field type)
		if err != nil {
			return nil, err
		}

		// append the field to the list of fields
		filterFields = append(filterFields, fields)

		// increment the index
		index++
	}
	return filterFields, nil
}

// makeField creates a new filter field from provided struct field
// It does not validate the field, it only creates a new filter field
// The only validation it does is validation against tag values correctness
// e.g. if lookup tag value is empty, it takes the struct field name as the lookup
// or if operator tag value is empty, it returns error
// or if the operator tag provided is not supported (does not exist in opmap),
// it returns error
func (s *scanner) makeField(reflectionValue reflect.Value,
	reflectionType reflect.StructField, parentField *reflect.StructField, index int) (field.IFilterField, error) {
	// get the tag value of the field
	lookupTagValue := reflectionType.Tag.Get(s.lookupTagName)
	collectionTagValue := reflectionType.Tag.Get(s.relationTagName)
	operatorTagValue := reflectionType.Tag.Get(s.operatorTagName)

	// if there is a parent field,
	// combine the parent field name and the current field name
	// to get the lookup value
	if parentField != nil {
		if parentLookupName := parentField.Tag.Get(s.lookupTagName); parentLookupName != "" {
			lookupTagValue = parentLookupName + "." + lookupTagValue
		} else {
			lookupTagValue = parentField.Name + "." + lookupTagValue
		}
	}

	// get operator from operator map
	op := s.operatorMap.Get(operatorTagValue)

	// if operator is not found, return error
	if op == nil {
		return nil, errors.Errorf("operator %s is not supported", operatorTagValue)
	}

	// if lookup tag value is empty, get the field name
	if lookupTagValue == "" {
		lookupTagValue = reflectionType.Name
	}

	for _, valid := range s.validators {
		if err := valid.Validate(reflectionValue, reflectionType); err != nil {
			return nil, err
		}
	}

	filterField := field.NewFilterField(
		collectionTagValue,
		reflectionValue.Kind().String(),
		lookupTagValue,
		reflectionValue.Interface(),
		op,
		index,
	)
	return filterField, nil
}

// NewScanner creates new scanner instance with provided options. Factory method.
func NewScanner(operatorMap operator.IOperatorMap,
	validators []validator.IValidator,
	lookupTagName string,
	operatorTagName string, relationTagName string) IScanner {
	return &scanner{
		validators:      validators,
		operatorMap:     operatorMap,
		lookupTagName:   lookupTagName,
		operatorTagName: operatorTagName,
		relationTagName: relationTagName,
	}
}
