package scanner

import (
	"github.com/jobsearch-demos/mongo-filter-struct/field"
	"github.com/jobsearch-demos/mongo-filter-struct/operator"
	"github.com/jobsearch-demos/mongo-filter-struct/validator"
	"reflect"
	"testing"
)

type TestStructWithIntPointer struct {
	Age *int `json:"age" bson:"age" filter:"age" operator:"eq"`
}

type TestStructWithInt struct {
	Age int `json:"age" bson:"age" filter:"age" operator:"eq"`
}

type TestStructWithFloatPointer struct {
	Age *float64 `json:"age" bson:"age" filter:"age" operator:"eq"`
}

type TestStructWithFloat struct {
	Age float64 `json:"age" bson:"age" filter:"age" operator:"eq"`
}

type TestStructWithBoolPointer struct {
	Active *bool `json:"active" bson:"active" filter:"active" operator:"eq"`
}

type TestStructWithBool struct {
	Active bool `json:"active" bson:"active" filter:"active" operator:"eq"`
}

type TestStructWithStringPointer struct {
	Name *string `json:"name" bson:"name" filter:"name" operator:"eq"`
}

type TestStructWithString struct {
	Name string `json:"name" bson:"name" filter:"name" operator:"eq"`
}

type TestStructWithSlicePointer struct {
	Names *[]string `json:"names" bson:"names" filter:"names" operator:"eq"`
}

type TestStructWithSlice struct {
	Names []string `json:"names" bson:"names" filter:"names" operator:"eq"`
}

type TestStructWithInt16Pointer struct {
	Age *int16 `json:"age" bson:"age" filter:"age" operator:"eq"`
}

type TestStructWithInt16 struct {
	Age int16 `json:"age" bson:"age" filter:"age" operator:"eq"`
}

type TestStructWithInt32Pointer struct {
	Age *int32 `json:"age" bson:"age" filter:"age" operator:"eq"`
}

type TestStructWithInt32 struct {
	Age int32 `json:"age" bson:"age" filter:"age" operator:"eq"`
}

type TestStructWithInt64Pointer struct {
	Age *int64 `json:"age" bson:"age" filter:"age" operator:"eq"`
}

type TestStructWithInt64 struct {
	Age int64 `json:"age" bson:"age" filter:"age" operator:"eq"`
}

type TestStructWithUintPointer struct {
	Age *uint `json:"age" bson:"age" filter:"age" operator:"eq"`
}

type TestStructWithUint struct {
	Age uint `json:"age" bson:"age" filter:"age" operator:"eq"`
}

type TestStructWithUint16Pointer struct {
	Age *uint16 `json:"age" bson:"age" filter:"age" operator:"eq"`
}

type TestStructWithUint16 struct {
	Age uint16 `json:"age" bson:"age" filter:"age" operator:"eq"`
}

type TestStructWithUint32Pointer struct {
	Age *uint32 `json:"age" bson:"age" filter:"age" operator:"eq"`
}

type TestStructWithUint32 struct {
	Age uint32 `json:"age" bson:"age" filter:"age" operator:"eq"`
}

type TestStructWithUint64Pointer struct {
	Age *uint64 `json:"age" bson:"age" filter:"age" operator:"eq"`
}

type TestStructWithUint64 struct {
	Age uint64 `json:"age" bson:"age" filter:"age" operator:"eq"`
}

type TestStructWithFloat32Pointer struct {
	Age *float32 `json:"age" bson:"age" filter:"age" operator:"eq"`
}

type TestStructWithFloat32 struct {
	Age float32 `json:"age" bson:"age" filter:"age" operator:"eq"`
}

type TestStructWithFloat64Pointer struct {
	Age *float64 `json:"age" bson:"age" filter:"age" operator:"eq"`
}

type TestStructWithFloat64 struct {
	Age float64 `json:"age" bson:"age" filter:"age" operator:"eq"`
}

type TestStructWithNestedStructPtr struct {
	User *TestStructWithInt `json:"user" bson:"user" filter:"user" operator:"eq"`
}

type TestStructWithNestedStruct struct {
	User TestStructWithInt `json:"user" bson:"user" filter:"user" operator:"eq"`
}

type TestStructWithCollectionName struct {
	Age int `json:"age" bson:"age" filter:"age" operator:"eq" collection:"users"`
}

func (t TestStructWithCollectionName) CollectionName() string {
	return "users"
}

