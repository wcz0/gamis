package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type CRUDCards struct {
	*BaseRenderer
}

func NewCRUDCards() *CRUDCards {
    a := &CRUDCards{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("mode", "cards")
    a.Set("type", "crud")
    return a
}


func (a *CRUDCards) Set(name string, value interface{}) *CRUDCards {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * name
 */
func (a *CRUDCards) Name(value interface{}) *CRUDCards {
    a.Set("name", value)
    return a
}

/**
 * stopAutoRefreshWhen
 */
func (a *CRUDCards) StopAutoRefreshWhen(value interface{}) *CRUDCards {
    a.Set("stopAutoRefreshWhen", value)
    return a
}

/**
 * filterTogglable
 */
func (a *CRUDCards) FilterTogglable(value interface{}) *CRUDCards {
    a.Set("filterTogglable", value)
    return a
}

/**
 * filterDefaultVisible
 */
func (a *CRUDCards) FilterDefaultVisible(value interface{}) *CRUDCards {
    a.Set("filterDefaultVisible", value)
    return a
}

/**
 * expandConfig
 */
func (a *CRUDCards) ExpandConfig(value interface{}) *CRUDCards {
    a.Set("expandConfig", value)
    return a
}

/**
 * itemClassName
 */
func (a *CRUDCards) ItemClassName(value interface{}) *CRUDCards {
    a.Set("itemClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *CRUDCards) StaticInputClassName(value interface{}) *CRUDCards {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * card
 */
func (a *CRUDCards) Card(value interface{}) *CRUDCards {
    a.Set("card", value)
    return a
}

/**
 * header
 */
func (a *CRUDCards) Header(value interface{}) *CRUDCards {
    a.Set("header", value)
    return a
}

/**
 * itemActions
 */
func (a *CRUDCards) ItemActions(value interface{}) *CRUDCards {
    a.Set("itemActions", value)
    return a
}

/**
 * innerClassName
 */
func (a *CRUDCards) InnerClassName(value interface{}) *CRUDCards {
    a.Set("innerClassName", value)
    return a
}

/**
 * quickSaveItemApi
 */
func (a *CRUDCards) QuickSaveItemApi(value interface{}) *CRUDCards {
    a.Set("quickSaveItemApi", value)
    return a
}

/**
 * toolbar
 */
func (a *CRUDCards) Toolbar(value interface{}) *CRUDCards {
    a.Set("toolbar", value)
    return a
}

/**
 * static
 */
func (a *CRUDCards) Static(value interface{}) *CRUDCards {
    a.Set("static", value)
    return a
}

/**
 * title
 */
func (a *CRUDCards) Title(value interface{}) *CRUDCards {
    a.Set("title", value)
    return a
}

/**
 * toolbarInline
 */
func (a *CRUDCards) ToolbarInline(value interface{}) *CRUDCards {
    a.Set("toolbarInline", value)
    return a
}

/**
 * perPageAvailable
 */
func (a *CRUDCards) PerPageAvailable(value interface{}) *CRUDCards {
    a.Set("perPageAvailable", value)
    return a
}

/**
 * alwaysShowPagination
 */
func (a *CRUDCards) AlwaysShowPagination(value interface{}) *CRUDCards {
    a.Set("alwaysShowPagination", value)
    return a
}

/**
 * loadMoreProps
 */
func (a *CRUDCards) LoadMoreProps(value interface{}) *CRUDCards {
    a.Set("loadMoreProps", value)
    return a
}

/**
 * useMobileUI
 */
func (a *CRUDCards) UseMobileUI(value interface{}) *CRUDCards {
    a.Set("useMobileUI", value)
    return a
}

/**
 * showFooter
 */
func (a *CRUDCards) ShowFooter(value interface{}) *CRUDCards {
    a.Set("showFooter", value)
    return a
}

/**
 * labelTpl
 */
func (a *CRUDCards) LabelTpl(value interface{}) *CRUDCards {
    a.Set("labelTpl", value)
    return a
}

/**
 * autoGenerateFilter
 */
func (a *CRUDCards) AutoGenerateFilter(value interface{}) *CRUDCards {
    a.Set("autoGenerateFilter", value)
    return a
}

/**
 * draggableOn
 */
func (a *CRUDCards) DraggableOn(value interface{}) *CRUDCards {
    a.Set("draggableOn", value)
    return a
}

/**
 * style
 */
func (a *CRUDCards) Style(value interface{}) *CRUDCards {
    a.Set("style", value)
    return a
}

/**
 * affixFooter
 */
func (a *CRUDCards) AffixFooter(value interface{}) *CRUDCards {
    a.Set("affixFooter", value)
    return a
}

/**
 * bulkActions
 */
func (a *CRUDCards) BulkActions(value interface{}) *CRUDCards {
    a.Set("bulkActions", value)
    return a
}

/**
 * draggable
 */
func (a *CRUDCards) Draggable(value interface{}) *CRUDCards {
    a.Set("draggable", value)
    return a
}

/**
 * initFetch
 */
func (a *CRUDCards) InitFetch(value interface{}) *CRUDCards {
    a.Set("initFetch", value)
    return a
}

/**
 * quickSaveApi
 */
func (a *CRUDCards) QuickSaveApi(value interface{}) *CRUDCards {
    a.Set("quickSaveApi", value)
    return a
}

/**
 * className
 */
func (a *CRUDCards) ClassName(value interface{}) *CRUDCards {
    a.Set("className", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *CRUDCards) StaticPlaceholder(value interface{}) *CRUDCards {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *CRUDCards) StaticLabelClassName(value interface{}) *CRUDCards {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * masonryLayout
 */
func (a *CRUDCards) MasonryLayout(value interface{}) *CRUDCards {
    a.Set("masonryLayout", value)
    return a
}

/**
 * saveOrderApi
 */
func (a *CRUDCards) SaveOrderApi(value interface{}) *CRUDCards {
    a.Set("saveOrderApi", value)
    return a
}

/**
 * syncLocation
 */
func (a *CRUDCards) SyncLocation(value interface{}) *CRUDCards {
    a.Set("syncLocation", value)
    return a
}

/**
 * parsePrimitiveQuery
 */
func (a *CRUDCards) ParsePrimitiveQuery(value interface{}) *CRUDCards {
    a.Set("parsePrimitiveQuery", value)
    return a
}

/**
 * selectable
 */
func (a *CRUDCards) Selectable(value interface{}) *CRUDCards {
    a.Set("selectable", value)
    return a
}

/**
 * hiddenOn
 */
func (a *CRUDCards) HiddenOn(value interface{}) *CRUDCards {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visibleOn
 */
func (a *CRUDCards) VisibleOn(value interface{}) *CRUDCards {
    a.Set("visibleOn", value)
    return a
}

/**
 * source
 */
func (a *CRUDCards) Source(value interface{}) *CRUDCards {
    a.Set("source", value)
    return a
}

/**
 * api
 */
func (a *CRUDCards) Api(value interface{}) *CRUDCards {
    a.Set("api", value)
    return a
}

/**
 * syncResponse2Query
 */
func (a *CRUDCards) SyncResponse2Query(value interface{}) *CRUDCards {
    a.Set("syncResponse2Query", value)
    return a
}

/**
 * perPage
 */
func (a *CRUDCards) PerPage(value interface{}) *CRUDCards {
    a.Set("perPage", value)
    return a
}

/**
 * checkOnItemClick
 */
func (a *CRUDCards) CheckOnItemClick(value interface{}) *CRUDCards {
    a.Set("checkOnItemClick", value)
    return a
}

/**
 * interval
 */
func (a *CRUDCards) Interval(value interface{}) *CRUDCards {
    a.Set("interval", value)
    return a
}

/**
 * orderField
 */
func (a *CRUDCards) OrderField(value interface{}) *CRUDCards {
    a.Set("orderField", value)
    return a
}

/**
 * autoJumpToTopOnPagerChange
 */
func (a *CRUDCards) AutoJumpToTopOnPagerChange(value interface{}) *CRUDCards {
    a.Set("autoJumpToTopOnPagerChange", value)
    return a
}

/**
 * onEvent
 */
func (a *CRUDCards) OnEvent(value interface{}) *CRUDCards {
    a.Set("onEvent", value)
    return a
}

/**
 * showHeader
 */
func (a *CRUDCards) ShowHeader(value interface{}) *CRUDCards {
    a.Set("showHeader", value)
    return a
}

/**
 * orderDir
 */
func (a *CRUDCards) OrderDir(value interface{}) *CRUDCards {
    a.Set("orderDir", value)
    return a
}

/**
 * maxKeepItemSelectionLength
 */
func (a *CRUDCards) MaxKeepItemSelectionLength(value interface{}) *CRUDCards {
    a.Set("maxKeepItemSelectionLength", value)
    return a
}

/**
 * disabled
 */
func (a *CRUDCards) Disabled(value interface{}) *CRUDCards {
    a.Set("disabled", value)
    return a
}

/**
 * id
 */
func (a *CRUDCards) Id(value interface{}) *CRUDCards {
    a.Set("id", value)
    return a
}

/**
 * staticOn
 */
func (a *CRUDCards) StaticOn(value interface{}) *CRUDCards {
    a.Set("staticOn", value)
    return a
}

/**
 * footerClassName
 */
func (a *CRUDCards) FooterClassName(value interface{}) *CRUDCards {
    a.Set("footerClassName", value)
    return a
}

/**
 * perPageField
 */
func (a *CRUDCards) PerPageField(value interface{}) *CRUDCards {
    a.Set("perPageField", value)
    return a
}

/**
 * totalField
 */
func (a *CRUDCards) TotalField(value interface{}) *CRUDCards {
    a.Set("totalField", value)
    return a
}

/**
 * hideQuickSaveBtn
 */
func (a *CRUDCards) HideQuickSaveBtn(value interface{}) *CRUDCards {
    a.Set("hideQuickSaveBtn", value)
    return a
}

/**
 * silentPolling
 */
func (a *CRUDCards) SilentPolling(value interface{}) *CRUDCards {
    a.Set("silentPolling", value)
    return a
}

/**
 * placeholder
 */
func (a *CRUDCards) Placeholder(value interface{}) *CRUDCards {
    a.Set("placeholder", value)
    return a
}

/**
 * itemDraggableOn
 */
func (a *CRUDCards) ItemDraggableOn(value interface{}) *CRUDCards {
    a.Set("itemDraggableOn", value)
    return a
}

/**
 * filter
 */
func (a *CRUDCards) Filter(value interface{}) *CRUDCards {
    a.Set("filter", value)
    return a
}

/**
 * pageField
 */
func (a *CRUDCards) PageField(value interface{}) *CRUDCards {
    a.Set("pageField", value)
    return a
}

/**
 * loadDataOnceFetchOnFilter
 */
func (a *CRUDCards) LoadDataOnceFetchOnFilter(value interface{}) *CRUDCards {
    a.Set("loadDataOnceFetchOnFilter", value)
    return a
}

/**
 * disabledOn
 */
func (a *CRUDCards) DisabledOn(value interface{}) *CRUDCards {
    a.Set("disabledOn", value)
    return a
}

/**
 * editorSetting
 */
func (a *CRUDCards) EditorSetting(value interface{}) *CRUDCards {
    a.Set("editorSetting", value)
    return a
}

/**
 * affixHeader
 */
func (a *CRUDCards) AffixHeader(value interface{}) *CRUDCards {
    a.Set("affixHeader", value)
    return a
}

/**
 * valueField
 */
func (a *CRUDCards) ValueField(value interface{}) *CRUDCards {
    a.Set("valueField", value)
    return a
}

/**
 * headerToolbar
 */
func (a *CRUDCards) HeaderToolbar(value interface{}) *CRUDCards {
    a.Set("headerToolbar", value)
    return a
}

/**
 * footerToolbar
 */
func (a *CRUDCards) FooterToolbar(value interface{}) *CRUDCards {
    a.Set("footerToolbar", value)
    return a
}

/**
 * keepItemSelectionOnPageChange
 */
func (a *CRUDCards) KeepItemSelectionOnPageChange(value interface{}) *CRUDCards {
    a.Set("keepItemSelectionOnPageChange", value)
    return a
}

/**
 * matchFunc
 */
func (a *CRUDCards) MatchFunc(value interface{}) *CRUDCards {
    a.Set("matchFunc", value)
    return a
}

/**
 * headerClassName
 */
func (a *CRUDCards) HeaderClassName(value interface{}) *CRUDCards {
    a.Set("headerClassName", value)
    return a
}

/**
 * visible
 */
func (a *CRUDCards) Visible(value interface{}) *CRUDCards {
    a.Set("visible", value)
    return a
}

/**
 * testid
 */
func (a *CRUDCards) Testid(value interface{}) *CRUDCards {
    a.Set("testid", value)
    return a
}

/**
 * orderBy
 */
func (a *CRUDCards) OrderBy(value interface{}) *CRUDCards {
    a.Set("orderBy", value)
    return a
}

/**
 * initFetchOn
 */
func (a *CRUDCards) InitFetchOn(value interface{}) *CRUDCards {
    a.Set("initFetchOn", value)
    return a
}

/**
 * loadDataOnce
 */
func (a *CRUDCards) LoadDataOnce(value interface{}) *CRUDCards {
    a.Set("loadDataOnce", value)
    return a
}

/**
 * autoFillHeight
 */
func (a *CRUDCards) AutoFillHeight(value interface{}) *CRUDCards {
    a.Set("autoFillHeight", value)
    return a
}

/**
 * multiple
 */
func (a *CRUDCards) Multiple(value interface{}) *CRUDCards {
    a.Set("multiple", value)
    return a
}

/**
 * loadingConfig
 */
func (a *CRUDCards) LoadingConfig(value interface{}) *CRUDCards {
    a.Set("loadingConfig", value)
    return a
}

/**
 */
func (a *CRUDCards) Mode(value interface{}) *CRUDCards {
    a.Set("mode", value)
    return a
}

/**
 * defaultParams
 */
func (a *CRUDCards) DefaultParams(value interface{}) *CRUDCards {
    a.Set("defaultParams", value)
    return a
}

/**
 * pageDirectionField
 */
func (a *CRUDCards) PageDirectionField(value interface{}) *CRUDCards {
    a.Set("pageDirectionField", value)
    return a
}

/**
 * hidden
 */
func (a *CRUDCards) Hidden(value interface{}) *CRUDCards {
    a.Set("hidden", value)
    return a
}

/**
 * staticClassName
 */
func (a *CRUDCards) StaticClassName(value interface{}) *CRUDCards {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *CRUDCards) StaticSchema(value interface{}) *CRUDCards {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *CRUDCards) Type(value interface{}) *CRUDCards {
    a.Set("type", value)
    return a
}

/**
 * itemCheckableOn
 */
func (a *CRUDCards) ItemCheckableOn(value interface{}) *CRUDCards {
    a.Set("itemCheckableOn", value)
    return a
}

/**
 * deferApi
 */
func (a *CRUDCards) DeferApi(value interface{}) *CRUDCards {
    a.Set("deferApi", value)
    return a
}

/**
 * messages
 */
func (a *CRUDCards) Messages(value interface{}) *CRUDCards {
    a.Set("messages", value)
    return a
}

/**
 * stopAutoRefreshWhenModalIsOpen
 */
func (a *CRUDCards) StopAutoRefreshWhenModalIsOpen(value interface{}) *CRUDCards {
    a.Set("stopAutoRefreshWhenModalIsOpen", value)
    return a
}

/**
 * hideCheckToggler
 */
func (a *CRUDCards) HideCheckToggler(value interface{}) *CRUDCards {
    a.Set("hideCheckToggler", value)
    return a
}

/**
 * footer
 */
func (a *CRUDCards) Footer(value interface{}) *CRUDCards {
    a.Set("footer", value)
    return a
}
