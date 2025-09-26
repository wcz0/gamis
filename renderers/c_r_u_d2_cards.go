package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type CRUD2Cards struct {
	*BaseRenderer
}

func NewCRUD2Cards() *CRUD2Cards {
    a := &CRUD2Cards{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "crud2")
    a.Set("mode", "cards")
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
/**
 * useMobileUI
 */
func (a *CRUD2Cards) UseMobileUI(value interface{}) *CRUD2Cards {
    a.Set("useMobileUI", value)
    return a
}

/**
 * card
 */
func (a *CRUD2Cards) Card(value interface{}) *CRUD2Cards {
    a.Set("card", value)
    return a
}

/**
 * valueField
 */
func (a *CRUD2Cards) ValueField(value interface{}) *CRUD2Cards {
    a.Set("valueField", value)
    return a
}

/**
 * loadType
 */
func (a *CRUD2Cards) LoadType(value interface{}) *CRUD2Cards {
    a.Set("loadType", value)
    return a
}

/**
 * showSelection
 */
func (a *CRUD2Cards) ShowSelection(value interface{}) *CRUD2Cards {
    a.Set("showSelection", value)
    return a
}

/**
 * quickSaveApi
 */
func (a *CRUD2Cards) QuickSaveApi(value interface{}) *CRUD2Cards {
    a.Set("quickSaveApi", value)
    return a
}

/**
 * name
 */
func (a *CRUD2Cards) Name(value interface{}) *CRUD2Cards {
    a.Set("name", value)
    return a
}

/**
 * syncResponse2Query
 */
func (a *CRUD2Cards) SyncResponse2Query(value interface{}) *CRUD2Cards {
    a.Set("syncResponse2Query", value)
    return a
}

/**
 * showHeader
 */
func (a *CRUD2Cards) ShowHeader(value interface{}) *CRUD2Cards {
    a.Set("showHeader", value)
    return a
}

/**
 * checkOnItemClick
 */
func (a *CRUD2Cards) CheckOnItemClick(value interface{}) *CRUD2Cards {
    a.Set("checkOnItemClick", value)
    return a
}

/**
 */
func (a *CRUD2Cards) Mode(value interface{}) *CRUD2Cards {
    a.Set("mode", value)
    return a
}

/**
 * syncLocation
 */
func (a *CRUD2Cards) SyncLocation(value interface{}) *CRUD2Cards {
    a.Set("syncLocation", value)
    return a
}

/**
 * pageField
 */
func (a *CRUD2Cards) PageField(value interface{}) *CRUD2Cards {
    a.Set("pageField", value)
    return a
}

/**
 * headerToolbarClassName
 */
func (a *CRUD2Cards) HeaderToolbarClassName(value interface{}) *CRUD2Cards {
    a.Set("headerToolbarClassName", value)
    return a
}

/**
 * hideCheckToggler
 */
func (a *CRUD2Cards) HideCheckToggler(value interface{}) *CRUD2Cards {
    a.Set("hideCheckToggler", value)
    return a
}

/**
 * silentPolling
 */
func (a *CRUD2Cards) SilentPolling(value interface{}) *CRUD2Cards {
    a.Set("silentPolling", value)
    return a
}

/**
 * stopAutoRefreshWhen
 */
func (a *CRUD2Cards) StopAutoRefreshWhen(value interface{}) *CRUD2Cards {
    a.Set("stopAutoRefreshWhen", value)
    return a
}

/**
 * footerToolbar
 */
func (a *CRUD2Cards) FooterToolbar(value interface{}) *CRUD2Cards {
    a.Set("footerToolbar", value)
    return a
}

/**
 * affixHeader
 */
func (a *CRUD2Cards) AffixHeader(value interface{}) *CRUD2Cards {
    a.Set("affixHeader", value)
    return a
}

/**
 * affixFooter
 */
func (a *CRUD2Cards) AffixFooter(value interface{}) *CRUD2Cards {
    a.Set("affixFooter", value)
    return a
}

/**
 * perPage
 */
func (a *CRUD2Cards) PerPage(value interface{}) *CRUD2Cards {
    a.Set("perPage", value)
    return a
}

/**
 * className
 */
func (a *CRUD2Cards) ClassName(value interface{}) *CRUD2Cards {
    a.Set("className", value)
    return a
}

/**
 * static
 */
func (a *CRUD2Cards) Static(value interface{}) *CRUD2Cards {
    a.Set("static", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *CRUD2Cards) StaticPlaceholder(value interface{}) *CRUD2Cards {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticSchema
 */
func (a *CRUD2Cards) StaticSchema(value interface{}) *CRUD2Cards {
    a.Set("staticSchema", value)
    return a
}

/**
 * header
 */
func (a *CRUD2Cards) Header(value interface{}) *CRUD2Cards {
    a.Set("header", value)
    return a
}

/**
 * interval
 */
func (a *CRUD2Cards) Interval(value interface{}) *CRUD2Cards {
    a.Set("interval", value)
    return a
}

/**
 * perPageField
 */
func (a *CRUD2Cards) PerPageField(value interface{}) *CRUD2Cards {
    a.Set("perPageField", value)
    return a
}

/**
 * disabledOn
 */
func (a *CRUD2Cards) DisabledOn(value interface{}) *CRUD2Cards {
    a.Set("disabledOn", value)
    return a
}

/**
 * visibleOn
 */
func (a *CRUD2Cards) VisibleOn(value interface{}) *CRUD2Cards {
    a.Set("visibleOn", value)
    return a
}

/**
 * footerClassName
 */
func (a *CRUD2Cards) FooterClassName(value interface{}) *CRUD2Cards {
    a.Set("footerClassName", value)
    return a
}

/**
 * autoFillHeight
 */
func (a *CRUD2Cards) AutoFillHeight(value interface{}) *CRUD2Cards {
    a.Set("autoFillHeight", value)
    return a
}

/**
 * loadingConfig
 */
func (a *CRUD2Cards) LoadingConfig(value interface{}) *CRUD2Cards {
    a.Set("loadingConfig", value)
    return a
}

/**
 * onEvent
 */
func (a *CRUD2Cards) OnEvent(value interface{}) *CRUD2Cards {
    a.Set("onEvent", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *CRUD2Cards) StaticInputClassName(value interface{}) *CRUD2Cards {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * testid
 */
func (a *CRUD2Cards) Testid(value interface{}) *CRUD2Cards {
    a.Set("testid", value)
    return a
}

/**
 * primaryField
 */
func (a *CRUD2Cards) PrimaryField(value interface{}) *CRUD2Cards {
    a.Set("primaryField", value)
    return a
}

/**
 * visible
 */
func (a *CRUD2Cards) Visible(value interface{}) *CRUD2Cards {
    a.Set("visible", value)
    return a
}

/**
 * itemDraggableOn
 */
func (a *CRUD2Cards) ItemDraggableOn(value interface{}) *CRUD2Cards {
    a.Set("itemDraggableOn", value)
    return a
}

/**
 * saveOrderApi
 */
func (a *CRUD2Cards) SaveOrderApi(value interface{}) *CRUD2Cards {
    a.Set("saveOrderApi", value)
    return a
}

/**
 * hideQuickSaveBtn
 */
func (a *CRUD2Cards) HideQuickSaveBtn(value interface{}) *CRUD2Cards {
    a.Set("hideQuickSaveBtn", value)
    return a
}

/**
 * keepItemSelectionOnPageChange
 */
func (a *CRUD2Cards) KeepItemSelectionOnPageChange(value interface{}) *CRUD2Cards {
    a.Set("keepItemSelectionOnPageChange", value)
    return a
}

/**
 * staticClassName
 */
func (a *CRUD2Cards) StaticClassName(value interface{}) *CRUD2Cards {
    a.Set("staticClassName", value)
    return a
}

/**
 * placeholder
 */
func (a *CRUD2Cards) Placeholder(value interface{}) *CRUD2Cards {
    a.Set("placeholder", value)
    return a
}

/**
 * multiple
 */
func (a *CRUD2Cards) Multiple(value interface{}) *CRUD2Cards {
    a.Set("multiple", value)
    return a
}

/**
 * autoJumpToTopOnPagerChange
 */
func (a *CRUD2Cards) AutoJumpToTopOnPagerChange(value interface{}) *CRUD2Cards {
    a.Set("autoJumpToTopOnPagerChange", value)
    return a
}

/**
 * showFooter
 */
func (a *CRUD2Cards) ShowFooter(value interface{}) *CRUD2Cards {
    a.Set("showFooter", value)
    return a
}

/**
 * source
 */
func (a *CRUD2Cards) Source(value interface{}) *CRUD2Cards {
    a.Set("source", value)
    return a
}

/**
 * title
 */
func (a *CRUD2Cards) Title(value interface{}) *CRUD2Cards {
    a.Set("title", value)
    return a
}

/**
 * itemCheckableOn
 */
func (a *CRUD2Cards) ItemCheckableOn(value interface{}) *CRUD2Cards {
    a.Set("itemCheckableOn", value)
    return a
}

/**
 * disabled
 */
func (a *CRUD2Cards) Disabled(value interface{}) *CRUD2Cards {
    a.Set("disabled", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *CRUD2Cards) StaticLabelClassName(value interface{}) *CRUD2Cards {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * editorSetting
 */
func (a *CRUD2Cards) EditorSetting(value interface{}) *CRUD2Cards {
    a.Set("editorSetting", value)
    return a
}

/**
 * masonryLayout
 */
func (a *CRUD2Cards) MasonryLayout(value interface{}) *CRUD2Cards {
    a.Set("masonryLayout", value)
    return a
}

/**
 * api
 */
func (a *CRUD2Cards) Api(value interface{}) *CRUD2Cards {
    a.Set("api", value)
    return a
}

/**
 * footerToolbarClassName
 */
func (a *CRUD2Cards) FooterToolbarClassName(value interface{}) *CRUD2Cards {
    a.Set("footerToolbarClassName", value)
    return a
}

/**
 * parsePrimitiveQuery
 */
func (a *CRUD2Cards) ParsePrimitiveQuery(value interface{}) *CRUD2Cards {
    a.Set("parsePrimitiveQuery", value)
    return a
}

/**
 * id
 */
func (a *CRUD2Cards) Id(value interface{}) *CRUD2Cards {
    a.Set("id", value)
    return a
}

/**
 * selectable
 */
func (a *CRUD2Cards) Selectable(value interface{}) *CRUD2Cards {
    a.Set("selectable", value)
    return a
}

/**
 * hidden
 */
func (a *CRUD2Cards) Hidden(value interface{}) *CRUD2Cards {
    a.Set("hidden", value)
    return a
}

/**
 * staticOn
 */
func (a *CRUD2Cards) StaticOn(value interface{}) *CRUD2Cards {
    a.Set("staticOn", value)
    return a
}

/**
 * style
 */
func (a *CRUD2Cards) Style(value interface{}) *CRUD2Cards {
    a.Set("style", value)
    return a
}

/**
 * itemClassName
 */
func (a *CRUD2Cards) ItemClassName(value interface{}) *CRUD2Cards {
    a.Set("itemClassName", value)
    return a
}

/**
 * loadDataOnce
 */
func (a *CRUD2Cards) LoadDataOnce(value interface{}) *CRUD2Cards {
    a.Set("loadDataOnce", value)
    return a
}

/**
 * headerToolbar
 */
func (a *CRUD2Cards) HeaderToolbar(value interface{}) *CRUD2Cards {
    a.Set("headerToolbar", value)
    return a
}

/**
 * pullRefresh
 */
func (a *CRUD2Cards) PullRefresh(value interface{}) *CRUD2Cards {
    a.Set("pullRefresh", value)
    return a
}

/**
 */
func (a *CRUD2Cards) Type(value interface{}) *CRUD2Cards {
    a.Set("type", value)
    return a
}

/**
 * headerClassName
 */
func (a *CRUD2Cards) HeaderClassName(value interface{}) *CRUD2Cards {
    a.Set("headerClassName", value)
    return a
}

/**
 * footer
 */
func (a *CRUD2Cards) Footer(value interface{}) *CRUD2Cards {
    a.Set("footer", value)
    return a
}

/**
 * quickSaveItemApi
 */
func (a *CRUD2Cards) QuickSaveItemApi(value interface{}) *CRUD2Cards {
    a.Set("quickSaveItemApi", value)
    return a
}

/**
 * hiddenOn
 */
func (a *CRUD2Cards) HiddenOn(value interface{}) *CRUD2Cards {
    a.Set("hiddenOn", value)
    return a
}
