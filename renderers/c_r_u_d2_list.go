package renderers

type CRUD2List struct {
	*BaseRenderer
}

func NewCRUD2List() *CRUD2List {
	a := &CRUD2List{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "crud2")
	return a
}

func (a *CRUD2List) Set(name string, value interface{}) *CRUD2List {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *CRUD2List) AffixFooter(value interface{}) *CRUD2List {
	a.Set("affixFooter", value)
	return a
}

func (a *CRUD2List) AffixHeader(value interface{}) *CRUD2List {
	a.Set("affixHeader", value)
	return a
}

func (a *CRUD2List) Api(value interface{}) *CRUD2List {
	a.Set("api", value)
	return a
}

func (a *CRUD2List) AutoFillHeight(value interface{}) *CRUD2List {
	a.Set("autoFillHeight", value)
	return a
}

func (a *CRUD2List) AutoJumpToTopOnPagerChange(value interface{}) *CRUD2List {
	a.Set("autoJumpToTopOnPagerChange", value)
	return a
}

func (a *CRUD2List) CheckOnItemClick(value interface{}) *CRUD2List {
	a.Set("checkOnItemClick", value)
	return a
}

func (a *CRUD2List) ClassName(value interface{}) *CRUD2List {
	a.Set("className", value)
	return a
}

func (a *CRUD2List) Disabled(value interface{}) *CRUD2List {
	a.Set("disabled", value)
	return a
}

func (a *CRUD2List) DisabledOn(value interface{}) *CRUD2List {
	a.Set("disabledOn", value)
	return a
}

func (a *CRUD2List) EditorSetting(value interface{}) *CRUD2List {
	a.Set("editorSetting", value)
	return a
}

func (a *CRUD2List) Footer(value interface{}) *CRUD2List {
	a.Set("footer", value)
	return a
}

func (a *CRUD2List) FooterClassName(value interface{}) *CRUD2List {
	a.Set("footerClassName", value)
	return a
}

func (a *CRUD2List) FooterToolbar(value interface{}) *CRUD2List {
	a.Set("footerToolbar", value)
	return a
}

func (a *CRUD2List) FooterToolbarClassName(value interface{}) *CRUD2List {
	a.Set("footerToolbarClassName", value)
	return a
}

func (a *CRUD2List) Header(value interface{}) *CRUD2List {
	a.Set("header", value)
	return a
}

func (a *CRUD2List) HeaderClassName(value interface{}) *CRUD2List {
	a.Set("headerClassName", value)
	return a
}

func (a *CRUD2List) HeaderToolbar(value interface{}) *CRUD2List {
	a.Set("headerToolbar", value)
	return a
}

func (a *CRUD2List) HeaderToolbarClassName(value interface{}) *CRUD2List {
	a.Set("headerToolbarClassName", value)
	return a
}

func (a *CRUD2List) Hidden(value interface{}) *CRUD2List {
	a.Set("hidden", value)
	return a
}

func (a *CRUD2List) HiddenOn(value interface{}) *CRUD2List {
	a.Set("hiddenOn", value)
	return a
}

func (a *CRUD2List) HideCheckToggler(value interface{}) *CRUD2List {
	a.Set("hideCheckToggler", value)
	return a
}

func (a *CRUD2List) HideQuickSaveBtn(value interface{}) *CRUD2List {
	a.Set("hideQuickSaveBtn", value)
	return a
}

func (a *CRUD2List) Id(value interface{}) *CRUD2List {
	a.Set("id", value)
	return a
}

func (a *CRUD2List) IndexBarOffset(value interface{}) *CRUD2List {
	a.Set("indexBarOffset", value)
	return a
}

func (a *CRUD2List) IndexField(value interface{}) *CRUD2List {
	a.Set("indexField", value)
	return a
}

func (a *CRUD2List) Interval(value interface{}) *CRUD2List {
	a.Set("interval", value)
	return a
}

func (a *CRUD2List) ItemAction(value interface{}) *CRUD2List {
	a.Set("itemAction", value)
	return a
}

func (a *CRUD2List) ItemCheckableOn(value interface{}) *CRUD2List {
	a.Set("itemCheckableOn", value)
	return a
}

func (a *CRUD2List) ItemDraggableOn(value interface{}) *CRUD2List {
	a.Set("itemDraggableOn", value)
	return a
}

func (a *CRUD2List) KeepItemSelectionOnPageChange(value interface{}) *CRUD2List {
	a.Set("keepItemSelectionOnPageChange", value)
	return a
}

func (a *CRUD2List) ListItem(value interface{}) *CRUD2List {
	a.Set("listItem", value)
	return a
}

func (a *CRUD2List) LoadDataOnce(value interface{}) *CRUD2List {
	a.Set("loadDataOnce", value)
	return a
}

func (a *CRUD2List) LoadType(value interface{}) *CRUD2List {
	a.Set("loadType", value)
	return a
}

func (a *CRUD2List) LoadingConfig(value interface{}) *CRUD2List {
	a.Set("loadingConfig", value)
	return a
}

func (a *CRUD2List) Mode(value interface{}) *CRUD2List {
	a.Set("mode", value)
	return a
}

