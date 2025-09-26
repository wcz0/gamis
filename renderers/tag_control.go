package renderers


/**
 * Tag 输入框 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/tag

 * @author wcz0
 * @version 6.2.2
 */
type TagControl struct {
	*BaseRenderer
}

func NewTagControl() *TagControl {
    a := &TagControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-tag")
    return a
}


func (a *TagControl) Set(name string, value interface{}) *TagControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * delimiter
 */
func (a *TagControl) Delimiter(value interface{}) *TagControl {
    a.Set("delimiter", value)
    return a
}

/**
 * staticOn
 */
func (a *TagControl) StaticOn(value interface{}) *TagControl {
    a.Set("staticOn", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *TagControl) StaticLabelClassName(value interface{}) *TagControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * optionsTip
 */
func (a *TagControl) OptionsTip(value interface{}) *TagControl {
    a.Set("optionsTip", value)
    return a
}

/**
 * source
 */
func (a *TagControl) Source(value interface{}) *TagControl {
    a.Set("source", value)
    return a
}

/**
 * multiple
 */
func (a *TagControl) Multiple(value interface{}) *TagControl {
    a.Set("multiple", value)
    return a
}

/**
 * desc
 */
func (a *TagControl) Desc(value interface{}) *TagControl {
    a.Set("desc", value)
    return a
}

/**
 * inline
 */
func (a *TagControl) Inline(value interface{}) *TagControl {
    a.Set("inline", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *TagControl) StaticPlaceholder(value interface{}) *TagControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticClassName
 */
func (a *TagControl) StaticClassName(value interface{}) *TagControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * editApi
 */
func (a *TagControl) EditApi(value interface{}) *TagControl {
    a.Set("editApi", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *TagControl) ReadOnlyOn(value interface{}) *TagControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * description
 */
func (a *TagControl) Description(value interface{}) *TagControl {
    a.Set("description", value)
    return a
}

/**
 * useMobileUI
 */
func (a *TagControl) UseMobileUI(value interface{}) *TagControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * overflowTagPopover
 */
func (a *TagControl) OverflowTagPopover(value interface{}) *TagControl {
    a.Set("overflowTagPopover", value)
    return a
}

/**
 * createBtnLabel
 */
func (a *TagControl) CreateBtnLabel(value interface{}) *TagControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * editControls
 */
func (a *TagControl) EditControls(value interface{}) *TagControl {
    a.Set("editControls", value)
    return a
}

/**
 * placeholder
 */
func (a *TagControl) Placeholder(value interface{}) *TagControl {
    a.Set("placeholder", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *TagControl) ClearValueOnHidden(value interface{}) *TagControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * style
 */
func (a *TagControl) Style(value interface{}) *TagControl {
    a.Set("style", value)
    return a
}

/**
 * size
 */
func (a *TagControl) Size(value interface{}) *TagControl {
    a.Set("size", value)
    return a
}

/**
 * deferApi
 */
func (a *TagControl) DeferApi(value interface{}) *TagControl {
    a.Set("deferApi", value)
    return a
}

/**
 * creatable
 */
func (a *TagControl) Creatable(value interface{}) *TagControl {
    a.Set("creatable", value)
    return a
}

/**
 * validateOnChange
 */
func (a *TagControl) ValidateOnChange(value interface{}) *TagControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * enableBatchAdd
 */
func (a *TagControl) EnableBatchAdd(value interface{}) *TagControl {
    a.Set("enableBatchAdd", value)
    return a
}

/**
 * labelOverflow
 */
func (a *TagControl) LabelOverflow(value interface{}) *TagControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * name
 */
func (a *TagControl) Name(value interface{}) *TagControl {
    a.Set("name", value)
    return a
}

/**
 * extraName
 */
func (a *TagControl) ExtraName(value interface{}) *TagControl {
    a.Set("extraName", value)
    return a
}

/**
 * disabledOn
 */
func (a *TagControl) DisabledOn(value interface{}) *TagControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * visible
 */
func (a *TagControl) Visible(value interface{}) *TagControl {
    a.Set("visible", value)
    return a
}

/**
 * valuesNoWrap
 */
func (a *TagControl) ValuesNoWrap(value interface{}) *TagControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * remark
 */
func (a *TagControl) Remark(value interface{}) *TagControl {
    a.Set("remark", value)
    return a
}

/**
 * disabled
 */
func (a *TagControl) Disabled(value interface{}) *TagControl {
    a.Set("disabled", value)
    return a
}

/**
 * hiddenOn
 */
func (a *TagControl) HiddenOn(value interface{}) *TagControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *TagControl) StaticInputClassName(value interface{}) *TagControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * separator
 */
func (a *TagControl) Separator(value interface{}) *TagControl {
    a.Set("separator", value)
    return a
}

/**
 * clearable
 */
func (a *TagControl) Clearable(value interface{}) *TagControl {
    a.Set("clearable", value)
    return a
}

/**
 * deferField
 */
func (a *TagControl) DeferField(value interface{}) *TagControl {
    a.Set("deferField", value)
    return a
}

/**
 * labelWidth
 */
func (a *TagControl) LabelWidth(value interface{}) *TagControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * readOnly
 */
func (a *TagControl) ReadOnly(value interface{}) *TagControl {
    a.Set("readOnly", value)
    return a
}

/**
 * validationErrors
 */
func (a *TagControl) ValidationErrors(value interface{}) *TagControl {
    a.Set("validationErrors", value)
    return a
}

/**
 */
func (a *TagControl) Type(value interface{}) *TagControl {
    a.Set("type", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *TagControl) Width(value interface{}) *TagControl {
    a.Set("width", value)
    return a
}

/**
 * initFetchOn
 */
func (a *TagControl) InitFetchOn(value interface{}) *TagControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * initFetch
 */
func (a *TagControl) InitFetch(value interface{}) *TagControl {
    a.Set("initFetch", value)
    return a
}

/**
 * deleteApi
 */
func (a *TagControl) DeleteApi(value interface{}) *TagControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * labelRemark
 */
func (a *TagControl) LabelRemark(value interface{}) *TagControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * horizontal
 */
func (a *TagControl) Horizontal(value interface{}) *TagControl {
    a.Set("horizontal", value)
    return a
}

/**
 * validateApi
 */
func (a *TagControl) ValidateApi(value interface{}) *TagControl {
    a.Set("validateApi", value)
    return a
}

/**
 * autoFill
 */
func (a *TagControl) AutoFill(value interface{}) *TagControl {
    a.Set("autoFill", value)
    return a
}

/**
 * resetValue
 */
func (a *TagControl) ResetValue(value interface{}) *TagControl {
    a.Set("resetValue", value)
    return a
}

/**
 * addDialog
 */
func (a *TagControl) AddDialog(value interface{}) *TagControl {
    a.Set("addDialog", value)
    return a
}

/**
 * editable
 */
func (a *TagControl) Editable(value interface{}) *TagControl {
    a.Set("editable", value)
    return a
}

/**
 * removable
 */
func (a *TagControl) Removable(value interface{}) *TagControl {
    a.Set("removable", value)
    return a
}

/**
 * validations
 */
func (a *TagControl) Validations(value interface{}) *TagControl {
    a.Set("validations", value)
    return a
}

/**
 * max
 */
func (a *TagControl) Max(value interface{}) *TagControl {
    a.Set("max", value)
    return a
}

/**
 * addControls
 */
func (a *TagControl) AddControls(value interface{}) *TagControl {
    a.Set("addControls", value)
    return a
}

/**
 * editDialog
 */
func (a *TagControl) EditDialog(value interface{}) *TagControl {
    a.Set("editDialog", value)
    return a
}

/**
 * labelClassName
 */
func (a *TagControl) LabelClassName(value interface{}) *TagControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *TagControl) StaticSchema(value interface{}) *TagControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * editorSetting
 */
func (a *TagControl) EditorSetting(value interface{}) *TagControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * selectFirst
 */
func (a *TagControl) SelectFirst(value interface{}) *TagControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * joinValues
 */
func (a *TagControl) JoinValues(value interface{}) *TagControl {
    a.Set("joinValues", value)
    return a
}

/**
 * clearValueOnSourceChange
 */
func (a *TagControl) ClearValueOnSourceChange(value interface{}) *TagControl {
    a.Set("clearValueOnSourceChange", value)
    return a
}

/**
 * label
 */
func (a *TagControl) Label(value interface{}) *TagControl {
    a.Set("label", value)
    return a
}

/**
 * submitOnChange
 */
func (a *TagControl) SubmitOnChange(value interface{}) *TagControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * initAutoFill
 */
func (a *TagControl) InitAutoFill(value interface{}) *TagControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * maxTagLength
 */
func (a *TagControl) MaxTagLength(value interface{}) *TagControl {
    a.Set("maxTagLength", value)
    return a
}

/**
 * maxTagCount
 */
func (a *TagControl) MaxTagCount(value interface{}) *TagControl {
    a.Set("maxTagCount", value)
    return a
}

/**
 * hint
 */
func (a *TagControl) Hint(value interface{}) *TagControl {
    a.Set("hint", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *TagControl) DescriptionClassName(value interface{}) *TagControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * addApi
 */
func (a *TagControl) AddApi(value interface{}) *TagControl {
    a.Set("addApi", value)
    return a
}

/**
 * deleteConfirmText
 */
func (a *TagControl) DeleteConfirmText(value interface{}) *TagControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * mode
 */
func (a *TagControl) Mode(value interface{}) *TagControl {
    a.Set("mode", value)
    return a
}

/**
 * value
 */
func (a *TagControl) Value(value interface{}) *TagControl {
    a.Set("value", value)
    return a
}

/**
 * hidden
 */
func (a *TagControl) Hidden(value interface{}) *TagControl {
    a.Set("hidden", value)
    return a
}

/**
 * visibleOn
 */
func (a *TagControl) VisibleOn(value interface{}) *TagControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * id
 */
func (a *TagControl) Id(value interface{}) *TagControl {
    a.Set("id", value)
    return a
}

/**
 * onEvent
 */
func (a *TagControl) OnEvent(value interface{}) *TagControl {
    a.Set("onEvent", value)
    return a
}

/**
 * checkAll
 */
func (a *TagControl) CheckAll(value interface{}) *TagControl {
    a.Set("checkAll", value)
    return a
}

/**
 * extractValue
 */
func (a *TagControl) ExtractValue(value interface{}) *TagControl {
    a.Set("extractValue", value)
    return a
}

/**
 * labelAlign
 */
func (a *TagControl) LabelAlign(value interface{}) *TagControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * inputClassName
 */
func (a *TagControl) InputClassName(value interface{}) *TagControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * required
 */
func (a *TagControl) Required(value interface{}) *TagControl {
    a.Set("required", value)
    return a
}

/**
 * static
 */
func (a *TagControl) Static(value interface{}) *TagControl {
    a.Set("static", value)
    return a
}

/**
 * dropdown
 */
func (a *TagControl) Dropdown(value interface{}) *TagControl {
    a.Set("dropdown", value)
    return a
}

/**
 * options
 */
func (a *TagControl) Options(value interface{}) *TagControl {
    a.Set("options", value)
    return a
}

/**
 * row
 */
func (a *TagControl) Row(value interface{}) *TagControl {
    a.Set("row", value)
    return a
}

/**
 * className
 */
func (a *TagControl) ClassName(value interface{}) *TagControl {
    a.Set("className", value)
    return a
}
