package renderers

type CRUDList struct {
	*BaseRenderer
}

func NewCRUDList() *CRUDList {
	a := &CRUDList{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "crud")
	return a
}

func (a *CRUDList) Set(name string, value interface{}) *CRUDList {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *CRUDList) AffixFooter(value interface{}) *CRUDList {
	a.Set("affixFooter", value)
	return a
}

func (a *CRUDList) AffixHeader(value interface{}) *CRUDList {
	a.Set("affixHeader", value)
	return a
}

func (a *CRUDList) AlwaysShowPagination(value interface{}) *CRUDList {
	a.Set("alwaysShowPagination", value)
	return a
}

func (a *CRUDList) Api(value interface{}) *CRUDList {
	a.Set("api", value)
	return a
}

func (a *CRUDList) AutoFillHeight(value interface{}) *CRUDList {
	a.Set("autoFillHeight", value)
	return a
}

func (a *CRUDList) AutoGenerateFilter(value interface{}) *CRUDList {
	a.Set("autoGenerateFilter", value)
	return a
}

func (a *CRUDList) AutoJumpToTopOnPagerChange(value interface{}) *CRUDList {
	a.Set("autoJumpToTopOnPagerChange", value)
	return a
}

func (a *CRUDList) BulkActions(value interface{}) *CRUDList {
	a.Set("bulkActions", value)
	return a
}

func (a *CRUDList) CheckOnItemClick(value interface{}) *CRUDList {
	a.Set("checkOnItemClick", value)
	return a
}

func (a *CRUDList) ClassName(value interface{}) *CRUDList {
	a.Set("className", value)
	return a
}

func (a *CRUDList) DefaultParams(value interface{}) *CRUDList {
	a.Set("defaultParams", value)
	return a
}

func (a *CRUDList) DeferApi(value interface{}) *CRUDList {
	a.Set("deferApi", value)
	return a
}

func (a *CRUDList) Disabled(value interface{}) *CRUDList {
	a.Set("disabled", value)
	return a
}

func (a *CRUDList) DisabledOn(value interface{}) *CRUDList {
	a.Set("disabledOn", value)
	return a
}

func (a *CRUDList) Draggable(value interface{}) *CRUDList {
	a.Set("draggable", value)
	return a
}

func (a *CRUDList) DraggableOn(value interface{}) *CRUDList {
	a.Set("draggableOn", value)
	return a
}

func (a *CRUDList) EditorSetting(value interface{}) *CRUDList {
	a.Set("editorSetting", value)
	return a
}

func (a *CRUDList) ExpandConfig(value interface{}) *CRUDList {
	a.Set("expandConfig", value)
	return a
}

func (a *CRUDList) Filter(value interface{}) *CRUDList {
	a.Set("filter", value)
	return a
}

func (a *CRUDList) FilterDefaultVisible(value interface{}) *CRUDList {
	a.Set("filterDefaultVisible", value)
	return a
}

func (a *CRUDList) FilterTogglable(value interface{}) *CRUDList {
	a.Set("filterTogglable", value)
	return a
}

func (a *CRUDList) Footer(value interface{}) *CRUDList {
	a.Set("footer", value)
	return a
}

func (a *CRUDList) FooterClassName(value interface{}) *CRUDList {
	a.Set("footerClassName", value)
	return a
}

func (a *CRUDList) FooterToolbar(value interface{}) *CRUDList {
	a.Set("footerToolbar", value)
	return a
}

func (a *CRUDList) Header(value interface{}) *CRUDList {
	a.Set("header", value)
	return a
}

func (a *CRUDList) HeaderClassName(value interface{}) *CRUDList {
	a.Set("headerClassName", value)
	return a
}

func (a *CRUDList) HeaderToolbar(value interface{}) *CRUDList {
	a.Set("headerToolbar", value)
	return a
}

func (a *CRUDList) Hidden(value interface{}) *CRUDList {
	a.Set("hidden", value)
	return a
}

func (a *CRUDList) HiddenOn(value interface{}) *CRUDList {
	a.Set("hiddenOn", value)
	return a
}

func (a *CRUDList) HideCheckToggler(value interface{}) *CRUDList {
	a.Set("hideCheckToggler", value)
	return a
}

func (a *CRUDList) HideQuickSaveBtn(value interface{}) *CRUDList {
	a.Set("hideQuickSaveBtn", value)
	return a
}