func (a *CRUD2List) Multiple(value interface{}) *CRUD2List {
	a.Set("multiple", value)
	return a
}

func (a *CRUD2List) Name(value interface{}) *CRUD2List {
	a.Set("name", value)
	return a
}

func (a *CRUD2List) OnEvent(value interface{}) *CRUD2List {
	a.Set("onEvent", value)
	return a
}

func (a *CRUD2List) PageField(value interface{}) *CRUD2List {
	a.Set("pageField", value)
	return a
}

func (a *CRUD2List) ParsePrimitiveQuery(value interface{}) *CRUD2List {
	a.Set("parsePrimitiveQuery", value)
	return a
}

func (a *CRUD2List) PerPage(value interface{}) *CRUD2List {
	a.Set("perPage", value)
	return a
}

func (a *CRUD2List) PerPageField(value interface{}) *CRUD2List {
	a.Set("perPageField", value)
	return a
}

func (a *CRUD2List) Placeholder(value interface{}) *CRUD2List {
	a.Set("placeholder", value)
	return a
}

func (a *CRUD2List) PrimaryField(value interface{}) *CRUD2List {
	a.Set("primaryField", value)
	return a
}

func (a *CRUD2List) PullRefresh(value interface{}) *CRUD2List {
	a.Set("pullRefresh", value)
	return a
}

func (a *CRUD2List) QuickSaveApi(value interface{}) *CRUD2List {
	a.Set("quickSaveApi", value)
	return a
}

func (a *CRUD2List) QuickSaveItemApi(value interface{}) *CRUD2List {
	a.Set("quickSaveItemApi", value)
	return a
}

func (a *CRUD2List) SaveOrderApi(value interface{}) *CRUD2List {
	a.Set("saveOrderApi", value)
	return a
}

func (a *CRUD2List) Selectable(value interface{}) *CRUD2List {
	a.Set("selectable", value)
	return a
}

func (a *CRUD2List) ShowFooter(value interface{}) *CRUD2List {
	a.Set("showFooter", value)
	return a
}

func (a *CRUD2List) ShowHeader(value interface{}) *CRUD2List {
	a.Set("showHeader", value)
	return a
}

func (a *CRUD2List) ShowIndexBar(value interface{}) *CRUD2List {
	a.Set("showIndexBar", value)
	return a
}

func (a *CRUD2List) ShowSelection(value interface{}) *CRUD2List {
	a.Set("showSelection", value)
	return a
}

func (a *CRUD2List) SilentPolling(value interface{}) *CRUD2List {
	a.Set("silentPolling", value)
	return a
}

func (a *CRUD2List) Size(value interface{}) *CRUD2List {
	a.Set("size", value)
	return a
}

func (a *CRUD2List) Source(value interface{}) *CRUD2List {
	a.Set("source", value)
	return a
}

func (a *CRUD2List) Static(value interface{}) *CRUD2List {
	a.Set("static", value)
	return a
}

func (a *CRUD2List) StaticClassName(value interface{}) *CRUD2List {
	a.Set("staticClassName", value)
	return a
}

func (a *CRUD2List) StaticInputClassName(value interface{}) *CRUD2List {
	a.Set("staticInputClassName", value)
	return a
}

func (a *CRUD2List) StaticLabelClassName(value interface{}) *CRUD2List {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *CRUD2List) StaticOn(value interface{}) *CRUD2List {
	a.Set("staticOn", value)
	return a
}

func (a *CRUD2List) StaticPlaceholder(value interface{}) *CRUD2List {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *CRUD2List) StaticSchema(value interface{}) *CRUD2List {
	a.Set("staticSchema", value)
	return a
}

func (a *CRUD2List) StopAutoRefreshWhen(value interface{}) *CRUD2List {
	a.Set("stopAutoRefreshWhen", value)
	return a
}

func (a *CRUD2List) Style(value interface{}) *CRUD2List {
	a.Set("style", value)
	return a
}

func (a *CRUD2List) SyncLocation(value interface{}) *CRUD2List {
	a.Set("syncLocation", value)
	return a
}

func (a *CRUD2List) SyncResponse2Query(value interface{}) *CRUD2List {
	a.Set("syncResponse2Query", value)
	return a
}

func (a *CRUD2List) TestIdBuilder(value interface{}) *CRUD2List {
	a.Set("testIdBuilder", value)
	return a
}

func (a *CRUD2List) Testid(value interface{}) *CRUD2List {
	a.Set("testid", value)
	return a
}

func (a *CRUD2List) Title(value interface{}) *CRUD2List {
	a.Set("title", value)
	return a
}

func (a *CRUD2List) Type(value interface{}) *CRUD2List {
	a.Set("type", value)
	return a
}

func (a *CRUD2List) UseMobileUI(value interface{}) *CRUD2List {
	a.Set("useMobileUI", value)
	return a
}

func (a *CRUD2List) ValueField(value interface{}) *CRUD2List {
	a.Set("valueField", value)
	return a
}

func (a *CRUD2List) Visible(value interface{}) *CRUD2List {
	a.Set("visible", value)
	return a
}

func (a *CRUD2List) VisibleOn(value interface{}) *CRUD2List {
	a.Set("visibleOn", value)
	return a
}
