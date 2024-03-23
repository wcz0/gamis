package renderers


/**
 * Transfer 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/transfer
 *

*/
type TransferControl struct {
	*BaseRenderer
}

func NewTransferControl() *TransferControl {
    a := &TransferControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "transfer")
    return a
}
/**
 * 控制编辑弹框设置项
 */
func (a *TransferControl) EditDialog(value string) *TransferControl {
    a.Set("editDialog", value)
    return a
}

/**
 * 占位符
 */
func (a *TransferControl) Placeholder(value string) *TransferControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 勾选展示模式
 * 可选值: table | list | tree | chained | associated
 */
func (a *TransferControl) SelectMode(value string) *TransferControl {
    a.Set("selectMode", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *TransferControl) Static(value string) *TransferControl {
    a.Set("static", value)
    return a
}

/**
 * 单个选项的高度，主要用于虚拟渲染
 */
func (a *TransferControl) ItemHeight(value string) *TransferControl {
    a.Set("itemHeight", value)
    return a
}

/**
 * 选项集合
 */
func (a *TransferControl) Options(value string) *TransferControl {
    a.Set("options", value)
    return a
}

/**
 * 是否为多选模式
 */
func (a *TransferControl) Multiple(value string) *TransferControl {
    a.Set("multiple", value)
    return a
}

/**
 * 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
 */
func (a *TransferControl) JoinValues(value string) *TransferControl {
    a.Set("joinValues", value)
    return a
}

/**
 * 是否只读
 */
func (a *TransferControl) ReadOnly(value string) *TransferControl {
    a.Set("readOnly", value)
    return a
}

/**
 */
func (a *TransferControl) Desc(value string) *TransferControl {
    a.Set("desc", value)
    return a
}

/**
 * 左侧的标题文字
 */
func (a *TransferControl) SelectTitle(value string) *TransferControl {
    a.Set("selectTitle", value)
    return a
}

/**
 * 用表达式来配置 source 接口初始要不要拉取
 */
func (a *TransferControl) InitFetchOn(value string) *TransferControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 是否为必填
 */
func (a *TransferControl) Required(value string) *TransferControl {
    a.Set("required", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *TransferControl) Hidden(value string) *TransferControl {
    a.Set("hidden", value)
    return a
}

/**
 * 是否默认都展开
 */
func (a *TransferControl) InitiallyOpen(value string) *TransferControl {
    a.Set("initiallyOpen", value)
    return a
}

/**
 * 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
 */
func (a *TransferControl) ExtractValue(value string) *TransferControl {
    a.Set("extractValue", value)
    return a
}

/**
 * 添加时调用的接口
 */
func (a *TransferControl) AddApi(value string) *TransferControl {
    a.Set("addApi", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *TransferControl) LabelWidth(value string) *TransferControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *TransferControl) Inline(value string) *TransferControl {
    a.Set("inline", value)
    return a
}

/**
 * 可搜索？
 */
func (a *TransferControl) Searchable(value string) *TransferControl {
    a.Set("searchable", value)
    return a
}

/**
 * 当在value值未匹配到当前options中的选项时，是否value值对应文本飘红显示
 */
func (a *TransferControl) ShowInvalidMatch(value string) *TransferControl {
    a.Set("showInvalidMatch", value)
    return a
}

/**
 * 是否可删除
 */
func (a *TransferControl) Removable(value string) *TransferControl {
    a.Set("removable", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *TransferControl) Mode(value string) *TransferControl {
    a.Set("mode", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *TransferControl) DisabledOn(value string) *TransferControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *TransferControl) StaticPlaceholder(value string) *TransferControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 在选项数量达到多少时开启虚拟渲染
 */
func (a *TransferControl) VirtualThreshold(value string) *TransferControl {
    a.Set("virtualThreshold", value)
    return a
}

/**
 * 分页配置，selectMode为默认和table才会生效
 */
func (a *TransferControl) Pagination(value string) *TransferControl {
    a.Set("pagination", value)
    return a
}

/**
 */
func (a *TransferControl) LoadingConfig(value string) *TransferControl {
    a.Set("loadingConfig", value)
    return a
}

/**
 */
func (a *TransferControl) InitAutoFill(value string) *TransferControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *TransferControl) ValidateOnChange(value string) *TransferControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *TransferControl) Width(value string) *TransferControl {
    a.Set("width", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *TransferControl) ExtraName(value string) *TransferControl {
    a.Set("extraName", value)
    return a
}

/**
 * 当 selectMode 为 associated 时用来定义右侧的选择模式
 * 可选值: table | list | tree | chained
 */
func (a *TransferControl) RightMode(value string) *TransferControl {
    a.Set("rightMode", value)
    return a
}

/**
 * 当 searchResultMode 为 table 时定义表格列信息。
 */
func (a *TransferControl) SearchResultColumns(value string) *TransferControl {
    a.Set("searchResultColumns", value)
    return a
}

/**
 * 右侧结果的标题文字
 */
func (a *TransferControl) ResultTitle(value string) *TransferControl {
    a.Set("resultTitle", value)
    return a
}

/**
 * 是否可以新增
 */
func (a *TransferControl) Creatable(value string) *TransferControl {
    a.Set("creatable", value)
    return a
}

/**
 * 配置 label className
 */
func (a *TransferControl) LabelClassName(value string) *TransferControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *TransferControl) Horizontal(value string) *TransferControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *TransferControl) EditorSetting(value string) *TransferControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *TransferControl) OnEvent(value string) *TransferControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 当 selectMode 为 associated 时用来定义左侧的选项
 */
