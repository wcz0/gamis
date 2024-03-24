package renderers


/**
 * Switch 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/switch

 * @author wcz0
 * @version 6.2.2
 */
type SwitchControl struct {
	*BaseRenderer
}

func NewSwitchControl() *SwitchControl {
    a := &SwitchControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "switch")
    return a
}
/**
 * 容器 css 类名
 */
func (a *SwitchControl) ClassName(value interface{}) *SwitchControl {
    a.Set("className", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *SwitchControl) DescriptionClassName(value interface{}) *SwitchControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *SwitchControl) StaticClassName(value interface{}) *SwitchControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *SwitchControl) Value(value interface{}) *SwitchControl {
    a.Set("value", value)
    return a
}

/**
 * 勾选值
 */
func (a *SwitchControl) TrueValue(value interface{}) *SwitchControl {
    a.Set("trueValue", value)
    return a
}

/**
 * 未勾选值
 */
func (a *SwitchControl) FalseValue(value interface{}) *SwitchControl {
    a.Set("falseValue", value)
    return a
}

/**
 */
func (a *SwitchControl) Testid(value interface{}) *SwitchControl {
    a.Set("testid", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *SwitchControl) DisabledOn(value interface{}) *SwitchControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *SwitchControl) ValidationErrors(value interface{}) *SwitchControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 是否禁用
 */
func (a *SwitchControl) Disabled(value interface{}) *SwitchControl {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *SwitchControl) Hidden(value interface{}) *SwitchControl {
    a.Set("hidden", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *SwitchControl) Horizontal(value interface{}) *SwitchControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *SwitchControl) ValidateApi(value interface{}) *SwitchControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 选项说明
 */
func (a *SwitchControl) Option(value interface{}) *SwitchControl {
    a.Set("option", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *SwitchControl) StaticPlaceholder(value interface{}) *SwitchControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *SwitchControl) StaticLabelClassName(value interface{}) *SwitchControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *SwitchControl) ExtraName(value interface{}) *SwitchControl {
    a.Set("extraName", value)
    return a
}

/**
 * 只读条件
 */
func (a *SwitchControl) ReadOnlyOn(value interface{}) *SwitchControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *SwitchControl) ClearValueOnHidden(value interface{}) *SwitchControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *SwitchControl) Width(value interface{}) *SwitchControl {
    a.Set("width", value)
    return a
}

/**
 * 是否显示
 */
func (a *SwitchControl) Visible(value interface{}) *SwitchControl {
    a.Set("visible", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *SwitchControl) Name(value interface{}) *SwitchControl {
    a.Set("name", value)
    return a
}

/**
 * 指定为多行文本输入框
 */
func (a *SwitchControl) Type(value interface{}) *SwitchControl {
    a.Set("type", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *SwitchControl) StaticOn(value interface{}) *SwitchControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 开关尺寸
 * 可选值: sm | md
 */
func (a *SwitchControl) Size(value interface{}) *SwitchControl {
    a.Set("size", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *SwitchControl) Hint(value interface{}) *SwitchControl {
    a.Set("hint", value)
    return a
}

/**
 * 配置 input className
 */
func (a *SwitchControl) InputClassName(value interface{}) *SwitchControl {
    a.Set("inputClassName", value)
    return a
}

/**
 */
func (a *SwitchControl) Validations(value interface{}) *SwitchControl {
    a.Set("validations", value)
    return a
}

/**
 * 配置 label className
 */
func (a *SwitchControl) LabelClassName(value interface{}) *SwitchControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 开启时显示的内容
 */
func (a *SwitchControl) OnText(value interface{}) *SwitchControl {
    a.Set("onText", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *SwitchControl) EditorSetting(value interface{}) *SwitchControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *SwitchControl) VisibleOn(value interface{}) *SwitchControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *SwitchControl) Id(value interface{}) *SwitchControl {
    a.Set("id", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *SwitchControl) UseMobileUI(value interface{}) *SwitchControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *SwitchControl) LabelRemark(value interface{}) *SwitchControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *SwitchControl) Inline(value interface{}) *SwitchControl {
    a.Set("inline", value)
    return a
}

/**
 * 是否为必填
 */
func (a *SwitchControl) Required(value interface{}) *SwitchControl {
    a.Set("required", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *SwitchControl) StaticInputClassName(value interface{}) *SwitchControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *SwitchControl) Style(value interface{}) *SwitchControl {
    a.Set("style", value)
    return a
}

/**
 * 是否只读
 */
func (a *SwitchControl) ReadOnly(value interface{}) *SwitchControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *SwitchControl) Mode(value interface{}) *SwitchControl {
    a.Set("mode", value)
    return a
}

/**
 * 占位符
 */
func (a *SwitchControl) Placeholder(value interface{}) *SwitchControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *SwitchControl) OnEvent(value interface{}) *SwitchControl {
    a.Set("onEvent", value)
    return a
}

/**
 */
func (a *SwitchControl) StaticSchema(value interface{}) *SwitchControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *SwitchControl) LabelWidth(value interface{}) *SwitchControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *SwitchControl) Remark(value interface{}) *SwitchControl {
    a.Set("remark", value)
    return a
}

/**
 * 是否处于加载状态
 */
func (a *SwitchControl) Loading(value interface{}) *SwitchControl {
    a.Set("loading", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *SwitchControl) Static(value interface{}) *SwitchControl {
    a.Set("static", value)
    return a
}

/**
 * 描述标题
 */
func (a *SwitchControl) LabelAlign(value interface{}) *SwitchControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *SwitchControl) SubmitOnChange(value interface{}) *SwitchControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *SwitchControl) Description(value interface{}) *SwitchControl {
    a.Set("description", value)
    return a
}

/**
 * 关闭时显示的内容
 */
func (a *SwitchControl) OffText(value interface{}) *SwitchControl {
    a.Set("offText", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *SwitchControl) HiddenOn(value interface{}) *SwitchControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 描述标题
 */
func (a *SwitchControl) Label(value interface{}) *SwitchControl {
    a.Set("label", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *SwitchControl) ValidateOnChange(value interface{}) *SwitchControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 */
func (a *SwitchControl) Desc(value interface{}) *SwitchControl {
    a.Set("desc", value)
    return a
}
