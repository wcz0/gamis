package renderers


/**
 * Static 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/static

 * @author wcz0
 * @version 6.2.2
 */
type StaticExactControl struct {
	*BaseRenderer
}

func NewStaticExactControl() *StaticExactControl {
    a := &StaticExactControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "static")
    return a
}


func (a *StaticExactControl) Set(name string, value interface{}) *StaticExactControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * autoFill
 */
func (a *StaticExactControl) AutoFill(value interface{}) *StaticExactControl {
    a.Set("autoFill", value)
    return a
}

/**
 * className
 */
func (a *StaticExactControl) ClassName(value interface{}) *StaticExactControl {
    a.Set("className", value)
    return a
}

/**
 */
func (a *StaticExactControl) Type(value interface{}) *StaticExactControl {
    a.Set("type", value)
    return a
}

/**
 * tpl
 */
func (a *StaticExactControl) Tpl(value interface{}) *StaticExactControl {
    a.Set("tpl", value)
    return a
}

/**
 * labelClassName
 */
func (a *StaticExactControl) LabelClassName(value interface{}) *StaticExactControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * extraName
 */
func (a *StaticExactControl) ExtraName(value interface{}) *StaticExactControl {
    a.Set("extraName", value)
    return a
}

/**
 * initAutoFill
 */
func (a *StaticExactControl) InitAutoFill(value interface{}) *StaticExactControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * disabled
 */
func (a *StaticExactControl) Disabled(value interface{}) *StaticExactControl {
    a.Set("disabled", value)
    return a
}

/**
 * popOver
 */
func (a *StaticExactControl) PopOver(value interface{}) *StaticExactControl {
    a.Set("popOver", value)
    return a
}

/**
 * copyable
 */
func (a *StaticExactControl) Copyable(value interface{}) *StaticExactControl {
    a.Set("copyable", value)
    return a
}

/**
 * borderMode
 */
func (a *StaticExactControl) BorderMode(value interface{}) *StaticExactControl {
    a.Set("borderMode", value)
    return a
}

/**
 * hint
 */
func (a *StaticExactControl) Hint(value interface{}) *StaticExactControl {
    a.Set("hint", value)
    return a
}

/**
 * description
 */
func (a *StaticExactControl) Description(value interface{}) *StaticExactControl {
    a.Set("description", value)
    return a
}

/**
 * inputClassName
 */
func (a *StaticExactControl) InputClassName(value interface{}) *StaticExactControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * disabledOn
 */
func (a *StaticExactControl) DisabledOn(value interface{}) *StaticExactControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * static
 */
func (a *StaticExactControl) Static(value interface{}) *StaticExactControl {
    a.Set("static", value)
    return a
}

/**
 * hiddenOn
 */
func (a *StaticExactControl) HiddenOn(value interface{}) *StaticExactControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticSchema
 */
func (a *StaticExactControl) StaticSchema(value interface{}) *StaticExactControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * useMobileUI
 */
func (a *StaticExactControl) UseMobileUI(value interface{}) *StaticExactControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * text
 */
