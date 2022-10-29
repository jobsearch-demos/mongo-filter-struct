// License: MIT
// Author: Kamran Valijonov
// Version: 1.0.0
// Date: 2022-10-29
// Description: Mongo Filter Builder
// This tool is used to build bson filter for mongodb based on provided struct.
// Motivation: I was tired of writing bson.M{} for every query and wanted
// something more elegant and easy to use like django-filter.

package filterbuilder

// IScanner is used to scan struct and find fields with tags
// Also it is used to check against several rules e.g. is field compatible with operator
// or is field a relation field
// It is used to find field name, operator and value
type IScanner interface {
	IsCompatible(fieldType string, operator string) bool
}

type scanner struct {
	operatorMap     map[string]string
	lookupTagName   string
	operatorTagName string
	relationTagName string
}

// IsCompatible checks if field is compatible with operator
func (s scanner) IsCompatible(fieldType string, operator string) bool {
	// TODO implement me
	panic("implement me")
}

// NewScanner creates new scanner instance with provided options. Factory method.
func NewScanner(operatorMap map[string]string, lookupTagName string, operatorTagName string, relationTagName string) IScanner {
	return &scanner{
		operatorMap:     operatorMap,
		lookupTagName:   lookupTagName,
		operatorTagName: operatorTagName,
		relationTagName: relationTagName,
	}
}
