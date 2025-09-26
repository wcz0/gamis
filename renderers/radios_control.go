package renderers


/**
 * Radio 单选框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/radios

 * @author wcz0
 * @version 6.2.2
 */
type RadiosControl struct {
	*BaseRenderer
}

func NewRadiosControl() *RadiosControl {
    a := &RadiosControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "radios")
    return a
}


func (a *RadiosControl) Set(name string, value interface{}) *RadiosControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * initFetch
 */
func (a *RadiosControl) InitFetch(value interface{}) *RadiosControl {
    a.Set("initFetch", value)
    return a
}

/**
 * readOnly
 */
func (a *RadiosControl) ReadOnly(value interface{}) *RadiosControl {
    a.Set("readOnly", value)
    return a
}

/**
 * validationErrors
 */
func (a *RadiosControl) ValidationErrors(value interface{}) *RadiosControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * size
 */
func (a *RadiosControl) Size(value interface{}) *RadiosControl {
    a.Set("size", value)
    return a
}

/**
 * labelClassName
 */
func (a *RadiosControl) LabelClassName(value interface{}) *RadiosControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * extraName
 */
func (a *RadiosControl) ExtraName(value interface{}) *RadiosControl {
    a.Set("extraName", value)
    return a
}

/**
 * mode
 */
func (a *RadiosControl) Mode(value interface{}) *RadiosControl {
    a.Set("mode", value)
    return a
}

/**
 * className
 */
func (a *RadiosControl) ClassName(value interface{}) *RadiosControl {
    a.Set("className", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *RadiosControl) StaticInputClassName(value interface{}) *RadiosControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * resetValue
 */
func (a *RadiosControl) ResetValue(value interface{}) *RadiosControl {
    a.Set("resetValue", value)
    return a
}

/**
 * initFetchOn
 */
func (a *RadiosControl) InitFetchOn(value interface{}) *RadiosControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * valuesNoWrap
 */
func (a *RadiosControl) ValuesNoWrap(value interface{}) *RadiosControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * addControls
 */
func (a *RadiosControl) AddControls(value interface{}) *RadiosControl {
    a.Set("addControls", value)
    return a
}

/**
 * editable
 */
func (a *RadiosControl) Editable(value interface{}) *RadiosControl {
    a.Set("editable", value)
    return a
}

/**
 * placeholder
 */
func (a *RadiosControl) Placeholder(value interface{}) *RadiosControl {
    a.Set("placeholder", value)
    return a
}

/**
 * initAutoFill
 */
func (a *RadiosControl) InitAutoFill(value interface{}) *RadiosControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * label
 */
func (a *RadiosControl) Label(value interface{}) *RadiosControl {
    a.Set("label", value)
    return a
}

/**
 * inputClassName
 */
func (a *RadiosControl) InputClassName(value interface{}) *RadiosControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *RadiosControl) StaticLabelClassName(value interface{}) *RadiosControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *RadiosControl) Type(value interface{}) *RadiosControl {
    a.Set("type", value)
    return a
}

/**
 * extractValue
 */
func (a *RadiosControl) ExtractValue(value interface{}) *RadiosControl {
    a.Set("extractValue", value)
    return a
}

/**
 * row
 */
func (a *RadiosControl) Row(value interface{}) *RadiosControl {
    a.Set("row", value)
    return a
}

/**
 * visible
 */
func (a *RadiosControl) Visible(value interface{}) *RadiosControl {
    a.Set("visible", value)
    return a
}

/**
 * onEvent
 */
func (a *RadiosControl) OnEvent(value interface{}) *RadiosControl {
    a.Set("onEvent", value)
    return a
}

/**
 * useMobileUI
 */
func (a *RadiosControl) UseMobileUI(value interface{}) *RadiosControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * delimiter
 */
func (a *RadiosControl) Delimiter(value interface{}) *RadiosControl {
    a.Set("delimiter", value)
    return a
}

/**
 * creatable
 */
func (a *RadiosControl) Creatable(value interface{}) *RadiosControl {
    a.Set("creatable", value)
    return a
}

/**
 * editDialog
 */
func (a *RadiosControl) EditDialog(value interface{}) *RadiosControl {
    a.Set("editDialog", value)
    return a
}

/**
 * hiddenOn
 */
func (a *RadiosControl) HiddenOn(value interface{}) *RadiosControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * joinValues
 */
func (a *RadiosControl) JoinValues(value interface{}) *RadiosControl {
    a.Set("joinValues", value)
    return a
}

/**
 * addDialog
 */
func (a *RadiosControl) AddDialog(value interface{}) *RadiosControl {
    a.Set("addDialog", value)
    return a
}

/**
 * labelOverflow
 */
func (a *RadiosControl) LabelOverflow(value interface{}) *RadiosControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * hint
 */
func (a *RadiosControl) Hint(value interface{}) *RadiosControl {
    a.Set("hint", value)
    return a
}

/**
 * validateApi
 */
func (a *RadiosControl) ValidateApi(value interface{}) *RadiosControl {
    a.Set("validateApi", value)
    return a
}

/**
 * disabledOn
 */
func (a *RadiosControl) DisabledOn(value interface{}) *RadiosControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * style
 */
func (a *RadiosControl) Style(value interface{}) *RadiosControl {
    a.Set("style", value)
    return a
}

/**
 * removable
 */
func (a *RadiosControl) Removable(value interface{}) *RadiosControl {
    a.Set("removable", value)
    return a
}

/**
 * labelRemark
 */
func (a *RadiosControl) LabelRemark(value interface{}) *RadiosControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *RadiosControl) Width(value interface{}) *RadiosControl {
    a.Set("width", value)
    return a
}

