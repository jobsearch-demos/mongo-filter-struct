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

type IFilterField interface {
	Validate() error
	Merge(field IFilterField) IFilterField
	Build() bson.D
}

type filterField struct {
	Collection string
	Type       string
	Name       string
	Operator   string
	Value      interface{}
}

func (f filterField) Validate() error {
	//TODO implement me
	panic("implement me")
}

func (f filterField) Merge(field IFilterField) IFilterField {
	//TODO implement me
	panic("implement me")
}

func (f filterField) Build() bson.D {
	//TODO implement me
	panic("implement me")
}

func NewFilterField(collection string, name string, operator string, value interface{}) IFilterField {
	return &filterField{
		Collection: collection,
		Name:       name,
		Operator:   operator,
		Value:      value,
	}
}
