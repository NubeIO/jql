package jsonql

import (
	"encoding/json"
	"strconv"
)

// JQL - JSON Query Lang struct encapsulating the JSON data.
type JQL struct {
	Data interface{}
}

type Result struct {
	Response interface{} `json:"response"`
	Error    string      `json:"error"`
	Count    int         `json:"count"`
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
func (j *JQL) Query(where string) *Result {
	parser := &Parser{
		Operators: sqlOperators,
	}
	tokens := parser.Tokenize(where)
	rpn, err := parser.ParseRPN(tokens)
	if err != nil {
		return &Result{Error: err.Error()}
	}
	count := 0

	switch v := j.Data.(type) {
	case []interface{}:
		var ret []interface{}
		for _, obj := range v {
			parser.SymbolTable = obj
			r, err := j.processObj(parser, *rpn)
			if err != nil {
				return &Result{Error: err.Error()}
			}
			if r {
				ret = append(ret, obj)
				count++
			}
		}
		return &Result{Response: ret, Count: count}
	case map[string]interface{}:
		parser.SymbolTable = v
		r, err := j.processObj(parser, *rpn)
		if err != nil {
			return &Result{Error: err.Error()}
		}
		if r {
			return &Result{Response: v, Count: 1}
		}
		return nil
	default:
		return &Result{Error: "failed to parse input data"}
	}
}

func (j *JQL) processObj(parser *Parser, rpn Lifo) (bool, error) {
	result, err := parser.Evaluate(&rpn, true)
	if err != nil {
		return false, nil
	}
	return strconv.ParseBool(result)
}
