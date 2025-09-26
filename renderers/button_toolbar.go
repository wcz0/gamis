package renderers


/**
 * Button Toolar 渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/button-toolbar

 * @author wcz0
 * @version 6.2.2
 */
type ButtonToolbar struct {
	*BaseRenderer
}

func NewButtonToolbar() *ButtonToolbar {
    a := &ButtonToolbar{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "button-toolbar")
    return a
}


func (a *ButtonToolbar) Set(name string, value interface{}) *ButtonToolbar {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * staticPlaceholder
 */
func (a *ButtonToolbar) StaticPlaceholder(value interface{}) *ButtonToolbar {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * editorSetting
 */
func (a *ButtonToolbar) EditorSetting(value interface{}) *ButtonToolbar {
    a.Set("editorSetting", value)
    return a
}

/**
 * useMobileUI
 */
func (a *ButtonToolbar) UseMobileUI(value interface{}) *ButtonToolbar {
    a.Set("useMobileUI", value)
    return a
}

/**
 * testid
 */
func (a *ButtonToolbar) Testid(value interface{}) *ButtonToolbar {
    a.Set("testid", value)
    return a
}

/**
 * labelWidth
 */
func (a *ButtonToolbar) LabelWidth(value interface{}) *ButtonToolbar {
    a.Set("labelWidth", value)
    return a
}

/**
 * hint
 */
func (a *ButtonToolbar) Hint(value interface{}) *ButtonToolbar {
    a.Set("hint", value)
    return a
}

/**
 * submitOnChange
 */
func (a *ButtonToolbar) SubmitOnChange(value interface{}) *ButtonToolbar {
    a.Set("submitOnChange", value)
    return a
}

/**
 * readOnly
 */
func (a *ButtonToolbar) ReadOnly(value interface{}) *ButtonToolbar {
    a.Set("readOnly", value)
    return a
}

/**
 * validations
 */
func (a *ButtonToolbar) Validations(value interface{}) *ButtonToolbar {
    a.Set("validations", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *ButtonToolbar) ClearValueOnHidden(value interface{}) *ButtonToolbar {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * row
 */
func (a *ButtonToolbar) Row(value interface{}) *ButtonToolbar {
    a.Set("row", value)
    return a
}

/**
 * horizontal
 */
func (a *ButtonToolbar) Horizontal(value interface{}) *ButtonToolbar {
    a.Set("horizontal", value)
    return a
}

/**
 * hiddenOn
 */
func (a *ButtonToolbar) HiddenOn(value interface{}) *ButtonToolbar {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visibleOn
 */
func (a *ButtonToolbar) VisibleOn(value interface{}) *ButtonToolbar {
    a.Set("visibleOn", value)
    return a
}

/**
 * id
 */
func (a *ButtonToolbar) Id(value interface{}) *ButtonToolbar {
    a.Set("id", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *ButtonToolbar) StaticLabelClassName(value interface{}) *ButtonToolbar {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * buttons
 */
func (a *ButtonToolbar) Buttons(value interface{}) *ButtonToolbar {
    a.Set("buttons", value)
    return a
}

/**
 * desc
 */
func (a *ButtonToolbar) Desc(value interface{}) *ButtonToolbar {
    a.Set("desc", value)
    return a
}

/**
 * mode
 */
func (a *ButtonToolbar) Mode(value interface{}) *ButtonToolbar {
    a.Set("mode", value)
    return a
}

/**
 * inline
 */
func (a *ButtonToolbar) Inline(value interface{}) *ButtonToolbar {
    a.Set("inline", value)
    return a
}

/**
 * inputClassName
 */
func (a *ButtonToolbar) InputClassName(value interface{}) *ButtonToolbar {
    a.Set("inputClassName", value)
    return a
}

/**
 * style
 */
func (a *ButtonToolbar) Style(value interface{}) *ButtonToolbar {
    a.Set("style", value)
    return a
}

/**
 * remark
 */
func (a *ButtonToolbar) Remark(value interface{}) *ButtonToolbar {
    a.Set("remark", value)
    return a
}

/**
 * value
 */
func (a *ButtonToolbar) Value(value interface{}) *ButtonToolbar {
    a.Set("value", value)
    return a
}

/**
 * visible
 */
func (a *ButtonToolbar) Visible(value interface{}) *ButtonToolbar {
    a.Set("visible", value)
    return a
}

/**
 * staticClassName
 */
func (a *ButtonToolbar) StaticClassName(value interface{}) *ButtonToolbar {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *ButtonToolbar) StaticSchema(value interface{}) *ButtonToolbar {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *ButtonToolbar) Type(value interface{}) *ButtonToolbar {
    a.Set("type", value)
    return a
}

/**
 * validateApi
 */
func (a *ButtonToolbar) ValidateApi(value interface{}) *ButtonToolbar {
    a.Set("validateApi", value)
    return a
}

/**
 * name
 */
func (a *ButtonToolbar) Name(value interface{}) *ButtonToolbar {
    a.Set("name", value)
    return a
}

/**
 * labelRemark
 */
func (a *ButtonToolbar) LabelRemark(value interface{}) *ButtonToolbar {
    a.Set("labelRemark", value)
    return a
}

/**
 * static
 */
func (a *ButtonToolbar) Static(value interface{}) *ButtonToolbar {
    a.Set("static", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *ButtonToolbar) StaticInputClassName(value interface{}) *ButtonToolbar {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * label
 */
func (a *ButtonToolbar) Label(value interface{}) *ButtonToolbar {
    a.Set("label", value)
    return a
}

/**
 * labelAlign
 */
func (a *ButtonToolbar) LabelAlign(value interface{}) *ButtonToolbar {
    a.Set("labelAlign", value)
    return a
}

/**
 * labelClassName
 */
func (a *ButtonToolbar) LabelClassName(value interface{}) *ButtonToolbar {
    a.Set("labelClassName", value)
    return a
}

/**
 * extraName
 */
func (a *ButtonToolbar) ExtraName(value interface{}) *ButtonToolbar {
    a.Set("extraName", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *ButtonToolbar) ReadOnlyOn(value interface{}) *ButtonToolbar {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * description
 */
func (a *ButtonToolbar) Description(value interface{}) *ButtonToolbar {
    a.Set("description", value)
    return a
}

/**
 * placeholder
 */
func (a *ButtonToolbar) Placeholder(value interface{}) *ButtonToolbar {
    a.Set("placeholder", value)
    return a
}

/**
 * validationErrors
 */
func (a *ButtonToolbar) ValidationErrors(value interface{}) *ButtonToolbar {
    a.Set("validationErrors", value)
    return a
}

/**
 * labelOverflow
 */
func (a *ButtonToolbar) LabelOverflow(value interface{}) *ButtonToolbar {
    a.Set("labelOverflow", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *ButtonToolbar) DescriptionClassName(value interface{}) *ButtonToolbar {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * required
 */
func (a *ButtonToolbar) Required(value interface{}) *ButtonToolbar {
    a.Set("required", value)
    return a
}

/**
 * autoFill
 */
func (a *ButtonToolbar) AutoFill(value interface{}) *ButtonToolbar {
    a.Set("autoFill", value)
    return a
}

/**
 * initAutoFill
 */
func (a *ButtonToolbar) InitAutoFill(value interface{}) *ButtonToolbar {
    a.Set("initAutoFill", value)
    return a
}

/**
 * className
 */
func (a *ButtonToolbar) ClassName(value interface{}) *ButtonToolbar {
    a.Set("className", value)
    return a
}

/**
 * disabledOn
 */
func (a *ButtonToolbar) DisabledOn(value interface{}) *ButtonToolbar {
    a.Set("disabledOn", value)
    return a
}

/**
 * hidden
 */
func (a *ButtonToolbar) Hidden(value interface{}) *ButtonToolbar {
    a.Set("hidden", value)
    return a
}

/**
 * validateOnChange
 */
func (a *ButtonToolbar) ValidateOnChange(value interface{}) *ButtonToolbar {
    a.Set("validateOnChange", value)
    return a
}

/**
 * disabled
 */
func (a *ButtonToolbar) Disabled(value interface{}) *ButtonToolbar {
    a.Set("disabled", value)
    return a
}

/**
 * onEvent
 */
func (a *ButtonToolbar) OnEvent(value interface{}) *ButtonToolbar {
    a.Set("onEvent", value)
    return a
}

/**
 * staticOn
 */
func (a *ButtonToolbar) StaticOn(value interface{}) *ButtonToolbar {
    a.Set("staticOn", value)
    return a
}
