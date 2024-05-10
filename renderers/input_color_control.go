package renderers


/**
 * Color 颜色选择框 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/color

 * @author wcz0
 * @version 6.2.2
 */
type InputColorControl struct {
	*BaseRenderer
}

func NewInputColorControl() *InputColorControl {
    a := &InputColorControl{
        BaseRenderer: NewBaseRenderer(),
    }
    a.Set("type", "input-color")
    return a
}

/**
 * 颜色格式
 * 可选值: hex | rgb | rgba | hsl
 */
func (a *InputColorControl) Format(value interface{}) *InputColorControl {
    a.Set("format", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *InputColorControl) EditorSetting(value interface{}) *InputColorControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 只读条件
 */
func (a *InputColorControl) ReadOnlyOn(value interface{}) *InputColorControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *InputColorControl) Mode(value interface{}) *InputColorControl {
    a.Set("mode", value)
    return a
}

/**
 */
func (a *InputColorControl) Validations(value interface{}) *InputColorControl {
    a.Set("validations", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *InputColorControl) Value(value interface{}) *InputColorControl {
    a.Set("value", value)
    return a
}

/**
 * 是否显示清除按钮
 */
func (a *InputColorControl) Clearable(value interface{}) *InputColorControl {
    a.Set("clearable", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *InputColorControl) StaticClassName(value interface{}) *InputColorControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 是否为必填
 */
func (a *InputColorControl) Required(value interface{}) *InputColorControl {
    a.Set("required", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *InputColorControl) ValidationErrors(value interface{}) *InputColorControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *InputColorControl) DisabledOn(value interface{}) *InputColorControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *InputColorControl) LabelRemark(value interface{}) *InputColorControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *InputColorControl) Width(value interface{}) *InputColorControl {
    a.Set("width", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *InputColorControl) StaticInputClassName(value interface{}) *InputColorControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *InputColorControl) Size(value interface{}) *InputColorControl {
    a.Set("size", value)
    return a
}

/**
 * 配置 input className
 */
func (a *InputColorControl) InputClassName(value interface{}) *InputColorControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *InputColorControl) ClearValueOnHidden(value interface{}) *InputColorControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 预设颜色，用户可以直接从预设中选。
 */
func (a *InputColorControl) PresetColors(value interface{}) *InputColorControl {
    a.Set("presetColors", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *InputColorControl) ClassName(value interface{}) *InputColorControl {
    a.Set("className", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *InputColorControl) VisibleOn(value interface{}) *InputColorControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *InputColorControl) Name(value interface{}) *InputColorControl {
    a.Set("name", value)
    return a
}

/**
 * 是否只读
 */
func (a *InputColorControl) ReadOnly(value interface{}) *InputColorControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 指定为颜色选择框
 */
func (a *InputColorControl) Type(value interface{}) *InputColorControl {
    a.Set("type", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *InputColorControl) Static(value interface{}) *InputColorControl {
    a.Set("static", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *InputColorControl) StaticOn(value interface{}) *InputColorControl {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *InputColorControl) StaticSchema(value interface{}) *InputColorControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *InputColorControl) Hidden(value interface{}) *InputColorControl {
    a.Set("hidden", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *InputColorControl) ValidateApi(value interface{}) *InputColorControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 描述标题
 */
func (a *InputColorControl) Label(value interface{}) *InputColorControl {
    a.Set("label", value)
    return a
}

/**
 * 是否禁用
 */
func (a *InputColorControl) Disabled(value interface{}) *InputColorControl {
    a.Set("disabled", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *InputColorControl) StaticPlaceholder(value interface{}) *InputColorControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 配置 label className
 */
func (a *InputColorControl) LabelClassName(value interface{}) *InputColorControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *InputColorControl) StaticLabelClassName(value interface{}) *InputColorControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *InputColorControl) TestIdBuilder(value interface{}) *InputColorControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *InputColorControl) Remark(value interface{}) *InputColorControl {
    a.Set("remark", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *InputColorControl) SubmitOnChange(value interface{}) *InputColorControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *InputColorControl) ValidateOnChange(value interface{}) *InputColorControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *InputColorControl) Horizontal(value interface{}) *InputColorControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *InputColorControl) Inline(value interface{}) *InputColorControl {
    a.Set("inline", value)
    return a
}

/**
 * 占位符
 */
func (a *InputColorControl) Placeholder(value interface{}) *InputColorControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *InputColorControl) HiddenOn(value interface{}) *InputColorControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *InputColorControl) Id(value interface{}) *InputColorControl {
    a.Set("id", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *InputColorControl) UseMobileUI(value interface{}) *InputColorControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *InputColorControl) ExtraName(value interface{}) *InputColorControl {
    a.Set("extraName", value)
    return a
}

/**
 */
func (a *InputColorControl) Desc(value interface{}) *InputColorControl {
    a.Set("desc", value)
    return a
}

/**
 * 选中颜色后是否关闭弹出层。
 */
func (a *InputColorControl) CloseOnSelect(value interface{}) *InputColorControl {
    a.Set("closeOnSelect", value)
    return a
}

/**
 * 是否显示
 */
func (a *InputColorControl) Visible(value interface{}) *InputColorControl {
    a.Set("visible", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *InputColorControl) DescriptionClassName(value interface{}) *InputColorControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *InputColorControl) AutoFill(value interface{}) *InputColorControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 是否允许用户输入颜色。
 */
func (a *InputColorControl) AllowCustomColor(value interface{}) *InputColorControl {
    a.Set("allowCustomColor", value)
    return a
}

/**
 * 描述标题
 */
func (a *InputColorControl) LabelAlign(value interface{}) *InputColorControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *InputColorControl) Hint(value interface{}) *InputColorControl {
    a.Set("hint", value)
    return a
}

/**
 */
func (a *InputColorControl) InitAutoFill(value interface{}) *InputColorControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *InputColorControl) LabelWidth(value interface{}) *InputColorControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *InputColorControl) Description(value interface{}) *InputColorControl {
    a.Set("description", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *InputColorControl) OnEvent(value interface{}) *InputColorControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 组件样式
 */
func (a *InputColorControl) Style(value interface{}) *InputColorControl {
    a.Set("style", value)
    return a
}
