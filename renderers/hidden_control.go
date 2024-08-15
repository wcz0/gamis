package renderers


/**
 * Hidden 隐藏域。功能性组件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/hidden

 * @author wcz0
 * @version 6.2.2
 */
type HiddenControl struct {
	*BaseRenderer
}

func NewHiddenControl() *HiddenControl {
    a := &HiddenControl{
        BaseRenderer: NewBaseRenderer(),
    }
    a.Set("type", "hidden")
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *HiddenControl) Inline(value interface{}) *HiddenControl {
    a.Set("inline", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *HiddenControl) Mode(value interface{}) *HiddenControl {
    a.Set("mode", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *HiddenControl) LabelRemark(value interface{}) *HiddenControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 是否禁用
 */
func (a *HiddenControl) Disabled(value interface{}) *HiddenControl {
    a.Set("disabled", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *HiddenControl) UseMobileUI(value interface{}) *HiddenControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *HiddenControl) Size(value interface{}) *HiddenControl {
    a.Set("size", value)
    return a
}

/**
 * 描述标题
 */
func (a *HiddenControl) LabelAlign(value interface{}) *HiddenControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *HiddenControl) Horizontal(value interface{}) *HiddenControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *HiddenControl) ValidateOnChange(value interface{}) *HiddenControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *HiddenControl) Description(value interface{}) *HiddenControl {
    a.Set("description", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *HiddenControl) ValidationErrors(value interface{}) *HiddenControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *HiddenControl) OnEvent(value interface{}) *HiddenControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 配置 label className
 */
func (a *HiddenControl) LabelClassName(value interface{}) *HiddenControl {
    a.Set("labelClassName", value)
    return a
}

/**
 */
func (a *HiddenControl) Desc(value interface{}) *HiddenControl {
    a.Set("desc", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *HiddenControl) Id(value interface{}) *HiddenControl {
    a.Set("id", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *HiddenControl) EditorSetting(value interface{}) *HiddenControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *HiddenControl) Hint(value interface{}) *HiddenControl {
    a.Set("hint", value)
    return a
}

/**
 * 是否只读
 */
func (a *HiddenControl) ReadOnly(value interface{}) *HiddenControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *HiddenControl) Hidden(value interface{}) *HiddenControl {
    a.Set("hidden", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *HiddenControl) StaticClassName(value interface{}) *HiddenControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *HiddenControl) Style(value interface{}) *HiddenControl {
    a.Set("style", value)
    return a
}

/**
 */
func (a *HiddenControl) TestIdBuilder(value interface{}) *HiddenControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *HiddenControl) SubmitOnChange(value interface{}) *HiddenControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 只读条件
 */
func (a *HiddenControl) ReadOnlyOn(value interface{}) *HiddenControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 配置 input className
 */
func (a *HiddenControl) InputClassName(value interface{}) *HiddenControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 表单项类型
 */
func (a *HiddenControl) Type(value interface{}) *HiddenControl {
    a.Set("type", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *HiddenControl) ClassName(value interface{}) *HiddenControl {
    a.Set("className", value)
    return a
}

/**
 */
func (a *HiddenControl) StaticSchema(value interface{}) *HiddenControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *HiddenControl) StaticLabelClassName(value interface{}) *HiddenControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 描述标题
 */
func (a *HiddenControl) Label(value interface{}) *HiddenControl {
    a.Set("label", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *HiddenControl) Value(value interface{}) *HiddenControl {
    a.Set("value", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *HiddenControl) LabelWidth(value interface{}) *HiddenControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *HiddenControl) Remark(value interface{}) *HiddenControl {
    a.Set("remark", value)
    return a
}

/**
 */
func (a *HiddenControl) Validations(value interface{}) *HiddenControl {
    a.Set("validations", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *HiddenControl) ClearValueOnHidden(value interface{}) *HiddenControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *HiddenControl) ValidateApi(value interface{}) *HiddenControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *HiddenControl) AutoFill(value interface{}) *HiddenControl {
    a.Set("autoFill", value)
    return a
}

/**
 */
func (a *HiddenControl) Row(value interface{}) *HiddenControl {
    a.Set("row", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *HiddenControl) StaticPlaceholder(value interface{}) *HiddenControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *HiddenControl) ExtraName(value interface{}) *HiddenControl {
    a.Set("extraName", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *HiddenControl) Width(value interface{}) *HiddenControl {
    a.Set("width", value)
    return a
}

/**
 * 占位符
 */
func (a *HiddenControl) Placeholder(value interface{}) *HiddenControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *HiddenControl) VisibleOn(value interface{}) *HiddenControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *HiddenControl) DescriptionClassName(value interface{}) *HiddenControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *HiddenControl) Name(value interface{}) *HiddenControl {
    a.Set("name", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *HiddenControl) StaticOn(value interface{}) *HiddenControl {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *HiddenControl) InitAutoFill(value interface{}) *HiddenControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 是否显示
 */
func (a *HiddenControl) Visible(value interface{}) *HiddenControl {
    a.Set("visible", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *HiddenControl) Static(value interface{}) *HiddenControl {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *HiddenControl) StaticInputClassName(value interface{}) *HiddenControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *HiddenControl) HiddenOn(value interface{}) *HiddenControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *HiddenControl) DisabledOn(value interface{}) *HiddenControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否为必填
 */
func (a *HiddenControl) Required(value interface{}) *HiddenControl {
    a.Set("required", value)
    return a
}
