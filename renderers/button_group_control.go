package renderers


/**
 * 按钮组控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/button-group

 * @author wcz0
 * @version 6.2.2
 */
type ButtonGroupControl struct {
	*BaseRenderer
}

func NewButtonGroupControl() *ButtonGroupControl {
    a := &ButtonGroupControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "button-group-select")
    return a
}


func (a *ButtonGroupControl) Set(name string, value interface{}) *ButtonGroupControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * inline
 */
func (a *ButtonGroupControl) Inline(value interface{}) *ButtonGroupControl {
    a.Set("inline", value)
    return a
}

/**
 * row
 */
func (a *ButtonGroupControl) Row(value interface{}) *ButtonGroupControl {
    a.Set("row", value)
    return a
}

/**
 * selectFirst
 */
func (a *ButtonGroupControl) SelectFirst(value interface{}) *ButtonGroupControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * addControls
 */
func (a *ButtonGroupControl) AddControls(value interface{}) *ButtonGroupControl {
    a.Set("addControls", value)
    return a
}

/**
 * style
 */
func (a *ButtonGroupControl) Style(value interface{}) *ButtonGroupControl {
    a.Set("style", value)
    return a
}

/**
 * vertical
 */
func (a *ButtonGroupControl) Vertical(value interface{}) *ButtonGroupControl {
    a.Set("vertical", value)
    return a
}

/**
 * mode
 */
func (a *ButtonGroupControl) Mode(value interface{}) *ButtonGroupControl {
    a.Set("mode", value)
    return a
}

/**
 * deleteConfirmText
 */
func (a *ButtonGroupControl) DeleteConfirmText(value interface{}) *ButtonGroupControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * disabled
 */
func (a *ButtonGroupControl) Disabled(value interface{}) *ButtonGroupControl {
    a.Set("disabled", value)
    return a
}

/**
 * id
 */
func (a *ButtonGroupControl) Id(value interface{}) *ButtonGroupControl {
    a.Set("id", value)
    return a
}

/**
 * onEvent
 */
func (a *ButtonGroupControl) OnEvent(value interface{}) *ButtonGroupControl {
    a.Set("onEvent", value)
    return a
}

/**
 * desc
 */
