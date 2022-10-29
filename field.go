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
	Build() IFilterField
	Output() bson.D
}

type filterField struct {
	collection string
	fieldType  string
	name       string
	operator   string
	value      interface{}
	output     *bson.D
}

// Merge merges two filter fields into a single one
func (f filterField) Merge(field IFilterField) IFilterField {
	//TODO implement me
	panic("implement me")
}

// Build builds a bson.D from a single filter field
func (f filterField) Build() IFilterField {
	//TODO implement me
	panic("implement me")
}

// Output returns the output of the filter field
func (f filterField) Output() bson.D {
	//TODO implement me
	panic("implement me")
}

func NewFilterField(collection string, name string, operator string, fieldType string, value interface{}) IFilterField {
	return &filterField{
		collection: collection,
		fieldType:  fieldType,
		name:       name,
		operator:   operator,
		value:      value,
	}
}
