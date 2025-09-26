package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type InputDateBaseControl struct {
	*BaseRenderer
}

func NewInputDateBaseControl() *InputDateBaseControl {
    a := &InputDateBaseControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-date")
    return a
}


func (a *InputDateBaseControl) Set(name string, value interface{}) *InputDateBaseControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * validateApi
 */
func (a *InputDateBaseControl) ValidateApi(value interface{}) *InputDateBaseControl {
    a.Set("validateApi", value)
    return a
}

/**
 * value
 */
func (a *InputDateBaseControl) Value(value interface{}) *InputDateBaseControl {
    a.Set("value", value)
    return a
}

/**
 * onEvent
 */
func (a *InputDateBaseControl) OnEvent(value interface{}) *InputDateBaseControl {
    a.Set("onEvent", value)
    return a
}

/**
 * row
 */
func (a *InputDateBaseControl) Row(value interface{}) *InputDateBaseControl {
    a.Set("row", value)
    return a
}

/**
 * disabledOn
 */
func (a *InputDateBaseControl) DisabledOn(value interface{}) *InputDateBaseControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticOn
 */
func (a *InputDateBaseControl) StaticOn(value interface{}) *InputDateBaseControl {
    a.Set("staticOn", value)
    return a
}

/**
 * style
 */
func (a *InputDateBaseControl) Style(value interface{}) *InputDateBaseControl {
    a.Set("style", value)
    return a
}

/**
 * editorSetting
 */
func (a *InputDateBaseControl) EditorSetting(value interface{}) *InputDateBaseControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * extraName
 */
func (a *InputDateBaseControl) ExtraName(value interface{}) *InputDateBaseControl {
    a.Set("extraName", value)
    return a
}

/**
 * emebed
 */
func (a *InputDateBaseControl) Emebed(value interface{}) *InputDateBaseControl {
    a.Set("emebed", value)
    return a
}

/**
 * validations
 */
func (a *InputDateBaseControl) Validations(value interface{}) *InputDateBaseControl {
    a.Set("validations", value)
    return a
}

/**
 * id
 */
func (a *InputDateBaseControl) Id(value interface{}) *InputDateBaseControl {
    a.Set("id", value)
    return a
}

/**
 * labelWidth
 */
func (a *InputDateBaseControl) LabelWidth(value interface{}) *InputDateBaseControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * name
 */
func (a *InputDateBaseControl) Name(value interface{}) *InputDateBaseControl {
    a.Set("name", value)
    return a
}

/**
 * hiddenOn
 */
func (a *InputDateBaseControl) HiddenOn(value interface{}) *InputDateBaseControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *InputDateBaseControl) StaticInputClassName(value interface{}) *InputDateBaseControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *InputDateBaseControl) Width(value interface{}) *InputDateBaseControl {
    a.Set("width", value)
    return a
}

/**
 * labelAlign
 */
func (a *InputDateBaseControl) LabelAlign(value interface{}) *InputDateBaseControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *InputDateBaseControl) DescriptionClassName(value interface{}) *InputDateBaseControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * disabled
 */
