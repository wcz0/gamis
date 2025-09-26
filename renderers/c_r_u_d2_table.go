package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type CRUD2Table struct {
	*BaseRenderer
}

func NewCRUD2Table() *CRUD2Table {
    a := &CRUD2Table{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "crud2")
    return a
}


func (a *CRUD2Table) Set(name string, value interface{}) *CRUD2Table {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * source
 */
func (a *CRUD2Table) Source(value interface{}) *CRUD2Table {
    a.Set("source", value)
    return a
}

/**
 * keyField
 */
func (a *CRUD2Table) KeyField(value interface{}) *CRUD2Table {
    a.Set("keyField", value)
    return a
}

/**
 * maxKeepItemSelectionLength
 */
func (a *CRUD2Table) MaxKeepItemSelectionLength(value interface{}) *CRUD2Table {
    a.Set("maxKeepItemSelectionLength", value)
    return a
}

/**
 * loadingConfig
 */
func (a *CRUD2Table) LoadingConfig(value interface{}) *CRUD2Table {
    a.Set("loadingConfig", value)
    return a
}

/**
 * silentPolling
 */
func (a *CRUD2Table) SilentPolling(value interface{}) *CRUD2Table {
    a.Set("silentPolling", value)
    return a
}

/**
 * parsePrimitiveQuery
 */
func (a *CRUD2Table) ParsePrimitiveQuery(value interface{}) *CRUD2Table {
    a.Set("parsePrimitiveQuery", value)
    return a
}

/**
 * tableLayout
 */
func (a *CRUD2Table) TableLayout(value interface{}) *CRUD2Table {
    a.Set("tableLayout", value)
    return a
}

/**
 */
func (a *CRUD2Table) Mode(value interface{}) *CRUD2Table {
    a.Set("mode", value)
    return a
}

/**
 * interval
 */
func (a *CRUD2Table) Interval(value interface{}) *CRUD2Table {
    a.Set("interval", value)
    return a
}

/**
 * columns
 */
func (a *CRUD2Table) Columns(value interface{}) *CRUD2Table {
    a.Set("columns", value)
    return a
}

/**
 * expandable
 */
func (a *CRUD2Table) Expandable(value interface{}) *CRUD2Table {
    a.Set("expandable", value)
    return a
}

/**
 * quickSaveApi
 */
func (a *CRUD2Table) QuickSaveApi(value interface{}) *CRUD2Table {
    a.Set("quickSaveApi", value)
    return a
}

/**
 * syncLocation
 */
func (a *CRUD2Table) SyncLocation(value interface{}) *CRUD2Table {
    a.Set("syncLocation", value)
    return a
}

/**
 * footerToolbarClassName
 */
func (a *CRUD2Table) FooterToolbarClassName(value interface{}) *CRUD2Table {
    a.Set("footerToolbarClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *CRUD2Table) StaticSchema(value interface{}) *CRUD2Table {
    a.Set("staticSchema", value)
    return a
}

/**
 * style
 */
func (a *CRUD2Table) Style(value interface{}) *CRUD2Table {
    a.Set("style", value)
    return a
}

/**
 * columnsTogglable
 */
func (a *CRUD2Table) ColumnsTogglable(value interface{}) *CRUD2Table {
    a.Set("columnsTogglable", value)
    return a
}

/**
 * lineHeight
 */
func (a *CRUD2Table) LineHeight(value interface{}) *CRUD2Table {
    a.Set("lineHeight", value)
    return a
}

/**
 * messages
 */
func (a *CRUD2Table) Messages(value interface{}) *CRUD2Table {
    a.Set("messages", value)
    return a
}

/**
 * stopAutoRefreshWhen
 */
func (a *CRUD2Table) StopAutoRefreshWhen(value interface{}) *CRUD2Table {
    a.Set("stopAutoRefreshWhen", value)
    return a
}

/**
 * loadType
 */
func (a *CRUD2Table) LoadType(value interface{}) *CRUD2Table {
    a.Set("loadType", value)
    return a
}

/**
 * perPageField
 */
func (a *CRUD2Table) PerPageField(value interface{}) *CRUD2Table {
    a.Set("perPageField", value)
    return a
}

/**
 * hiddenOn
 */
func (a *CRUD2Table) HiddenOn(value interface{}) *CRUD2Table {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *CRUD2Table) StaticLabelClassName(value interface{}) *CRUD2Table {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * showHeader
 */
func (a *CRUD2Table) ShowHeader(value interface{}) *CRUD2Table {
    a.Set("showHeader", value)
    return a
}

/**
 * autoFillHeight
 */
func (a *CRUD2Table) AutoFillHeight(value interface{}) *CRUD2Table {
    a.Set("autoFillHeight", value)
    return a
}

/**
 * canAccessSuperData
 */
func (a *CRUD2Table) CanAccessSuperData(value interface{}) *CRUD2Table {
    a.Set("canAccessSuperData", value)
    return a
}

/**
 * rowSelection
 */
func (a *CRUD2Table) RowSelection(value interface{}) *CRUD2Table {
    a.Set("rowSelection", value)
    return a
}

/**
 * loading
 */
func (a *CRUD2Table) Loading(value interface{}) *CRUD2Table {
    a.Set("loading", value)
    return a
}

/**
 * quickSaveItemApi
 */
func (a *CRUD2Table) QuickSaveItemApi(value interface{}) *CRUD2Table {
    a.Set("quickSaveItemApi", value)
    return a
}

/**
 * disabled
 */
func (a *CRUD2Table) Disabled(value interface{}) *CRUD2Table {
    a.Set("disabled", value)
    return a
}

/**
 * hideQuickSaveBtn
 */
func (a *CRUD2Table) HideQuickSaveBtn(value interface{}) *CRUD2Table {
    a.Set("hideQuickSaveBtn", value)
    return a
}

/**
 * hidden
 */
func (a *CRUD2Table) Hidden(value interface{}) *CRUD2Table {
    a.Set("hidden", value)
    return a
}

/**
 * useMobileUI
 */
func (a *CRUD2Table) UseMobileUI(value interface{}) *CRUD2Table {
    a.Set("useMobileUI", value)
    return a
}

/**
 * selectable
 */
func (a *CRUD2Table) Selectable(value interface{}) *CRUD2Table {
    a.Set("selectable", value)
    return a
}

/**
 * primaryField
 */
func (a *CRUD2Table) PrimaryField(value interface{}) *CRUD2Table {
    a.Set("primaryField", value)
    return a
}

/**
 * lazyRenderAfter
 */
func (a *CRUD2Table) LazyRenderAfter(value interface{}) *CRUD2Table {
    a.Set("lazyRenderAfter", value)
    return a
}

/**
 * name
 */
func (a *CRUD2Table) Name(value interface{}) *CRUD2Table {
    a.Set("name", value)
    return a
}

/**
 * headerToolbar
 */
func (a *CRUD2Table) HeaderToolbar(value interface{}) *CRUD2Table {
    a.Set("headerToolbar", value)
    return a
}

/**
 * title
 */
func (a *CRUD2Table) Title(value interface{}) *CRUD2Table {
    a.Set("title", value)
    return a
}

/**
 * itemBadge
 */
func (a *CRUD2Table) ItemBadge(value interface{}) *CRUD2Table {
    a.Set("itemBadge", value)
    return a
}

/**
 * multiple
 */
func (a *CRUD2Table) Multiple(value interface{}) *CRUD2Table {
    a.Set("multiple", value)
    return a
}

/**
 * footerToolbar
 */
func (a *CRUD2Table) FooterToolbar(value interface{}) *CRUD2Table {
    a.Set("footerToolbar", value)
    return a
}

/**
 * pullRefresh
 */
func (a *CRUD2Table) PullRefresh(value interface{}) *CRUD2Table {
    a.Set("pullRefresh", value)
    return a
}

/**
 * id
 */
func (a *CRUD2Table) Id(value interface{}) *CRUD2Table {
    a.Set("id", value)
    return a
}

/**
 * onEvent
 */
func (a *CRUD2Table) OnEvent(value interface{}) *CRUD2Table {
    a.Set("onEvent", value)
    return a
}

/**
 * static
 */
func (a *CRUD2Table) Static(value interface{}) *CRUD2Table {
    a.Set("static", value)
    return a
}

/**
 * testid
 */
func (a *CRUD2Table) Testid(value interface{}) *CRUD2Table {
    a.Set("testid", value)
    return a
}

/**
 * keepItemSelectionOnPageChange
 */
func (a *CRUD2Table) KeepItemSelectionOnPageChange(value interface{}) *CRUD2Table {
    a.Set("keepItemSelectionOnPageChange", value)
    return a
}

/**
 * api
 */
func (a *CRUD2Table) Api(value interface{}) *CRUD2Table {
    a.Set("api", value)
    return a
}

/**
 * headerToolbarClassName
 */
func (a *CRUD2Table) HeaderToolbarClassName(value interface{}) *CRUD2Table {
    a.Set("headerToolbarClassName", value)
    return a
}

/**
 * syncResponse2Query
 */
func (a *CRUD2Table) SyncResponse2Query(value interface{}) *CRUD2Table {
    a.Set("syncResponse2Query", value)
    return a
}

/**
 * disabledOn
 */
func (a *CRUD2Table) DisabledOn(value interface{}) *CRUD2Table {
    a.Set("disabledOn", value)
    return a
}

/**
 * editorSetting
 */
func (a *CRUD2Table) EditorSetting(value interface{}) *CRUD2Table {
    a.Set("editorSetting", value)
    return a
}

/**
 * popOverContainer
 */
func (a *CRUD2Table) PopOverContainer(value interface{}) *CRUD2Table {
    a.Set("popOverContainer", value)
    return a
}

/**
 * childrenColumnName
 */
func (a *CRUD2Table) ChildrenColumnName(value interface{}) *CRUD2Table {
    a.Set("childrenColumnName", value)
    return a
}

/**
 * rowClassNameExpr
 */
func (a *CRUD2Table) RowClassNameExpr(value interface{}) *CRUD2Table {
    a.Set("rowClassNameExpr", value)
    return a
}

/**
 * bordered
 */
func (a *CRUD2Table) Bordered(value interface{}) *CRUD2Table {
    a.Set("bordered", value)
    return a
}

/**
 * actions
 */
func (a *CRUD2Table) Actions(value interface{}) *CRUD2Table {
    a.Set("actions", value)
    return a
}

/**
 * showSelection
 */
func (a *CRUD2Table) ShowSelection(value interface{}) *CRUD2Table {
    a.Set("showSelection", value)
    return a
}

/**
 * staticOn
 */
func (a *CRUD2Table) StaticOn(value interface{}) *CRUD2Table {
    a.Set("staticOn", value)
    return a
}

/**
 * sticky
 */
func (a *CRUD2Table) Sticky(value interface{}) *CRUD2Table {
    a.Set("sticky", value)
    return a
}

/**
 * footer
 */
func (a *CRUD2Table) Footer(value interface{}) *CRUD2Table {
    a.Set("footer", value)
    return a
}

/**
 * reload
 */
func (a *CRUD2Table) Reload(value interface{}) *CRUD2Table {
    a.Set("reload", value)
    return a
}

/**
 * loadDataOnce
 */
func (a *CRUD2Table) LoadDataOnce(value interface{}) *CRUD2Table {
    a.Set("loadDataOnce", value)
    return a
}

/**
 * pageField
 */
func (a *CRUD2Table) PageField(value interface{}) *CRUD2Table {
    a.Set("pageField", value)
    return a
}

/**
 * visibleOn
 */
func (a *CRUD2Table) VisibleOn(value interface{}) *CRUD2Table {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *CRUD2Table) StaticInputClassName(value interface{}) *CRUD2Table {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * showBadge
 */
func (a *CRUD2Table) ShowBadge(value interface{}) *CRUD2Table {
    a.Set("showBadge", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *CRUD2Table) StaticPlaceholder(value interface{}) *CRUD2Table {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticClassName
 */
func (a *CRUD2Table) StaticClassName(value interface{}) *CRUD2Table {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *CRUD2Table) Type(value interface{}) *CRUD2Table {
    a.Set("type", value)
    return a
}

/**
 * saveOrderApi
 */
func (a *CRUD2Table) SaveOrderApi(value interface{}) *CRUD2Table {
    a.Set("saveOrderApi", value)
    return a
}

/**
 * autoJumpToTopOnPagerChange
 */
func (a *CRUD2Table) AutoJumpToTopOnPagerChange(value interface{}) *CRUD2Table {
    a.Set("autoJumpToTopOnPagerChange", value)
    return a
}

/**
 * visible
 */
func (a *CRUD2Table) Visible(value interface{}) *CRUD2Table {
    a.Set("visible", value)
    return a
}

/**
 * perPage
 */
func (a *CRUD2Table) PerPage(value interface{}) *CRUD2Table {
    a.Set("perPage", value)
    return a
}

/**
 * className
 */
func (a *CRUD2Table) ClassName(value interface{}) *CRUD2Table {
    a.Set("className", value)
    return a
}
