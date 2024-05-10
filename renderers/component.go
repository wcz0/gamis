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

/**
 * 配置快速编辑功能
 */
 func (a *Component) QuickEdit(value string) *Component {
    a.Set("quickEdit", value)
    return a
}


/**
 * 指定行合并表达式
 */
func (a *Component) RowSpanExpr(value string) *Component {
    a.Set("rowSpanExpr", value)
    return a
}

/**
 * 指定列合并表达式
 */
func (a *Component) ColSpanExpr(value string) *Component {
    a.Set("colSpanExpr", value)
    return a
}

/**
 * 快速搜索
 */
func (a *Component) Searchable(value string) *Component {
    a.Set("searchable", value)
    return a
}

/**
 * 兼容table列筛选
 */
func (a *Component) Filterable(value string) *Component {
    a.Set("filterable", value)
    return a
}

/**
 * 内容居左、居中、居右
 */
func (a *Component) Align(value string) *Component {
    a.Set("align", value)
    return a
}

/**
 * 表头分组
 */
func (a *Component) Children(value string) *Component {
    a.Set("children", value)
    return a
}

/**
 * 兼容table快速排序
 */
func (a *Component) Sortable(value string) *Component {
    a.Set("sortable", value)
    return a
}

/**
 * 是否固定在左侧/右侧
 */
func (a *Component) Fixed(value string) *Component {
    a.Set("fixed", value)
    return a
}

/**
 * 单元格样式
 */
func (a *Component) ClassNameExpr(value string) *Component {
    a.Set("classNameExpr", value)
    return a
}

/**
 */
func (a *Component) Width(value string) *Component {
    a.Set("width", value)
    return a
}



/**
 * 可复制
 */
func (a *Component) Copyable(value string) *Component {
    a.Set("copyable", value)
    return a
}

/**
 * 快速排序
 */
func (a *Component) Sorter(value string) *Component {
    a.Set("sorter", value)
    return a
}

/**
 * 当前列是否展示
 */
func (a *Component) Toggled(value string) *Component {
    a.Set("toggled", value)
    return a
}

/**
 * 表头单元格样式
 */
func (a *Component) TitleClassName(value string) *Component {
    a.Set("titleClassName", value)
    return a
}

/**
 * 表格列单元格是否可以获取父级数据域值，默认为true，该配置对当前列内单元格生效
 */
func (a *Component) CanAccessSuperData(value string) *Component {
    a.Set("canAccessSuperData", value)
    return a
}

