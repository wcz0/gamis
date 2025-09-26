package renderers

type AutoGenerateFilter struct {
	*BaseRenderer
}

func NewAutoGenerateFilter() *AutoGenerateFilter {
	a := &AutoGenerateFilter{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *AutoGenerateFilter) Set(name string, value interface{}) *AutoGenerateFilter {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *AutoGenerateFilter) ColumnsNum(value interface{}) *AutoGenerateFilter {
	a.Set("columnsNum", value)
	return a
}

func (a *AutoGenerateFilter) DefaultCollapsed(value interface{}) *AutoGenerateFilter {
	a.Set("defaultCollapsed", value)
	return a
}

func (a *AutoGenerateFilter) EnableBulkActions(value interface{}) *AutoGenerateFilter {
	a.Set("enableBulkActions", value)
	return a
}

func (a *AutoGenerateFilter) EnableBulkActionsOn(value interface{}) *AutoGenerateFilter {
	a.Set("enableBulkActionsOn", value)
	return a
}

func (a *AutoGenerateFilter) ShowBtnToolbar(value interface{}) *AutoGenerateFilter {
	a.Set("showBtnToolbar", value)
	return a
}
