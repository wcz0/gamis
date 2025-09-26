package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type RangeControl struct {
	*BaseRenderer
}

func NewRangeControl() *RangeControl {
    a := &RangeControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-range")
    return a
}


func (a *RangeControl) Set(name string, value interface{}) *RangeControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * labelWidth
 */
func (a *RangeControl) LabelWidth(value interface{}) *RangeControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * onEvent
 */
func (a *RangeControl) OnEvent(value interface{}) *RangeControl {
    a.Set("onEvent", value)
    return a
}

/**
 * placeholder
 */
func (a *RangeControl) Placeholder(value interface{}) *RangeControl {
    a.Set("placeholder", value)
    return a
}

/**
 * hidden
 */
func (a *RangeControl) Hidden(value interface{}) *RangeControl {
    a.Set("hidden", value)
    return a
}

/**
 * step
 */
func (a *RangeControl) Step(value interface{}) *RangeControl {
    a.Set("step", value)
    return a
}

/**
 * parts
 */
func (a *RangeControl) Parts(value interface{}) *RangeControl {
    a.Set("parts", value)
    return a
}

/**
 * joinValues
 */
func (a *RangeControl) JoinValues(value interface{}) *RangeControl {
    a.Set("joinValues", value)
    return a
}

/**
 * delimiter
 */
func (a *RangeControl) Delimiter(value interface{}) *RangeControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *RangeControl) Width(value interface{}) *RangeControl {
    a.Set("width", value)
    return a
}

/**
 * labelAlign
 */
func (a *RangeControl) LabelAlign(value interface{}) *RangeControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * submitOnChange
 */
func (a *RangeControl) SubmitOnChange(value interface{}) *RangeControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * required
 */
func (a *RangeControl) Required(value interface{}) *RangeControl {
    a.Set("required", value)
    return a
}

/**
 * disabledOn
 */
func (a *RangeControl) DisabledOn(value interface{}) *RangeControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *RangeControl) StaticInputClassName(value interface{}) *RangeControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * max
 */
func (a *RangeControl) Max(value interface{}) *RangeControl {
    a.Set("max", value)
    return a
}

/**
 * showInput
 */
func (a *RangeControl) ShowInput(value interface{}) *RangeControl {
    a.Set("showInput", value)
    return a
}

/**
 * labelOverflow
 */
func (a *RangeControl) LabelOverflow(value interface{}) *RangeControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * name
 */
func (a *RangeControl) Name(value interface{}) *RangeControl {
    a.Set("name", value)
    return a
}

/**
 * validateApi
 */
func (a *RangeControl) ValidateApi(value interface{}) *RangeControl {
    a.Set("validateApi", value)
    return a
}

/**
 * staticSchema
 */
func (a *RangeControl) StaticSchema(value interface{}) *RangeControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * extraName
 */
func (a *RangeControl) ExtraName(value interface{}) *RangeControl {
    a.Set("extraName", value)
    return a
}

/**
 * validations
 */
func (a *RangeControl) Validations(value interface{}) *RangeControl {
    a.Set("validations", value)
    return a
}

/**
 * autoFill
 */
func (a *RangeControl) AutoFill(value interface{}) *RangeControl {
    a.Set("autoFill", value)
    return a
}

/**
 * className
 */
func (a *RangeControl) ClassName(value interface{}) *RangeControl {
    a.Set("className", value)
    return a
}

/**
 * style
 */
func (a *RangeControl) Style(value interface{}) *RangeControl {
    a.Set("style", value)
    return a
}

/**
 * editorSetting
 */
func (a *RangeControl) EditorSetting(value interface{}) *RangeControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * tooltipVisible
 */
func (a *RangeControl) TooltipVisible(value interface{}) *RangeControl {
    a.Set("tooltipVisible", value)
    return a
}

/**
 * useMobileUI
 */
func (a *RangeControl) UseMobileUI(value interface{}) *RangeControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * visibleOn
 */
func (a *RangeControl) VisibleOn(value interface{}) *RangeControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *RangeControl) StaticClassName(value interface{}) *RangeControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *RangeControl) ReadOnlyOn(value interface{}) *RangeControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * description
 */
func (a *RangeControl) Description(value interface{}) *RangeControl {
    a.Set("description", value)
    return a
}

/**
 * mode
 */
