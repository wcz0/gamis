package renderers


/**
 * RichText 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/input-rich-text
 *

*/
type RichTextControl struct {
	*BaseRenderer
}

func NewRichTextControl() *RichTextControl {
    a := &RichTextControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-rich-text")
    return a
}
/**
 * 是否禁用
 */
func (a *RichTextControl) Disabled(value string) *RichTextControl {
    a.Set("disabled", value)
    return a
}

/**
 * 是否显示
 */
func (a *RichTextControl) Visible(value string) *RichTextControl {
    a.Set("visible", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *RichTextControl) UseMobileUI(value string) *RichTextControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 边框模式，全边框，还是半边框，或者没边框。
 * 可选值: full | half | none
 */
func (a *RichTextControl) BorderMode(value string) *RichTextControl {
    a.Set("borderMode", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *RichTextControl) DisabledOn(value string) *RichTextControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 配置 label className
 */
func (a *RichTextControl) LabelClassName(value string) *RichTextControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *RichTextControl) Name(value string) *RichTextControl {
    a.Set("name", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *RichTextControl) ValidationErrors(value string) *RichTextControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *RichTextControl) ClassName(value string) *RichTextControl {
    a.Set("className", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *RichTextControl) Hint(value string) *RichTextControl {
    a.Set("hint", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *RichTextControl) StaticOn(value string) *RichTextControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 是否为必填
 */
func (a *RichTextControl) Required(value string) *RichTextControl {
    a.Set("required", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *RichTextControl) Description(value string) *RichTextControl {
    a.Set("description", value)
    return a
}

/**
 */
func (a *RichTextControl) Desc(value string) *RichTextControl {
    a.Set("desc", value)
    return a
}

/**
 * 表单项类型
 */
func (a *RichTextControl) Type(value string) *RichTextControl {
    a.Set("type", value)
    return a
}

/**
 * 是否只读
 */
func (a *RichTextControl) ReadOnly(value string) *RichTextControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *RichTextControl) Mode(value string) *RichTextControl {
    a.Set("mode", value)
    return a
}

/**
 */
func (a *RichTextControl) Validations(value string) *RichTextControl {
    a.Set("validations", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *RichTextControl) StaticPlaceholder(value string) *RichTextControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *RichTextControl) ValidateOnChange(value string) *RichTextControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *RichTextControl) DescriptionClassName(value string) *RichTextControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *RichTextControl) Value(value string) *RichTextControl {
    a.Set("value", value)
    return a
}

/**
 */
func (a *RichTextControl) StaticSchema(value string) *RichTextControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *RichTextControl) Size(value string) *RichTextControl {
    a.Set("size", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *RichTextControl) ClearValueOnHidden(value string) *RichTextControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 接收器的字段名
 */
func (a *RichTextControl) FileField(value string) *RichTextControl {
    a.Set("fileField", value)
    return a
}

/**
 * tinymce 或 froala 的配置
 */
func (a *RichTextControl) Options(value string) *RichTextControl {
    a.Set("options", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *RichTextControl) Hidden(value string) *RichTextControl {
    a.Set("hidden", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *RichTextControl) HiddenOn(value string) *RichTextControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *RichTextControl) Id(value string) *RichTextControl {
    a.Set("id", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *RichTextControl) SubmitOnChange(value string) *RichTextControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *RichTextControl) Inline(value string) *RichTextControl {
    a.Set("inline", value)
    return a
}

/**
 * 占位符
 */
func (a *RichTextControl) Placeholder(value string) *RichTextControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *RichTextControl) OnEvent(value string) *RichTextControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *RichTextControl) Static(value string) *RichTextControl {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *RichTextControl) StaticLabelClassName(value string) *RichTextControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 描述标题
 */
func (a *RichTextControl) Label(value string) *RichTextControl {
    a.Set("label", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *RichTextControl) LabelRemark(value string) *RichTextControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *RichTextControl) ValidateApi(value string) *RichTextControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 描述标题
 */
func (a *RichTextControl) LabelAlign(value string) *RichTextControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *RichTextControl) LabelWidth(value string) *RichTextControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 只读条件
 */
func (a *RichTextControl) ReadOnlyOn(value string) *RichTextControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 视频保存 API
 */
func (a *RichTextControl) VideoReceiver(value string) *RichTextControl {
    a.Set("videoReceiver", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *RichTextControl) VisibleOn(value string) *RichTextControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 组件样式
 */
func (a *RichTextControl) Style(value string) *RichTextControl {
    a.Set("style", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *RichTextControl) EditorSetting(value string) *RichTextControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 图片保存 API
 */
func (a *RichTextControl) Receiver(value string) *RichTextControl {
    a.Set("receiver", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *RichTextControl) StaticInputClassName(value string) *RichTextControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *RichTextControl) Horizontal(value string) *RichTextControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *RichTextControl) Width(value string) *RichTextControl {
    a.Set("width", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *RichTextControl) Remark(value string) *RichTextControl {
    a.Set("remark", value)
    return a
}

/**
 * 配置 input className
 */
func (a *RichTextControl) InputClassName(value string) *RichTextControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *RichTextControl) StaticClassName(value string) *RichTextControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *RichTextControl) ExtraName(value string) *RichTextControl {
    a.Set("extraName", value)
    return a
}

/**
 * 编辑器类型
 * 可选值: froala | tinymce
 */
func (a *RichTextControl) Vendor(value string) *RichTextControl {
    a.Set("vendor", value)
    return a
}
