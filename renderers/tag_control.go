package renderers


/**
 * Tag 输入框 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/tag
 *

*/
type TagControl struct {
	*BaseRenderer
}

func NewTagControl() *TagControl {
    a := &TagControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-tag")
    return a
}
/**
 * 分割符
 */
func (a *TagControl) Delimiter(value string) *TagControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *TagControl) Value(value string) *TagControl {
    a.Set("value", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *TagControl) DisabledOn(value string) *TagControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 选项修改的表单项
 */
func (a *TagControl) EditControls(value string) *TagControl {
    a.Set("editControls", value)
    return a
}

/**
 * 配置 label className
 */
func (a *TagControl) LabelClassName(value string) *TagControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *TagControl) Horizontal(value string) *TagControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 收纳标签的Popover配置
 */
func (a *TagControl) OverflowTagPopover(value string) *TagControl {
    a.Set("overflowTagPopover", value)
    return a
}

/**
 * 默认选择选项第一个值。
 */
func (a *TagControl) SelectFirst(value string) *TagControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * 控制编辑弹框设置项
 */
func (a *TagControl) EditDialog(value string) *TagControl {
    a.Set("editDialog", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *TagControl) VisibleOn(value string) *TagControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *TagControl) StaticOn(value string) *TagControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *TagControl) Static(value string) *TagControl {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *TagControl) StaticClassName(value string) *TagControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 点清除按钮时，将表单项设置成当前配置的值。
 */
func (a *TagControl) ResetValue(value string) *TagControl {
    a.Set("resetValue", value)
    return a
}

/**
 * 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
 */
func (a *TagControl) DeferApi(value string) *TagControl {
    a.Set("deferApi", value)
    return a
}

/**
 * 选项删除提示文字。
 */
func (a *TagControl) DeleteConfirmText(value string) *TagControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *TagControl) AutoFill(value string) *TagControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *TagControl) Inline(value string) *TagControl {
    a.Set("inline", value)
    return a
}

/**
 */
func (a *TagControl) Validations(value string) *TagControl {
    a.Set("validations", value)
    return a
}

/**
 * 是否开启批量添加模式
 */
func (a *TagControl) EnableBatchAdd(value string) *TagControl {
    a.Set("enableBatchAdd", value)
    return a
}

/**
 * 选项集合
 */
func (a *TagControl) Options(value string) *TagControl {
    a.Set("options", value)
    return a
}

/**
 * 用表达式来配置 source 接口初始要不要拉取
 */
func (a *TagControl) InitFetchOn(value string) *TagControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
 */
func (a *TagControl) ExtractValue(value string) *TagControl {
    a.Set("extractValue", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *TagControl) ValidateOnChange(value string) *TagControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 */
func (a *TagControl) Desc(value string) *TagControl {
    a.Set("desc", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *TagControl) StaticLabelClassName(value string) *TagControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *TagControl) StaticSchema(value string) *TagControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 允许添加的标签的最大数量
 */
func (a *TagControl) Max(value string) *TagControl {
    a.Set("max", value)
    return a
}

/**
 * 是否为多选模式
 */
func (a *TagControl) Multiple(value string) *TagControl {
    a.Set("multiple", value)
    return a
}

/**
 * 是否可清除。
 */
func (a *TagControl) Clearable(value string) *TagControl {
    a.Set("clearable", value)
    return a
}

/**
 * 控制新增弹框设置项
 */
func (a *TagControl) AddDialog(value string) *TagControl {
    a.Set("addDialog", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *TagControl) LabelWidth(value string) *TagControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 配置 input className
 */
func (a *TagControl) InputClassName(value string) *TagControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 多选模式，值太多时是否避免折行
 */
func (a *TagControl) ValuesNoWrap(value string) *TagControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 */
func (a *TagControl) InitAutoFill(value string) *TagControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *TagControl) DescriptionClassName(value string) *TagControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *TagControl) Remark(value string) *TagControl {
    a.Set("remark", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *TagControl) LabelRemark(value string) *TagControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 描述标题
 */
func (a *TagControl) LabelAlign(value string) *TagControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *TagControl) Description(value string) *TagControl {
    a.Set("description", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *TagControl) Mode(value string) *TagControl {
    a.Set("mode", value)
    return a
}

/**
 * 占位符
 */
func (a *TagControl) Placeholder(value string) *TagControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *TagControl) HiddenOn(value string) *TagControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *TagControl) StaticPlaceholder(value string) *TagControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *TagControl) StaticInputClassName(value string) *TagControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 添加时调用的接口
 */
