package renderers


/**
 * 公式功能控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/formula

 * @author wcz0
 * @version 6.2.2
 */
type FormulaControl struct {
	*BaseRenderer
}

func NewFormulaControl() *FormulaControl {
    a := &FormulaControl{
        BaseRenderer: NewBaseRenderer(),
    }
    a.Set("type", "formula")
    return a
}

/**
 * 公式
 */
func (a *FormulaControl) Formula(value interface{}) *FormulaControl {
    a.Set("formula", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *FormulaControl) StaticClassName(value interface{}) *FormulaControl {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *FormulaControl) TestIdBuilder(value interface{}) *FormulaControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *FormulaControl) Mode(value interface{}) *FormulaControl {
    a.Set("mode", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *FormulaControl) Inline(value interface{}) *FormulaControl {
    a.Set("inline", value)
    return a
}

/**
 * 指定为公式功能控件。
 */
func (a *FormulaControl) Type(value interface{}) *FormulaControl {
    a.Set("type", value)
    return a
}

/**
 * 配置 input className
 */
func (a *FormulaControl) InputClassName(value interface{}) *FormulaControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 占位符
 */
func (a *FormulaControl) Placeholder(value interface{}) *FormulaControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *FormulaControl) VisibleOn(value interface{}) *FormulaControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *FormulaControl) StaticOn(value interface{}) *FormulaControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *FormulaControl) HiddenOn(value interface{}) *FormulaControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 字段名，公式结果将作用到此处指定的变量中去
 */
func (a *FormulaControl) Name(value interface{}) *FormulaControl {
    a.Set("name", value)
    return a
}

/**
 * 只读条件
 */
func (a *FormulaControl) ReadOnlyOn(value interface{}) *FormulaControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *FormulaControl) ValidateApi(value interface{}) *FormulaControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *FormulaControl) DisabledOn(value interface{}) *FormulaControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *FormulaControl) StaticLabelClassName(value interface{}) *FormulaControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *FormulaControl) EditorSetting(value interface{}) *FormulaControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *FormulaControl) Horizontal(value interface{}) *FormulaControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *FormulaControl) ClassName(value interface{}) *FormulaControl {
    a.Set("className", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *FormulaControl) Remark(value interface{}) *FormulaControl {
    a.Set("remark", value)
    return a
}

/**
 */
func (a *FormulaControl) Validations(value interface{}) *FormulaControl {
    a.Set("validations", value)
    return a
}

/**
 * 是否初始应用
 */
func (a *FormulaControl) InitSet(value interface{}) *FormulaControl {
    a.Set("initSet", value)
    return a
}

/**
 * 是否显示
 */
func (a *FormulaControl) Visible(value interface{}) *FormulaControl {
    a.Set("visible", value)
    return a
}

/**
 * 当某个按钮的目标指定为此值后，会触发一次公式应用。这个机制可以在 autoSet 为 false 时用来手动触发
 */
func (a *FormulaControl) Id(value interface{}) *FormulaControl {
    a.Set("id", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *FormulaControl) OnEvent(value interface{}) *FormulaControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 组件样式
 */
func (a *FormulaControl) Style(value interface{}) *FormulaControl {
    a.Set("style", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *FormulaControl) Value(value interface{}) *FormulaControl {
    a.Set("value", value)
    return a
}

/**
 * 触发公式的作用条件，如 data.xxx == \"a\" 或者 ${xx}
 */
func (a *FormulaControl) Condition(value interface{}) *FormulaControl {
    a.Set("condition", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *FormulaControl) Hidden(value interface{}) *FormulaControl {
    a.Set("hidden", value)
    return a
}

/**
 * 配置 label className
 */
func (a *FormulaControl) LabelClassName(value interface{}) *FormulaControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *FormulaControl) SubmitOnChange(value interface{}) *FormulaControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *FormulaControl) Description(value interface{}) *FormulaControl {
    a.Set("description", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *FormulaControl) ValidationErrors(value interface{}) *FormulaControl {
    a.Set("validationErrors", value)
    return a
}

/**
 */
func (a *FormulaControl) InitAutoFill(value interface{}) *FormulaControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *FormulaControl) UseMobileUI(value interface{}) *FormulaControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 描述标题
 */
func (a *FormulaControl) Label(value interface{}) *FormulaControl {
    a.Set("label", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *FormulaControl) ClearValueOnHidden(value interface{}) *FormulaControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *FormulaControl) Width(value interface{}) *FormulaControl {
    a.Set("width", value)
    return a
}

/**
 * 是否禁用
 */
func (a *FormulaControl) Disabled(value interface{}) *FormulaControl {
    a.Set("disabled", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *FormulaControl) ExtraName(value interface{}) *FormulaControl {
    a.Set("extraName", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *FormulaControl) Hint(value interface{}) *FormulaControl {
    a.Set("hint", value)
    return a
}

/**
 */
func (a *FormulaControl) Row(value interface{}) *FormulaControl {
    a.Set("row", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *FormulaControl) Size(value interface{}) *FormulaControl {
    a.Set("size", value)
    return a
}

/**
 * 是否只读
 */
func (a *FormulaControl) ReadOnly(value interface{}) *FormulaControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *FormulaControl) ValidateOnChange(value interface{}) *FormulaControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 是否自动应用
 */
func (a *FormulaControl) AutoSet(value interface{}) *FormulaControl {
    a.Set("autoSet", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *FormulaControl) StaticPlaceholder(value interface{}) *FormulaControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *FormulaControl) StaticInputClassName(value interface{}) *FormulaControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否为必填
 */
func (a *FormulaControl) Required(value interface{}) *FormulaControl {
    a.Set("required", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *FormulaControl) Static(value interface{}) *FormulaControl {
    a.Set("static", value)
    return a
}

/**
 */
func (a *FormulaControl) StaticSchema(value interface{}) *FormulaControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *FormulaControl) LabelRemark(value interface{}) *FormulaControl {
    a.Set("labelRemark", value)
    return a
}

/**
 */
func (a *FormulaControl) Desc(value interface{}) *FormulaControl {
    a.Set("desc", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *FormulaControl) DescriptionClassName(value interface{}) *FormulaControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 描述标题
 */
func (a *FormulaControl) LabelAlign(value interface{}) *FormulaControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *FormulaControl) LabelWidth(value interface{}) *FormulaControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *FormulaControl) AutoFill(value interface{}) *FormulaControl {
    a.Set("autoFill", value)
    return a
}
