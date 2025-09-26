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


func (a *EditorControl) Set(name string, value interface{}) *EditorControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * placeholder
 */
func (a *EditorControl) Placeholder(value interface{}) *EditorControl {
    a.Set("placeholder", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *EditorControl) StaticInputClassName(value interface{}) *EditorControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * size
 */
func (a *EditorControl) Size(value interface{}) *EditorControl {
    a.Set("size", value)
    return a
}

/**
 * desc
 */
func (a *EditorControl) Desc(value interface{}) *EditorControl {
    a.Set("desc", value)
    return a
}

/**
 * mode
 */
func (a *EditorControl) Mode(value interface{}) *EditorControl {
    a.Set("mode", value)
    return a
}

/**
 * static
 */
func (a *EditorControl) Static(value interface{}) *EditorControl {
    a.Set("static", value)
    return a
}

/**
 * staticOn
 */
func (a *EditorControl) StaticOn(value interface{}) *EditorControl {
    a.Set("staticOn", value)
    return a
}

/**
 * label
 */
func (a *EditorControl) Label(value interface{}) *EditorControl {
    a.Set("label", value)
    return a
}

/**
 * labelOverflow
 */
func (a *EditorControl) LabelOverflow(value interface{}) *EditorControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * autoFill
 */
func (a *EditorControl) AutoFill(value interface{}) *EditorControl {
    a.Set("autoFill", value)
    return a
}

/**
 * hiddenOn
 */
func (a *EditorControl) HiddenOn(value interface{}) *EditorControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * labelAlign
 */
func (a *EditorControl) LabelAlign(value interface{}) *EditorControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * hint
 */
func (a *EditorControl) Hint(value interface{}) *EditorControl {
    a.Set("hint", value)
    return a
}

/**
 * value
 */
func (a *EditorControl) Value(value interface{}) *EditorControl {
    a.Set("value", value)
    return a
}

/**
 * validateApi
 */
func (a *EditorControl) ValidateApi(value interface{}) *EditorControl {
    a.Set("validateApi", value)
    return a
}

/**
 * editorSetting
 */
func (a *EditorControl) EditorSetting(value interface{}) *EditorControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * editorDidMount
 */
func (a *EditorControl) EditorDidMount(value interface{}) *EditorControl {
    a.Set("editorDidMount", value)
    return a
}

/**
 * labelWidth
 */
func (a *EditorControl) LabelWidth(value interface{}) *EditorControl {
    a.Set("labelWidth", value)
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
 * labelClassName
 */
func (a *EditorControl) LabelClassName(value interface{}) *EditorControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * name
 */
func (a *EditorControl) Name(value interface{}) *EditorControl {
    a.Set("name", value)
    return a
}

/**
 * extraName
 */
func (a *EditorControl) ExtraName(value interface{}) *EditorControl {
    a.Set("extraName", value)
    return a
}

/**
 * labelRemark
 */
func (a *EditorControl) LabelRemark(value interface{}) *EditorControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * required
 */
func (a *EditorControl) Required(value interface{}) *EditorControl {
    a.Set("required", value)
    return a
}

/**
 * initAutoFill
 */
func (a *EditorControl) InitAutoFill(value interface{}) *EditorControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * disabled
 */
func (a *EditorControl) Disabled(value interface{}) *EditorControl {
    a.Set("disabled", value)
    return a
}

/**
 * visibleOn
 */
func (a *EditorControl) VisibleOn(value interface{}) *EditorControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * readOnly
 */
func (a *EditorControl) ReadOnly(value interface{}) *EditorControl {
    a.Set("readOnly", value)
    return a
}

/**
 * description
 */
func (a *EditorControl) Description(value interface{}) *EditorControl {
    a.Set("description", value)
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
 * staticPlaceholder
 */
func (a *EditorControl) StaticPlaceholder(value interface{}) *EditorControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * remark
 */
func (a *EditorControl) Remark(value interface{}) *EditorControl {
    a.Set("remark", value)
    return a
}

/**
 * validateOnChange
 */
func (a *EditorControl) ValidateOnChange(value interface{}) *EditorControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *EditorControl) DescriptionClassName(value interface{}) *EditorControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * validations
 */
func (a *EditorControl) Validations(value interface{}) *EditorControl {
    a.Set("validations", value)
    return a
}

/**
 * className
 */
func (a *EditorControl) ClassName(value interface{}) *EditorControl {
    a.Set("className", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *EditorControl) StaticLabelClassName(value interface{}) *EditorControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * style
 */
func (a *EditorControl) Style(value interface{}) *EditorControl {
    a.Set("style", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *EditorControl) ClearValueOnHidden(value interface{}) *EditorControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * row
 */
func (a *EditorControl) Row(value interface{}) *EditorControl {
    a.Set("row", value)
    return a
}

/**
 * id
 */
func (a *EditorControl) Id(value interface{}) *EditorControl {
    a.Set("id", value)
    return a
}

/**
 * onEvent
 */
func (a *EditorControl) OnEvent(value interface{}) *EditorControl {
    a.Set("onEvent", value)
    return a
}

/**
 * horizontal
 */
func (a *EditorControl) Horizontal(value interface{}) *EditorControl {
    a.Set("horizontal", value)
    return a
}

/**
 * disabledOn
 */
func (a *EditorControl) DisabledOn(value interface{}) *EditorControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * visible
 */
func (a *EditorControl) Visible(value interface{}) *EditorControl {
    a.Set("visible", value)
    return a
}

/**
 * language
 */
func (a *EditorControl) Language(value interface{}) *EditorControl {
    a.Set("language", value)
    return a
}

/**
 * submitOnChange
 */
func (a *EditorControl) SubmitOnChange(value interface{}) *EditorControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * inline
 */
func (a *EditorControl) Inline(value interface{}) *EditorControl {
    a.Set("inline", value)
    return a
}

/**
 * staticClassName
 */
func (a *EditorControl) StaticClassName(value interface{}) *EditorControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *EditorControl) ReadOnlyOn(value interface{}) *EditorControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * inputClassName
 */
func (a *EditorControl) InputClassName(value interface{}) *EditorControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * validationErrors
 */
func (a *EditorControl) ValidationErrors(value interface{}) *EditorControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * hidden
 */
func (a *EditorControl) Hidden(value interface{}) *EditorControl {
    a.Set("hidden", value)
    return a
}

/**
 * allowFullscreen
 */
func (a *EditorControl) AllowFullscreen(value interface{}) *EditorControl {
    a.Set("allowFullscreen", value)
    return a
}

/**
 * staticSchema
 */
func (a *EditorControl) StaticSchema(value interface{}) *EditorControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * useMobileUI
 */
func (a *EditorControl) UseMobileUI(value interface{}) *EditorControl {
    a.Set("useMobileUI", value)
    return a
}