/**
 */
 func (a *Component) TestIdBuilder(value interface{}) *Component {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 列集合
 */
func (a *Component) Columns(value interface{}) *Component {
    a.Set("columns", value)
    return a
}

/**
 * 垂直对齐方式
 * 可选值: top | middle | bottom | between
 */
func (a *Component) Valign(value interface{}) *Component {
    a.Set("valign", value)
    return a
}

/**
 * 水平间距
 * 可选值: xs | sm | base | none | md | lg
 */
func (a *Component) Gap(value interface{}) *Component {
    a.Set("gap", value)
    return a
}

/**
 * 展示列显示开关，自动即：列数量大于或等于5个时自动开启
 */
func (a *Component) ColumnsTogglable(value interface{}) *Component {
    a.Set("columnsTogglable", value)
    return a
}

/**
 * 表格 CSS 类名
 */
func (a *Component) TableClassName(value interface{}) *Component {
    a.Set("tableClassName", value)
    return a
}

/**
 * 是否开启底部展示功能，适合移动端展示
 */
func (a *Component) Footable(value interface{}) *Component {
    a.Set("footable", value)
    return a
}

/**
 * 顶部总结行
 */
func (a *Component) PrefixRow(value interface{}) *Component {
    a.Set("prefixRow", value)
    return a
}

/**
 * 合并单元格配置，配置数字表示从左到右的多少列自动合并单元格。
 */
func (a *Component) CombineNum(value interface{}) *Component {
    a.Set("combineNum", value)
    return a
}

/**
 * 行角标
 */
func (a *Component) ItemBadge(value interface{}) *Component {
    a.Set("itemBadge", value)
    return a
}

/**
 * table layout
 * 可选值: fixed | auto
 */
func (a *Component) TableLayout(value interface{}) *Component {
    a.Set("tableLayout", value)
    return a
}

/**
 * 底部总结行
 */
func (a *Component) AffixRow(value interface{}) *Component {
    a.Set("affixRow", value)
    return a
}

/**
 * 是否可调整列宽
 */
func (a *Component) Resizable(value interface{}) *Component {
    a.Set("resizable", value)
    return a
}

/**
 * 合并单元格配置，配置从第几列开始合并。
 */
func (a *Component) CombineFromIndex(value interface{}) *Component {
    a.Set("combineFromIndex", value)
    return a
}

/**
 * 行样式表表达式
 */
func (a *Component) RowClassNameExpr(value interface{}) *Component {
    a.Set("rowClassNameExpr", value)
    return a
}

/**
 * 修改的时候是否直接提交表单。
 */
func (a *Component) SubmitOnChange(value interface{}) *Component {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 表单label的对齐方式
 */
func (a *Component) LabelAlign(value interface{}) *Component {
    a.Set("labelAlign", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *Component) LabelWidth(value interface{}) *Component {
    a.Set("labelWidth", value)
    return a
}

/**
 * 设置了initAsyncApi以后，默认拉取的时间间隔
 */
func (a *Component) InitCheckInterval(value interface{}) *Component {
    a.Set("initCheckInterval", value)
    return a
}

/**
 * 禁用回车提交
 */
func (a *Component) PreventEnterSubmit(value interface{}) *Component {
    a.Set("preventEnterSubmit", value)
    return a
}

/**
 * 提交成功后清空本地缓存
 */
func (a *Component) ClearPersistDataAfterSubmit(value interface{}) *Component {
    a.Set("clearPersistDataAfterSubmit", value)
    return a
}

/**
 * 设置主键 id, 当设置后，检测表单是否完成时（asyncApi），只会携带此数据。
 */
func (a *Component) PrimaryField(value interface{}) *Component {
    a.Set("primaryField", value)
    return a
}

/**
 * 设置了initAsyncApi后，默认会从返回数据的data.finished来判断是否完成，也可以设置成其他的xxx，就会从data.xxx中获取
 */
func (a *Component) InitFinishedField(value interface{}) *Component {
    a.Set("initFinishedField", value)
    return a
}

/**
 * 是否开启本地缓存
 */
func (a *Component) PersistData(value interface{}) *Component {
    a.Set("persistData", value)
    return a
}

/**
 */
func (a *Component) Redirect(value interface{}) *Component {
    a.Set("redirect", value)
    return a
}

/**
 * 表单初始先提交一次，联动的时候有用
 */
func (a *Component) SubmitOnInit(value interface{}) *Component {
    a.Set("submitOnInit", value)
    return a
}

/**
 */
func (a *Component) Tabs(value interface{}) *Component {
    a.Set("tabs", value)
    return a
}

/**
 * 页面离开提示，为了防止页面不小心跳转而导致表单没有保存。
 */
func (a *Component) PromptPageLeave(value interface{}) *Component {
    a.Set("promptPageLeave", value)
    return a
}

/**
 * 提交完后重置表单
 */
func (a *Component) ResetAfterSubmit(value interface{}) *Component {
    a.Set("resetAfterSubmit", value)
    return a
}

/**
 * Form 也可以配置 feedback。
 */
func (a *Component) Feedback(value interface{}) *Component {
    a.Set("feedback", value)
    return a
}

/**
 * 设置此属性后，表单提交发送保存接口后，还会继续轮询请求该接口，直到返回 finished 属性为 true 才 结束。
 */
func (a *Component) AsyncApi(value interface{}) *Component {
    a.Set("asyncApi", value)
    return a
}

/**
 * 是否自动将第一个表单元素聚焦。
 */
func (a *Component) AutoFocus(value interface{}) *Component {
    a.Set("autoFocus", value)
    return a
}

/**
 * 是否用 panel 包裹起来
 */
func (a *Component) WrapWithPanel(value interface{}) *Component {
    a.Set("wrapWithPanel", value)
    return a
}

/**
 * Debug面板配置
 */
func (a *Component) DebugConfig(value interface{}) *Component {
    a.Set("debugConfig", value)
    return a
}

/**
 * 如果是水平排版，这个属性可以细化水平排版的左右宽度占比。
 */
func (a *Component) Horizontal(value interface{}) *Component {
    a.Set("horizontal", value)
    return a
}

/**
 * 配置容器 panel className
 */
func (a *Component) PanelClassName(value interface{}) *Component {
    a.Set("panelClassName", value)
    return a
}
/**
 */
func (a *Component) FieldSet(value interface{}) *Component {
    a.Set("fieldSet", value)
    return a
}

/**
 * 是否开启调试，开启后会在顶部实时显示表单项数据。
 */
func (a *Component) Debug(value interface{}) *Component {
    a.Set("debug", value)
    return a
}

/**
 * 按钮集合，会固定在底部显示。
 */
func (a *Component) Actions(value interface{}) *Component {
    a.Set("actions", value)
    return a
}

/**
 * 开启本地缓存后限制保存哪些 key
 */
func (a *Component) PersistDataKeys(value interface{}) *Component {
    a.Set("persistDataKeys", value)
    return a
}

/**
 * 提交后清空表单
 */
func (a *Component) ClearAfterSubmit(value interface{}) *Component {
    a.Set("clearAfterSubmit", value)
    return a
}

/**
 * Form 用来获取初始数据的 api,与initApi不同的是，会一直轮询请求该接口，直到返回 finished 属性为 true 才 结束。
 */
func (a *Component) InitAsyncApi(value interface{}) *Component {
    a.Set("initAsyncApi", value)
    return a
}

/**
 * 如果决定结束的字段名不是 `finished` 请设置此属性，比如 `is_success`
 */
func (a *Component) FinishedField(value interface{}) *Component {
    a.Set("finishedField", value)
    return a
}

/**
 * 组合校验规则，选填
 */
func (a *Component) Rules(value interface{}) *Component {
    a.Set("rules", value)
    return a
}

/**
 * 轮询请求的时间间隔，默认为 3秒。设置 asyncApi 才有效
 */
func (a *Component) CheckInterval(value interface{}) *Component {
    a.Set("checkInterval", value)
    return a
}

/**
 * 具体的提示信息，选填。
 */
func (a *Component) PromptPageLeaveMessage(value interface{}) *Component {
    a.Set("promptPageLeaveMessage", value)
    return a
}

/**
 * 表单项显示为几列
 */
func (a *Component) ColumnCount(value interface{}) *Component {
    a.Set("columnCount", value)
    return a
}

/**
 * 默认的提交按钮名称，如果设置成空，则可以把默认按钮去掉。
 */
func (a *Component) SubmitText(value interface{}) *Component {
    a.Set("submitText", value)
    return a
}

/**
 * 数据录入配置，自动填充或者参照录入
 */
 func (a *Component) AutoFill(value interface{}) *Component {
	a.Set("autoFill", value)
	return a
}

/**
 * 默认值
 */
func (a *Component) DefaultValue(value interface{}) *Component {
	a.Set("defaultValue", value)
	return a
}

/**
 * 表单项描述
 */
func (a *Component) Description(value interface{}) *Component {
	a.Set("description", value)
	return a
}

/**
 * 是否内联
 */
func (a *Component) Inline(value bool) *Component {
	a.Set("inline", value)
	return a
}

/**
 * 表单控制器类名
 */
func (a *Component) InputClassName(value interface{}) *Component {
	a.Set("inputClassName", value)
	return a
}

/**
 * key 的提示信息的
 */
func (a *Component) KeyPlaceholder(value interface{}) *Component {
	a.Set("keyPlaceholder", value)
	return a
}

/**
 * label 的类名
 */
func (a *Component) LabelClassName(value interface{}) *Component {
	a.Set("labelClassName", value)
	return a
}

/**
 * 表单项标签描述
 */
func (a *Component) LabelRemark(value interface{}) *Component {
	a.Set("labelRemark", value)
	return a
}

/**
 * 通过表达式来配置当前表单项是否为必填。
 */
func (a *Component) RequiredOn(value interface{}) *Component {
	a.Set("requiredOn", value)
	return a
}

/**
 * 表单校验接口
 */
func (a *Component) ValidateApi(value interface{}) *Component {
	a.Set("validateApi", value)
	return a
}

/**
 * 表单项值格式验证，支持设置多个，多个规则用英文逗号隔开。
 */
func (a *Component) Validations(value interface{}) *Component {
	a.Set("validations", value)
	return a
}

/**
 * 表单默认值
 */
func (a *Component) Value(value interface{}) *Component {
	a.Set("value", value)
	return a
}

/**
 * value 的提示信息的
 */
func (a *Component) ValuePlaceholder(value interface{}) *Component {
	a.Set("valuePlaceholder", value)
	return a
}

/**
 * 值类型
 */
func (a *Component) ValueType(value interface{}) *Component {
	a.Set("valueType", value)
	return a
}

/**
 * 确认删除时的提示
 */
func (a *Component) DeleteConfirmText(value interface{}) *Component {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * 新增按钮CSS类名
 */
func (a *Component) AddButtonClassName(value interface{}) *Component {
    a.Set("addButtonClassName", value)
    return a
}

/**
 * Add at top
 */
func (a *Component) Addattop(value interface{}) *Component {
    a.Set("addattop", value)
    return a
}

/**
 * 当扁平化开启的时候，是否用分隔符的形式发送给后端，否则采用array的方式
 */
func (a *Component) JoinValues(value interface{}) *Component {
    a.Set("joinValues", value)
    return a
}

/**
 * 是否可多选
 */
func (a *Component) Multiple(value interface{}) *Component {
    a.Set("multiple", value)
    return a
}

/**
 */
func (a *Component) Desc(value interface{}) *Component {
    a.Set("desc", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *Component) ClearValueOnHidden(value interface{}) *Component {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 数组输入框的子项
 */
func (a *Component) Items(value interface{}) *Component {
    a.Set("items", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *Component) Hint(value interface{}) *Component {
    a.Set("hint", value)
    return a
}

/**
 * 限制最小个数
 */
func (a *Component) MinLength(value interface{}) *Component {
    a.Set("minLength", value)
    return a
}

/**
 * 如果是水平排版，这个属性可以细化水平排版的左右宽度占比。
 */
func (a *Component) SubFormHorizontal(value interface{}) *Component {
    a.Set("subFormHorizontal", value)
    return a
}

/**
 * 允许为空，如果子表单项里面配置验证器，且又是单条模式。可以允许用户选择清空（不填）。
 */
func (a *Component) Nullable(value interface{}) *Component {
    a.Set("nullable", value)
    return a
}

/**
 */
func (a *Component) InitAutoFill(value interface{}) *Component {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 限制最大个数
 */
func (a *Component) MaxLength(value interface{}) *Component {
    a.Set("maxLength", value)
    return a
}

/**
 * 符合某类条件后才渲染的schema
 */
func (a *Component) Conditions(value interface{}) *Component {
    a.Set("conditions", value)
    return a
}

/**
 * 数据比较多，比较卡时，可以试试开启。
 */
func (a *Component) LazyLoad(value interface{}) *Component {
    a.Set("lazyLoad", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *Component) ValidateOnChange(value interface{}) *Component {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *Component) DescriptionClassName(value interface{}) *Component {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 是否可切换条件，配合`conditions`使用
 */
func (a *Component) TypeSwitchable(value interface{}) *Component {
    a.Set("typeSwitchable", value)
    return a
}

/**
 * 内部单组表单项的类名
 */
func (a *Component) FormClassName(value interface{}) *Component {
    a.Set("formClassName", value)
    return a
}

/**
 * 新增按钮文字
 */
func (a *Component) AddButtonText(value interface{}) *Component {
    a.Set("addButtonText", value)
    return a
}

/**
 * Tabs 的展示模式。
 * 可选值:  | line | card | radio
 */
func (a *Component) TabsStyle(value interface{}) *Component {
    a.Set("tabsStyle", value)
    return a
}

/**
 * 是否可删除
 */
func (a *Component) Removable(value interface{}) *Component {
    a.Set("removable", value)
    return a
}

/**
 * 配置同步字段。只有 `strictMode` 为 `false` 时有效。 如果 Combo 层级比较深，底层的获取外层的数据可能不同步。 但是给 combo 配置这个属性就能同步下来。输入格式：`["os"]`
 */
func (a *Component) SyncFields(value interface{}) *Component {
    a.Set("syncFields", value)
    return a
}

/**
 * 是否只读
 */
func (a *Component) ReadOnly(value interface{}) *Component {
    a.Set("readOnly", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *Component) ValidationErrors(value interface{}) *Component {
    a.Set("validationErrors", value)
    return a
}

/**
 * 单组表单项初始值。默认为 `{}`
 */
func (a *Component) Scaffold(value interface{}) *Component {
    a.Set("scaffold", value)
    return a
}

/**
 * 删除时调用的api
 */
func (a *Component) DeleteApi(value interface{}) *Component {
    a.Set("deleteApi", value)
    return a
}

/**
 * 是否可新增
 */
func (a *Component) Addable(value interface{}) *Component {
    a.Set("addable", value)
    return a
}

/**
 * 可拖拽排序的提示信息。
 */
func (a *Component) DraggableTip(value interface{}) *Component {
    a.Set("draggableTip", value)
    return a
}

/**
 * 采用 Tabs 展示方式？
 */
func (a *Component) TabsMode(value interface{}) *Component {
    a.Set("tabsMode", value)
    return a
}

/**
 * 只读条件
 */
func (a *Component) ReadOnlyOn(value interface{}) *Component {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 选项卡标题的生成模板。
 */
func (a *Component) TabsLabelTpl(value interface{}) *Component {
    a.Set("tabsLabelTpl", value)
    return a
}

/**
 * 严格模式，为了性能默认不开的。
 */
func (a *Component) StrictMode(value interface{}) *Component {
    a.Set("strictMode", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *Component) ExtraName(value interface{}) *Component {
    a.Set("extraName", value)
    return a
}

/**
 * 是否将结果扁平化(去掉name),只有当controls的length为1且multiple为true的时候才有效
 */
func (a *Component) Flat(value interface{}) *Component {
    a.Set("flat", value)
    return a
}

/**
 * 当扁平化开启并且joinValues为true时，用什么分隔符
 */
func (a *Component) Delimiter(value interface{}) *Component {
    a.Set("delimiter", value)
    return a
}

/**
 * 子表单的模式。
 * 可选值: normal | horizontal | inline
 */
func (a *Component) SubFormMode(value interface{}) *Component {
    a.Set("subFormMode", value)
    return a
}

/**
 * 是否含有边框
 */
func (a *Component) NoBorder(value interface{}) *Component {
    a.Set("noBorder", value)
    return a
}

/**
 * 是否多行模式，默认一行展示完
 */
func (a *Component) MultiLine(value interface{}) *Component {
    a.Set("multiLine", value)
    return a
}

/**
 */
func (a *Component) UpdatePristineAfterStoreDataReInit(value interface{}) *Component {
    a.Set("updatePristineAfterStoreDataReInit", value)
    return a
}

/**
 * 勾选值
 */
func (a *Component) TrueValue(value interface{}) *Component {
    a.Set("trueValue", value)
    return a
}

/**
 */
func (a *Component) Checked(value interface{}) *Component {
    a.Set("checked", value)
    return a
}

/**
 * 可选值: default | button
 */
func (a *Component) OptionType(value interface{}) *Component {
    a.Set("optionType", value)
    return a
}

/**
 */
func (a *Component) Partial(value interface{}) *Component {
    a.Set("partial", value)
    return a
}

/**
 * 选项说明
 */
func (a *Component) Option(value interface{}) *Component {
    a.Set("option", value)
    return a
}

/**
 * 未勾选值
 */
func (a *Component) FalseValue(value interface{}) *Component {
    a.Set("falseValue", value)
    return a
}

/**
 * 控制新增弹框设置项
 */
func (a *Component) AddDialog(value interface{}) *Component {
    a.Set("addDialog", value)
    return a
}
/**
 * 多选模式，值太多时是否避免折行
 */
func (a *Component) ValuesNoWrap(value interface{}) *Component {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * 点清除按钮时，将表单项设置成当前配置的值。
 */
func (a *Component) ResetValue(value interface{}) *Component {
    a.Set("resetValue", value)
    return a
}

/**
 * 是否可清除。
 */
func (a *Component) Clearable(value interface{}) *Component {
    a.Set("clearable", value)
    return a
}

/**
 * 新增文字
 */
func (a *Component) CreateBtnLabel(value interface{}) *Component {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * 默认选择选项第一个值。
 */
func (a *Component) SelectFirst(value interface{}) *Component {
    a.Set("selectFirst", value)
    return a
}

/**
 * 编辑时调用的 API
 */
func (a *Component) EditApi(value interface{}) *Component {
    a.Set("editApi", value)
    return a
}

/**
 * 选项修改的表单项
 */
func (a *Component) EditControls(value interface{}) *Component {
    a.Set("editControls", value)
    return a
}

/**
 * 新增时的表单项。
 */
func (a *Component) AddControls(value interface{}) *Component {
    a.Set("addControls", value)
    return a
}

/**
 * 是否可以新增
 */
func (a *Component) Creatable(value interface{}) *Component {
    a.Set("creatable", value)
    return a
}

/**
 * 控制编辑弹框设置项
 */
func (a *Component) EditDialog(value interface{}) *Component {
    a.Set("editDialog", value)
    return a
}

/**
 * 选项集合
 */
func (a *Component) Options(value interface{}) *Component {
    a.Set("options", value)
    return a
}

/**
 * 懒加载字段
 */
func (a *Component) DeferField(value interface{}) *Component {
    a.Set("deferField", value)
    return a
}

/**
 * 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
 */
func (a *Component) ExtractValue(value interface{}) *Component {
    a.Set("extractValue", value)
    return a
}

/**
 * 添加时调用的接口
 */
func (a *Component) AddApi(value interface{}) *Component {
    a.Set("addApi", value)
    return a
}

/**
 * 是否可以编辑
 */
func (a *Component) Editable(value interface{}) *Component {
    a.Set("editable", value)
    return a
}

/**
 * 数据展示模式 无限加载 or 分页
 * 可选值: more | pagination
 */
func (a *Component) LoadType(value interface{}) *Component {
    a.Set("loadType", value)
    return a
}

/**
 * 是否可以选择数据，外部事件动作
 */
func (a *Component) Selectable(value interface{}) *Component {
    a.Set("selectable", value)
    return a
}

/**
 * 当一次性渲染太多列上有用，默认为 100，可以用来提升表格渲染性能
 */
func (a *Component) LazyRenderAfter(value interface{}) *Component {
    a.Set("lazyRenderAfter", value)
    return a
}

/**
 * 是否展示边框
 */
func (a *Component) Bordered(value interface{}) *Component {
    a.Set("bordered", value)
    return a
}

/**
 * 批量操作最大限制数
 */
func (a *Component) MaxKeepItemSelectionLength(value interface{}) *Component {
    a.Set("maxKeepItemSelectionLength", value)
    return a
}

/**
 * 是否固定内容行高度
 */
func (a *Component) LineHeight(value interface{}) *Component {
    a.Set("lineHeight", value)
    return a
}

/**
 * 顶部区域CSS类名
 */
func (a *Component) HeaderToolbarClassName(value interface{}) *Component {
    a.Set("headerToolbarClassName", value)
    return a
}

/**
 * 指定挂载dom
 */
func (a *Component) PopOverContainer(value interface{}) *Component {
    a.Set("popOverContainer", value)
    return a
}

/**
 * 数据源嵌套自定义字段名
 */
func (a *Component) ChildrenColumnName(value interface{}) *Component {
    a.Set("childrenColumnName", value)
    return a
}

/**
 * 表格行可展开配置
 */
func (a *Component) Expandable(value interface{}) *Component {
    a.Set("expandable", value)
    return a
}

/**
 * 粘性头部
 */
func (a *Component) Sticky(value interface{}) *Component {
    a.Set("sticky", value)
    return a
}

/**
 * 是否展示行角标
 */
func (a *Component) ShowBadge(value interface{}) *Component {
    a.Set("showBadge", value)
    return a
}

/**
 * 底部区域CSS类名
 */
func (a *Component) FooterToolbarClassName(value interface{}) *Component {
    a.Set("footerToolbarClassName", value)
    return a
}

/**
 * 表格可选择配置
 */
func (a *Component) RowSelection(value interface{}) *Component {
    a.Set("rowSelection", value)
    return a
}

/**
 * 加载中
 */
func (a *Component) Loading(value interface{}) *Component {
    a.Set("loading", value)
    return a
}

/**
 * 是否展示已选数据区域，仅当selectable为 true 时生效
 */
func (a *Component) ShowSelection(value interface{}) *Component {
    a.Set("showSelection", value)
    return a
}

/**
 * 多选、嵌套展开记录的ID字段名 默认id
 */
func (a *Component) KeyField(value interface{}) *Component {
    a.Set("keyField", value)
    return a
}

/**
 * 按钮样式级别
 */
func (a *Component) BtnLevel(value interface{}) *Component {
    a.Set("btnLevel", value)
    return a
}

/**
 * 平铺展示？
 */
func (a *Component) Tiled(value interface{}) *Component {
    a.Set("tiled", value)
    return a
}

/**
 */
func (a *Component) BtnClassName(value interface{}) *Component {
    a.Set("btnClassName", value)
    return a
}

/**
 * 垂直展示？
 */
func (a *Component) Vertical(value interface{}) *Component {
    a.Set("vertical", value)
    return a
}

/**
 * 按钮集合
 */
func (a *Component) Buttons(value interface{}) *Component {
    a.Set("buttons", value)
    return a
}

/**
 */
func (a *Component) BtnActiveClassName(value interface{}) *Component {
    a.Set("btnActiveClassName", value)
    return a
}

/**
 * 按钮选中的样式级别
 */
func (a *Component) BtnActiveLevel(value interface{}) *Component {
    a.Set("btnActiveLevel", value)
    return a
}

/**
 */
func (a *Component) IgnoreConfirm(value interface{}) *Component {
    a.Set("ignoreConfirm", value)
    return a
}

/**
 * 是否开启请求隔离, 主要用于隔离联动CRUD, Service的请求
 */
func (a *Component) IsolateScope(value interface{}) *Component {
    a.Set("isolateScope", value)
    return a
}