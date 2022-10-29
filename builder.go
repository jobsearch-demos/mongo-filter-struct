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

type IFilterBuilder interface {
	Build() bson.D
	AddField(field IFilterField) IFilterBuilder
	AddFields(fields []IFilterField) IFilterBuilder
	RemoveField(field IFilterField) IFilterBuilder
	RemoveFieldByName(name string) IFilterBuilder
	RemoveFieldByIndex(index int) IFilterBuilder
	GetFieldsByName(name string) []IFilterField
	GetFieldByIndex(index int) IFilterField
	GetFields() []IFilterField
}

type filterBuilder struct {
	fields []IFilterField
}

func (f filterBuilder) Build() bson.D {
	//TODO implement me
	panic("implement me")
}

func (f filterBuilder) AddField(field IFilterField) IFilterBuilder {
	//TODO implement me
	panic("implement me")
}

func (f filterBuilder) AddFields(fields []IFilterField) IFilterBuilder {
	//TODO implement me
	panic("implement me")
}

func (f filterBuilder) RemoveField(field IFilterField) IFilterBuilder {
	//TODO implement me
	panic("implement me")
}

func (f filterBuilder) RemoveFieldByName(name string) IFilterBuilder {
	//TODO implement me
	panic("implement me")
}

func (f filterBuilder) RemoveFieldByIndex(index int) IFilterBuilder {
	//TODO implement me
	panic("implement me")
}

func (f filterBuilder) GetFieldsByName(name string) []IFilterField {
	//TODO implement me
	panic("implement me")
}

func (f filterBuilder) GetFieldByIndex(index int) IFilterField {
	//TODO implement me
	panic("implement me")
}

func (f filterBuilder) GetFields() []IFilterField {
	//TODO implement me
	panic("implement me")
}

func NewFilterBuilder() IFilterBuilder {
	return &filterBuilder{}
}
