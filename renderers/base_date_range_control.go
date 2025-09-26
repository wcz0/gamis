package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type BaseDateRangeControl struct {
	*BaseRenderer
}

func NewBaseDateRangeControl() *BaseDateRangeControl {
    a := &BaseDateRangeControl{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}


func (a *BaseDateRangeControl) Set(name string, value interface{}) *BaseDateRangeControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * staticSchema
 */
func (a *BaseDateRangeControl) StaticSchema(value interface{}) *BaseDateRangeControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * maxDate
 */
func (a *BaseDateRangeControl) MaxDate(value interface{}) *BaseDateRangeControl {
    a.Set("maxDate", value)
    return a
}

/**
 * labelRemark
 */
func (a *BaseDateRangeControl) LabelRemark(value interface{}) *BaseDateRangeControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * mode
 */
func (a *BaseDateRangeControl) Mode(value interface{}) *BaseDateRangeControl {
    a.Set("mode", value)
    return a
}

/**
 * validations
 */
func (a *BaseDateRangeControl) Validations(value interface{}) *BaseDateRangeControl {
    a.Set("validations", value)
    return a
}

/**
 * autoFill
 */
func (a *BaseDateRangeControl) AutoFill(value interface{}) *BaseDateRangeControl {
    a.Set("autoFill", value)
    return a
}

/**
 * hint
 */
func (a *BaseDateRangeControl) Hint(value interface{}) *BaseDateRangeControl {
    a.Set("hint", value)
    return a
}

/**
 * submitOnChange
 */
func (a *BaseDateRangeControl) SubmitOnChange(value interface{}) *BaseDateRangeControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * placeholder
 */
func (a *BaseDateRangeControl) Placeholder(value interface{}) *BaseDateRangeControl {
    a.Set("placeholder", value)
    return a
}

/**
 * valueFormat
 */
func (a *BaseDateRangeControl) ValueFormat(value interface{}) *BaseDateRangeControl {
    a.Set("valueFormat", value)
    return a
}

/**
 * label
 */
func (a *BaseDateRangeControl) Label(value interface{}) *BaseDateRangeControl {
    a.Set("label", value)
    return a
}

/**
 * disabled
 */
func (a *BaseDateRangeControl) Disabled(value interface{}) *BaseDateRangeControl {
    a.Set("disabled", value)
    return a
}

/**
 * hidden
 */
func (a *BaseDateRangeControl) Hidden(value interface{}) *BaseDateRangeControl {
    a.Set("hidden", value)
    return a
}

/**
 * style
 */
func (a *BaseDateRangeControl) Style(value interface{}) *BaseDateRangeControl {
    a.Set("style", value)
    return a
}

/**
 * joinValues
 */
func (a *BaseDateRangeControl) JoinValues(value interface{}) *BaseDateRangeControl {
    a.Set("joinValues", value)
    return a
}

/**
 * type
 */
func (a *BaseDateRangeControl) Type(value interface{}) *BaseDateRangeControl {
    a.Set("type", value)
    return a
}

/**
 * borderMode
 */
func (a *BaseDateRangeControl) BorderMode(value interface{}) *BaseDateRangeControl {
    a.Set("borderMode", value)
    return a
}

/**
 * shortcuts
 */
func (a *BaseDateRangeControl) Shortcuts(value interface{}) *BaseDateRangeControl {
    a.Set("shortcuts", value)
    return a
}

/**
 * readOnly
 */
func (a *BaseDateRangeControl) ReadOnly(value interface{}) *BaseDateRangeControl {
    a.Set("readOnly", value)
    return a
}

/**
 * static
 */
func (a *BaseDateRangeControl) Static(value interface{}) *BaseDateRangeControl {
    a.Set("static", value)
    return a
}

/**
 * horizontal
 */
func (a *BaseDateRangeControl) Horizontal(value interface{}) *BaseDateRangeControl {
    a.Set("horizontal", value)
    return a
}

/**
 * required
 */
func (a *BaseDateRangeControl) Required(value interface{}) *BaseDateRangeControl {
    a.Set("required", value)
    return a
}

/**
 * initAutoFill
 */
func (a *BaseDateRangeControl) InitAutoFill(value interface{}) *BaseDateRangeControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *BaseDateRangeControl) StaticLabelClassName(value interface{}) *BaseDateRangeControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *BaseDateRangeControl) StaticInputClassName(value interface{}) *BaseDateRangeControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * editorSetting
 */
