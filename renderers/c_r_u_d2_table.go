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
/**
 * 静默拉取
 */
func (a *CRUD2Table) SilentPolling(value string) *CRUD2Table {
    a.Set("silentPolling", value)
    return a
}

/**
 * 是否可以多选数据，仅当selectable为 true 时生效
 */
func (a *CRUD2Table) Multiple(value string) *CRUD2Table {
    a.Set("multiple", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *CRUD2Table) HiddenOn(value string) *CRUD2Table {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *CRUD2Table) StaticInputClassName(value string) *CRUD2Table {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *CRUD2Table) ClassName(value string) *CRUD2Table {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *CRUD2Table) Hidden(value string) *CRUD2Table {
    a.Set("hidden", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *CRUD2Table) Id(value string) *CRUD2Table {
    a.Set("id", value)
    return a
}

/**
 * 初始化数据 API
 */
func (a *CRUD2Table) Api(value string) *CRUD2Table {
    a.Set("api", value)
    return a
}

/**
 * 设置自动刷新时间
 */
func (a *CRUD2Table) Interval(value string) *CRUD2Table {
    a.Set("interval", value)
    return a
}

/**
 * 是否固定内容行高度
 */
func (a *CRUD2Table) LineHeight(value string) *CRUD2Table {
    a.Set("lineHeight", value)
    return a
}

/**
 * 指定表尾
 */
func (a *CRUD2Table) Footer(value string) *CRUD2Table {
    a.Set("footer", value)
    return a
}

/**
 * 操作列配置
 */
func (a *CRUD2Table) Actions(value string) *CRUD2Table {
    a.Set("actions", value)
    return a
}

/**
 * 数据展示模式 无限加载 or 分页
 * 可选值: more | pagination
 */
func (a *CRUD2Table) LoadType(value string) *CRUD2Table {
    a.Set("loadType", value)
    return a
}

/**
 * 无限加载时，根据此项设置其每页加载数量，可以不限制
 */
func (a *CRUD2Table) PerPage(value string) *CRUD2Table {
    a.Set("perPage", value)
    return a
}

/**
 * 是否隐藏快速编辑的按钮。
 */
func (a *CRUD2Table) HideQuickSaveBtn(value string) *CRUD2Table {
    a.Set("hideQuickSaveBtn", value)
    return a
}

/**
 * 指定挂载dom
 */
func (a *CRUD2Table) PopOverContainer(value string) *CRUD2Table {
    a.Set("popOverContainer", value)
    return a
}

/**
 * 数据源嵌套自定义字段名
 */
func (a *CRUD2Table) ChildrenColumnName(value string) *CRUD2Table {
    a.Set("childrenColumnName", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *CRUD2Table) StaticOn(value string) *CRUD2Table {
    a.Set("staticOn", value)
    return a
}

/**
 * 自定义行样式
 */
func (a *CRUD2Table) RowClassNameExpr(value string) *CRUD2Table {
    a.Set("rowClassNameExpr", value)
    return a
}

/**
 * 设置分页页码字段名。
 */
func (a *CRUD2Table) PageField(value string) *CRUD2Table {
    a.Set("pageField", value)
    return a
}

/**
 * 表格行可展开配置
 */
func (a *CRUD2Table) Expandable(value string) *CRUD2Table {
    a.Set("expandable", value)
    return a
}

/**
 * 表格是否可以获取父级数据域值，默认为false
 */
func (a *CRUD2Table) CanAccessSuperData(value string) *CRUD2Table {
    a.Set("canAccessSuperData", value)
    return a
}

/**
 * 是否为前端单次加载模式，可以用来实现前端分页。
 */
func (a *CRUD2Table) LoadDataOnce(value string) *CRUD2Table {
    a.Set("loadDataOnce", value)
    return a
}

/**
 * 内容区域占满屏幕剩余空间
 */
func (a *CRUD2Table) AutoFillHeight(value string) *CRUD2Table {
    a.Set("autoFillHeight", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *CRUD2Table) DisabledOn(value string) *CRUD2Table {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *CRUD2Table) Visible(value string) *CRUD2Table {
    a.Set("visible", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *CRUD2Table) EditorSetting(value string) *CRUD2Table {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *CRUD2Table) LoadingConfig(value string) *CRUD2Table {
    a.Set("loadingConfig", value)
    return a
}

/**
 * 加载中
 */
func (a *CRUD2Table) Loading(value string) *CRUD2Table {
    a.Set("loading", value)
    return a
}

/**
 * 是否展示表头
 */
func (a *CRUD2Table) ShowHeader(value string) *CRUD2Table {
    a.Set("showHeader", value)
    return a
}

/**
 * 快速编辑后用来批量保存的 API
 */
func (a *CRUD2Table) QuickSaveApi(value string) *CRUD2Table {
    a.Set("quickSaveApi", value)
    return a
}

/**
 * 是否将过滤条件的参数同步到地址栏,默认为true
 */
func (a *CRUD2Table) SyncLocation(value string) *CRUD2Table {
    a.Set("syncLocation", value)
    return a
}

/**
 * 底部区域CSS类名
 */
func (a *CRUD2Table) FooterToolbarClassName(value string) *CRUD2Table {
    a.Set("footerToolbarClassName", value)
    return a
}

/**
 */
func (a *CRUD2Table) StaticSchema(value string) *CRUD2Table {
    a.Set("staticSchema", value)
    return a
}

/**
 * 指定内容区的展示模式。
 */
func (a *CRUD2Table) Mode(value string) *CRUD2Table {
    a.Set("mode", value)
    return a
}

/**
 * 是否展示已选数据区域，仅当selectable为 true 时生效
 */
func (a *CRUD2Table) ShowSelection(value string) *CRUD2Table {
    a.Set("showSelection", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *CRUD2Table) Static(value string) *CRUD2Table {
    a.Set("static", value)
    return a
}

/**
 * 底部区域
 */
func (a *CRUD2Table) FooterToolbar(value string) *CRUD2Table {
    a.Set("footerToolbar", value)
    return a
}

/**
 * 是否开启Query信息转换，开启后将会对url中的Query进行转换，默认开启，默认仅转化布尔值
 */
func (a *CRUD2Table) ParsePrimitiveQuery(value string) *CRUD2Table {
    a.Set("parsePrimitiveQuery", value)
    return a
}

/**
 * 是否展示行角标
 */
func (a *CRUD2Table) ShowBadge(value string) *CRUD2Table {
    a.Set("showBadge", value)
    return a
}

/**
 * 快速编辑配置成及时保存时使用的 API
 */
func (a *CRUD2Table) QuickSaveItemApi(value string) *CRUD2Table {
    a.Set("quickSaveItemApi", value)
    return a
}

/**
 * 是否禁用
 */
func (a *CRUD2Table) Disabled(value string) *CRUD2Table {
    a.Set("disabled", value)
    return a
}

/**
 * 可选值: fixed | auto
 */
func (a *CRUD2Table) TableLayout(value string) *CRUD2Table {
    a.Set("tableLayout", value)
    return a
}

/**
 * 保存排序的 api
 */
func (a *CRUD2Table) SaveOrderApi(value string) *CRUD2Table {
    a.Set("saveOrderApi", value)
    return a
}

/**
 */
func (a *CRUD2Table) Name(value string) *CRUD2Table {
    a.Set("name", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *CRUD2Table) StaticClassName(value string) *CRUD2Table {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *CRUD2Table) StopAutoRefreshWhen(value string) *CRUD2Table {
    a.Set("stopAutoRefreshWhen", value)
    return a
}

/**
 * 顶部区域
 */
func (a *CRUD2Table) HeaderToolbar(value string) *CRUD2Table {
    a.Set("headerToolbar", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *CRUD2Table) OnEvent(value string) *CRUD2Table {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *CRUD2Table) StaticLabelClassName(value string) *CRUD2Table {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 表格可选择配置
 */
func (a *CRUD2Table) RowSelection(value string) *CRUD2Table {
    a.Set("rowSelection", value)
    return a
}

/**
 * 粘性头部
 */
func (a *CRUD2Table) Sticky(value string) *CRUD2Table {
    a.Set("sticky", value)
    return a
}

/**
 * 批量操作最大限制数
 */
func (a *CRUD2Table) MaxKeepItemSelectionLength(value string) *CRUD2Table {
    a.Set("maxKeepItemSelectionLength", value)
    return a
}

/**
 * 翻页时是否保留用户已选的数据
 */
func (a *CRUD2Table) KeepItemSelectionOnPageChange(value string) *CRUD2Table {
    a.Set("keepItemSelectionOnPageChange", value)
    return a
}

/**
 * 行标识符，默认为id
 */
func (a *CRUD2Table) PrimaryField(value string) *CRUD2Table {
    a.Set("primaryField", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *CRUD2Table) VisibleOn(value string) *CRUD2Table {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否自动跳顶部，当切分页的时候。
 */
func (a *CRUD2Table) AutoJumpToTopOnPagerChange(value string) *CRUD2Table {
    a.Set("autoJumpToTopOnPagerChange", value)
    return a
}

/**
 * 顶部区域CSS类名
 */
func (a *CRUD2Table) HeaderToolbarClassName(value string) *CRUD2Table {
    a.Set("headerToolbarClassName", value)
    return a
}

/**
 * 也可以直接从环境变量中读取，但是不太推荐。
 */
func (a *CRUD2Table) Source(value string) *CRUD2Table {
    a.Set("source", value)
    return a
}

/**
 * 表格可自定义列
 */
func (a *CRUD2Table) ColumnsTogglable(value string) *CRUD2Table {
    a.Set("columnsTogglable", value)
    return a
}

/**
 * 重新加载的组件名称
 */
func (a *CRUD2Table) Reload(value string) *CRUD2Table {
    a.Set("reload", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *CRUD2Table) UseMobileUI(value string) *CRUD2Table {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 设置分页一页显示的多少条数据的字段名。
 */
func (a *CRUD2Table) PerPageField(value string) *CRUD2Table {
    a.Set("perPageField", value)
    return a
}

/**
 * 多选、嵌套展开记录的ID字段名 默认id
 */
func (a *CRUD2Table) KeyField(value string) *CRUD2Table {
    a.Set("keyField", value)
    return a
}

/**
 * 是否展示边框
 */
func (a *CRUD2Table) Bordered(value string) *CRUD2Table {
    a.Set("bordered", value)
    return a
}

/**
 * 当一次性渲染太多列上有用，默认为 100，可以用来提升表格渲染性能
 */
func (a *CRUD2Table) LazyRenderAfter(value string) *CRUD2Table {
    a.Set("lazyRenderAfter", value)
    return a
}

/**
 * 是否可以选择数据，外部事件动作
 */
func (a *CRUD2Table) Selectable(value string) *CRUD2Table {
    a.Set("selectable", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *CRUD2Table) StaticPlaceholder(value string) *CRUD2Table {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 指定为 CRUD2 渲染器。
 */
func (a *CRUD2Table) Type(value string) *CRUD2Table {
    a.Set("type", value)
    return a
}

/**
 * 是否将接口返回的内容自动同步到地址栏，前提是开启了同步地址栏。
 */
func (a *CRUD2Table) SyncResponse2Query(value string) *CRUD2Table {
    a.Set("syncResponse2Query", value)
    return a
}

/**
 * 表格标题
 */
func (a *CRUD2Table) Title(value string) *CRUD2Table {
    a.Set("title", value)
    return a
}

/**
 * 表格列配置
 */
func (a *CRUD2Table) Columns(value string) *CRUD2Table {
    a.Set("columns", value)
    return a
}

/**
 * 行角标内容
 */
func (a *CRUD2Table) ItemBadge(value string) *CRUD2Table {
    a.Set("itemBadge", value)
    return a
}

/**
 * 接口报错信息配置
 */
func (a *CRUD2Table) Messages(value string) *CRUD2Table {
    a.Set("messages", value)
    return a
}

/**
 * 组件样式
 */
func (a *CRUD2Table) Style(value string) *CRUD2Table {
    a.Set("style", value)
    return a
}
