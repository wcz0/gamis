package renderers


/**
 * UUID 功能性组件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/uuid

 * @author wcz0
 * @version 6.2.2
 */
type UUIDControl struct {
	*BaseRenderer
}

func NewUUIDControl() *UUIDControl {
    a := &UUIDControl{
        BaseRenderer: NewBaseRenderer(),
    }
    a.Set("type", "uuid")
    return a
}
/**
 * 验证失败的提示信息
 */
func (a *UUIDControl) ValidationErrors(value interface{}) *UUIDControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 表单项类型
 */
func (a *UUIDControl) Type(value interface{}) *UUIDControl {
    a.Set("type", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *UUIDControl) Hidden(value interface{}) *UUIDControl {
    a.Set("hidden", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *UUIDControl) HiddenOn(value interface{}) *UUIDControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *UUIDControl) ValidateOnChange(value interface{}) *UUIDControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *UUIDControl) Description(value interface{}) *UUIDControl {
    a.Set("description", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *UUIDControl) Inline(value interface{}) *UUIDControl {
    a.Set("inline", value)
    return a
}

/**
 * 是否为必填
 */
func (a *UUIDControl) Required(value interface{}) *UUIDControl {
    a.Set("required", value)
    return a
}

/**
 */
func (a *UUIDControl) StaticSchema(value interface{}) *UUIDControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 描述标题
 */
func (a *UUIDControl) Label(value interface{}) *UUIDControl {
    a.Set("label", value)
    return a
}

/**
 * 是否只读
 */
func (a *UUIDControl) ReadOnly(value interface{}) *UUIDControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 只读条件
 */
func (a *UUIDControl) ReadOnlyOn(value interface{}) *UUIDControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *UUIDControl) Mode(value interface{}) *UUIDControl {
    a.Set("mode", value)
    return a
}

/**
 */
func (a *UUIDControl) Validations(value interface{}) *UUIDControl {
    a.Set("validations", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *UUIDControl) Value(value interface{}) *UUIDControl {
    a.Set("value", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *UUIDControl) StaticClassName(value interface{}) *UUIDControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *UUIDControl) EditorSetting(value interface{}) *UUIDControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 描述标题
 */
func (a *UUIDControl) LabelAlign(value interface{}) *UUIDControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *UUIDControl) LabelWidth(value interface{}) *UUIDControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *UUIDControl) SubmitOnChange(value interface{}) *UUIDControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *UUIDControl) DescriptionClassName(value interface{}) *UUIDControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *UUIDControl) ValidateApi(value interface{}) *UUIDControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *UUIDControl) Width(value interface{}) *UUIDControl {
    a.Set("width", value)
    return a
}

/**
 */
func (a *UUIDControl) Desc(value interface{}) *UUIDControl {
    a.Set("desc", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *UUIDControl) DisabledOn(value interface{}) *UUIDControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *UUIDControl) StaticOn(value interface{}) *UUIDControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *UUIDControl) Name(value interface{}) *UUIDControl {
    a.Set("name", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *UUIDControl) ExtraName(value interface{}) *UUIDControl {
    a.Set("extraName", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *UUIDControl) LabelRemark(value interface{}) *UUIDControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *UUIDControl) Hint(value interface{}) *UUIDControl {
    a.Set("hint", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *UUIDControl) ClassName(value interface{}) *UUIDControl {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用
 */
func (a *UUIDControl) Disabled(value interface{}) *UUIDControl {
    a.Set("disabled", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *UUIDControl) Size(value interface{}) *UUIDControl {
    a.Set("size", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *UUIDControl) Horizontal(value interface{}) *UUIDControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *UUIDControl) Id(value interface{}) *UUIDControl {
    a.Set("id", value)
    return a
}

/**
 * 配置 label className
 */
func (a *UUIDControl) LabelClassName(value interface{}) *UUIDControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *UUIDControl) ClearValueOnHidden(value interface{}) *UUIDControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 配置 input className
 */
func (a *UUIDControl) InputClassName(value interface{}) *UUIDControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 占位符
 */
func (a *UUIDControl) Placeholder(value interface{}) *UUIDControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *UUIDControl) VisibleOn(value interface{}) *UUIDControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *UUIDControl) OnEvent(value interface{}) *UUIDControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *UUIDControl) Static(value interface{}) *UUIDControl {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *UUIDControl) StaticInputClassName(value interface{}) *UUIDControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *UUIDControl) Remark(value interface{}) *UUIDControl {
    a.Set("remark", value)
    return a
}

/**
 * 是否显示
 */
func (a *UUIDControl) Visible(value interface{}) *UUIDControl {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *UUIDControl) StaticPlaceholder(value interface{}) *UUIDControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *UUIDControl) StaticLabelClassName(value interface{}) *UUIDControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *UUIDControl) Style(value interface{}) *UUIDControl {
    a.Set("style", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *UUIDControl) UseMobileUI(value interface{}) *UUIDControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 长度，默认 uuid 的长度是 36，如果不需要那么长，可以设置这个来缩短
 */
func (a *UUIDControl) Length(value interface{}) *UUIDControl {
    a.Set("length", value)
    return a
}
