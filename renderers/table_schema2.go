package renderers

type TableSchema2 struct {
	*BaseRenderer
}

func NewTableSchema2() *TableSchema2 {
	a := &TableSchema2{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "table2")
	return a
}

func (a *TableSchema2) Set(name string, value interface{}) *TableSchema2 {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *TableSchema2) Actions(value interface{}) *TableSchema2 {
	a.Set("actions", value)
	return a
}

func (a *TableSchema2) AutoFillHeight(value interface{}) *TableSchema2 {
	a.Set("autoFillHeight", value)
	return a
}

func (a *TableSchema2) Bordered(value interface{}) *TableSchema2 {
	a.Set("bordered", value)
	return a
}

func (a *TableSchema2) CanAccessSuperData(value interface{}) *TableSchema2 {
	a.Set("canAccessSuperData", value)
	return a
}

func (a *TableSchema2) ChildrenColumnName(value interface{}) *TableSchema2 {
	a.Set("childrenColumnName", value)
	return a
}

func (a *TableSchema2) ClassName(value interface{}) *TableSchema2 {
	a.Set("className", value)
	return a
}

func (a *TableSchema2) Columns(value interface{}) *TableSchema2 {
	a.Set("columns", value)
	return a
}

func (a *TableSchema2) ColumnsTogglable(value interface{}) *TableSchema2 {
	a.Set("columnsTogglable", value)
	return a
}

func (a *TableSchema2) Disabled(value interface{}) *TableSchema2 {
	a.Set("disabled", value)
	return a
}

func (a *TableSchema2) DisabledOn(value interface{}) *TableSchema2 {
	a.Set("disabledOn", value)
	return a
}

func (a *TableSchema2) EditorSetting(value interface{}) *TableSchema2 {
	a.Set("editorSetting", value)
	return a
}

func (a *TableSchema2) Expandable(value interface{}) *TableSchema2 {
	a.Set("expandable", value)
	return a
}

func (a *TableSchema2) Footer(value interface{}) *TableSchema2 {
	a.Set("footer", value)
	return a
}

func (a *TableSchema2) Hidden(value interface{}) *TableSchema2 {
	a.Set("hidden", value)
	return a
}

func (a *TableSchema2) HiddenOn(value interface{}) *TableSchema2 {
	a.Set("hiddenOn", value)
	return a
}

func (a *TableSchema2) Id(value interface{}) *TableSchema2 {
	a.Set("id", value)
	return a
}

func (a *TableSchema2) ItemBadge(value interface{}) *TableSchema2 {
	a.Set("itemBadge", value)
	return a
}

func (a *TableSchema2) KeepItemSelectionOnPageChange(value interface{}) *TableSchema2 {
	a.Set("keepItemSelectionOnPageChange", value)
	return a
}

func (a *TableSchema2) KeyField(value interface{}) *TableSchema2 {
	a.Set("keyField", value)
	return a
}

func (a *TableSchema2) LazyRenderAfter(value interface{}) *TableSchema2 {
	a.Set("lazyRenderAfter", value)
	return a
}

func (a *TableSchema2) LineHeight(value interface{}) *TableSchema2 {
	a.Set("lineHeight", value)
	return a
}

func (a *TableSchema2) Loading(value interface{}) *TableSchema2 {
	a.Set("loading", value)
	return a
}

func (a *TableSchema2) MaxKeepItemSelectionLength(value interface{}) *TableSchema2 {
	a.Set("maxKeepItemSelectionLength", value)
	return a
}

func (a *TableSchema2) Messages(value interface{}) *TableSchema2 {
	a.Set("messages", value)
	return a
}

func (a *TableSchema2) Multiple(value interface{}) *TableSchema2 {
	a.Set("multiple", value)
	return a
}

func (a *TableSchema2) OnEvent(value interface{}) *TableSchema2 {
	a.Set("onEvent", value)
	return a
}

func (a *TableSchema2) PopOverContainer(value interface{}) *TableSchema2 {
	a.Set("popOverContainer", value)
	return a
}

func (a *TableSchema2) PrimaryField(value interface{}) *TableSchema2 {
	a.Set("primaryField", value)
	return a
}

func (a *TableSchema2) QuickSaveApi(value interface{}) *TableSchema2 {
	a.Set("quickSaveApi", value)
	return a
}

func (a *TableSchema2) QuickSaveItemApi(value interface{}) *TableSchema2 {
	a.Set("quickSaveItemApi", value)
	return a
}

func (a *TableSchema2) Reload(value interface{}) *TableSchema2 {
	a.Set("reload", value)
	return a
}

func (a *TableSchema2) RowClassNameExpr(value interface{}) *TableSchema2 {
	a.Set("rowClassNameExpr", value)
	return a
}

func (a *TableSchema2) RowSelection(value interface{}) *TableSchema2 {
	a.Set("rowSelection", value)
	return a
}

func (a *TableSchema2) Selectable(value interface{}) *TableSchema2 {
	a.Set("selectable", value)
	return a
}

func (a *TableSchema2) ShowBadge(value interface{}) *TableSchema2 {
	a.Set("showBadge", value)
	return a
}

func (a *TableSchema2) ShowHeader(value interface{}) *TableSchema2 {
	a.Set("showHeader", value)
	return a
}

func (a *TableSchema2) Source(value interface{}) *TableSchema2 {
	a.Set("source", value)
	return a
}

func (a *TableSchema2) Static(value interface{}) *TableSchema2 {
	a.Set("static", value)
	return a
}

func (a *TableSchema2) StaticClassName(value interface{}) *TableSchema2 {
	a.Set("staticClassName", value)
	return a
}

func (a *TableSchema2) StaticInputClassName(value interface{}) *TableSchema2 {
	a.Set("staticInputClassName", value)
	return a
}

func (a *TableSchema2) StaticLabelClassName(value interface{}) *TableSchema2 {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *TableSchema2) StaticOn(value interface{}) *TableSchema2 {
	a.Set("staticOn", value)
	return a
}

func (a *TableSchema2) StaticPlaceholder(value interface{}) *TableSchema2 {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *TableSchema2) StaticSchema(value interface{}) *TableSchema2 {
	a.Set("staticSchema", value)
	return a
}

func (a *TableSchema2) Sticky(value interface{}) *TableSchema2 {
	a.Set("sticky", value)
	return a
}

func (a *TableSchema2) Style(value interface{}) *TableSchema2 {
	a.Set("style", value)
	return a
}

func (a *TableSchema2) TableLayout(value interface{}) *TableSchema2 {
	a.Set("tableLayout", value)
	return a
}

func (a *TableSchema2) TestIdBuilder(value interface{}) *TableSchema2 {
	a.Set("testIdBuilder", value)
	return a
}

func (a *TableSchema2) Testid(value interface{}) *TableSchema2 {
	a.Set("testid", value)
	return a
}

func (a *TableSchema2) Title(value interface{}) *TableSchema2 {
	a.Set("title", value)
	return a
}

func (a *TableSchema2) Type(value interface{}) *TableSchema2 {
	a.Set("type", value)
	return a
}

func (a *TableSchema2) UseMobileUI(value interface{}) *TableSchema2 {
	a.Set("useMobileUI", value)
	return a
}

func (a *TableSchema2) Visible(value interface{}) *TableSchema2 {
	a.Set("visible", value)
	return a
}

func (a *TableSchema2) VisibleOn(value interface{}) *TableSchema2 {
	a.Set("visibleOn", value)
	return a
}
