package renderers


/**
 * Text 文本输入框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/input-text

 * @author wcz0
 * @version 6.2.2
 */
type TextControl struct {
	*BaseRenderer
}

func NewTextControl() *TextControl {
    a := &TextControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-text")
    return a
}


func (a *TextControl) Set(name string, value interface{}) *TextControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * joinValues
 */
func (a *TextControl) JoinValues(value interface{}) *TextControl {
    a.Set("joinValues", value)
    return a
}

/**
 * deleteApi
 */
func (a *TextControl) DeleteApi(value interface{}) *TextControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * inputClassName
 */
func (a *TextControl) InputClassName(value interface{}) *TextControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * required
 */
func (a *TextControl) Required(value interface{}) *TextControl {
    a.Set("required", value)
    return a
}

/**
 * disabledOn
 */
func (a *TextControl) DisabledOn(value interface{}) *TextControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *TextControl) StaticClassName(value interface{}) *TextControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * clearable
 */
func (a *TextControl) Clearable(value interface{}) *TextControl {
    a.Set("clearable", value)
    return a
}

/**
 * placeholder
 */
func (a *TextControl) Placeholder(value interface{}) *TextControl {
    a.Set("placeholder", value)
    return a
}

/**
 * transform
 */
func (a *TextControl) Transform(value interface{}) *TextControl {
    a.Set("transform", value)
    return a
}

/**
 * deferApi
 */
func (a *TextControl) DeferApi(value interface{}) *TextControl {
    a.Set("deferApi", value)
    return a
}

/**
 * addApi
 */
func (a *TextControl) AddApi(value interface{}) *TextControl {
    a.Set("addApi", value)
    return a
}

/**
 * editControls
 */
func (a *TextControl) EditControls(value interface{}) *TextControl {
    a.Set("editControls", value)
    return a
}

/**
 * hint
 */
func (a *TextControl) Hint(value interface{}) *TextControl {
    a.Set("hint", value)
    return a
}

/**
 * validateApi
 */