func (a *ButtonGroupControl) Desc(value interface{}) *ButtonGroupControl {
    a.Set("desc", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *ButtonGroupControl) DescriptionClassName(value interface{}) *ButtonGroupControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * addApi
 */
func (a *ButtonGroupControl) AddApi(value interface{}) *ButtonGroupControl {
    a.Set("addApi", value)
    return a
}

/**
 * creatable
 */
func (a *ButtonGroupControl) Creatable(value interface{}) *ButtonGroupControl {
    a.Set("creatable", value)
    return a
}

/**
 * editable
 */
func (a *ButtonGroupControl) Editable(value interface{}) *ButtonGroupControl {
    a.Set("editable", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *ButtonGroupControl) StaticInputClassName(value interface{}) *ButtonGroupControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * readOnly
 */
func (a *ButtonGroupControl) ReadOnly(value interface{}) *ButtonGroupControl {
    a.Set("readOnly", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *ButtonGroupControl) ReadOnlyOn(value interface{}) *ButtonGroupControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * autoFill
 */
func (a *ButtonGroupControl) AutoFill(value interface{}) *ButtonGroupControl {
    a.Set("autoFill", value)
    return a
}

/**
 * editControls
 */
func (a *ButtonGroupControl) EditControls(value interface{}) *ButtonGroupControl {
    a.Set("editControls", value)
    return a
}

/**
 * multiple
 */
func (a *ButtonGroupControl) Multiple(value interface{}) *ButtonGroupControl {
    a.Set("multiple", value)
    return a
}

/**
 * joinValues
 */
func (a *ButtonGroupControl) JoinValues(value interface{}) *ButtonGroupControl {
    a.Set("joinValues", value)
    return a
}

/**
 * valuesNoWrap
 */
func (a *ButtonGroupControl) ValuesNoWrap(value interface{}) *ButtonGroupControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * deferApi
 */
func (a *ButtonGroupControl) DeferApi(value interface{}) *ButtonGroupControl {
    a.Set("deferApi", value)
    return a
}

/**
 * hiddenOn
 */
func (a *ButtonGroupControl) HiddenOn(value interface{}) *ButtonGroupControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * testid
 */
func (a *ButtonGroupControl) Testid(value interface{}) *ButtonGroupControl {
    a.Set("testid", value)
    return a
}

/**
 * labelWidth
 */
func (a *ButtonGroupControl) LabelWidth(value interface{}) *ButtonGroupControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * labelClassName
 */
func (a *ButtonGroupControl) LabelClassName(value interface{}) *ButtonGroupControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * required
 */
func (a *ButtonGroupControl) Required(value interface{}) *ButtonGroupControl {
    a.Set("required", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *ButtonGroupControl) Width(value interface{}) *ButtonGroupControl {
    a.Set("width", value)
    return a
}

/**
 * hint
 */
func (a *ButtonGroupControl) Hint(value interface{}) *ButtonGroupControl {
    a.Set("hint", value)
    return a
}

/**
 * source
 */
func (a *ButtonGroupControl) Source(value interface{}) *ButtonGroupControl {
    a.Set("source", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *ButtonGroupControl) StaticPlaceholder(value interface{}) *ButtonGroupControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * btnClassName
 */
func (a *ButtonGroupControl) BtnClassName(value interface{}) *ButtonGroupControl {
    a.Set("btnClassName", value)
    return a
}

/**
 * labelAlign
 */
func (a *ButtonGroupControl) LabelAlign(value interface{}) *ButtonGroupControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * labelRemark
 */
func (a *ButtonGroupControl) LabelRemark(value interface{}) *ButtonGroupControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * validateOnChange
 */
func (a *ButtonGroupControl) ValidateOnChange(value interface{}) *ButtonGroupControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * initAutoFill
 */
func (a *ButtonGroupControl) InitAutoFill(value interface{}) *ButtonGroupControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * editApi
 */
func (a *ButtonGroupControl) EditApi(value interface{}) *ButtonGroupControl {
    a.Set("editApi", value)
    return a
}

/**
 * deleteApi
 */
func (a *ButtonGroupControl) DeleteApi(value interface{}) *ButtonGroupControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * staticOn
 */
func (a *ButtonGroupControl) StaticOn(value interface{}) *ButtonGroupControl {
    a.Set("staticOn", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *ButtonGroupControl) StaticLabelClassName(value interface{}) *ButtonGroupControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * btnLevel
 */
func (a *ButtonGroupControl) BtnLevel(value interface{}) *ButtonGroupControl {
    a.Set("btnLevel", value)
    return a
}

/**
 * tiled
 */
func (a *ButtonGroupControl) Tiled(value interface{}) *ButtonGroupControl {
    a.Set("tiled", value)
    return a
}

/**
 * name
 */
func (a *ButtonGroupControl) Name(value interface{}) *ButtonGroupControl {
    a.Set("name", value)
    return a
}

/**
 * inputClassName
 */
func (a *ButtonGroupControl) InputClassName(value interface{}) *ButtonGroupControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * options
 */
func (a *ButtonGroupControl) Options(value interface{}) *ButtonGroupControl {
    a.Set("options", value)
    return a
}

/**
 * visible
 */
func (a *ButtonGroupControl) Visible(value interface{}) *ButtonGroupControl {
    a.Set("visible", value)
    return a
}

/**
 * static
 */
func (a *ButtonGroupControl) Static(value interface{}) *ButtonGroupControl {
    a.Set("static", value)
    return a
}

/**
 * extraName
 */
func (a *ButtonGroupControl) ExtraName(value interface{}) *ButtonGroupControl {
    a.Set("extraName", value)
    return a
}

/**
 * validationErrors
 */
func (a *ButtonGroupControl) ValidationErrors(value interface{}) *ButtonGroupControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * delimiter
 */
func (a *ButtonGroupControl) Delimiter(value interface{}) *ButtonGroupControl {
    a.Set("delimiter", value)
    return a
}

/**
 * className
 */
func (a *ButtonGroupControl) ClassName(value interface{}) *ButtonGroupControl {
    a.Set("className", value)
    return a
}

/**
 * hidden
 */
func (a *ButtonGroupControl) Hidden(value interface{}) *ButtonGroupControl {
    a.Set("hidden", value)
    return a
}

/**
 * btnActiveClassName
 */
func (a *ButtonGroupControl) BtnActiveClassName(value interface{}) *ButtonGroupControl {
    a.Set("btnActiveClassName", value)
    return a
}

/**
 * label
 */
func (a *ButtonGroupControl) Label(value interface{}) *ButtonGroupControl {
    a.Set("label", value)
    return a
}

/**
 * remark
 */
func (a *ButtonGroupControl) Remark(value interface{}) *ButtonGroupControl {
    a.Set("remark", value)
    return a
}

/**
 * submitOnChange
 */
func (a *ButtonGroupControl) SubmitOnChange(value interface{}) *ButtonGroupControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * extractValue
 */
func (a *ButtonGroupControl) ExtractValue(value interface{}) *ButtonGroupControl {
    a.Set("extractValue", value)
    return a
}

/**
 * clearable
 */
func (a *ButtonGroupControl) Clearable(value interface{}) *ButtonGroupControl {
    a.Set("clearable", value)
    return a
}

/**
 * disabledOn
 */
func (a *ButtonGroupControl) DisabledOn(value interface{}) *ButtonGroupControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * size
 */
func (a *ButtonGroupControl) Size(value interface{}) *ButtonGroupControl {
    a.Set("size", value)
    return a
}

/**
 * description
 */
func (a *ButtonGroupControl) Description(value interface{}) *ButtonGroupControl {
    a.Set("description", value)
    return a
}

/**
 * initFetchOn
 */
func (a *ButtonGroupControl) InitFetchOn(value interface{}) *ButtonGroupControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * deferField
 */
func (a *ButtonGroupControl) DeferField(value interface{}) *ButtonGroupControl {
    a.Set("deferField", value)
    return a
}

/**
 * editDialog
 */
func (a *ButtonGroupControl) EditDialog(value interface{}) *ButtonGroupControl {
    a.Set("editDialog", value)
    return a
}

/**
 * useMobileUI
 */
func (a *ButtonGroupControl) UseMobileUI(value interface{}) *ButtonGroupControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * btnActiveLevel
 */
func (a *ButtonGroupControl) BtnActiveLevel(value interface{}) *ButtonGroupControl {
    a.Set("btnActiveLevel", value)
    return a
}

/**
 * labelOverflow
 */
func (a *ButtonGroupControl) LabelOverflow(value interface{}) *ButtonGroupControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * validations
 */
func (a *ButtonGroupControl) Validations(value interface{}) *ButtonGroupControl {
    a.Set("validations", value)
    return a
}

/**
 * checkAll
 */
func (a *ButtonGroupControl) CheckAll(value interface{}) *ButtonGroupControl {
    a.Set("checkAll", value)
    return a
}

/**
 * addDialog
 */
func (a *ButtonGroupControl) AddDialog(value interface{}) *ButtonGroupControl {
    a.Set("addDialog", value)
    return a
}

/**
 * removable
 */
func (a *ButtonGroupControl) Removable(value interface{}) *ButtonGroupControl {
    a.Set("removable", value)
    return a
}

/**
 * visibleOn
 */
func (a *ButtonGroupControl) VisibleOn(value interface{}) *ButtonGroupControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *ButtonGroupControl) StaticClassName(value interface{}) *ButtonGroupControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *ButtonGroupControl) StaticSchema(value interface{}) *ButtonGroupControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * editorSetting
 */
func (a *ButtonGroupControl) EditorSetting(value interface{}) *ButtonGroupControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * placeholder
 */
func (a *ButtonGroupControl) Placeholder(value interface{}) *ButtonGroupControl {
    a.Set("placeholder", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *ButtonGroupControl) ClearValueOnHidden(value interface{}) *ButtonGroupControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * clearValueOnSourceChange
 */
func (a *ButtonGroupControl) ClearValueOnSourceChange(value interface{}) *ButtonGroupControl {
    a.Set("clearValueOnSourceChange", value)
    return a
}

/**
 */
func (a *ButtonGroupControl) Type(value interface{}) *ButtonGroupControl {
    a.Set("type", value)
    return a
}

/**
 * buttons
 */
func (a *ButtonGroupControl) Buttons(value interface{}) *ButtonGroupControl {
    a.Set("buttons", value)
    return a
}

/**
 * initFetch
 */
func (a *ButtonGroupControl) InitFetch(value interface{}) *ButtonGroupControl {
    a.Set("initFetch", value)
    return a
}

/**
 * resetValue
 */
func (a *ButtonGroupControl) ResetValue(value interface{}) *ButtonGroupControl {
    a.Set("resetValue", value)
    return a
}

/**
 * horizontal
 */
func (a *ButtonGroupControl) Horizontal(value interface{}) *ButtonGroupControl {
    a.Set("horizontal", value)
    return a
}

/**
 * value
 */
func (a *ButtonGroupControl) Value(value interface{}) *ButtonGroupControl {
    a.Set("value", value)
    return a
}

/**
 * validateApi
 */
func (a *ButtonGroupControl) ValidateApi(value interface{}) *ButtonGroupControl {
    a.Set("validateApi", value)
    return a
}

/**
 * createBtnLabel
 */
func (a *ButtonGroupControl) CreateBtnLabel(value interface{}) *ButtonGroupControl {
    a.Set("createBtnLabel", value)
    return a
}