func TestScanner_Scan(t *testing.T) {
	integer := 73
	integerPointer := &integer
	float := 73.73
	floatPointer := &float
	boolean := true
	booleanPointer := &boolean
	stringValue := "test"
	stringPointer := &stringValue
	slice := []string{"test"}
	slicePointer := &slice
	integer16 := int16(73)
	integer16Pointer := &integer16
	integer32 := int32(73)
	integer32Pointer := &integer32
	integer64 := int64(73)
	integer64Pointer := &integer64
	uinteger := uint(73)
	uintegerPointer := &uinteger
	uinteger16 := uint16(73)
	uinteger16Pointer := &uinteger16
	uinteger32 := uint32(73)
	uinteger32Pointer := &uinteger32
	uinteger64 := uint64(73)
	uinteger64Pointer := &uinteger64
	double32 := float32(73.73)
	double32Pointer := &double32
	double64 := float64(73.73)
	double64Pointer := &double64

	tests := []struct {
		name    string
		strct   interface{}
		wantErr bool
		want    []field.IFilterField
	}{
		{
			name: "Scan struct with pointer to int",
			strct: TestStructWithIntPointer{
				Age: integerPointer,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Int.String(),
					"age", *integerPointer,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with int",
			strct: TestStructWithInt{
				Age: integer,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Int.String(),
					"age", integer,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with pointer to float",
			strct: TestStructWithFloatPointer{
				Age: floatPointer,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Float64.String(),
					"age", *floatPointer,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with float",
			strct: TestStructWithFloat{
				Age: float,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Float64.String(),
					"age", float,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with pointer to bool",
			strct: TestStructWithBoolPointer{
				Active: booleanPointer,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Bool.String(),
					"active", *booleanPointer,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with bool",
			strct: TestStructWithBool{
				Active: boolean,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Bool.String(),
					"active", boolean,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with pointer to string",
			strct: TestStructWithStringPointer{
				Name: stringPointer,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.String.String(),
					"name", *stringPointer,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with string",
			strct: TestStructWithString{
				Name: stringValue,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.String.String(),
					"name", stringValue,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with pointer to slice",
			strct: TestStructWithSlicePointer{
				Names: slicePointer,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Slice.String(),
					"names", *slicePointer,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with slice",
			strct: TestStructWithSlice{
				Names: slice,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Slice.String(),
					"names", slice,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with int32 pointer",
			strct: TestStructWithInt32Pointer{
				Age: integer32Pointer,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Int32.String(),
					"age", *integer32Pointer,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with int32",
			strct: TestStructWithInt32{
				Age: integer32,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Int32.String(),
					"age", integer32,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with int64 pointer",
			strct: TestStructWithInt64Pointer{
				Age: integer64Pointer,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Int64.String(),
					"age", *integer64Pointer,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with int64",
			strct: TestStructWithInt64{
				Age: integer64,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Int64.String(),
					"age", integer64,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with uint32 pointer",
			strct: TestStructWithUint32Pointer{
				Age: uinteger32Pointer,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Uint32.String(),
					"age", *uinteger32Pointer,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with uint32",
			strct: TestStructWithUint32{
				Age: uinteger32,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Uint32.String(),
					"age", uinteger32,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with uint64 pointer",
			strct: TestStructWithUint64Pointer{
				Age: uinteger64Pointer,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Uint64.String(),
					"age", *uinteger64Pointer,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with uint64",
			strct: TestStructWithUint64{
				Age: uinteger64,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Uint64.String(),
					"age", uinteger64,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with float32 pointer",
			strct: TestStructWithFloat32Pointer{
				Age: double32Pointer,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Float32.String(),
					"age", *double32Pointer,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with float32",
			strct: TestStructWithFloat32{
				Age: double32,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Float32.String(),
					"age", double32,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with float64 pointer",
			strct: TestStructWithFloat64Pointer{
				Age: double64Pointer,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Float64.String(),
					"age", *double64Pointer,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with float64",
			strct: TestStructWithFloat64{
				Age: double64,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Float64.String(),
					"age", double64,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with int16 pointer",
			strct: TestStructWithInt16Pointer{
				Age: integer16Pointer,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Int16.String(),
					"age", *integer16Pointer,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with int16",
			strct: TestStructWithInt16{
				Age: integer16,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Int16.String(),
					"age", integer16,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with uint16 pointer",
			strct: TestStructWithUint16Pointer{
				Age: uinteger16Pointer,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Uint16.String(),
					"age", *uinteger16Pointer,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with uint16",
			strct: TestStructWithUint16{
				Age: uinteger16,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Uint16.String(),
					"age", uinteger16,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with uint pointer",
			strct: TestStructWithUintPointer{
				Age: uintegerPointer,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Uint.String(),
					"age", *uintegerPointer,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with uint",
			strct: TestStructWithUint{
				Age: uinteger,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Uint.String(),
					"age", uinteger,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with uint",
			strct: TestStructWithUint{
				Age: uinteger,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Uint.String(),
					"age", uinteger,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with nested struct pointer",
			strct: TestStructWithNestedStructPtr{
				User: &TestStructWithInt{
					Age: integer,
				},
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Int.String(),
					"user.age", integer,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with nested struct",
			strct: TestStructWithNestedStruct{
				User: TestStructWithInt{
					Age: integer,
				},
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("",
					reflect.Int.String(),
					"user.age", integer,
					operator.EQOperator{}, 0),
			},
		},
		{
			name: "Scan struct with collection name",
			strct: TestStructWithCollectionName{
				Age: integer,
			},
			wantErr: false,
			want: []field.IFilterField{
				field.NewFilterField("users",
					reflect.Int.String(),
					"age", integer,
					operator.EQOperator{}, 0),
			},
		},
	}
	// create validators
	opValidator := validator.NewOperatorValidator(operator.NewOperatorMap(), "operator")

	// gather all validators in one slice
	validators := []validator.IValidator{opValidator}

	// create a new scanner
	scan := NewScanner(operator.NewOperatorMap(), validators,
		"filter", "operator", "join")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := scan.Scan(tt.strct, nil, 0)
			if (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i, v := range got {
				if !reflect.DeepEqual(v, tt.want[i]) {
					t.Errorf("Scan() got = %v, want %v", v, tt.want[i])
				}
			}
		})
	}
}
