package renderers

type Table struct {
	*BaseRenderer
}

func NewTable() *Table {
	a := &Table{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "table")
	return a
}

func (a *Table) Set(name string, value interface{}) *Table {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Table) AffixFooter(value interface{}) *Table {
	a.Set("affixFooter", value)
	return a
}

func (a *Table) AffixHeader(value interface{}) *Table {
	a.Set("affixHeader", value)
	return a
}

func (a *Table) AffixRow(value interface{}) *Table {
	a.Set("affixRow", value)
	return a
}

func (a *Table) AutoFillHeight(value interface{}) *Table {
	a.Set("autoFillHeight", value)
	return a
}

func (a *Table) AutoGenerateFilter(value interface{}) *Table {
	a.Set("autoGenerateFilter", value)
	return a
}

func (a *Table) CanAccessSuperData(value interface{}) *Table {
	a.Set("canAccessSuperData", value)
	return a
}

func (a *Table) ClassName(value interface{}) *Table {
	a.Set("className", value)
	return a
}

func (a *Table) Columns(value interface{}) *Table {
	a.Set("columns", value)
	return a
}

func (a *Table) ColumnsTogglable(value interface{}) *Table {
	a.Set("columnsTogglable", value)
	return a
}

func (a *Table) CombineFromIndex(value interface{}) *Table {
	a.Set("combineFromIndex", value)
	return a
}

func (a *Table) CombineNum(value interface{}) *Table {
	a.Set("combineNum", value)
	return a
}

func (a *Table) Data(value interface{}) *Table {
	a.Set("data", value)
	return a
}

func (a *Table) DeferApi(value interface{}) *Table {
	a.Set("deferApi", value)
	return a
}

func (a *Table) Disabled(value interface{}) *Table {
	a.Set("disabled", value)
	return a
}

func (a *Table) DisabledOn(value interface{}) *Table {
	a.Set("disabledOn", value)
	return a
}

func (a *Table) EditorSetting(value interface{}) *Table {
	a.Set("editorSetting", value)
	return a
}

func (a *Table) Footable(value interface{}) *Table {
	a.Set("footable", value)
	return a
}

func (a *Table) FooterClassName(value interface{}) *Table {
	a.Set("footerClassName", value)
	return a
}

func (a *Table) HeaderClassName(value interface{}) *Table {
	a.Set("headerClassName", value)
	return a
}

func (a *Table) Hidden(value interface{}) *Table {
	a.Set("hidden", value)
	return a
}

func (a *Table) HiddenOn(value interface{}) *Table {
	a.Set("hiddenOn", value)
	return a
}

func (a *Table) Id(value interface{}) *Table {
	a.Set("id", value)
	return a
}

func (a *Table) ItemBadge(value interface{}) *Table {
	a.Set("itemBadge", value)
	return a
}

func (a *Table) OnEvent(value interface{}) *Table {
	a.Set("onEvent", value)
	return a
}

func (a *Table) Placeholder(value interface{}) *Table {
	a.Set("placeholder", value)
	return a
}

func (a *Table) PrefixRow(value interface{}) *Table {
	a.Set("prefixRow", value)
	return a
}

func (a *Table) Resizable(value interface{}) *Table {
	a.Set("resizable", value)
	return a
}

func (a *Table) RowClassNameExpr(value interface{}) *Table {
	a.Set("rowClassNameExpr", value)
	return a
}

func (a *Table) ShowFooter(value interface{}) *Table {
	a.Set("showFooter", value)
	return a
}

func (a *Table) ShowHeader(value interface{}) *Table {
	a.Set("showHeader", value)
	return a
}

func (a *Table) ShowIndex(value interface{}) *Table {
	a.Set("showIndex", value)
	return a
}

func (a *Table) Source(value interface{}) *Table {
	a.Set("source", value)
	return a
}

func (a *Table) Static(value interface{}) *Table {
	a.Set("static", value)
	return a
}

func (a *Table) StaticClassName(value interface{}) *Table {
	a.Set("staticClassName", value)
	return a
}

func (a *Table) StaticInputClassName(value interface{}) *Table {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Table) StaticLabelClassName(value interface{}) *Table {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Table) StaticOn(value interface{}) *Table {
	a.Set("staticOn", value)
	return a
}

func (a *Table) StaticPlaceholder(value interface{}) *Table {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Table) StaticSchema(value interface{}) *Table {
	a.Set("staticSchema", value)
	return a
}

func (a *Table) Style(value interface{}) *Table {
	a.Set("style", value)
	return a
}

func (a *Table) TableClassName(value interface{}) *Table {
	a.Set("tableClassName", value)
	return a
}

func (a *Table) TableLayout(value interface{}) *Table {
	a.Set("tableLayout", value)
	return a
}

func (a *Table) TestIdBuilder(value interface{}) *Table {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Table) Testid(value interface{}) *Table {
	a.Set("testid", value)
	return a
}

func (a *Table) Title(value interface{}) *Table {
	a.Set("title", value)
	return a
}

func (a *Table) ToolbarClassName(value interface{}) *Table {
	a.Set("toolbarClassName", value)
	return a
}

func (a *Table) Type(value interface{}) *Table {
	a.Set("type", value)
	return a
}

func (a *Table) UseMobileUI(value interface{}) *Table {
	a.Set("useMobileUI", value)
	return a
}

func (a *Table) Visible(value interface{}) *Table {
	a.Set("visible", value)
	return a
}

func (a *Table) VisibleOn(value interface{}) *Table {
	a.Set("visibleOn", value)
	return a
}
