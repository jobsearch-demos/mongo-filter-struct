package validator_test

import (
	"github.com/jobsearch-demos/mongo-filter-struct/operator"
	"github.com/jobsearch-demos/mongo-filter-struct/validator"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type testStrEQ struct {
	Name string `json:"name" bson:"name" filter:"name" operator:"eq"`
}

type testStrNE struct {
	Name string `json:"name" bson:"name" filter:"name" operator:"ne"`
}

type testStrIN struct {
	Name string `json:"name" bson:"name" filter:"name" operator:"in"`
}

type testStrNIN struct {
	Name string `json:"name" bson:"name" filter:"name" operator:"nin"`
}

type testStrRegex struct {
	Name string `json:"name" bson:"name" filter:"name" operator:"regex"`
}

type testStrLT struct {
	Name string `json:"name" bson:"name" filter:"name" operator:"lt"`
}

type testStrLTE struct {
	Name string `json:"name" bson:"name" filter:"name" operator:"lte"`
}

type testStrGT struct {
	Name string `json:"name" bson:"name" filter:"name" operator:"gt"`
}

type testStrGTE struct {
	Name string `json:"name" bson:"name" filter:"name" operator:"gte"`
}

type testStrNotSupported struct {
	Name string `json:"name" bson:"name" filter:"name" operator:"not_supported"`
}

type testIntEQ struct {
	Age int `json:"age" bson:"age" filter:"age" operator:"eq"`
}

type testIntNE struct {
	Age int `json:"age" bson:"age" filter:"age" operator:"ne"`
}

type testIntIN struct {
	Age int `json:"age" bson:"age" filter:"age" operator:"in"`
}

type testIntNIN struct {
	Age int `json:"age" bson:"age" filter:"age" operator:"nin"`
}

type testIntRegex struct {
	Age int `json:"age" bson:"age" filter:"age" operator:"regex"`
}

type testIntLT struct {
	Age int `json:"age" bson:"age" filter:"age" operator:"lt"`
}

type testIntLTE struct {
	Age int `json:"age" bson:"age" filter:"age" operator:"lte"`
}

type testIntGT struct {
	Age int `json:"age" bson:"age" filter:"age" operator:"gt"`
}

type testIntGTE struct {
	Age int `json:"age" bson:"age" filter:"age" operator:"gte"`
}

type testIntNotSupported struct {
	Age int `json:"age" bson:"age" filter:"age" operator:"not_supported"`
}

type testFloatEQ struct {
	Age float64 `json:"age" bson:"age" filter:"age" operator:"eq"`
}

type testFloatNE struct {
	Age float64 `json:"age" bson:"age" filter:"age" operator:"ne"`
}

type testFloatIN struct {
	Age float64 `json:"age" bson:"age" filter:"age" operator:"in"`
}

type testFloatNIN struct {
	Age float64 `json:"age" bson:"age" filter:"age" operator:"nin"`
}

type testFloatRegex struct {
	Age float64 `json:"age" bson:"age" filter:"age" operator:"regex"`
}

type testFloatLT struct {
	Age float64 `json:"age" bson:"age" filter:"age" operator:"lt"`
}

type testFloatLTE struct {
	Age float64 `json:"age" bson:"age" filter:"age" operator:"lte"`
}

type testFloatGT struct {
	Age float64 `json:"age" bson:"age" filter:"age" operator:"gt"`
}

type testFloatGTE struct {
	Age float64 `json:"age" bson:"age" filter:"age" operator:"gte"`
}

type testFloatNotSupported struct {
	Age float64 `json:"age" bson:"age" filter:"age" operator:"not_supported"`
}

type testBoolEQ struct {
	Active bool `json:"active" bson:"active" filter:"active" operator:"eq"`
}

type testBoolNE struct {
	Active bool `json:"active" bson:"active" filter:"active" operator:"ne"`
}

type testBoolIN struct {
	Active bool `json:"active" bson:"active" filter:"active" operator:"in"`
}

type testBoolNIN struct {
	Active bool `json:"active" bson:"active" filter:"active" operator:"nin"`
}

type testBoolRegex struct {
	Active bool `json:"active" bson:"active" filter:"active" operator:"regex"`
}

type testBoolLT struct {
	Active bool `json:"active" bson:"active" filter:"active" operator:"lt"`
}

type testBoolLTE struct {
	Active bool `json:"active" bson:"active" filter:"active" operator:"lte"`
}

type testBoolGT struct {
	Active bool `json:"active" bson:"active" filter:"active" operator:"gt"`
}

type testBoolGTE struct {
	Active bool `json:"active" bson:"active" filter:"active" operator:"gte"`
}

type testBoolNotSupported struct {
	Active bool `json:"active" bson:"active" filter:"active" operator:"not_supported"`
}

func TestValidator(t *testing.T) {
	tests := []struct {
		Name      string
		Structure interface{}
		WantErr   bool
	}{{
		Name: "String can use eq operator",
		Structure: testStrEQ{
			Name: "test",
		},
		WantErr: false},
		{
			Name: "String can use ne operator",
			Structure: testStrNE{
				Name: "test",
			},
			WantErr: false,
		},
		{
			Name: "String can use in operator",
			Structure: testStrIN{
				Name: "test",
			},
			WantErr: false,
		},
		{
			Name: "String can use nin operator",
			Structure: testStrNIN{
				Name: "test",
			},
			WantErr: false,
		},
		{
			Name: "String can use regex operator",
			Structure: testStrRegex{
				Name: "test",
			},
			WantErr: false,
		},
		{
			Name: "String cannot use lt operator",
			Structure: testStrLT{
				Name: "test",
			},
			WantErr: true,
		},
		{
			Name: "String cannot use lte operator",
			Structure: testStrLTE{
				Name: "test",
			},
			WantErr: true,
		},
		{
			Name: "String cannot use gt operator",
			Structure: testStrGT{
				Name: "test",
			},
			WantErr: true,
		},
		{
			Name: "String cannot use gte operator",
			Structure: testStrGTE{
				Name: "test",
			},
			WantErr: true,
		},
		{
			Name: "String cannot use not supported operator",
			Structure: testStrNotSupported{
				Name: "test",
			},
			WantErr: true,
		},
		{
			Name: "Int can use eq operator",
			Structure: testIntEQ{
				Age: 10,
			},
			WantErr: false,
		},
		{
			Name: "Int can use ne operator",
			Structure: testIntNE{
				Age: 10,
			},
			WantErr: false,
		},
		{
			Name: "Int cannot use in operator",
			Structure: testIntIN{
				Age: 10,
			},
			WantErr: true,
		},
		{
			Name: "Int cannot use nin operator",
			Structure: testIntNIN{
				Age: 10,
			},
			WantErr: true,
		},
		{
			Name: "Int cannot use regex operator",
			Structure: testIntRegex{
				Age: 10,
			},
			WantErr: true,
		},
		{
			Name: "Int can use lt operator",
			Structure: testIntLT{
				Age: 10,
			},
			WantErr: false,
		},
		{
			Name: "Int can use lte operator",
			Structure: testIntLTE{
				Age: 10,
			},
			WantErr: false,
		},
		{
			Name: "Int can use gt operator",
			Structure: testIntGT{
				Age: 10,
			},
			WantErr: false,
		},
		{
			Name: "Int can use gte operator",
			Structure: testIntGTE{
				Age: 10,
			},
			WantErr: false,
		},
		{
			Name: "Int cannot use not supported operator",
			Structure: testIntNotSupported{
				Age: 10,
			},
			WantErr: true,
		},
		{
			Name: "Bool can use eq operator",
			Structure: testBoolEQ{
				Active: true,
			},
			WantErr: false,
		},
		{
			Name: "Bool can use ne operator",
			Structure: testBoolNE{
				Active: true,
			},
			WantErr: false,
		},
		{
			Name: "Bool cannot use in operator",
			Structure: testBoolIN{
				Active: true,
			},
			WantErr: true,
		},
		{
			Name: "Bool cannot use nin operator",
			Structure: testBoolNIN{
				Active: true,
			},
			WantErr: true,
		},
		{
			Name: "Bool cannot use regex operator",
			Structure: testBoolRegex{
				Active: true,
			},
			WantErr: true,
		},
		{
			Name: "Bool cannot use lt operator",
			Structure: testBoolLT{
				Active: true,
			},
			WantErr: true,
		},
		{
			Name: "Bool cannot use lte operator",
			Structure: testBoolLTE{
				Active: true,
			},
			WantErr: true,
		},
		{
			Name: "Bool cannot use gt operator",
			Structure: testBoolGT{
				Active: true,
			},
			WantErr: true,
		},
		{
			Name: "Bool cannot use gte operator",
			Structure: testBoolGTE{
				Active: true,
			},
			WantErr: true,
		},
		{
			Name: "Bool cannot use not supported operator",
			Structure: testBoolNotSupported{
				Active: true,
			},
			WantErr: true,
		},
		{
			Name: "Float can use eq operator",
			Structure: testFloatEQ{
				Age: 1.75,
			},
			WantErr: false,
		},
		{
			Name: "Float can use ne operator",
			Structure: testFloatNE{
				Age: 1.75,
			},
			WantErr: false,
		},
		{
			Name: "Float cannot use in operator",
			Structure: testFloatIN{
				Age: 1.75,
			},
			WantErr: true,
		},
		{
			Name: "Float cannot use nin operator",
			Structure: testFloatNIN{
				Age: 1.75,
			},
			WantErr: true,
		},
		{
			Name: "Float cannot use regex operator",
			Structure: testFloatRegex{
				Age: 1.75,
			},
			WantErr: true,
		},
		{
			Name: "Float can use lt operator",
			Structure: testFloatLT{
				Age: 1.75,
			},
			WantErr: false,
		},
		{
			Name: "Float can use lte operator",
			Structure: testFloatLTE{
				Age: 1.75,
			},
			WantErr: false,
		},
		{
			Name: "Float can use gt operator",
			Structure: testFloatGT{
				Age: 1.75,
			},
			WantErr: false,
		},
		{
			Name: "Float can use gte operator",
			Structure: testFloatGTE{
				Age: 1.75,
			},
			WantErr: false,
		},
		{
			Name: "Float cannot use not supported operator",
			Structure: testFloatNotSupported{
				Age: 1.75,
			},
			WantErr: true,
		},
	}

	opValidator := validator.NewOperatorValidator(operator.NewOperatorMap(), "operator")
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			rv, rt := reflect.ValueOf(test.Structure), reflect.TypeOf(test.Structure)

			for i := 0; i < rt.NumField(); i++ {
				field := rt.Field(i)
				value := rv.Field(i)

				err := opValidator.Validate(value, field)

				if test.WantErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
			}
		})
	}
}
