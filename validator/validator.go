package validator

import (
	"github.com/jobsearch-demos/mongo-filter-struct/operator"
	"github.com/pkg/errors"
	"reflect"
)

type IValidator interface {
	Validate(reflectionValue reflect.Value,
		reflectionType reflect.StructField) error
}

type operatorValidator struct {
	opMap           operator.IOperatorMap
	operatorTagName string
}

func (v *operatorValidator) Validate(reflectionValue reflect.Value,
	reflectionType reflect.StructField) error {
	// get operator tag value
	operatorTagValue := reflectionType.Tag.Get(v.operatorTagName)

	// if operator tag value is empty, return error
	if operatorTagValue == "" {
		return errors.Errorf("operator tag value is empty")
	}

	// get operator from operator map
	op := v.opMap.Get(operatorTagValue)

	// if operator tag value is not in the operator map, return error
	if op == nil {
		return errors.Errorf("operator %s is not supported", operatorTagValue)
	}

	// if operator is not compatible with the field, return error
	if !op.IsCompatible(reflectionValue.Kind()) {
		return errors.Errorf("operator %s is not compatible with field %s of type %s", operatorTagValue, reflectionType.Name, reflectionValue.Kind())
	}
	return nil
}

func NewOperatorValidator(opMap operator.IOperatorMap, operatorTagName string) IValidator {
	return &operatorValidator{
		operatorTagName: operatorTagName,
		opMap:           opMap,
	}
}
