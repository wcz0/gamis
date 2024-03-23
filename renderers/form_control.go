package renderers


/**
 * Control 表单项包裹 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/control
 *

*/
type FormControl struct {
	*BaseRenderer
}

func NewFormControl() *FormControl {
    a := &FormControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "control")
    return a
}
/**
 * 描述标题
 */
func (a *FormControl) Label(value string) *FormControl {
    a.Set("label", value)
    return a
}

/**
 * 配置 label className
 */
func (a *FormControl) LabelClassName(value string) *FormControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 配置 input className
 */
func (a *FormControl) InputClassName(value string) *FormControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 是否禁用
 */
func (a *FormControl) Disabled(value string) *FormControl {
    a.Set("disabled", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *FormControl) Id(value string) *FormControl {
    a.Set("id", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *FormControl) StaticOn(value string) *FormControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *FormControl) StaticInputClassName(value string) *FormControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *FormControl) EditorSetting(value string) *FormControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *FormControl) ValidationErrors(value string) *FormControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *FormControl) ClearValueOnHidden(value string) *FormControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * FormItem 内容
 */
func (a *FormControl) Body(value string) *FormControl {
    a.Set("body", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *FormControl) Inline(value string) *FormControl {
    a.Set("inline", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *FormControl) Static(value string) *FormControl {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *FormControl) StaticClassName(value string) *FormControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *FormControl) LabelWidth(value string) *FormControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *FormControl) ExtraName(value string) *FormControl {
    a.Set("extraName", value)
    return a
}

/**
 */
func (a *FormControl) Validations(value string) *FormControl {
    a.Set("validations", value)
    return a
}

/**
 * 是否显示
 */
func (a *FormControl) Visible(value string) *FormControl {
    a.Set("visible", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *FormControl) Name(value string) *FormControl {
    a.Set("name", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *FormControl) Remark(value string) *FormControl {
    a.Set("remark", value)
    return a
}

/**
 * 只读条件
 */
func (a *FormControl) ReadOnlyOn(value string) *FormControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *FormControl) ValidateOnChange(value string) *FormControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 占位符
 */
func (a *FormControl) Placeholder(value string) *FormControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *FormControl) Value(value string) *FormControl {
    a.Set("value", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *FormControl) ClassName(value string) *FormControl {
    a.Set("className", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *FormControl) Size(value string) *FormControl {
    a.Set("size", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *FormControl) Hint(value string) *FormControl {
    a.Set("hint", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *FormControl) SubmitOnChange(value string) *FormControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *FormControl) Horizontal(value string) *FormControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 是否为必填
 */
func (a *FormControl) Required(value string) *FormControl {
    a.Set("required", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *FormControl) Width(value string) *FormControl {
    a.Set("width", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *FormControl) DisabledOn(value string) *FormControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *FormControl) HiddenOn(value string) *FormControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *FormControl) StaticPlaceholder(value string) *FormControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *FormControl) UseMobileUI(value string) *FormControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 描述标题
 */
func (a *FormControl) LabelAlign(value string) *FormControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 是否只读
 */
func (a *FormControl) ReadOnly(value string) *FormControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *FormControl) Mode(value string) *FormControl {
    a.Set("mode", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *FormControl) ValidateApi(value string) *FormControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *FormControl) OnEvent(value string) *FormControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *FormControl) StaticLabelClassName(value string) *FormControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *FormControl) StaticSchema(value string) *FormControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 组件样式
 */
func (a *FormControl) Style(value string) *FormControl {
    a.Set("style", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *FormControl) LabelRemark(value string) *FormControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *FormControl) Hidden(value string) *FormControl {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *FormControl) VisibleOn(value string) *FormControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *FormControl) Description(value string) *FormControl {
    a.Set("description", value)
    return a
}

/**
 */
func (a *FormControl) Desc(value string) *FormControl {
    a.Set("desc", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *FormControl) DescriptionClassName(value string) *FormControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 表单项类型
 */
func (a *FormControl) Type(value string) *FormControl {
    a.Set("type", value)
    return a
}
