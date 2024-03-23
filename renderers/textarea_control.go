package renderers


/**
 * TextArea 多行文本输入框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/textarea
 *

*/
type TextareaControl struct {
	*BaseRenderer
}

func NewTextareaControl() *TextareaControl {
    a := &TextareaControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "textarea")
    return a
}
/**
 * 指定为多行文本输入框
 */
func (a *TextareaControl) Type(value string) *TextareaControl {
    a.Set("type", value)
    return a
}

/**
 * 是否显示计数
 */
func (a *TextareaControl) ShowCounter(value string) *TextareaControl {
    a.Set("showCounter", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *TextareaControl) Size(value string) *TextareaControl {
    a.Set("size", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *TextareaControl) SubmitOnChange(value string) *TextareaControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *TextareaControl) Description(value string) *TextareaControl {
    a.Set("description", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *TextareaControl) ValidateApi(value string) *TextareaControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 是否禁用
 */
func (a *TextareaControl) Disabled(value string) *TextareaControl {
    a.Set("disabled", value)
    return a
}

/**
 * 描述标题
 */
func (a *TextareaControl) Label(value string) *TextareaControl {
    a.Set("label", value)
    return a
}

/**
 * 描述标题
 */
func (a *TextareaControl) LabelAlign(value string) *TextareaControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *TextareaControl) UseMobileUI(value string) *TextareaControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 配置 input className
 */
func (a *TextareaControl) InputClassName(value string) *TextareaControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *TextareaControl) ValidationErrors(value string) *TextareaControl {
    a.Set("validationErrors", value)
    return a
}

/**
 */
func (a *TextareaControl) Validations(value string) *TextareaControl {
    a.Set("validations", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *TextareaControl) Value(value string) *TextareaControl {
    a.Set("value", value)
    return a
}

/**
 */
func (a *TextareaControl) Testid(value string) *TextareaControl {
    a.Set("testid", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *TextareaControl) DisabledOn(value string) *TextareaControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *TextareaControl) StaticOn(value string) *TextareaControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 是否为必填
 */
func (a *TextareaControl) Required(value string) *TextareaControl {
    a.Set("required", value)
    return a
}

/**
 * 重置值
 */
func (a *TextareaControl) ResetValue(value string) *TextareaControl {
    a.Set("resetValue", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *TextareaControl) StaticInputClassName(value string) *TextareaControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 最小行数
 */
func (a *TextareaControl) MinRows(value string) *TextareaControl {
    a.Set("minRows", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *TextareaControl) Inline(value string) *TextareaControl {
    a.Set("inline", value)
    return a
}

/**
 * 最大行数
 */
func (a *TextareaControl) MaxRows(value string) *TextareaControl {
    a.Set("maxRows", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *TextareaControl) OnEvent(value string) *TextareaControl {
    a.Set("onEvent", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *TextareaControl) LabelWidth(value string) *TextareaControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *TextareaControl) ValidateOnChange(value string) *TextareaControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *TextareaControl) Name(value string) *TextareaControl {
    a.Set("name", value)
    return a
}

/**
 * 输入内容是否可清除
 */
func (a *TextareaControl) Clearable(value string) *TextareaControl {
    a.Set("clearable", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *TextareaControl) ClassName(value string) *TextareaControl {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *TextareaControl) Hidden(value string) *TextareaControl {
    a.Set("hidden", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *TextareaControl) StaticPlaceholder(value string) *TextareaControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *TextareaControl) StaticClassName(value string) *TextareaControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 只读条件
 */
func (a *TextareaControl) ReadOnlyOn(value string) *TextareaControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *TextareaControl) ClearValueOnHidden(value string) *TextareaControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 是否显示
 */
func (a *TextareaControl) Visible(value string) *TextareaControl {
    a.Set("visible", value)
    return a
}

/**
 * 占位符
 */
func (a *TextareaControl) Placeholder(value string) *TextareaControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 组件样式
 */
func (a *TextareaControl) Style(value string) *TextareaControl {
    a.Set("style", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *TextareaControl) DescriptionClassName(value string) *TextareaControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *TextareaControl) Id(value string) *TextareaControl {
    a.Set("id", value)
    return a
}

/**
 */
func (a *TextareaControl) Desc(value string) *TextareaControl {
    a.Set("desc", value)
    return a
}

/**
 * 限制文字个数
 */
func (a *TextareaControl) MaxLength(value string) *TextareaControl {
    a.Set("maxLength", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *TextareaControl) VisibleOn(value string) *TextareaControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *TextareaControl) LabelRemark(value string) *TextareaControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 是否只读
 */
func (a *TextareaControl) ReadOnly(value string) *TextareaControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 配置 label className
 */
func (a *TextareaControl) LabelClassName(value string) *TextareaControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *TextareaControl) Remark(value string) *TextareaControl {
    a.Set("remark", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *TextareaControl) Hint(value string) *TextareaControl {
    a.Set("hint", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *TextareaControl) Static(value string) *TextareaControl {
    a.Set("static", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *TextareaControl) EditorSetting(value string) *TextareaControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *TextareaControl) Horizontal(value string) *TextareaControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *TextareaControl) HiddenOn(value string) *TextareaControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *TextareaControl) StaticLabelClassName(value string) *TextareaControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 边框模式，全边框，还是半边框，或者没边框。
 * 可选值: full | half | none
 */
func (a *TextareaControl) BorderMode(value string) *TextareaControl {
    a.Set("borderMode", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *TextareaControl) Width(value string) *TextareaControl {
    a.Set("width", value)
    return a
}

/**
 */
func (a *TextareaControl) StaticSchema(value string) *TextareaControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *TextareaControl) ExtraName(value string) *TextareaControl {
    a.Set("extraName", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *TextareaControl) Mode(value string) *TextareaControl {
    a.Set("mode", value)
    return a
}