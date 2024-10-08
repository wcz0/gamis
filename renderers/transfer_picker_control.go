package renderers


/**
 * TransferPicker 穿梭器的弹框形态 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/transfer-picker

 * @author wcz0
 * @version 6.2.2
 */
type TransferPickerControl struct {
	*BaseRenderer
}

func NewTransferPickerControl() *TransferPickerControl {
    a := &TransferPickerControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "transfer-picker")
    return a
}


func (a *TransferPickerControl) Set(name string, value interface{}) *TransferPickerControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 是否显示表达式
 */
func (a *TransferPickerControl) VisibleOn(value interface{}) *TransferPickerControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *TransferPickerControl) OnEvent(value interface{}) *TransferPickerControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 编辑时调用的 API
 */
func (a *TransferPickerControl) EditApi(value interface{}) *TransferPickerControl {
    a.Set("editApi", value)
    return a
}

/**
 * 边框模式，全边框，还是半边框，或者没边框。
 * 可选值: full | half | none
 */
func (a *TransferPickerControl) BorderMode(value interface{}) *TransferPickerControl {
    a.Set("borderMode", value)
    return a
}

/**
 * 是否显示剪头
 */
func (a *TransferPickerControl) ShowArrow(value interface{}) *TransferPickerControl {
    a.Set("showArrow", value)
    return a
}

/**
 * 左侧的标题文字
 */
func (a *TransferPickerControl) SelectTitle(value interface{}) *TransferPickerControl {
    a.Set("selectTitle", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *TransferPickerControl) Hint(value interface{}) *TransferPickerControl {
    a.Set("hint", value)
    return a
}

/**
 * 选项删除提示文字。
 */
func (a *TransferPickerControl) DeleteConfirmText(value interface{}) *TransferPickerControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * 勾选展示模式
 * 可选值: table | list | tree | chained | associated
 */
func (a *TransferPickerControl) SelectMode(value interface{}) *TransferPickerControl {
    a.Set("selectMode", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *TransferPickerControl) ValidationErrors(value interface{}) *TransferPickerControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 分割符
 */
func (a *TransferPickerControl) Delimiter(value interface{}) *TransferPickerControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *TransferPickerControl) Inline(value interface{}) *TransferPickerControl {
    a.Set("inline", value)
    return a
}

/**
 * 选项删除 API
 */
func (a *TransferPickerControl) DeleteApi(value interface{}) *TransferPickerControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * 可排序？
 */
func (a *TransferPickerControl) Sortable(value interface{}) *TransferPickerControl {
    a.Set("sortable", value)
    return a
}

/**
 * 搜索结果展示模式
 * 可选值: table | list | tree | chained
 */
func (a *TransferPickerControl) SearchResultMode(value interface{}) *TransferPickerControl {
    a.Set("searchResultMode", value)
    return a
}

/**
 * 弹窗大小
 * 可选值: xs | sm | md | lg | xl | full
 */
func (a *TransferPickerControl) PickerSize(value interface{}) *TransferPickerControl {
    a.Set("pickerSize", value)
    return a
}

/**
 * 懒加载字段
 */
func (a *TransferPickerControl) DeferField(value interface{}) *TransferPickerControl {
    a.Set("deferField", value)
    return a
}

/**
 * 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
 */
func (a *TransferPickerControl) DeferApi(value interface{}) *TransferPickerControl {
    a.Set("deferApi", value)
    return a
}

/**
 * 配置 label className
 */
func (a *TransferPickerControl) LabelClassName(value interface{}) *TransferPickerControl {
    a.Set("labelClassName", value)
    return a
}

/**
 */
func (a *TransferPickerControl) InitAutoFill(value interface{}) *TransferPickerControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *TransferPickerControl) UseMobileUI(value interface{}) *TransferPickerControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 分页配置，selectMode为默认和table才会生效
 */
func (a *TransferPickerControl) Pagination(value interface{}) *TransferPickerControl {
    a.Set("pagination", value)
    return a
}

/**
 * 是否为必填
 */
func (a *TransferPickerControl) Required(value interface{}) *TransferPickerControl {
    a.Set("required", value)
    return a
}

/**
 * 描述标题
 */
