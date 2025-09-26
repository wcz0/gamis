package renderers

type TableColumn struct {
	*BaseRenderer
}

func NewTableColumn() *TableColumn {
	a := &TableColumn{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *TableColumn) Set(name string, value interface{}) *TableColumn {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *TableColumn) Align(value interface{}) *TableColumn {
	a.Set("align", value)
	return a
}

func (a *TableColumn) Breakpoint(value interface{}) *TableColumn {
	a.Set("breakpoint", value)
	return a
}

func (a *TableColumn) CanAccessSuperData(value interface{}) *TableColumn {
	a.Set("canAccessSuperData", value)
	return a
}

func (a *TableColumn) ClassName(value interface{}) *TableColumn {
	a.Set("className", value)
	return a
}

func (a *TableColumn) ClassNameExpr(value interface{}) *TableColumn {
	a.Set("classNameExpr", value)
	return a
}

func (a *TableColumn) Copyable(value interface{}) *TableColumn {
	a.Set("copyable", value)
	return a
}

func (a *TableColumn) Filterable(value interface{}) *TableColumn {
	a.Set("filterable", value)
	return a
}

func (a *TableColumn) Fixed(value interface{}) *TableColumn {
	a.Set("fixed", value)
	return a
}

func (a *TableColumn) HeaderAlign(value interface{}) *TableColumn {
	a.Set("headerAlign", value)
	return a
}

func (a *TableColumn) InnerStyle(value interface{}) *TableColumn {
	a.Set("innerStyle", value)
	return a
}

func (a *TableColumn) Label(value interface{}) *TableColumn {
	a.Set("label", value)
	return a
}

func (a *TableColumn) LabelClassName(value interface{}) *TableColumn {
	a.Set("labelClassName", value)
	return a
}

func (a *TableColumn) LazyRenderAfter(value interface{}) *TableColumn {
	a.Set("lazyRenderAfter", value)
	return a
}

func (a *TableColumn) Name(value interface{}) *TableColumn {
	a.Set("name", value)
	return a
}

func (a *TableColumn) PopOver(value interface{}) *TableColumn {
	a.Set("popOver", value)
	return a
}

func (a *TableColumn) QuickEdit(value interface{}) *TableColumn {
	a.Set("quickEdit", value)
	return a
}

func (a *TableColumn) QuickEditOnUpdate(value interface{}) *TableColumn {
	a.Set("quickEditOnUpdate", value)
	return a
}

func (a *TableColumn) Remark(value interface{}) *TableColumn {
	a.Set("remark", value)
	return a
}

func (a *TableColumn) Searchable(value interface{}) *TableColumn {
	a.Set("searchable", value)
	return a
}

func (a *TableColumn) Sortable(value interface{}) *TableColumn {
	a.Set("sortable", value)
	return a
}

func (a *TableColumn) Toggled(value interface{}) *TableColumn {
	a.Set("toggled", value)
	return a
}

func (a *TableColumn) Type(value interface{}) *TableColumn {
	a.Set("type", value)
	return a
}

func (a *TableColumn) Unique(value interface{}) *TableColumn {
	a.Set("unique", value)
	return a
}

func (a *TableColumn) VAlign(value interface{}) *TableColumn {
	a.Set("vAlign", value)
	return a
}

func (a *TableColumn) Value(value interface{}) *TableColumn {
	a.Set("value", value)
	return a
}

func (a *TableColumn) Width(value interface{}) *TableColumn {
	a.Set("width", value)
	return a
}
