package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type InputSignature struct {
	*BaseRenderer
}

func NewInputSignature() *InputSignature {
    a := &InputSignature{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-signature")
    return a
}


func (a *InputSignature) Set(name string, value interface{}) *InputSignature {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * clearValueOnHidden
 */
func (a *InputSignature) ClearValueOnHidden(value interface{}) *InputSignature {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * confirmBtnIcon
 */
func (a *InputSignature) ConfirmBtnIcon(value interface{}) *InputSignature {
    a.Set("confirmBtnIcon", value)
    return a
}

/**
 * placeholder
 */
func (a *InputSignature) Placeholder(value interface{}) *InputSignature {
    a.Set("placeholder", value)
    return a
}

/**
 * validations
 */
func (a *InputSignature) Validations(value interface{}) *InputSignature {
    a.Set("validations", value)
    return a
}

/**
 * submitOnChange
 */
func (a *InputSignature) SubmitOnChange(value interface{}) *InputSignature {
    a.Set("submitOnChange", value)
    return a
}

/**
 * visibleOn
 */
func (a *InputSignature) VisibleOn(value interface{}) *InputSignature {
    a.Set("visibleOn", value)
    return a
}

/**
 * ebmedCancelLabel
 */
func (a *InputSignature) EbmedCancelLabel(value interface{}) *InputSignature {
    a.Set("ebmedCancelLabel", value)
    return a
}

/**
 * embedBtnIcon
 */
func (a *InputSignature) EmbedBtnIcon(value interface{}) *InputSignature {
    a.Set("embedBtnIcon", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *InputSignature) DescriptionClassName(value interface{}) *InputSignature {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * inline
 */
func (a *InputSignature) Inline(value interface{}) *InputSignature {
    a.Set("inline", value)
    return a
}

/**
 * initAutoFill
 */
func (a *InputSignature) InitAutoFill(value interface{}) *InputSignature {
    a.Set("initAutoFill", value)
    return a
}

/**
 * hidden
 */
func (a *InputSignature) Hidden(value interface{}) *InputSignature {
    a.Set("hidden", value)
    return a
}

/**
 * size
 */
func (a *InputSignature) Size(value interface{}) *InputSignature {
    a.Set("size", value)
    return a
}

/**
 * description
 */
func (a *InputSignature) Description(value interface{}) *InputSignature {
    a.Set("description", value)
    return a
}

/**
 * disabledOn
 */
func (a *InputSignature) DisabledOn(value interface{}) *InputSignature {
    a.Set("disabledOn", value)
    return a
}

/**
 * labelClassName
 */
func (a *InputSignature) LabelClassName(value interface{}) *InputSignature {
    a.Set("labelClassName", value)
    return a
}

/**
 * desc
 */
func (a *InputSignature) Desc(value interface{}) *InputSignature {
    a.Set("desc", value)
    return a
}

/**
 * disabled
 */
func (a *InputSignature) Disabled(value interface{}) *InputSignature {
    a.Set("disabled", value)
    return a
}

/**
 * id
 */
func (a *InputSignature) Id(value interface{}) *InputSignature {
    a.Set("id", value)
    return a
}

/**
 * staticOn
 */
func (a *InputSignature) StaticOn(value interface{}) *InputSignature {
    a.Set("staticOn", value)
    return a
}

/**
 * height
 */
func (a *InputSignature) Height(value interface{}) *InputSignature {
    a.Set("height", value)
    return a
}

/**
 * clearBtnIcon
 */
func (a *InputSignature) ClearBtnIcon(value interface{}) *InputSignature {
    a.Set("clearBtnIcon", value)
    return a
}

/**
 * mode
 */
func (a *InputSignature) Mode(value interface{}) *InputSignature {
    a.Set("mode", value)
    return a
}

/**
 * validateOnChange
 */
func (a *InputSignature) ValidateOnChange(value interface{}) *InputSignature {
    a.Set("validateOnChange", value)
    return a
}

/**
 * labelWidth
 */
func (a *InputSignature) LabelWidth(value interface{}) *InputSignature {
    a.Set("labelWidth", value)
    return a
}

/**
 * labelRemark
 */
func (a *InputSignature) LabelRemark(value interface{}) *InputSignature {
    a.Set("labelRemark", value)
    return a
}

/**
 * staticClassName
 */
func (a *InputSignature) StaticClassName(value interface{}) *InputSignature {
    a.Set("staticClassName", value)
    return a
}

/**
 * color
 */
func (a *InputSignature) Color(value interface{}) *InputSignature {
    a.Set("color", value)
    return a
}

/**
 * inputClassName
 */
func (a *InputSignature) InputClassName(value interface{}) *InputSignature {
    a.Set("inputClassName", value)
    return a
}

/**
 * autoFill
 */
func (a *InputSignature) AutoFill(value interface{}) *InputSignature {
    a.Set("autoFill", value)
    return a
}

/**
 * width
 */
func (a *InputSignature) Width(value interface{}) *InputSignature {
    a.Set("width", value)
    return a
}

/**
 * bgColor
 */
func (a *InputSignature) BgColor(value interface{}) *InputSignature {
    a.Set("bgColor", value)
    return a
}

/**
 * validationErrors
 */
func (a *InputSignature) ValidationErrors(value interface{}) *InputSignature {
    a.Set("validationErrors", value)
    return a
}

/**
 * validateApi
 */
func (a *InputSignature) ValidateApi(value interface{}) *InputSignature {
    a.Set("validateApi", value)
    return a
}

/**
 */
func (a *InputSignature) Type(value interface{}) *InputSignature {
    a.Set("type", value)
    return a
}

/**
 * undoBtnLabel
 */
func (a *InputSignature) UndoBtnLabel(value interface{}) *InputSignature {
    a.Set("undoBtnLabel", value)
    return a
}

/**
 * horizontal
 */
func (a *InputSignature) Horizontal(value interface{}) *InputSignature {
    a.Set("horizontal", value)
    return a
}

/**
 * static
 */
func (a *InputSignature) Static(value interface{}) *InputSignature {
    a.Set("static", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *InputSignature) StaticLabelClassName(value interface{}) *InputSignature {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *InputSignature) StaticInputClassName(value interface{}) *InputSignature {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * style
 */
func (a *InputSignature) Style(value interface{}) *InputSignature {
    a.Set("style", value)
    return a
}

/**
 * uploadApi
 */
func (a *InputSignature) UploadApi(value interface{}) *InputSignature {
    a.Set("uploadApi", value)
    return a
}

/**
 * labelAlign
 */
func (a *InputSignature) LabelAlign(value interface{}) *InputSignature {
    a.Set("labelAlign", value)
    return a
}

/**
 * hint
 */
func (a *InputSignature) Hint(value interface{}) *InputSignature {
    a.Set("hint", value)
    return a
}

/**
 * hiddenOn
 */
func (a *InputSignature) HiddenOn(value interface{}) *InputSignature {
    a.Set("hiddenOn", value)
    return a
}

/**
 * onEvent
 */
func (a *InputSignature) OnEvent(value interface{}) *InputSignature {
    a.Set("onEvent", value)
    return a
}

/**
 * confirmBtnLabel
 */
func (a *InputSignature) ConfirmBtnLabel(value interface{}) *InputSignature {
    a.Set("confirmBtnLabel", value)
    return a
}

/**
 * embedConfirmIcon
 */
func (a *InputSignature) EmbedConfirmIcon(value interface{}) *InputSignature {
    a.Set("embedConfirmIcon", value)
    return a
}

/**
 * labelOverflow
 */
func (a *InputSignature) LabelOverflow(value interface{}) *InputSignature {
    a.Set("labelOverflow", value)
    return a
}

/**
 * value
 */
func (a *InputSignature) Value(value interface{}) *InputSignature {
    a.Set("value", value)
    return a
}

/**
 * name
 */
func (a *InputSignature) Name(value interface{}) *InputSignature {
    a.Set("name", value)
    return a
}

/**
 * staticSchema
 */
func (a *InputSignature) StaticSchema(value interface{}) *InputSignature {
    a.Set("staticSchema", value)
    return a
}

/**
 * clearBtnLabel
 */
func (a *InputSignature) ClearBtnLabel(value interface{}) *InputSignature {
    a.Set("clearBtnLabel", value)
    return a
}

/**
 * extraName
 */
func (a *InputSignature) ExtraName(value interface{}) *InputSignature {
    a.Set("extraName", value)
    return a
}

/**
 * row
 */
func (a *InputSignature) Row(value interface{}) *InputSignature {
    a.Set("row", value)
    return a
}

/**
 * undoBtnIcon
 */
func (a *InputSignature) UndoBtnIcon(value interface{}) *InputSignature {
    a.Set("undoBtnIcon", value)
    return a
}

/**
 * embedConfirmLabel
 */
func (a *InputSignature) EmbedConfirmLabel(value interface{}) *InputSignature {
    a.Set("embedConfirmLabel", value)
    return a
}

/**
 * label
 */
func (a *InputSignature) Label(value interface{}) *InputSignature {
    a.Set("label", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *InputSignature) ReadOnlyOn(value interface{}) *InputSignature {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *InputSignature) StaticPlaceholder(value interface{}) *InputSignature {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * editorSetting
 */
func (a *InputSignature) EditorSetting(value interface{}) *InputSignature {
    a.Set("editorSetting", value)
    return a
}

/**
 * useMobileUI
 */
func (a *InputSignature) UseMobileUI(value interface{}) *InputSignature {
    a.Set("useMobileUI", value)
    return a
}

/**
 * remark
 */
func (a *InputSignature) Remark(value interface{}) *InputSignature {
    a.Set("remark", value)
    return a
}

/**
 * readOnly
 */
func (a *InputSignature) ReadOnly(value interface{}) *InputSignature {
    a.Set("readOnly", value)
    return a
}

/**
 * required
 */
func (a *InputSignature) Required(value interface{}) *InputSignature {
    a.Set("required", value)
    return a
}

/**
 * className
 */
func (a *InputSignature) ClassName(value interface{}) *InputSignature {
    a.Set("className", value)
    return a
}

/**
 * visible
 */
func (a *InputSignature) Visible(value interface{}) *InputSignature {
    a.Set("visible", value)
    return a
}

/**
 * embed
 */
func (a *InputSignature) Embed(value interface{}) *InputSignature {
    a.Set("embed", value)
    return a
}

/**
 * ebmedCancelIcon
 */
func (a *InputSignature) EbmedCancelIcon(value interface{}) *InputSignature {
    a.Set("ebmedCancelIcon", value)
    return a
}

/**
 * embedBtnLabel
 */
func (a *InputSignature) EmbedBtnLabel(value interface{}) *InputSignature {
    a.Set("embedBtnLabel", value)
    return a
}