func (a *TextControl) ValidateApi(value interface{}) *TextControl {
    a.Set("validateApi", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *TextControl) StaticPlaceholder(value interface{}) *TextControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * initFetch
 */
func (a *TextControl) InitFetch(value interface{}) *TextControl {
    a.Set("initFetch", value)
    return a
}

/**
 * editable
 */
func (a *TextControl) Editable(value interface{}) *TextControl {
    a.Set("editable", value)
    return a
}

/**
 * horizontal
 */
func (a *TextControl) Horizontal(value interface{}) *TextControl {
    a.Set("horizontal", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *TextControl) ClearValueOnHidden(value interface{}) *TextControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * hidden
 */
func (a *TextControl) Hidden(value interface{}) *TextControl {
    a.Set("hidden", value)
    return a
}

/**
 * staticOn
 */
func (a *TextControl) StaticOn(value interface{}) *TextControl {
    a.Set("staticOn", value)
    return a
}

/**
 * borderMode
 */
func (a *TextControl) BorderMode(value interface{}) *TextControl {
    a.Set("borderMode", value)
    return a
}

/**
 * selectFirst
 */
func (a *TextControl) SelectFirst(value interface{}) *TextControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * labelOverflow
 */
func (a *TextControl) LabelOverflow(value interface{}) *TextControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * extraName
 */
func (a *TextControl) ExtraName(value interface{}) *TextControl {
    a.Set("extraName", value)
    return a
}

/**
 * submitOnChange
 */
func (a *TextControl) SubmitOnChange(value interface{}) *TextControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * inline
 */
func (a *TextControl) Inline(value interface{}) *TextControl {
    a.Set("inline", value)
    return a
}

/**
 * disabled
 */
func (a *TextControl) Disabled(value interface{}) *TextControl {
    a.Set("disabled", value)
    return a
}

/**
 * addOn
 */
func (a *TextControl) AddOn(value interface{}) *TextControl {
    a.Set("addOn", value)
    return a
}

/**
 * clearValueOnEmpty
 */
func (a *TextControl) ClearValueOnEmpty(value interface{}) *TextControl {
    a.Set("clearValueOnEmpty", value)
    return a
}

/**
 * createBtnLabel
 */
func (a *TextControl) CreateBtnLabel(value interface{}) *TextControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * clearValueOnSourceChange
 */
func (a *TextControl) ClearValueOnSourceChange(value interface{}) *TextControl {
    a.Set("clearValueOnSourceChange", value)
    return a
}

/**
 * minLength
 */
func (a *TextControl) MinLength(value interface{}) *TextControl {
    a.Set("minLength", value)
    return a
}

/**
 * addControls
 */
func (a *TextControl) AddControls(value interface{}) *TextControl {
    a.Set("addControls", value)
    return a
}

/**
 * deleteConfirmText
 */
func (a *TextControl) DeleteConfirmText(value interface{}) *TextControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * description
 */
func (a *TextControl) Description(value interface{}) *TextControl {
    a.Set("description", value)
    return a
}

/**
 * className
 */
func (a *TextControl) ClassName(value interface{}) *TextControl {
    a.Set("className", value)
    return a
}

/**
 * visibleOn
 */
func (a *TextControl) VisibleOn(value interface{}) *TextControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * useMobileUI
 */
func (a *TextControl) UseMobileUI(value interface{}) *TextControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * trimContents
 */
func (a *TextControl) TrimContents(value interface{}) *TextControl {
    a.Set("trimContents", value)
    return a
}

/**
 * nativeAutoComplete
 */
func (a *TextControl) NativeAutoComplete(value interface{}) *TextControl {
    a.Set("nativeAutoComplete", value)
    return a
}

/**
 * creatable
 */
func (a *TextControl) Creatable(value interface{}) *TextControl {
    a.Set("creatable", value)
    return a
}

/**
 * labelAlign
 */
func (a *TextControl) LabelAlign(value interface{}) *TextControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * editorSetting
 */
func (a *TextControl) EditorSetting(value interface{}) *TextControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * maxLength
 */
func (a *TextControl) MaxLength(value interface{}) *TextControl {
    a.Set("maxLength", value)
    return a
}

/**
 * prefix
 */
func (a *TextControl) Prefix(value interface{}) *TextControl {
    a.Set("prefix", value)
    return a
}

/**
 * inputControlClassName
 */
func (a *TextControl) InputControlClassName(value interface{}) *TextControl {
    a.Set("inputControlClassName", value)
    return a
}

/**
 * checkAll
 */
func (a *TextControl) CheckAll(value interface{}) *TextControl {
    a.Set("checkAll", value)
    return a
}

/**
 * delimiter
 */
func (a *TextControl) Delimiter(value interface{}) *TextControl {
    a.Set("delimiter", value)
    return a
}

/**
 * label
 */
func (a *TextControl) Label(value interface{}) *TextControl {
    a.Set("label", value)
    return a
}

/**
 * remark
 */
func (a *TextControl) Remark(value interface{}) *TextControl {
    a.Set("remark", value)
    return a
}

/**
 * validationErrors
 */
func (a *TextControl) ValidationErrors(value interface{}) *TextControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * visible
 */
func (a *TextControl) Visible(value interface{}) *TextControl {
    a.Set("visible", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *TextControl) StaticLabelClassName(value interface{}) *TextControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * showCounter
 */
func (a *TextControl) ShowCounter(value interface{}) *TextControl {
    a.Set("showCounter", value)
    return a
}

/**
 * initFetchOn
 */
func (a *TextControl) InitFetchOn(value interface{}) *TextControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * extractValue
 */
func (a *TextControl) ExtractValue(value interface{}) *TextControl {
    a.Set("extractValue", value)
    return a
}

/**
 * deferField
 */
func (a *TextControl) DeferField(value interface{}) *TextControl {
    a.Set("deferField", value)
    return a
}

/**
 * labelWidth
 */
func (a *TextControl) LabelWidth(value interface{}) *TextControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * validations
 */
func (a *TextControl) Validations(value interface{}) *TextControl {
    a.Set("validations", value)
    return a
}

/**
 * hiddenOn
 */
func (a *TextControl) HiddenOn(value interface{}) *TextControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * autoComplete
 */
func (a *TextControl) AutoComplete(value interface{}) *TextControl {
    a.Set("autoComplete", value)
    return a
}

/**
 * nativeInputClassName
 */
func (a *TextControl) NativeInputClassName(value interface{}) *TextControl {
    a.Set("nativeInputClassName", value)
    return a
}

/**
 * initAutoFill
 */
func (a *TextControl) InitAutoFill(value interface{}) *TextControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * staticSchema
 */
func (a *TextControl) StaticSchema(value interface{}) *TextControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * suffix
 */
func (a *TextControl) Suffix(value interface{}) *TextControl {
    a.Set("suffix", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *TextControl) Width(value interface{}) *TextControl {
    a.Set("width", value)
    return a
}

/**
 * options
 */
func (a *TextControl) Options(value interface{}) *TextControl {
    a.Set("options", value)
    return a
}

/**
 * source
 */
func (a *TextControl) Source(value interface{}) *TextControl {
    a.Set("source", value)
    return a
}

/**
 * valuesNoWrap
 */
func (a *TextControl) ValuesNoWrap(value interface{}) *TextControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * labelClassName
 */
func (a *TextControl) LabelClassName(value interface{}) *TextControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * name
 */
func (a *TextControl) Name(value interface{}) *TextControl {
    a.Set("name", value)
    return a
}

/**
 * validateOnChange
 */
func (a *TextControl) ValidateOnChange(value interface{}) *TextControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * desc
 */
func (a *TextControl) Desc(value interface{}) *TextControl {
    a.Set("desc", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *TextControl) StaticInputClassName(value interface{}) *TextControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * editApi
 */
func (a *TextControl) EditApi(value interface{}) *TextControl {
    a.Set("editApi", value)
    return a
}

/**
 * style
 */
func (a *TextControl) Style(value interface{}) *TextControl {
    a.Set("style", value)
    return a
}

/**
 * resetValue
 */
func (a *TextControl) ResetValue(value interface{}) *TextControl {
    a.Set("resetValue", value)
    return a
}

/**
 * addDialog
 */
func (a *TextControl) AddDialog(value interface{}) *TextControl {
    a.Set("addDialog", value)
    return a
}

/**
 * removable
 */
func (a *TextControl) Removable(value interface{}) *TextControl {
    a.Set("removable", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *TextControl) DescriptionClassName(value interface{}) *TextControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * mode
 */
func (a *TextControl) Mode(value interface{}) *TextControl {
    a.Set("mode", value)
    return a
}

/**
 * multiple
 */
func (a *TextControl) Multiple(value interface{}) *TextControl {
    a.Set("multiple", value)
    return a
}

/**
 * labelRemark
 */
func (a *TextControl) LabelRemark(value interface{}) *TextControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * readOnly
 */
func (a *TextControl) ReadOnly(value interface{}) *TextControl {
    a.Set("readOnly", value)
    return a
}

/**
 * autoFill
 */
func (a *TextControl) AutoFill(value interface{}) *TextControl {
    a.Set("autoFill", value)
    return a
}

/**
 * row
 */
func (a *TextControl) Row(value interface{}) *TextControl {
    a.Set("row", value)
    return a
}

/**
 * id
 */
func (a *TextControl) Id(value interface{}) *TextControl {
    a.Set("id", value)
    return a
}

/**
 * static
 */
func (a *TextControl) Static(value interface{}) *TextControl {
    a.Set("static", value)
    return a
}

/**
 * 可选值: input-text | input-email | input-url | input-password | native-date | native-time | native-number
 */
func (a *TextControl) Type(value interface{}) *TextControl {
    a.Set("type", value)
    return a
}

/**
 * editDialog
 */
func (a *TextControl) EditDialog(value interface{}) *TextControl {
    a.Set("editDialog", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *TextControl) ReadOnlyOn(value interface{}) *TextControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * value
 */
func (a *TextControl) Value(value interface{}) *TextControl {
    a.Set("value", value)
    return a
}

/**
 * onEvent
 */
func (a *TextControl) OnEvent(value interface{}) *TextControl {
    a.Set("onEvent", value)
    return a
}

/**
 * size
 */
func (a *TextControl) Size(value interface{}) *TextControl {
    a.Set("size", value)
    return a
}
