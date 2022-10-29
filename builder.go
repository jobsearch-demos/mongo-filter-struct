// License: MIT
// Author: Kamran Valijonov
// Version: 1.0.0
// Date: 2022-10-29
// Description: Mongo Filter Builder
// This tool is used to build bson filter for mongodb based on provided struct.
// Motivation: I was tired of writing bson.M{} for every query and wanted
// something more elegant and easy to use like django-filter.

package filterbuilder

import (
	"go.mongodb.org/mongo-driver/bson"
)

// IFilterBuilder is used to build bson filter for mongodb based on provided struct.
// Its main responsibility is to construct a proper bson.D from provided struct
// by scanning each field of the struct, storing them in a list and then merging
// them into a single bson.D object. It also provides convenience methods like add/remove fields.
// To confirm to SRP, it delegates the responsibility of scanning struct fields to IScanner
// and the responsibility of building bson.D from each field to IFilterField
type IFilterBuilder interface {
	SetFields(fields []IFilterField) IFilterBuilder
	AddFields(fields []IFilterField) IFilterBuilder
	AddField(field IFilterField) IFilterBuilder
	GetFields() []IFilterField
	GetFieldByIndex(index int) IFilterField
	GetFieldsByName(name string) []IFilterField
	RemoveField(field IFilterField) IFilterBuilder
	RemoveFieldByIndex(index int) IFilterBuilder
	RemoveFieldByName(name string) IFilterBuilder
	MergeDuplicateFields() IFilterBuilder
	Build() IFilterBuilder
	Output() bson.D
}

type filterBuilder struct {
	fields             []IFilterField
	output             bson.D
	input              interface{}
	modificationNeeded bool
}

func (f *filterBuilder) SetFields(fields []IFilterField) IFilterBuilder {
	f.fields = fields
	return f
}

// Build is used to build bson filter for mongodb based on provided struct.
func (f *filterBuilder) Build() IFilterBuilder {
	// TODO implement me
	panic("implement me")
}

// AddField adds a new field to the filter.
func (f *filterBuilder) AddField(field IFilterField) IFilterBuilder {
	f.fields = append(f.fields, field)
	return f
}

// AddFields adds a list of fields to the filter.
func (f *filterBuilder) AddFields(fields []IFilterField) IFilterBuilder {
	f.fields = append(f.fields, fields...)
	return f
}

// RemoveField removes a field from the filter.
func (f *filterBuilder) RemoveField(field IFilterField) IFilterBuilder {
	f.fields = append(f.fields[:field.GetIndex()], f.fields[field.GetIndex()+1:]...)
	return f
}

// RemoveFieldByName removes a field from the filter by its name.
func (f *filterBuilder) RemoveFieldByName(name string) IFilterBuilder {
	for idx, field := range f.fields {
		if field.GetName() == name {
			f.fields = append(f.fields[:idx], f.fields[idx+1:]...)
		}
	}
	return f
}

// RemoveFieldByIndex removes a field from the filter by its index.
func (f *filterBuilder) RemoveFieldByIndex(index int) IFilterBuilder {
	f.fields = append(f.fields[:index], f.fields[index+1:]...)
	return f
}

// GetFieldsByName returns a list of fields if their names match the provided one.
func (f *filterBuilder) GetFieldsByName(name string) []IFilterField {
	var fields []IFilterField
	for _, field := range f.fields {
		if field.GetName() == name {
			fields = append(fields, field)
		}
	}
	return fields
}

// GetFieldByIndex returns a field by its index.
func (f *filterBuilder) GetFieldByIndex(index int) IFilterField {
	return f.fields[index]
}

// GetFields returns a list of all fields.
func (f *filterBuilder) GetFields() []IFilterField {
	return f.fields
}

// MergeDuplicateFields merges duplicate fields into a single field
func (f *filterBuilder) MergeDuplicateFields() IFilterBuilder {
	for _, field := range f.fields {
		for _, field2 := range f.fields {
			if field.GetName() == field2.GetName() && field.GetIndex() != field2.GetIndex() {
				field.Merge(field2)
				f.RemoveField(field2)
			}
		}
	}
	return f
}

// Output returns the final bson.D object
func (f *filterBuilder) Output() bson.D {
	return f.output
}

// NewFilterBuilder creates a new instance of IFilterBuilder
func NewFilterBuilder() IFilterBuilder {
	return &filterBuilder{
		fields:             []IFilterField{},
		output:             bson.D{},
		modificationNeeded: false,
	}
}
