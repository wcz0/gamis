package renderers


/**

*/
type CRUD2List struct {
	*BaseRenderer
}

func NewCRUD2List() *CRUD2List {
    a := &CRUD2List{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("mode", "list")
    a.Set("type", "crud2")
    return a
}
/**
 * 顶部区域
 */
func (a *CRUD2List) Header(value string) *CRUD2List {
    a.Set("header", value)
    return a
}

/**
 * 是否固底
 */
func (a *CRUD2List) AffixFooter(value string) *CRUD2List {
    a.Set("affixFooter", value)
    return a
}

/**
 */
func (a *CRUD2List) StopAutoRefreshWhen(value string) *CRUD2List {
    a.Set("stopAutoRefreshWhen", value)
    return a
}

/**
 * 快速编辑后用来批量保存的 API
 */
func (a *CRUD2List) QuickSaveApi(value string) *CRUD2List {
    a.Set("quickSaveApi", value)
    return a
}

/**
 * 是否隐藏快速编辑的按钮。
 */
func (a *CRUD2List) HideQuickSaveBtn(value string) *CRUD2List {
    a.Set("hideQuickSaveBtn", value)
    return a
}

/**
 * 点击列表项的行为
 */
func (a *CRUD2List) ItemAction(value string) *CRUD2List {
    a.Set("itemAction", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *CRUD2List) ClassName(value string) *CRUD2List {
    a.Set("className", value)
    return a
}

/**
 */
func (a *CRUD2List) Name(value string) *CRUD2List {
    a.Set("name", value)
    return a
}

/**
 * 是否将接口返回的内容自动同步到地址栏，前提是开启了同步地址栏。
 */
func (a *CRUD2List) SyncResponse2Query(value string) *CRUD2List {
    a.Set("syncResponse2Query", value)
    return a
}

/**
 * 也可以直接从环境变量中读取，但是不太推荐。
 */
func (a *CRUD2List) Source(value string) *CRUD2List {
    a.Set("source", value)
    return a
}

/**
 * 可以用来作为值的字段
 */
func (a *CRUD2List) ValueField(value string) *CRUD2List {
    a.Set("valueField", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *CRUD2List) StaticClassName(value string) *CRUD2List {
    a.Set("staticClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *CRUD2List) Style(value string) *CRUD2List {
    a.Set("style", value)
    return a
}

/**
 */
func (a *CRUD2List) LoadingConfig(value string) *CRUD2List {
    a.Set("loadingConfig", value)
    return a
}

/**
 * 设置自动刷新时间
 */
func (a *CRUD2List) Interval(value string) *CRUD2List {
    a.Set("interval", value)
    return a
}

/**
 * 是否可以选择数据，外部事件动作
 */
func (a *CRUD2List) Selectable(value string) *CRUD2List {
    a.Set("selectable", value)
    return a
}

/**
 * 底部区域
 */
func (a *CRUD2List) Footer(value string) *CRUD2List {
    a.Set("footer", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *CRUD2List) StaticOn(value string) *CRUD2List {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *CRUD2List) StaticInputClassName(value string) *CRUD2List {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否固顶
 */
func (a *CRUD2List) AffixHeader(value string) *CRUD2List {
    a.Set("affixHeader", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *CRUD2List) VisibleOn(value string) *CRUD2List {
    a.Set("visibleOn", value)
    return a
}

/**
 * 指定为 CRUD2 渲染器。
 */
func (a *CRUD2List) Type(value string) *CRUD2List {
    a.Set("type", value)
    return a
}

/**
 * 快速编辑配置成及时保存时使用的 API
 */
func (a *CRUD2List) QuickSaveItemApi(value string) *CRUD2List {
    a.Set("quickSaveItemApi", value)
    return a
}

/**
 * 点击列表单行时，是否选择
 */
func (a *CRUD2List) CheckOnItemClick(value string) *CRUD2List {
    a.Set("checkOnItemClick", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *CRUD2List) Id(value string) *CRUD2List {
    a.Set("id", value)
    return a
}

/**
 * 设置分页一页显示的多少条数据的字段名。
 */
func (a *CRUD2List) PerPageField(value string) *CRUD2List {
    a.Set("perPageField", value)
    return a
}

/**
 * 内容区域占满屏幕剩余空间
 */
func (a *CRUD2List) AutoFillHeight(value string) *CRUD2List {
    a.Set("autoFillHeight", value)
    return a
}

/**
 * 是否开启Query信息转换，开启后将会对url中的Query进行转换，默认开启，默认仅转化布尔值
 */
func (a *CRUD2List) ParsePrimitiveQuery(value string) *CRUD2List {
    a.Set("parsePrimitiveQuery", value)
    return a
}

/**
 * 标题
 */
func (a *CRUD2List) Title(value string) *CRUD2List {
    a.Set("title", value)
    return a
}

/**
 * 静默拉取
 */
func (a *CRUD2List) SilentPolling(value string) *CRUD2List {
    a.Set("silentPolling", value)
    return a
}

/**
 * 翻页时是否保留用户已选的数据
 */
func (a *CRUD2List) KeepItemSelectionOnPageChange(value string) *CRUD2List {
    a.Set("keepItemSelectionOnPageChange", value)
    return a
}

/**
 * 行标识符，默认为id
 */
func (a *CRUD2List) PrimaryField(value string) *CRUD2List {
    a.Set("primaryField", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *CRUD2List) Static(value string) *CRUD2List {
    a.Set("static", value)
    return a
}

/**
 * 是否将过滤条件的参数同步到地址栏,默认为true
 */
func (a *CRUD2List) SyncLocation(value string) *CRUD2List {
    a.Set("syncLocation", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *CRUD2List) HiddenOn(value string) *CRUD2List {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否展示已选数据区域，仅当selectable为 true 时生效
 */
func (a *CRUD2List) ShowSelection(value string) *CRUD2List {
    a.Set("showSelection", value)
    return a
}

/**
 * 是否禁用
 */
func (a *CRUD2List) Disabled(value string) *CRUD2List {
    a.Set("disabled", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *CRUD2List) EditorSetting(value string) *CRUD2List {
    a.Set("editorSetting", value)
    return a
}

/**
 * 底部区域CSS类名
 */
func (a *CRUD2List) FooterToolbarClassName(value string) *CRUD2List {
    a.Set("footerToolbarClassName", value)
    return a
}

/**
 * 底部区域类名
 */
func (a *CRUD2List) FooterClassName(value string) *CRUD2List {
    a.Set("footerClassName", value)
    return a
}

/**
 * 无数据提示
 */
func (a *CRUD2List) Placeholder(value string) *CRUD2List {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *CRUD2List) DisabledOn(value string) *CRUD2List {
    a.Set("disabledOn", value)
    return a
}

/**
 * 指定内容区的展示模式。
 */
func (a *CRUD2List) Mode(value string) *CRUD2List {
    a.Set("mode", value)
    return a
}

/**
 * 保存排序的 api
 */
func (a *CRUD2List) SaveOrderApi(value string) *CRUD2List {
    a.Set("saveOrderApi", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *CRUD2List) UseMobileUI(value string) *CRUD2List {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 设置分页页码字段名。
 */
func (a *CRUD2List) PageField(value string) *CRUD2List {
    a.Set("pageField", value)
    return a
}

/**
 * 底部区域
 */
func (a *CRUD2List) FooterToolbar(value string) *CRUD2List {
    a.Set("footerToolbar", value)
    return a
}

/**
 * 是否隐藏勾选框
 */
func (a *CRUD2List) HideCheckToggler(value string) *CRUD2List {
    a.Set("hideCheckToggler", value)
    return a
}

/**
 * 配置某项是否可以点选
 */
func (a *CRUD2List) ItemCheckableOn(value string) *CRUD2List {
    a.Set("itemCheckableOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *CRUD2List) OnEvent(value string) *CRUD2List {
    a.Set("onEvent", value)
    return a
}

/**
 */
func (a *CRUD2List) StaticSchema(value string) *CRUD2List {
    a.Set("staticSchema", value)
    return a
}

/**
 * 数据展示模式 无限加载 or 分页
 * 可选值: more | pagination
 */
func (a *CRUD2List) LoadType(value string) *CRUD2List {
    a.Set("loadType", value)
    return a
}

/**
 * 是否可以多选数据，仅当selectable为 true 时生效
 */
func (a *CRUD2List) Multiple(value string) *CRUD2List {
    a.Set("multiple", value)
    return a
}

/**
 * 是否自动跳顶部，当切分页的时候。
 */
func (a *CRUD2List) AutoJumpToTopOnPagerChange(value string) *CRUD2List {
    a.Set("autoJumpToTopOnPagerChange", value)
    return a
}

/**
 * 单条数据展示内容配置
 */
func (a *CRUD2List) ListItem(value string) *CRUD2List {
    a.Set("listItem", value)
    return a
}

/**
 * 大小
 * 可选值: sm | base
 */
func (a *CRUD2List) Size(value string) *CRUD2List {
    a.Set("size", value)
    return a
}

/**
 * 是否显示
 */
func (a *CRUD2List) Visible(value string) *CRUD2List {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *CRUD2List) StaticPlaceholder(value string) *CRUD2List {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 初始化数据 API
 */
func (a *CRUD2List) Api(value string) *CRUD2List {
    a.Set("api", value)
    return a
}

/**
 * 是否为前端单次加载模式，可以用来实现前端分页。
 */
func (a *CRUD2List) LoadDataOnce(value string) *CRUD2List {
    a.Set("loadDataOnce", value)
    return a
}

/**
 * 顶部区域类名
 */
func (a *CRUD2List) HeaderClassName(value string) *CRUD2List {
    a.Set("headerClassName", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *CRUD2List) StaticLabelClassName(value string) *CRUD2List {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 顶部区域CSS类名
 */
func (a *CRUD2List) HeaderToolbarClassName(value string) *CRUD2List {
    a.Set("headerToolbarClassName", value)
    return a
}

/**
 * 是否显示底部
 */
func (a *CRUD2List) ShowFooter(value string) *CRUD2List {
    a.Set("showFooter", value)
    return a
}

/**
 * 是否显示头部
 */
func (a *CRUD2List) ShowHeader(value string) *CRUD2List {
    a.Set("showHeader", value)
    return a
}

/**
 * 配置某项是否可拖拽排序，前提是要开启拖拽功能
 */
func (a *CRUD2List) ItemDraggableOn(value string) *CRUD2List {
    a.Set("itemDraggableOn", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *CRUD2List) Hidden(value string) *CRUD2List {
    a.Set("hidden", value)
    return a
}

/**
 * 无限加载时，根据此项设置其每页加载数量，可以不限制
 */
func (a *CRUD2List) PerPage(value string) *CRUD2List {
    a.Set("perPage", value)
    return a
}

/**
 * 顶部区域
 */
func (a *CRUD2List) HeaderToolbar(value string) *CRUD2List {
    a.Set("headerToolbar", value)
    return a
}
