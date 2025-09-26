package renderers


/**
 * Location 选点组件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/location

 * @author wcz0
 * @version 6.2.2
 */
type LocationControl struct {
	*BaseRenderer
}

func NewLocationControl() *LocationControl {
    a := &LocationControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "location-picker")
    return a
}


func (a *LocationControl) Set(name string, value interface{}) *LocationControl {
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
func (a *LocationControl) SubmitOnChange(value interface{}) *LocationControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * inputClassName
 */
func (a *LocationControl) InputClassName(value interface{}) *LocationControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * validateApi
 */
func (a *LocationControl) ValidateApi(value interface{}) *LocationControl {
    a.Set("validateApi", value)
    return a
}

/**
 * ak
 */
func (a *LocationControl) Ak(value interface{}) *LocationControl {
    a.Set("ak", value)
    return a
}

/**
 * readOnly
 */
func (a *LocationControl) ReadOnly(value interface{}) *LocationControl {
    a.Set("readOnly", value)
    return a
}

/**
 * autoSelectCurrentLoc
 */
func (a *LocationControl) AutoSelectCurrentLoc(value interface{}) *LocationControl {
    a.Set("autoSelectCurrentLoc", value)
    return a
}

/**
 * onlySelectCurrentLoc
 */
func (a *LocationControl) OnlySelectCurrentLoc(value interface{}) *LocationControl {
    a.Set("onlySelectCurrentLoc", value)
    return a
}

/**
 * disabledOn
 */
func (a *LocationControl) DisabledOn(value interface{}) *LocationControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * id
 */
func (a *LocationControl) Id(value interface{}) *LocationControl {
    a.Set("id", value)
    return a
}

/**
 * name
 */
func (a *LocationControl) Name(value interface{}) *LocationControl {
    a.Set("name", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *LocationControl) ReadOnlyOn(value interface{}) *LocationControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * mode
 */
func (a *LocationControl) Mode(value interface{}) *LocationControl {
    a.Set("mode", value)
    return a
}

/**
 * initAutoFill
 */
func (a *LocationControl) InitAutoFill(value interface{}) *LocationControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * hint
 */
func (a *LocationControl) Hint(value interface{}) *LocationControl {
    a.Set("hint", value)
    return a
}

/**
 * disabled
 */
func (a *LocationControl) Disabled(value interface{}) *LocationControl {
    a.Set("disabled", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *LocationControl) Width(value interface{}) *LocationControl {
    a.Set("width", value)
    return a
}

/**
 * labelWidth
 */
func (a *LocationControl) LabelWidth(value interface{}) *LocationControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * value
 */
func (a *LocationControl) Value(value interface{}) *LocationControl {
    a.Set("value", value)
    return a
}

/**
 * row
 */
func (a *LocationControl) Row(value interface{}) *LocationControl {
    a.Set("row", value)
    return a
}

/**
 * staticSchema
 */
func (a *LocationControl) StaticSchema(value interface{}) *LocationControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * useMobileUI
 */
func (a *LocationControl) UseMobileUI(value interface{}) *LocationControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * validations
 */
func (a *LocationControl) Validations(value interface{}) *LocationControl {
    a.Set("validations", value)
    return a
}

/**
 * visibleOn
 */
func (a *LocationControl) VisibleOn(value interface{}) *LocationControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * size
 */
func (a *LocationControl) Size(value interface{}) *LocationControl {
    a.Set("size", value)
    return a
}

/**
 * labelAlign
 */
func (a *LocationControl) LabelAlign(value interface{}) *LocationControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * labelClassName
 */
func (a *LocationControl) LabelClassName(value interface{}) *LocationControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *LocationControl) DescriptionClassName(value interface{}) *LocationControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * validationErrors
 */
func (a *LocationControl) ValidationErrors(value interface{}) *LocationControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * hidden
 */
func (a *LocationControl) Hidden(value interface{}) *LocationControl {
    a.Set("hidden", value)
    return a
}

/**
 * staticClassName
 */
func (a *LocationControl) StaticClassName(value interface{}) *LocationControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *LocationControl) StaticInputClassName(value interface{}) *LocationControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *LocationControl) Type(value interface{}) *LocationControl {
    a.Set("type", value)
    return a
}