func (a *CRUDList) Id(value interface{}) *CRUDList {
	a.Set("id", value)
	return a
}

func (a *CRUDList) IndexBarOffset(value interface{}) *CRUDList {
	a.Set("indexBarOffset", value)
	return a
}

func (a *CRUDList) IndexField(value interface{}) *CRUDList {
	a.Set("indexField", value)
	return a
}

func (a *CRUDList) InitFetch(value interface{}) *CRUDList {
	a.Set("initFetch", value)
	return a
}

func (a *CRUDList) InitFetchOn(value interface{}) *CRUDList {
	a.Set("initFetchOn", value)
	return a
}

func (a *CRUDList) InnerClassName(value interface{}) *CRUDList {
	a.Set("innerClassName", value)
	return a
}

func (a *CRUDList) Interval(value interface{}) *CRUDList {
	a.Set("interval", value)
	return a
}

func (a *CRUDList) ItemAction(value interface{}) *CRUDList {
	a.Set("itemAction", value)
	return a
}

func (a *CRUDList) ItemActions(value interface{}) *CRUDList {
	a.Set("itemActions", value)
	return a
}

func (a *CRUDList) ItemCheckableOn(value interface{}) *CRUDList {
	a.Set("itemCheckableOn", value)
	return a
}

func (a *CRUDList) ItemDraggableOn(value interface{}) *CRUDList {
	a.Set("itemDraggableOn", value)
	return a
}

func (a *CRUDList) KeepItemSelectionOnPageChange(value interface{}) *CRUDList {
	a.Set("keepItemSelectionOnPageChange", value)
	return a
}

func (a *CRUDList) LabelTpl(value interface{}) *CRUDList {
	a.Set("labelTpl", value)
	return a
}

func (a *CRUDList) ListItem(value interface{}) *CRUDList {
	a.Set("listItem", value)
	return a
}

func (a *CRUDList) LoadDataOnce(value interface{}) *CRUDList {
	a.Set("loadDataOnce", value)
	return a
}

func (a *CRUDList) LoadDataOnceFetchOnFilter(value interface{}) *CRUDList {
	a.Set("loadDataOnceFetchOnFilter", value)
	return a
}

func (a *CRUDList) LoadMoreProps(value interface{}) *CRUDList {
	a.Set("loadMoreProps", value)
	return a
}

func (a *CRUDList) LoadingConfig(value interface{}) *CRUDList {
	a.Set("loadingConfig", value)
	return a
}

func (a *CRUDList) MatchFunc(value interface{}) *CRUDList {
	a.Set("matchFunc", value)
	return a
}

func (a *CRUDList) Messages(value interface{}) *CRUDList {
	a.Set("messages", value)
	return a
}

func (a *CRUDList) Mode(value interface{}) *CRUDList {
	a.Set("mode", value)
	return a
}

func (a *CRUDList) Multiple(value interface{}) *CRUDList {
	a.Set("multiple", value)
	return a
}

func (a *CRUDList) Name(value interface{}) *CRUDList {
	a.Set("name", value)
	return a
}

func (a *CRUDList) OnEvent(value interface{}) *CRUDList {
	a.Set("onEvent", value)
	return a
}

func (a *CRUDList) OrderBy(value interface{}) *CRUDList {
	a.Set("orderBy", value)
	return a
}

func (a *CRUDList) OrderDir(value interface{}) *CRUDList {
	a.Set("orderDir", value)
	return a
}

func (a *CRUDList) OrderField(value interface{}) *CRUDList {
	a.Set("orderField", value)
	return a
}

func (a *CRUDList) PageDirectionField(value interface{}) *CRUDList {
	a.Set("pageDirectionField", value)
	return a
}

func (a *CRUDList) PageField(value interface{}) *CRUDList {
	a.Set("pageField", value)
	return a
}

func (a *CRUDList) ParsePrimitiveQuery(value interface{}) *CRUDList {
	a.Set("parsePrimitiveQuery", value)
	return a
}

func (a *CRUDList) PerPage(value interface{}) *CRUDList {
	a.Set("perPage", value)
	return a
}

func (a *CRUDList) PerPageAvailable(value interface{}) *CRUDList {
	a.Set("perPageAvailable", value)
	return a
}

func (a *CRUDList) PerPageField(value interface{}) *CRUDList {
	a.Set("perPageField", value)
	return a
}

