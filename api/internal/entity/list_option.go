package entity

import "fmt"

type ListOption struct {
	Filter Filter `json:"filter"`
}

type Filter struct {
	Operator Operator   `json:"operator"`
	Operands []*Operand `json:"operands"`
}

func (f *Filter) Validate() bool {
	if f.Operands == nil || len(f.Operands) == 0 {
		return false
	}
	switch f.Operator {
	case Operators.And:
		return true
	}

	return false
}

type Operator string

var Operators = struct {
	And Operator
	Eq  Operator
}{
	And: "and",
	Eq:  "eq",
}

type Operand struct {
	Operator Operator    `json:"operator"`
	Field    string      `json:"field"`
	Value    interface{} `json:"value"`
}

func (o *Operand) Eval(value interface{}) (bool, error) {
	switch o.Operator {
	case Operators.Eq:
		return o.Value == value, nil
	}

	return false, fmt.Errorf("unsupported operator type for the operand: %v", o)
}
