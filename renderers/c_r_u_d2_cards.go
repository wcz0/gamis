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

    a.Set("mode", "cards")
    a.Set("type", "crud2")
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
 * 是否静态展示表达式
 */
func (a *CRUD2Cards) StaticOn(value interface{}) *CRUD2Cards {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *CRUD2Cards) StaticLabelClassName(value interface{}) *CRUD2Cards {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *CRUD2Cards) Style(value interface{}) *CRUD2Cards {
    a.Set("style", value)
    return a
}

/**
 * 设置分页页码字段名。
 */
func (a *CRUD2Cards) PageField(value interface{}) *CRUD2Cards {
    a.Set("pageField", value)
    return a
}

/**
 * 配置某项是否可以点选
 */
func (a *CRUD2Cards) ItemCheckableOn(value interface{}) *CRUD2Cards {
    a.Set("itemCheckableOn", value)
    return a
}

/**
 * 是否为前端单次加载模式，可以用来实现前端分页。
 */
func (a *CRUD2Cards) LoadDataOnce(value interface{}) *CRUD2Cards {
    a.Set("loadDataOnce", value)
    return a
}

/**
 * 是否固顶
 */
func (a *CRUD2Cards) AffixHeader(value interface{}) *CRUD2Cards {
    a.Set("affixHeader", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *CRUD2Cards) VisibleOn(value interface{}) *CRUD2Cards {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *CRUD2Cards) StaticInputClassName(value interface{}) *CRUD2Cards {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 快速编辑配置成及时保存时使用的 API
 */
func (a *CRUD2Cards) QuickSaveItemApi(value interface{}) *CRUD2Cards {
    a.Set("quickSaveItemApi", value)
    return a
}

/**
 * 也可以直接从环境变量中读取，但是不太推荐。
 */
func (a *CRUD2Cards) Source(value interface{}) *CRUD2Cards {
    a.Set("source", value)
    return a
}

/**
 * 底部区域
 */
func (a *CRUD2Cards) Footer(value interface{}) *CRUD2Cards {
    a.Set("footer", value)
    return a
}

/**
 * 指定内容区的展示模式。
 */
func (a *CRUD2Cards) Mode(value interface{}) *CRUD2Cards {
    a.Set("mode", value)
    return a
}

/**
 * 初始化数据 API
 */
func (a *CRUD2Cards) Api(value interface{}) *CRUD2Cards {
    a.Set("api", value)
    return a
}

/**
 * 底部区域
 */
func (a *CRUD2Cards) FooterToolbar(value interface{}) *CRUD2Cards {
    a.Set("footerToolbar", value)
    return a
}

/**
 * 底部区域CSS类名
 */
func (a *CRUD2Cards) FooterToolbarClassName(value interface{}) *CRUD2Cards {
    a.Set("footerToolbarClassName", value)
    return a
}

/**
 * 卡片 CSS 类名
 */
func (a *CRUD2Cards) ItemClassName(value interface{}) *CRUD2Cards {
    a.Set("itemClassName", value)
    return a
}

/**
 * 无数据提示
 */
func (a *CRUD2Cards) Placeholder(value interface{}) *CRUD2Cards {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否显示底部
 */
func (a *CRUD2Cards) ShowFooter(value interface{}) *CRUD2Cards {
    a.Set("showFooter", value)
    return a
}

/**
 * 是否显示头部
 */
func (a *CRUD2Cards) ShowHeader(value interface{}) *CRUD2Cards {
    a.Set("showHeader", value)
    return a
}

/**
 * 标题
 */
func (a *CRUD2Cards) Title(value interface{}) *CRUD2Cards {
    a.Set("title", value)
    return a
}

/**
 * 是否禁用
 */
func (a *CRUD2Cards) Disabled(value interface{}) *CRUD2Cards {
    a.Set("disabled", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *CRUD2Cards) StaticClassName(value interface{}) *CRUD2Cards {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静默拉取
 */
func (a *CRUD2Cards) SilentPolling(value interface{}) *CRUD2Cards {
    a.Set("silentPolling", value)
    return a
}

/**
 * 底部 CSS 类名
 */
func (a *CRUD2Cards) FooterClassName(value interface{}) *CRUD2Cards {
    a.Set("footerClassName", value)
    return a
}

/**
 * 点击卡片的时候是否勾选卡片。
 */
func (a *CRUD2Cards) CheckOnItemClick(value interface{}) *CRUD2Cards {
    a.Set("checkOnItemClick", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *CRUD2Cards) OnEvent(value interface{}) *CRUD2Cards {
    a.Set("onEvent", value)
    return a
}

/**
 * 保存排序的 api
 */
func (a *CRUD2Cards) SaveOrderApi(value interface{}) *CRUD2Cards {
    a.Set("saveOrderApi", value)
    return a
}

/**
 * 顶部区域
 */
func (a *CRUD2Cards) HeaderToolbar(value interface{}) *CRUD2Cards {
    a.Set("headerToolbar", value)
    return a
}

/**
 * 头部 CSS 类名
 */
func (a *CRUD2Cards) HeaderClassName(value interface{}) *CRUD2Cards {
    a.Set("headerClassName", value)
    return a
}

/**
 */
func (a *CRUD2Cards) TestIdBuilder(value interface{}) *CRUD2Cards {
    a.Set("testIdBuilder", value)
    return a
}

/**
 */
func (a *CRUD2Cards) StopAutoRefreshWhen(value interface{}) *CRUD2Cards {
    a.Set("stopAutoRefreshWhen", value)
    return a
}

/**
 * 数据展示模式 无限加载 or 分页
 * 可选值: more | pagination
 */
func (a *CRUD2Cards) LoadType(value interface{}) *CRUD2Cards {
    a.Set("loadType", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *CRUD2Cards) Id(value interface{}) *CRUD2Cards {
    a.Set("id", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *CRUD2Cards) Hidden(value interface{}) *CRUD2Cards {
    a.Set("hidden", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *CRUD2Cards) ClassName(value interface{}) *CRUD2Cards {
    a.Set("className", value)
    return a
}

/**
 * 设置自动刷新时间
 */
func (a *CRUD2Cards) Interval(value interface{}) *CRUD2Cards {
    a.Set("interval", value)
    return a
}

/**
 */
func (a *CRUD2Cards) Name(value interface{}) *CRUD2Cards {
    a.Set("name", value)
    return a
}

/**
 * 是否隐藏快速编辑的按钮。
 */
func (a *CRUD2Cards) HideQuickSaveBtn(value interface{}) *CRUD2Cards {
    a.Set("hideQuickSaveBtn", value)
    return a
}

/**
 * 是否自动跳顶部，当切分页的时候。
 */
func (a *CRUD2Cards) AutoJumpToTopOnPagerChange(value interface{}) *CRUD2Cards {
    a.Set("autoJumpToTopOnPagerChange", value)
    return a
}

/**
 * 是否展示已选数据区域，仅当selectable为 true 时生效
 */
func (a *CRUD2Cards) ShowSelection(value interface{}) *CRUD2Cards {
    a.Set("showSelection", value)
    return a
}

/**
 * 是否将接口返回的内容自动同步到地址栏，前提是开启了同步地址栏。
 */
func (a *CRUD2Cards) SyncResponse2Query(value interface{}) *CRUD2Cards {
    a.Set("syncResponse2Query", value)
    return a
}

/**
 * 行标识符，默认为id
 */
func (a *CRUD2Cards) PrimaryField(value interface{}) *CRUD2Cards {
    a.Set("primaryField", value)
    return a
}

/**
 */
func (a *CRUD2Cards) StaticSchema(value interface{}) *CRUD2Cards {
    a.Set("staticSchema", value)
    return a
}

/**
 * 是否隐藏勾选框
 */
func (a *CRUD2Cards) HideCheckToggler(value interface{}) *CRUD2Cards {
    a.Set("hideCheckToggler", value)
    return a
}

/**
 * 顶部区域
 */
func (a *CRUD2Cards) Header(value interface{}) *CRUD2Cards {
    a.Set("header", value)
    return a
}

/**
 * 配置某项是否可拖拽排序，前提是要开启拖拽功能
 */
func (a *CRUD2Cards) ItemDraggableOn(value interface{}) *CRUD2Cards {
    a.Set("itemDraggableOn", value)
    return a
}

/**
 * 可以用来作为值的字段
 */
func (a *CRUD2Cards) ValueField(value interface{}) *CRUD2Cards {
    a.Set("valueField", value)
    return a
}

/**
 */
func (a *CRUD2Cards) Testid(value interface{}) *CRUD2Cards {
    a.Set("testid", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *CRUD2Cards) EditorSetting(value interface{}) *CRUD2Cards {
    a.Set("editorSetting", value)
    return a
}

/**
 * 无限加载时，根据此项设置其每页加载数量，可以不限制
 */
func (a *CRUD2Cards) PerPage(value interface{}) *CRUD2Cards {
    a.Set("perPage", value)
    return a
}

/**
 */
func (a *CRUD2Cards) Card(value interface{}) *CRUD2Cards {
    a.Set("card", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *CRUD2Cards) StaticPlaceholder(value interface{}) *CRUD2Cards {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 是否可以选择数据，外部事件动作
 */
func (a *CRUD2Cards) Selectable(value interface{}) *CRUD2Cards {
    a.Set("selectable", value)
    return a
}

/**
 * 是否将过滤条件的参数同步到地址栏,默认为true
 */
func (a *CRUD2Cards) SyncLocation(value interface{}) *CRUD2Cards {
    a.Set("syncLocation", value)
    return a
}

/**
 * 设置分页一页显示的多少条数据的字段名。
 */
func (a *CRUD2Cards) PerPageField(value interface{}) *CRUD2Cards {
    a.Set("perPageField", value)
    return a
}

/**
 * 是否开启Query信息转换，开启后将会对url中的Query进行转换，默认开启，默认仅转化布尔值
 */
func (a *CRUD2Cards) ParsePrimitiveQuery(value interface{}) *CRUD2Cards {
    a.Set("parsePrimitiveQuery", value)
    return a
}

/**
 * 是否为瀑布流布局？
 */
func (a *CRUD2Cards) MasonryLayout(value interface{}) *CRUD2Cards {
    a.Set("masonryLayout", value)
    return a
}

/**
 * 快速编辑后用来批量保存的 API
 */
func (a *CRUD2Cards) QuickSaveApi(value interface{}) *CRUD2Cards {
    a.Set("quickSaveApi", value)
    return a
}

/**
 * 翻页时是否保留用户已选的数据
 */
func (a *CRUD2Cards) KeepItemSelectionOnPageChange(value interface{}) *CRUD2Cards {
    a.Set("keepItemSelectionOnPageChange", value)
    return a
}

/**
 */
func (a *CRUD2Cards) LoadingConfig(value interface{}) *CRUD2Cards {
    a.Set("loadingConfig", value)
    return a
}

/**
 * 是否显示
 */
func (a *CRUD2Cards) Visible(value interface{}) *CRUD2Cards {
    a.Set("visible", value)
    return a
}

/**
 * 指定为 CRUD2 渲染器。
 */
func (a *CRUD2Cards) Type(value interface{}) *CRUD2Cards {
    a.Set("type", value)
    return a
}

/**
 * 顶部区域CSS类名
 */
func (a *CRUD2Cards) HeaderToolbarClassName(value interface{}) *CRUD2Cards {
    a.Set("headerToolbarClassName", value)
    return a
}

/**
 * 内容区域占满屏幕剩余空间
 */
func (a *CRUD2Cards) AutoFillHeight(value interface{}) *CRUD2Cards {
    a.Set("autoFillHeight", value)
    return a
}

/**
 * 是否固底
 */
func (a *CRUD2Cards) AffixFooter(value interface{}) *CRUD2Cards {
    a.Set("affixFooter", value)
    return a
}

/**
 * 是否可以多选数据，仅当selectable为 true 时生效
 */
func (a *CRUD2Cards) Multiple(value interface{}) *CRUD2Cards {
    a.Set("multiple", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *CRUD2Cards) HiddenOn(value interface{}) *CRUD2Cards {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *CRUD2Cards) Static(value interface{}) *CRUD2Cards {
    a.Set("static", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *CRUD2Cards) UseMobileUI(value interface{}) *CRUD2Cards {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *CRUD2Cards) DisabledOn(value interface{}) *CRUD2Cards {
    a.Set("disabledOn", value)
    return a
}
