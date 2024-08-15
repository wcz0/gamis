package renderers


/**
 * TabsTransferPicker 穿梭器的弹框形态 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/tabs-transfer-picker

 * @author wcz0
 * @version 6.2.2
 */
type TabsTransferPickerControl struct {
	*BaseRenderer
}

func NewTabsTransferPickerControl() *TabsTransferPickerControl {
    a := &TabsTransferPickerControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "tabs-transfer-picker")
    return a
}


func (a *TabsTransferPickerControl) Set(name string, value interface{}) *TabsTransferPickerControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 控制新增弹框设置项
 */
func (a *TabsTransferPickerControl) AddDialog(value interface{}) *TabsTransferPickerControl {
    a.Set("addDialog", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *TabsTransferPickerControl) ValidationErrors(value interface{}) *TabsTransferPickerControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *TabsTransferPickerControl) StaticClassName(value interface{}) *TabsTransferPickerControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 是否显示剪头
 */
func (a *TabsTransferPickerControl) ShowArrow(value interface{}) *TabsTransferPickerControl {
    a.Set("showArrow", value)
    return a
}

/**
 * 结果（右则）列表的检索功能，当设置为true时，可以通过输入检索模糊匹配检索内容
 */
func (a *TabsTransferPickerControl) ResultSearchable(value interface{}) *TabsTransferPickerControl {
    a.Set("resultSearchable", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *TabsTransferPickerControl) StaticOn(value interface{}) *TabsTransferPickerControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 是否可删除
 */
func (a *TabsTransferPickerControl) Removable(value interface{}) *TabsTransferPickerControl {
    a.Set("removable", value)
    return a
}

/**
 * 多选模式，值太多时是否避免折行
 */
func (a *TabsTransferPickerControl) ValuesNoWrap(value interface{}) *TabsTransferPickerControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * 懒加载字段
 */
func (a *TabsTransferPickerControl) DeferField(value interface{}) *TabsTransferPickerControl {
    a.Set("deferField", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *TabsTransferPickerControl) Description(value interface{}) *TabsTransferPickerControl {
    a.Set("description", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *TabsTransferPickerControl) Horizontal(value interface{}) *TabsTransferPickerControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 描述标题
 */
func (a *TabsTransferPickerControl) Label(value interface{}) *TabsTransferPickerControl {
    a.Set("label", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *TabsTransferPickerControl) HiddenOn(value interface{}) *TabsTransferPickerControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 可用来通过 API 拉取 options。
 */
func (a *TabsTransferPickerControl) Source(value interface{}) *TabsTransferPickerControl {
    a.Set("source", value)
    return a
}

/**
 */
func (a *TabsTransferPickerControl) Desc(value interface{}) *TabsTransferPickerControl {
    a.Set("desc", value)
    return a
}

/**
 * 新增文字
 */
func (a *TabsTransferPickerControl) CreateBtnLabel(value interface{}) *TabsTransferPickerControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *TabsTransferPickerControl) Hidden(value interface{}) *TabsTransferPickerControl {
    a.Set("hidden", value)
    return a
}

/**
 * 分页配置，selectMode为默认和table才会生效
 */
func (a *TabsTransferPickerControl) Pagination(value interface{}) *TabsTransferPickerControl {
    a.Set("pagination", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *TabsTransferPickerControl) Hint(value interface{}) *TabsTransferPickerControl {
    a.Set("hint", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *TabsTransferPickerControl) Value(value interface{}) *TabsTransferPickerControl {
    a.Set("value", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *TabsTransferPickerControl) StaticLabelClassName(value interface{}) *TabsTransferPickerControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 点清除按钮时，将表单项设置成当前配置的值。
 */
func (a *TabsTransferPickerControl) ResetValue(value interface{}) *TabsTransferPickerControl {
    a.Set("resetValue", value)
    return a
}

/**
 */
func (a *TabsTransferPickerControl) Type(value interface{}) *TabsTransferPickerControl {
    a.Set("type", value)
    return a
}

/**
 * 是否禁用
 */
func (a *TabsTransferPickerControl) Disabled(value interface{}) *TabsTransferPickerControl {
    a.Set("disabled", value)
    return a
}

/**
 * 配置 source 接口初始拉不拉取。
 */
func (a *TabsTransferPickerControl) InitFetch(value interface{}) *TabsTransferPickerControl {
    a.Set("initFetch", value)
    return a
}

/**
 * 只读条件
 */
func (a *TabsTransferPickerControl) ReadOnlyOn(value interface{}) *TabsTransferPickerControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 组件样式
 */
func (a *TabsTransferPickerControl) Style(value interface{}) *TabsTransferPickerControl {
    a.Set("style", value)
    return a
}

/**
 * 用表达式来配置 source 接口初始要不要拉取
 */
func (a *TabsTransferPickerControl) InitFetchOn(value interface{}) *TabsTransferPickerControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *TabsTransferPickerControl) DisabledOn(value interface{}) *TabsTransferPickerControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *TabsTransferPickerControl) StaticPlaceholder(value interface{}) *TabsTransferPickerControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 可排序？
 */
func (a *TabsTransferPickerControl) Sortable(value interface{}) *TabsTransferPickerControl {
    a.Set("sortable", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *TabsTransferPickerControl) LabelRemark(value interface{}) *TabsTransferPickerControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *TabsTransferPickerControl) DescriptionClassName(value interface{}) *TabsTransferPickerControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 */
func (a *TabsTransferPickerControl) Validations(value interface{}) *TabsTransferPickerControl {
    a.Set("validations", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *TabsTransferPickerControl) ValidateApi(value interface{}) *TabsTransferPickerControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 选项修改的表单项
 */
func (a *TabsTransferPickerControl) EditControls(value interface{}) *TabsTransferPickerControl {
    a.Set("editControls", value)
    return a
}

/**
 * 选项删除 API
 */
func (a *TabsTransferPickerControl) DeleteApi(value interface{}) *TabsTransferPickerControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * 结果面板是否追踪显示
 */
func (a *TabsTransferPickerControl) ResultListModeFollowSelect(value interface{}) *TabsTransferPickerControl {
    a.Set("resultListModeFollowSelect", value)
    return a
}

/**
 * 右侧结果的标题文字
 */
func (a *TabsTransferPickerControl) ResultTitle(value interface{}) *TabsTransferPickerControl {
    a.Set("resultTitle", value)
    return a
}

/**
 * 是否显示
 */
func (a *TabsTransferPickerControl) Visible(value interface{}) *TabsTransferPickerControl {
    a.Set("visible", value)
    return a
}

/**
 * 是否可以新增
 */
func (a *TabsTransferPickerControl) Creatable(value interface{}) *TabsTransferPickerControl {
    a.Set("creatable", value)
    return a
}

/**
 * ui级联关系，true代表级联选中，false代表不级联，默认为true
 */
func (a *TabsTransferPickerControl) AutoCheckChildren(value interface{}) *TabsTransferPickerControl {
    a.Set("autoCheckChildren", value)
    return a
}

/**
 */
func (a *TabsTransferPickerControl) TestIdBuilder(value interface{}) *TabsTransferPickerControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 默认选择选项第一个值。
 */
func (a *TabsTransferPickerControl) SelectFirst(value interface{}) *TabsTransferPickerControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * 搜索结果展示模式
 * 可选值: table | list | tree | chained
 */
func (a *TabsTransferPickerControl) SearchResultMode(value interface{}) *TabsTransferPickerControl {
    a.Set("searchResultMode", value)
    return a
}

/**
 * 占位符
 */
func (a *TabsTransferPickerControl) Placeholder(value interface{}) *TabsTransferPickerControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *TabsTransferPickerControl) Id(value interface{}) *TabsTransferPickerControl {
    a.Set("id", value)
    return a
}

/**
 * 选项集合
 */
func (a *TabsTransferPickerControl) Options(value interface{}) *TabsTransferPickerControl {
    a.Set("options", value)
    return a
}

/**
 * 是否为多选模式
 */
func (a *TabsTransferPickerControl) Multiple(value interface{}) *TabsTransferPickerControl {
    a.Set("multiple", value)
    return a
}

/**
 * 用来丰富选项展示
 */
func (a *TabsTransferPickerControl) MenuTpl(value interface{}) *TabsTransferPickerControl {
    a.Set("menuTpl", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *TabsTransferPickerControl) Static(value interface{}) *TabsTransferPickerControl {
    a.Set("static", value)
    return a
}

/**
 * 新增时的表单项。
 */
func (a *TabsTransferPickerControl) AddControls(value interface{}) *TabsTransferPickerControl {
    a.Set("addControls", value)
    return a
}

/**
 * 当 selectMode 为 table 时定义表格列信息。
 */
func (a *TabsTransferPickerControl) Columns(value interface{}) *TabsTransferPickerControl {
    a.Set("columns", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *TabsTransferPickerControl) Size(value interface{}) *TabsTransferPickerControl {
    a.Set("size", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *TabsTransferPickerControl) ClearValueOnHidden(value interface{}) *TabsTransferPickerControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 */
func (a *TabsTransferPickerControl) InitAutoFill(value interface{}) *TabsTransferPickerControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 是否默认都展开
 */
func (a *TabsTransferPickerControl) InitiallyOpen(value interface{}) *TabsTransferPickerControl {
    a.Set("initiallyOpen", value)
    return a
}

/**
 * 配置 input className
 */
func (a *TabsTransferPickerControl) InputClassName(value interface{}) *TabsTransferPickerControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 是否为必填
 */
func (a *TabsTransferPickerControl) Required(value interface{}) *TabsTransferPickerControl {
    a.Set("required", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *TabsTransferPickerControl) UseMobileUI(value interface{}) *TabsTransferPickerControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
 */
func (a *TabsTransferPickerControl) JoinValues(value interface{}) *TabsTransferPickerControl {
    a.Set("joinValues", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *TabsTransferPickerControl) Remark(value interface{}) *TabsTransferPickerControl {
    a.Set("remark", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *TabsTransferPickerControl) SubmitOnChange(value interface{}) *TabsTransferPickerControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *TabsTransferPickerControl) Mode(value interface{}) *TabsTransferPickerControl {
    a.Set("mode", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *TabsTransferPickerControl) ClassName(value interface{}) *TabsTransferPickerControl {
    a.Set("className", value)
    return a
}

/**
 * 当 searchResultMode 为 table 时定义表格列信息。
 */
func (a *TabsTransferPickerControl) SearchResultColumns(value interface{}) *TabsTransferPickerControl {
    a.Set("searchResultColumns", value)
    return a
}

/**
 * 用来丰富值的展示
 */
func (a *TabsTransferPickerControl) ValueTpl(value interface{}) *TabsTransferPickerControl {
    a.Set("valueTpl", value)
    return a
}

/**
 * 树形模式下，仅选中子节点
 */
func (a *TabsTransferPickerControl) OnlyChildren(value interface{}) *TabsTransferPickerControl {
    a.Set("onlyChildren", value)
    return a
}

/**
 * 统计数字
 */
func (a *TabsTransferPickerControl) Statistics(value interface{}) *TabsTransferPickerControl {
    a.Set("statistics", value)
    return a
}

/**
 * 在选项数量达到多少时开启虚拟渲染
 */
func (a *TabsTransferPickerControl) VirtualThreshold(value interface{}) *TabsTransferPickerControl {
    a.Set("virtualThreshold", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *TabsTransferPickerControl) LabelWidth(value interface{}) *TabsTransferPickerControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *TabsTransferPickerControl) VisibleOn(value interface{}) *TabsTransferPickerControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否只读
 */
func (a *TabsTransferPickerControl) ReadOnly(value interface{}) *TabsTransferPickerControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *TabsTransferPickerControl) StaticInputClassName(value interface{}) *TabsTransferPickerControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
 */
func (a *TabsTransferPickerControl) ExtractValue(value interface{}) *TabsTransferPickerControl {
    a.Set("extractValue", value)
    return a
}

/**
 */
func (a *TabsTransferPickerControl) LoadingConfig(value interface{}) *TabsTransferPickerControl {
    a.Set("loadingConfig", value)
    return a
}

/**
 * 描述标题
 */
func (a *TabsTransferPickerControl) LabelAlign(value interface{}) *TabsTransferPickerControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 是否可以编辑
 */
func (a *TabsTransferPickerControl) Editable(value interface{}) *TabsTransferPickerControl {
    a.Set("editable", value)
    return a
}

/**
 * 右侧列表搜索框提示
 */
func (a *TabsTransferPickerControl) ResultSearchPlaceholder(value interface{}) *TabsTransferPickerControl {
    a.Set("resultSearchPlaceholder", value)
    return a
}

/**
 * 搜索 API
 */
func (a *TabsTransferPickerControl) SearchApi(value interface{}) *TabsTransferPickerControl {
    a.Set("searchApi", value)
    return a
}

/**
 */
func (a *TabsTransferPickerControl) StaticSchema(value interface{}) *TabsTransferPickerControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 选项删除提示文字。
 */
func (a *TabsTransferPickerControl) DeleteConfirmText(value interface{}) *TabsTransferPickerControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * 当 selectMode 为 associated 时用来定义左侧的选择模式
 * 可选值: tree | list
 */
func (a *TabsTransferPickerControl) LeftMode(value interface{}) *TabsTransferPickerControl {
    a.Set("leftMode", value)
    return a
}

/**
 * 当在value值未匹配到当前options中的选项时，是否value值对应文本飘红显示
 */
func (a *TabsTransferPickerControl) ShowInvalidMatch(value interface{}) *TabsTransferPickerControl {
    a.Set("showInvalidMatch", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *TabsTransferPickerControl) ExtraName(value interface{}) *TabsTransferPickerControl {
    a.Set("extraName", value)
    return a
}

/**
 * 控制编辑弹框设置项
 */
func (a *TabsTransferPickerControl) EditDialog(value interface{}) *TabsTransferPickerControl {
    a.Set("editDialog", value)
    return a
}

/**
 * 是否可清除。
 */
func (a *TabsTransferPickerControl) Clearable(value interface{}) *TabsTransferPickerControl {
    a.Set("clearable", value)
    return a
}

/**
 * 添加时调用的接口
 */
func (a *TabsTransferPickerControl) AddApi(value interface{}) *TabsTransferPickerControl {
    a.Set("addApi", value)
    return a
}

/**
 * 单个选项的高度，主要用于虚拟渲染
 */
func (a *TabsTransferPickerControl) ItemHeight(value interface{}) *TabsTransferPickerControl {
    a.Set("itemHeight", value)
    return a
}

/**
 * 配置 label className
 */
func (a *TabsTransferPickerControl) LabelClassName(value interface{}) *TabsTransferPickerControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *TabsTransferPickerControl) Inline(value interface{}) *TabsTransferPickerControl {
    a.Set("inline", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *TabsTransferPickerControl) OnEvent(value interface{}) *TabsTransferPickerControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 可搜索？
 */
func (a *TabsTransferPickerControl) Searchable(value interface{}) *TabsTransferPickerControl {
    a.Set("searchable", value)
    return a
}

/**
 * 左侧的标题文字
 */
func (a *TabsTransferPickerControl) SelectTitle(value interface{}) *TabsTransferPickerControl {
    a.Set("selectTitle", value)
    return a
}

/**
 */
func (a *TabsTransferPickerControl) Row(value interface{}) *TabsTransferPickerControl {
    a.Set("row", value)
    return a
}

/**
 * 分割符
 */
func (a *TabsTransferPickerControl) Delimiter(value interface{}) *TabsTransferPickerControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
 */
func (a *TabsTransferPickerControl) DeferApi(value interface{}) *TabsTransferPickerControl {
    a.Set("deferApi", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *TabsTransferPickerControl) Width(value interface{}) *TabsTransferPickerControl {
    a.Set("width", value)
    return a
}

/**
 * 当 selectMode 为 associated 时用来定义左侧的选项
 */
func (a *TabsTransferPickerControl) LeftOptions(value interface{}) *TabsTransferPickerControl {
    a.Set("leftOptions", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *TabsTransferPickerControl) Name(value interface{}) *TabsTransferPickerControl {
    a.Set("name", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *TabsTransferPickerControl) ValidateOnChange(value interface{}) *TabsTransferPickerControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *TabsTransferPickerControl) AutoFill(value interface{}) *TabsTransferPickerControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 左侧列表搜索框提示
 */
func (a *TabsTransferPickerControl) SearchPlaceholder(value interface{}) *TabsTransferPickerControl {
    a.Set("searchPlaceholder", value)
    return a
}

/**
 * 编辑时调用的 API
 */
func (a *TabsTransferPickerControl) EditApi(value interface{}) *TabsTransferPickerControl {
    a.Set("editApi", value)
    return a
}

/**
 * 勾选展示模式
 * 可选值: table | list | tree | chained | associated
 */
func (a *TabsTransferPickerControl) SelectMode(value interface{}) *TabsTransferPickerControl {
    a.Set("selectMode", value)
    return a
}

/**
 * 当 selectMode 为 associated 时用来定义右侧的选择模式
 * 可选值: table | list | tree | chained
 */
func (a *TabsTransferPickerControl) RightMode(value interface{}) *TabsTransferPickerControl {
    a.Set("rightMode", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *TabsTransferPickerControl) EditorSetting(value interface{}) *TabsTransferPickerControl {
    a.Set("editorSetting", value)
    return a
}