func (a *TransferControl) LeftOptions(value string) *TransferControl {
    a.Set("leftOptions", value)
    return a
}

/**
 * 统计数字
 */
func (a *TransferControl) Statistics(value string) *TransferControl {
    a.Set("statistics", value)
    return a
}

/**
 * 只读条件
 */
func (a *TransferControl) ReadOnlyOn(value string) *TransferControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 右侧列表搜索框提示
 */
func (a *TransferControl) ResultSearchPlaceholder(value string) *TransferControl {
    a.Set("resultSearchPlaceholder", value)
    return a
}

/**
 * 点清除按钮时，将表单项设置成当前配置的值。
 */
func (a *TransferControl) ResetValue(value string) *TransferControl {
    a.Set("resetValue", value)
    return a
}

/**
 * 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
 */
func (a *TransferControl) DeferApi(value string) *TransferControl {
    a.Set("deferApi", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *TransferControl) Description(value string) *TransferControl {
    a.Set("description", value)
    return a
}

/**
 * 结果（右则）列表的检索功能，当设置为true时，可以通过输入检索模糊匹配检索内容
 */
func (a *TransferControl) ResultSearchable(value string) *TransferControl {
    a.Set("resultSearchable", value)
    return a
}

/**
 * 编辑时调用的 API
 */
func (a *TransferControl) EditApi(value string) *TransferControl {
    a.Set("editApi", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *TransferControl) StaticInputClassName(value string) *TransferControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *TransferControl) StaticOn(value string) *TransferControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 左侧列表搜索框提示
 */
func (a *TransferControl) SearchPlaceholder(value string) *TransferControl {
    a.Set("searchPlaceholder", value)
    return a
}

/**
 * 新增时的表单项。
 */
func (a *TransferControl) AddControls(value string) *TransferControl {
    a.Set("addControls", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *TransferControl) Size(value string) *TransferControl {
    a.Set("size", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *TransferControl) HiddenOn(value string) *TransferControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 搜索 API
 */
func (a *TransferControl) SearchApi(value string) *TransferControl {
    a.Set("searchApi", value)
    return a
}

/**
 * 是否显示剪头
 */
func (a *TransferControl) ShowArrow(value string) *TransferControl {
    a.Set("showArrow", value)
    return a
}

/**
 */
func (a *TransferControl) StaticSchema(value string) *TransferControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *TransferControl) Remark(value string) *TransferControl {
    a.Set("remark", value)
    return a
}

/**
 * 懒加载字段
 */
func (a *TransferControl) DeferField(value string) *TransferControl {
    a.Set("deferField", value)
    return a
}

/**
 * 选项修改的表单项
 */
func (a *TransferControl) EditControls(value string) *TransferControl {
    a.Set("editControls", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *TransferControl) AutoFill(value string) *TransferControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *TransferControl) Hint(value string) *TransferControl {
    a.Set("hint", value)
    return a
}

/**
 * 配置 input className
 */
func (a *TransferControl) InputClassName(value string) *TransferControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *TransferControl) Value(value string) *TransferControl {
    a.Set("value", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *TransferControl) VisibleOn(value string) *TransferControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 用来丰富值的展示
 */
func (a *TransferControl) ValueTpl(value string) *TransferControl {
    a.Set("valueTpl", value)
    return a
}

/**
 * ui级联关系，true代表级联选中，false代表不级联，默认为true
 */
func (a *TransferControl) AutoCheckChildren(value string) *TransferControl {
    a.Set("autoCheckChildren", value)
    return a
}

/**
 * 是否可清除。
 */
func (a *TransferControl) Clearable(value string) *TransferControl {
    a.Set("clearable", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *TransferControl) UseMobileUI(value string) *TransferControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 结果面板是否追踪显示
 */
func (a *TransferControl) ResultListModeFollowSelect(value string) *TransferControl {
    a.Set("resultListModeFollowSelect", value)
    return a
}

/**
 * 搜索结果展示模式
 * 可选值: table | list | tree | chained
 */
func (a *TransferControl) SearchResultMode(value string) *TransferControl {
    a.Set("searchResultMode", value)
    return a
}

/**
 * 用来丰富选项展示
 */
func (a *TransferControl) MenuTpl(value string) *TransferControl {
    a.Set("menuTpl", value)
    return a
}

/**
 * 配置 source 接口初始拉不拉取。
 */
func (a *TransferControl) InitFetch(value string) *TransferControl {
    a.Set("initFetch", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *TransferControl) ClearValueOnHidden(value string) *TransferControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *TransferControl) StaticLabelClassName(value string) *TransferControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 多选模式，值太多时是否避免折行
 */
func (a *TransferControl) ValuesNoWrap(value string) *TransferControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * 新增文字
 */
func (a *TransferControl) CreateBtnLabel(value string) *TransferControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *TransferControl) ValidationErrors(value string) *TransferControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 描述标题
 */
func (a *TransferControl) Label(value string) *TransferControl {
    a.Set("label", value)
    return a
}

/**
 * 描述标题
 */
func (a *TransferControl) LabelAlign(value string) *TransferControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *TransferControl) DescriptionClassName(value string) *TransferControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 是否显示
 */
func (a *TransferControl) Visible(value string) *TransferControl {
    a.Set("visible", value)
    return a
}

/**
 * 控制新增弹框设置项
 */
func (a *TransferControl) AddDialog(value string) *TransferControl {
    a.Set("addDialog", value)
    return a
}

/**
 */
func (a *TransferControl) Validations(value string) *TransferControl {
    a.Set("validations", value)
    return a
}

/**
 * 是否禁用
 */
func (a *TransferControl) Disabled(value string) *TransferControl {
    a.Set("disabled", value)
    return a
}

/**
 * 选项删除提示文字。
 */
func (a *TransferControl) DeleteConfirmText(value string) *TransferControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *TransferControl) StaticClassName(value string) *TransferControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 可用来通过 API 拉取 options。
 */