/**
 * deferField
 */
func (a *RadiosControl) DeferField(value interface{}) *RadiosControl {
    a.Set("deferField", value)
    return a
}

/**
 * source
 */
func (a *RadiosControl) Source(value interface{}) *RadiosControl {
    a.Set("source", value)
    return a
}

/**
 * editApi
 */
func (a *RadiosControl) EditApi(value interface{}) *RadiosControl {
    a.Set("editApi", value)
    return a
}

/**
 * name
 */
func (a *RadiosControl) Name(value interface{}) *RadiosControl {
    a.Set("name", value)
    return a
}

/**
 * deleteConfirmText
 */
func (a *RadiosControl) DeleteConfirmText(value interface{}) *RadiosControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *RadiosControl) StaticPlaceholder(value interface{}) *RadiosControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * value
 */
func (a *RadiosControl) Value(value interface{}) *RadiosControl {
    a.Set("value", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *RadiosControl) ClearValueOnHidden(value interface{}) *RadiosControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * options
 */
func (a *RadiosControl) Options(value interface{}) *RadiosControl {
    a.Set("options", value)
    return a
}

/**
 * checkAll
 */
func (a *RadiosControl) CheckAll(value interface{}) *RadiosControl {
    a.Set("checkAll", value)
    return a
}

/**
 * clearValueOnSourceChange
 */
func (a *RadiosControl) ClearValueOnSourceChange(value interface{}) *RadiosControl {
    a.Set("clearValueOnSourceChange", value)
    return a
}

/**
 * remark
 */
func (a *RadiosControl) Remark(value interface{}) *RadiosControl {
    a.Set("remark", value)
    return a
}

/**
 * submitOnChange
 */
func (a *RadiosControl) SubmitOnChange(value interface{}) *RadiosControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * inline
 */
func (a *RadiosControl) Inline(value interface{}) *RadiosControl {
    a.Set("inline", value)
    return a
}

/**
 * clearable
 */
func (a *RadiosControl) Clearable(value interface{}) *RadiosControl {
    a.Set("clearable", value)
    return a
}

/**
 * createBtnLabel
 */
func (a *RadiosControl) CreateBtnLabel(value interface{}) *RadiosControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * autoFill
 */
func (a *RadiosControl) AutoFill(value interface{}) *RadiosControl {
    a.Set("autoFill", value)
    return a
}

/**
 * hidden
 */
func (a *RadiosControl) Hidden(value interface{}) *RadiosControl {
    a.Set("hidden", value)
    return a
}

/**
 * editorSetting
 */
func (a *RadiosControl) EditorSetting(value interface{}) *RadiosControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * columnsCount
 */
func (a *RadiosControl) ColumnsCount(value interface{}) *RadiosControl {
    a.Set("columnsCount", value)
    return a
}

/**
 * selectFirst
 */
func (a *RadiosControl) SelectFirst(value interface{}) *RadiosControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * deferApi
 */
func (a *RadiosControl) DeferApi(value interface{}) *RadiosControl {
    a.Set("deferApi", value)
    return a
}

/**
 * desc
 */
func (a *RadiosControl) Desc(value interface{}) *RadiosControl {
    a.Set("desc", value)
    return a
}

/**
 * required
 */
func (a *RadiosControl) Required(value interface{}) *RadiosControl {
    a.Set("required", value)
    return a
}

/**
 * visibleOn
 */
func (a *RadiosControl) VisibleOn(value interface{}) *RadiosControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * id
 */
func (a *RadiosControl) Id(value interface{}) *RadiosControl {
    a.Set("id", value)
    return a
}

/**
 * static
 */
func (a *RadiosControl) Static(value interface{}) *RadiosControl {
    a.Set("static", value)
    return a
}

/**
 * editControls
 */
func (a *RadiosControl) EditControls(value interface{}) *RadiosControl {
    a.Set("editControls", value)
    return a
}

/**
 * horizontal
 */
func (a *RadiosControl) Horizontal(value interface{}) *RadiosControl {
    a.Set("horizontal", value)
    return a
}

/**
 * staticOn
 */
func (a *RadiosControl) StaticOn(value interface{}) *RadiosControl {
    a.Set("staticOn", value)
    return a
}

/**
 * addApi
 */
func (a *RadiosControl) AddApi(value interface{}) *RadiosControl {
    a.Set("addApi", value)
    return a
}

/**
 * deleteApi
 */
func (a *RadiosControl) DeleteApi(value interface{}) *RadiosControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * labelAlign
 */
func (a *RadiosControl) LabelAlign(value interface{}) *RadiosControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *RadiosControl) ReadOnlyOn(value interface{}) *RadiosControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * description
 */
func (a *RadiosControl) Description(value interface{}) *RadiosControl {
    a.Set("description", value)
    return a
}

/**
 * validations
 */
func (a *RadiosControl) Validations(value interface{}) *RadiosControl {
    a.Set("validations", value)
    return a
}

/**
 * staticClassName
 */
func (a *RadiosControl) StaticClassName(value interface{}) *RadiosControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *RadiosControl) StaticSchema(value interface{}) *RadiosControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * multiple
 */
func (a *RadiosControl) Multiple(value interface{}) *RadiosControl {
    a.Set("multiple", value)
    return a
}

/**
 * labelWidth
 */
func (a *RadiosControl) LabelWidth(value interface{}) *RadiosControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * validateOnChange
 */
func (a *RadiosControl) ValidateOnChange(value interface{}) *RadiosControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *RadiosControl) DescriptionClassName(value interface{}) *RadiosControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * disabled
 */
func (a *RadiosControl) Disabled(value interface{}) *RadiosControl {
    a.Set("disabled", value)
    return a
}
