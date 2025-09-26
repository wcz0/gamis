package renderers


/**
 * TransferPicker 穿梭器的弹框形态 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/transfer-picker

 * @author wcz0
 * @version 6.2.2
 */
type TransferPickerControl struct {
	*BaseRenderer
}

func NewTransferPickerControl() *TransferPickerControl {
    a := &TransferPickerControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "transfer-picker")
    return a
}


func (a *TransferPickerControl) Set(name string, value interface{}) *TransferPickerControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * editable
 */
func (a *TransferPickerControl) Editable(value interface{}) *TransferPickerControl {
    a.Set("editable", value)
    return a
}

/**
 * hiddenOn
 */
func (a *TransferPickerControl) HiddenOn(value interface{}) *TransferPickerControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * borderMode
 */
func (a *TransferPickerControl) BorderMode(value interface{}) *TransferPickerControl {
    a.Set("borderMode", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *TransferPickerControl) ReadOnlyOn(value interface{}) *TransferPickerControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * value
 */
func (a *TransferPickerControl) Value(value interface{}) *TransferPickerControl {
    a.Set("value", value)
    return a
}

/**
 * visibleOn
 */
func (a *TransferPickerControl) VisibleOn(value interface{}) *TransferPickerControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * leftOptions
 */
func (a *TransferPickerControl) LeftOptions(value interface{}) *TransferPickerControl {
    a.Set("leftOptions", value)
    return a
}

/**
 * statistics
 */
func (a *TransferPickerControl) Statistics(value interface{}) *TransferPickerControl {
    a.Set("statistics", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *TransferPickerControl) Width(value interface{}) *TransferPickerControl {
    a.Set("width", value)
    return a
}

/**
 * selectFirst
 */
func (a *TransferPickerControl) SelectFirst(value interface{}) *TransferPickerControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * useMobileUI
 */
func (a *TransferPickerControl) UseMobileUI(value interface{}) *TransferPickerControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * resultSearchPlaceholder
 */
func (a *TransferPickerControl) ResultSearchPlaceholder(value interface{}) *TransferPickerControl {
    a.Set("resultSearchPlaceholder", value)
    return a
}

/**
 * onlyChildren
 */
func (a *TransferPickerControl) OnlyChildren(value interface{}) *TransferPickerControl {
    a.Set("onlyChildren", value)
    return a
}

/**
 * pagination
 */
func (a *TransferPickerControl) Pagination(value interface{}) *TransferPickerControl {
    a.Set("pagination", value)
    return a
}

/**
 * addControls
 */
func (a *TransferPickerControl) AddControls(value interface{}) *TransferPickerControl {
    a.Set("addControls", value)
    return a
}

/**
 * inputClassName
 */
func (a *TransferPickerControl) InputClassName(value interface{}) *TransferPickerControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * extractValue
 */
func (a *TransferPickerControl) ExtractValue(value interface{}) *TransferPickerControl {
    a.Set("extractValue", value)
    return a
}

/**
 * hint
 */
func (a *TransferPickerControl) Hint(value interface{}) *TransferPickerControl {
    a.Set("hint", value)
    return a
}

/**
 * searchResultMode
 */
func (a *TransferPickerControl) SearchResultMode(value interface{}) *TransferPickerControl {
    a.Set("searchResultMode", value)
    return a
}

/**
 * label
 */
func (a *TransferPickerControl) Label(value interface{}) *TransferPickerControl {
    a.Set("label", value)
    return a
}

/**
 * horizontal
 */
func (a *TransferPickerControl) Horizontal(value interface{}) *TransferPickerControl {
    a.Set("horizontal", value)
    return a
}

/**
 * delimiter
 */
func (a *TransferPickerControl) Delimiter(value interface{}) *TransferPickerControl {
    a.Set("delimiter", value)
    return a
}

/**
 * mode
 */
func (a *TransferPickerControl) Mode(value interface{}) *TransferPickerControl {
    a.Set("mode", value)
    return a
}

/**
 * inline
 */
func (a *TransferPickerControl) Inline(value interface{}) *TransferPickerControl {
    a.Set("inline", value)
    return a
}

/**
 * id
 */
func (a *TransferPickerControl) Id(value interface{}) *TransferPickerControl {
    a.Set("id", value)
    return a
}

/**
 * extraName
 */
func (a *TransferPickerControl) ExtraName(value interface{}) *TransferPickerControl {
    a.Set("extraName", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *TransferPickerControl) StaticPlaceholder(value interface{}) *TransferPickerControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *TransferPickerControl) StaticLabelClassName(value interface{}) *TransferPickerControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * menuTpl
 */
func (a *TransferPickerControl) MenuTpl(value interface{}) *TransferPickerControl {
    a.Set("menuTpl", value)
    return a
}

/**
 * creatable
 */
func (a *TransferPickerControl) Creatable(value interface{}) *TransferPickerControl {
    a.Set("creatable", value)
    return a
}

/**
 * className
 */
func (a *TransferPickerControl) ClassName(value interface{}) *TransferPickerControl {
    a.Set("className", value)
    return a
}

/**
 * resultSearchable
 */
func (a *TransferPickerControl) ResultSearchable(value interface{}) *TransferPickerControl {
    a.Set("resultSearchable", value)
    return a
}

/**
 * deleteApi
 */
func (a *TransferPickerControl) DeleteApi(value interface{}) *TransferPickerControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * checkAll
 */
func (a *TransferPickerControl) CheckAll(value interface{}) *TransferPickerControl {
    a.Set("checkAll", value)
    return a
}

/**
 * validationErrors
 */
func (a *TransferPickerControl) ValidationErrors(value interface{}) *TransferPickerControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * staticSchema
 */
func (a *TransferPickerControl) StaticSchema(value interface{}) *TransferPickerControl {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *TransferPickerControl) Type(value interface{}) *TransferPickerControl {
    a.Set("type", value)
    return a
}

/**
 * itemHeight
 */
func (a *TransferPickerControl) ItemHeight(value interface{}) *TransferPickerControl {
    a.Set("itemHeight", value)
    return a
}

/**
 * required
 */
func (a *TransferPickerControl) Required(value interface{}) *TransferPickerControl {
    a.Set("required", value)
    return a
}

/**
 * editorSetting
 */
func (a *TransferPickerControl) EditorSetting(value interface{}) *TransferPickerControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * leftMode
 */
func (a *TransferPickerControl) LeftMode(value interface{}) *TransferPickerControl {
    a.Set("leftMode", value)
    return a
}

/**
 * searchable
 */
func (a *TransferPickerControl) Searchable(value interface{}) *TransferPickerControl {
    a.Set("searchable", value)
    return a
}

/**
 * resultTitle
 */
func (a *TransferPickerControl) ResultTitle(value interface{}) *TransferPickerControl {
    a.Set("resultTitle", value)
    return a
}

/**
 * valuesNoWrap
 */
func (a *TransferPickerControl) ValuesNoWrap(value interface{}) *TransferPickerControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * addDialog
 */
func (a *TransferPickerControl) AddDialog(value interface{}) *TransferPickerControl {
    a.Set("addDialog", value)
    return a
}

/**
 * validateOnChange
 */
func (a *TransferPickerControl) ValidateOnChange(value interface{}) *TransferPickerControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * hidden
 */
func (a *TransferPickerControl) Hidden(value interface{}) *TransferPickerControl {
    a.Set("hidden", value)
    return a
}

/**
 * pickerSize
 */
func (a *TransferPickerControl) PickerSize(value interface{}) *TransferPickerControl {
    a.Set("pickerSize", value)
    return a
}

/**
 * clearable
 */
func (a *TransferPickerControl) Clearable(value interface{}) *TransferPickerControl {
    a.Set("clearable", value)
    return a
}

/**
 * loadingConfig
 */
func (a *TransferPickerControl) LoadingConfig(value interface{}) *TransferPickerControl {
    a.Set("loadingConfig", value)
    return a
}

/**
 * initFetch
 */
func (a *TransferPickerControl) InitFetch(value interface{}) *TransferPickerControl {
    a.Set("initFetch", value)
    return a
}

/**
 * sortable
 */
func (a *TransferPickerControl) Sortable(value interface{}) *TransferPickerControl {
    a.Set("sortable", value)
    return a
}

/**
 * searchApi
 */
func (a *TransferPickerControl) SearchApi(value interface{}) *TransferPickerControl {
    a.Set("searchApi", value)
    return a
}

/**
 * description
 */
func (a *TransferPickerControl) Description(value interface{}) *TransferPickerControl {
    a.Set("description", value)
    return a
}

/**
 * row
 */
func (a *TransferPickerControl) Row(value interface{}) *TransferPickerControl {
    a.Set("row", value)
    return a
}

/**
 * initiallyOpen
 */
func (a *TransferPickerControl) InitiallyOpen(value interface{}) *TransferPickerControl {
    a.Set("initiallyOpen", value)
    return a
}

/**
 * autoCheckChildren
 */
func (a *TransferPickerControl) AutoCheckChildren(value interface{}) *TransferPickerControl {
    a.Set("autoCheckChildren", value)
    return a
}

/**
 * searchResultColumns
 */
func (a *TransferPickerControl) SearchResultColumns(value interface{}) *TransferPickerControl {
    a.Set("searchResultColumns", value)
    return a
}

/**
 * addApi
 */
func (a *TransferPickerControl) AddApi(value interface{}) *TransferPickerControl {
    a.Set("addApi", value)
    return a
}

/**
 * desc
 */
func (a *TransferPickerControl) Desc(value interface{}) *TransferPickerControl {
    a.Set("desc", value)
    return a
}

/**
 * selectTitle
 */
func (a *TransferPickerControl) SelectTitle(value interface{}) *TransferPickerControl {
    a.Set("selectTitle", value)
    return a
}

/**
 * options
 */
func (a *TransferPickerControl) Options(value interface{}) *TransferPickerControl {
    a.Set("options", value)
    return a
}

/**
 * joinValues
 */
func (a *TransferPickerControl) JoinValues(value interface{}) *TransferPickerControl {
    a.Set("joinValues", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *TransferPickerControl) ClearValueOnHidden(value interface{}) *TransferPickerControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * autoFill
 */
func (a *TransferPickerControl) AutoFill(value interface{}) *TransferPickerControl {
    a.Set("autoFill", value)
    return a
}

/**
 * labelAlign
 */
func (a *TransferPickerControl) LabelAlign(value interface{}) *TransferPickerControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * remark
 */
func (a *TransferPickerControl) Remark(value interface{}) *TransferPickerControl {
    a.Set("remark", value)
    return a
}

/**
 * validateApi
 */
func (a *TransferPickerControl) ValidateApi(value interface{}) *TransferPickerControl {
    a.Set("validateApi", value)
    return a
}

/**
 * staticOn
 */
func (a *TransferPickerControl) StaticOn(value interface{}) *TransferPickerControl {
    a.Set("staticOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *TransferPickerControl) StaticInputClassName(value interface{}) *TransferPickerControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * rightMode
 */
func (a *TransferPickerControl) RightMode(value interface{}) *TransferPickerControl {
    a.Set("rightMode", value)
    return a
}

/**
 * deleteConfirmText
 */
func (a *TransferPickerControl) DeleteConfirmText(value interface{}) *TransferPickerControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * labelClassName
 */
func (a *TransferPickerControl) LabelClassName(value interface{}) *TransferPickerControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * initAutoFill
 */
func (a *TransferPickerControl) InitAutoFill(value interface{}) *TransferPickerControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * disabledOn
 */
func (a *TransferPickerControl) DisabledOn(value interface{}) *TransferPickerControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *TransferPickerControl) StaticClassName(value interface{}) *TransferPickerControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * virtualThreshold
 */
func (a *TransferPickerControl) VirtualThreshold(value interface{}) *TransferPickerControl {
    a.Set("virtualThreshold", value)
    return a
}

/**
 * createBtnLabel
 */
func (a *TransferPickerControl) CreateBtnLabel(value interface{}) *TransferPickerControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * readOnly
 */
func (a *TransferPickerControl) ReadOnly(value interface{}) *TransferPickerControl {
    a.Set("readOnly", value)
    return a
}

/**
 * selectMode
 */
func (a *TransferPickerControl) SelectMode(value interface{}) *TransferPickerControl {
    a.Set("selectMode", value)
    return a
}

/**
 * columns
 */
func (a *TransferPickerControl) Columns(value interface{}) *TransferPickerControl {
    a.Set("columns", value)
    return a
}

/**
 * editApi
 */
func (a *TransferPickerControl) EditApi(value interface{}) *TransferPickerControl {
    a.Set("editApi", value)
    return a
}

/**
 * name
 */
func (a *TransferPickerControl) Name(value interface{}) *TransferPickerControl {
    a.Set("name", value)
    return a
}

/**
 * validations
 */
func (a *TransferPickerControl) Validations(value interface{}) *TransferPickerControl {
    a.Set("validations", value)
    return a
}

/**
 * editControls
 */
func (a *TransferPickerControl) EditControls(value interface{}) *TransferPickerControl {
    a.Set("editControls", value)
    return a
}

/**
 * labelOverflow
 */
func (a *TransferPickerControl) LabelOverflow(value interface{}) *TransferPickerControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * showArrow
 */
func (a *TransferPickerControl) ShowArrow(value interface{}) *TransferPickerControl {
    a.Set("showArrow", value)
    return a
}

/**
 * multiple
 */
func (a *TransferPickerControl) Multiple(value interface{}) *TransferPickerControl {
    a.Set("multiple", value)
    return a
}

/**
 * resetValue
 */
func (a *TransferPickerControl) ResetValue(value interface{}) *TransferPickerControl {
    a.Set("resetValue", value)
    return a
}

/**
 * labelRemark
 */
func (a *TransferPickerControl) LabelRemark(value interface{}) *TransferPickerControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * submitOnChange
 */
func (a *TransferPickerControl) SubmitOnChange(value interface{}) *TransferPickerControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * static
 */
func (a *TransferPickerControl) Static(value interface{}) *TransferPickerControl {
    a.Set("static", value)
    return a
}

/**
 * style
 */
func (a *TransferPickerControl) Style(value interface{}) *TransferPickerControl {
    a.Set("style", value)
    return a
}

/**
 * resultListModeFollowSelect
 */
func (a *TransferPickerControl) ResultListModeFollowSelect(value interface{}) *TransferPickerControl {
    a.Set("resultListModeFollowSelect", value)
    return a
}

/**
 * source
 */
func (a *TransferPickerControl) Source(value interface{}) *TransferPickerControl {
    a.Set("source", value)
    return a
}

/**
 * clearValueOnSourceChange
 */
func (a *TransferPickerControl) ClearValueOnSourceChange(value interface{}) *TransferPickerControl {
    a.Set("clearValueOnSourceChange", value)
    return a
}

/**
 * labelWidth
 */
func (a *TransferPickerControl) LabelWidth(value interface{}) *TransferPickerControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * editDialog
 */
func (a *TransferPickerControl) EditDialog(value interface{}) *TransferPickerControl {
    a.Set("editDialog", value)
    return a
}

/**
 * onEvent
 */
func (a *TransferPickerControl) OnEvent(value interface{}) *TransferPickerControl {
    a.Set("onEvent", value)
    return a
}

/**
 * valueTpl
 */
func (a *TransferPickerControl) ValueTpl(value interface{}) *TransferPickerControl {
    a.Set("valueTpl", value)
    return a
}

/**
 * initFetchOn
 */
func (a *TransferPickerControl) InitFetchOn(value interface{}) *TransferPickerControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * removable
 */
func (a *TransferPickerControl) Removable(value interface{}) *TransferPickerControl {
    a.Set("removable", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *TransferPickerControl) DescriptionClassName(value interface{}) *TransferPickerControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * placeholder
 */
func (a *TransferPickerControl) Placeholder(value interface{}) *TransferPickerControl {
    a.Set("placeholder", value)
    return a
}

/**
 * deferApi
 */
func (a *TransferPickerControl) DeferApi(value interface{}) *TransferPickerControl {
    a.Set("deferApi", value)
    return a
}

/**
 * disabled
 */
func (a *TransferPickerControl) Disabled(value interface{}) *TransferPickerControl {
    a.Set("disabled", value)
    return a
}

/**
 * size
 */
func (a *TransferPickerControl) Size(value interface{}) *TransferPickerControl {
    a.Set("size", value)
    return a
}

/**
 * visible
 */
func (a *TransferPickerControl) Visible(value interface{}) *TransferPickerControl {
    a.Set("visible", value)
    return a
}

/**
 * searchPlaceholder
 */
func (a *TransferPickerControl) SearchPlaceholder(value interface{}) *TransferPickerControl {
    a.Set("searchPlaceholder", value)
    return a
}

/**
 * showInvalidMatch
 */
func (a *TransferPickerControl) ShowInvalidMatch(value interface{}) *TransferPickerControl {
    a.Set("showInvalidMatch", value)
    return a
}

/**
 * deferField
 */
func (a *TransferPickerControl) DeferField(value interface{}) *TransferPickerControl {
    a.Set("deferField", value)
    return a
}
