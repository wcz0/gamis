package renderers


/**
 * 链式下拉框 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/chain-select

 * @author wcz0
 * @version 6.2.2
 */
type ChainedSelectControl struct {
	*BaseRenderer
}

func NewChainedSelectControl() *ChainedSelectControl {
    a := &ChainedSelectControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "chained-select")
    return a
}


func (a *ChainedSelectControl) Set(name string, value interface{}) *ChainedSelectControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * resetValue
 */
func (a *ChainedSelectControl) ResetValue(value interface{}) *ChainedSelectControl {
    a.Set("resetValue", value)
    return a
}

/**
 * description
 */
func (a *ChainedSelectControl) Description(value interface{}) *ChainedSelectControl {
    a.Set("description", value)
    return a
}

/**
 * size
 */
func (a *ChainedSelectControl) Size(value interface{}) *ChainedSelectControl {
    a.Set("size", value)
    return a
}

/**
 * selectFirst
 */
func (a *ChainedSelectControl) SelectFirst(value interface{}) *ChainedSelectControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * deferField
 */
func (a *ChainedSelectControl) DeferField(value interface{}) *ChainedSelectControl {
    a.Set("deferField", value)
    return a
}

/**
 * addDialog
 */
func (a *ChainedSelectControl) AddDialog(value interface{}) *ChainedSelectControl {
    a.Set("addDialog", value)
    return a
}

/**
 * editable
 */
func (a *ChainedSelectControl) Editable(value interface{}) *ChainedSelectControl {
    a.Set("editable", value)
    return a
}

/**
 * editControls
 */
func (a *ChainedSelectControl) EditControls(value interface{}) *ChainedSelectControl {
    a.Set("editControls", value)
    return a
}

/**
 * deleteApi
 */
func (a *ChainedSelectControl) DeleteApi(value interface{}) *ChainedSelectControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * validateApi
 */
func (a *ChainedSelectControl) ValidateApi(value interface{}) *ChainedSelectControl {
    a.Set("validateApi", value)
    return a
}

/**
 * deferApi
 */
func (a *ChainedSelectControl) DeferApi(value interface{}) *ChainedSelectControl {
    a.Set("deferApi", value)
    return a
}

/**
 * clearValueOnSourceChange
 */
func (a *ChainedSelectControl) ClearValueOnSourceChange(value interface{}) *ChainedSelectControl {
    a.Set("clearValueOnSourceChange", value)
    return a
}

/**
 * label
 */
