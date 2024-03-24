package renderers


/**
 * Editor 代码编辑器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/editor

 * @author wcz0
 * @version 6.2.2
 */
type EditorControl struct {
	*BaseRenderer
}

func NewEditorControl() *EditorControl {
    a := &EditorControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "editor")
    return a
}
/**
 * 在Table中调整宽度
 */
func (a *EditorControl) Width(value interface{}) *EditorControl {
    a.Set("width", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *EditorControl) Horizontal(value interface{}) *EditorControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 可选值: editor | bat-editor | c-editor | coffeescript-editor | cpp-editor | csharp-editor | css-editor | dockerfile-editor | fsharp-editor | go-editor | handlebars-editor | html-editor | ini-editor | java-editor | javascript-editor | json-editor | less-editor | lua-editor | markdown-editor | msdax-editor | objective-c-editor | php-editor | plaintext-editor | postiats-editor | powershell-editor | pug-editor | python-editor | r-editor | razor-editor | ruby-editor | sb-editor | scss-editor | sol-editor | sql-editor | swift-editor | typescript-editor | vb-editor | xml-editor | yaml-editor
 */
func (a *EditorControl) Type(value interface{}) *EditorControl {
    a.Set("type", value)
    return a
}

/**
 * 获取编辑器底层实例
 */
func (a *EditorControl) EditorDidMount(value interface{}) *EditorControl {
    a.Set("editorDidMount", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *EditorControl) ClearValueOnHidden(value interface{}) *EditorControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *EditorControl) Static(value interface{}) *EditorControl {
    a.Set("static", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *EditorControl) UseMobileUI(value interface{}) *EditorControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *EditorControl) ExtraName(value interface{}) *EditorControl {
    a.Set("extraName", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *EditorControl) Remark(value interface{}) *EditorControl {
    a.Set("remark", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *EditorControl) LabelRemark(value interface{}) *EditorControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *EditorControl) ValidateOnChange(value interface{}) *EditorControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 占位符
 */
func (a *EditorControl) Placeholder(value interface{}) *EditorControl {
    a.Set("placeholder", value)
    return a
}

/**
 */
func (a *EditorControl) Validations(value interface{}) *EditorControl {
    a.Set("validations", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *EditorControl) EditorSetting(value interface{}) *EditorControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 是否展示全屏模式开关
 */
func (a *EditorControl) AllowFullscreen(value interface{}) *EditorControl {
    a.Set("allowFullscreen", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *EditorControl) HiddenOn(value interface{}) *EditorControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *EditorControl) VisibleOn(value interface{}) *EditorControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *EditorControl) OnEvent(value interface{}) *EditorControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 组件样式
 */
func (a *EditorControl) Style(value interface{}) *EditorControl {
    a.Set("style", value)
    return a
}

/**
 * 语言类型
 * 可选值: bat | c | coffeescript | cpp | csharp | css | dockerfile | fsharp | go | handlebars | html | ini | java | javascript | json | less | lua | markdown | msdax | objective-c | php | plaintext | postiats | powershell | pug | python | r | razor | ruby | sb | scss | shell | sol | sql | swift | typescript | vb | xml | yaml
 */
func (a *EditorControl) Language(value interface{}) *EditorControl {
    a.Set("language", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *EditorControl) Hidden(value interface{}) *EditorControl {
    a.Set("hidden", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *EditorControl) ValidationErrors(value interface{}) *EditorControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *EditorControl) Value(value interface{}) *EditorControl {
    a.Set("value", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *EditorControl) ClassName(value interface{}) *EditorControl {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用
 */
func (a *EditorControl) Disabled(value interface{}) *EditorControl {
    a.Set("disabled", value)
    return a
}

/**
 */
func (a *EditorControl) Desc(value interface{}) *EditorControl {
    a.Set("desc", value)
    return a
}

/**
 * 描述标题
 */
func (a *EditorControl) LabelAlign(value interface{}) *EditorControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *EditorControl) Hint(value interface{}) *EditorControl {
    a.Set("hint", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *EditorControl) StaticInputClassName(value interface{}) *EditorControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *EditorControl) DisabledOn(value interface{}) *EditorControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *EditorControl) StaticLabelClassName(value interface{}) *EditorControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *EditorControl) Mode(value interface{}) *EditorControl {
    a.Set("mode", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *EditorControl) ValidateApi(value interface{}) *EditorControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 是否显示
 */
func (a *EditorControl) Visible(value interface{}) *EditorControl {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *EditorControl) StaticClassName(value interface{}) *EditorControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 编辑器大小
 * 可选值: sm | md | lg | xl | xxl
 */
func (a *EditorControl) Size(value interface{}) *EditorControl {
    a.Set("size", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *EditorControl) Id(value interface{}) *EditorControl {
    a.Set("id", value)
    return a
}

/**
 */
func (a *EditorControl) StaticSchema(value interface{}) *EditorControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *EditorControl) LabelWidth(value interface{}) *EditorControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *EditorControl) Name(value interface{}) *EditorControl {
    a.Set("name", value)
    return a
}

/**
 * 只读条件
 */
func (a *EditorControl) ReadOnlyOn(value interface{}) *EditorControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *EditorControl) StaticOn(value interface{}) *EditorControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *EditorControl) StaticPlaceholder(value interface{}) *EditorControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *EditorControl) SubmitOnChange(value interface{}) *EditorControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *EditorControl) Inline(value interface{}) *EditorControl {
    a.Set("inline", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *EditorControl) DescriptionClassName(value interface{}) *EditorControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 描述标题
 */
func (a *EditorControl) Label(value interface{}) *EditorControl {
    a.Set("label", value)
    return a
}

/**
 * 配置 label className
 */
func (a *EditorControl) LabelClassName(value interface{}) *EditorControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *EditorControl) Description(value interface{}) *EditorControl {
    a.Set("description", value)
    return a
}

/**
 * 是否为必填
 */
func (a *EditorControl) Required(value interface{}) *EditorControl {
    a.Set("required", value)
    return a
}

/**
 * 是否只读
 */
func (a *EditorControl) ReadOnly(value interface{}) *EditorControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 配置 input className
 */
func (a *EditorControl) InputClassName(value interface{}) *EditorControl {
    a.Set("inputClassName", value)
    return a
}
