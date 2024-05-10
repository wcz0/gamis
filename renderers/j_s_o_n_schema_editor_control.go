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
 * 事件动作配置
 */
func (a *JSONSchemaEditorControl) OnEvent(value interface{}) *JSONSchemaEditorControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *JSONSchemaEditorControl) StaticOn(value interface{}) *JSONSchemaEditorControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 顶层是否允许修改类型
 */
func (a *JSONSchemaEditorControl) RootTypeMutable(value interface{}) *JSONSchemaEditorControl {
    a.Set("rootTypeMutable", value)
    return a
}

/**
 * 是否只读
 */
func (a *JSONSchemaEditorControl) ReadOnly(value interface{}) *JSONSchemaEditorControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 只读条件
 */
func (a *JSONSchemaEditorControl) ReadOnlyOn(value interface{}) *JSONSchemaEditorControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *JSONSchemaEditorControl) Description(value interface{}) *JSONSchemaEditorControl {
    a.Set("description", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *JSONSchemaEditorControl) LabelWidth(value interface{}) *JSONSchemaEditorControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *JSONSchemaEditorControl) SubmitOnChange(value interface{}) *JSONSchemaEditorControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 */
func (a *JSONSchemaEditorControl) Desc(value interface{}) *JSONSchemaEditorControl {
    a.Set("desc", value)
    return a
}

/**
 * 配置 input className
 */
func (a *JSONSchemaEditorControl) InputClassName(value interface{}) *JSONSchemaEditorControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *JSONSchemaEditorControl) StaticInputClassName(value interface{}) *JSONSchemaEditorControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 可以理解为类型模板，方便快速定义复杂类型
 */
func (a *JSONSchemaEditorControl) Definitions(value interface{}) *JSONSchemaEditorControl {
    a.Set("definitions", value)
    return a
}

/**
 * 自定义详情配置面板如：{   boolean: [      {type: "input-text", name: "aa", label: "AA" }   ] }当配置布尔字段详情时，就会出现以上配置
 */
func (a *JSONSchemaEditorControl) AdvancedSettings(value interface{}) *JSONSchemaEditorControl {
    a.Set("advancedSettings", value)
    return a
}

/**
 * 描述标题
 */
func (a *JSONSchemaEditorControl) LabelAlign(value interface{}) *JSONSchemaEditorControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 配置 label className
 */
