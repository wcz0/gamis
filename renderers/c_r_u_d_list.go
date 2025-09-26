package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type CRUDList struct {
	*BaseRenderer
}

func NewCRUDList() *CRUDList {
    a := &CRUDList{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("mode", "list")
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
/**
 * header
 */
func (a *CRUDList) Header(value interface{}) *CRUDList {
    a.Set("header", value)
    return a
}

/**
 * headerClassName
 */
func (a *CRUDList) HeaderClassName(value interface{}) *CRUDList {
    a.Set("headerClassName", value)
    return a
}

/**
 * footerToolbar
 */
func (a *CRUDList) FooterToolbar(value interface{}) *CRUDList {
    a.Set("footerToolbar", value)
    return a
}

/**
 * loadDataOnce
 */
func (a *CRUDList) LoadDataOnce(value interface{}) *CRUDList {
    a.Set("loadDataOnce", value)
    return a
}

/**
 * visibleOn
 */
func (a *CRUDList) VisibleOn(value interface{}) *CRUDList {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticOn
 */
func (a *CRUDList) StaticOn(value interface{}) *CRUDList {
    a.Set("staticOn", value)
    return a
}

/**
 * listItem
 */
func (a *CRUDList) ListItem(value interface{}) *CRUDList {
    a.Set("listItem", value)
    return a
}

/**
 * showIndexBar
 */
func (a *CRUDList) ShowIndexBar(value interface{}) *CRUDList {
    a.Set("showIndexBar", value)
    return a
}

/**
 */
func (a *CRUDList) Mode(value interface{}) *CRUDList {
    a.Set("mode", value)
    return a
}

/**
 * pageField
 */
func (a *CRUDList) PageField(value interface{}) *CRUDList {
    a.Set("pageField", value)
    return a
}

/**
 * disabled
 */
func (a *CRUDList) Disabled(value interface{}) *CRUDList {
    a.Set("disabled", value)
    return a
}

/**
 * itemActions
 */
func (a *CRUDList) ItemActions(value interface{}) *CRUDList {
    a.Set("itemActions", value)
    return a
}

/**
 * perPageAvailable
 */
func (a *CRUDList) PerPageAvailable(value interface{}) *CRUDList {
    a.Set("perPageAvailable", value)
    return a
}

/**
 * visible
 */
func (a *CRUDList) Visible(value interface{}) *CRUDList {
    a.Set("visible", value)
    return a
}

/**
 * footerClassName
 */
func (a *CRUDList) FooterClassName(value interface{}) *CRUDList {
    a.Set("footerClassName", value)
    return a
}

/**
 * checkOnItemClick
 */
func (a *CRUDList) CheckOnItemClick(value interface{}) *CRUDList {
    a.Set("checkOnItemClick", value)
    return a
}

/**
 * indexField
 */
func (a *CRUDList) IndexField(value interface{}) *CRUDList {
    a.Set("indexField", value)
    return a
}

/**
 * syncLocation
 */
func (a *CRUDList) SyncLocation(value interface{}) *CRUDList {
    a.Set("syncLocation", value)
    return a
}

/**
 * stopAutoRefreshWhen
 */
func (a *CRUDList) StopAutoRefreshWhen(value interface{}) *CRUDList {
    a.Set("stopAutoRefreshWhen", value)
    return a
}

/**
 * autoFillHeight
 */
func (a *CRUDList) AutoFillHeight(value interface{}) *CRUDList {
    a.Set("autoFillHeight", value)
    return a
}

/**
 * disabledOn
 */
func (a *CRUDList) DisabledOn(value interface{}) *CRUDList {
    a.Set("disabledOn", value)
    return a
}

/**
 * static
 */
func (a *CRUDList) Static(value interface{}) *CRUDList {
    a.Set("static", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *CRUDList) StaticLabelClassName(value interface{}) *CRUDList {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * hideCheckToggler
 */
func (a *CRUDList) HideCheckToggler(value interface{}) *CRUDList {
    a.Set("hideCheckToggler", value)
    return a
}

/**
 * affixFooter
 */
func (a *CRUDList) AffixFooter(value interface{}) *CRUDList {
    a.Set("affixFooter", value)
    return a
}

/**
 * selectable
 */
func (a *CRUDList) Selectable(value interface{}) *CRUDList {
    a.Set("selectable", value)
    return a
}

/**
 * itemAction
 */
func (a *CRUDList) ItemAction(value interface{}) *CRUDList {
    a.Set("itemAction", value)
    return a
}

/**
 * api
 */
func (a *CRUDList) Api(value interface{}) *CRUDList {
    a.Set("api", value)
    return a
}

/**
 * deferApi
 */
func (a *CRUDList) DeferApi(value interface{}) *CRUDList {
    a.Set("deferApi", value)
    return a
}

/**
 * name
 */
func (a *CRUDList) Name(value interface{}) *CRUDList {
    a.Set("name", value)
    return a
}

/**
 * innerClassName
 */
func (a *CRUDList) InnerClassName(value interface{}) *CRUDList {
    a.Set("innerClassName", value)
    return a
}

/**
 * filterDefaultVisible
 */
func (a *CRUDList) FilterDefaultVisible(value interface{}) *CRUDList {
    a.Set("filterDefaultVisible", value)
    return a
}

/**
 * style
 */
func (a *CRUDList) Style(value interface{}) *CRUDList {
    a.Set("style", value)
    return a
}

/**
 * editorSetting
 */
func (a *CRUDList) EditorSetting(value interface{}) *CRUDList {
    a.Set("editorSetting", value)
    return a
}

/**
 * footer
 */
func (a *CRUDList) Footer(value interface{}) *CRUDList {
    a.Set("footer", value)
    return a
}

/**
 * size
 */
func (a *CRUDList) Size(value interface{}) *CRUDList {
    a.Set("size", value)
    return a
}

/**
 * initFetch
 */
func (a *CRUDList) InitFetch(value interface{}) *CRUDList {
    a.Set("initFetch", value)
    return a
}

/**
 * quickSaveApi
 */
func (a *CRUDList) QuickSaveApi(value interface{}) *CRUDList {
    a.Set("quickSaveApi", value)
    return a
}

/**
 * autoJumpToTopOnPagerChange
 */
func (a *CRUDList) AutoJumpToTopOnPagerChange(value interface{}) *CRUDList {
    a.Set("autoJumpToTopOnPagerChange", value)
    return a
}

/**
 * affixHeader
 */
func (a *CRUDList) AffixHeader(value interface{}) *CRUDList {
    a.Set("affixHeader", value)
    return a
}

/**
 * indexBarOffset
 */
func (a *CRUDList) IndexBarOffset(value interface{}) *CRUDList {
    a.Set("indexBarOffset", value)
    return a
}

/**
 * orderDir
 */
func (a *CRUDList) OrderDir(value interface{}) *CRUDList {
    a.Set("orderDir", value)
    return a
}

/**
 * toolbarInline
 */
func (a *CRUDList) ToolbarInline(value interface{}) *CRUDList {
    a.Set("toolbarInline", value)
    return a
}

/**
 * autoGenerateFilter
 */
func (a *CRUDList) AutoGenerateFilter(value interface{}) *CRUDList {
    a.Set("autoGenerateFilter", value)
    return a
}

/**
 * loadMoreProps
 */
func (a *CRUDList) LoadMoreProps(value interface{}) *CRUDList {
    a.Set("loadMoreProps", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *CRUDList) StaticInputClassName(value interface{}) *CRUDList {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *CRUDList) StaticSchema(value interface{}) *CRUDList {
    a.Set("staticSchema", value)
    return a
}

/**
 * itemCheckableOn
 */
func (a *CRUDList) ItemCheckableOn(value interface{}) *CRUDList {
    a.Set("itemCheckableOn", value)
    return a
}

/**
 * perPage
 */
func (a *CRUDList) PerPage(value interface{}) *CRUDList {
    a.Set("perPage", value)
    return a
}

/**
 * draggable
 */
func (a *CRUDList) Draggable(value interface{}) *CRUDList {
    a.Set("draggable", value)
    return a
}

/**
 * filter
 */
func (a *CRUDList) Filter(value interface{}) *CRUDList {
    a.Set("filter", value)
    return a
}

/**
 * orderField
 */
func (a *CRUDList) OrderField(value interface{}) *CRUDList {
    a.Set("orderField", value)
    return a
}

/**
 * totalField
 */
func (a *CRUDList) TotalField(value interface{}) *CRUDList {
    a.Set("totalField", value)
    return a
}

/**
 * id
 */
func (a *CRUDList) Id(value interface{}) *CRUDList {
    a.Set("id", value)
    return a
}

/**
 * testid
 */
func (a *CRUDList) Testid(value interface{}) *CRUDList {
    a.Set("testid", value)
    return a
}

/**
 * filterTogglable
 */
func (a *CRUDList) FilterTogglable(value interface{}) *CRUDList {
    a.Set("filterTogglable", value)
    return a
}

/**
 * matchFunc
 */
func (a *CRUDList) MatchFunc(value interface{}) *CRUDList {
    a.Set("matchFunc", value)
    return a
}

/**
 * hiddenOn
 */
func (a *CRUDList) HiddenOn(value interface{}) *CRUDList {
    a.Set("hiddenOn", value)
    return a
}

/**
 * placeholder
 */
func (a *CRUDList) Placeholder(value interface{}) *CRUDList {
    a.Set("placeholder", value)
    return a
}

/**
 * orderBy
 */
func (a *CRUDList) OrderBy(value interface{}) *CRUDList {
    a.Set("orderBy", value)
    return a
}

/**
 * defaultParams
 */
func (a *CRUDList) DefaultParams(value interface{}) *CRUDList {
    a.Set("defaultParams", value)
    return a
}

/**
 * toolbar
 */
func (a *CRUDList) Toolbar(value interface{}) *CRUDList {
    a.Set("toolbar", value)
    return a
}

/**
 * headerToolbar
 */
func (a *CRUDList) HeaderToolbar(value interface{}) *CRUDList {
    a.Set("headerToolbar", value)
    return a
}

/**
 * stopAutoRefreshWhenModalIsOpen
 */
func (a *CRUDList) StopAutoRefreshWhenModalIsOpen(value interface{}) *CRUDList {
    a.Set("stopAutoRefreshWhenModalIsOpen", value)
    return a
}

/**
 * multiple
 */
func (a *CRUDList) Multiple(value interface{}) *CRUDList {
    a.Set("multiple", value)
    return a
}

/**
 * onEvent
 */
func (a *CRUDList) OnEvent(value interface{}) *CRUDList {
    a.Set("onEvent", value)
    return a
}

/**
 * showHeader
 */
func (a *CRUDList) ShowHeader(value interface{}) *CRUDList {
    a.Set("showHeader", value)
    return a
}

/**
 * itemDraggableOn
 */
func (a *CRUDList) ItemDraggableOn(value interface{}) *CRUDList {
    a.Set("itemDraggableOn", value)
    return a
}

/**
 * quickSaveItemApi
 */
func (a *CRUDList) QuickSaveItemApi(value interface{}) *CRUDList {
    a.Set("quickSaveItemApi", value)
    return a
}

/**
 * saveOrderApi
 */
func (a *CRUDList) SaveOrderApi(value interface{}) *CRUDList {
    a.Set("saveOrderApi", value)
    return a
}

/**
 * hideQuickSaveBtn
 */
func (a *CRUDList) HideQuickSaveBtn(value interface{}) *CRUDList {
    a.Set("hideQuickSaveBtn", value)
    return a
}

/**
 * silentPolling
 */
func (a *CRUDList) SilentPolling(value interface{}) *CRUDList {
    a.Set("silentPolling", value)
    return a
}

/**
 * expandConfig
 */
func (a *CRUDList) ExpandConfig(value interface{}) *CRUDList {
    a.Set("expandConfig", value)
    return a
}

/**
 * title
 */
func (a *CRUDList) Title(value interface{}) *CRUDList {
    a.Set("title", value)
    return a
}

/**
 * showFooter
 */
func (a *CRUDList) ShowFooter(value interface{}) *CRUDList {
    a.Set("showFooter", value)
    return a
}

/**
 * interval
 */
func (a *CRUDList) Interval(value interface{}) *CRUDList {
    a.Set("interval", value)
    return a
}

/**
 * loadDataOnceFetchOnFilter
 */
func (a *CRUDList) LoadDataOnceFetchOnFilter(value interface{}) *CRUDList {
    a.Set("loadDataOnceFetchOnFilter", value)
    return a
}

/**
 * staticClassName
 */
func (a *CRUDList) StaticClassName(value interface{}) *CRUDList {
    a.Set("staticClassName", value)
    return a
}

/**
 * valueField
 */
func (a *CRUDList) ValueField(value interface{}) *CRUDList {
    a.Set("valueField", value)
    return a
}

/**
 * bulkActions
 */
func (a *CRUDList) BulkActions(value interface{}) *CRUDList {
    a.Set("bulkActions", value)
    return a
}

/**
 * initFetchOn
 */
func (a *CRUDList) InitFetchOn(value interface{}) *CRUDList {
    a.Set("initFetchOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *CRUDList) StaticPlaceholder(value interface{}) *CRUDList {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *CRUDList) Type(value interface{}) *CRUDList {
    a.Set("type", value)
    return a
}

/**
 * source
 */
func (a *CRUDList) Source(value interface{}) *CRUDList {
    a.Set("source", value)
    return a
}

/**
 * draggableOn
 */
func (a *CRUDList) DraggableOn(value interface{}) *CRUDList {
    a.Set("draggableOn", value)
    return a
}

/**
 * perPageField
 */
func (a *CRUDList) PerPageField(value interface{}) *CRUDList {
    a.Set("perPageField", value)
    return a
}

/**
 * pageDirectionField
 */
func (a *CRUDList) PageDirectionField(value interface{}) *CRUDList {
    a.Set("pageDirectionField", value)
    return a
}

/**
 * messages
 */
func (a *CRUDList) Messages(value interface{}) *CRUDList {
    a.Set("messages", value)
    return a
}

/**
 * syncResponse2Query
 */
func (a *CRUDList) SyncResponse2Query(value interface{}) *CRUDList {
    a.Set("syncResponse2Query", value)
    return a
}

/**
 * className
 */
func (a *CRUDList) ClassName(value interface{}) *CRUDList {
    a.Set("className", value)
    return a
}

/**
 * useMobileUI
 */
func (a *CRUDList) UseMobileUI(value interface{}) *CRUDList {
    a.Set("useMobileUI", value)
    return a
}

/**
 * keepItemSelectionOnPageChange
 */
func (a *CRUDList) KeepItemSelectionOnPageChange(value interface{}) *CRUDList {
    a.Set("keepItemSelectionOnPageChange", value)
    return a
}

/**
 * maxKeepItemSelectionLength
 */
func (a *CRUDList) MaxKeepItemSelectionLength(value interface{}) *CRUDList {
    a.Set("maxKeepItemSelectionLength", value)
    return a
}

/**
 * labelTpl
 */
func (a *CRUDList) LabelTpl(value interface{}) *CRUDList {
    a.Set("labelTpl", value)
    return a
}

/**
 * alwaysShowPagination
 */
func (a *CRUDList) AlwaysShowPagination(value interface{}) *CRUDList {
    a.Set("alwaysShowPagination", value)
    return a
}

/**
 * parsePrimitiveQuery
 */
func (a *CRUDList) ParsePrimitiveQuery(value interface{}) *CRUDList {
    a.Set("parsePrimitiveQuery", value)
    return a
}

/**
 * hidden
 */
func (a *CRUDList) Hidden(value interface{}) *CRUDList {
    a.Set("hidden", value)
    return a
}
