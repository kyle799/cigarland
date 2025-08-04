package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

func QueryDB(db *gorm.DB, table string, filters ...SelectionFilter) (cigars []*Cigar, qErr error) {
	log.Printf("Setting db table for query: %s", table)
	db = db.Table(table)
	queryString, queryArgs, err := GenerateDBQuery(filters...)
	if err != nil {
		log.Printf("Error generating query string: %s", err)
		qErr = fmt.Errorf("Error generating query string: %w", err)
		return nil, qErr
	}
	log.Printf("Querying with query string: \"%s\"", queryString)
	log.Printf("Query args:")
	for _, arg := range queryArgs {
		log.Printf("\t %s", arg)
	}
	queryDB := db.Where(queryString, queryArgs...).Find(&cigars)
	if queryDB.Error != nil {
		qErr = fmt.Errorf("Error querying DB: %w", queryDB.Error)
		return nil, qErr
	}
	return cigars, nil
}

func GenerateDBQuery(filters ...SelectionFilter) (filterString string, filterArgs []any, err error) {
	filterFormat := "%s %s ? %s"
	filterBuilder := strings.Builder{}
	filterArgs = make([]any, len(filters))
	for idx, filter := range filters {
		valid, vErr := ValidateFilterOperator(filter)
		if !valid {
			err = fmt.Errorf("Error validating filter: %w", vErr)
			return "", make([]any, 0), err
		}
		if ok := ValidateLogicalOperator(filter); !ok {
			err = fmt.Errorf("Error, invalid logical operator provided: %s", filter.Logical)
			return "", make([]any, 0), err
		}
		fmt.Fprintf(&filterBuilder, filterFormat, filter.Column, filter.Operator, filter.Logical)
		filterArgs[idx] = filter.Value
	}
	filterString = filterBuilder.String()
	return filterString, filterArgs, nil
}

/*
func (s SelectionFilter) OperatorString() string {
	switch s.Operator {
	case "gt":
		return ">="
	case "
		outputWriter := ParseOutputValue()
	}
}
*/

func ValidateFilter(filter SelectionFilter) (isValid bool, err error) {
	switch filter.Value.(type) {
	case float64:
		_, isValid = ValueOperatorMap[TypeFloat][filter.Operator]
		if !isValid {
			err = fmt.Errorf("Error, invalid float operator: \"%s\" provided", filter.Operator)
		}
	case int:
		_, isValid = ValueOperatorMap[TypeInt][filter.Operator]
		if !isValid {
			err = fmt.Errorf("Error, invalid int operator: \"%s\" provided", filter.Operator)
		}
	case string:
		_, isValid = ValueOperatorMap[TypeString][filter.Operator]
		if !isValid {
			err = fmt.Errorf("Error, invalid string operator: \"%s\" provided", filter.Operator)
		}
	case bool:
		_, isValid = ValueOperatorMap[TypeBool][filter.Operator]
		if !isValid {
			err = fmt.Errorf("Error, invalid boolean operator: \"%s\" provided", filter.Operator)
		}
	default:
		err = fmt.Errorf("Error unrecognized type for filter value: \"%s\"", reflect.TypeOf(filter.Value))

	}
	return isValid, err
}

func ValidateFilterOperator(filter SelectionFilter) (isValid bool, err error) {
	switch filter.Value.(type) {
	case float64:
		_, isValid = ValueOperatorMap[TypeFloat][filter.Operator]
		if !isValid {
			err = fmt.Errorf("Error, invalid float operator: %s provided", filter.Operator)
		}
	case int:
		_, isValid = ValueOperatorMap[TypeInt][filter.Operator]
		if !isValid {
			err = fmt.Errorf("Error, invalid int operator: %s provided", filter.Operator)
		}
	case string:
		_, isValid = ValueOperatorMap[TypeString][filter.Operator]
		if !isValid {
			err = fmt.Errorf("Error, invalid string operator: %s provided", filter.Operator)
		}
	case bool:
		_, isValid = ValueOperatorMap[TypeBool][filter.Operator]
		if !isValid {
			err = fmt.Errorf("Error, invalid boolean operator: %s provided", filter.Operator)
		}
	default:
		err = fmt.Errorf("Error unrecognized type for filter value: %s", reflect.TypeOf(filter.Value))

	}
	return isValid, err
}

func ValidateLogicalOperator(filter SelectionFilter) (isValid bool) {
	if _, ok := LogicalOperatorMap[filter.Logical]; !ok {
		return false
	}
	return true
}