func (a *JSONSchemaEditorControl) LabelClassName(value interface{}) *JSONSchemaEditorControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *JSONSchemaEditorControl) Hidden(value interface{}) *JSONSchemaEditorControl {
    a.Set("hidden", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *JSONSchemaEditorControl) Static(value interface{}) *JSONSchemaEditorControl {
    a.Set("static", value)
    return a
}

/**
 * 是否禁用
 */
func (a *JSONSchemaEditorControl) Disabled(value interface{}) *JSONSchemaEditorControl {
    a.Set("disabled", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *JSONSchemaEditorControl) UseMobileUI(value interface{}) *JSONSchemaEditorControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 顶层类型信息是否隐藏
 */
func (a *JSONSchemaEditorControl) ShowRootInfo(value interface{}) *JSONSchemaEditorControl {
    a.Set("showRootInfo", value)
    return a
}

/**
 * 禁用类型，默认禁用了 null 类型
 */
func (a *JSONSchemaEditorControl) DisabledTypes(value interface{}) *JSONSchemaEditorControl {
    a.Set("disabledTypes", value)
    return a
}

/**
 * 描述标题
 */
func (a *JSONSchemaEditorControl) Label(value interface{}) *JSONSchemaEditorControl {
    a.Set("label", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *JSONSchemaEditorControl) Inline(value interface{}) *JSONSchemaEditorControl {
    a.Set("inline", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *JSONSchemaEditorControl) StaticLabelClassName(value interface{}) *JSONSchemaEditorControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *JSONSchemaEditorControl) Style(value interface{}) *JSONSchemaEditorControl {
    a.Set("style", value)
    return a
}

/**
 * 开启详情配置
 */
func (a *JSONSchemaEditorControl) EnableAdvancedSetting(value interface{}) *JSONSchemaEditorControl {
    a.Set("enableAdvancedSetting", value)
    return a
}

/**
 * 是否为迷你模式，会隐藏一些不必要的元素
 */
func (a *JSONSchemaEditorControl) Mini(value interface{}) *JSONSchemaEditorControl {
    a.Set("mini", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *JSONSchemaEditorControl) DisabledOn(value interface{}) *JSONSchemaEditorControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *JSONSchemaEditorControl) Visible(value interface{}) *JSONSchemaEditorControl {
    a.Set("visible", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *JSONSchemaEditorControl) DescriptionClassName(value interface{}) *JSONSchemaEditorControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *JSONSchemaEditorControl) VisibleOn(value interface{}) *JSONSchemaEditorControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *JSONSchemaEditorControl) EditorSetting(value interface{}) *JSONSchemaEditorControl {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *JSONSchemaEditorControl) TestIdBuilder(value interface{}) *JSONSchemaEditorControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *JSONSchemaEditorControl) ValidateOnChange(value interface{}) *JSONSchemaEditorControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *JSONSchemaEditorControl) ValidateApi(value interface{}) *JSONSchemaEditorControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *JSONSchemaEditorControl) ClassName(value interface{}) *JSONSchemaEditorControl {
    a.Set("className", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *JSONSchemaEditorControl) StaticClassName(value interface{}) *JSONSchemaEditorControl {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *JSONSchemaEditorControl) Validations(value interface{}) *JSONSchemaEditorControl {
    a.Set("validations", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *JSONSchemaEditorControl) AutoFill(value interface{}) *JSONSchemaEditorControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *JSONSchemaEditorControl) ClearValueOnHidden(value interface{}) *JSONSchemaEditorControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *JSONSchemaEditorControl) HiddenOn(value interface{}) *JSONSchemaEditorControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 各属性输入控件的占位提示文本{   key: "key placeholder",   title: "title placeholder",   description: "description placeholder",   default: "default placeholder" }
 */
func (a *JSONSchemaEditorControl) Placeholder(value interface{}) *JSONSchemaEditorControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *JSONSchemaEditorControl) Remark(value interface{}) *JSONSchemaEditorControl {
    a.Set("remark", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *JSONSchemaEditorControl) Size(value interface{}) *JSONSchemaEditorControl {
    a.Set("size", value)
    return a
}

/**
 */
func (a *JSONSchemaEditorControl) StaticSchema(value interface{}) *JSONSchemaEditorControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *JSONSchemaEditorControl) Name(value interface{}) *JSONSchemaEditorControl {
    a.Set("name", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *JSONSchemaEditorControl) Value(value interface{}) *JSONSchemaEditorControl {
    a.Set("value", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *JSONSchemaEditorControl) ValidationErrors(value interface{}) *JSONSchemaEditorControl {
    a.Set("validationErrors", value)
    return a
}

/**
 */
func (a *JSONSchemaEditorControl) InitAutoFill(value interface{}) *JSONSchemaEditorControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *JSONSchemaEditorControl) Width(value interface{}) *JSONSchemaEditorControl {
    a.Set("width", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *JSONSchemaEditorControl) Hint(value interface{}) *JSONSchemaEditorControl {
    a.Set("hint", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *JSONSchemaEditorControl) Horizontal(value interface{}) *JSONSchemaEditorControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *JSONSchemaEditorControl) ExtraName(value interface{}) *JSONSchemaEditorControl {
    a.Set("extraName", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *JSONSchemaEditorControl) StaticPlaceholder(value interface{}) *JSONSchemaEditorControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 指定为 JSON Schema Editor
 */
func (a *JSONSchemaEditorControl) Type(value interface{}) *JSONSchemaEditorControl {
    a.Set("type", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *JSONSchemaEditorControl) LabelRemark(value interface{}) *JSONSchemaEditorControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 是否为必填
 */
func (a *JSONSchemaEditorControl) Required(value interface{}) *JSONSchemaEditorControl {
    a.Set("required", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *JSONSchemaEditorControl) Mode(value interface{}) *JSONSchemaEditorControl {
    a.Set("mode", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *JSONSchemaEditorControl) Id(value interface{}) *JSONSchemaEditorControl {
    a.Set("id", value)
    return a
}
