package renderers

type Expandable struct {
	*BaseRenderer
}

func NewExpandable() *Expandable {
	a := &Expandable{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *Expandable) Set(name string, value interface{}) *Expandable {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Expandable) ExpandableOn(value interface{}) *Expandable {
	a.Set("expandableOn", value)
	return a
}

func (a *Expandable) ExpandedRowClassNameExpr(value interface{}) *Expandable {
	a.Set("expandedRowClassNameExpr", value)
	return a
}

func (a *Expandable) ExpandedRowKeys(value interface{}) *Expandable {
	a.Set("expandedRowKeys", value)
	return a
}

func (a *Expandable) ExpandedRowKeysExpr(value interface{}) *Expandable {
	a.Set("expandedRowKeysExpr", value)
	return a
}

func (a *Expandable) KeyField(value interface{}) *Expandable {
	a.Set("keyField", value)
	return a
}

func (a *Expandable) Type(value interface{}) *Expandable {
	a.Set("type", value)
	return a
}