func (a *TransferPickerControl) Label(value interface{}) *TransferPickerControl {
    a.Set("label", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *TransferPickerControl) DescriptionClassName(value interface{}) *TransferPickerControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 占位符
 */
func (a *TransferPickerControl) Placeholder(value interface{}) *TransferPickerControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否为多选模式
 */
func (a *TransferPickerControl) Multiple(value interface{}) *TransferPickerControl {
    a.Set("multiple", value)
    return a
}

/**
 * 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
 */
func (a *TransferPickerControl) ExtractValue(value interface{}) *TransferPickerControl {
    a.Set("extractValue", value)
    return a
}

/**
 * 当 selectMode 为 associated 时用来定义左侧的选择模式
 * 可选值: tree | list
 */
func (a *TransferPickerControl) LeftMode(value interface{}) *TransferPickerControl {
    a.Set("leftMode", value)
    return a
}

/**
 * 统计数字
 */
func (a *TransferPickerControl) Statistics(value interface{}) *TransferPickerControl {
    a.Set("statistics", value)
    return a
}

/**
 * 在选项数量达到多少时开启虚拟渲染
 */
func (a *TransferPickerControl) VirtualThreshold(value interface{}) *TransferPickerControl {
    a.Set("virtualThreshold", value)
    return a
}

/**
 * 是否禁用
 */
func (a *TransferPickerControl) Disabled(value interface{}) *TransferPickerControl {
    a.Set("disabled", value)
    return a
}

/**
 * 点清除按钮时，将表单项设置成当前配置的值。
 */
func (a *TransferPickerControl) ResetValue(value interface{}) *TransferPickerControl {
    a.Set("resetValue", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *TransferPickerControl) LabelRemark(value interface{}) *TransferPickerControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *TransferPickerControl) StaticInputClassName(value interface{}) *TransferPickerControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 可用来通过 API 拉取 options。
 */
func (a *TransferPickerControl) Source(value interface{}) *TransferPickerControl {
    a.Set("source", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *TransferPickerControl) AutoFill(value interface{}) *TransferPickerControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *TransferPickerControl) ClassName(value interface{}) *TransferPickerControl {
    a.Set("className", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *TransferPickerControl) StaticPlaceholder(value interface{}) *TransferPickerControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 添加时调用的接口
 */
func (a *TransferPickerControl) AddApi(value interface{}) *TransferPickerControl {
    a.Set("addApi", value)
    return a
}

/**
 * 只读条件
 */
func (a *TransferPickerControl) ReadOnlyOn(value interface{}) *TransferPickerControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 */
func (a *TransferPickerControl) LoadingConfig(value interface{}) *TransferPickerControl {
    a.Set("loadingConfig", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *TransferPickerControl) HiddenOn(value interface{}) *TransferPickerControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 右侧列表搜索框提示
 */
func (a *TransferPickerControl) ResultSearchPlaceholder(value interface{}) *TransferPickerControl {
    a.Set("resultSearchPlaceholder", value)
    return a
}

/**
 * 是否默认都展开
 */
func (a *TransferPickerControl) InitiallyOpen(value interface{}) *TransferPickerControl {
    a.Set("initiallyOpen", value)
    return a
}

/**
 * 选项修改的表单项
 */
func (a *TransferPickerControl) EditControls(value interface{}) *TransferPickerControl {
    a.Set("editControls", value)
    return a
}

/**
 * 当 selectMode 为 associated 时用来定义左侧的选项
 */
func (a *TransferPickerControl) LeftOptions(value interface{}) *TransferPickerControl {
    a.Set("leftOptions", value)
    return a
}

/**
 * 用来丰富选项展示
 */
func (a *TransferPickerControl) MenuTpl(value interface{}) *TransferPickerControl {
    a.Set("menuTpl", value)
    return a
}

/**
 * 配置 source 接口初始拉不拉取。
 */
func (a *TransferPickerControl) InitFetch(value interface{}) *TransferPickerControl {
    a.Set("initFetch", value)
    return a
}

/**
 * 新增文字
 */
func (a *TransferPickerControl) CreateBtnLabel(value interface{}) *TransferPickerControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * 是否可以编辑
 */
func (a *TransferPickerControl) Editable(value interface{}) *TransferPickerControl {
    a.Set("editable", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *TransferPickerControl) Remark(value interface{}) *TransferPickerControl {
    a.Set("remark", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *TransferPickerControl) ValidateOnChange(value interface{}) *TransferPickerControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 是否可清除。
 */
func (a *TransferPickerControl) Clearable(value interface{}) *TransferPickerControl {
    a.Set("clearable", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *TransferPickerControl) SubmitOnChange(value interface{}) *TransferPickerControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 结果（右则）列表的检索功能，当设置为true时，可以通过输入检索模糊匹配检索内容
 */
func (a *TransferPickerControl) ResultSearchable(value interface{}) *TransferPickerControl {
    a.Set("resultSearchable", value)
    return a
}

/**
 * 树形模式下，仅选中子节点
 */
func (a *TransferPickerControl) OnlyChildren(value interface{}) *TransferPickerControl {
    a.Set("onlyChildren", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *TransferPickerControl) Size(value interface{}) *TransferPickerControl {
    a.Set("size", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *TransferPickerControl) LabelWidth(value interface{}) *TransferPickerControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 当 searchResultMode 为 table 时定义表格列信息。
 */
func (a *TransferPickerControl) SearchResultColumns(value interface{}) *TransferPickerControl {
    a.Set("searchResultColumns", value)
    return a
}

/**
 * 单个选项的高度，主要用于虚拟渲染
 */
func (a *TransferPickerControl) ItemHeight(value interface{}) *TransferPickerControl {
    a.Set("itemHeight", value)
    return a
}

/**
 * 当在value值未匹配到当前options中的选项时，是否value值对应文本飘红显示
 */
func (a *TransferPickerControl) ShowInvalidMatch(value interface{}) *TransferPickerControl {
    a.Set("showInvalidMatch", value)
    return a
}

/**
 * 右侧结果的标题文字
 */
func (a *TransferPickerControl) ResultTitle(value interface{}) *TransferPickerControl {
    a.Set("resultTitle", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *TransferPickerControl) ExtraName(value interface{}) *TransferPickerControl {
    a.Set("extraName", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *TransferPickerControl) ValidateApi(value interface{}) *TransferPickerControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 选项集合
 */
func (a *TransferPickerControl) Options(value interface{}) *TransferPickerControl {
    a.Set("options", value)
    return a
}

/**
 * 用表达式来配置 source 接口初始要不要拉取
 */
func (a *TransferPickerControl) InitFetchOn(value interface{}) *TransferPickerControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 控制新增弹框设置项
 */
func (a *TransferPickerControl) AddDialog(value interface{}) *TransferPickerControl {
    a.Set("addDialog", value)
    return a
}

/**
 * 搜索 API
 */
func (a *TransferPickerControl) SearchApi(value interface{}) *TransferPickerControl {
    a.Set("searchApi", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *TransferPickerControl) Value(value interface{}) *TransferPickerControl {
    a.Set("value", value)
    return a
}

/**
 * 组件样式
 */
func (a *TransferPickerControl) Style(value interface{}) *TransferPickerControl {
    a.Set("style", value)
    return a
}

/**
 * 当 selectMode 为 associated 时用来定义右侧的选择模式
 * 可选值: table | list | tree | chained
 */
func (a *TransferPickerControl) RightMode(value interface{}) *TransferPickerControl {
    a.Set("rightMode", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *TransferPickerControl) StaticClassName(value interface{}) *TransferPickerControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 是否可删除
 */
func (a *TransferPickerControl) Removable(value interface{}) *TransferPickerControl {
    a.Set("removable", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *TransferPickerControl) Name(value interface{}) *TransferPickerControl {
    a.Set("name", value)
    return a
}

/**
 * 用来丰富值的展示
 */
func (a *TransferPickerControl) ValueTpl(value interface{}) *TransferPickerControl {
    a.Set("valueTpl", value)
    return a
}

/**
 */
func (a *TransferPickerControl) TestIdBuilder(value interface{}) *TransferPickerControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *TransferPickerControl) Id(value interface{}) *TransferPickerControl {
    a.Set("id", value)
    return a
}

/**
 * 结果面板是否追踪显示
 */
func (a *TransferPickerControl) ResultListModeFollowSelect(value interface{}) *TransferPickerControl {
    a.Set("resultListModeFollowSelect", value)
    return a
}

/**
 * 左侧列表搜索框提示
 */
func (a *TransferPickerControl) SearchPlaceholder(value interface{}) *TransferPickerControl {
    a.Set("searchPlaceholder", value)
    return a
}

/**
 * 是否只读
 */
func (a *TransferPickerControl) ReadOnly(value interface{}) *TransferPickerControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *TransferPickerControl) StaticLabelClassName(value interface{}) *TransferPickerControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 是否显示
 */
func (a *TransferPickerControl) Visible(value interface{}) *TransferPickerControl {
    a.Set("visible", value)
    return a
}

/**
 * 控制编辑弹框设置项
 */
func (a *TransferPickerControl) EditDialog(value interface{}) *TransferPickerControl {
    a.Set("editDialog", value)
    return a
}

/**
 */
func (a *TransferPickerControl) Desc(value interface{}) *TransferPickerControl {
    a.Set("desc", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *TransferPickerControl) Horizontal(value interface{}) *TransferPickerControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *TransferPickerControl) Hidden(value interface{}) *TransferPickerControl {
    a.Set("hidden", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *TransferPickerControl) Static(value interface{}) *TransferPickerControl {
    a.Set("static", value)
    return a
}

/**
 * 描述标题
 */
func (a *TransferPickerControl) LabelAlign(value interface{}) *TransferPickerControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 配置 input className
 */
func (a *TransferPickerControl) InputClassName(value interface{}) *TransferPickerControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 是否可以新增
 */
func (a *TransferPickerControl) Creatable(value interface{}) *TransferPickerControl {
    a.Set("creatable", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *TransferPickerControl) DisabledOn(value interface{}) *TransferPickerControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 默认选择选项第一个值。
 */
func (a *TransferPickerControl) SelectFirst(value interface{}) *TransferPickerControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *TransferPickerControl) Width(value interface{}) *TransferPickerControl {
    a.Set("width", value)
    return a
}

/**
 */
func (a *TransferPickerControl) Row(value interface{}) *TransferPickerControl {
    a.Set("row", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *TransferPickerControl) EditorSetting(value interface{}) *TransferPickerControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
 */
func (a *TransferPickerControl) JoinValues(value interface{}) *TransferPickerControl {
    a.Set("joinValues", value)
    return a
}

/**
 * 新增时的表单项。
 */
func (a *TransferPickerControl) AddControls(value interface{}) *TransferPickerControl {
    a.Set("addControls", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *TransferPickerControl) Mode(value interface{}) *TransferPickerControl {
    a.Set("mode", value)
    return a
}

/**
 */
func (a *TransferPickerControl) Validations(value interface{}) *TransferPickerControl {
    a.Set("validations", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *TransferPickerControl) ClearValueOnHidden(value interface{}) *TransferPickerControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 */
func (a *TransferPickerControl) Type(value interface{}) *TransferPickerControl {
    a.Set("type", value)
    return a
}

/**
 * 可搜索？
 */
func (a *TransferPickerControl) Searchable(value interface{}) *TransferPickerControl {
    a.Set("searchable", value)
    return a
}

/**
 * ui级联关系，true代表级联选中，false代表不级联，默认为true
 */
func (a *TransferPickerControl) AutoCheckChildren(value interface{}) *TransferPickerControl {
    a.Set("autoCheckChildren", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *TransferPickerControl) Description(value interface{}) *TransferPickerControl {
    a.Set("description", value)
    return a
}

/**
 * 多选模式，值太多时是否避免折行
 */
func (a *TransferPickerControl) ValuesNoWrap(value interface{}) *TransferPickerControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * 当 selectMode 为 table 时定义表格列信息。
 */
func (a *TransferPickerControl) Columns(value interface{}) *TransferPickerControl {
    a.Set("columns", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *TransferPickerControl) StaticOn(value interface{}) *TransferPickerControl {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *TransferPickerControl) StaticSchema(value interface{}) *TransferPickerControl {
    a.Set("staticSchema", value)
    return a
}
