package renderers


/**
 * 年份选择控件

 * @author wcz0
 * @version 6.2.2
 */
type YearControl struct {
	*BaseRenderer
}

func NewYearControl() *YearControl {
    a := &YearControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-year")
    return a
}


func (a *YearControl) Set(name string, value interface{}) *YearControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * submitOnChange
 */
func (a *YearControl) SubmitOnChange(value interface{}) *YearControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *YearControl) DescriptionClassName(value interface{}) *YearControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * utc
 */
func (a *YearControl) Utc(value interface{}) *YearControl {
    a.Set("utc", value)
    return a
}

/**
 * labelAlign
 */
func (a *YearControl) LabelAlign(value interface{}) *YearControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * horizontal
 */
func (a *YearControl) Horizontal(value interface{}) *YearControl {
    a.Set("horizontal", value)
    return a
}

/**
 * inline
 */
func (a *YearControl) Inline(value interface{}) *YearControl {
    a.Set("inline", value)
    return a
}

/**
 * initAutoFill
 */
func (a *YearControl) InitAutoFill(value interface{}) *YearControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * disabledOn
 */
func (a *YearControl) DisabledOn(value interface{}) *YearControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticOn
 */
func (a *YearControl) StaticOn(value interface{}) *YearControl {
    a.Set("staticOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *YearControl) StaticInputClassName(value interface{}) *YearControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * description
 */
func (a *YearControl) Description(value interface{}) *YearControl {
    a.Set("description", value)
    return a
}

/**
 * autoFill
 */
func (a *YearControl) AutoFill(value interface{}) *YearControl {
    a.Set("autoFill", value)
    return a
}

/**
 * className
 */
func (a *YearControl) ClassName(value interface{}) *YearControl {
    a.Set("className", value)
    return a
}

/**
 * disabled
 */
func (a *YearControl) Disabled(value interface{}) *YearControl {
    a.Set("disabled", value)
    return a
}

/**
 * staticSchema
 */
func (a *YearControl) StaticSchema(value interface{}) *YearControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * size
 */
func (a *YearControl) Size(value interface{}) *YearControl {
    a.Set("size", value)
    return a
}

/**
 * emebed
 */
func (a *YearControl) Emebed(value interface{}) *YearControl {
    a.Set("emebed", value)
    return a
}

/**
 * shortcuts
 */
func (a *YearControl) Shortcuts(value interface{}) *YearControl {
    a.Set("shortcuts", value)
    return a
}

/**
 * editorSetting
 */
func (a *YearControl) EditorSetting(value interface{}) *YearControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * inputFormat
 */
func (a *YearControl) InputFormat(value interface{}) *YearControl {
    a.Set("inputFormat", value)
    return a
}

/**
 * remark
 */
func (a *YearControl) Remark(value interface{}) *YearControl {
    a.Set("remark", value)
    return a
}

/**
 * placeholder
 */
func (a *YearControl) Placeholder(value interface{}) *YearControl {
    a.Set("placeholder", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *YearControl) StaticPlaceholder(value interface{}) *YearControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * inputForbid
 */
func (a *YearControl) InputForbid(value interface{}) *YearControl {
    a.Set("inputForbid", value)
    return a
}

/**
 * name
 */
func (a *YearControl) Name(value interface{}) *YearControl {
    a.Set("name", value)
    return a
}

/**
 * extraName
 */
func (a *YearControl) ExtraName(value interface{}) *YearControl {
    a.Set("extraName", value)
    return a
}

/**
 * readOnly
 */
func (a *YearControl) ReadOnly(value interface{}) *YearControl {
    a.Set("readOnly", value)
    return a
}

/**
 * visible
 */
func (a *YearControl) Visible(value interface{}) *YearControl {
    a.Set("visible", value)
    return a
}

/**
 * staticClassName
 */
func (a *YearControl) StaticClassName(value interface{}) *YearControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *YearControl) StaticLabelClassName(value interface{}) *YearControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *YearControl) Type(value interface{}) *YearControl {
    a.Set("type", value)
    return a
}

/**
 * label
 */
func (a *YearControl) Label(value interface{}) *YearControl {
    a.Set("label", value)
    return a
}

/**
 * mode
 */
func (a *YearControl) Mode(value interface{}) *YearControl {
    a.Set("mode", value)
    return a
}

/**
 * onEvent
 */
func (a *YearControl) OnEvent(value interface{}) *YearControl {
    a.Set("onEvent", value)
    return a
}

/**
 * valueFormat
 */
func (a *YearControl) ValueFormat(value interface{}) *YearControl {
    a.Set("valueFormat", value)
    return a
}

/**
 * validationErrors
 */
func (a *YearControl) ValidationErrors(value interface{}) *YearControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * hidden
 */
func (a *YearControl) Hidden(value interface{}) *YearControl {
    a.Set("hidden", value)
    return a
}

/**
 * id
 */
func (a *YearControl) Id(value interface{}) *YearControl {
    a.Set("id", value)
    return a
}

/**
 * value
 */
func (a *YearControl) Value(value interface{}) *YearControl {
    a.Set("value", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *YearControl) ClearValueOnHidden(value interface{}) *YearControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * labelClassName
 */
func (a *YearControl) LabelClassName(value interface{}) *YearControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * labelRemark
 */
func (a *YearControl) LabelRemark(value interface{}) *YearControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * validateOnChange
 */
func (a *YearControl) ValidateOnChange(value interface{}) *YearControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * visibleOn
 */
func (a *YearControl) VisibleOn(value interface{}) *YearControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * disabledDate
 */
func (a *YearControl) DisabledDate(value interface{}) *YearControl {
    a.Set("disabledDate", value)
    return a
}

/**
 * labelWidth
 */
func (a *YearControl) LabelWidth(value interface{}) *YearControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * hiddenOn
 */
func (a *YearControl) HiddenOn(value interface{}) *YearControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * static
 */
func (a *YearControl) Static(value interface{}) *YearControl {
    a.Set("static", value)
    return a
}

/**
 * format
 */
func (a *YearControl) Format(value interface{}) *YearControl {
    a.Set("format", value)
    return a
}

/**
 * required
 */
func (a *YearControl) Required(value interface{}) *YearControl {
    a.Set("required", value)
    return a
}

/**
 * borderMode
 */
func (a *YearControl) BorderMode(value interface{}) *YearControl {
    a.Set("borderMode", value)
    return a
}

/**
 * style
 */
func (a *YearControl) Style(value interface{}) *YearControl {
    a.Set("style", value)
    return a
}

/**
 * clearable
 */
func (a *YearControl) Clearable(value interface{}) *YearControl {
    a.Set("clearable", value)
    return a
}

/**
 * hint
 */
func (a *YearControl) Hint(value interface{}) *YearControl {
    a.Set("hint", value)
    return a
}

/**
 * labelOverflow
 */
func (a *YearControl) LabelOverflow(value interface{}) *YearControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *YearControl) ReadOnlyOn(value interface{}) *YearControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * desc
 */
func (a *YearControl) Desc(value interface{}) *YearControl {
    a.Set("desc", value)
    return a
}

/**
 * validations
 */
func (a *YearControl) Validations(value interface{}) *YearControl {
    a.Set("validations", value)
    return a
}

/**
 * validateApi
 */
func (a *YearControl) ValidateApi(value interface{}) *YearControl {
    a.Set("validateApi", value)
    return a
}

/**
 * row
 */
func (a *YearControl) Row(value interface{}) *YearControl {
    a.Set("row", value)
    return a
}

/**
 * useMobileUI
 */
func (a *YearControl) UseMobileUI(value interface{}) *YearControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *YearControl) Width(value interface{}) *YearControl {
    a.Set("width", value)
    return a
}

/**
 * inputClassName
 */
func (a *YearControl) InputClassName(value interface{}) *YearControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * displayFormat
 */
func (a *YearControl) DisplayFormat(value interface{}) *YearControl {
    a.Set("displayFormat", value)
    return a
}