func (a *InputDateBaseControl) Disabled(value interface{}) *InputDateBaseControl {
    a.Set("disabled", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *InputDateBaseControl) StaticLabelClassName(value interface{}) *InputDateBaseControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * remark
 */
func (a *InputDateBaseControl) Remark(value interface{}) *InputDateBaseControl {
    a.Set("remark", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *InputDateBaseControl) ReadOnlyOn(value interface{}) *InputDateBaseControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * autoFill
 */
func (a *InputDateBaseControl) AutoFill(value interface{}) *InputDateBaseControl {
    a.Set("autoFill", value)
    return a
}

/**
 * clearable
 */
func (a *InputDateBaseControl) Clearable(value interface{}) *InputDateBaseControl {
    a.Set("clearable", value)
    return a
}

/**
 * utc
 */
func (a *InputDateBaseControl) Utc(value interface{}) *InputDateBaseControl {
    a.Set("utc", value)
    return a
}

/**
 * labelClassName
 */
func (a *InputDateBaseControl) LabelClassName(value interface{}) *InputDateBaseControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * submitOnChange
 */
func (a *InputDateBaseControl) SubmitOnChange(value interface{}) *InputDateBaseControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * readOnly
 */
func (a *InputDateBaseControl) ReadOnly(value interface{}) *InputDateBaseControl {
    a.Set("readOnly", value)
    return a
}

/**
 * mode
 */
func (a *InputDateBaseControl) Mode(value interface{}) *InputDateBaseControl {
    a.Set("mode", value)
    return a
}

/**
 * placeholder
 */
func (a *InputDateBaseControl) Placeholder(value interface{}) *InputDateBaseControl {
    a.Set("placeholder", value)
    return a
}

/**
 * className
 */
func (a *InputDateBaseControl) ClassName(value interface{}) *InputDateBaseControl {
    a.Set("className", value)
    return a
}

/**
 * staticSchema
 */
func (a *InputDateBaseControl) StaticSchema(value interface{}) *InputDateBaseControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * format
 */
func (a *InputDateBaseControl) Format(value interface{}) *InputDateBaseControl {
    a.Set("format", value)
    return a
}

/**
 * hint
 */
func (a *InputDateBaseControl) Hint(value interface{}) *InputDateBaseControl {
    a.Set("hint", value)
    return a
}

/**
 * desc
 */
func (a *InputDateBaseControl) Desc(value interface{}) *InputDateBaseControl {
    a.Set("desc", value)
    return a
}

/**
 * inputClassName
 */
func (a *InputDateBaseControl) InputClassName(value interface{}) *InputDateBaseControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * validationErrors
 */
func (a *InputDateBaseControl) ValidationErrors(value interface{}) *InputDateBaseControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * displayFormat
 */
func (a *InputDateBaseControl) DisplayFormat(value interface{}) *InputDateBaseControl {
    a.Set("displayFormat", value)
    return a
}

/**
 * borderMode
 */
func (a *InputDateBaseControl) BorderMode(value interface{}) *InputDateBaseControl {
    a.Set("borderMode", value)
    return a
}

/**
 * shortcuts
 */
func (a *InputDateBaseControl) Shortcuts(value interface{}) *InputDateBaseControl {
    a.Set("shortcuts", value)
    return a
}

/**
 * inputForbid
 */
func (a *InputDateBaseControl) InputForbid(value interface{}) *InputDateBaseControl {
    a.Set("inputForbid", value)
    return a
}

/**
 * inline
 */
func (a *InputDateBaseControl) Inline(value interface{}) *InputDateBaseControl {
    a.Set("inline", value)
    return a
}

/**
 * inputFormat
 */
func (a *InputDateBaseControl) InputFormat(value interface{}) *InputDateBaseControl {
    a.Set("inputFormat", value)
    return a
}

/**
 * labelRemark
 */
func (a *InputDateBaseControl) LabelRemark(value interface{}) *InputDateBaseControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * required
 */
func (a *InputDateBaseControl) Required(value interface{}) *InputDateBaseControl {
    a.Set("required", value)
    return a
}

/**
 * 可选值: input-date | input-datetime | input-time | input-month | input-quarter | input-year
 */
func (a *InputDateBaseControl) Type(value interface{}) *InputDateBaseControl {
    a.Set("type", value)
    return a
}

/**
 * label
 */
func (a *InputDateBaseControl) Label(value interface{}) *InputDateBaseControl {
    a.Set("label", value)
    return a
}

/**
 * static
 */
func (a *InputDateBaseControl) Static(value interface{}) *InputDateBaseControl {
    a.Set("static", value)
    return a
}

/**
 * useMobileUI
 */
func (a *InputDateBaseControl) UseMobileUI(value interface{}) *InputDateBaseControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * valueFormat
 */
func (a *InputDateBaseControl) ValueFormat(value interface{}) *InputDateBaseControl {
    a.Set("valueFormat", value)
    return a
}

/**
 * validateOnChange
 */
func (a *InputDateBaseControl) ValidateOnChange(value interface{}) *InputDateBaseControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * visible
 */
func (a *InputDateBaseControl) Visible(value interface{}) *InputDateBaseControl {
    a.Set("visible", value)
    return a
}

/**
 * description
 */
func (a *InputDateBaseControl) Description(value interface{}) *InputDateBaseControl {
    a.Set("description", value)
    return a
}

/**
 * hidden
 */
func (a *InputDateBaseControl) Hidden(value interface{}) *InputDateBaseControl {
    a.Set("hidden", value)
    return a
}

/**
 * labelOverflow
 */
func (a *InputDateBaseControl) LabelOverflow(value interface{}) *InputDateBaseControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * horizontal
 */
func (a *InputDateBaseControl) Horizontal(value interface{}) *InputDateBaseControl {
    a.Set("horizontal", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *InputDateBaseControl) ClearValueOnHidden(value interface{}) *InputDateBaseControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * visibleOn
 */
func (a *InputDateBaseControl) VisibleOn(value interface{}) *InputDateBaseControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * disabledDate
 */
func (a *InputDateBaseControl) DisabledDate(value interface{}) *InputDateBaseControl {
    a.Set("disabledDate", value)
    return a
}

/**
 * initAutoFill
 */
func (a *InputDateBaseControl) InitAutoFill(value interface{}) *InputDateBaseControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *InputDateBaseControl) StaticPlaceholder(value interface{}) *InputDateBaseControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticClassName
 */
func (a *InputDateBaseControl) StaticClassName(value interface{}) *InputDateBaseControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * size
 */
func (a *InputDateBaseControl) Size(value interface{}) *InputDateBaseControl {
    a.Set("size", value)
    return a
}
