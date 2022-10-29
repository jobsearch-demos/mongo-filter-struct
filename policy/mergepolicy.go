// License: GNU General Public License v3.0
// Author: Kamran Valijonov
// Version: 1.0.0
// Date: 2022-10-29
// Description: Mongo Filter Builder
// This tool is used to build bson filter for mongodb based on provided struct.
// Motivation: I was tired of writing bson.M{} for every query and wanted
// something more elegant and easy to use like django-filter.

package policy

import (
	"github.com/jobsearch-demos/mongo-filter-struct/field"
	"go.mongodb.org/mongo-driver/bson"
)

// IMergePolicy is used to define the way of merging two fields (IFilterField).
// It is moved to a separate interface to ensure that the
// IFilterField interface is not polluted with the logic
// related to merging fields, since there are different
// types of merging (override, and, or, etc.) and each of them
// has its own logic. It is injected into the IFilterField
// as a dependency to perform different merge operations.
type IMergePolicy interface {
	Merge(left, right field.IFilterField) bson.D
}

// overrideMergePolicy merges two fields (IFilterField) from different collections
// using the `override` method. (i.e. the right field overrides the left field)
type overrideMergePolicy struct {
	method string
}

func (m *overrideMergePolicy) Merge(left, right field.IFilterField) bson.D {
	panic("implement me")
}

func NewOverrideMergePolicy() IMergePolicy {
	return &overrideMergePolicy{
		method: "override",
	}
}

// andMergePolicy merges two fields (IFilterField) from different collections
// using the `and` method. (i.e. the left and right fields are merged using the and operator)
type andMergePolicy struct {
	method string
}

func (m *andMergePolicy) Merge(left, right field.IFilterField) bson.D {
	panic("implement me")
}

func NewAndMergePolicy() IMergePolicy {
	return &andMergePolicy{
		method: "and",
	}
}

// orMergePolicy merges two fields (IFilterField) from different collections
// using the `or` method. (i.e. the left and right fields are merged using the or operator)
type orMergePolicy struct {
	method string
}

func (m *orMergePolicy) Merge(left, right field.IFilterField) bson.D {
	panic("implement me")
}

func NewOrMergePolicy() IMergePolicy {
	return &orMergePolicy{
		method: "or",
	}
}
