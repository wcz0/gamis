package renderers

type CRUDTable struct {
	*BaseRenderer
}

func NewCRUDTable() *CRUDTable {
	a := &CRUDTable{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "crud")
	return a
}

func (a *CRUDTable) Set(name string, value interface{}) *CRUDTable {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *CRUDTable) AffixFooter(value interface{}) *CRUDTable {
	a.Set("affixFooter", value)
	return a
}

func (a *CRUDTable) AffixHeader(value interface{}) *CRUDTable {
	a.Set("affixHeader", value)
	return a
}

func (a *CRUDTable) AffixRow(value interface{}) *CRUDTable {
	a.Set("affixRow", value)
	return a
}

func (a *CRUDTable) AlwaysShowPagination(value interface{}) *CRUDTable {
	a.Set("alwaysShowPagination", value)
	return a
}

func (a *CRUDTable) Api(value interface{}) *CRUDTable {
	a.Set("api", value)
	return a
}

func (a *CRUDTable) AutoFillHeight(value interface{}) *CRUDTable {
	a.Set("autoFillHeight", value)
	return a
}

func (a *CRUDTable) AutoGenerateFilter(value interface{}) *CRUDTable {
	a.Set("autoGenerateFilter", value)
	return a
}

func (a *CRUDTable) AutoJumpToTopOnPagerChange(value interface{}) *CRUDTable {
	a.Set("autoJumpToTopOnPagerChange", value)
	return a
}

func (a *CRUDTable) BulkActions(value interface{}) *CRUDTable {
	a.Set("bulkActions", value)
	return a
}

func (a *CRUDTable) CanAccessSuperData(value interface{}) *CRUDTable {
	a.Set("canAccessSuperData", value)
	return a
}

func (a *CRUDTable) ClassName(value interface{}) *CRUDTable {
	a.Set("className", value)
	return a
}

func (a *CRUDTable) Columns(value interface{}) *CRUDTable {
	a.Set("columns", value)
	return a
}

func (a *CRUDTable) ColumnsTogglable(value interface{}) *CRUDTable {
	a.Set("columnsTogglable", value)
	return a
}

func (a *CRUDTable) CombineFromIndex(value interface{}) *CRUDTable {
	a.Set("combineFromIndex", value)
	return a
}

func (a *CRUDTable) CombineNum(value interface{}) *CRUDTable {
	a.Set("combineNum", value)
	return a
}

func (a *CRUDTable) Data(value interface{}) *CRUDTable {
	a.Set("data", value)
	return a
}

func (a *CRUDTable) DefaultParams(value interface{}) *CRUDTable {
	a.Set("defaultParams", value)
	return a
}

func (a *CRUDTable) DeferApi(value interface{}) *CRUDTable {
	a.Set("deferApi", value)
	return a
}

func (a *CRUDTable) Disabled(value interface{}) *CRUDTable {
	a.Set("disabled", value)
	return a
}

func (a *CRUDTable) DisabledOn(value interface{}) *CRUDTable {
	a.Set("disabledOn", value)
	return a
}

func (a *CRUDTable) Draggable(value interface{}) *CRUDTable {
	a.Set("draggable", value)
	return a
}

func (a *CRUDTable) DraggableOn(value interface{}) *CRUDTable {
	a.Set("draggableOn", value)
	return a
}

func (a *CRUDTable) EditorSetting(value interface{}) *CRUDTable {
	a.Set("editorSetting", value)
	return a
}

func (a *CRUDTable) ExpandConfig(value interface{}) *CRUDTable {
	a.Set("expandConfig", value)
	return a
}

func (a *CRUDTable) Filter(value interface{}) *CRUDTable {
	a.Set("filter", value)
	return a
}

func (a *CRUDTable) FilterDefaultVisible(value interface{}) *CRUDTable {
	a.Set("filterDefaultVisible", value)
	return a
}

func (a *CRUDTable) FilterTogglable(value interface{}) *CRUDTable {
	a.Set("filterTogglable", value)
	return a
}

func (a *CRUDTable) Footable(value interface{}) *CRUDTable {
	a.Set("footable", value)
	return a
}

func (a *CRUDTable) FooterClassName(value interface{}) *CRUDTable {
	a.Set("footerClassName", value)
	return a
}

func (a *CRUDTable) FooterToolbar(value interface{}) *CRUDTable {
	a.Set("footerToolbar", value)
	return a
}

func (a *CRUDTable) HeaderClassName(value interface{}) *CRUDTable {
	a.Set("headerClassName", value)
	return a
}

func (a *CRUDTable) HeaderToolbar(value interface{}) *CRUDTable {
	a.Set("headerToolbar", value)
	return a
}

