package renderers


/**
 * Table 表格渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/table

 * @author wcz0
 * @version 6.2.2
 */
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
/**
 * disabled
 */
func (a *BaseTable) Disabled(value interface{}) *BaseTable {
    a.Set("disabled", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *BaseTable) StaticPlaceholder(value interface{}) *BaseTable {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *BaseTable) StaticInputClassName(value interface{}) *BaseTable {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * style
 */
func (a *BaseTable) Style(value interface{}) *BaseTable {
    a.Set("style", value)
    return a
}

/**
 * editorSetting
 */
func (a *BaseTable) EditorSetting(value interface{}) *BaseTable {
    a.Set("editorSetting", value)
    return a
}

/**
 * title
 */
func (a *BaseTable) Title(value interface{}) *BaseTable {
    a.Set("title", value)
    return a
}

/**
 * autoFillHeight
 */
func (a *BaseTable) AutoFillHeight(value interface{}) *BaseTable {
    a.Set("autoFillHeight", value)
    return a
}

/**
 * staticOn
 */
func (a *BaseTable) StaticOn(value interface{}) *BaseTable {
    a.Set("staticOn", value)
    return a
}

/**
 * hiddenOn
 */
func (a *BaseTable) HiddenOn(value interface{}) *BaseTable {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visibleOn
 */
func (a *BaseTable) VisibleOn(value interface{}) *BaseTable {
    a.Set("visibleOn", value)
    return a
}

/**
 * onEvent
 */
func (a *BaseTable) OnEvent(value interface{}) *BaseTable {
    a.Set("onEvent", value)
    return a
}

/**
 * combineNum
 */
func (a *BaseTable) CombineNum(value interface{}) *BaseTable {
    a.Set("combineNum", value)
    return a
}

/**
 * prefixRow
 */
func (a *BaseTable) PrefixRow(value interface{}) *BaseTable {
    a.Set("prefixRow", value)
    return a
}

/**
 * affixRow
 */
func (a *BaseTable) AffixRow(value interface{}) *BaseTable {
    a.Set("affixRow", value)
    return a
}

/**
 * persistKey
 */
func (a *BaseTable) PersistKey(value interface{}) *BaseTable {
    a.Set("persistKey", value)
    return a
}

/**
 * className
 */
func (a *BaseTable) ClassName(value interface{}) *BaseTable {
    a.Set("className", value)
    return a
}

/**
 * hidden
 */
func (a *BaseTable) Hidden(value interface{}) *BaseTable {
    a.Set("hidden", value)
    return a
}

/**
 * showFooter
 */
func (a *BaseTable) ShowFooter(value interface{}) *BaseTable {
    a.Set("showFooter", value)
    return a
}

/**
 * static
 */
func (a *BaseTable) Static(value interface{}) *BaseTable {
    a.Set("static", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *BaseTable) StaticLabelClassName(value interface{}) *BaseTable {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * columnsTogglable
 */
func (a *BaseTable) ColumnsTogglable(value interface{}) *BaseTable {
    a.Set("columnsTogglable", value)
    return a
}

/**
 * placeholder
 */
func (a *BaseTable) Placeholder(value interface{}) *BaseTable {
    a.Set("placeholder", value)
    return a
}

/**
 * tableClassName
 */
func (a *BaseTable) TableClassName(value interface{}) *BaseTable {
    a.Set("tableClassName", value)
    return a
}

/**
 * resizable
 */
func (a *BaseTable) Resizable(value interface{}) *BaseTable {
    a.Set("resizable", value)
    return a
}

/**
 * tableLayout
 */
func (a *BaseTable) TableLayout(value interface{}) *BaseTable {
    a.Set("tableLayout", value)
    return a
}

/**
 * staticClassName
 */
func (a *BaseTable) StaticClassName(value interface{}) *BaseTable {
    a.Set("staticClassName", value)
    return a
}

/**
 * footerClassName
 */
func (a *BaseTable) FooterClassName(value interface{}) *BaseTable {
    a.Set("footerClassName", value)
    return a
}

/**
 * showIndex
 */
func (a *BaseTable) ShowIndex(value interface{}) *BaseTable {
    a.Set("showIndex", value)
    return a
}

/**
 * showHeader
 */
func (a *BaseTable) ShowHeader(value interface{}) *BaseTable {
    a.Set("showHeader", value)
    return a
}

/**
 * toolbarClassName
 */
func (a *BaseTable) ToolbarClassName(value interface{}) *BaseTable {
    a.Set("toolbarClassName", value)
    return a
}

/**
 * id
 */
func (a *BaseTable) Id(value interface{}) *BaseTable {
    a.Set("id", value)
    return a
}

/**
 * useMobileUI
 */
func (a *BaseTable) UseMobileUI(value interface{}) *BaseTable {
    a.Set("useMobileUI", value)
    return a
}

/**
 * type
 */
func (a *BaseTable) Type(value interface{}) *BaseTable {
    a.Set("type", value)
    return a
}

/**
 * columns
 */
func (a *BaseTable) Columns(value interface{}) *BaseTable {
    a.Set("columns", value)
    return a
}

/**
 * autoGenerateFilter
 */
func (a *BaseTable) AutoGenerateFilter(value interface{}) *BaseTable {
    a.Set("autoGenerateFilter", value)
    return a
}

/**
 * testid
 */
func (a *BaseTable) Testid(value interface{}) *BaseTable {
    a.Set("testid", value)
    return a
}

/**
 * affixHeader
 */
func (a *BaseTable) AffixHeader(value interface{}) *BaseTable {
    a.Set("affixHeader", value)
    return a
}

/**
 * footable
 */
func (a *BaseTable) Footable(value interface{}) *BaseTable {
    a.Set("footable", value)
    return a
}

/**
 * source
 */
func (a *BaseTable) Source(value interface{}) *BaseTable {
    a.Set("source", value)
    return a
}

/**
 * itemBadge
 */
func (a *BaseTable) ItemBadge(value interface{}) *BaseTable {
    a.Set("itemBadge", value)
    return a
}

/**
 * deferApi
 */
func (a *BaseTable) DeferApi(value interface{}) *BaseTable {
    a.Set("deferApi", value)
    return a
}

/**
 * disabledOn
 */
func (a *BaseTable) DisabledOn(value interface{}) *BaseTable {
    a.Set("disabledOn", value)
    return a
}

/**
 * visible
 */
func (a *BaseTable) Visible(value interface{}) *BaseTable {
    a.Set("visible", value)
    return a
}

/**
 * staticSchema
 */
func (a *BaseTable) StaticSchema(value interface{}) *BaseTable {
    a.Set("staticSchema", value)
    return a
}

/**
 * affixFooter
 */
func (a *BaseTable) AffixFooter(value interface{}) *BaseTable {
    a.Set("affixFooter", value)
    return a
}

/**
 * headerClassName
 */
func (a *BaseTable) HeaderClassName(value interface{}) *BaseTable {
    a.Set("headerClassName", value)
    return a
}

/**
 * combineFromIndex
 */
func (a *BaseTable) CombineFromIndex(value interface{}) *BaseTable {
    a.Set("combineFromIndex", value)
    return a
}

/**
 * rowClassNameExpr
 */
func (a *BaseTable) RowClassNameExpr(value interface{}) *BaseTable {
    a.Set("rowClassNameExpr", value)
    return a
}

/**
 * canAccessSuperData
 */
func (a *BaseTable) CanAccessSuperData(value interface{}) *BaseTable {
    a.Set("canAccessSuperData", value)
    return a
}
