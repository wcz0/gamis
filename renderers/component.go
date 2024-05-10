package renderers

type Component struct {
	*BaseRenderer
}

func NewComponent(typeStr string) *Component {
	a := &Component{
		BaseRenderer: NewBaseRenderer(),
	}
	a.Set("type", typeStr)
	return a
}

/**
 * 【必填】这是 action 最核心的配置，来指定该 action 的作用类型，支持：ajax、link、url、drawer、dialog、confirm、cancel、prev、next、copy、close。
 */
 func (a *Component) ActionType(value interface{}) *Component {
	a.Set("actionType", value)
	return a
}

/**
 * 给按钮高亮添加类名。
 */
func (a *Component) ActiveClassName(value interface{}) *Component {
	a.Set("activeClassName", value)
	return a
}

/**
 * 按钮高亮时的样式，配置支持同level。
 */
func (a *Component) ActiveLevel(value interface{}) *Component {
	a.Set("activeLevel", value)
	return a
}

func (a *Component) Api(value interface{}) *Component {
	a.Set("api", value)
	return a
}

/**
 * 添加类名。
 */
func (a *Component) ClassName(value interface{}) *Component {
	a.Set("className", value)
	return a
}

/**
 * 当action配置在dialog或drawer的actions中时，配置为true指定此次操作完后关闭当前dialog或drawer。当值为字符串，并且是祖先层弹框的名字的时候，会把祖先弹框关闭掉。
 */
func (a *Component) Close(value interface{}) *Component {
	a.Set("close", value)
	return a
}

/**
 * 当设置后，操作在开始前会询问用户。可用 '$[xxx]' 取值。
 */
func (a *Component) ConfirmText(value interface{}) *Component {
	a.Set("confirmText", value)
	return a
}

/**
 * 被禁用后鼠标停留时弹出该段文字，也可以配置对象类型：字段为title和content。可用 '$[xxx]' 取值。
 */
func (a *Component) DisabledTip(value interface{}) *Component {
	a.Set("disabledTip", value)
	return a
}

/**
 * 设置图标，例如fa fa-plus。
 */
func (a *Component) Icon(value interface{}) *Component {
	a.Set("icon", value)
	return a
}

/**
 * 给图标上添加类名。
 */
func (a *Component) IconClassName(value interface{}) *Component {
	a.Set("iconClassName", value)
	return a
}

/**
 * 按钮文本。可用 '$[xxx]' 取值。
 */
func (a *Component) Label(value interface{}) *Component {
	a.Set("label", value)
	return a
}

/**
 * 按钮样式，支持：link /primary/secondary/info/success/warning/danger/light/dark/default。
 */
func (a *Component) Level(value interface{}) *Component {
	a.Set("level", value)
	return a
}

func (a *Component) Link(value interface{}) *Component {
	a.Set("link", value)
	return a
}

/**
 * 指定此次操作完后，需要刷新的目标组件名字（组件的name值，自己配置的），多个请用, 号隔开。
 */
func (a *Component) Reload(value interface{}) *Component {
	a.Set("reload", value)
	return a
}

/**
 * 配置字符串数组，指定在form中进行操作之前，需要指定的字段名的表单项通过验证
 */
func (a *Component) Required(value interface{}) *Component {
	a.Set("required", value)
	return a
}

/**
 * 在按钮文本右侧设置图标，例如fa fa-plus。
 */
func (a *Component) RightIcon(value interface{}) *Component {
	a.Set("rightIcon", value)
	return a
}

/**
 * 给右侧图标上添加类名。
 */
func (a *Component) RightIconClassName(value interface{}) *Component {
	a.Set("rightIconClassName", value)
	return a
}

/**
 * 按钮大小，支持：xs、sm、md、lg。
 */
func (a *Component) Size(value interface{}) *Component {
	a.Set("size", value)
	return a
}

/**
 * 鼠标停留时弹出该段文字，也可以配置对象类型：字段为title和content。可用 '$[xxx]' 取值。
 */
