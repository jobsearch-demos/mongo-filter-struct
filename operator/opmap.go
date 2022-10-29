// License: GNU General Public License v3.0
// Author: Kamran Valijonov
// Version: 1.0.0
// Date: 2022-10-29
// Description: Mongo Filter Builder
// This tool is used to build bson filter for mongodb based on provided struct.
// Motivation: I was tired of writing bson.M{} for every query and wanted
// something more elegant and easy to use like django-filter.

package operator

type IOperatorMap interface {
	Get(name string) IOperator
	Set(name string, operator IOperator) IOperator
	SetSource(source map[string]IOperator)
}

type operatorMap struct {
	source map[string]IOperator
}

func (o operatorMap) Get(name string) IOperator {
	return o.source[name]
}

func (o operatorMap) Set(name string, operator IOperator) IOperator {
	o.source[name] = operator
	return operator
}

func (o operatorMap) SetSource(source map[string]IOperator) {
	o.source = source
}

func NewOperatorMap() IOperatorMap {
	return &operatorMap{
		source: map[string]IOperator{
			"eq":    EQOperator{},
			"regex": RegexOperator{},
			"lt":    LTOperator{},
			"lte":   LTEOperator{},
			"gt":    GTOperator{},
			"gte":   GTEOperator{},
			"ne":    NEOperator{},
			"in":    INOperator{},
			"nin":   NINOperator{},
		},
	}
}