func (a *CRUDTable) Hidden(value interface{}) *CRUDTable {
	a.Set("hidden", value)
	return a
}

func (a *CRUDTable) HiddenOn(value interface{}) *CRUDTable {
	a.Set("hiddenOn", value)
	return a
}

func (a *CRUDTable) HideQuickSaveBtn(value interface{}) *CRUDTable {
	a.Set("hideQuickSaveBtn", value)
	return a
}

func (a *CRUDTable) Id(value interface{}) *CRUDTable {
	a.Set("id", value)
	return a
}

func (a *CRUDTable) InitFetch(value interface{}) *CRUDTable {
	a.Set("initFetch", value)
	return a
}

func (a *CRUDTable) InitFetchOn(value interface{}) *CRUDTable {
	a.Set("initFetchOn", value)
	return a
}

func (a *CRUDTable) InnerClassName(value interface{}) *CRUDTable {
	a.Set("innerClassName", value)
	return a
}

func (a *CRUDTable) Interval(value interface{}) *CRUDTable {
	a.Set("interval", value)
	return a
}

func (a *CRUDTable) ItemActions(value interface{}) *CRUDTable {
	a.Set("itemActions", value)
	return a
}

func (a *CRUDTable) ItemBadge(value interface{}) *CRUDTable {
	a.Set("itemBadge", value)
	return a
}

func (a *CRUDTable) ItemCheckableOn(value interface{}) *CRUDTable {
	a.Set("itemCheckableOn", value)
	return a
}

func (a *CRUDTable) KeepItemSelectionOnPageChange(value interface{}) *CRUDTable {
	a.Set("keepItemSelectionOnPageChange", value)
	return a
}

func (a *CRUDTable) LabelTpl(value interface{}) *CRUDTable {
	a.Set("labelTpl", value)
	return a
}

func (a *CRUDTable) LoadDataOnce(value interface{}) *CRUDTable {
	a.Set("loadDataOnce", value)
	return a
}

func (a *CRUDTable) LoadDataOnceFetchOnFilter(value interface{}) *CRUDTable {
	a.Set("loadDataOnceFetchOnFilter", value)
	return a
}

func (a *CRUDTable) LoadMoreProps(value interface{}) *CRUDTable {
	a.Set("loadMoreProps", value)
	return a
}

func (a *CRUDTable) LoadingConfig(value interface{}) *CRUDTable {
	a.Set("loadingConfig", value)
	return a
}

func (a *CRUDTable) MatchFunc(value interface{}) *CRUDTable {
	a.Set("matchFunc", value)
	return a
}

func (a *CRUDTable) Messages(value interface{}) *CRUDTable {
	a.Set("messages", value)
	return a
}

func (a *CRUDTable) Mode(value interface{}) *CRUDTable {
	a.Set("mode", value)
	return a
}

func (a *CRUDTable) Multiple(value interface{}) *CRUDTable {
	a.Set("multiple", value)
	return a
}

func (a *CRUDTable) Name(value interface{}) *CRUDTable {
	a.Set("name", value)
	return a
}

func (a *CRUDTable) OnEvent(value interface{}) *CRUDTable {
	a.Set("onEvent", value)
	return a
}

func (a *CRUDTable) OrderBy(value interface{}) *CRUDTable {
	a.Set("orderBy", value)
	return a
}

func (a *CRUDTable) OrderDir(value interface{}) *CRUDTable {
	a.Set("orderDir", value)
	return a
}

func (a *CRUDTable) OrderField(value interface{}) *CRUDTable {
	a.Set("orderField", value)
	return a
}

func (a *CRUDTable) PageDirectionField(value interface{}) *CRUDTable {
	a.Set("pageDirectionField", value)
	return a
}

func (a *CRUDTable) PageField(value interface{}) *CRUDTable {
	a.Set("pageField", value)
	return a
}

func (a *CRUDTable) ParsePrimitiveQuery(value interface{}) *CRUDTable {
	a.Set("parsePrimitiveQuery", value)
	return a
}

func (a *CRUDTable) PerPage(value interface{}) *CRUDTable {
	a.Set("perPage", value)
	return a
}

func (a *CRUDTable) PerPageAvailable(value interface{}) *CRUDTable {
	a.Set("perPageAvailable", value)
	return a
}

func (a *CRUDTable) PerPageField(value interface{}) *CRUDTable {
	a.Set("perPageField", value)
	return a
}

func (a *CRUDTable) Placeholder(value interface{}) *CRUDTable {
	a.Set("placeholder", value)
	return a
}

func (a *CRUDTable) PrefixRow(value interface{}) *CRUDTable {
	a.Set("prefixRow", value)
	return a
}

func (a *CRUDTable) QuickSaveApi(value interface{}) *CRUDTable {
	a.Set("quickSaveApi", value)
	return a
}