func (a *Component) Tooltip(value interface{}) *Component {
	a.Set("tooltip", value)
	return a
}

/**
 * 如果配置了tooltip或者disabledTip，指定提示信息位置，可配置top、bottom、left、right。
 */
func (a *Component) TooltipPlacement(value interface{}) *Component {
	a.Set("tooltipPlacement", value)
	return a
}

/**
 * 指定为 action 渲染器。
 */
func (a *Component) Type(value interface{}) *Component {
	a.Set("type", value)
	return a
}

/**
 * 是否将弹框中数据 merge 到父级作用域。
 */
func (a *Component) MergeData(value interface{}) *Component {
    a.Set("mergeData", value)
    return a
}

/**
 * 是否显示loading效果
 */
func (a *Component) LoadingOn(value interface{}) *Component {
    a.Set("loadingOn", value)
    return a
}

/**
 * 子内容
 */
func (a *Component) Body(value interface{}) *Component {
    a.Set("body", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Component) StaticLabelClassName(value interface{}) *Component {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 角标
 */
func (a *Component) Badge(value interface{}) *Component {
    a.Set("badge", value)
    return a
}

/**
 * 键盘快捷键
 */
func (a *Component) HotKey(value interface{}) *Component {
    a.Set("hotKey", value)
    return a
}

/**
 */
func (a *Component) StaticSchema(value interface{}) *Component {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *Component) Primary(value interface{}) *Component {
    a.Set("primary", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Component) StaticPlaceholder(value interface{}) *Component {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Component) UseMobileUI(value interface{}) *Component {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 可以指定让谁来触发这个动作。
 */
func (a *Component) Target(value interface{}) *Component {
    a.Set("target", value)
    return a
}

/**
 */
func (a *Component) DownloadFileName(value interface{}) *Component {
    a.Set("downloadFileName", value)
    return a
}

/**
 * 主要用于用户行为跟踪里区分是哪个按钮
 */
func (a *Component) Id(value interface{}) *Component {
    a.Set("id", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Component) OnEvent(value interface{}) *Component {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Component) StaticClassName(value interface{}) *Component {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Component) StaticInputClassName(value interface{}) *Component {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *Component) Style(value interface{}) *Component {
    a.Set("style", value)
    return a
}


/**
 * 是否禁用
 */
func (a *Component) Disabled(value interface{}) *Component {
    a.Set("disabled", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Component) VisibleOn(value interface{}) *Component {
    a.Set("visibleOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Component) EditorSetting(value interface{}) *Component {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *Component) Testid(value interface{}) *Component {
    a.Set("testid", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Component) DisabledOn(value interface{}) *Component {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Component) Static(value interface{}) *Component {
    a.Set("static", value)
    return a
}

/**
 * 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击，如果此属性配置成 false，则没有点选成员也能点击。
 */
func (a *Component) RequireSelected(value interface{}) *Component {
    a.Set("requireSelected", value)
    return a
}

/**
 * 自定义事件处理函数
 */
func (a *Component) OnClick(value interface{}) *Component {
    a.Set("onClick", value)
    return a
}

/**
 * loading 上的css 类名
 */
func (a *Component) LoadingClassName(value interface{}) *Component {
    a.Set("loadingClassName", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Component) HiddenOn(value interface{}) *Component {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Component) StaticOn(value interface{}) *Component {
    a.Set("staticOn", value)
    return a
}

/**
 * 倒计时文字自定义
 */
func (a *Component) CountDownTpl(value interface{}) *Component {
    a.Set("countDownTpl", value)
    return a
}

/**
 * 是否显示
 */
func (a *Component) Visible(value interface{}) *Component {
    a.Set("visible", value)
    return a
}

/**
 * 点击后的禁止倒计时（秒）
 */
func (a *Component) CountDown(value interface{}) *Component {
    a.Set("countDown", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Component) Hidden(value interface{}) *Component {
    a.Set("hidden", value)
    return a
}

/**
 * 是否为块状展示，默认为内联。
 */
func (a *Component) Block(value interface{}) *Component {
    a.Set("block", value)
    return a
}

/**
 */
 func (a *Component) Name(value interface{}) *Component {
    a.Set("name", value)
    return a
}

/**
 */
func (a *Component) Definitions(value interface{}) *Component {
    a.Set("definitions", value)
    return a
}

/**
 * 边栏区域
 */
func (a *Component) Aside(value interface{}) *Component {
    a.Set("aside", value)
    return a
}

/**
 * 边栏内容是否粘住，即不跟随滚动。
 */
func (a *Component) AsideSticky(value interface{}) *Component {
    a.Set("asideSticky", value)
    return a
}

/**
 * 页面顶部区域，当存在 title 时在右上角显示。
 */
func (a *Component) Toolbar(value interface{}) *Component {
    a.Set("toolbar", value)
    return a
}

/**
 * css 变量
 */
func (a *Component) CssVars(value interface{}) *Component {
    a.Set("cssVars", value)
    return a
}

/**
 * 下拉刷新配置
 */
func (a *Component) PullRefresh(value interface{}) *Component {
    a.Set("pullRefresh", value)
    return a
}

/**
 * 内容区 css 类名
 */
func (a *Component) BodyClassName(value interface{}) *Component {
    a.Set("bodyClassName", value)
    return a
}

/**
 * 边栏最小宽度
 */
func (a *Component) AsideMaxWidth(value interface{}) *Component {
    a.Set("asideMaxWidth", value)
    return a
}

/**
 * 页面描述, 标题旁边会出现个小图标，放上去会显示这个属性配置的内容。
 */
func (a *Component) Remark(value interface{}) *Component {
    a.Set("remark", value)
    return a
}

/**
 * 边栏是否允许拖动
 */
func (a *Component) AsideResizor(value interface{}) *Component {
    a.Set("asideResizor", value)
    return a
}

/**
 * 边栏最小宽度
 */
func (a *Component) AsideMinWidth(value interface{}) *Component {
    a.Set("asideMinWidth", value)
    return a
}

/**
 */
func (a *Component) Messages(value interface{}) *Component {
    a.Set("messages", value)
    return a
}

/**
 * 自定义页面级别样式表
 */
func (a *Component) Css(value interface{}) *Component {
    a.Set("css", value)
    return a
}

/**
 * 移动端下的样式表
 */
func (a *Component) MobileCSS(value interface{}) *Component {
    a.Set("mobileCSS", value)
    return a
}

/**
 * 配置 header 容器 className
 */
func (a *Component) HeaderClassName(value interface{}) *Component {
    a.Set("headerClassName", value)
    return a
}

/**
 * 是否默认就拉取？
 */
func (a *Component) InitFetch(value interface{}) *Component {
    a.Set("initFetch", value)
    return a
}

/**
 * 是否显示错误信息，默认是显示的。
 */
func (a *Component) ShowErrorMsg(value interface{}) *Component {
    a.Set("showErrorMsg", value)
    return a
}

/**
 */
func (a *Component) LoadingConfig(value interface{}) *Component {
    a.Set("loadingConfig", value)
    return a
}

/**
 * 页面级别的初始数据
 */
func (a *Component) Data(value interface{}) *Component {
    a.Set("data", value)
    return a
}

/**
 * 配置停止轮询的条件。
 */
func (a *Component) StopAutoRefreshWhen(value interface{}) *Component {
    a.Set("stopAutoRefreshWhen", value)
    return a
}

/**
 * 默认不设置自动感觉内容来决定要不要展示这些区域 如果配置了，以配置为主。
 */
func (a *Component) Regions(value interface{}) *Component {
    a.Set("regions", value)
    return a
}

/**
 * 页面副标题
 */
func (a *Component) SubTitle(value interface{}) *Component {
    a.Set("subTitle", value)
    return a
}

/**
 * 边栏区 css 类名
 */
func (a *Component) AsideClassName(value interface{}) *Component {
    a.Set("asideClassName", value)
    return a
}

/**
 * 配置轮询间隔，配置后 initApi 将轮询加载。
 */
func (a *Component) Interval(value interface{}) *Component {
    a.Set("interval", value)
    return a
}

/**
 * 页面标题
 */
func (a *Component) Title(value interface{}) *Component {
    a.Set("title", value)
    return a
}

/**
 * 是否默认就拉取表达式
 */
func (a *Component) InitFetchOn(value interface{}) *Component {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 配置 toolbar 容器 className
 */
func (a *Component) ToolbarClassName(value interface{}) *Component {
    a.Set("toolbarClassName", value)
    return a
}

/**
 * 是否要静默加载，也就是说不显示进度
 */
func (a *Component) SilentPolling(value interface{}) *Component {
    a.Set("silentPolling", value)
    return a
}

/**
 * 页面初始化的时候，可以设置一个 API 让其取拉取，发送数据会携带当前 data 数据（包含地址栏参数），获取得数据会合并到 data 中，供组件内使用。
 */
func (a *Component) InitApi(value interface{}) *Component {
    a.Set("initApi", value)
    return a
}

/**
 * 可以用来作为值的字段
 */
 func (a *Component) ValueField(value interface{}) *Component {
    a.Set("valueField", value)
    return a
}

/**
 * 默认排序字段
 */
func (a *Component) OrderBy(value interface{}) *Component {
    a.Set("orderBy", value)
    return a
}

/**
 * 是否将过滤条件的参数同步到地址栏,默认为true
 */
func (a *Component) SyncLocation(value interface{}) *Component {
    a.Set("syncLocation", value)
    return a
}

/**
 */
func (a *Component) Card(value interface{}) *Component {
    a.Set("card", value)
    return a
}

/**
 * 底部 CSS 类名
 */
func (a *Component) FooterClassName(value interface{}) *Component {
    a.Set("footerClassName", value)
    return a
}

/**
 * 顶部区域
 */
func (a *Component) Header(value interface{}) *Component {
    a.Set("header", value)
    return a
}

/**
 * 单条操作
 */
func (a *Component) ItemActions(value interface{}) *Component {
    a.Set("itemActions", value)
    return a
}

/**
 * 设置分页页码字段名。
 */
func (a *Component) PageField(value interface{}) *Component {
    a.Set("pageField", value)
    return a
}

/**
 * 是否固底
 */
func (a *Component) AffixFooter(value interface{}) *Component {
    a.Set("affixFooter", value)
    return a
}

/**
 * 每页显示多少个空间成员的配置如： [10, 20, 50, 100]。
 */
func (a *Component) PerPageAvailable(value interface{}) *Component {
    a.Set("perPageAvailable", value)
    return a
}

/**
 * 在开启loadDataOnce时，当修改过滤条件时是否重新请求api如果没有配置，当查询条件表单触发的会重新请求 api，当是列过滤或者是 search-box 触发的则不重新请求 api 如果配置为 true，则不管是什么触发都会重新请求 api 如果配置为 false 则不管是什么触发都不会重新请求 api
 */
func (a *Component) LoadDataOnceFetchOnFilter(value interface{}) *Component {
    a.Set("loadDataOnceFetchOnFilter", value)
    return a
}

/**
 * 如果时内嵌模式，可以通过这个来配置默认的展开选项。
 */
func (a *Component) ExpandConfig(value interface{}) *Component {
    a.Set("expandConfig", value)
    return a
}

/**
 * 设置分页一页显示的多少条数据的字段名。
 */
func (a *Component) PerPageField(value interface{}) *Component {
    a.Set("perPageField", value)
    return a
}

/**
 * 底部工具栏
 */
func (a *Component) FooterToolbar(value interface{}) *Component {
    a.Set("footerToolbar", value)
    return a
}

/**
 * 快速编辑配置成及时保存时使用的 API
 */
func (a *Component) QuickSaveItemApi(value interface{}) *Component {
    a.Set("quickSaveItemApi", value)
    return a
}

/**
 */
func (a *Component) StopAutoRefreshWhenModalIsOpen(value interface{}) *Component {
    a.Set("stopAutoRefreshWhenModalIsOpen", value)
    return a
}

/**
 * 是否为瀑布流布局？
 */
func (a *Component) MasonryLayout(value interface{}) *Component {
    a.Set("masonryLayout", value)
    return a
}

/**
 * 批量操作
 */
func (a *Component) BulkActions(value interface{}) *Component {
    a.Set("bulkActions", value)
    return a
}

/**
 * 每页个数，默认为 10，如果不是请设置。
 */
func (a *Component) PerPage(value interface{}) *Component {
    a.Set("perPage", value)
    return a
}

/**
 * 过滤器表单
 */
func (a *Component) Filter(value interface{}) *Component {
    a.Set("filter", value)
    return a
}

/**
 * 是否固顶
 */
func (a *Component) AffixHeader(value interface{}) *Component {
    a.Set("affixHeader", value)
    return a
}

/**
 * 分页的时候是否保留用户选择。
 */
func (a *Component) KeepItemSelectionOnPageChange(value interface{}) *Component {
    a.Set("keepItemSelectionOnPageChange", value)
    return a
}

/**
 * 是否开启Query信息转换，开启后将会对url中的Query进行转换，默认开启，默认仅转化布尔值
 */
func (a *Component) ParsePrimitiveQuery(value interface{}) *Component {
    a.Set("parsePrimitiveQuery", value)
    return a
}

/**
 * 自定义搜索匹配函数，当开启loadDataOnce时，会基于该函数计算的匹配结果进行过滤，主要用于处理列字段类型较为复杂或者字段值格式和后端返回不一致的场景
 */
func (a *Component) MatchFunc(value interface{}) *Component {
    a.Set("matchFunc", value)
    return a
}

/**
 * 指定内容区的展示模式。
 */
func (a *Component) Mode(value interface{}) *Component {
    a.Set("mode", value)
    return a
}

/**
 * 保存排序的 api
 */
func (a *Component) SaveOrderApi(value interface{}) *Component {
    a.Set("saveOrderApi", value)
    return a
}

/**
 * 也可以直接从环境变量中读取，但是不太推荐。
 */
func (a *Component) Source(value interface{}) *Component {
    a.Set("source", value)
    return a
}

/**
 * 配置内部 DOM 的 className
 */
func (a *Component) InnerClassName(value interface{}) *Component {
    a.Set("innerClassName", value)
    return a
}

/**
 * 设置分页方向的字段名。单位简单分页时清楚时向前还是向后翻页。
 */
func (a *Component) PageDirectionField(value interface{}) *Component {
    a.Set("pageDirectionField", value)
    return a
}

/**
 * 顶部工具栏
 */
func (a *Component) HeaderToolbar(value interface{}) *Component {
    a.Set("headerToolbar", value)
    return a
}

/**
 * 是否隐藏快速编辑的按钮。
 */
func (a *Component) HideQuickSaveBtn(value interface{}) *Component {
    a.Set("hideQuickSaveBtn", value)
    return a
}

/**
 * 卡片 CSS 类名
 */
func (a *Component) ItemClassName(value interface{}) *Component {
    a.Set("itemClassName", value)
    return a
}

/**
 * 是否为前端单次加载模式，可以用来实现前端分页。
 */
func (a *Component) LoadDataOnce(value interface{}) *Component {
    a.Set("loadDataOnce", value)
    return a
}

/**
 * 懒加载 API，当行数据中用 defer: true 标记了，则其孩子节点将会用这个 API 来拉取数据。
 */
func (a *Component) DeferApi(value interface{}) *Component {
    a.Set("deferApi", value)
    return a
}

/**
 * 是否将接口返回的内容自动同步到地址栏，前提是开启了同步地址栏。
 */
func (a *Component) SyncResponse2Query(value interface{}) *Component {
    a.Set("syncResponse2Query", value)
    return a
}

/**
 * 点击卡片的时候是否勾选卡片。
 */
func (a *Component) CheckOnItemClick(value interface{}) *Component {
    a.Set("checkOnItemClick", value)
    return a
}

/**
 * 底部区域
 */
func (a *Component) Footer(value interface{}) *Component {
    a.Set("footer", value)
    return a
}

/**
 * 配置某项是否可以点选
 */
func (a *Component) ItemCheckableOn(value interface{}) *Component {
    a.Set("itemCheckableOn", value)
    return a
}

/**
 * 是否可通过拖拽排序
 */
func (a *Component) Draggable(value interface{}) *Component {
    a.Set("draggable", value)
    return a
}

/**
 * 是否可通过拖拽排序，通过表达式来配置
 */
func (a *Component) DraggableOn(value interface{}) *Component {
    a.Set("draggableOn", value)
    return a
}

/**
 * 无数据提示
 */
func (a *Component) Placeholder(value interface{}) *Component {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否显示底部
 */
func (a *Component) ShowFooter(value interface{}) *Component {
    a.Set("showFooter", value)
    return a
}

/**
 * 是否显示头部
 */
func (a *Component) ShowHeader(value interface{}) *Component {
    a.Set("showHeader", value)
    return a
}

/**
 * 设置用来确定位置的字段名，设置后新的顺序将被赋值到该字段中。
 */
func (a *Component) OrderField(value interface{}) *Component {
    a.Set("orderField", value)
    return a
}

/**
 */
func (a *Component) FilterDefaultVisible(value interface{}) *Component {
    a.Set("filterDefaultVisible", value)
    return a
}

/**
 * 开启查询区域，会根据列元素的searchable属性值，自动生成查询条件表单
 */
func (a *Component) AutoGenerateFilter(value interface{}) *Component {
    a.Set("autoGenerateFilter", value)
    return a
}

/**
 * 是否隐藏勾选框
 */
func (a *Component) HideCheckToggler(value interface{}) *Component {
    a.Set("hideCheckToggler", value)
    return a
}

/**
 * 可以默认给定初始参数如： {\"perPage\": 24}
 */
func (a *Component) DefaultParams(value interface{}) *Component {
    a.Set("defaultParams", value)
    return a
}

/**
 * 快速编辑后用来批量保存的 API
 */
func (a *Component) QuickSaveApi(value interface{}) *Component {
    a.Set("quickSaveApi", value)
    return a
}

/**
 */
func (a *Component) FilterTogglable(value interface{}) *Component {
    a.Set("filterTogglable", value)
    return a
}

/**
 * 配置某项是否可拖拽排序，前提是要开启拖拽功能
 */
func (a *Component) ItemDraggableOn(value interface{}) *Component {
    a.Set("itemDraggableOn", value)
    return a
}

/**
 * 默认排序方向
 * 可选值: asc | desc
 */
func (a *Component) OrderDir(value interface{}) *Component {
    a.Set("orderDir", value)
    return a
}

/**
 * 是否自动跳顶部，当切分页的时候。
 */
func (a *Component) AutoJumpToTopOnPagerChange(value interface{}) *Component {
    a.Set("autoJumpToTopOnPagerChange", value)
    return a
}

/**
 * 当配置 keepItemSelectionOnPageChange 时有用，用来配置已勾选项的文案。
 */
func (a *Component) LabelTpl(value interface{}) *Component {
    a.Set("labelTpl", value)
    return a
}

/**
 * 默认只有当分页数大于 1 是才显示，如果总是想显示请配置。
 */
func (a *Component) AlwaysShowPagination(value interface{}) *Component {
    a.Set("alwaysShowPagination", value)
    return a
}

/**
 * 内容区域占满屏幕剩余空间
 */
func (a *Component) AutoFillHeight(value interface{}) *Component {
    a.Set("autoFillHeight", value)
    return a
}

