package renderers

type RowSelection struct {
	*BaseRenderer
}

func NewRowSelection() *RowSelection {
	a := &RowSelection{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *RowSelection) Set(name string, value interface{}) *RowSelection {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *RowSelection) ColumnWidth(value interface{}) *RowSelection {
	a.Set("columnWidth", value)
	return a
}

func (a *RowSelection) DisableOn(value interface{}) *RowSelection {
	a.Set("disableOn", value)
	return a
}

func (a *RowSelection) KeyField(value interface{}) *RowSelection {
	a.Set("keyField", value)
	return a
}

func (a *RowSelection) RowClick(value interface{}) *RowSelection {
	a.Set("rowClick", value)
	return a
}

func (a *RowSelection) SelectedRowKeys(value interface{}) *RowSelection {
	a.Set("selectedRowKeys", value)
	return a
}

func (a *RowSelection) SelectedRowKeysExpr(value interface{}) *RowSelection {
	a.Set("selectedRowKeysExpr", value)
	return a
}

func (a *RowSelection) Selections(value interface{}) *RowSelection {
	a.Set("selections", value)
	return a
}

func (a *RowSelection) Type(value interface{}) *RowSelection {
	a.Set("type", value)
	return a
}
