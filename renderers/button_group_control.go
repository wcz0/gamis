package renderers


/**
 * 按钮组控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/button-group
 *

*/
type ButtonGroupControl struct {
	*BaseRenderer
}

func NewButtonGroupControl() *ButtonGroupControl {
    a := &ButtonGroupControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "button-group-select")
    return a
}
/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *ButtonGroupControl) Horizontal(value string) *ButtonGroupControl {
    a.Set("horizontal", value)
    return a
}

/**
 */
func (a *ButtonGroupControl) BtnActiveClassName(value string) *ButtonGroupControl {
    a.Set("btnActiveClassName", value)
    return a
}

/**
 * 平铺展示？
 */
func (a *ButtonGroupControl) Tiled(value string) *ButtonGroupControl {
    a.Set("tiled", value)
    return a
}

/**
 */
func (a *ButtonGroupControl) Type(value string) *ButtonGroupControl {
    a.Set("type", value)
    return a
}

/**
 * 可用来通过 API 拉取 options。
 */
func (a *ButtonGroupControl) Source(value string) *ButtonGroupControl {
    a.Set("source", value)
    return a
}

/**
 * 配置 label className
 */
func (a *ButtonGroupControl) LabelClassName(value string) *ButtonGroupControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *ButtonGroupControl) Description(value string) *ButtonGroupControl {
    a.Set("description", value)
    return a
}

/**
 * 选项删除 API
 */
func (a *ButtonGroupControl) DeleteApi(value string) *ButtonGroupControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *ButtonGroupControl) Static(value string) *ButtonGroupControl {
    a.Set("static", value)
    return a
}

/**
 */
func (a *ButtonGroupControl) Desc(value string) *ButtonGroupControl {
    a.Set("desc", value)
    return a
}

/**
 * 编辑时调用的 API
 */
func (a *ButtonGroupControl) EditApi(value string) *ButtonGroupControl {
    a.Set("editApi", value)
    return a
}

/**
 * 选项修改的表单项
 */
func (a *ButtonGroupControl) EditControls(value string) *ButtonGroupControl {
    a.Set("editControls", value)
    return a
}

/**
 * 是否只读
 */
func (a *ButtonGroupControl) ReadOnly(value string) *ButtonGroupControl {
    a.Set("readOnly", value)
    return a
}

/**
 */
func (a *ButtonGroupControl) Validations(value string) *ButtonGroupControl {
    a.Set("validations", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *ButtonGroupControl) ValidateApi(value string) *ButtonGroupControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *ButtonGroupControl) StaticOn(value string) *ButtonGroupControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *ButtonGroupControl) EditorSetting(value string) *ButtonGroupControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 是否可清除。
 */
func (a *ButtonGroupControl) Clearable(value string) *ButtonGroupControl {
    a.Set("clearable", value)
    return a
}

/**
 * 添加时调用的接口
 */
func (a *ButtonGroupControl) AddApi(value string) *ButtonGroupControl {
    a.Set("addApi", value)
    return a
}

/**
 */
func (a *ButtonGroupControl) InitAutoFill(value string) *ButtonGroupControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *ButtonGroupControl) Value(value string) *ButtonGroupControl {
    a.Set("value", value)
    return a
}

/**
 * 是否为禁用状态。
 */
func (a *ButtonGroupControl) Disabled(value string) *ButtonGroupControl {
    a.Set("disabled", value)
    return a
}

/**
 * 通过 JS 表达式来配置当前表单项的禁用状态。
 */
func (a *ButtonGroupControl) DisabledOn(value string) *ButtonGroupControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *ButtonGroupControl) Visible(value string) *ButtonGroupControl {
    a.Set("visible", value)
    return a
}

/**
 */
func (a *ButtonGroupControl) StaticSchema(value string) *ButtonGroupControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 控制新增弹框设置项
 */
func (a *ButtonGroupControl) AddDialog(value string) *ButtonGroupControl {
    a.Set("addDialog", value)
    return a
}

/**
 * 选项删除提示文字。
 */
func (a *ButtonGroupControl) DeleteConfirmText(value string) *ButtonGroupControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * 描述标题
 */
func (a *ButtonGroupControl) Label(value string) *ButtonGroupControl {
    a.Set("label", value)
    return a
}

/**
 * 按钮大小
 * 可选值: xs | sm | md | lg
 */