func (a *RangeControl) Mode(value interface{}) *RangeControl {
    a.Set("mode", value)
    return a
}

/**
 * horizontal
 */
func (a *RangeControl) Horizontal(value interface{}) *RangeControl {
    a.Set("horizontal", value)
    return a
}

/**
 * validationErrors
 */
func (a *RangeControl) ValidationErrors(value interface{}) *RangeControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * initAutoFill
 */
func (a *RangeControl) InitAutoFill(value interface{}) *RangeControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * row
 */
func (a *RangeControl) Row(value interface{}) *RangeControl {
    a.Set("row", value)
    return a
}

/**
 * visible
 */
func (a *RangeControl) Visible(value interface{}) *RangeControl {
    a.Set("visible", value)
    return a
}

/**
 * static
 */
func (a *RangeControl) Static(value interface{}) *RangeControl {
    a.Set("static", value)
    return a
}

/**
 * size
 */
func (a *RangeControl) Size(value interface{}) *RangeControl {
    a.Set("size", value)
    return a
}

/**
 * label
 */
func (a *RangeControl) Label(value interface{}) *RangeControl {
    a.Set("label", value)
    return a
}

/**
 * inline
 */
func (a *RangeControl) Inline(value interface{}) *RangeControl {
    a.Set("inline", value)
    return a
}

/**
 * hiddenOn
 */
func (a *RangeControl) HiddenOn(value interface{}) *RangeControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * id
 */
func (a *RangeControl) Id(value interface{}) *RangeControl {
    a.Set("id", value)
    return a
}

/**
 * showSteps
 */
func (a *RangeControl) ShowSteps(value interface{}) *RangeControl {
    a.Set("showSteps", value)
    return a
}

/**
 * marks
 */
func (a *RangeControl) Marks(value interface{}) *RangeControl {
    a.Set("marks", value)
    return a
}

/**
 * clearable
 */
func (a *RangeControl) Clearable(value interface{}) *RangeControl {
    a.Set("clearable", value)
    return a
}

/**
 * labelRemark
 */
func (a *RangeControl) LabelRemark(value interface{}) *RangeControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * desc
 */
func (a *RangeControl) Desc(value interface{}) *RangeControl {
    a.Set("desc", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *RangeControl) StaticLabelClassName(value interface{}) *RangeControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * min
 */
func (a *RangeControl) Min(value interface{}) *RangeControl {
    a.Set("min", value)
    return a
}

/**
 */
func (a *RangeControl) Type(value interface{}) *RangeControl {
    a.Set("type", value)
    return a
}

/**
 * hint
 */
func (a *RangeControl) Hint(value interface{}) *RangeControl {
    a.Set("hint", value)
    return a
}

/**
 * readOnly
 */
func (a *RangeControl) ReadOnly(value interface{}) *RangeControl {
    a.Set("readOnly", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *RangeControl) DescriptionClassName(value interface{}) *RangeControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *RangeControl) ClearValueOnHidden(value interface{}) *RangeControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * unit
 */
func (a *RangeControl) Unit(value interface{}) *RangeControl {
    a.Set("unit", value)
    return a
}

/**
 * multiple
 */
func (a *RangeControl) Multiple(value interface{}) *RangeControl {
    a.Set("multiple", value)
    return a
}

/**
 * labelClassName
 */
func (a *RangeControl) LabelClassName(value interface{}) *RangeControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * validateOnChange
 */
func (a *RangeControl) ValidateOnChange(value interface{}) *RangeControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * inputClassName
 */
func (a *RangeControl) InputClassName(value interface{}) *RangeControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * staticOn
 */
func (a *RangeControl) StaticOn(value interface{}) *RangeControl {
    a.Set("staticOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *RangeControl) StaticPlaceholder(value interface{}) *RangeControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * tooltipPlacement
 */
func (a *RangeControl) TooltipPlacement(value interface{}) *RangeControl {
    a.Set("tooltipPlacement", value)
    return a
}

/**
 * remark
 */
func (a *RangeControl) Remark(value interface{}) *RangeControl {
    a.Set("remark", value)
    return a
}

/**
 * value
 */
func (a *RangeControl) Value(value interface{}) *RangeControl {
    a.Set("value", value)
    return a
}

/**
 * disabled
 */
func (a *RangeControl) Disabled(value interface{}) *RangeControl {
    a.Set("disabled", value)
    return a
}