func (a *BaseDateRangeControl) EditorSetting(value interface{}) *BaseDateRangeControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * minDate
 */
func (a *BaseDateRangeControl) MinDate(value interface{}) *BaseDateRangeControl {
    a.Set("minDate", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *BaseDateRangeControl) ClearValueOnHidden(value interface{}) *BaseDateRangeControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * hiddenOn
 */
func (a *BaseDateRangeControl) HiddenOn(value interface{}) *BaseDateRangeControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * format
 */
func (a *BaseDateRangeControl) Format(value interface{}) *BaseDateRangeControl {
    a.Set("format", value)
    return a
}

/**
 * ranges
 */
func (a *BaseDateRangeControl) Ranges(value interface{}) *BaseDateRangeControl {
    a.Set("ranges", value)
    return a
}

/**
 * inputClassName
 */
func (a *BaseDateRangeControl) InputClassName(value interface{}) *BaseDateRangeControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * delimiter
 */
func (a *BaseDateRangeControl) Delimiter(value interface{}) *BaseDateRangeControl {
    a.Set("delimiter", value)
    return a
}

/**
 * displayFormat
 */
func (a *BaseDateRangeControl) DisplayFormat(value interface{}) *BaseDateRangeControl {
    a.Set("displayFormat", value)
    return a
}

/**
 * maxDuration
 */
func (a *BaseDateRangeControl) MaxDuration(value interface{}) *BaseDateRangeControl {
    a.Set("maxDuration", value)
    return a
}

/**
 * transform
 */
func (a *BaseDateRangeControl) Transform(value interface{}) *BaseDateRangeControl {
    a.Set("transform", value)
    return a
}

/**
 * size
 */
func (a *BaseDateRangeControl) Size(value interface{}) *BaseDateRangeControl {
    a.Set("size", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *BaseDateRangeControl) DescriptionClassName(value interface{}) *BaseDateRangeControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * value
 */
func (a *BaseDateRangeControl) Value(value interface{}) *BaseDateRangeControl {
    a.Set("value", value)
    return a
}

/**
 * staticOn
 */
func (a *BaseDateRangeControl) StaticOn(value interface{}) *BaseDateRangeControl {
    a.Set("staticOn", value)
    return a
}

/**
 * startPlaceholder
 */
func (a *BaseDateRangeControl) StartPlaceholder(value interface{}) *BaseDateRangeControl {
    a.Set("startPlaceholder", value)
    return a
}

/**
 * animation
 */
func (a *BaseDateRangeControl) Animation(value interface{}) *BaseDateRangeControl {
    a.Set("animation", value)
    return a
}

/**
 * useMobileUI
 */
func (a *BaseDateRangeControl) UseMobileUI(value interface{}) *BaseDateRangeControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * extraName
 */
func (a *BaseDateRangeControl) ExtraName(value interface{}) *BaseDateRangeControl {
    a.Set("extraName", value)
    return a
}

/**
 * remark
 */
func (a *BaseDateRangeControl) Remark(value interface{}) *BaseDateRangeControl {
    a.Set("remark", value)
    return a
}

/**
 * validateApi
 */
func (a *BaseDateRangeControl) ValidateApi(value interface{}) *BaseDateRangeControl {
    a.Set("validateApi", value)
    return a
}

/**
 * visibleOn
 */
func (a *BaseDateRangeControl) VisibleOn(value interface{}) *BaseDateRangeControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * id
 */
func (a *BaseDateRangeControl) Id(value interface{}) *BaseDateRangeControl {
    a.Set("id", value)
    return a
}

/**
 * inputFormat
 */
func (a *BaseDateRangeControl) InputFormat(value interface{}) *BaseDateRangeControl {
    a.Set("inputFormat", value)
    return a
}

/**
 * endPlaceholder
 */
func (a *BaseDateRangeControl) EndPlaceholder(value interface{}) *BaseDateRangeControl {
    a.Set("endPlaceholder", value)
    return a
}

/**
 * labelAlign
 */
func (a *BaseDateRangeControl) LabelAlign(value interface{}) *BaseDateRangeControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * labelOverflow
 */
func (a *BaseDateRangeControl) LabelOverflow(value interface{}) *BaseDateRangeControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * description
 */
func (a *BaseDateRangeControl) Description(value interface{}) *BaseDateRangeControl {
    a.Set("description", value)
    return a
}

/**
 * inline
 */
func (a *BaseDateRangeControl) Inline(value interface{}) *BaseDateRangeControl {
    a.Set("inline", value)
    return a
}

/**
 * validationErrors
 */
func (a *BaseDateRangeControl) ValidationErrors(value interface{}) *BaseDateRangeControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * className
 */
func (a *BaseDateRangeControl) ClassName(value interface{}) *BaseDateRangeControl {
    a.Set("className", value)
    return a
}

/**
 * visible
 */
func (a *BaseDateRangeControl) Visible(value interface{}) *BaseDateRangeControl {
    a.Set("visible", value)
    return a
}

/**
 * popOverContainerSelector
 */
func (a *BaseDateRangeControl) PopOverContainerSelector(value interface{}) *BaseDateRangeControl {
    a.Set("popOverContainerSelector", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *BaseDateRangeControl) ReadOnlyOn(value interface{}) *BaseDateRangeControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * onEvent
 */
func (a *BaseDateRangeControl) OnEvent(value interface{}) *BaseDateRangeControl {
    a.Set("onEvent", value)
    return a
}

/**
 * name
 */
func (a *BaseDateRangeControl) Name(value interface{}) *BaseDateRangeControl {
    a.Set("name", value)
    return a
}

/**
 * validateOnChange
 */
func (a *BaseDateRangeControl) ValidateOnChange(value interface{}) *BaseDateRangeControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * staticClassName
 */
func (a *BaseDateRangeControl) StaticClassName(value interface{}) *BaseDateRangeControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * embed
 */
func (a *BaseDateRangeControl) Embed(value interface{}) *BaseDateRangeControl {
    a.Set("embed", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *BaseDateRangeControl) Width(value interface{}) *BaseDateRangeControl {
    a.Set("width", value)
    return a
}

/**
 * labelWidth
 */
func (a *BaseDateRangeControl) LabelWidth(value interface{}) *BaseDateRangeControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * row
 */
func (a *BaseDateRangeControl) Row(value interface{}) *BaseDateRangeControl {
    a.Set("row", value)
    return a
}

/**
 * minDuration
 */
func (a *BaseDateRangeControl) MinDuration(value interface{}) *BaseDateRangeControl {
    a.Set("minDuration", value)
    return a
}

/**
 * labelClassName
 */
func (a *BaseDateRangeControl) LabelClassName(value interface{}) *BaseDateRangeControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * desc
 */
func (a *BaseDateRangeControl) Desc(value interface{}) *BaseDateRangeControl {
    a.Set("desc", value)
    return a
}

/**
 * disabledOn
 */
func (a *BaseDateRangeControl) DisabledOn(value interface{}) *BaseDateRangeControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *BaseDateRangeControl) StaticPlaceholder(value interface{}) *BaseDateRangeControl {
    a.Set("staticPlaceholder", value)
    return a
}
