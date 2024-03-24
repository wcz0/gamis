package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
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
/**
 * 是否显示头部
 */
func (a *CRUDTable) ShowHeader(value interface{}) *CRUDTable {
    a.Set("showHeader", value)
    return a
}

/**
 * 默认只有当分页数大于 1 是才显示，如果总是想显示请配置。
 */
func (a *CRUDTable) AlwaysShowPagination(value interface{}) *CRUDTable {
    a.Set("alwaysShowPagination", value)
    return a
}

/**
 * 是否显示底部
 */
func (a *CRUDTable) ShowFooter(value interface{}) *CRUDTable {
    a.Set("showFooter", value)
    return a
}

/**
 * 顶部总结行
 */
func (a *CRUDTable) PrefixRow(value interface{}) *CRUDTable {
    a.Set("prefixRow", value)
    return a
}

/**
 * table layout
 * 可选值: fixed | auto
 */
func (a *CRUDTable) TableLayout(value interface{}) *CRUDTable {
    a.Set("tableLayout", value)
    return a
}

/**
 * 是否禁用
 */
func (a *CRUDTable) Disabled(value interface{}) *CRUDTable {
    a.Set("disabled", value)
    return a
}

/**
 * 指定为 CRUD 渲染器。
 */
func (a *CRUDTable) Type(value interface{}) *CRUDTable {
    a.Set("type", value)
    return a
}

/**
 * 每页显示多少个空间成员的配置如： [10, 20, 50, 100]。
 */
func (a *CRUDTable) PerPageAvailable(value interface{}) *CRUDTable {
    a.Set("perPageAvailable", value)
    return a
}

/**
 * 静默拉取
 */
func (a *CRUDTable) SilentPolling(value interface{}) *CRUDTable {
    a.Set("silentPolling", value)
    return a
}

/**
 * 在开启loadDataOnce时，当修改过滤条件时是否重新请求api如果没有配置，当查询条件表单触发的会重新请求 api，当是列过滤或者是 search-box 触发的则不重新请求 api 如果配置为 true，则不管是什么触发都会重新请求 api 如果配置为 false 则不管是什么触发都不会重新请求 api
 */