func (a *CRUDTable) QuickSaveItemApi(value interface{}) *CRUDTable {
	a.Set("quickSaveItemApi", value)
	return a
}

func (a *CRUDTable) Resizable(value interface{}) *CRUDTable {
	a.Set("resizable", value)
	return a
}

func (a *CRUDTable) RowClassNameExpr(value interface{}) *CRUDTable {
	a.Set("rowClassNameExpr", value)
	return a
}

func (a *CRUDTable) SaveOrderApi(value interface{}) *CRUDTable {
	a.Set("saveOrderApi", value)
	return a
}

func (a *CRUDTable) Selectable(value interface{}) *CRUDTable {
	a.Set("selectable", value)
	return a
}

func (a *CRUDTable) ShowFooter(value interface{}) *CRUDTable {
	a.Set("showFooter", value)
	return a
}

func (a *CRUDTable) ShowHeader(value interface{}) *CRUDTable {
	a.Set("showHeader", value)
	return a
}

func (a *CRUDTable) ShowIndex(value interface{}) *CRUDTable {
	a.Set("showIndex", value)
	return a
}

func (a *CRUDTable) SilentPolling(value interface{}) *CRUDTable {
	a.Set("silentPolling", value)
	return a
}

func (a *CRUDTable) Source(value interface{}) *CRUDTable {
	a.Set("source", value)
	return a
}

func (a *CRUDTable) Static(value interface{}) *CRUDTable {
	a.Set("static", value)
	return a
}

func (a *CRUDTable) StaticClassName(value interface{}) *CRUDTable {
	a.Set("staticClassName", value)
	return a
}

func (a *CRUDTable) StaticInputClassName(value interface{}) *CRUDTable {
	a.Set("staticInputClassName", value)
	return a
}

func (a *CRUDTable) StaticLabelClassName(value interface{}) *CRUDTable {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *CRUDTable) StaticOn(value interface{}) *CRUDTable {
	a.Set("staticOn", value)
	return a
}

func (a *CRUDTable) StaticPlaceholder(value interface{}) *CRUDTable {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *CRUDTable) StaticSchema(value interface{}) *CRUDTable {
	a.Set("staticSchema", value)
	return a
}

func (a *CRUDTable) StopAutoRefreshWhen(value interface{}) *CRUDTable {
	a.Set("stopAutoRefreshWhen", value)
	return a
}

func (a *CRUDTable) StopAutoRefreshWhenModalIsOpen(value interface{}) *CRUDTable {
	a.Set("stopAutoRefreshWhenModalIsOpen", value)
	return a
}

func (a *CRUDTable) Style(value interface{}) *CRUDTable {
	a.Set("style", value)
	return a
}

func (a *CRUDTable) SyncLocation(value interface{}) *CRUDTable {
	a.Set("syncLocation", value)
	return a
}

func (a *CRUDTable) SyncResponse2Query(value interface{}) *CRUDTable {
	a.Set("syncResponse2Query", value)
	return a
}

func (a *CRUDTable) TableClassName(value interface{}) *CRUDTable {
	a.Set("tableClassName", value)
	return a
}

func (a *CRUDTable) TableLayout(value interface{}) *CRUDTable {
	a.Set("tableLayout", value)
	return a
}

func (a *CRUDTable) TestIdBuilder(value interface{}) *CRUDTable {
	a.Set("testIdBuilder", value)
	return a
}

func (a *CRUDTable) Testid(value interface{}) *CRUDTable {
	a.Set("testid", value)
	return a
}

func (a *CRUDTable) Title(value interface{}) *CRUDTable {
	a.Set("title", value)
	return a
}

func (a *CRUDTable) Toolbar(value interface{}) *CRUDTable {
	a.Set("toolbar", value)
	return a
}

func (a *CRUDTable) ToolbarClassName(value interface{}) *CRUDTable {
	a.Set("toolbarClassName", value)
	return a
}

func (a *CRUDTable) ToolbarInline(value interface{}) *CRUDTable {
	a.Set("toolbarInline", value)
	return a
}

func (a *CRUDTable) TotalField(value interface{}) *CRUDTable {
	a.Set("totalField", value)
	return a
}

func (a *CRUDTable) Type(value interface{}) *CRUDTable {
	a.Set("type", value)
	return a
}

func (a *CRUDTable) UseMobileUI(value interface{}) *CRUDTable {
	a.Set("useMobileUI", value)
	return a
}

func (a *CRUDTable) Visible(value interface{}) *CRUDTable {
	a.Set("visible", value)
	return a
}

func (a *CRUDTable) VisibleOn(value interface{}) *CRUDTable {
	a.Set("visibleOn", value)
	return a
}
