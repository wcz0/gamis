package renderers

type TableView struct {
	*BaseRenderer
}

func NewTableView() *TableView {
	a := &TableView{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "table-view")
	return a
}

func (a *TableView) Set(name string, value interface{}) *TableView {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *TableView) Type(value interface{}) *TableView {
	a.Set("type", value)
	return a
}