func (a *CRUDTable) LoadDataOnceFetchOnFilter(value interface{}) *CRUDTable {
    a.Set("loadDataOnceFetchOnFilter", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *CRUDTable) StaticInputClassName(value interface{}) *CRUDTable {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *CRUDTable) LoadingConfig(value interface{}) *CRUDTable {
    a.Set("loadingConfig", value)
    return a
}

/**
 * 初始是否拉取
 */
func (a *CRUDTable) InitFetch(value interface{}) *CRUDTable {
    a.Set("initFetch", value)
    return a
}

/**
 * 底部工具栏
 */
func (a *CRUDTable) FooterToolbar(value interface{}) *CRUDTable {
    a.Set("footerToolbar", value)
    return a
}

/**
 * 也可以直接从环境变量中读取，但是不太推荐。
 */
func (a *CRUDTable) Source(value interface{}) *CRUDTable {
    a.Set("source", value)
    return a
}

/**
 * 工具栏 CSS 类名
 */
func (a *CRUDTable) ToolbarClassName(value interface{}) *CRUDTable {
    a.Set("toolbarClassName", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *CRUDTable) DisabledOn(value interface{}) *CRUDTable {
    a.Set("disabledOn", value)
    return a
}

/**
 * 单条操作
 */
func (a *CRUDTable) ItemActions(value interface{}) *CRUDTable {
    a.Set("itemActions", value)
    return a
}

/**
 * 是否可通过拖拽排序，通过表达式来配置
 */
func (a *CRUDTable) DraggableOn(value interface{}) *CRUDTable {
    a.Set("draggableOn", value)
    return a
}

/**
 * 保存排序的 api
 */
func (a *CRUDTable) SaveOrderApi(value interface{}) *CRUDTable {
    a.Set("saveOrderApi", value)
    return a
}

/**
 * 是否将过滤条件的参数同步到地址栏,默认为true
 */
func (a *CRUDTable) SyncLocation(value interface{}) *CRUDTable {
    a.Set("syncLocation", value)
    return a
}

/**
 * 是否固定表头
 */
func (a *CRUDTable) AffixHeader(value interface{}) *CRUDTable {
    a.Set("affixHeader", value)
    return a
}

/**
 * 是否固底
 */
func (a *CRUDTable) AffixFooter(value interface{}) *CRUDTable {
    a.Set("affixFooter", value)
    return a
}

/**
 * 是否开启底部展示功能，适合移动端展示
 */
func (a *CRUDTable) Footable(value interface{}) *CRUDTable {
    a.Set("footable", value)
    return a
}

/**
 * 占位符
 */
func (a *CRUDTable) Placeholder(value interface{}) *CRUDTable {
    a.Set("placeholder", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *CRUDTable) StaticLabelClassName(value interface{}) *CRUDTable {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 设置自动刷新时间
 */
func (a *CRUDTable) Interval(value interface{}) *CRUDTable {
    a.Set("interval", value)
    return a
}

/**
 * 内容区域占满屏幕剩余空间
 */
func (a *CRUDTable) AutoFillHeight(value interface{}) *CRUDTable {
    a.Set("autoFillHeight", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *CRUDTable) UseMobileUI(value interface{}) *CRUDTable {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *CRUDTable) StopAutoRefreshWhenModalIsOpen(value interface{}) *CRUDTable {
    a.Set("stopAutoRefreshWhenModalIsOpen", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *CRUDTable) Static(value interface{}) *CRUDTable {
    a.Set("static", value)
    return a
}

/**
 */
func (a *CRUDTable) StaticSchema(value interface{}) *CRUDTable {
    a.Set("staticSchema", value)
    return a
}

/**
 * 初始是否拉取，用表达式来配置。
 */
func (a *CRUDTable) InitFetchOn(value interface{}) *CRUDTable {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 快速编辑后用来批量保存的 API
 */
func (a *CRUDTable) QuickSaveApi(value interface{}) *CRUDTable {
    a.Set("quickSaveApi", value)
    return a
}

/**
 */
func (a *CRUDTable) Messages(value interface{}) *CRUDTable {
    a.Set("messages", value)
    return a
}

/**
 * 如果时内嵌模式，可以通过这个来配置默认的展开选项。
 */
func (a *CRUDTable) ExpandConfig(value interface{}) *CRUDTable {
    a.Set("expandConfig", value)
    return a
}

/**
 * 表格的列信息
 */
func (a *CRUDTable) Columns(value interface{}) *CRUDTable {
    a.Set("columns", value)
    return a
}

/**
 * 合并单元格配置，配置数字表示从左到右的多少列自动合并单元格。
 */
func (a *CRUDTable) CombineNum(value interface{}) *CRUDTable {
    a.Set("combineNum", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *CRUDTable) VisibleOn(value interface{}) *CRUDTable {
    a.Set("visibleOn", value)
    return a
}

/**
 * 设置分页一页显示的多少条数据的字段名。
 */
func (a *CRUDTable) PerPageField(value interface{}) *CRUDTable {
    a.Set("perPageField", value)
    return a
}

/**
 * 是否自动跳顶部，当切分页的时候。
 */
func (a *CRUDTable) AutoJumpToTopOnPagerChange(value interface{}) *CRUDTable {
    a.Set("autoJumpToTopOnPagerChange", value)
    return a
}

/**
 * 当配置 keepItemSelectionOnPageChange 时有用，用来配置已勾选项的文案。
 */
func (a *CRUDTable) LabelTpl(value interface{}) *CRUDTable {
    a.Set("labelTpl", value)
    return a
}

/**
 * 合并单元格配置，配置从第几列开始合并。
 */
func (a *CRUDTable) CombineFromIndex(value interface{}) *CRUDTable {
    a.Set("combineFromIndex", value)
    return a
}

/**
 * 是否可调整列宽
 */
func (a *CRUDTable) Resizable(value interface{}) *CRUDTable {
    a.Set("resizable", value)
    return a
}

/**
 * 是否显示
 */
func (a *CRUDTable) Visible(value interface{}) *CRUDTable {
    a.Set("visible", value)
    return a
}

/**
 */
func (a *CRUDTable) Name(value interface{}) *CRUDTable {
    a.Set("name", value)
    return a
}

/**
 */
func (a *CRUDTable) FilterTogglable(value interface{}) *CRUDTable {
    a.Set("filterTogglable", value)
    return a
}

/**
 * 展示列显示开关，自动即：列数量大于或等于5个时自动开启
 */
func (a *CRUDTable) ColumnsTogglable(value interface{}) *CRUDTable {
    a.Set("columnsTogglable", value)
    return a
}

/**
 * 顶部外层 CSS 类名
 */
func (a *CRUDTable) HeaderClassName(value interface{}) *CRUDTable {
    a.Set("headerClassName", value)
    return a
}

/**
 * 开启查询区域，会根据列元素的searchable属性值，自动生成查询条件表单
 */
func (a *CRUDTable) AutoGenerateFilter(value interface{}) *CRUDTable {
    a.Set("autoGenerateFilter", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *CRUDTable) Hidden(value interface{}) *CRUDTable {
    a.Set("hidden", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *CRUDTable) StaticPlaceholder(value interface{}) *CRUDTable {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *CRUDTable) StaticClassName(value interface{}) *CRUDTable {
    a.Set("staticClassName", value)
    return a
}

/**
 * 默认排序字段
 */
func (a *CRUDTable) OrderBy(value interface{}) *CRUDTable {
    a.Set("orderBy", value)
    return a
}

/**
 * 配置内部 DOM 的 className
 */
func (a *CRUDTable) InnerClassName(value interface{}) *CRUDTable {
    a.Set("innerClassName", value)
    return a
}

/**
 * 设置分页方向的字段名。单位简单分页时清楚时向前还是向后翻页。
 */
func (a *CRUDTable) PageDirectionField(value interface{}) *CRUDTable {
    a.Set("pageDirectionField", value)
    return a
}

/**
 * 懒加载 API，当行数据中用 defer: true 标记了，则其孩子节点将会用这个 API 来拉取数据。
 */
func (a *CRUDTable) DeferApi(value interface{}) *CRUDTable {
    a.Set("deferApi", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *CRUDTable) ClassName(value interface{}) *CRUDTable {
    a.Set("className", value)
    return a
}

/**
 * 每页个数，默认为 10，如果不是请设置。
 */
func (a *CRUDTable) PerPage(value interface{}) *CRUDTable {
    a.Set("perPage", value)
    return a
}

/**
 * 是否可通过拖拽排序
 */
func (a *CRUDTable) Draggable(value interface{}) *CRUDTable {
    a.Set("draggable", value)
    return a
}

/**
 */
func (a *CRUDTable) StopAutoRefreshWhen(value interface{}) *CRUDTable {
    a.Set("stopAutoRefreshWhen", value)
    return a
}

/**
 * 是否为前端单次加载模式，可以用来实现前端分页。
 */
func (a *CRUDTable) LoadDataOnce(value interface{}) *CRUDTable {
    a.Set("loadDataOnce", value)
    return a
}

/**
 * 设置数据
 */
func (a *CRUDTable) Data(value interface{}) *CRUDTable {
    a.Set("data", value)
    return a
}

/**
 * 底部总结行
 */
func (a *CRUDTable) AffixRow(value interface{}) *CRUDTable {
    a.Set("affixRow", value)
    return a
}

/**
 * 组件样式
 */
func (a *CRUDTable) Style(value interface{}) *CRUDTable {
    a.Set("style", value)
    return a
}

/**
 * 初始化数据 API
 */
func (a *CRUDTable) Api(value interface{}) *CRUDTable {
    a.Set("api", value)
    return a
}

/**
 * 默认排序方向
 * 可选值: asc | desc
 */
func (a *CRUDTable) OrderDir(value interface{}) *CRUDTable {
    a.Set("orderDir", value)
    return a
}

/**
 * 快速编辑配置成及时保存时使用的 API
 */
func (a *CRUDTable) QuickSaveItemApi(value interface{}) *CRUDTable {
    a.Set("quickSaveItemApi", value)
    return a
}

/**
 * 标题
 */
func (a *CRUDTable) Title(value interface{}) *CRUDTable {
    a.Set("title", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *CRUDTable) StaticOn(value interface{}) *CRUDTable {
    a.Set("staticOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *CRUDTable) EditorSetting(value interface{}) *CRUDTable {
    a.Set("editorSetting", value)
    return a
}

/**
 * 是否开启Query信息转换，开启后将会对url中的Query进行转换，默认开启，默认仅转化布尔值
 */
func (a *CRUDTable) ParsePrimitiveQuery(value interface{}) *CRUDTable {
    a.Set("parsePrimitiveQuery", value)
    return a
}

/**
 * 行样式表表达式
 */
func (a *CRUDTable) RowClassNameExpr(value interface{}) *CRUDTable {
    a.Set("rowClassNameExpr", value)
    return a
}

/**
 * 指定内容区的展示模式。
 */
func (a *CRUDTable) Mode(value interface{}) *CRUDTable {
    a.Set("mode", value)
    return a
}

/**
 * 可以默认给定初始参数如： {\"perPage\": 24}
 */
func (a *CRUDTable) DefaultParams(value interface{}) *CRUDTable {
    a.Set("defaultParams", value)
    return a
}

/**
 * 设置用来确定位置的字段名，设置后新的顺序将被赋值到该字段中。
 */
func (a *CRUDTable) OrderField(value interface{}) *CRUDTable {
    a.Set("orderField", value)
    return a
}

/**
 * 是否将接口返回的内容自动同步到地址栏，前提是开启了同步地址栏。
 */
func (a *CRUDTable) SyncResponse2Query(value interface{}) *CRUDTable {
    a.Set("syncResponse2Query", value)
    return a
}

/**
 * 自定义搜索匹配函数，当开启loadDataOnce时，会基于该函数计算的匹配结果进行过滤，主要用于处理列字段类型较为复杂或者字段值格式和后端返回不一致的场景
 */
func (a *CRUDTable) MatchFunc(value interface{}) *CRUDTable {
    a.Set("matchFunc", value)
    return a
}

/**
 * 底部外层 CSS 类名
 */
func (a *CRUDTable) FooterClassName(value interface{}) *CRUDTable {
    a.Set("footerClassName", value)
    return a
}

/**
 * 表格 CSS 类名
 */
func (a *CRUDTable) TableClassName(value interface{}) *CRUDTable {
    a.Set("tableClassName", value)
    return a
}

/**
 * 行角标
 */
func (a *CRUDTable) ItemBadge(value interface{}) *CRUDTable {
    a.Set("itemBadge", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *CRUDTable) HiddenOn(value interface{}) *CRUDTable {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *CRUDTable) Id(value interface{}) *CRUDTable {
    a.Set("id", value)
    return a
}

/**
 * 批量操作
 */
func (a *CRUDTable) BulkActions(value interface{}) *CRUDTable {
    a.Set("bulkActions", value)
    return a
}

/**
 */
func (a *CRUDTable) FilterDefaultVisible(value interface{}) *CRUDTable {
    a.Set("filterDefaultVisible", value)
    return a
}

/**
 * 表格是否可以获取父级数据域值，默认为false
 */
func (a *CRUDTable) CanAccessSuperData(value interface{}) *CRUDTable {
    a.Set("canAccessSuperData", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *CRUDTable) OnEvent(value interface{}) *CRUDTable {
    a.Set("onEvent", value)
    return a
}

/**
 * 过滤器表单
 */
func (a *CRUDTable) Filter(value interface{}) *CRUDTable {
    a.Set("filter", value)
    return a
}

/**
 * 设置分页页码字段名。
 */
func (a *CRUDTable) PageField(value interface{}) *CRUDTable {
    a.Set("pageField", value)
    return a
}

/**
 * 顶部工具栏
 */
func (a *CRUDTable) HeaderToolbar(value interface{}) *CRUDTable {
    a.Set("headerToolbar", value)
    return a
}

/**
 * 是否隐藏快速编辑的按钮。
 */
func (a *CRUDTable) HideQuickSaveBtn(value interface{}) *CRUDTable {
    a.Set("hideQuickSaveBtn", value)
    return a
}

/**
 * 分页的时候是否保留用户选择。
 */
func (a *CRUDTable) KeepItemSelectionOnPageChange(value interface{}) *CRUDTable {
    a.Set("keepItemSelectionOnPageChange", value)
    return a
}
