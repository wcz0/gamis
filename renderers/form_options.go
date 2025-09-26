package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type FormOptions struct {
	*BaseRenderer
}

func NewFormOptions() *FormOptions {
    a := &FormOptions{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}


func (a *FormOptions) Set(name string, value interface{}) *FormOptions {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * initFetchOn
 */
func (a *FormOptions) InitFetchOn(value interface{}) *FormOptions {
    a.Set("initFetchOn", value)
    return a
}

/**
 * editApi
 */
func (a *FormOptions) EditApi(value interface{}) *FormOptions {
    a.Set("editApi", value)
    return a
}

/**
 * label
 */
func (a *FormOptions) Label(value interface{}) *FormOptions {
    a.Set("label", value)
    return a
}

/**
 * validateOnChange
 */
func (a *FormOptions) ValidateOnChange(value interface{}) *FormOptions {
    a.Set("validateOnChange", value)
    return a
}

/**
 * mode
 */
func (a *FormOptions) Mode(value interface{}) *FormOptions {
    a.Set("mode", value)
    return a
}

/**
 * horizontal
 */
func (a *FormOptions) Horizontal(value interface{}) *FormOptions {
    a.Set("horizontal", value)
    return a
}

/**
 * staticClassName
 */
func (a *FormOptions) StaticClassName(value interface{}) *FormOptions {
    a.Set("staticClassName", value)
    return a
}

/**
 * remark
 */
func (a *FormOptions) Remark(value interface{}) *FormOptions {
    a.Set("remark", value)
    return a
}

/**
 * readOnly
 */
func (a *FormOptions) ReadOnly(value interface{}) *FormOptions {
    a.Set("readOnly", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *FormOptions) DescriptionClassName(value interface{}) *FormOptions {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * hiddenOn
 */
func (a *FormOptions) HiddenOn(value interface{}) *FormOptions {
    a.Set("hiddenOn", value)
    return a
}

/**
 * onEvent
 */
func (a *FormOptions) OnEvent(value interface{}) *FormOptions {
    a.Set("onEvent", value)
    return a
}

/**
 * staticSchema
 */
func (a *FormOptions) StaticSchema(value interface{}) *FormOptions {
    a.Set("staticSchema", value)
    return a
}

/**
 * selectFirst
 */
func (a *FormOptions) SelectFirst(value interface{}) *FormOptions {
    a.Set("selectFirst", value)
    return a
}

/**
 * resetValue
 */
func (a *FormOptions) ResetValue(value interface{}) *FormOptions {
    a.Set("resetValue", value)
    return a
}

/**
 * labelClassName
 */
func (a *FormOptions) LabelClassName(value interface{}) *FormOptions {
    a.Set("labelClassName", value)
    return a
}

/**
 * autoFill
 */
func (a *FormOptions) AutoFill(value interface{}) *FormOptions {
    a.Set("autoFill", value)
    return a
}

/**
 * creatable
 */
func (a *FormOptions) Creatable(value interface{}) *FormOptions {
    a.Set("creatable", value)
    return a
}

/**
 * visible
 */
func (a *FormOptions) Visible(value interface{}) *FormOptions {
    a.Set("visible", value)
    return a
}

/**
 * static
 */
func (a *FormOptions) Static(value interface{}) *FormOptions {
    a.Set("static", value)
    return a
}

/**
 * type
 */
func (a *FormOptions) Type(value interface{}) *FormOptions {
    a.Set("type", value)
    return a
}

/**
 * options
 */
func (a *FormOptions) Options(value interface{}) *FormOptions {
    a.Set("options", value)
    return a
}

/**
 * editControls
 */
func (a *FormOptions) EditControls(value interface{}) *FormOptions {
    a.Set("editControls", value)
    return a
}

/**
 * className
 */
func (a *FormOptions) ClassName(value interface{}) *FormOptions {
    a.Set("className", value)
    return a
}

/**
 * deferField
 */
func (a *FormOptions) DeferField(value interface{}) *FormOptions {
    a.Set("deferField", value)
    return a
}

/**
 * addApi
 */
func (a *FormOptions) AddApi(value interface{}) *FormOptions {
    a.Set("addApi", value)
    return a
}

/**
 * validationErrors
 */
func (a *FormOptions) ValidationErrors(value interface{}) *FormOptions {
    a.Set("validationErrors", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *FormOptions) StaticPlaceholder(value interface{}) *FormOptions {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *FormOptions) StaticLabelClassName(value interface{}) *FormOptions {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *FormOptions) StaticInputClassName(value interface{}) *FormOptions {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * multiple
 */
func (a *FormOptions) Multiple(value interface{}) *FormOptions {
    a.Set("multiple", value)
    return a
}

/**
 * delimiter
 */
func (a *FormOptions) Delimiter(value interface{}) *FormOptions {
    a.Set("delimiter", value)
    return a
}

/**
 * editDialog
 */
func (a *FormOptions) EditDialog(value interface{}) *FormOptions {
    a.Set("editDialog", value)
    return a
}

/**
 * desc
 */
func (a *FormOptions) Desc(value interface{}) *FormOptions {
    a.Set("desc", value)
    return a
}

/**
 * required
 */
func (a *FormOptions) Required(value interface{}) *FormOptions {
    a.Set("required", value)
    return a
}

/**
 * addDialog
 */
func (a *FormOptions) AddDialog(value interface{}) *FormOptions {
    a.Set("addDialog", value)
    return a
}

/**
 * name
 */
func (a *FormOptions) Name(value interface{}) *FormOptions {
    a.Set("name", value)
    return a
}

/**
 * value
 */
func (a *FormOptions) Value(value interface{}) *FormOptions {
    a.Set("value", value)
    return a
}

/**
 * initFetch
 */
func (a *FormOptions) InitFetch(value interface{}) *FormOptions {
    a.Set("initFetch", value)
    return a
}

/**
 * joinValues
 */
func (a *FormOptions) JoinValues(value interface{}) *FormOptions {
    a.Set("joinValues", value)
    return a
}

/**
 * editable
 */
func (a *FormOptions) Editable(value interface{}) *FormOptions {
    a.Set("editable", value)
    return a
}

/**
 * hidden
 */
func (a *FormOptions) Hidden(value interface{}) *FormOptions {
    a.Set("hidden", value)
    return a
}

/**
 * id
 */
func (a *FormOptions) Id(value interface{}) *FormOptions {
    a.Set("id", value)
    return a
}

/**
 * deferApi
 */
func (a *FormOptions) DeferApi(value interface{}) *FormOptions {
    a.Set("deferApi", value)
    return a
}

/**
 * labelRemark
 */
func (a *FormOptions) LabelRemark(value interface{}) *FormOptions {
    a.Set("labelRemark", value)
    return a
}

/**
 * validations
 */
func (a *FormOptions) Validations(value interface{}) *FormOptions {
    a.Set("validations", value)
    return a
}

/**
 * size
 */
func (a *FormOptions) Size(value interface{}) *FormOptions {
    a.Set("size", value)
    return a
}

/**
 * clearable
 */
func (a *FormOptions) Clearable(value interface{}) *FormOptions {
    a.Set("clearable", value)
    return a
}

/**
 * createBtnLabel
 */
func (a *FormOptions) CreateBtnLabel(value interface{}) *FormOptions {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * labelOverflow
 */
func (a *FormOptions) LabelOverflow(value interface{}) *FormOptions {
    a.Set("labelOverflow", value)
    return a
}

/**
 * extraName
 */
func (a *FormOptions) ExtraName(value interface{}) *FormOptions {
    a.Set("extraName", value)
    return a
}

/**
 * submitOnChange
 */
func (a *FormOptions) SubmitOnChange(value interface{}) *FormOptions {
    a.Set("submitOnChange", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *FormOptions) ReadOnlyOn(value interface{}) *FormOptions {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * description
 */
func (a *FormOptions) Description(value interface{}) *FormOptions {
    a.Set("description", value)
    return a
}

/**
 * disabled
 */
func (a *FormOptions) Disabled(value interface{}) *FormOptions {
    a.Set("disabled", value)
    return a
}

/**
 * extractValue
 */
func (a *FormOptions) ExtractValue(value interface{}) *FormOptions {
    a.Set("extractValue", value)
    return a
}

/**
 * deleteApi
 */
func (a *FormOptions) DeleteApi(value interface{}) *FormOptions {
    a.Set("deleteApi", value)
    return a
}

/**
 * deleteConfirmText
 */
func (a *FormOptions) DeleteConfirmText(value interface{}) *FormOptions {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * hint
 */
func (a *FormOptions) Hint(value interface{}) *FormOptions {
    a.Set("hint", value)
    return a
}

/**
 * editorSetting
 */
func (a *FormOptions) EditorSetting(value interface{}) *FormOptions {
    a.Set("editorSetting", value)
    return a
}

/**
 * useMobileUI
 */
func (a *FormOptions) UseMobileUI(value interface{}) *FormOptions {
    a.Set("useMobileUI", value)
    return a
}

/**
 * valuesNoWrap
 */
func (a *FormOptions) ValuesNoWrap(value interface{}) *FormOptions {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * removable
 */
func (a *FormOptions) Removable(value interface{}) *FormOptions {
    a.Set("removable", value)
    return a
}

/**
 * inputClassName
 */
func (a *FormOptions) InputClassName(value interface{}) *FormOptions {
    a.Set("inputClassName", value)
    return a
}

/**
 * placeholder
 */
func (a *FormOptions) Placeholder(value interface{}) *FormOptions {
    a.Set("placeholder", value)
    return a
}

/**
 * validateApi
 */
func (a *FormOptions) ValidateApi(value interface{}) *FormOptions {
    a.Set("validateApi", value)
    return a
}

/**
 * style
 */
func (a *FormOptions) Style(value interface{}) *FormOptions {
    a.Set("style", value)
    return a
}

/**
 * source
 */
func (a *FormOptions) Source(value interface{}) *FormOptions {
    a.Set("source", value)
    return a
}

/**
 * checkAll
 */
func (a *FormOptions) CheckAll(value interface{}) *FormOptions {
    a.Set("checkAll", value)
    return a
}

/**
 * clearValueOnSourceChange
 */
func (a *FormOptions) ClearValueOnSourceChange(value interface{}) *FormOptions {
    a.Set("clearValueOnSourceChange", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *FormOptions) ClearValueOnHidden(value interface{}) *FormOptions {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * visibleOn
 */
func (a *FormOptions) VisibleOn(value interface{}) *FormOptions {
    a.Set("visibleOn", value)
    return a
}

/**
 * addControls
 */
func (a *FormOptions) AddControls(value interface{}) *FormOptions {
    a.Set("addControls", value)
    return a
}

/**
 * initAutoFill
 */
func (a *FormOptions) InitAutoFill(value interface{}) *FormOptions {
    a.Set("initAutoFill", value)
    return a
}

/**
 * row
 */
func (a *FormOptions) Row(value interface{}) *FormOptions {
    a.Set("row", value)
    return a
}

/**
 * labelAlign
 */
func (a *FormOptions) LabelAlign(value interface{}) *FormOptions {
    a.Set("labelAlign", value)
    return a
}

/**
 * labelWidth
 */
func (a *FormOptions) LabelWidth(value interface{}) *FormOptions {
    a.Set("labelWidth", value)
    return a
}

/**
 * inline
 */
func (a *FormOptions) Inline(value interface{}) *FormOptions {
    a.Set("inline", value)
    return a
}

/**
 * disabledOn
 */
func (a *FormOptions) DisabledOn(value interface{}) *FormOptions {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticOn
 */
func (a *FormOptions) StaticOn(value interface{}) *FormOptions {
    a.Set("staticOn", value)
    return a
}
