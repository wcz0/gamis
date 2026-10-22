package renderers

type BaseTable struct {
	*BaseRenderer
}

func NewBaseTable() *BaseTable {
	a := &BaseTable{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *BaseTable) Set(name string, value interface{}) *BaseTable {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *BaseTable) AffixFooter(value interface{}) *BaseTable {
	a.Set("affixFooter", value)
	return a
}

func (a *BaseTable) AffixHeader(value interface{}) *BaseTable {
	a.Set("affixHeader", value)
	return a
}

func (a *BaseTable) AffixRow(value interface{}) *BaseTable {
	a.Set("affixRow", value)
	return a
}

func (a *BaseTable) AutoFillHeight(value interface{}) *BaseTable {
	a.Set("autoFillHeight", value)
	return a
}

func (a *BaseTable) AutoGenerateFilter(value interface{}) *BaseTable {
	a.Set("autoGenerateFilter", value)
	return a
}

func (a *BaseTable) CanAccessSuperData(value interface{}) *BaseTable {
	a.Set("canAccessSuperData", value)
	return a
}

func (a *BaseTable) ClassName(value interface{}) *BaseTable {
	a.Set("className", value)
	return a
}

func (a *BaseTable) Columns(value interface{}) *BaseTable {
	a.Set("columns", value)
	return a
}

func (a *BaseTable) ColumnsTogglable(value interface{}) *BaseTable {
	a.Set("columnsTogglable", value)
	return a
}

func (a *BaseTable) CombineFromIndex(value interface{}) *BaseTable {
	a.Set("combineFromIndex", value)
	return a
}

func (a *BaseTable) CombineNum(value interface{}) *BaseTable {
	a.Set("combineNum", value)
	return a
}

func (a *BaseTable) DeferApi(value interface{}) *BaseTable {
	a.Set("deferApi", value)
	return a
}

func (a *BaseTable) Disabled(value interface{}) *BaseTable {
	a.Set("disabled", value)
	return a
}

func (a *BaseTable) DisabledOn(value interface{}) *BaseTable {
	a.Set("disabledOn", value)
	return a
}

func (a *BaseTable) EditorSetting(value interface{}) *BaseTable {
	a.Set("editorSetting", value)
	return a
}

func (a *BaseTable) Footable(value interface{}) *BaseTable {
	a.Set("footable", value)
	return a
}

func (a *BaseTable) FooterClassName(value interface{}) *BaseTable {
	a.Set("footerClassName", value)
	return a
}

func (a *BaseTable) HeaderClassName(value interface{}) *BaseTable {
	a.Set("headerClassName", value)
	return a
}

func (a *BaseTable) Hidden(value interface{}) *BaseTable {
	a.Set("hidden", value)
	return a
}

func (a *BaseTable) HiddenOn(value interface{}) *BaseTable {
	a.Set("hiddenOn", value)
	return a
}

func (a *BaseTable) Id(value interface{}) *BaseTable {
	a.Set("id", value)
	return a
}

func (a *BaseTable) ItemBadge(value interface{}) *BaseTable {
	a.Set("itemBadge", value)
	return a
}

func (a *BaseTable) OnEvent(value interface{}) *BaseTable {
	a.Set("onEvent", value)
	return a
}

func (a *BaseTable) PersistKey(value interface{}) *BaseTable {
	a.Set("persistKey", value)
	return a
}

func (a *BaseTable) Placeholder(value interface{}) *BaseTable {
	a.Set("placeholder", value)
	return a
}

func (a *BaseTable) PrefixRow(value interface{}) *BaseTable {
	a.Set("prefixRow", value)
	return a
}

func (a *BaseTable) Resizable(value interface{}) *BaseTable {
	a.Set("resizable", value)
	return a
}

func (a *BaseTable) RowClassNameExpr(value interface{}) *BaseTable {
	a.Set("rowClassNameExpr", value)
	return a
}

func (a *BaseTable) ShowFooter(value interface{}) *BaseTable {
	a.Set("showFooter", value)
	return a
}

func (a *BaseTable) ShowHeader(value interface{}) *BaseTable {
	a.Set("showHeader", value)
	return a
}

func (a *BaseTable) ShowIndex(value interface{}) *BaseTable {
	a.Set("showIndex", value)
	return a
}

func (a *BaseTable) Source(value interface{}) *BaseTable {
	a.Set("source", value)
	return a
}

func (a *BaseTable) Static(value interface{}) *BaseTable {
	a.Set("static", value)
	return a
}

func (a *BaseTable) StaticClassName(value interface{}) *BaseTable {
	a.Set("staticClassName", value)
	return a
}

func (a *BaseTable) StaticInputClassName(value interface{}) *BaseTable {
	a.Set("staticInputClassName", value)
	return a
}

func (a *BaseTable) StaticLabelClassName(value interface{}) *BaseTable {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *BaseTable) StaticOn(value interface{}) *BaseTable {
	a.Set("staticOn", value)
	return a
}

func (a *BaseTable) StaticPlaceholder(value interface{}) *BaseTable {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *BaseTable) StaticSchema(value interface{}) *BaseTable {
	a.Set("staticSchema", value)
	return a
}

func (a *BaseTable) Style(value interface{}) *BaseTable {
	a.Set("style", value)
	return a
}

func (a *BaseTable) TableClassName(value interface{}) *BaseTable {
	a.Set("tableClassName", value)
	return a
}

func (a *BaseTable) TableLayout(value interface{}) *BaseTable {
	a.Set("tableLayout", value)
	return a
}

func (a *BaseTable) Testid(value interface{}) *BaseTable {
	a.Set("testid", value)
	return a
}

func (a *BaseTable) Title(value interface{}) *BaseTable {
	a.Set("title", value)
	return a
}

func (a *BaseTable) ToolbarClassName(value interface{}) *BaseTable {
	a.Set("toolbarClassName", value)
	return a
}

func (a *BaseTable) Type(value interface{}) *BaseTable {
	a.Set("type", value)
	return a
}

func (a *BaseTable) UseMobileUI(value interface{}) *BaseTable {
	a.Set("useMobileUI", value)
	return a
}

func (a *BaseTable) Visible(value interface{}) *BaseTable {
	a.Set("visible", value)
	return a
}

func (a *BaseTable) VisibleOn(value interface{}) *BaseTable {
	a.Set("visibleOn", value)
	return a
}
