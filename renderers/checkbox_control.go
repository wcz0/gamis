package renderers


/**
 * Checkbox 勾选框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/checkbox

 * @author wcz0
 * @version 6.2.2
 */
type CheckboxControl struct {
	*BaseRenderer
}

func NewCheckboxControl() *CheckboxControl {
    a := &CheckboxControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "checkbox")
    return a
}
/**
 * 编辑器配置，运行时可以忽略
 */
func (a *CheckboxControl) EditorSetting(value interface{}) *CheckboxControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *CheckboxControl) Size(value interface{}) *CheckboxControl {
    a.Set("size", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *CheckboxControl) Name(value interface{}) *CheckboxControl {
    a.Set("name", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *CheckboxControl) ExtraName(value interface{}) *CheckboxControl {
    a.Set("extraName", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *CheckboxControl) Remark(value interface{}) *CheckboxControl {
    a.Set("remark", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *CheckboxControl) ValidateOnChange(value interface{}) *CheckboxControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *CheckboxControl) DescriptionClassName(value interface{}) *CheckboxControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *CheckboxControl) Horizontal(value interface{}) *CheckboxControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *CheckboxControl) ValidateApi(value interface{}) *CheckboxControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 是否禁用
 */
func (a *CheckboxControl) Disabled(value interface{}) *CheckboxControl {
    a.Set("disabled", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *CheckboxControl) LabelWidth(value interface{}) *CheckboxControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 配置 label className
 */
func (a *CheckboxControl) LabelClassName(value interface{}) *CheckboxControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 勾选值
 */
func (a *CheckboxControl) TrueValue(value interface{}) *CheckboxControl {
    a.Set("trueValue", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *CheckboxControl) Width(value interface{}) *CheckboxControl {
    a.Set("width", value)
    return a
}

/**
 * 是否只读
 */
func (a *CheckboxControl) ReadOnly(value interface{}) *CheckboxControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *CheckboxControl) Description(value interface{}) *CheckboxControl {
    a.Set("description", value)
    return a
}

/**
 * 占位符
 */
func (a *CheckboxControl) Placeholder(value interface{}) *CheckboxControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *CheckboxControl) StaticInputClassName(value interface{}) *CheckboxControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 配置 input className
 */
func (a *CheckboxControl) InputClassName(value interface{}) *CheckboxControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *CheckboxControl) VisibleOn(value interface{}) *CheckboxControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *CheckboxControl) OnEvent(value interface{}) *CheckboxControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *CheckboxControl) StaticOn(value interface{}) *CheckboxControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *CheckboxControl) LabelRemark(value interface{}) *CheckboxControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *CheckboxControl) ClearValueOnHidden(value interface{}) *CheckboxControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 选项说明
 */
func (a *CheckboxControl) Option(value interface{}) *CheckboxControl {
    a.Set("option", value)
    return a
}

/**
 */
func (a *CheckboxControl) Checked(value interface{}) *CheckboxControl {
    a.Set("checked", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *CheckboxControl) StaticClassName(value interface{}) *CheckboxControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *CheckboxControl) Value(value interface{}) *CheckboxControl {
    a.Set("value", value)
    return a
}

/**
 * 是否显示
 */
func (a *CheckboxControl) Visible(value interface{}) *CheckboxControl {
    a.Set("visible", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *CheckboxControl) Id(value interface{}) *CheckboxControl {
    a.Set("id", value)
    return a
}

/**
 * 组件样式
 */
func (a *CheckboxControl) Style(value interface{}) *CheckboxControl {
    a.Set("style", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *CheckboxControl) ClassName(value interface{}) *CheckboxControl {
    a.Set("className", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *CheckboxControl) StaticPlaceholder(value interface{}) *CheckboxControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *CheckboxControl) StaticSchema(value interface{}) *CheckboxControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *CheckboxControl) UseMobileUI(value interface{}) *CheckboxControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *CheckboxControl) SubmitOnChange(value interface{}) *CheckboxControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 */
func (a *CheckboxControl) Desc(value interface{}) *CheckboxControl {
    a.Set("desc", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *CheckboxControl) StaticLabelClassName(value interface{}) *CheckboxControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 描述标题
 */
func (a *CheckboxControl) LabelAlign(value interface{}) *CheckboxControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *CheckboxControl) Hint(value interface{}) *CheckboxControl {
    a.Set("hint", value)
    return a
}

/**
 * 指定为多行文本输入框
 */
func (a *CheckboxControl) Type(value interface{}) *CheckboxControl {
    a.Set("type", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *CheckboxControl) HiddenOn(value interface{}) *CheckboxControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 角标
 */
func (a *CheckboxControl) Badge(value interface{}) *CheckboxControl {
    a.Set("badge", value)
    return a
}

/**
 * 是否为必填
 */
func (a *CheckboxControl) Required(value interface{}) *CheckboxControl {
    a.Set("required", value)
    return a
}

/**
 */
func (a *CheckboxControl) Validations(value interface{}) *CheckboxControl {
    a.Set("validations", value)
    return a
}

/**
 */
func (a *CheckboxControl) Partial(value interface{}) *CheckboxControl {
    a.Set("partial", value)
    return a
}

/**
 * 可选值: default | button
 */
func (a *CheckboxControl) OptionType(value interface{}) *CheckboxControl {
    a.Set("optionType", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *CheckboxControl) Static(value interface{}) *CheckboxControl {
    a.Set("static", value)
    return a
}

/**
 * 描述标题
 */
func (a *CheckboxControl) Label(value interface{}) *CheckboxControl {
    a.Set("label", value)
    return a
}

/**
 * 只读条件
 */
func (a *CheckboxControl) ReadOnlyOn(value interface{}) *CheckboxControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 未勾选值
 */
func (a *CheckboxControl) FalseValue(value interface{}) *CheckboxControl {
    a.Set("falseValue", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *CheckboxControl) Mode(value interface{}) *CheckboxControl {
    a.Set("mode", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *CheckboxControl) ValidationErrors(value interface{}) *CheckboxControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *CheckboxControl) DisabledOn(value interface{}) *CheckboxControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *CheckboxControl) Hidden(value interface{}) *CheckboxControl {
    a.Set("hidden", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *CheckboxControl) Inline(value interface{}) *CheckboxControl {
    a.Set("inline", value)
    return a
}