/**
 * label
 */
func (a *LocationControl) Label(value interface{}) *LocationControl {
    a.Set("label", value)
    return a
}

/**
 * validateOnChange
 */
func (a *LocationControl) ValidateOnChange(value interface{}) *LocationControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * onEvent
 */
func (a *LocationControl) OnEvent(value interface{}) *LocationControl {
    a.Set("onEvent", value)
    return a
}

/**
 * staticOn
 */
func (a *LocationControl) StaticOn(value interface{}) *LocationControl {
    a.Set("staticOn", value)
    return a
}

/**
 * editorSetting
 */
func (a *LocationControl) EditorSetting(value interface{}) *LocationControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * getLocationPlaceholder
 */
func (a *LocationControl) GetLocationPlaceholder(value interface{}) *LocationControl {
    a.Set("getLocationPlaceholder", value)
    return a
}

/**
 * labelOverflow
 */
func (a *LocationControl) LabelOverflow(value interface{}) *LocationControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * labelRemark
 */
func (a *LocationControl) LabelRemark(value interface{}) *LocationControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * horizontal
 */
func (a *LocationControl) Horizontal(value interface{}) *LocationControl {
    a.Set("horizontal", value)
    return a
}

/**
 * desc
 */
func (a *LocationControl) Desc(value interface{}) *LocationControl {
    a.Set("desc", value)
    return a
}

/**
 * inline
 */
func (a *LocationControl) Inline(value interface{}) *LocationControl {
    a.Set("inline", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *LocationControl) ClearValueOnHidden(value interface{}) *LocationControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *LocationControl) StaticPlaceholder(value interface{}) *LocationControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * remark
 */
func (a *LocationControl) Remark(value interface{}) *LocationControl {
    a.Set("remark", value)
    return a
}

/**
 * autoFill
 */
func (a *LocationControl) AutoFill(value interface{}) *LocationControl {
    a.Set("autoFill", value)
    return a
}

/**
 * className
 */
func (a *LocationControl) ClassName(value interface{}) *LocationControl {
    a.Set("className", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *LocationControl) StaticLabelClassName(value interface{}) *LocationControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * vendor
 */
func (a *LocationControl) Vendor(value interface{}) *LocationControl {
    a.Set("vendor", value)
    return a
}

/**
 * hideViewControl
 */
func (a *LocationControl) HideViewControl(value interface{}) *LocationControl {
    a.Set("hideViewControl", value)
    return a
}

/**
 * visible
 */
func (a *LocationControl) Visible(value interface{}) *LocationControl {
    a.Set("visible", value)
    return a
}

/**
 * style
 */
func (a *LocationControl) Style(value interface{}) *LocationControl {
    a.Set("style", value)
    return a
}

/**
 * placeholder
 */
func (a *LocationControl) Placeholder(value interface{}) *LocationControl {
    a.Set("placeholder", value)
    return a
}

/**
 * required
 */
func (a *LocationControl) Required(value interface{}) *LocationControl {
    a.Set("required", value)
    return a
}

/**
 * hiddenOn
 */
func (a *LocationControl) HiddenOn(value interface{}) *LocationControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * extraName
 */
func (a *LocationControl) ExtraName(value interface{}) *LocationControl {
    a.Set("extraName", value)
    return a
}

/**
 * static
 */
func (a *LocationControl) Static(value interface{}) *LocationControl {
    a.Set("static", value)
    return a
}

/**
 * description
 */
func (a *LocationControl) Description(value interface{}) *LocationControl {
    a.Set("description", value)
    return a
}
