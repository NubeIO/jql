package jsonql

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// JQL - JSON Query Lang struct encapsulating the JSON data.
type JQL struct {
	Data interface{}
}

type Result struct {
	Response interface{}
	Count    int
}

func New() *JQL {
	return &JQL{}
}

// NewStringData - creates a new &JQL from raw JSON string
func (j *JQL) NewStringData(jsonString string) (*JQL, error) {
	var data = new(interface{})
	err := json.Unmarshal([]byte(jsonString), data)
	if err != nil {
		return nil, err
	}
	j.Data = *data
	return j, nil
}

// NewData - creates a new &JQL from an array of interface{} or a map of [string]interface{}
func (j *JQL) NewData(data interface{}) *JQL {
	j.Data = data
	return j
}

// Query - queries against the JSON using the conditions specified in the where stirng.
func (j *JQL) Query(where string) (*Result, error) {
	parser := &Parser{
		Operators: sqlOperators,
	}
	tokens := parser.Tokenize(where)
	rpn, err := parser.ParseRPN(tokens)
	if err != nil {
		return nil, err
	}
	count := 0

	switch v := j.Data.(type) {
	case []interface{}:
		var ret []interface{}
		for _, obj := range v {
			parser.SymbolTable = obj
			r, err := j.processObj(parser, *rpn)
			if err != nil {
				return nil, err
			}
			if r {
				ret = append(ret, obj)
				count++
			}
		}
		return &Result{Response: ret, Count: count}, nil
	case map[string]interface{}:
		parser.SymbolTable = v
		r, err := j.processObj(parser, *rpn)
		if err != nil {
			return nil, err
		}
		if r {
			return &Result{Response: v, Count: 1}, nil
		}
		return nil, nil
	default:
		return nil, fmt.Errorf("failed to parse input data")
	}
}

func (j *JQL) processObj(parser *Parser, rpn Lifo) (bool, error) {
	result, err := parser.Evaluate(&rpn, true)
	if err != nil {
		return false, nil
	}
	return strconv.ParseBool(result)
}
