package renderers


/**
 * Repeat 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/repeat

 * @author wcz0
 * @version 6.2.2
 */
type RepeatControl struct {
	*BaseRenderer
}

func NewRepeatControl() *RepeatControl {
    a := &RepeatControl{
        BaseRenderer: NewBaseRenderer(),
    }
    a.Set("type", "input-repeat")
    return a
}

/**
 * 描述标题
 */
func (a *RepeatControl) LabelAlign(value interface{}) *RepeatControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *RepeatControl) Name(value interface{}) *RepeatControl {
    a.Set("name", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *RepeatControl) Description(value interface{}) *RepeatControl {
    a.Set("description", value)
    return a
}

/**
 */
func (a *RepeatControl) InitAutoFill(value interface{}) *RepeatControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *RepeatControl) Value(value interface{}) *RepeatControl {
    a.Set("value", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *RepeatControl) Id(value interface{}) *RepeatControl {
    a.Set("id", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *RepeatControl) StaticOn(value interface{}) *RepeatControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 是否只读
 */
func (a *RepeatControl) ReadOnly(value interface{}) *RepeatControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *RepeatControl) Mode(value interface{}) *RepeatControl {
    a.Set("mode", value)
    return a
}

/**
 * 占位符
 */
func (a *RepeatControl) Placeholder(value interface{}) *RepeatControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *RepeatControl) StaticInputClassName(value interface{}) *RepeatControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 配置 label className
 */
func (a *RepeatControl) LabelClassName(value interface{}) *RepeatControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *RepeatControl) ExtraName(value interface{}) *RepeatControl {
    a.Set("extraName", value)
    return a
}

/**
 * 只读条件
 */
func (a *RepeatControl) ReadOnlyOn(value interface{}) *RepeatControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 */
func (a *RepeatControl) Options(value interface{}) *RepeatControl {
    a.Set("options", value)
    return a
}

/**
 */
func (a *RepeatControl) StaticSchema(value interface{}) *RepeatControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *RepeatControl) EditorSetting(value interface{}) *RepeatControl {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *RepeatControl) Desc(value interface{}) *RepeatControl {
    a.Set("desc", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *RepeatControl) Inline(value interface{}) *RepeatControl {
    a.Set("inline", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *RepeatControl) ValidationErrors(value interface{}) *RepeatControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *RepeatControl) Size(value interface{}) *RepeatControl {
    a.Set("size", value)
    return a
}

/**
 * 描述标题
 */
func (a *RepeatControl) Label(value interface{}) *RepeatControl {
    a.Set("label", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *RepeatControl) LabelRemark(value interface{}) *RepeatControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *RepeatControl) SubmitOnChange(value interface{}) *RepeatControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 组件样式
 */
func (a *RepeatControl) Style(value interface{}) *RepeatControl {
    a.Set("style", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *RepeatControl) UseMobileUI(value interface{}) *RepeatControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *RepeatControl) Row(value interface{}) *RepeatControl {
    a.Set("row", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *RepeatControl) VisibleOn(value interface{}) *RepeatControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *RepeatControl) StaticPlaceholder(value interface{}) *RepeatControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *RepeatControl) ValidateApi(value interface{}) *RepeatControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *RepeatControl) ClearValueOnHidden(value interface{}) *RepeatControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *RepeatControl) AutoFill(value interface{}) *RepeatControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *RepeatControl) DisabledOn(value interface{}) *RepeatControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *RepeatControl) Static(value interface{}) *RepeatControl {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *RepeatControl) StaticClassName(value interface{}) *RepeatControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *RepeatControl) Remark(value interface{}) *RepeatControl {
    a.Set("remark", value)
    return a
}

/**
 * 配置 input className
 */
func (a *RepeatControl) InputClassName(value interface{}) *RepeatControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *RepeatControl) ValidateOnChange(value interface{}) *RepeatControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *RepeatControl) Width(value interface{}) *RepeatControl {
    a.Set("width", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *RepeatControl) ClassName(value interface{}) *RepeatControl {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用
 */
func (a *RepeatControl) Disabled(value interface{}) *RepeatControl {
    a.Set("disabled", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *RepeatControl) StaticLabelClassName(value interface{}) *RepeatControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *RepeatControl) HiddenOn(value interface{}) *RepeatControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *RepeatControl) Hint(value interface{}) *RepeatControl {
    a.Set("hint", value)
    return a
}

/**
 * 表单项类型
 */
func (a *RepeatControl) Type(value interface{}) *RepeatControl {
    a.Set("type", value)
    return a
}

/**
 * 是否显示
 */
func (a *RepeatControl) Visible(value interface{}) *RepeatControl {
    a.Set("visible", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *RepeatControl) OnEvent(value interface{}) *RepeatControl {
    a.Set("onEvent", value)
    return a
}

/**
 */
func (a *RepeatControl) TestIdBuilder(value interface{}) *RepeatControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *RepeatControl) LabelWidth(value interface{}) *RepeatControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *RepeatControl) Horizontal(value interface{}) *RepeatControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *RepeatControl) DescriptionClassName(value interface{}) *RepeatControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 是否为必填
 */
func (a *RepeatControl) Required(value interface{}) *RepeatControl {
    a.Set("required", value)
    return a
}

/**
 */
func (a *RepeatControl) Validations(value interface{}) *RepeatControl {
    a.Set("validations", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *RepeatControl) Hidden(value interface{}) *RepeatControl {
    a.Set("hidden", value)
    return a
}