func (a *TransferControl) Source(value string) *TransferControl {
    a.Set("source", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *TransferControl) SubmitOnChange(value string) *TransferControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *TransferControl) ValidateApi(value string) *TransferControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 表单项类型
 */
func (a *TransferControl) Type(value string) *TransferControl {
    a.Set("type", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *TransferControl) ClassName(value string) *TransferControl {
    a.Set("className", value)
    return a
}

/**
 * 当 selectMode 为 associated 时用来定义左侧的选择模式
 * 可选值: tree | list
 */
func (a *TransferControl) LeftMode(value string) *TransferControl {
    a.Set("leftMode", value)
    return a
}

/**
 * 是否可以编辑
 */
func (a *TransferControl) Editable(value string) *TransferControl {
    a.Set("editable", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *TransferControl) Name(value string) *TransferControl {
    a.Set("name", value)
    return a
}

/**
 * 当 selectMode 为 table 时定义表格列信息。
 */
func (a *TransferControl) Columns(value string) *TransferControl {
    a.Set("columns", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *TransferControl) LabelRemark(value string) *TransferControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 默认选择选项第一个值。
 */
func (a *TransferControl) SelectFirst(value string) *TransferControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * 分割符
 */
func (a *TransferControl) Delimiter(value string) *TransferControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 选项删除 API
 */
func (a *TransferControl) DeleteApi(value string) *TransferControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *TransferControl) Id(value string) *TransferControl {
    a.Set("id", value)
    return a
}

/**
 * 组件样式
 */
func (a *TransferControl) Style(value string) *TransferControl {
    a.Set("style", value)
    return a
}

/**
 * 可排序？
 */
func (a *TransferControl) Sortable(value string) *TransferControl {
    a.Set("sortable", value)
    return a
}

/**
 * 树形模式下，仅选中子节点
 */
func (a *TransferControl) OnlyChildren(value string) *TransferControl {
    a.Set("onlyChildren", value)
    return a
}
