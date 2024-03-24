package renderers


/**
 * 图标选择器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/icon-picker

 * @author wcz0
 * @version 6.2.2
 */
type IconPickerControl struct {
	*BaseRenderer
}

func NewIconPickerControl() *IconPickerControl {
    a := &IconPickerControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "icon-picker")
    return a
}
/**
 * 静态展示表单项Value类名
 */
func (a *IconPickerControl) StaticInputClassName(value interface{}) *IconPickerControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 描述标题
 */
func (a *IconPickerControl) Label(value interface{}) *IconPickerControl {
    a.Set("label", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *IconPickerControl) SubmitOnChange(value interface{}) *IconPickerControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *IconPickerControl) ValidationErrors(value interface{}) *IconPickerControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *IconPickerControl) Width(value interface{}) *IconPickerControl {
    a.Set("width", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *IconPickerControl) DisabledOn(value interface{}) *IconPickerControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *IconPickerControl) StaticPlaceholder(value interface{}) *IconPickerControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 描述标题
 */
func (a *IconPickerControl) LabelAlign(value interface{}) *IconPickerControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *IconPickerControl) Value(value interface{}) *IconPickerControl {
    a.Set("value", value)
    return a
}

/**
 * 表单项类型
 */
func (a *IconPickerControl) Type(value interface{}) *IconPickerControl {
    a.Set("type", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *IconPickerControl) Hidden(value interface{}) *IconPickerControl {
    a.Set("hidden", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *IconPickerControl) StaticOn(value interface{}) *IconPickerControl {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *IconPickerControl) StaticSchema(value interface{}) *IconPickerControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 组件样式
 */
func (a *IconPickerControl) Style(value interface{}) *IconPickerControl {
    a.Set("style", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *IconPickerControl) Size(value interface{}) *IconPickerControl {
    a.Set("size", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *IconPickerControl) LabelWidth(value interface{}) *IconPickerControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *IconPickerControl) Mode(value interface{}) *IconPickerControl {
    a.Set("mode", value)
    return a
}

/**
 * 是否为必填
 */
func (a *IconPickerControl) Required(value interface{}) *IconPickerControl {
    a.Set("required", value)
    return a
}

/**
 */
func (a *IconPickerControl) Validations(value interface{}) *IconPickerControl {
    a.Set("validations", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *IconPickerControl) ClassName(value interface{}) *IconPickerControl {
    a.Set("className", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *IconPickerControl) OnEvent(value interface{}) *IconPickerControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *IconPickerControl) Static(value interface{}) *IconPickerControl {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *IconPickerControl) StaticLabelClassName(value interface{}) *IconPickerControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 占位符
 */
func (a *IconPickerControl) Placeholder(value interface{}) *IconPickerControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *IconPickerControl) VisibleOn(value interface{}) *IconPickerControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *IconPickerControl) EditorSetting(value interface{}) *IconPickerControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *IconPickerControl) Hint(value interface{}) *IconPickerControl {
    a.Set("hint", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *IconPickerControl) Inline(value interface{}) *IconPickerControl {
    a.Set("inline", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *IconPickerControl) Name(value interface{}) *IconPickerControl {
    a.Set("name", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *IconPickerControl) ExtraName(value interface{}) *IconPickerControl {
    a.Set("extraName", value)
    return a
}

/**
 * 是否只读
 */
func (a *IconPickerControl) ReadOnly(value interface{}) *IconPickerControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *IconPickerControl) ClearValueOnHidden(value interface{}) *IconPickerControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 配置 label className
 */
func (a *IconPickerControl) LabelClassName(value interface{}) *IconPickerControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *IconPickerControl) Remark(value interface{}) *IconPickerControl {
    a.Set("remark", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *IconPickerControl) LabelRemark(value interface{}) *IconPickerControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 只读条件
 */
func (a *IconPickerControl) ReadOnlyOn(value interface{}) *IconPickerControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *IconPickerControl) ValidateOnChange(value interface{}) *IconPickerControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 */
func (a *IconPickerControl) Desc(value interface{}) *IconPickerControl {
    a.Set("desc", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *IconPickerControl) DescriptionClassName(value interface{}) *IconPickerControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *IconPickerControl) Horizontal(value interface{}) *IconPickerControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 配置 input className
 */
func (a *IconPickerControl) InputClassName(value interface{}) *IconPickerControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *IconPickerControl) ValidateApi(value interface{}) *IconPickerControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 是否禁用
 */
func (a *IconPickerControl) Disabled(value interface{}) *IconPickerControl {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *IconPickerControl) HiddenOn(value interface{}) *IconPickerControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *IconPickerControl) Visible(value interface{}) *IconPickerControl {
    a.Set("visible", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *IconPickerControl) Id(value interface{}) *IconPickerControl {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *IconPickerControl) StaticClassName(value interface{}) *IconPickerControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *IconPickerControl) UseMobileUI(value interface{}) *IconPickerControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *IconPickerControl) Description(value interface{}) *IconPickerControl {
    a.Set("description", value)
    return a
}
