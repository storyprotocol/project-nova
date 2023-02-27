package entity

type KeyValue interface {
	GetKey() string
	GetValue() interface{}
}

type ListOption struct {
	Filter *Filter `json:"filter"`
}

type Filter struct {
	Operator FilterOperator `json:"operator"`
	Operands []*Operand     `json:"operands"`
}

func (f *Filter) Validate() bool {
	if !f.Operator.Validate() {
		return false
	}

	if f.Operands == nil || len(f.Operands) == 0 {
		return false
	}

	for _, operand := range f.Operands {
		if !operand.Validate() {
			return false
		}
	}
	return true
}

func (f *Filter) Eval(kvs []KeyValue) bool {
	switch f.Operator {
	case FilterOperators.And:
		return f.evalAnd(kvs)
	}
	return false
}

func (f *Filter) evalAnd(kvs []KeyValue) bool {
	matchedCount := 0
	for _, operand := range f.Operands {
		matched := false
		for _, kv := range kvs {
			if operand.Field == kv.GetKey() {
				passed := operand.Eval(kv.GetValue())
				if passed {
					matched = true
					break
				}
			}
		}
		if !matched {
			return false
		}
		matchedCount++
	}

	return matchedCount == len(f.Operands)
}

type FilterOperator string

var FilterOperators = struct {
	And FilterOperator
}{
	And: "and",
}

func (f FilterOperator) Validate() bool {
	switch f {
	case FilterOperators.And:
		return true
	}
	return false
}

type Operand struct {
	Operator OperandOperator `json:"operator"`
	Field    string          `json:"field"`
	Value    interface{}     `json:"value"`
}

func (o *Operand) Validate() bool {
	return o.Operator.Validate()
}

func (o *Operand) Eval(value interface{}) bool {
	switch o.Operator {
	case OperandOperators.Eq:
		return o.Value == value
	}
	return false
}

type OperandOperator string

var OperandOperators = struct {
	Eq OperandOperator
}{
	Eq: "eq",
}

func (f OperandOperator) Validate() bool {
	switch f {
	case OperandOperators.Eq:
		return true
	}
	return false
}
