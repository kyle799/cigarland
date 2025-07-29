package main

import (
	"fmt"
	"strings"
)

func QueryDB(filters ...SelectionFilter) {
}

func GenerateDBQuery(filters ...SelectionFilter) (filterString string, filterArgs []any) {
	filterFormat := "%s %s ? %s"
	filterBuilder := strings.Builder{}
	filterArgs = make([]any, len(filters))
	for idx, filter := range filters {
		fmt.Fprintf(&filterBuilder, filterFormat, filter.Column, filter.Operator, filter.Logical)
		filterArgs[idx] = filter.Value
	}
	filterString = filterBuilder.String()
	return filterString, filterArgs
}

/*
func (s SelectionFilter) OperatorString() string {
	switch s.Operator {
	case "gt":
		return ">="
	case "
	}
}
*/

func ValidateOperator() {}
