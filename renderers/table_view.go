package renderers

type TableView struct {
	*BaseRenderer
}

func NewTableView() *TableView {
	t := &TableView{
		BaseRenderer: NewBaseRenderer(),
	}
	t.Set("type", "table_view")
	return t
}

func (t *TableView) Set(name string, value interface{}) *TableView {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	t.AmisSchema[name] = value
	return t
}

func (t *TableView) Type(value interface{}) *TableView {
	t.Set("type", value)
	return t
}
