// License: MIT
// Author: Kamran Valijonov
// Version: 1.0.0
// Date: 2022-10-29
// Description: Mongo Filter Builder
// This tool is used to build bson filter for mongodb based on provided struct.
// Motivation: I was tired of writing bson.M{} for every query and wanted
// something more elegant and easy to use like django-filter.

package filterbuilder

import "go.mongodb.org/mongo-driver/bson"

// IFilterField is used to build bson filter for mongodb based on provided struct.
// Its main responsibility is to construct a proper bson.D from a provided single struct field.
// In case of field being duplicated, it merges them into a single bson.D object.
type IFilterField interface {
	Merge(field IFilterField) IFilterField
	GetIndex() int
	GetName() string
	GetCollection() string
	GetType() string
	GetOperator() IOperator
	GetLevel() int
	GetValue() interface{}
	Build() IFilterField
	Output() bson.D
}

type filterField struct {
	collection string
	fieldType  string
	name       string
	value      interface{}
	operator   IOperator
	level      int
	index      int
	output     bson.D
}

func (f *filterField) GetName() string {
	return f.name
}

func (f *filterField) GetCollection() string {
	return f.collection
}

func (f *filterField) GetType() string {
	return f.fieldType
}

func (f *filterField) GetOperator() IOperator {
	return f.operator
}

func (f *filterField) GetLevel() int {
	return f.level
}

func (f *filterField) GetValue() interface{} {
	return f.value
}

func (f *filterField) GetIndex() int {
	return f.index
}

// Merge merges two filter fields into a single one
func (f *filterField) Merge(field IFilterField) IFilterField {
	panic("implement me")
}

// Build builds a bson.D from a single filter field
func (f *filterField) Build() IFilterField {
	//TODO implement me
	panic("implement me")
}

// Output returns the output of the filter field
func (f *filterField) Output() bson.D {
	return f.output
}

// NewFilterField creates a new filter field
func NewFilterField(collection string, fieldType string, name string,
	value interface{}, operator IOperator, level int, index int) IFilterField {
	return &filterField{
		collection: collection,
		fieldType:  fieldType,
		name:       name,
		value:      value,
		operator:   operator,
		level:      level,
		index:      index,
	}
}
