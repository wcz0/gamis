package renderers

type ConditionGroupValue struct {
	*BaseRenderer
}

func NewConditionGroupValue() *ConditionGroupValue {
	a := &ConditionGroupValue{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *ConditionGroupValue) Set(name string, value interface{}) *ConditionGroupValue {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *ConditionGroupValue) Children(value interface{}) *ConditionGroupValue {
	a.Set("children", value)
	return a
}

func (a *ConditionGroupValue) Conjunction(value interface{}) *ConditionGroupValue {
	a.Set("conjunction", value)
	return a
}

func (a *ConditionGroupValue) Id(value interface{}) *ConditionGroupValue {
	a.Set("id", value)
	return a
}

func (a *ConditionGroupValue) If(value interface{}) *ConditionGroupValue {
	a.Set("if", value)
	return a
}

func (a *ConditionGroupValue) Not(value interface{}) *ConditionGroupValue {
	a.Set("not", value)
	return a
}
