package renderers


/**
 * UserSelect 移动端人员选择。

 * @author wcz0
 * @version 6.2.2
 */
type UserSelectControl struct {
	*BaseRenderer
}

func NewUserSelectControl() *UserSelectControl {
    a := &UserSelectControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "users-select")
    return a
}


func (a *UserSelectControl) Set(name string, value interface{}) *UserSelectControl {
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
func (a *UserSelectControl) LabelWidth(value interface{}) *UserSelectControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * validateOnChange
 */
func (a *UserSelectControl) ValidateOnChange(value interface{}) *UserSelectControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * disabledOn
 */
func (a *UserSelectControl) DisabledOn(value interface{}) *UserSelectControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *UserSelectControl) StaticPlaceholder(value interface{}) *UserSelectControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * delimiter
 */
func (a *UserSelectControl) Delimiter(value interface{}) *UserSelectControl {
    a.Set("delimiter", value)
    return a
}

/**
 * addApi
 */
func (a *UserSelectControl) AddApi(value interface{}) *UserSelectControl {
    a.Set("addApi", value)
    return a
}

/**
 * remark
 */
func (a *UserSelectControl) Remark(value interface{}) *UserSelectControl {
    a.Set("remark", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *UserSelectControl) ClearValueOnHidden(value interface{}) *UserSelectControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * static
 */
func (a *UserSelectControl) Static(value interface{}) *UserSelectControl {
    a.Set("static", value)
    return a
}

/**
 * label
 */
func (a *UserSelectControl) Label(value interface{}) *UserSelectControl {
    a.Set("label", value)
    return a
}

/**
 * placeholder
 */
func (a *UserSelectControl) Placeholder(value interface{}) *UserSelectControl {
    a.Set("placeholder", value)
    return a
}

/**
 * autoFill
 */
func (a *UserSelectControl) AutoFill(value interface{}) *UserSelectControl {
    a.Set("autoFill", value)
    return a
}

/**
 * row
 */
func (a *UserSelectControl) Row(value interface{}) *UserSelectControl {
    a.Set("row", value)
    return a
}

/**
 * className
 */
func (a *UserSelectControl) ClassName(value interface{}) *UserSelectControl {
    a.Set("className", value)
    return a
}

/**
 * labelRemark
 */
func (a *UserSelectControl) LabelRemark(value interface{}) *UserSelectControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * hint
 */
func (a *UserSelectControl) Hint(value interface{}) *UserSelectControl {
    a.Set("hint", value)
    return a
}

/**
 * staticSchema
 */
func (a *UserSelectControl) StaticSchema(value interface{}) *UserSelectControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * editable
 */
func (a *UserSelectControl) Editable(value interface{}) *UserSelectControl {
    a.Set("editable", value)
    return a
}

/**
 * editDialog
 */
func (a *UserSelectControl) EditDialog(value interface{}) *UserSelectControl {
    a.Set("editDialog", value)
    return a
}

/**
 * multiple
 */
func (a *UserSelectControl) Multiple(value interface{}) *UserSelectControl {
    a.Set("multiple", value)
    return a
}

/**
 * addDialog
 */
func (a *UserSelectControl) AddDialog(value interface{}) *UserSelectControl {
    a.Set("addDialog", value)
    return a
}

/**
 * desc
 */
func (a *UserSelectControl) Desc(value interface{}) *UserSelectControl {
    a.Set("desc", value)
    return a
}

/**
 * validationErrors
 */
func (a *UserSelectControl) ValidationErrors(value interface{}) *UserSelectControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * value
 */
func (a *UserSelectControl) Value(value interface{}) *UserSelectControl {
    a.Set("value", value)
    return a
}

/**
 * staticOn
 */
func (a *UserSelectControl) StaticOn(value interface{}) *UserSelectControl {
    a.Set("staticOn", value)
    return a
}

/**
 * options
 */
func (a *UserSelectControl) Options(value interface{}) *UserSelectControl {
    a.Set("options", value)
    return a
}

/**
 * selectFirst
 */
func (a *UserSelectControl) SelectFirst(value interface{}) *UserSelectControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * checkAll
 */
func (a *UserSelectControl) CheckAll(value interface{}) *UserSelectControl {
    a.Set("checkAll", value)
    return a
}

/**
 * validations
 */
func (a *UserSelectControl) Validations(value interface{}) *UserSelectControl {
    a.Set("validations", value)
    return a
}

/**
 * style
 */
func (a *UserSelectControl) Style(value interface{}) *UserSelectControl {
    a.Set("style", value)
    return a
}

/**
 * useMobileUI
 */
func (a *UserSelectControl) UseMobileUI(value interface{}) *UserSelectControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * editApi
 */
func (a *UserSelectControl) EditApi(value interface{}) *UserSelectControl {
    a.Set("editApi", value)
    return a
}

/**
 * labelAlign
 */
func (a *UserSelectControl) LabelAlign(value interface{}) *UserSelectControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * labelOverflow
 */
func (a *UserSelectControl) LabelOverflow(value interface{}) *UserSelectControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * mode
 */
func (a *UserSelectControl) Mode(value interface{}) *UserSelectControl {
    a.Set("mode", value)
    return a
}

/**
 * disabled
 */
func (a *UserSelectControl) Disabled(value interface{}) *UserSelectControl {
    a.Set("disabled", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *UserSelectControl) Width(value interface{}) *UserSelectControl {
    a.Set("width", value)
    return a
}

/**
 * hiddenOn
 */
func (a *UserSelectControl) HiddenOn(value interface{}) *UserSelectControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * initFetchOn
 */
func (a *UserSelectControl) InitFetchOn(value interface{}) *UserSelectControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * resetValue
 */
func (a *UserSelectControl) ResetValue(value interface{}) *UserSelectControl {
    a.Set("resetValue", value)
    return a
}

/**
 * addControls
 */
func (a *UserSelectControl) AddControls(value interface{}) *UserSelectControl {
    a.Set("addControls", value)
    return a
}

/**
 * name
 */
func (a *UserSelectControl) Name(value interface{}) *UserSelectControl {
    a.Set("name", value)
    return a
}

/**
 * inputClassName
 */
func (a *UserSelectControl) InputClassName(value interface{}) *UserSelectControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * visibleOn
 */
func (a *UserSelectControl) VisibleOn(value interface{}) *UserSelectControl {
    a.Set("visibleOn", value)
    return a
}

/**
 */
func (a *UserSelectControl) Type(value interface{}) *UserSelectControl {
    a.Set("type", value)
    return a
}

/**
 * source
 */
func (a *UserSelectControl) Source(value interface{}) *UserSelectControl {
    a.Set("source", value)
    return a
}

/**
 * deleteApi
 */
func (a *UserSelectControl) DeleteApi(value interface{}) *UserSelectControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * deleteConfirmText
 */
func (a *UserSelectControl) DeleteConfirmText(value interface{}) *UserSelectControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * extraName
 */
func (a *UserSelectControl) ExtraName(value interface{}) *UserSelectControl {
    a.Set("extraName", value)
    return a
}

/**
 * horizontal
 */
func (a *UserSelectControl) Horizontal(value interface{}) *UserSelectControl {
    a.Set("horizontal", value)
    return a
}

/**
 * onEvent
 */
func (a *UserSelectControl) OnEvent(value interface{}) *UserSelectControl {
    a.Set("onEvent", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *UserSelectControl) StaticInputClassName(value interface{}) *UserSelectControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * initFetch
 */
func (a *UserSelectControl) InitFetch(value interface{}) *UserSelectControl {
    a.Set("initFetch", value)
    return a
}

/**
 * creatable
 */
func (a *UserSelectControl) Creatable(value interface{}) *UserSelectControl {
    a.Set("creatable", value)
    return a
}

/**
 * clearValueOnSourceChange
 */
func (a *UserSelectControl) ClearValueOnSourceChange(value interface{}) *UserSelectControl {
    a.Set("clearValueOnSourceChange", value)
    return a
}

/**
 * visible
 */
func (a *UserSelectControl) Visible(value interface{}) *UserSelectControl {
    a.Set("visible", value)
    return a
}

/**
 * size
 */
func (a *UserSelectControl) Size(value interface{}) *UserSelectControl {
    a.Set("size", value)
    return a
}

/**
 * joinValues
 */
func (a *UserSelectControl) JoinValues(value interface{}) *UserSelectControl {
    a.Set("joinValues", value)
    return a
}

/**
 * description
 */
func (a *UserSelectControl) Description(value interface{}) *UserSelectControl {
    a.Set("description", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *UserSelectControl) DescriptionClassName(value interface{}) *UserSelectControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * id
 */
func (a *UserSelectControl) Id(value interface{}) *UserSelectControl {
    a.Set("id", value)
    return a
}

/**
 * createBtnLabel
 */
func (a *UserSelectControl) CreateBtnLabel(value interface{}) *UserSelectControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * deferField
 */
func (a *UserSelectControl) DeferField(value interface{}) *UserSelectControl {
    a.Set("deferField", value)
    return a
}

/**
 * deferApi
 */
func (a *UserSelectControl) DeferApi(value interface{}) *UserSelectControl {
    a.Set("deferApi", value)
    return a
}

/**
 * staticClassName
 */
func (a *UserSelectControl) StaticClassName(value interface{}) *UserSelectControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *UserSelectControl) StaticLabelClassName(value interface{}) *UserSelectControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * validateApi
 */
func (a *UserSelectControl) ValidateApi(value interface{}) *UserSelectControl {
    a.Set("validateApi", value)
    return a
}

/**
 * clearable
 */
func (a *UserSelectControl) Clearable(value interface{}) *UserSelectControl {
    a.Set("clearable", value)
    return a
}

/**
 * labelClassName
 */
func (a *UserSelectControl) LabelClassName(value interface{}) *UserSelectControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * submitOnChange
 */
func (a *UserSelectControl) SubmitOnChange(value interface{}) *UserSelectControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * initAutoFill
 */
func (a *UserSelectControl) InitAutoFill(value interface{}) *UserSelectControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * editControls
 */
func (a *UserSelectControl) EditControls(value interface{}) *UserSelectControl {
    a.Set("editControls", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *UserSelectControl) ReadOnlyOn(value interface{}) *UserSelectControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * inline
 */
func (a *UserSelectControl) Inline(value interface{}) *UserSelectControl {
    a.Set("inline", value)
    return a
}

/**
 * required
 */
func (a *UserSelectControl) Required(value interface{}) *UserSelectControl {
    a.Set("required", value)
    return a
}

/**
 * editorSetting
 */
func (a *UserSelectControl) EditorSetting(value interface{}) *UserSelectControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * extractValue
 */
func (a *UserSelectControl) ExtractValue(value interface{}) *UserSelectControl {
    a.Set("extractValue", value)
    return a
}

/**
 * removable
 */
func (a *UserSelectControl) Removable(value interface{}) *UserSelectControl {
    a.Set("removable", value)
    return a
}

/**
 * readOnly
 */
func (a *UserSelectControl) ReadOnly(value interface{}) *UserSelectControl {
    a.Set("readOnly", value)
    return a
}

/**
 * hidden
 */
func (a *UserSelectControl) Hidden(value interface{}) *UserSelectControl {
    a.Set("hidden", value)
    return a
}

/**
 * valuesNoWrap
 */
func (a *UserSelectControl) ValuesNoWrap(value interface{}) *UserSelectControl {
    a.Set("valuesNoWrap", value)
    return a
}
