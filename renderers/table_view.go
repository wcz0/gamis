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

func (t *TableView) Type(value interface{}) *TableView {
	t.Set("type", value)
	return t
}
