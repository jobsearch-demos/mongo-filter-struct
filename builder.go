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

// IFilterBuilder is used to build bson filter for mongodb based on provided struct.
// Its main responsibility is to construct a proper bson.D from provided struct
// by scanning each field of the struct, storing them in a list and then merging
// them into a single bson.D object. It also provides convenience methods like add/remove fields.
// To confirm to SRP, it delegates the responsibility of scanning struct fields to IScanner
// and the responsibility of building bson.D from each field to IFilterField
type IFilterBuilder interface {
	Build(filterStruct interface{}) IFilterBuilder
	AddField(field IFilterField) IFilterBuilder
	AddFields(fields []IFilterField) IFilterBuilder
	RemoveField(field IFilterField) IFilterBuilder
	RemoveFieldByName(name string) IFilterBuilder
	RemoveFieldByIndex(index int) IFilterBuilder
	GetFieldsByName(name string) []IFilterField
	GetFieldByIndex(index int) IFilterField
	GetFields() []IFilterField
	MergeDuplicateFields() IFilterBuilder
	Output() bson.D
}

type filterBuilder struct {
	fields             []IFilterField
	scanner            IScanner
	output             *bson.D
	input              interface{}
	modificationNeeded bool
}

// Build is used to build bson filter for mongodb based on provided struct.
func (f filterBuilder) Build(filterStruct interface{}) IFilterBuilder {
	//TODO implement me
	panic("implement me")
}

// AddField adds a new field to the filter
func (f filterBuilder) AddField(field IFilterField) IFilterBuilder {
	//TODO implement me
	panic("implement me")
}

// AddFields adds a list of fields to the filter
func (f filterBuilder) AddFields(fields []IFilterField) IFilterBuilder {
	//TODO implement me
	panic("implement me")
}

// RemoveField removes a field from the filter
func (f filterBuilder) RemoveField(field IFilterField) IFilterBuilder {
	//TODO implement me
	panic("implement me")
}

// RemoveFieldByName removes a field from the filter by its name
func (f filterBuilder) RemoveFieldByName(name string) IFilterBuilder {
	//TODO implement me
	panic("implement me")
}

// RemoveFieldByIndex removes a field from the filter by its index
func (f filterBuilder) RemoveFieldByIndex(index int) IFilterBuilder {
	//TODO implement me
	panic("implement me")
}

// GetFieldsByName returns a list of fields if their names match the provided one
func (f filterBuilder) GetFieldsByName(name string) []IFilterField {
	//TODO implement me
	panic("implement me")
}

// GetFieldByIndex returns a field by its index
func (f filterBuilder) GetFieldByIndex(index int) IFilterField {
	//TODO implement me
	panic("implement me")
}

// GetFields returns a list of all fields
func (f filterBuilder) GetFields() []IFilterField {
	//TODO implement me
	panic("implement me")
}

// MergeDuplicateFields merges duplicate fields into a single field
func (f filterBuilder) MergeDuplicateFields() IFilterBuilder {
	//TODO implement me
	panic("implement me")
}

// Output returns the final bson.D object
func (f filterBuilder) Output() bson.D {
	//TODO implement me
	panic("implement me")
}

// NewFilterBuilder creates a new instance of IFilterBuilder
func NewFilterBuilder() IFilterBuilder {
	return &filterBuilder{}
}