func (a *ChainedSelectControl) Label(value interface{}) *ChainedSelectControl {
    a.Set("label", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *ChainedSelectControl) StaticLabelClassName(value interface{}) *ChainedSelectControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * style
 */
func (a *ChainedSelectControl) Style(value interface{}) *ChainedSelectControl {
    a.Set("style", value)
    return a
}

/**
 * editorSetting
 */
func (a *ChainedSelectControl) EditorSetting(value interface{}) *ChainedSelectControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * removable
 */
func (a *ChainedSelectControl) Removable(value interface{}) *ChainedSelectControl {
    a.Set("removable", value)
    return a
}

/**
 * labelAlign
 */
func (a *ChainedSelectControl) LabelAlign(value interface{}) *ChainedSelectControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * inputClassName
 */
func (a *ChainedSelectControl) InputClassName(value interface{}) *ChainedSelectControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * row
 */
func (a *ChainedSelectControl) Row(value interface{}) *ChainedSelectControl {
    a.Set("row", value)
    return a
}

/**
 * joinValues
 */
func (a *ChainedSelectControl) JoinValues(value interface{}) *ChainedSelectControl {
    a.Set("joinValues", value)
    return a
}

/**
 * deleteConfirmText
 */
func (a *ChainedSelectControl) DeleteConfirmText(value interface{}) *ChainedSelectControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * name
 */
func (a *ChainedSelectControl) Name(value interface{}) *ChainedSelectControl {
    a.Set("name", value)
    return a
}

/**
 * placeholder
 */
func (a *ChainedSelectControl) Placeholder(value interface{}) *ChainedSelectControl {
    a.Set("placeholder", value)
    return a
}

/**
 * visibleOn
 */
func (a *ChainedSelectControl) VisibleOn(value interface{}) *ChainedSelectControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * onEvent
 */
func (a *ChainedSelectControl) OnEvent(value interface{}) *ChainedSelectControl {
    a.Set("onEvent", value)
    return a
}

/**
 * staticSchema
 */
func (a *ChainedSelectControl) StaticSchema(value interface{}) *ChainedSelectControl {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *ChainedSelectControl) Type(value interface{}) *ChainedSelectControl {
    a.Set("type", value)
    return a
}

/**
 * delimiter
 */
func (a *ChainedSelectControl) Delimiter(value interface{}) *ChainedSelectControl {
    a.Set("delimiter", value)
    return a
}

/**
 * createBtnLabel
 */
func (a *ChainedSelectControl) CreateBtnLabel(value interface{}) *ChainedSelectControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * inline
 */
func (a *ChainedSelectControl) Inline(value interface{}) *ChainedSelectControl {
    a.Set("inline", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *ChainedSelectControl) StaticPlaceholder(value interface{}) *ChainedSelectControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticClassName
 */
func (a *ChainedSelectControl) StaticClassName(value interface{}) *ChainedSelectControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *ChainedSelectControl) ReadOnlyOn(value interface{}) *ChainedSelectControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * mode
 */
func (a *ChainedSelectControl) Mode(value interface{}) *ChainedSelectControl {
    a.Set("mode", value)
    return a
}

/**
 * checkAll
 */
func (a *ChainedSelectControl) CheckAll(value interface{}) *ChainedSelectControl {
    a.Set("checkAll", value)
    return a
}

/**
 * labelOverflow
 */
func (a *ChainedSelectControl) LabelOverflow(value interface{}) *ChainedSelectControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * labelClassName
 */
func (a *ChainedSelectControl) LabelClassName(value interface{}) *ChainedSelectControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * validateOnChange
 */
func (a *ChainedSelectControl) ValidateOnChange(value interface{}) *ChainedSelectControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * required
 */
func (a *ChainedSelectControl) Required(value interface{}) *ChainedSelectControl {
    a.Set("required", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *ChainedSelectControl) ClearValueOnHidden(value interface{}) *ChainedSelectControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * source
 */
func (a *ChainedSelectControl) Source(value interface{}) *ChainedSelectControl {
    a.Set("source", value)
    return a
}

/**
 * labelRemark
 */
func (a *ChainedSelectControl) LabelRemark(value interface{}) *ChainedSelectControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * staticOn
 */
func (a *ChainedSelectControl) StaticOn(value interface{}) *ChainedSelectControl {
    a.Set("staticOn", value)
    return a
}

/**
 * options
 */
func (a *ChainedSelectControl) Options(value interface{}) *ChainedSelectControl {
    a.Set("options", value)
    return a
}

/**
 * valuesNoWrap
 */
func (a *ChainedSelectControl) ValuesNoWrap(value interface{}) *ChainedSelectControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * extractValue
 */
func (a *ChainedSelectControl) ExtractValue(value interface{}) *ChainedSelectControl {
    a.Set("extractValue", value)
    return a
}

/**
 * extraName
 */
func (a *ChainedSelectControl) ExtraName(value interface{}) *ChainedSelectControl {
    a.Set("extraName", value)
    return a
}

/**
 * disabledOn
 */
func (a *ChainedSelectControl) DisabledOn(value interface{}) *ChainedSelectControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * hiddenOn
 */
func (a *ChainedSelectControl) HiddenOn(value interface{}) *ChainedSelectControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * addApi
 */
func (a *ChainedSelectControl) AddApi(value interface{}) *ChainedSelectControl {
    a.Set("addApi", value)
    return a
}

/**
 * static
 */
func (a *ChainedSelectControl) Static(value interface{}) *ChainedSelectControl {
    a.Set("static", value)
    return a
}

/**
 * initFetchOn
 */
func (a *ChainedSelectControl) InitFetchOn(value interface{}) *ChainedSelectControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * creatable
 */
func (a *ChainedSelectControl) Creatable(value interface{}) *ChainedSelectControl {
    a.Set("creatable", value)
    return a
}

/**
 * editApi
 */
func (a *ChainedSelectControl) EditApi(value interface{}) *ChainedSelectControl {
    a.Set("editApi", value)
    return a
}

/**
 * hidden
 */
func (a *ChainedSelectControl) Hidden(value interface{}) *ChainedSelectControl {
    a.Set("hidden", value)
    return a
}

/**
 * useMobileUI
 */
func (a *ChainedSelectControl) UseMobileUI(value interface{}) *ChainedSelectControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *ChainedSelectControl) Width(value interface{}) *ChainedSelectControl {
    a.Set("width", value)
    return a
}

/**
 * multiple
 */
func (a *ChainedSelectControl) Multiple(value interface{}) *ChainedSelectControl {
    a.Set("multiple", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *ChainedSelectControl) DescriptionClassName(value interface{}) *ChainedSelectControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * addControls
 */
func (a *ChainedSelectControl) AddControls(value interface{}) *ChainedSelectControl {
    a.Set("addControls", value)
    return a
}

/**
 * hint
 */
func (a *ChainedSelectControl) Hint(value interface{}) *ChainedSelectControl {
    a.Set("hint", value)
    return a
}

/**
 * submitOnChange
 */
func (a *ChainedSelectControl) SubmitOnChange(value interface{}) *ChainedSelectControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * readOnly
 */
func (a *ChainedSelectControl) ReadOnly(value interface{}) *ChainedSelectControl {
    a.Set("readOnly", value)
    return a
}

/**
 * validations
 */
func (a *ChainedSelectControl) Validations(value interface{}) *ChainedSelectControl {
    a.Set("validations", value)
    return a
}

/**
 * value
 */
func (a *ChainedSelectControl) Value(value interface{}) *ChainedSelectControl {
    a.Set("value", value)
    return a
}

/**
 * initFetch
 */
func (a *ChainedSelectControl) InitFetch(value interface{}) *ChainedSelectControl {
    a.Set("initFetch", value)
    return a
}

/**
 * labelWidth
 */
func (a *ChainedSelectControl) LabelWidth(value interface{}) *ChainedSelectControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * remark
 */
func (a *ChainedSelectControl) Remark(value interface{}) *ChainedSelectControl {
    a.Set("remark", value)
    return a
}

/**
 * desc
 */
func (a *ChainedSelectControl) Desc(value interface{}) *ChainedSelectControl {
    a.Set("desc", value)
    return a
}

/**
 * horizontal
 */
func (a *ChainedSelectControl) Horizontal(value interface{}) *ChainedSelectControl {
    a.Set("horizontal", value)
    return a
}

/**
 * validationErrors
 */
func (a *ChainedSelectControl) ValidationErrors(value interface{}) *ChainedSelectControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * editDialog
 */
func (a *ChainedSelectControl) EditDialog(value interface{}) *ChainedSelectControl {
    a.Set("editDialog", value)
    return a
}

/**
 * autoFill
 */
func (a *ChainedSelectControl) AutoFill(value interface{}) *ChainedSelectControl {
    a.Set("autoFill", value)
    return a
}

/**
 * initAutoFill
 */
func (a *ChainedSelectControl) InitAutoFill(value interface{}) *ChainedSelectControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * className
 */
func (a *ChainedSelectControl) ClassName(value interface{}) *ChainedSelectControl {
    a.Set("className", value)
    return a
}

/**
 * disabled
 */
func (a *ChainedSelectControl) Disabled(value interface{}) *ChainedSelectControl {
    a.Set("disabled", value)
    return a
}

/**
 * visible
 */
func (a *ChainedSelectControl) Visible(value interface{}) *ChainedSelectControl {
    a.Set("visible", value)
    return a
}

/**
 * id
 */
func (a *ChainedSelectControl) Id(value interface{}) *ChainedSelectControl {
    a.Set("id", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *ChainedSelectControl) StaticInputClassName(value interface{}) *ChainedSelectControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * clearable
 */
func (a *ChainedSelectControl) Clearable(value interface{}) *ChainedSelectControl {
    a.Set("clearable", value)
    return a
}
