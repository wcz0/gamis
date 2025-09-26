package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type FormBaseControl struct {
	*BaseRenderer
}

func NewFormBaseControl() *FormBaseControl {
    a := &FormBaseControl{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}


func (a *FormBaseControl) Set(name string, value interface{}) *FormBaseControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * hidden
 */
func (a *FormBaseControl) Hidden(value interface{}) *FormBaseControl {
    a.Set("hidden", value)
    return a
}

/**
 * staticOn
 */
func (a *FormBaseControl) StaticOn(value interface{}) *FormBaseControl {
    a.Set("staticOn", value)
    return a
}

/**
 * name
 */
func (a *FormBaseControl) Name(value interface{}) *FormBaseControl {
    a.Set("name", value)
    return a
}

/**
 * labelRemark
 */
func (a *FormBaseControl) LabelRemark(value interface{}) *FormBaseControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * hint
 */
func (a *FormBaseControl) Hint(value interface{}) *FormBaseControl {
    a.Set("hint", value)
    return a
}

/**
 * validateOnChange
 */
func (a *FormBaseControl) ValidateOnChange(value interface{}) *FormBaseControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * hiddenOn
 */
func (a *FormBaseControl) HiddenOn(value interface{}) *FormBaseControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *FormBaseControl) StaticPlaceholder(value interface{}) *FormBaseControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * style
 */
func (a *FormBaseControl) Style(value interface{}) *FormBaseControl {
    a.Set("style", value)
    return a
}

/**
 * editorSetting
 */
func (a *FormBaseControl) EditorSetting(value interface{}) *FormBaseControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * label
 */
func (a *FormBaseControl) Label(value interface{}) *FormBaseControl {
    a.Set("label", value)
    return a
}

/**
 * labelClassName
 */
func (a *FormBaseControl) LabelClassName(value interface{}) *FormBaseControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * validationErrors
 */
func (a *FormBaseControl) ValidationErrors(value interface{}) *FormBaseControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * value
 */
func (a *FormBaseControl) Value(value interface{}) *FormBaseControl {
    a.Set("value", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *FormBaseControl) ClearValueOnHidden(value interface{}) *FormBaseControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * initAutoFill
 */
func (a *FormBaseControl) InitAutoFill(value interface{}) *FormBaseControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * disabled
 */
func (a *FormBaseControl) Disabled(value interface{}) *FormBaseControl {
    a.Set("disabled", value)
    return a
}

/**
 * labelWidth
 */
func (a *FormBaseControl) LabelWidth(value interface{}) *FormBaseControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *FormBaseControl) ReadOnlyOn(value interface{}) *FormBaseControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * description
 */
func (a *FormBaseControl) Description(value interface{}) *FormBaseControl {
    a.Set("description", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *FormBaseControl) DescriptionClassName(value interface{}) *FormBaseControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * autoFill
 */
func (a *FormBaseControl) AutoFill(value interface{}) *FormBaseControl {
    a.Set("autoFill", value)
    return a
}

/**
 * visible
 */
func (a *FormBaseControl) Visible(value interface{}) *FormBaseControl {
    a.Set("visible", value)
    return a
}

/**
 * id
 */
func (a *FormBaseControl) Id(value interface{}) *FormBaseControl {
    a.Set("id", value)
    return a
}

/**
 * onEvent
 */
func (a *FormBaseControl) OnEvent(value interface{}) *FormBaseControl {
    a.Set("onEvent", value)
    return a
}

/**
 * readOnly
 */
func (a *FormBaseControl) ReadOnly(value interface{}) *FormBaseControl {
    a.Set("readOnly", value)
    return a
}

/**
 * desc
 */
func (a *FormBaseControl) Desc(value interface{}) *FormBaseControl {
    a.Set("desc", value)
    return a
}

/**
 * mode
 */
func (a *FormBaseControl) Mode(value interface{}) *FormBaseControl {
    a.Set("mode", value)
    return a
}

/**
 * className
 */
func (a *FormBaseControl) ClassName(value interface{}) *FormBaseControl {
    a.Set("className", value)
    return a
}

/**
 * visibleOn
 */
func (a *FormBaseControl) VisibleOn(value interface{}) *FormBaseControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * useMobileUI
 */
func (a *FormBaseControl) UseMobileUI(value interface{}) *FormBaseControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * type
 */
func (a *FormBaseControl) Type(value interface{}) *FormBaseControl {
    a.Set("type", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *FormBaseControl) Width(value interface{}) *FormBaseControl {
    a.Set("width", value)
    return a
}

/**
 * submitOnChange
 */
func (a *FormBaseControl) SubmitOnChange(value interface{}) *FormBaseControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * required
 */
func (a *FormBaseControl) Required(value interface{}) *FormBaseControl {
    a.Set("required", value)
    return a
}

/**
 * validateApi
 */
func (a *FormBaseControl) ValidateApi(value interface{}) *FormBaseControl {
    a.Set("validateApi", value)
    return a
}

/**
 * disabledOn
 */
func (a *FormBaseControl) DisabledOn(value interface{}) *FormBaseControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * static
 */
func (a *FormBaseControl) Static(value interface{}) *FormBaseControl {
    a.Set("static", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *FormBaseControl) StaticLabelClassName(value interface{}) *FormBaseControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *FormBaseControl) StaticInputClassName(value interface{}) *FormBaseControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *FormBaseControl) StaticSchema(value interface{}) *FormBaseControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * validations
 */
func (a *FormBaseControl) Validations(value interface{}) *FormBaseControl {
    a.Set("validations", value)
    return a
}

/**
 * size
 */
func (a *FormBaseControl) Size(value interface{}) *FormBaseControl {
    a.Set("size", value)
    return a
}

/**
 * horizontal
 */
func (a *FormBaseControl) Horizontal(value interface{}) *FormBaseControl {
    a.Set("horizontal", value)
    return a
}

/**
 * inputClassName
 */
func (a *FormBaseControl) InputClassName(value interface{}) *FormBaseControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * row
 */
func (a *FormBaseControl) Row(value interface{}) *FormBaseControl {
    a.Set("row", value)
    return a
}

/**
 * staticClassName
 */
func (a *FormBaseControl) StaticClassName(value interface{}) *FormBaseControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * labelAlign
 */
func (a *FormBaseControl) LabelAlign(value interface{}) *FormBaseControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * labelOverflow
 */
func (a *FormBaseControl) LabelOverflow(value interface{}) *FormBaseControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * extraName
 */
func (a *FormBaseControl) ExtraName(value interface{}) *FormBaseControl {
    a.Set("extraName", value)
    return a
}

/**
 * remark
 */
func (a *FormBaseControl) Remark(value interface{}) *FormBaseControl {
    a.Set("remark", value)
    return a
}

/**
 * inline
 */
func (a *FormBaseControl) Inline(value interface{}) *FormBaseControl {
    a.Set("inline", value)
    return a
}

/**
 * placeholder
 */
func (a *FormBaseControl) Placeholder(value interface{}) *FormBaseControl {
    a.Set("placeholder", value)
    return a
}
