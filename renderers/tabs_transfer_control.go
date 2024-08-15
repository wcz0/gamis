package renderers


/**
 * TabsTransfer 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/tabs-transfer

 * @author wcz0
 * @version 6.2.2
 */
type TabsTransferControl struct {
	*BaseRenderer
}

func NewTabsTransferControl() *TabsTransferControl {
    a := &TabsTransferControl{
        BaseRenderer: NewBaseRenderer(),
    }
    a.Set("type", "tabs-transfer")
    return a
}

/**
 * 当 searchResultMode 为 table 时定义表格列信息。
 */
func (a *TabsTransferControl) SearchResultColumns(value interface{}) *TabsTransferControl {
    a.Set("searchResultColumns", value)
    return a
}

/**
 */
func (a *TabsTransferControl) Row(value interface{}) *TabsTransferControl {
    a.Set("row", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *TabsTransferControl) StaticClassName(value interface{}) *TabsTransferControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 默认选择选项第一个值。
 */
func (a *TabsTransferControl) SelectFirst(value interface{}) *TabsTransferControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * 用表达式来配置 source 接口初始要不要拉取
 */
func (a *TabsTransferControl) InitFetchOn(value interface{}) *TabsTransferControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *TabsTransferControl) OnEvent(value interface{}) *TabsTransferControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 右侧列表搜索框提示
 */
func (a *TabsTransferControl) ResultSearchPlaceholder(value interface{}) *TabsTransferControl {
    a.Set("resultSearchPlaceholder", value)
    return a
}

/**
 */
func (a *TabsTransferControl) Validations(value interface{}) *TabsTransferControl {
    a.Set("validations", value)
    return a
}

/**
 * 是否可清除。
 */
func (a *TabsTransferControl) Clearable(value interface{}) *TabsTransferControl {
    a.Set("clearable", value)
    return a
}

/**
 * 新增时的表单项。
 */
func (a *TabsTransferControl) AddControls(value interface{}) *TabsTransferControl {
    a.Set("addControls", value)
    return a
}

/**
 * 是否可以新增
 */
func (a *TabsTransferControl) Creatable(value interface{}) *TabsTransferControl {
    a.Set("creatable", value)
    return a
}

/**
 * ui级联关系，true代表级联选中，false代表不级联，默认为true
 */
func (a *TabsTransferControl) AutoCheckChildren(value interface{}) *TabsTransferControl {
    a.Set("autoCheckChildren", value)
    return a
}

/**
 * 只读条件
 */
func (a *TabsTransferControl) ReadOnlyOn(value interface{}) *TabsTransferControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
 */
func (a *TabsTransferControl) JoinValues(value interface{}) *TabsTransferControl {
    a.Set("joinValues", value)
    return a
}

/**
 * 是否默认都展开
 */
