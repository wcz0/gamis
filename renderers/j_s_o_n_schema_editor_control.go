package renderers


/**
 * JSON Schema Editor 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/json-schema-editor

* @author wcz0
* @version 6.2.2
*/
type JSONSchemaEditorControl struct {
	*BaseRenderer
}

func NewJSONSchemaEditorControl() *JSONSchemaEditorControl {
    a := &JSONSchemaEditorControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "json-schema-editor")
    return a
}
/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *JSONSchemaEditorControl) Mode(value string) *JSONSchemaEditorControl {
    a.Set("mode", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *JSONSchemaEditorControl) DisabledOn(value string) *JSONSchemaEditorControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *JSONSchemaEditorControl) StaticInputClassName(value string) *JSONSchemaEditorControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 可以理解为类型模板，方便快速定义复杂类型
 */
func (a *JSONSchemaEditorControl) Definitions(value string) *JSONSchemaEditorControl {
    a.Set("definitions", value)
    return a
}

/**
 * 自定义详情配置面板如：{   boolean: [      {type: "input-text", name: "aa", label: "AA" }   ] }当配置布尔字段详情时，就会出现以上配置
 */
func (a *JSONSchemaEditorControl) AdvancedSettings(value string) *JSONSchemaEditorControl {
    a.Set("advancedSettings", value)
    return a
}

/**
 * 各属性输入控件的占位提示文本{   key: "key placeholder",   title: "title placeholder",   description: "description placeholder",   default: "default placeholder" }
 */
func (a *JSONSchemaEditorControl) Placeholder(value string) *JSONSchemaEditorControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *JSONSchemaEditorControl) SubmitOnChange(value string) *JSONSchemaEditorControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 只读条件
 */
func (a *JSONSchemaEditorControl) ReadOnlyOn(value string) *JSONSchemaEditorControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *JSONSchemaEditorControl) Description(value string) *JSONSchemaEditorControl {
    a.Set("description", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *JSONSchemaEditorControl) Horizontal(value string) *JSONSchemaEditorControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *JSONSchemaEditorControl) ClearValueOnHidden(value string) *JSONSchemaEditorControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *JSONSchemaEditorControl) HiddenOn(value string) *JSONSchemaEditorControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 顶层类型信息是否隐藏
 */
func (a *JSONSchemaEditorControl) ShowRootInfo(value string) *JSONSchemaEditorControl {
    a.Set("showRootInfo", value)
    return a
}

/**
 * 是否为必填
 */
func (a *JSONSchemaEditorControl) Required(value string) *JSONSchemaEditorControl {
    a.Set("required", value)
    return a
}

/**
 */
func (a *JSONSchemaEditorControl) Validations(value string) *JSONSchemaEditorControl {
    a.Set("validations", value)
    return a
}

/**
 * 配置 label className
 */
func (a *JSONSchemaEditorControl) LabelClassName(value string) *JSONSchemaEditorControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *JSONSchemaEditorControl) Id(value string) *JSONSchemaEditorControl {
    a.Set("id", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *JSONSchemaEditorControl) StaticOn(value string) *JSONSchemaEditorControl {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *JSONSchemaEditorControl) StaticSchema(value string) *JSONSchemaEditorControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *JSONSchemaEditorControl) UseMobileUI(value string) *JSONSchemaEditorControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *JSONSchemaEditorControl) Inline(value string) *JSONSchemaEditorControl {
    a.Set("inline", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *JSONSchemaEditorControl) Static(value string) *JSONSchemaEditorControl {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *JSONSchemaEditorControl) StaticClassName(value string) *JSONSchemaEditorControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *JSONSchemaEditorControl) Style(value string) *JSONSchemaEditorControl {
    a.Set("style", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *JSONSchemaEditorControl) Name(value string) *JSONSchemaEditorControl {
    a.Set("name", value)
    return a
}

/**
 * 是否禁用
 */
func (a *JSONSchemaEditorControl) Disabled(value string) *JSONSchemaEditorControl {
    a.Set("disabled", value)
    return a
}

/**
 * 是否只读
 */
func (a *JSONSchemaEditorControl) ReadOnly(value string) *JSONSchemaEditorControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *JSONSchemaEditorControl) DescriptionClassName(value string) *JSONSchemaEditorControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *JSONSchemaEditorControl) Hint(value string) *JSONSchemaEditorControl {
    a.Set("hint", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *JSONSchemaEditorControl) ValidateApi(value string) *JSONSchemaEditorControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *JSONSchemaEditorControl) ValidationErrors(value string) *JSONSchemaEditorControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *JSONSchemaEditorControl) StaticLabelClassName(value string) *JSONSchemaEditorControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 配置 input className
 */
func (a *JSONSchemaEditorControl) InputClassName(value string) *JSONSchemaEditorControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *JSONSchemaEditorControl) ClassName(value string) *JSONSchemaEditorControl {
    a.Set("className", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *JSONSchemaEditorControl) Width(value string) *JSONSchemaEditorControl {
    a.Set("width", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *JSONSchemaEditorControl) LabelWidth(value string) *JSONSchemaEditorControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *JSONSchemaEditorControl) ValidateOnChange(value string) *JSONSchemaEditorControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *JSONSchemaEditorControl) StaticPlaceholder(value string) *JSONSchemaEditorControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 开启详情配置
 */
func (a *JSONSchemaEditorControl) EnableAdvancedSetting(value string) *JSONSchemaEditorControl {
    a.Set("enableAdvancedSetting", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *JSONSchemaEditorControl) Size(value string) *JSONSchemaEditorControl {
    a.Set("size", value)
    return a
}

/**
 * 描述标题
 */
func (a *JSONSchemaEditorControl) LabelAlign(value string) *JSONSchemaEditorControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *JSONSchemaEditorControl) EditorSetting(value string) *JSONSchemaEditorControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 禁用类型，默认禁用了 null 类型
 */
func (a *JSONSchemaEditorControl) DisabledTypes(value string) *JSONSchemaEditorControl {
    a.Set("disabledTypes", value)
    return a
}

/**
 * 指定为 JSON Schema Editor
 */
func (a *JSONSchemaEditorControl) Type(value string) *JSONSchemaEditorControl {
    a.Set("type", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *JSONSchemaEditorControl) LabelRemark(value string) *JSONSchemaEditorControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 顶层是否允许修改类型
 */
func (a *JSONSchemaEditorControl) RootTypeMutable(value string) *JSONSchemaEditorControl {
    a.Set("rootTypeMutable", value)
    return a
}

/**
 * 描述标题
 */
func (a *JSONSchemaEditorControl) Label(value string) *JSONSchemaEditorControl {
    a.Set("label", value)
    return a
}

/**
 */
func (a *JSONSchemaEditorControl) Desc(value string) *JSONSchemaEditorControl {
    a.Set("desc", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *JSONSchemaEditorControl) OnEvent(value string) *JSONSchemaEditorControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *JSONSchemaEditorControl) Value(value string) *JSONSchemaEditorControl {
    a.Set("value", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *JSONSchemaEditorControl) Hidden(value string) *JSONSchemaEditorControl {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示
 */
func (a *JSONSchemaEditorControl) Visible(value string) *JSONSchemaEditorControl {
    a.Set("visible", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *JSONSchemaEditorControl) VisibleOn(value string) *JSONSchemaEditorControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *JSONSchemaEditorControl) Remark(value string) *JSONSchemaEditorControl {
    a.Set("remark", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *JSONSchemaEditorControl) ExtraName(value string) *JSONSchemaEditorControl {
    a.Set("extraName", value)
    return a
}