func (a *CRUDList) Placeholder(value interface{}) *CRUDList {
	a.Set("placeholder", value)
	return a
}

func (a *CRUDList) QuickSaveApi(value interface{}) *CRUDList {
	a.Set("quickSaveApi", value)
	return a
}

func (a *CRUDList) QuickSaveItemApi(value interface{}) *CRUDList {
	a.Set("quickSaveItemApi", value)
	return a
}

func (a *CRUDList) SaveOrderApi(value interface{}) *CRUDList {
	a.Set("saveOrderApi", value)
	return a
}

func (a *CRUDList) Selectable(value interface{}) *CRUDList {
	a.Set("selectable", value)
	return a
}

func (a *CRUDList) ShowFooter(value interface{}) *CRUDList {
	a.Set("showFooter", value)
	return a
}

func (a *CRUDList) ShowHeader(value interface{}) *CRUDList {
	a.Set("showHeader", value)
	return a
}

func (a *CRUDList) ShowIndexBar(value interface{}) *CRUDList {
	a.Set("showIndexBar", value)
	return a
}

func (a *CRUDList) SilentPolling(value interface{}) *CRUDList {
	a.Set("silentPolling", value)
	return a
}

func (a *CRUDList) Size(value interface{}) *CRUDList {
	a.Set("size", value)
	return a
}

func (a *CRUDList) Source(value interface{}) *CRUDList {
	a.Set("source", value)
	return a
}

func (a *CRUDList) Static(value interface{}) *CRUDList {
	a.Set("static", value)
	return a
}

func (a *CRUDList) StaticClassName(value interface{}) *CRUDList {
	a.Set("staticClassName", value)
	return a
}

func (a *CRUDList) StaticInputClassName(value interface{}) *CRUDList {
	a.Set("staticInputClassName", value)
	return a
}

func (a *CRUDList) StaticLabelClassName(value interface{}) *CRUDList {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *CRUDList) StaticOn(value interface{}) *CRUDList {
	a.Set("staticOn", value)
	return a
}

func (a *CRUDList) StaticPlaceholder(value interface{}) *CRUDList {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *CRUDList) StaticSchema(value interface{}) *CRUDList {
	a.Set("staticSchema", value)
	return a
}

func (a *CRUDList) StopAutoRefreshWhen(value interface{}) *CRUDList {
	a.Set("stopAutoRefreshWhen", value)
	return a
}

func (a *CRUDList) StopAutoRefreshWhenModalIsOpen(value interface{}) *CRUDList {
	a.Set("stopAutoRefreshWhenModalIsOpen", value)
	return a
}

func (a *CRUDList) Style(value interface{}) *CRUDList {
	a.Set("style", value)
	return a
}

func (a *CRUDList) SyncLocation(value interface{}) *CRUDList {
	a.Set("syncLocation", value)
	return a
}

func (a *CRUDList) SyncResponse2Query(value interface{}) *CRUDList {
	a.Set("syncResponse2Query", value)
	return a
}

func (a *CRUDList) TestIdBuilder(value interface{}) *CRUDList {
	a.Set("testIdBuilder", value)
	return a
}

func (a *CRUDList) Testid(value interface{}) *CRUDList {
	a.Set("testid", value)
	return a
}

func (a *CRUDList) Title(value interface{}) *CRUDList {
	a.Set("title", value)
	return a
}

func (a *CRUDList) Toolbar(value interface{}) *CRUDList {
	a.Set("toolbar", value)
	return a
}

func (a *CRUDList) ToolbarInline(value interface{}) *CRUDList {
	a.Set("toolbarInline", value)
	return a
}

func (a *CRUDList) TotalField(value interface{}) *CRUDList {
	a.Set("totalField", value)
	return a
}

func (a *CRUDList) Type(value interface{}) *CRUDList {
	a.Set("type", value)
	return a
}

func (a *CRUDList) UseMobileUI(value interface{}) *CRUDList {
	a.Set("useMobileUI", value)
	return a
}

func (a *CRUDList) ValueField(value interface{}) *CRUDList {
	a.Set("valueField", value)
	return a
}

func (a *CRUDList) Visible(value interface{}) *CRUDList {
	a.Set("visible", value)
	return a
}

func (a *CRUDList) VisibleOn(value interface{}) *CRUDList {
	a.Set("visibleOn", value)
	return a
}