func (a *TabsTransferControl) InitiallyOpen(value interface{}) *TabsTransferControl {
    a.Set("initiallyOpen", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *TabsTransferControl) Name(value interface{}) *TabsTransferControl {
    a.Set("name", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *TabsTransferControl) ValidateApi(value interface{}) *TabsTransferControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 配置 label className
 */
func (a *TabsTransferControl) LabelClassName(value interface{}) *TabsTransferControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 多选模式，值太多时是否避免折行
 */
func (a *TabsTransferControl) ValuesNoWrap(value interface{}) *TabsTransferControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * 右侧结果的标题文字
 */
func (a *TabsTransferControl) ResultTitle(value interface{}) *TabsTransferControl {
    a.Set("resultTitle", value)
    return a
}

/**
 * 是否显示
 */
func (a *TabsTransferControl) Visible(value interface{}) *TabsTransferControl {
    a.Set("visible", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *TabsTransferControl) LabelWidth(value interface{}) *TabsTransferControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *TabsTransferControl) ValidationErrors(value interface{}) *TabsTransferControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 可排序？
 */
func (a *TabsTransferControl) Sortable(value interface{}) *TabsTransferControl {
    a.Set("sortable", value)
    return a
}

/**
 * 统计数字
 */
func (a *TabsTransferControl) Statistics(value interface{}) *TabsTransferControl {
    a.Set("statistics", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *TabsTransferControl) LabelRemark(value interface{}) *TabsTransferControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 点清除按钮时，将表单项设置成当前配置的值。
 */
func (a *TabsTransferControl) ResetValue(value interface{}) *TabsTransferControl {
    a.Set("resetValue", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *TabsTransferControl) Width(value interface{}) *TabsTransferControl {
    a.Set("width", value)
    return a
}

/**
 * 编辑时调用的 API
 */
func (a *TabsTransferControl) EditApi(value interface{}) *TabsTransferControl {
    a.Set("editApi", value)
    return a
}

/**
 * 是否显示剪头
 */
func (a *TabsTransferControl) ShowArrow(value interface{}) *TabsTransferControl {
    a.Set("showArrow", value)
    return a
}

/**
 * 搜索 API
 */
func (a *TabsTransferControl) SearchApi(value interface{}) *TabsTransferControl {
    a.Set("searchApi", value)
    return a
}

/**
 * 分页配置，selectMode为默认和table才会生效
 */
func (a *TabsTransferControl) Pagination(value interface{}) *TabsTransferControl {
    a.Set("pagination", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *TabsTransferControl) ValidateOnChange(value interface{}) *TabsTransferControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 可用来通过 API 拉取 options。
 */
func (a *TabsTransferControl) Source(value interface{}) *TabsTransferControl {
    a.Set("source", value)
    return a
}

/**
 */
func (a *TabsTransferControl) LoadingConfig(value interface{}) *TabsTransferControl {
    a.Set("loadingConfig", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *TabsTransferControl) StaticInputClassName(value interface{}) *TabsTransferControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *TabsTransferControl) UseMobileUI(value interface{}) *TabsTransferControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 选项集合
 */
func (a *TabsTransferControl) Options(value interface{}) *TabsTransferControl {
    a.Set("options", value)
    return a
}

/**
 * 当 selectMode 为 associated 时用来定义左侧的选项
 */
func (a *TabsTransferControl) LeftOptions(value interface{}) *TabsTransferControl {
    a.Set("leftOptions", value)
    return a
}

/**
 * 描述标题
 */
func (a *TabsTransferControl) Label(value interface{}) *TabsTransferControl {
    a.Set("label", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *TabsTransferControl) ClearValueOnHidden(value interface{}) *TabsTransferControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
 */
func (a *TabsTransferControl) ExtractValue(value interface{}) *TabsTransferControl {
    a.Set("extractValue", value)
    return a
}

/**
 * 可搜索？
 */
func (a *TabsTransferControl) Searchable(value interface{}) *TabsTransferControl {
    a.Set("searchable", value)
    return a
}

/**
 * 左侧的标题文字
 */
func (a *TabsTransferControl) SelectTitle(value interface{}) *TabsTransferControl {
    a.Set("selectTitle", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *TabsTransferControl) Remark(value interface{}) *TabsTransferControl {
    a.Set("remark", value)
    return a
}

/**
 * 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
 */
func (a *TabsTransferControl) DeferApi(value interface{}) *TabsTransferControl {
    a.Set("deferApi", value)
    return a
}

/**
 * 占位符
 */
func (a *TabsTransferControl) Placeholder(value interface{}) *TabsTransferControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *TabsTransferControl) Hidden(value interface{}) *TabsTransferControl {
    a.Set("hidden", value)
    return a
}

/**
 * 左侧列表搜索框提示
 */
func (a *TabsTransferControl) SearchPlaceholder(value interface{}) *TabsTransferControl {
    a.Set("searchPlaceholder", value)
    return a
}

/**
 * 当在value值未匹配到当前options中的选项时，是否value值对应文本飘红显示
 */
func (a *TabsTransferControl) ShowInvalidMatch(value interface{}) *TabsTransferControl {
    a.Set("showInvalidMatch", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *TabsTransferControl) Inline(value interface{}) *TabsTransferControl {
    a.Set("inline", value)
    return a
}

/**
 * 配置 input className
 */
func (a *TabsTransferControl) InputClassName(value interface{}) *TabsTransferControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *TabsTransferControl) Description(value interface{}) *TabsTransferControl {
    a.Set("description", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *TabsTransferControl) AutoFill(value interface{}) *TabsTransferControl {
    a.Set("autoFill", value)
    return a
}

/**
 */
func (a *TabsTransferControl) Type(value interface{}) *TabsTransferControl {
    a.Set("type", value)
    return a
}

/**
 * 结果（右则）列表的检索功能，当设置为true时，可以通过输入检索模糊匹配检索内容
 */
func (a *TabsTransferControl) ResultSearchable(value interface{}) *TabsTransferControl {
    a.Set("resultSearchable", value)
    return a
}

/**
 * 树形模式下，仅选中子节点
 */
func (a *TabsTransferControl) OnlyChildren(value interface{}) *TabsTransferControl {
    a.Set("onlyChildren", value)
    return a
}

/**
 * 是否只读
 */
func (a *TabsTransferControl) ReadOnly(value interface{}) *TabsTransferControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *TabsTransferControl) StaticPlaceholder(value interface{}) *TabsTransferControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 选项删除 API
 */
func (a *TabsTransferControl) DeleteApi(value interface{}) *TabsTransferControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * 是否可以编辑
 */
func (a *TabsTransferControl) Editable(value interface{}) *TabsTransferControl {
    a.Set("editable", value)
    return a
}

/**
 */
func (a *TabsTransferControl) StaticSchema(value interface{}) *TabsTransferControl {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *TabsTransferControl) TestIdBuilder(value interface{}) *TabsTransferControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 结果面板是否追踪显示
 */
func (a *TabsTransferControl) ResultListModeFollowSelect(value interface{}) *TabsTransferControl {
    a.Set("resultListModeFollowSelect", value)
    return a
}

/**
 * 搜索结果展示模式
 * 可选值: table | list | tree | chained
 */
func (a *TabsTransferControl) SearchResultMode(value interface{}) *TabsTransferControl {
    a.Set("searchResultMode", value)
    return a
}

/**
 * 用来丰富选项展示
 */
func (a *TabsTransferControl) MenuTpl(value interface{}) *TabsTransferControl {
    a.Set("menuTpl", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *TabsTransferControl) DescriptionClassName(value interface{}) *TabsTransferControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *TabsTransferControl) Static(value interface{}) *TabsTransferControl {
    a.Set("static", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *TabsTransferControl) EditorSetting(value interface{}) *TabsTransferControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 控制编辑弹框设置项
 */
func (a *TabsTransferControl) EditDialog(value interface{}) *TabsTransferControl {
    a.Set("editDialog", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *TabsTransferControl) Mode(value interface{}) *TabsTransferControl {
    a.Set("mode", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *TabsTransferControl) Value(value interface{}) *TabsTransferControl {
    a.Set("value", value)
    return a
}

/**
 * 配置 source 接口初始拉不拉取。
 */
func (a *TabsTransferControl) InitFetch(value interface{}) *TabsTransferControl {
    a.Set("initFetch", value)
    return a
}

/**
 * 是否禁用
 */
func (a *TabsTransferControl) Disabled(value interface{}) *TabsTransferControl {
    a.Set("disabled", value)
    return a
}

/**
 * 单个选项的高度，主要用于虚拟渲染
 */
func (a *TabsTransferControl) ItemHeight(value interface{}) *TabsTransferControl {
    a.Set("itemHeight", value)
    return a
}

/**
 */
func (a *TabsTransferControl) Desc(value interface{}) *TabsTransferControl {
    a.Set("desc", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *TabsTransferControl) DisabledOn(value interface{}) *TabsTransferControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 在选项数量达到多少时开启虚拟渲染
 */
func (a *TabsTransferControl) VirtualThreshold(value interface{}) *TabsTransferControl {
    a.Set("virtualThreshold", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *TabsTransferControl) Hint(value interface{}) *TabsTransferControl {
    a.Set("hint", value)
    return a
}

/**
 * 是否为多选模式
 */
func (a *TabsTransferControl) Multiple(value interface{}) *TabsTransferControl {
    a.Set("multiple", value)
    return a
}

/**
 * 控制新增弹框设置项
 */
func (a *TabsTransferControl) AddDialog(value interface{}) *TabsTransferControl {
    a.Set("addDialog", value)
    return a
}

/**
 * 勾选展示模式
 * 可选值: table | list | tree | chained | associated
 */
func (a *TabsTransferControl) SelectMode(value interface{}) *TabsTransferControl {
    a.Set("selectMode", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *TabsTransferControl) Size(value interface{}) *TabsTransferControl {
    a.Set("size", value)
    return a
}

/**
 */
func (a *TabsTransferControl) InitAutoFill(value interface{}) *TabsTransferControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 当 selectMode 为 associated 时用来定义左侧的选择模式
 * 可选值: tree | list
 */
func (a *TabsTransferControl) LeftMode(value interface{}) *TabsTransferControl {
    a.Set("leftMode", value)
    return a
}

/**
 * 是否为必填
 */
func (a *TabsTransferControl) Required(value interface{}) *TabsTransferControl {
    a.Set("required", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *TabsTransferControl) VisibleOn(value interface{}) *TabsTransferControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *TabsTransferControl) StaticLabelClassName(value interface{}) *TabsTransferControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 选项删除提示文字。
 */
func (a *TabsTransferControl) DeleteConfirmText(value interface{}) *TabsTransferControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * 选项修改的表单项
 */
func (a *TabsTransferControl) EditControls(value interface{}) *TabsTransferControl {
    a.Set("editControls", value)
    return a
}

/**
 * 懒加载字段
 */
func (a *TabsTransferControl) DeferField(value interface{}) *TabsTransferControl {
    a.Set("deferField", value)
    return a
}

/**
 * 添加时调用的接口
 */
func (a *TabsTransferControl) AddApi(value interface{}) *TabsTransferControl {
    a.Set("addApi", value)
    return a
}

/**
 * 新增文字
 */
func (a *TabsTransferControl) CreateBtnLabel(value interface{}) *TabsTransferControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * 当 selectMode 为 associated 时用来定义右侧的选择模式
 * 可选值: table | list | tree | chained
 */
func (a *TabsTransferControl) RightMode(value interface{}) *TabsTransferControl {
    a.Set("rightMode", value)
    return a
}

/**
 * 用来丰富值的展示
 */
func (a *TabsTransferControl) ValueTpl(value interface{}) *TabsTransferControl {
    a.Set("valueTpl", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *TabsTransferControl) HiddenOn(value interface{}) *TabsTransferControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 组件样式
 */
func (a *TabsTransferControl) Style(value interface{}) *TabsTransferControl {
    a.Set("style", value)
    return a
}

/**
 * 分割符
 */
func (a *TabsTransferControl) Delimiter(value interface{}) *TabsTransferControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 描述标题
 */
func (a *TabsTransferControl) LabelAlign(value interface{}) *TabsTransferControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *TabsTransferControl) Id(value interface{}) *TabsTransferControl {
    a.Set("id", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *TabsTransferControl) ExtraName(value interface{}) *TabsTransferControl {
    a.Set("extraName", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *TabsTransferControl) Horizontal(value interface{}) *TabsTransferControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 是否可删除
 */
func (a *TabsTransferControl) Removable(value interface{}) *TabsTransferControl {
    a.Set("removable", value)
    return a
}

/**
 * 当 selectMode 为 table 时定义表格列信息。
 */
func (a *TabsTransferControl) Columns(value interface{}) *TabsTransferControl {
    a.Set("columns", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *TabsTransferControl) SubmitOnChange(value interface{}) *TabsTransferControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *TabsTransferControl) ClassName(value interface{}) *TabsTransferControl {
    a.Set("className", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *TabsTransferControl) StaticOn(value interface{}) *TabsTransferControl {
    a.Set("staticOn", value)
    return a
}
