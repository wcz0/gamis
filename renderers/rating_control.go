package renderers


/**
 * Rating 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/rating

 * @author wcz0
 * @version 6.2.2
 */
type RatingControl struct {
	*BaseRenderer
}

func NewRatingControl() *RatingControl {
    a := &RatingControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-rating")
    return a
}


func (a *RatingControl) Set(name string, value interface{}) *RatingControl {
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
func (a *RatingControl) LabelWidth(value interface{}) *RatingControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * name
 */
func (a *RatingControl) Name(value interface{}) *RatingControl {
    a.Set("name", value)
    return a
}

/**
 * validateOnChange
 */
func (a *RatingControl) ValidateOnChange(value interface{}) *RatingControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *RatingControl) DescriptionClassName(value interface{}) *RatingControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * required
 */
func (a *RatingControl) Required(value interface{}) *RatingControl {
    a.Set("required", value)
    return a
}

/**
 * disabledOn
 */
func (a *RatingControl) DisabledOn(value interface{}) *RatingControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *RatingControl) StaticPlaceholder(value interface{}) *RatingControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * count
 */
func (a *RatingControl) Count(value interface{}) *RatingControl {
    a.Set("count", value)
    return a
}

/**
 * staticOn
 */
func (a *RatingControl) StaticOn(value interface{}) *RatingControl {
    a.Set("staticOn", value)
    return a
}

/**
 * labelOverflow
 */
func (a *RatingControl) LabelOverflow(value interface{}) *RatingControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * mode
 */
func (a *RatingControl) Mode(value interface{}) *RatingControl {
    a.Set("mode", value)
    return a
}

/**
 * initAutoFill
 */
func (a *RatingControl) InitAutoFill(value interface{}) *RatingControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * row
 */
func (a *RatingControl) Row(value interface{}) *RatingControl {
    a.Set("row", value)
    return a
}

/**
 * hidden
 */
func (a *RatingControl) Hidden(value interface{}) *RatingControl {
    a.Set("hidden", value)
    return a
}

/**
 * allowClear
 */
func (a *RatingControl) AllowClear(value interface{}) *RatingControl {
    a.Set("allowClear", value)
    return a
}

/**
 * char
 */
