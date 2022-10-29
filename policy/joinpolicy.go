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

// IJoinPolicy is used to build bson filter which
// joins two fields (IFilterField) from different collections.
// It is moved to a separate interface to ensure that the
// IFilterField interface is not polluted with the logic
// related to joining fields, since there are different
// types of joins (left, right, inner, etc.) and each of them
// has its own logic.
// It is injected into the IFilterField as a dependency
// to perform the join.
type IJoinPolicy interface {
	Join(right, left field.IFilterField) bson.D
}

// leftOuterJoinPolicy joins two fields (IFilterField) from different collections
// using the left outer join method. (i.e. all records from the left table are returned,
// and the matching records from the right table are returned if any; if there is no match,
// null is returned in the result set.)
type leftOuterJoinPolicy struct {
	method string
	fields []field.IFilterField
}

func (j *leftOuterJoinPolicy) Join(right, left field.IFilterField) bson.D {
	panic("implement me")
}

func NewLeftOuterJoinPolicy() IJoinPolicy {
	return &leftOuterJoinPolicy{
		method: "leftOuter",
	}
}
