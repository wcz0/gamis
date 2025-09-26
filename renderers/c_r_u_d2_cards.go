package renderers

type CRUD2Cards struct {
	*BaseRenderer
}

func NewCRUD2Cards() *CRUD2Cards {
	a := &CRUD2Cards{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "crud2")
	return a
}

func (a *CRUD2Cards) Set(name string, value interface{}) *CRUD2Cards {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *CRUD2Cards) AffixFooter(value interface{}) *CRUD2Cards {
	a.Set("affixFooter", value)
	return a
}

func (a *CRUD2Cards) AffixHeader(value interface{}) *CRUD2Cards {
	a.Set("affixHeader", value)
	return a
}

func (a *CRUD2Cards) Api(value interface{}) *CRUD2Cards {
	a.Set("api", value)
	return a
}

func (a *CRUD2Cards) AutoFillHeight(value interface{}) *CRUD2Cards {
	a.Set("autoFillHeight", value)
	return a
}

func (a *CRUD2Cards) AutoJumpToTopOnPagerChange(value interface{}) *CRUD2Cards {
	a.Set("autoJumpToTopOnPagerChange", value)
	return a
}

func (a *CRUD2Cards) Card(value interface{}) *CRUD2Cards {
	a.Set("card", value)
	return a
}

func (a *CRUD2Cards) CheckOnItemClick(value interface{}) *CRUD2Cards {
	a.Set("checkOnItemClick", value)
	return a
}

func (a *CRUD2Cards) ClassName(value interface{}) *CRUD2Cards {
	a.Set("className", value)
	return a
}

func (a *CRUD2Cards) Disabled(value interface{}) *CRUD2Cards {
	a.Set("disabled", value)
	return a
}

func (a *CRUD2Cards) DisabledOn(value interface{}) *CRUD2Cards {
	a.Set("disabledOn", value)
	return a
}

func (a *CRUD2Cards) EditorSetting(value interface{}) *CRUD2Cards {
	a.Set("editorSetting", value)
	return a
}

func (a *CRUD2Cards) Footer(value interface{}) *CRUD2Cards {
	a.Set("footer", value)
	return a
}

func (a *CRUD2Cards) FooterClassName(value interface{}) *CRUD2Cards {
	a.Set("footerClassName", value)
	return a
}

func (a *CRUD2Cards) FooterToolbar(value interface{}) *CRUD2Cards {
	a.Set("footerToolbar", value)
	return a
}

func (a *CRUD2Cards) FooterToolbarClassName(value interface{}) *CRUD2Cards {
	a.Set("footerToolbarClassName", value)
	return a
}

func (a *CRUD2Cards) Header(value interface{}) *CRUD2Cards {
	a.Set("header", value)
	return a
}

func (a *CRUD2Cards) HeaderClassName(value interface{}) *CRUD2Cards {
	a.Set("headerClassName", value)
	return a
}

func (a *CRUD2Cards) HeaderToolbar(value interface{}) *CRUD2Cards {
	a.Set("headerToolbar", value)
	return a
}

func (a *CRUD2Cards) HeaderToolbarClassName(value interface{}) *CRUD2Cards {
	a.Set("headerToolbarClassName", value)
	return a
}

func (a *CRUD2Cards) Hidden(value interface{}) *CRUD2Cards {
	a.Set("hidden", value)
	return a
}

func (a *CRUD2Cards) HiddenOn(value interface{}) *CRUD2Cards {
	a.Set("hiddenOn", value)
	return a
}

func (a *CRUD2Cards) HideCheckToggler(value interface{}) *CRUD2Cards {
	a.Set("hideCheckToggler", value)
	return a
}

func (a *CRUD2Cards) HideQuickSaveBtn(value interface{}) *CRUD2Cards {
	a.Set("hideQuickSaveBtn", value)
	return a
}

func (a *CRUD2Cards) Id(value interface{}) *CRUD2Cards {
	a.Set("id", value)
	return a
}

func (a *CRUD2Cards) Interval(value interface{}) *CRUD2Cards {
	a.Set("interval", value)
	return a
}

func (a *CRUD2Cards) ItemCheckableOn(value interface{}) *CRUD2Cards {
	a.Set("itemCheckableOn", value)
	return a
}

func (a *CRUD2Cards) ItemClassName(value interface{}) *CRUD2Cards {
	a.Set("itemClassName", value)
	return a
}

func (a *CRUD2Cards) ItemDraggableOn(value interface{}) *CRUD2Cards {
	a.Set("itemDraggableOn", value)
	return a
}

func (a *CRUD2Cards) KeepItemSelectionOnPageChange(value interface{}) *CRUD2Cards {
	a.Set("keepItemSelectionOnPageChange", value)
	return a
}

func (a *CRUD2Cards) LoadDataOnce(value interface{}) *CRUD2Cards {
	a.Set("loadDataOnce", value)
	return a
}

func (a *CRUD2Cards) LoadType(value interface{}) *CRUD2Cards {
	a.Set("loadType", value)
	return a
}

func (a *CRUD2Cards) LoadingConfig(value interface{}) *CRUD2Cards {
	a.Set("loadingConfig", value)
	return a
}

func (a *CRUD2Cards) MasonryLayout(value interface{}) *CRUD2Cards {
	a.Set("masonryLayout", value)
	return a
}

func (a *CRUD2Cards) Mode(value interface{}) *CRUD2Cards {
	a.Set("mode", value)
	return a
}

