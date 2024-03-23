package renderers


/**
 * Radio 单选框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/radios
 *

*/
type RadioControl struct {
	*BaseRenderer
}

func NewRadioControl() *RadioControl {
    a := &RadioControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "radio")
    return a
}
/**
 */
func (a *RadioControl) Partial(value string) *RadioControl {
    a.Set("partial", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *RadioControl) StaticInputClassName(value string) *RadioControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *RadioControl) Mode(value string) *RadioControl {
    a.Set("mode", value)
    return a
}

/**
 * 是否为必填
 */
func (a *RadioControl) Required(value string) *RadioControl {
    a.Set("required", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *RadioControl) ClearValueOnHidden(value string) *RadioControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *RadioControl) Hidden(value string) *RadioControl {
    a.Set("hidden", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *RadioControl) SubmitOnChange(value string) *RadioControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *RadioControl) Inline(value string) *RadioControl {
    a.Set("inline", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *RadioControl) Size(value string) *RadioControl {
    a.Set("size", value)
    return a
}

/**
 * 配置 label className
 */
func (a *RadioControl) LabelClassName(value string) *RadioControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *RadioControl) Width(value string) *RadioControl {
    a.Set("width", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *RadioControl) ValidateOnChange(value string) *RadioControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *RadioControl) Horizontal(value string) *RadioControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 未勾选值
 */
func (a *RadioControl) FalseValue(value string) *RadioControl {
    a.Set("falseValue", value)
    return a
}

/**
 * 选项说明
 */
func (a *RadioControl) Option(value string) *RadioControl {
    a.Set("option", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *RadioControl) VisibleOn(value string) *RadioControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *RadioControl) Id(value string) *RadioControl {
    a.Set("id", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *RadioControl) EditorSetting(value string) *RadioControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *RadioControl) ValidateApi(value string) *RadioControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 可选值: default | button
 */
func (a *RadioControl) OptionType(value string) *RadioControl {
    a.Set("optionType", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *RadioControl) StaticLabelClassName(value string) *RadioControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *RadioControl) Remark(value string) *RadioControl {
    a.Set("remark", value)
    return a
}

/**
 */
func (a *RadioControl) Desc(value string) *RadioControl {
    a.Set("desc", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *RadioControl) StaticOn(value string) *RadioControl {
    a.Set("staticOn", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *RadioControl) LabelWidth(value string) *RadioControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *RadioControl) ExtraName(value string) *RadioControl {
    a.Set("extraName", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *RadioControl) StaticPlaceholder(value string) *RadioControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *RadioControl) LabelRemark(value string) *RadioControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *RadioControl) ValidationErrors(value string) *RadioControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *RadioControl) Name(value string) *RadioControl {
    a.Set("name", value)
    return a
}

/**
 * 角标
 */
func (a *RadioControl) Badge(value string) *RadioControl {
    a.Set("badge", value)
    return a
}

/**
 * 只读条件
 */
func (a *RadioControl) ReadOnlyOn(value string) *RadioControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *RadioControl) Value(value string) *RadioControl {
    a.Set("value", value)
    return a
}

/**
 * 配置 input className
 */
func (a *RadioControl) InputClassName(value string) *RadioControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 指定为多行文本输入框
 */
func (a *RadioControl) Type(value string) *RadioControl {
    a.Set("type", value)
    return a
}

/**
 * 是否禁用
 */
func (a *RadioControl) Disabled(value string) *RadioControl {
    a.Set("disabled", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *RadioControl) UseMobileUI(value string) *RadioControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *RadioControl) Hint(value string) *RadioControl {
    a.Set("hint", value)
    return a
}

/**
 * 描述标题
 */
func (a *RadioControl) Label(value string) *RadioControl {
    a.Set("label", value)
    return a
}

/**
 */
func (a *RadioControl) Validations(value string) *RadioControl {
    a.Set("validations", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *RadioControl) Static(value string) *RadioControl {
    a.Set("static", value)
    return a
}

/**
 * 是否只读
 */
func (a *RadioControl) ReadOnly(value string) *RadioControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 占位符
 */
func (a *RadioControl) Placeholder(value string) *RadioControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 勾选值
 */
func (a *RadioControl) TrueValue(value string) *RadioControl {
    a.Set("trueValue", value)
    return a
}

/**
 * 是否显示
 */
func (a *RadioControl) Visible(value string) *RadioControl {
    a.Set("visible", value)
    return a
}

/**
 * 组件样式
 */
func (a *RadioControl) Style(value string) *RadioControl {
    a.Set("style", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *RadioControl) Description(value string) *RadioControl {
    a.Set("description", value)
    return a
}

/**
 */
func (a *RadioControl) StaticSchema(value string) *RadioControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 描述标题
 */
func (a *RadioControl) LabelAlign(value string) *RadioControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *RadioControl) ClassName(value string) *RadioControl {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *RadioControl) DisabledOn(value string) *RadioControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *RadioControl) StaticClassName(value string) *RadioControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *RadioControl) HiddenOn(value string) *RadioControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *RadioControl) OnEvent(value string) *RadioControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *RadioControl) DescriptionClassName(value string) *RadioControl {
    a.Set("descriptionClassName", value)
    return a
}
