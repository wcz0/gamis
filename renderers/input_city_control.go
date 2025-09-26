package renderers


/**
 * City 城市选择框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/city

 * @author wcz0
 * @version 6.2.2
 */
type InputCityControl struct {
	*BaseRenderer
}

func NewInputCityControl() *InputCityControl {
    a := &InputCityControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-city")
    return a
}


func (a *InputCityControl) Set(name string, value interface{}) *InputCityControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * descriptionClassName
 */
func (a *InputCityControl) DescriptionClassName(value interface{}) *InputCityControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * validations
 */
func (a *InputCityControl) Validations(value interface{}) *InputCityControl {
    a.Set("validations", value)
    return a
}

/**
 * allowDistrict
 */
func (a *InputCityControl) AllowDistrict(value interface{}) *InputCityControl {
    a.Set("allowDistrict", value)
    return a
}

/**
 * loadingConfig
 */
func (a *InputCityControl) LoadingConfig(value interface{}) *InputCityControl {
    a.Set("loadingConfig", value)
    return a
}

/**
 * validateOnChange
 */
func (a *InputCityControl) ValidateOnChange(value interface{}) *InputCityControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * className
 */
func (a *InputCityControl) ClassName(value interface{}) *InputCityControl {
    a.Set("className", value)
    return a
}

/**
 * disabled
 */
func (a *InputCityControl) Disabled(value interface{}) *InputCityControl {
    a.Set("disabled", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *InputCityControl) StaticLabelClassName(value interface{}) *InputCityControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * placeholder
 */
func (a *InputCityControl) Placeholder(value interface{}) *InputCityControl {
    a.Set("placeholder", value)
    return a
}

/**
 * validationErrors
 */
func (a *InputCityControl) ValidationErrors(value interface{}) *InputCityControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * staticSchema
 */
func (a *InputCityControl) StaticSchema(value interface{}) *InputCityControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * itemClassName
 */
func (a *InputCityControl) ItemClassName(value interface{}) *InputCityControl {
    a.Set("itemClassName", value)
    return a
}

/**
 * desc
 */
func (a *InputCityControl) Desc(value interface{}) *InputCityControl {
    a.Set("desc", value)
    return a
}

/**
 * hidden
 */
func (a *InputCityControl) Hidden(value interface{}) *InputCityControl {
    a.Set("hidden", value)
    return a
}

/**
 * labelOverflow
 */
func (a *InputCityControl) LabelOverflow(value interface{}) *InputCityControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * labelClassName
 */
func (a *InputCityControl) LabelClassName(value interface{}) *InputCityControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * inputClassName
 */
func (a *InputCityControl) InputClassName(value interface{}) *InputCityControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *InputCityControl) StaticInputClassName(value interface{}) *InputCityControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * extraName
 */
func (a *InputCityControl) ExtraName(value interface{}) *InputCityControl {
    a.Set("extraName", value)
    return a
}

/**
 * submitOnChange
 */
func (a *InputCityControl) SubmitOnChange(value interface{}) *InputCityControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * mode
 */
func (a *InputCityControl) Mode(value interface{}) *InputCityControl {
    a.Set("mode", value)
    return a
}

/**
 */
func (a *InputCityControl) Type(value interface{}) *InputCityControl {
    a.Set("type", value)
    return a
}

/**
 * hint
 */
func (a *InputCityControl) Hint(value interface{}) *InputCityControl {
    a.Set("hint", value)
    return a
}

/**
 * style
 */
func (a *InputCityControl) Style(value interface{}) *InputCityControl {
    a.Set("style", value)
    return a
}

/**
 * horizontal
 */
func (a *InputCityControl) Horizontal(value interface{}) *InputCityControl {
    a.Set("horizontal", value)
    return a
}

/**
 * row
 */
func (a *InputCityControl) Row(value interface{}) *InputCityControl {
    a.Set("row", value)
    return a
}

/**
 * id
 */
func (a *InputCityControl) Id(value interface{}) *InputCityControl {
    a.Set("id", value)
    return a
}

/**
 * delimiter
 */
func (a *InputCityControl) Delimiter(value interface{}) *InputCityControl {
    a.Set("delimiter", value)
    return a
}

/**
 * label
 */
func (a *InputCityControl) Label(value interface{}) *InputCityControl {
    a.Set("label", value)
    return a
}