func (a *ButtonGroupControl) Size(value string) *ButtonGroupControl {
    a.Set("size", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *ButtonGroupControl) StaticLabelClassName(value string) *ButtonGroupControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 是否为多选模式
 */
func (a *ButtonGroupControl) Multiple(value string) *ButtonGroupControl {
    a.Set("multiple", value)
    return a
}

/**
 * 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
 */
func (a *ButtonGroupControl) ExtractValue(value string) *ButtonGroupControl {
    a.Set("extractValue", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *ButtonGroupControl) ExtraName(value string) *ButtonGroupControl {
    a.Set("extraName", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *ButtonGroupControl) Inline(value string) *ButtonGroupControl {
    a.Set("inline", value)
    return a
}

/**
 * 控制编辑弹框设置项
 */
func (a *ButtonGroupControl) EditDialog(value string) *ButtonGroupControl {
    a.Set("editDialog", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *ButtonGroupControl) AutoFill(value string) *ButtonGroupControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 描述标题
 */
func (a *ButtonGroupControl) LabelAlign(value string) *ButtonGroupControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 分割符
 */
func (a *ButtonGroupControl) Delimiter(value string) *ButtonGroupControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *ButtonGroupControl) UseMobileUI(value string) *ButtonGroupControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 是否为必填
 */
func (a *ButtonGroupControl) Required(value string) *ButtonGroupControl {
    a.Set("required", value)
    return a
}

/**
 * 垂直展示？
 */
func (a *ButtonGroupControl) Vertical(value string) *ButtonGroupControl {
    a.Set("vertical", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *ButtonGroupControl) Name(value string) *ButtonGroupControl {
    a.Set("name", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *ButtonGroupControl) ValidateOnChange(value string) *ButtonGroupControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *ButtonGroupControl) ValidationErrors(value string) *ButtonGroupControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *ButtonGroupControl) StaticPlaceholder(value string) *ButtonGroupControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 点清除按钮时，将表单项设置成当前配置的值。
 */
func (a *ButtonGroupControl) ResetValue(value string) *ButtonGroupControl {
    a.Set("resetValue", value)
    return a
}

/**
 * 是否可以新增
 */
func (a *ButtonGroupControl) Creatable(value string) *ButtonGroupControl {
    a.Set("creatable", value)
    return a
}

/**
 * 是否可删除
 */
func (a *ButtonGroupControl) Removable(value string) *ButtonGroupControl {
    a.Set("removable", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *ButtonGroupControl) ClearValueOnHidden(value string) *ButtonGroupControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *ButtonGroupControl) ClassName(value string) *ButtonGroupControl {
    a.Set("className", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *ButtonGroupControl) Id(value string) *ButtonGroupControl {
    a.Set("id", value)
    return a
}

/**
 * 按钮样式级别
 */
func (a *ButtonGroupControl) BtnLevel(value string) *ButtonGroupControl {
    a.Set("btnLevel", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *ButtonGroupControl) Width(value string) *ButtonGroupControl {
    a.Set("width", value)
    return a
}

/**
 * 新增文字
 */
func (a *ButtonGroupControl) CreateBtnLabel(value string) *ButtonGroupControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *ButtonGroupControl) SubmitOnChange(value string) *ButtonGroupControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 只读条件
 */
func (a *ButtonGroupControl) ReadOnlyOn(value string) *ButtonGroupControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 配置 input className
 */
func (a *ButtonGroupControl) InputClassName(value string) *ButtonGroupControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *ButtonGroupControl) StaticClassName(value string) *ButtonGroupControl {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *ButtonGroupControl) BtnClassName(value string) *ButtonGroupControl {
    a.Set("btnClassName", value)
    return a
}

/**
 * 按钮集合
 */
func (a *ButtonGroupControl) Buttons(value string) *ButtonGroupControl {
    a.Set("buttons", value)
    return a
}

/**
 * 按钮选中的样式级别
 */
func (a *ButtonGroupControl) BtnActiveLevel(value string) *ButtonGroupControl {
    a.Set("btnActiveLevel", value)
    return a
}

/**
 * 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
 */
func (a *ButtonGroupControl) DeferApi(value string) *ButtonGroupControl {
    a.Set("deferApi", value)
    return a
}

/**
 * 新增时的表单项。
 */
func (a *ButtonGroupControl) AddControls(value string) *ButtonGroupControl {
    a.Set("addControls", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *ButtonGroupControl) LabelWidth(value string) *ButtonGroupControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *ButtonGroupControl) DescriptionClassName(value string) *ButtonGroupControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *ButtonGroupControl) Mode(value string) *ButtonGroupControl {
    a.Set("mode", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *ButtonGroupControl) StaticInputClassName(value string) *ButtonGroupControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 多选模式，值太多时是否避免折行
 */
func (a *ButtonGroupControl) ValuesNoWrap(value string) *ButtonGroupControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *ButtonGroupControl) Remark(value string) *ButtonGroupControl {
    a.Set("remark", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *ButtonGroupControl) Hint(value string) *ButtonGroupControl {
    a.Set("hint", value)
    return a
}

/**
 * 占位符
 */
func (a *ButtonGroupControl) Placeholder(value string) *ButtonGroupControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *ButtonGroupControl) Hidden(value string) *ButtonGroupControl {
    a.Set("hidden", value)
    return a
}

/**
 * 通过 JS 表达式来配置当前表单项是否显示
 */
func (a *ButtonGroupControl) VisibleOn(value string) *ButtonGroupControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 配置 source 接口初始拉不拉取。
 */
func (a *ButtonGroupControl) InitFetch(value string) *ButtonGroupControl {
    a.Set("initFetch", value)
    return a
}

/**
 * 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
 */
func (a *ButtonGroupControl) JoinValues(value string) *ButtonGroupControl {
    a.Set("joinValues", value)
    return a
}

/**
 * 是否可以编辑
 */
func (a *ButtonGroupControl) Editable(value string) *ButtonGroupControl {
    a.Set("editable", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *ButtonGroupControl) OnEvent(value string) *ButtonGroupControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 选项集合
 */
func (a *ButtonGroupControl) Options(value string) *ButtonGroupControl {
    a.Set("options", value)
    return a
}

/**
 * 默认选择选项第一个值。
 */
func (a *ButtonGroupControl) SelectFirst(value string) *ButtonGroupControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *ButtonGroupControl) LabelRemark(value string) *ButtonGroupControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 组件样式
 */
func (a *ButtonGroupControl) Style(value string) *ButtonGroupControl {
    a.Set("style", value)
    return a
}

/**
 * 用表达式来配置 source 接口初始要不要拉取
 */
func (a *ButtonGroupControl) InitFetchOn(value string) *ButtonGroupControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 懒加载字段
 */
func (a *ButtonGroupControl) DeferField(value string) *ButtonGroupControl {
    a.Set("deferField", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *ButtonGroupControl) HiddenOn(value string) *ButtonGroupControl {
    a.Set("hiddenOn", value)
    return a
}
