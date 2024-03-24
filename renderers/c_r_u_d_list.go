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
/**
 * 配置某项是否可以点选
 */
func (a *CRUDList) ItemCheckableOn(value string) *CRUDList {
    a.Set("itemCheckableOn", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *CRUDList) Hidden(value string) *CRUDList {
    a.Set("hidden", value)
    return a
}

/**
 * 指定为 CRUD 渲染器。
 */
func (a *CRUDList) Type(value string) *CRUDList {
    a.Set("type", value)
    return a
}

/**
 * 可以默认给定初始参数如： {\"perPage\": 24}
 */
func (a *CRUDList) DefaultParams(value string) *CRUDList {
    a.Set("defaultParams", value)
    return a
}

/**
 * 初始是否拉取
 */
func (a *CRUDList) InitFetch(value string) *CRUDList {
    a.Set("initFetch", value)
    return a
}

/**
 * 是否固顶
 */
func (a *CRUDList) AffixHeader(value string) *CRUDList {
    a.Set("affixHeader", value)
    return a
}

/**
 * 初始化数据 API
 */
func (a *CRUDList) Api(value string) *CRUDList {
    a.Set("api", value)
    return a
}

/**
 * 懒加载 API，当行数据中用 defer: true 标记了，则其孩子节点将会用这个 API 来拉取数据。
 */
func (a *CRUDList) DeferApi(value string) *CRUDList {
    a.Set("deferApi", value)
    return a
}

/**
 */
func (a *CRUDList) Name(value string) *CRUDList {
    a.Set("name", value)
    return a
}

/**
 * 分页的时候是否保留用户选择。
 */
func (a *CRUDList) KeepItemSelectionOnPageChange(value string) *CRUDList {
    a.Set("keepItemSelectionOnPageChange", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *CRUDList) EditorSetting(value string) *CRUDList {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *CRUDList) Messages(value string) *CRUDList {
    a.Set("messages", value)
    return a
}

/**
 * 是否自动跳顶部，当切分页的时候。
 */
func (a *CRUDList) AutoJumpToTopOnPagerChange(value string) *CRUDList {
    a.Set("autoJumpToTopOnPagerChange", value)
    return a
}

/**
 * 是否隐藏勾选框
 */
func (a *CRUDList) HideCheckToggler(value string) *CRUDList {
    a.Set("hideCheckToggler", value)
    return a
}

/**
 * 是否固底
 */
func (a *CRUDList) AffixFooter(value string) *CRUDList {
    a.Set("affixFooter", value)
    return a
}

/**
 * 大小
 * 可选值: sm | base
 */
func (a *CRUDList) Size(value string) *CRUDList {
    a.Set("size", value)
    return a
}

/**
 * 是否禁用
 */
func (a *CRUDList) Disabled(value string) *CRUDList {
    a.Set("disabled", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *CRUDList) StaticInputClassName(value string) *CRUDList {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否将接口返回的内容自动同步到地址栏，前提是开启了同步地址栏。
 */
func (a *CRUDList) SyncResponse2Query(value string) *CRUDList {
    a.Set("syncResponse2Query", value)
    return a
}

/**
 * 如果时内嵌模式，可以通过这个来配置默认的展开选项。
 */
func (a *CRUDList) ExpandConfig(value string) *CRUDList {
    a.Set("expandConfig", value)
    return a
}

/**
 * 批量操作
 */
func (a *CRUDList) BulkActions(value string) *CRUDList {
    a.Set("bulkActions", value)
    return a
}

/**
 */
func (a *CRUDList) StopAutoRefreshWhen(value string) *CRUDList {
    a.Set("stopAutoRefreshWhen", value)
    return a
}

/**
 * 是否显示
 */
func (a *CRUDList) Visible(value string) *CRUDList {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *CRUDList) StaticClassName(value string) *CRUDList {
    a.Set("staticClassName", value)
    return a
}

/**
 * 在开启loadDataOnce时，当修改过滤条件时是否重新请求api如果没有配置，当查询条件表单触发的会重新请求 api，当是列过滤或者是 search-box 触发的则不重新请求 api 如果配置为 true，则不管是什么触发都会重新请求 api 如果配置为 false 则不管是什么触发都不会重新请求 api
 */
func (a *CRUDList) LoadDataOnceFetchOnFilter(value string) *CRUDList {
    a.Set("loadDataOnceFetchOnFilter", value)
    return a
}

/**
 * 顶部区域
 */
func (a *CRUDList) Header(value string) *CRUDList {
    a.Set("header", value)
    return a
}

/**
 * 单条数据展示内容配置
 */
func (a *CRUDList) ListItem(value string) *CRUDList {
    a.Set("listItem", value)
    return a
}

/**
 * 配置内部 DOM 的 className
 */
func (a *CRUDList) InnerClassName(value string) *CRUDList {
    a.Set("innerClassName", value)
    return a
}

/**
 */
func (a *CRUDList) FilterTogglable(value string) *CRUDList {
    a.Set("filterTogglable", value)
    return a
}

/**
 */
func (a *CRUDList) LoadingConfig(value string) *CRUDList {
    a.Set("loadingConfig", value)
    return a
}

/**
 * 每页个数，默认为 10，如果不是请设置。
 */
func (a *CRUDList) PerPage(value string) *CRUDList {
    a.Set("perPage", value)
    return a
}

/**
 * 默认排序方向
 * 可选值: asc | desc
 */
func (a *CRUDList) OrderDir(value string) *CRUDList {
    a.Set("orderDir", value)
    return a
}

/**
 * 开启查询区域，会根据列元素的searchable属性值，自动生成查询条件表单
 */
func (a *CRUDList) AutoGenerateFilter(value string) *CRUDList {
    a.Set("autoGenerateFilter", value)
    return a
}

/**
 * 是否开启Query信息转换，开启后将会对url中的Query进行转换，默认开启，默认仅转化布尔值
 */
func (a *CRUDList) ParsePrimitiveQuery(value string) *CRUDList {
    a.Set("parsePrimitiveQuery", value)
    return a
}

/**
 * 点击列表单行时，是否选择
 */
func (a *CRUDList) CheckOnItemClick(value string) *CRUDList {
    a.Set("checkOnItemClick", value)
    return a
}

/**
 * 点击列表项的行为
 */
func (a *CRUDList) ItemAction(value string) *CRUDList {
    a.Set("itemAction", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *CRUDList) HiddenOn(value string) *CRUDList {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *CRUDList) StaticPlaceholder(value string) *CRUDList {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 设置分页方向的字段名。单位简单分页时清楚时向前还是向后翻页。
 */
func (a *CRUDList) PageDirectionField(value string) *CRUDList {
    a.Set("pageDirectionField", value)
    return a
}

/**
 * 组件样式
 */
func (a *CRUDList) Style(value string) *CRUDList {
    a.Set("style", value)
    return a
}

/**
 * 是否可通过拖拽排序
 */
func (a *CRUDList) Draggable(value string) *CRUDList {
    a.Set("draggable", value)
    return a
}

/**
 * 是否隐藏快速编辑的按钮。
 */
func (a *CRUDList) HideQuickSaveBtn(value string) *CRUDList {
    a.Set("hideQuickSaveBtn", value)
    return a
}

/**
 * 标题
 */
func (a *CRUDList) Title(value string) *CRUDList {
    a.Set("title", value)
    return a
}

/**
 * 底部区域
 */
func (a *CRUDList) Footer(value string) *CRUDList {
    a.Set("footer", value)
    return a
}

/**
 * 也可以直接从环境变量中读取，但是不太推荐。
 */
func (a *CRUDList) Source(value string) *CRUDList {
    a.Set("source", value)
    return a
}

/**
 * 是否显示底部
 */
func (a *CRUDList) ShowFooter(value string) *CRUDList {
    a.Set("showFooter", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *CRUDList) VisibleOn(value string) *CRUDList {
    a.Set("visibleOn", value)
    return a
}

/**
 */
func (a *CRUDList) FilterDefaultVisible(value string) *CRUDList {
    a.Set("filterDefaultVisible", value)
    return a
}

/**
 * 顶部区域类名
 */
func (a *CRUDList) HeaderClassName(value string) *CRUDList {
    a.Set("headerClassName", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *CRUDList) OnEvent(value string) *CRUDList {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *CRUDList) Static(value string) *CRUDList {
    a.Set("static", value)
    return a
}

/**
 * 自定义搜索匹配函数，当开启loadDataOnce时，会基于该函数计算的匹配结果进行过滤，主要用于处理列字段类型较为复杂或者字段值格式和后端返回不一致的场景
 */
func (a *CRUDList) MatchFunc(value string) *CRUDList {
    a.Set("matchFunc", value)
    return a
}

/**
 * 是否显示头部
 */
func (a *CRUDList) ShowHeader(value string) *CRUDList {
    a.Set("showHeader", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *CRUDList) StaticLabelClassName(value string) *CRUDList {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *CRUDList) UseMobileUI(value string) *CRUDList {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 设置分页页码字段名。
 */
func (a *CRUDList) PageField(value string) *CRUDList {
    a.Set("pageField", value)
    return a
}

/**
 * 快速编辑配置成及时保存时使用的 API
 */
func (a *CRUDList) QuickSaveItemApi(value string) *CRUDList {
    a.Set("quickSaveItemApi", value)
    return a
}

/**
 * 快速编辑后用来批量保存的 API
 */
func (a *CRUDList) QuickSaveApi(value string) *CRUDList {
    a.Set("quickSaveApi", value)
    return a
}

/**
 * 顶部工具栏
 */
func (a *CRUDList) HeaderToolbar(value string) *CRUDList {
    a.Set("headerToolbar", value)
    return a
}

/**
 * 配置某项是否可拖拽排序，前提是要开启拖拽功能
 */
func (a *CRUDList) ItemDraggableOn(value string) *CRUDList {
    a.Set("itemDraggableOn", value)
    return a
}

/**
 * 指定内容区的展示模式。
 */
func (a *CRUDList) Mode(value string) *CRUDList {
    a.Set("mode", value)
    return a
}

/**
 * 默认排序字段
 */
func (a *CRUDList) OrderBy(value string) *CRUDList {
    a.Set("orderBy", value)
    return a
}

/**
 * 是否可通过拖拽排序，通过表达式来配置
 */
func (a *CRUDList) DraggableOn(value string) *CRUDList {
    a.Set("draggableOn", value)
    return a
}

/**
 * 设置用来确定位置的字段名，设置后新的顺序将被赋值到该字段中。
 */
func (a *CRUDList) OrderField(value string) *CRUDList {
    a.Set("orderField", value)
    return a
}

/**
 * 内容区域占满屏幕剩余空间
 */
func (a *CRUDList) AutoFillHeight(value string) *CRUDList {
    a.Set("autoFillHeight", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *CRUDList) StaticOn(value string) *CRUDList {
    a.Set("staticOn", value)
    return a
}

/**
 * 初始是否拉取，用表达式来配置。
 */
func (a *CRUDList) InitFetchOn(value string) *CRUDList {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 设置分页一页显示的多少条数据的字段名。
 */
func (a *CRUDList) PerPageField(value string) *CRUDList {
    a.Set("perPageField", value)
    return a
}

/**
 * 是否将过滤条件的参数同步到地址栏,默认为true
 */
func (a *CRUDList) SyncLocation(value string) *CRUDList {
    a.Set("syncLocation", value)
    return a
}

/**
 * 每页显示多少个空间成员的配置如： [10, 20, 50, 100]。
 */
func (a *CRUDList) PerPageAvailable(value string) *CRUDList {
    a.Set("perPageAvailable", value)
    return a
}

/**
 * 静默拉取
 */
func (a *CRUDList) SilentPolling(value string) *CRUDList {
    a.Set("silentPolling", value)
    return a
}

/**
 */
func (a *CRUDList) StopAutoRefreshWhenModalIsOpen(value string) *CRUDList {
    a.Set("stopAutoRefreshWhenModalIsOpen", value)
    return a
}

/**
 * 当配置 keepItemSelectionOnPageChange 时有用，用来配置已勾选项的文案。
 */
func (a *CRUDList) LabelTpl(value string) *CRUDList {
    a.Set("labelTpl", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *CRUDList) ClassName(value string) *CRUDList {
    a.Set("className", value)
    return a
}

/**
 */
func (a *CRUDList) StaticSchema(value string) *CRUDList {
    a.Set("staticSchema", value)
    return a
}

/**
 * 过滤器表单
 */
func (a *CRUDList) Filter(value string) *CRUDList {
    a.Set("filter", value)
    return a
}

/**
 * 设置自动刷新时间
 */
func (a *CRUDList) Interval(value string) *CRUDList {
    a.Set("interval", value)
    return a
}

/**
 * 无数据提示
 */
func (a *CRUDList) Placeholder(value string) *CRUDList {
    a.Set("placeholder", value)
    return a
}

/**
 * 可以用来作为值的字段
 */
func (a *CRUDList) ValueField(value string) *CRUDList {
    a.Set("valueField", value)
    return a
}

/**
 * 保存排序的 api
 */
func (a *CRUDList) SaveOrderApi(value string) *CRUDList {
    a.Set("saveOrderApi", value)
    return a
}

/**
 * 底部工具栏
 */
func (a *CRUDList) FooterToolbar(value string) *CRUDList {
    a.Set("footerToolbar", value)
    return a
}

/**
 * 默认只有当分页数大于 1 是才显示，如果总是想显示请配置。
 */
func (a *CRUDList) AlwaysShowPagination(value string) *CRUDList {
    a.Set("alwaysShowPagination", value)
    return a
}

/**
 * 底部区域类名
 */
func (a *CRUDList) FooterClassName(value string) *CRUDList {
    a.Set("footerClassName", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *CRUDList) DisabledOn(value string) *CRUDList {
    a.Set("disabledOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *CRUDList) Id(value string) *CRUDList {
    a.Set("id", value)
    return a
}

/**
 * 单条操作
 */
func (a *CRUDList) ItemActions(value string) *CRUDList {
    a.Set("itemActions", value)
    return a
}

/**
 * 是否为前端单次加载模式，可以用来实现前端分页。
 */
func (a *CRUDList) LoadDataOnce(value string) *CRUDList {
    a.Set("loadDataOnce", value)
    return a
}