/**
 * labelAlign
 */
func (a *InputCityControl) LabelAlign(value interface{}) *InputCityControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * allowCity
 */
func (a *InputCityControl) AllowCity(value interface{}) *InputCityControl {
    a.Set("allowCity", value)
    return a
}

/**
 * labelWidth
 */
func (a *InputCityControl) LabelWidth(value interface{}) *InputCityControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * labelRemark
 */
func (a *InputCityControl) LabelRemark(value interface{}) *InputCityControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * size
 */
func (a *InputCityControl) Size(value interface{}) *InputCityControl {
    a.Set("size", value)
    return a
}

/**
 * searchable
 */
func (a *InputCityControl) Searchable(value interface{}) *InputCityControl {
    a.Set("searchable", value)
    return a
}

/**
 * staticOn
 */
func (a *InputCityControl) StaticOn(value interface{}) *InputCityControl {
    a.Set("staticOn", value)
    return a
}

/**
 * extractValue
 */
func (a *InputCityControl) ExtractValue(value interface{}) *InputCityControl {
    a.Set("extractValue", value)
    return a
}

/**
 * allowStreet
 */
func (a *InputCityControl) AllowStreet(value interface{}) *InputCityControl {
    a.Set("allowStreet", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *InputCityControl) Width(value interface{}) *InputCityControl {
    a.Set("width", value)
    return a
}

/**
 * description
 */
func (a *InputCityControl) Description(value interface{}) *InputCityControl {
    a.Set("description", value)
    return a
}

/**
 * initAutoFill
 */
func (a *InputCityControl) InitAutoFill(value interface{}) *InputCityControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * visible
 */
func (a *InputCityControl) Visible(value interface{}) *InputCityControl {
    a.Set("visible", value)
    return a
}

/**
 * useMobileUI
 */
func (a *InputCityControl) UseMobileUI(value interface{}) *InputCityControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * joinValues
 */
func (a *InputCityControl) JoinValues(value interface{}) *InputCityControl {
    a.Set("joinValues", value)
    return a
}

/**
 * inline
 */
func (a *InputCityControl) Inline(value interface{}) *InputCityControl {
    a.Set("inline", value)
    return a
}

/**
 * onEvent
 */
func (a *InputCityControl) OnEvent(value interface{}) *InputCityControl {
    a.Set("onEvent", value)
    return a
}

/**
 * static
 */
func (a *InputCityControl) Static(value interface{}) *InputCityControl {
    a.Set("static", value)
    return a
}

/**
 * name
 */
func (a *InputCityControl) Name(value interface{}) *InputCityControl {
    a.Set("name", value)
    return a
}

/**
 * readOnly
 */
func (a *InputCityControl) ReadOnly(value interface{}) *InputCityControl {
    a.Set("readOnly", value)
    return a
}

/**
 * disabledOn
 */
func (a *InputCityControl) DisabledOn(value interface{}) *InputCityControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * hiddenOn
 */
func (a *InputCityControl) HiddenOn(value interface{}) *InputCityControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *InputCityControl) StaticClassName(value interface{}) *InputCityControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * editorSetting
 */
func (a *InputCityControl) EditorSetting(value interface{}) *InputCityControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * required
 */
func (a *InputCityControl) Required(value interface{}) *InputCityControl {
    a.Set("required", value)
    return a
}

/**
 * value
 */
func (a *InputCityControl) Value(value interface{}) *InputCityControl {
    a.Set("value", value)
    return a
}

/**
 * validateApi
 */
func (a *InputCityControl) ValidateApi(value interface{}) *InputCityControl {
    a.Set("validateApi", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *InputCityControl) StaticPlaceholder(value interface{}) *InputCityControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *InputCityControl) ReadOnlyOn(value interface{}) *InputCityControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * remark
 */
func (a *InputCityControl) Remark(value interface{}) *InputCityControl {
    a.Set("remark", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *InputCityControl) ClearValueOnHidden(value interface{}) *InputCityControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * autoFill
 */
func (a *InputCityControl) AutoFill(value interface{}) *InputCityControl {
    a.Set("autoFill", value)
    return a
}

/**
 * visibleOn
 */
func (a *InputCityControl) VisibleOn(value interface{}) *InputCityControl {
    a.Set("visibleOn", value)
    return a
}
