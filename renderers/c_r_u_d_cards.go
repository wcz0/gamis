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
    a.Set("type", "crud")
    a.Set("mode", "cards")
    return a
}

/**
 * 可以默认给定初始参数如： {\"perPage\": 24}
 */
func (a *CRUDCards) DefaultParams(value interface{}) *CRUDCards {
    a.Set("defaultParams", value)
    return a
}

/**
 * 保存排序的 api
 */
func (a *CRUDCards) SaveOrderApi(value interface{}) *CRUDCards {
    a.Set("saveOrderApi", value)
    return a
}

/**
 * 顶部工具栏
 */
func (a *CRUDCards) HeaderToolbar(value interface{}) *CRUDCards {
    a.Set("headerToolbar", value)
    return a
}

/**
 */
func (a *CRUDCards) Messages(value interface{}) *CRUDCards {
    a.Set("messages", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *CRUDCards) ClassName(value interface{}) *CRUDCards {
    a.Set("className", value)
    return a
}

/**
 * 指定为 CRUD 渲染器。
 */
func (a *CRUDCards) Type(value interface{}) *CRUDCards {
    a.Set("type", value)
    return a
}

/**
 * 是否为前端单次加载模式，可以用来实现前端分页。
 */
func (a *CRUDCards) LoadDataOnce(value interface{}) *CRUDCards {
    a.Set("loadDataOnce", value)
    return a
}

/**
 * 配置某项是否可以点选
 */
func (a *CRUDCards) ItemCheckableOn(value interface{}) *CRUDCards {
    a.Set("itemCheckableOn", value)
    return a
}

/**
 * 设置分页一页显示的多少条数据的字段名。
 */
func (a *CRUDCards) PerPageField(value interface{}) *CRUDCards {
    a.Set("perPageField", value)
    return a
}

/**
 * 默认只有当分页数大于 1 是才显示，如果总是想显示请配置。
 */
func (a *CRUDCards) AlwaysShowPagination(value interface{}) *CRUDCards {
    a.Set("alwaysShowPagination", value)
    return a
}

/**
 * 开启查询区域，会根据列元素的searchable属性值，自动生成查询条件表单
 */
func (a *CRUDCards) AutoGenerateFilter(value interface{}) *CRUDCards {
    a.Set("autoGenerateFilter", value)
    return a
}

/**
 * 可以用来作为值的字段
 */
func (a *CRUDCards) ValueField(value interface{}) *CRUDCards {
    a.Set("valueField", value)
    return a
}

/**
 * 每页显示多少个空间成员的配置如： [10, 20, 50, 100]。
 */
func (a *CRUDCards) PerPageAvailable(value interface{}) *CRUDCards {
    a.Set("perPageAvailable", value)
    return a
}

/**
 */
func (a *CRUDCards) TestIdBuilder(value interface{}) *CRUDCards {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 初始是否拉取
 */
func (a *CRUDCards) InitFetch(value interface{}) *CRUDCards {
    a.Set("initFetch", value)
    return a
}

/**
 * 是否自动跳顶部，当切分页的时候。
 */
func (a *CRUDCards) AutoJumpToTopOnPagerChange(value interface{}) *CRUDCards {
    a.Set("autoJumpToTopOnPagerChange", value)
    return a
}

/**
 */
func (a *CRUDCards) StopAutoRefreshWhen(value interface{}) *CRUDCards {
    a.Set("stopAutoRefreshWhen", value)
    return a
}

/**
 * 是否为瀑布流布局？
 */
func (a *CRUDCards) MasonryLayout(value interface{}) *CRUDCards {
    a.Set("masonryLayout", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *CRUDCards) StaticPlaceholder(value interface{}) *CRUDCards {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 分页的时候是否保留用户选择。
 */
func (a *CRUDCards) KeepItemSelectionOnPageChange(value interface{}) *CRUDCards {
    a.Set("keepItemSelectionOnPageChange", value)
    return a
}

/**
 * 如果时内嵌模式，可以通过这个来配置默认的展开选项。
 */
func (a *CRUDCards) ExpandConfig(value interface{}) *CRUDCards {
    a.Set("expandConfig", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *CRUDCards) StaticOn(value interface{}) *CRUDCards {
    a.Set("staticOn", value)
    return a
}

/**
 * 懒加载 API，当行数据中用 defer: true 标记了，则其孩子节点将会用这个 API 来拉取数据。
 */
func (a *CRUDCards) DeferApi(value interface{}) *CRUDCards {
    a.Set("deferApi", value)
    return a
}

/**
 * 是否固底
 */
func (a *CRUDCards) AffixFooter(value interface{}) *CRUDCards {
    a.Set("affixFooter", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *CRUDCards) HiddenOn(value interface{}) *CRUDCards {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示头部
 */
func (a *CRUDCards) ShowHeader(value interface{}) *CRUDCards {
    a.Set("showHeader", value)
    return a
}

/**
 * 也可以直接从环境变量中读取，但是不太推荐。
 */
func (a *CRUDCards) Source(value interface{}) *CRUDCards {
    a.Set("source", value)
    return a
}

/**
 * 初始化数据 API
 */
func (a *CRUDCards) Api(value interface{}) *CRUDCards {
    a.Set("api", value)
    return a
}

/**
 * 静默拉取
 */
func (a *CRUDCards) SilentPolling(value interface{}) *CRUDCards {
    a.Set("silentPolling", value)
    return a
}

/**
 */
func (a *CRUDCards) Card(value interface{}) *CRUDCards {
    a.Set("card", value)
    return a
}

/**
 * 头部 CSS 类名
 */
func (a *CRUDCards) HeaderClassName(value interface{}) *CRUDCards {
    a.Set("headerClassName", value)
    return a
}

/**
 * 是否可通过拖拽排序
 */
func (a *CRUDCards) Draggable(value interface{}) *CRUDCards {
    a.Set("draggable", value)
    return a
}

/**
 * 是否隐藏快速编辑的按钮。
 */
func (a *CRUDCards) HideQuickSaveBtn(value interface{}) *CRUDCards {
    a.Set("hideQuickSaveBtn", value)
    return a
}

/**
 * 当配置 keepItemSelectionOnPageChange 时有用，用来配置已勾选项的文案。
 */
func (a *CRUDCards) LabelTpl(value interface{}) *CRUDCards {
    a.Set("labelTpl", value)
    return a
}

/**
 * 自定义搜索匹配函数，当开启loadDataOnce时，会基于该函数计算的匹配结果进行过滤，主要用于处理列字段类型较为复杂或者字段值格式和后端返回不一致的场景
 */
func (a *CRUDCards) MatchFunc(value interface{}) *CRUDCards {
    a.Set("matchFunc", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *CRUDCards) Hidden(value interface{}) *CRUDCards {
    a.Set("hidden", value)
    return a
}

/**
 * 每页个数，默认为 10，如果不是请设置。
 */
func (a *CRUDCards) PerPage(value interface{}) *CRUDCards {
    a.Set("perPage", value)
    return a
}

/**
 * 初始是否拉取，用表达式来配置。
 */
func (a *CRUDCards) InitFetchOn(value interface{}) *CRUDCards {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 配置内部 DOM 的 className
 */
func (a *CRUDCards) InnerClassName(value interface{}) *CRUDCards {
    a.Set("innerClassName", value)
    return a
}

/**
 * 设置分页页码字段名。
 */
func (a *CRUDCards) PageField(value interface{}) *CRUDCards {
    a.Set("pageField", value)
    return a
}

/**
 */
func (a *CRUDCards) StopAutoRefreshWhenModalIsOpen(value interface{}) *CRUDCards {
    a.Set("stopAutoRefreshWhenModalIsOpen", value)
    return a
}

/**
 * 是否开启Query信息转换，开启后将会对url中的Query进行转换，默认开启，默认仅转化布尔值
 */
func (a *CRUDCards) ParsePrimitiveQuery(value interface{}) *CRUDCards {
    a.Set("parsePrimitiveQuery", value)
    return a
}

/**
 * 单条操作
 */
func (a *CRUDCards) ItemActions(value interface{}) *CRUDCards {
    a.Set("itemActions", value)
    return a
}

/**
 * 是否可通过拖拽排序，通过表达式来配置
 */
func (a *CRUDCards) DraggableOn(value interface{}) *CRUDCards {
    a.Set("draggableOn", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *CRUDCards) UseMobileUI(value interface{}) *CRUDCards {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 是否将过滤条件的参数同步到地址栏,默认为true
 */
func (a *CRUDCards) SyncLocation(value interface{}) *CRUDCards {
    a.Set("syncLocation", value)
    return a
}

/**
 * 是否固顶
 */
func (a *CRUDCards) AffixHeader(value interface{}) *CRUDCards {
    a.Set("affixHeader", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *CRUDCards) Static(value interface{}) *CRUDCards {
    a.Set("static", value)
    return a
}

/**
 * 点击卡片的时候是否勾选卡片。
 */
func (a *CRUDCards) CheckOnItemClick(value interface{}) *CRUDCards {
    a.Set("checkOnItemClick", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *CRUDCards) VisibleOn(value interface{}) *CRUDCards {
    a.Set("visibleOn", value)
    return a
}

/**
 */
func (a *CRUDCards) FilterTogglable(value interface{}) *CRUDCards {
    a.Set("filterTogglable", value)
    return a
}

/**
 * 底部 CSS 类名
 */
func (a *CRUDCards) FooterClassName(value interface{}) *CRUDCards {
    a.Set("footerClassName", value)
    return a
}

/**
 * 无数据提示
 */
func (a *CRUDCards) Placeholder(value interface{}) *CRUDCards {
    a.Set("placeholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *CRUDCards) StaticClassName(value interface{}) *CRUDCards {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *CRUDCards) StaticInputClassName(value interface{}) *CRUDCards {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 指定内容区的展示模式。
 */
func (a *CRUDCards) Mode(value interface{}) *CRUDCards {
    a.Set("mode", value)
    return a
}

/**
 * 默认排序方向
 * 可选值: asc | desc
 */
func (a *CRUDCards) OrderDir(value interface{}) *CRUDCards {
    a.Set("orderDir", value)
    return a
}

/**
 */
func (a *CRUDCards) Name(value interface{}) *CRUDCards {
    a.Set("name", value)
    return a
}

/**
 * 设置自动刷新时间
 */
func (a *CRUDCards) Interval(value interface{}) *CRUDCards {
    a.Set("interval", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *CRUDCards) Id(value interface{}) *CRUDCards {
    a.Set("id", value)
    return a
}

/**
 * 设置用来确定位置的字段名，设置后新的顺序将被赋值到该字段中。
 */
func (a *CRUDCards) OrderField(value interface{}) *CRUDCards {
    a.Set("orderField", value)
    return a
}

/**
 * 是否将接口返回的内容自动同步到地址栏，前提是开启了同步地址栏。
 */
func (a *CRUDCards) SyncResponse2Query(value interface{}) *CRUDCards {
    a.Set("syncResponse2Query", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *CRUDCards) OnEvent(value interface{}) *CRUDCards {
    a.Set("onEvent", value)
    return a
}

/**
 */
func (a *CRUDCards) StaticSchema(value interface{}) *CRUDCards {
    a.Set("staticSchema", value)
    return a
}

/**
 * 快速编辑后用来批量保存的 API
 */
func (a *CRUDCards) QuickSaveApi(value interface{}) *CRUDCards {
    a.Set("quickSaveApi", value)
    return a
}

/**
 * 内容区域占满屏幕剩余空间
 */
func (a *CRUDCards) AutoFillHeight(value interface{}) *CRUDCards {
    a.Set("autoFillHeight", value)
    return a
}

/**
 * 是否显示底部
 */
func (a *CRUDCards) ShowFooter(value interface{}) *CRUDCards {
    a.Set("showFooter", value)
    return a
}

/**
 * 顶部区域
 */
func (a *CRUDCards) Header(value interface{}) *CRUDCards {
    a.Set("header", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *CRUDCards) DisabledOn(value interface{}) *CRUDCards {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *CRUDCards) Visible(value interface{}) *CRUDCards {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *CRUDCards) StaticLabelClassName(value interface{}) *CRUDCards {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *CRUDCards) EditorSetting(value interface{}) *CRUDCards {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *CRUDCards) LoadingConfig(value interface{}) *CRUDCards {
    a.Set("loadingConfig", value)
    return a
}

/**
 * 默认排序字段
 */
func (a *CRUDCards) OrderBy(value interface{}) *CRUDCards {
    a.Set("orderBy", value)
    return a
}

/**
 * 卡片 CSS 类名
 */
func (a *CRUDCards) ItemClassName(value interface{}) *CRUDCards {
    a.Set("itemClassName", value)
    return a
}

/**
 * 是否禁用
 */
func (a *CRUDCards) Disabled(value interface{}) *CRUDCards {
    a.Set("disabled", value)
    return a
}

/**
 * 过滤器表单
 */
func (a *CRUDCards) Filter(value interface{}) *CRUDCards {
    a.Set("filter", value)
    return a
}

/**
 * 批量操作
 */
func (a *CRUDCards) BulkActions(value interface{}) *CRUDCards {
    a.Set("bulkActions", value)
    return a
}

/**
 * 快速编辑配置成及时保存时使用的 API
 */
func (a *CRUDCards) QuickSaveItemApi(value interface{}) *CRUDCards {
    a.Set("quickSaveItemApi", value)
    return a
}

/**
 */
func (a *CRUDCards) FilterDefaultVisible(value interface{}) *CRUDCards {
    a.Set("filterDefaultVisible", value)
    return a
}

/**
 * 是否隐藏勾选框
 */
func (a *CRUDCards) HideCheckToggler(value interface{}) *CRUDCards {
    a.Set("hideCheckToggler", value)
    return a
}

/**
 * 底部区域
 */
func (a *CRUDCards) Footer(value interface{}) *CRUDCards {
    a.Set("footer", value)
    return a
}

/**
 */
func (a *CRUDCards) Testid(value interface{}) *CRUDCards {
    a.Set("testid", value)
    return a
}

/**
 * 组件样式
 */
func (a *CRUDCards) Style(value interface{}) *CRUDCards {
    a.Set("style", value)
    return a
}

/**
 * 设置分页方向的字段名。单位简单分页时清楚时向前还是向后翻页。
 */
func (a *CRUDCards) PageDirectionField(value interface{}) *CRUDCards {
    a.Set("pageDirectionField", value)
    return a
}

/**
 * 底部工具栏
 */
func (a *CRUDCards) FooterToolbar(value interface{}) *CRUDCards {
    a.Set("footerToolbar", value)
    return a
}

/**
 * 在开启loadDataOnce时，当修改过滤条件时是否重新请求api如果没有配置，当查询条件表单触发的会重新请求 api，当是列过滤或者是 search-box 触发的则不重新请求 api 如果配置为 true，则不管是什么触发都会重新请求 api 如果配置为 false 则不管是什么触发都不会重新请求 api
 */
func (a *CRUDCards) LoadDataOnceFetchOnFilter(value interface{}) *CRUDCards {
    a.Set("loadDataOnceFetchOnFilter", value)
    return a
}

/**
 * 标题
 */
func (a *CRUDCards) Title(value interface{}) *CRUDCards {
    a.Set("title", value)
    return a
}

/**
 * 配置某项是否可拖拽排序，前提是要开启拖拽功能
 */
func (a *CRUDCards) ItemDraggableOn(value interface{}) *CRUDCards {
    a.Set("itemDraggableOn", value)
    return a
}