func (a *StaticExactControl) Text(value interface{}) *StaticExactControl {
    a.Set("text", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *StaticExactControl) Width(value interface{}) *StaticExactControl {
    a.Set("width", value)
    return a
}

/**
 * required
 */
func (a *StaticExactControl) Required(value interface{}) *StaticExactControl {
    a.Set("required", value)
    return a
}

/**
 * validationErrors
 */
func (a *StaticExactControl) ValidationErrors(value interface{}) *StaticExactControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * validateApi
 */
func (a *StaticExactControl) ValidateApi(value interface{}) *StaticExactControl {
    a.Set("validateApi", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *StaticExactControl) StaticLabelClassName(value interface{}) *StaticExactControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *StaticExactControl) StaticInputClassName(value interface{}) *StaticExactControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * style
 */
func (a *StaticExactControl) Style(value interface{}) *StaticExactControl {
    a.Set("style", value)
    return a
}

/**
 * desc
 */
func (a *StaticExactControl) Desc(value interface{}) *StaticExactControl {
    a.Set("desc", value)
    return a
}

/**
 * size
 */
func (a *StaticExactControl) Size(value interface{}) *StaticExactControl {
    a.Set("size", value)
    return a
}

/**
 * remark
 */
func (a *StaticExactControl) Remark(value interface{}) *StaticExactControl {
    a.Set("remark", value)
    return a
}

/**
 * mode
 */
func (a *StaticExactControl) Mode(value interface{}) *StaticExactControl {
    a.Set("mode", value)
    return a
}

/**
 * visibleOn
 */
func (a *StaticExactControl) VisibleOn(value interface{}) *StaticExactControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * labelAlign
 */
func (a *StaticExactControl) LabelAlign(value interface{}) *StaticExactControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * placeholder
 */
func (a *StaticExactControl) Placeholder(value interface{}) *StaticExactControl {
    a.Set("placeholder", value)
    return a
}

/**
 * validateOnChange
 */
func (a *StaticExactControl) ValidateOnChange(value interface{}) *StaticExactControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * validations
 */
func (a *StaticExactControl) Validations(value interface{}) *StaticExactControl {
    a.Set("validations", value)
    return a
}

/**
 * id
 */
func (a *StaticExactControl) Id(value interface{}) *StaticExactControl {
    a.Set("id", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *StaticExactControl) StaticPlaceholder(value interface{}) *StaticExactControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *StaticExactControl) ReadOnlyOn(value interface{}) *StaticExactControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *StaticExactControl) DescriptionClassName(value interface{}) *StaticExactControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * hidden
 */
func (a *StaticExactControl) Hidden(value interface{}) *StaticExactControl {
    a.Set("hidden", value)
    return a
}

/**
 * editorSetting
 */
func (a *StaticExactControl) EditorSetting(value interface{}) *StaticExactControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * labelWidth
 */
func (a *StaticExactControl) LabelWidth(value interface{}) *StaticExactControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * row
 */
func (a *StaticExactControl) Row(value interface{}) *StaticExactControl {
    a.Set("row", value)
    return a
}

/**
 * name
 */
func (a *StaticExactControl) Name(value interface{}) *StaticExactControl {
    a.Set("name", value)
    return a
}

/**
 * labelRemark
 */
func (a *StaticExactControl) LabelRemark(value interface{}) *StaticExactControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * readOnly
 */
func (a *StaticExactControl) ReadOnly(value interface{}) *StaticExactControl {
    a.Set("readOnly", value)
    return a
}

/**
 * visible
 */
func (a *StaticExactControl) Visible(value interface{}) *StaticExactControl {
    a.Set("visible", value)
    return a
}

/**
 * label
 */
func (a *StaticExactControl) Label(value interface{}) *StaticExactControl {
    a.Set("label", value)
    return a
}

/**
 * horizontal
 */
func (a *StaticExactControl) Horizontal(value interface{}) *StaticExactControl {
    a.Set("horizontal", value)
    return a
}

/**
 * onEvent
 */
func (a *StaticExactControl) OnEvent(value interface{}) *StaticExactControl {
    a.Set("onEvent", value)
    return a
}

/**
 * staticOn
 */
func (a *StaticExactControl) StaticOn(value interface{}) *StaticExactControl {
    a.Set("staticOn", value)
    return a
}

/**
 * submitOnChange
 */
func (a *StaticExactControl) SubmitOnChange(value interface{}) *StaticExactControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * inline
 */
func (a *StaticExactControl) Inline(value interface{}) *StaticExactControl {
    a.Set("inline", value)
    return a
}

/**
 * staticClassName
 */
func (a *StaticExactControl) StaticClassName(value interface{}) *StaticExactControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * labelOverflow
 */
func (a *StaticExactControl) LabelOverflow(value interface{}) *StaticExactControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * value
 */
func (a *StaticExactControl) Value(value interface{}) *StaticExactControl {
    a.Set("value", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *StaticExactControl) ClearValueOnHidden(value interface{}) *StaticExactControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * quickEdit
 */
func (a *StaticExactControl) QuickEdit(value interface{}) *StaticExactControl {
    a.Set("quickEdit", value)
    return a
}