func (a *TagControl) AddApi(value string) *TagControl {
    a.Set("addApi", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *TagControl) Size(value string) *TagControl {
    a.Set("size", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *TagControl) Hint(value string) *TagControl {
    a.Set("hint", value)
    return a
}

/**
 * 是否只读
 */
func (a *TagControl) ReadOnly(value string) *TagControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 只读条件
 */
func (a *TagControl) ReadOnlyOn(value string) *TagControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 组件样式
 */
func (a *TagControl) Style(value string) *TagControl {
    a.Set("style", value)
    return a
}

/**
 * 是否显示
 */
func (a *TagControl) Visible(value string) *TagControl {
    a.Set("visible", value)
    return a
}

/**
 * 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
 */
func (a *TagControl) JoinValues(value string) *TagControl {
    a.Set("joinValues", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *TagControl) ExtraName(value string) *TagControl {
    a.Set("extraName", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *TagControl) ValidationErrors(value string) *TagControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *TagControl) ValidateApi(value string) *TagControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *TagControl) Hidden(value string) *TagControl {
    a.Set("hidden", value)
    return a
}

/**
 * 新增文字
 */
func (a *TagControl) CreateBtnLabel(value string) *TagControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * 是否可以编辑
 */
func (a *TagControl) Editable(value string) *TagControl {
    a.Set("editable", value)
    return a
}

/**
 * 是否可删除
 */
func (a *TagControl) Removable(value string) *TagControl {
    a.Set("removable", value)
    return a
}

/**
 * 是否为下拉模式
 */
func (a *TagControl) Dropdown(value string) *TagControl {
    a.Set("dropdown", value)
    return a
}

/**
 * 标签的最大展示数量，超出数量后以收纳浮层的方式展示，仅在多选模式开启后生效
 */
func (a *TagControl) MaxTagCount(value string) *TagControl {
    a.Set("maxTagCount", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *TagControl) UseMobileUI(value string) *TagControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 选项提示信息
 */
func (a *TagControl) OptionsTip(value string) *TagControl {
    a.Set("optionsTip", value)
    return a
}

/**
 * 配置 source 接口初始拉不拉取。
 */
func (a *TagControl) InitFetch(value string) *TagControl {
    a.Set("initFetch", value)
    return a
}

/**
 * 懒加载字段
 */
func (a *TagControl) DeferField(value string) *TagControl {
    a.Set("deferField", value)
    return a
}

/**
 * 选项删除 API
 */
func (a *TagControl) DeleteApi(value string) *TagControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * 描述标题
 */
func (a *TagControl) Label(value string) *TagControl {
    a.Set("label", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *TagControl) OnEvent(value string) *TagControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否可以新增
 */
func (a *TagControl) Creatable(value string) *TagControl {
    a.Set("creatable", value)
    return a
}

/**
 * 是否禁用
 */
func (a *TagControl) Disabled(value string) *TagControl {
    a.Set("disabled", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *TagControl) EditorSetting(value string) *TagControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 单个标签的最大文本长度
 */
func (a *TagControl) MaxTagLength(value string) *TagControl {
    a.Set("maxTagLength", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *TagControl) Name(value string) *TagControl {
    a.Set("name", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *TagControl) SubmitOnChange(value string) *TagControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *TagControl) Id(value string) *TagControl {
    a.Set("id", value)
    return a
}

/**
 * 表单项类型
 */
func (a *TagControl) Type(value string) *TagControl {
    a.Set("type", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *TagControl) Width(value string) *TagControl {
    a.Set("width", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *TagControl) ClassName(value string) *TagControl {
    a.Set("className", value)
    return a
}

/**
 * 开启批量添加后，输入多个标签的分隔符，支持传入多个符号，默认为"-"
 */
func (a *TagControl) Separator(value string) *TagControl {
    a.Set("separator", value)
    return a
}

/**
 * 可用来通过 API 拉取 options。
 */
func (a *TagControl) Source(value string) *TagControl {
    a.Set("source", value)
    return a
}

/**
 * 新增时的表单项。
 */
func (a *TagControl) AddControls(value string) *TagControl {
    a.Set("addControls", value)
    return a
}

/**
 * 编辑时调用的 API
 */
func (a *TagControl) EditApi(value string) *TagControl {
    a.Set("editApi", value)
    return a
}

/**
 * 是否为必填
 */
func (a *TagControl) Required(value string) *TagControl {
    a.Set("required", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *TagControl) ClearValueOnHidden(value string) *TagControl {
    a.Set("clearValueOnHidden", value)
    return a
}