func (a *RatingControl) Char(value interface{}) *RatingControl {
    a.Set("char", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *RatingControl) ReadOnlyOn(value interface{}) *RatingControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 */
func (a *RatingControl) Type(value interface{}) *RatingControl {
    a.Set("type", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *RatingControl) Width(value interface{}) *RatingControl {
    a.Set("width", value)
    return a
}

/**
 * description
 */
func (a *RatingControl) Description(value interface{}) *RatingControl {
    a.Set("description", value)
    return a
}

/**
 * horizontal
 */
func (a *RatingControl) Horizontal(value interface{}) *RatingControl {
    a.Set("horizontal", value)
    return a
}

/**
 * className
 */
func (a *RatingControl) ClassName(value interface{}) *RatingControl {
    a.Set("className", value)
    return a
}

/**
 * readonly
 */
func (a *RatingControl) Readonly(value interface{}) *RatingControl {
    a.Set("readonly", value)
    return a
}

/**
 * size
 */
func (a *RatingControl) Size(value interface{}) *RatingControl {
    a.Set("size", value)
    return a
}

/**
 * extraName
 */
func (a *RatingControl) ExtraName(value interface{}) *RatingControl {
    a.Set("extraName", value)
    return a
}

/**
 * remark
 */
func (a *RatingControl) Remark(value interface{}) *RatingControl {
    a.Set("remark", value)
    return a
}

/**
 * visibleOn
 */
func (a *RatingControl) VisibleOn(value interface{}) *RatingControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * colors
 */
func (a *RatingControl) Colors(value interface{}) *RatingControl {
    a.Set("colors", value)
    return a
}

/**
 * charClassName
 */
func (a *RatingControl) CharClassName(value interface{}) *RatingControl {
    a.Set("charClassName", value)
    return a
}

/**
 * labelRemark
 */
func (a *RatingControl) LabelRemark(value interface{}) *RatingControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * readOnly
 */
func (a *RatingControl) ReadOnly(value interface{}) *RatingControl {
    a.Set("readOnly", value)
    return a
}

/**
 * placeholder
 */
func (a *RatingControl) Placeholder(value interface{}) *RatingControl {
    a.Set("placeholder", value)
    return a
}

/**
 * onEvent
 */
func (a *RatingControl) OnEvent(value interface{}) *RatingControl {
    a.Set("onEvent", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *RatingControl) StaticLabelClassName(value interface{}) *RatingControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * style
 */
func (a *RatingControl) Style(value interface{}) *RatingControl {
    a.Set("style", value)
    return a
}

/**
 * value
 */
func (a *RatingControl) Value(value interface{}) *RatingControl {
    a.Set("value", value)
    return a
}

/**
 * textClassName
 */
func (a *RatingControl) TextClassName(value interface{}) *RatingControl {
    a.Set("textClassName", value)
    return a
}

/**
 * validationErrors
 */
func (a *RatingControl) ValidationErrors(value interface{}) *RatingControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * validations
 */
func (a *RatingControl) Validations(value interface{}) *RatingControl {
    a.Set("validations", value)
    return a
}

/**
 * validateApi
 */
func (a *RatingControl) ValidateApi(value interface{}) *RatingControl {
    a.Set("validateApi", value)
    return a
}

/**
 * submitOnChange
 */
func (a *RatingControl) SubmitOnChange(value interface{}) *RatingControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * inline
 */
func (a *RatingControl) Inline(value interface{}) *RatingControl {
    a.Set("inline", value)
    return a
}

/**
 * inputClassName
 */
func (a *RatingControl) InputClassName(value interface{}) *RatingControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * label
 */
func (a *RatingControl) Label(value interface{}) *RatingControl {
    a.Set("label", value)
    return a
}

/**
 * disabled
 */
func (a *RatingControl) Disabled(value interface{}) *RatingControl {
    a.Set("disabled", value)
    return a
}

/**
 * hiddenOn
 */
func (a *RatingControl) HiddenOn(value interface{}) *RatingControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * static
 */
func (a *RatingControl) Static(value interface{}) *RatingControl {
    a.Set("static", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *RatingControl) StaticInputClassName(value interface{}) *RatingControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * texts
 */
func (a *RatingControl) Texts(value interface{}) *RatingControl {
    a.Set("texts", value)
    return a
}

/**
 * labelAlign
 */
func (a *RatingControl) LabelAlign(value interface{}) *RatingControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * staticClassName
 */
func (a *RatingControl) StaticClassName(value interface{}) *RatingControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *RatingControl) StaticSchema(value interface{}) *RatingControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * textPosition
 */
func (a *RatingControl) TextPosition(value interface{}) *RatingControl {
    a.Set("textPosition", value)
    return a
}

/**
 * half
 */
func (a *RatingControl) Half(value interface{}) *RatingControl {
    a.Set("half", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *RatingControl) ClearValueOnHidden(value interface{}) *RatingControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * labelClassName
 */
func (a *RatingControl) LabelClassName(value interface{}) *RatingControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * hint
 */
func (a *RatingControl) Hint(value interface{}) *RatingControl {
    a.Set("hint", value)
    return a
}

/**
 * desc
 */
func (a *RatingControl) Desc(value interface{}) *RatingControl {
    a.Set("desc", value)
    return a
}

/**
 * autoFill
 */
func (a *RatingControl) AutoFill(value interface{}) *RatingControl {
    a.Set("autoFill", value)
    return a
}

/**
 * editorSetting
 */
func (a *RatingControl) EditorSetting(value interface{}) *RatingControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * visible
 */
func (a *RatingControl) Visible(value interface{}) *RatingControl {
    a.Set("visible", value)
    return a
}

/**
 * useMobileUI
 */
func (a *RatingControl) UseMobileUI(value interface{}) *RatingControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * id
 */
func (a *RatingControl) Id(value interface{}) *RatingControl {
    a.Set("id", value)
    return a
}

/**
 * inactiveColor
 */
func (a *RatingControl) InactiveColor(value interface{}) *RatingControl {
    a.Set("inactiveColor", value)
    return a
}
