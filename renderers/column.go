package renderers

type Column struct {
	*BaseRenderer
}

func NewColumn() *Column {
	a := &Column{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *Column) Set(name string, value interface{}) *Column {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Column) Align(value interface{}) *Column {
	a.Set("align", value)
	return a
}

func (a *Column) CanAccessSuperData(value interface{}) *Column {
	a.Set("canAccessSuperData", value)
	return a
}

func (a *Column) Children(value interface{}) *Column {
	a.Set("children", value)
	return a
}

func (a *Column) ClassName(value interface{}) *Column {
	a.Set("className", value)
	return a
}

func (a *Column) ClassNameExpr(value interface{}) *Column {
	a.Set("classNameExpr", value)
	return a
}

func (a *Column) ColSpanExpr(value interface{}) *Column {
	a.Set("colSpanExpr", value)
	return a
}

func (a *Column) Copyable(value interface{}) *Column {
	a.Set("copyable", value)
	return a
}

func (a *Column) Filterable(value interface{}) *Column {
	a.Set("filterable", value)
	return a
}

func (a *Column) Fixed(value interface{}) *Column {
	a.Set("fixed", value)
	return a
}

func (a *Column) HeaderAlign(value interface{}) *Column {
	a.Set("headerAlign", value)
	return a
}

func (a *Column) Name(value interface{}) *Column {
	a.Set("name", value)
	return a
}

func (a *Column) QuickEdit(value interface{}) *Column {
	a.Set("quickEdit", value)
	return a
}

func (a *Column) Remark(value interface{}) *Column {
	a.Set("remark", value)
	return a
}

func (a *Column) RowSpanExpr(value interface{}) *Column {
	a.Set("rowSpanExpr", value)
	return a
}

func (a *Column) Searchable(value interface{}) *Column {
	a.Set("searchable", value)
	return a
}

func (a *Column) Sortable(value interface{}) *Column {
	a.Set("sortable", value)
	return a
}

func (a *Column) Sorter(value interface{}) *Column {
	a.Set("sorter", value)
	return a
}

func (a *Column) Title(value interface{}) *Column {
	a.Set("title", value)
	return a
}

func (a *Column) TitleClassName(value interface{}) *Column {
	a.Set("titleClassName", value)
	return a
}

func (a *Column) Toggled(value interface{}) *Column {
	a.Set("toggled", value)
	return a
}

func (a *Column) Type(value interface{}) *Column {
	a.Set("type", value)
	return a
}

func (a *Column) VAlign(value interface{}) *Column {
	a.Set("vAlign", value)
	return a
}

func (a *Column) Width(value interface{}) *Column {
	a.Set("width", value)
	return a
}