func (a *CRUD2Cards) Multiple(value interface{}) *CRUD2Cards {
	a.Set("multiple", value)
	return a
}

func (a *CRUD2Cards) Name(value interface{}) *CRUD2Cards {
	a.Set("name", value)
	return a
}

func (a *CRUD2Cards) OnEvent(value interface{}) *CRUD2Cards {
	a.Set("onEvent", value)
	return a
}

func (a *CRUD2Cards) PageField(value interface{}) *CRUD2Cards {
	a.Set("pageField", value)
	return a
}

func (a *CRUD2Cards) ParsePrimitiveQuery(value interface{}) *CRUD2Cards {
	a.Set("parsePrimitiveQuery", value)
	return a
}

func (a *CRUD2Cards) PerPage(value interface{}) *CRUD2Cards {
	a.Set("perPage", value)
	return a
}

func (a *CRUD2Cards) PerPageField(value interface{}) *CRUD2Cards {
	a.Set("perPageField", value)
	return a
}

func (a *CRUD2Cards) Placeholder(value interface{}) *CRUD2Cards {
	a.Set("placeholder", value)
	return a
}

func (a *CRUD2Cards) PrimaryField(value interface{}) *CRUD2Cards {
	a.Set("primaryField", value)
	return a
}

func (a *CRUD2Cards) PullRefresh(value interface{}) *CRUD2Cards {
	a.Set("pullRefresh", value)
	return a
}

func (a *CRUD2Cards) QuickSaveApi(value interface{}) *CRUD2Cards {
	a.Set("quickSaveApi", value)
	return a
}

func (a *CRUD2Cards) QuickSaveItemApi(value interface{}) *CRUD2Cards {
	a.Set("quickSaveItemApi", value)
	return a
}

func (a *CRUD2Cards) SaveOrderApi(value interface{}) *CRUD2Cards {
	a.Set("saveOrderApi", value)
	return a
}

func (a *CRUD2Cards) Selectable(value interface{}) *CRUD2Cards {
	a.Set("selectable", value)
	return a
}

func (a *CRUD2Cards) ShowFooter(value interface{}) *CRUD2Cards {
	a.Set("showFooter", value)
	return a
}

func (a *CRUD2Cards) ShowHeader(value interface{}) *CRUD2Cards {
	a.Set("showHeader", value)
	return a
}

func (a *CRUD2Cards) ShowSelection(value interface{}) *CRUD2Cards {
	a.Set("showSelection", value)
	return a
}

func (a *CRUD2Cards) SilentPolling(value interface{}) *CRUD2Cards {
	a.Set("silentPolling", value)
	return a
}

func (a *CRUD2Cards) Source(value interface{}) *CRUD2Cards {
	a.Set("source", value)
	return a
}

func (a *CRUD2Cards) Static(value interface{}) *CRUD2Cards {
	a.Set("static", value)
	return a
}

func (a *CRUD2Cards) StaticClassName(value interface{}) *CRUD2Cards {
	a.Set("staticClassName", value)
	return a
}

func (a *CRUD2Cards) StaticInputClassName(value interface{}) *CRUD2Cards {
	a.Set("staticInputClassName", value)
	return a
}

func (a *CRUD2Cards) StaticLabelClassName(value interface{}) *CRUD2Cards {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *CRUD2Cards) StaticOn(value interface{}) *CRUD2Cards {
	a.Set("staticOn", value)
	return a
}

func (a *CRUD2Cards) StaticPlaceholder(value interface{}) *CRUD2Cards {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *CRUD2Cards) StaticSchema(value interface{}) *CRUD2Cards {
	a.Set("staticSchema", value)
	return a
}

func (a *CRUD2Cards) StopAutoRefreshWhen(value interface{}) *CRUD2Cards {
	a.Set("stopAutoRefreshWhen", value)
	return a
}

func (a *CRUD2Cards) Style(value interface{}) *CRUD2Cards {
	a.Set("style", value)
	return a
}

func (a *CRUD2Cards) SyncLocation(value interface{}) *CRUD2Cards {
	a.Set("syncLocation", value)
	return a
}

func (a *CRUD2Cards) SyncResponse2Query(value interface{}) *CRUD2Cards {
	a.Set("syncResponse2Query", value)
	return a
}

func (a *CRUD2Cards) TestIdBuilder(value interface{}) *CRUD2Cards {
	a.Set("testIdBuilder", value)
	return a
}

func (a *CRUD2Cards) Testid(value interface{}) *CRUD2Cards {
	a.Set("testid", value)
	return a
}

func (a *CRUD2Cards) Title(value interface{}) *CRUD2Cards {
	a.Set("title", value)
	return a
}

func (a *CRUD2Cards) Type(value interface{}) *CRUD2Cards {
	a.Set("type", value)
	return a
}

func (a *CRUD2Cards) UseMobileUI(value interface{}) *CRUD2Cards {
	a.Set("useMobileUI", value)
	return a
}

func (a *CRUD2Cards) ValueField(value interface{}) *CRUD2Cards {
	a.Set("valueField", value)
	return a
}

func (a *CRUD2Cards) Visible(value interface{}) *CRUD2Cards {
	a.Set("visible", value)
	return a
}

func (a *CRUD2Cards) VisibleOn(value interface{}) *CRUD2Cards {
	a.Set("visibleOn", value)
	return a
}
