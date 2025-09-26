package renderers


/**
 * TabsTransferPicker 穿梭器的弹框形态 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/tabs-transfer-picker

 * @author wcz0
 * @version 6.2.2
 */
type TabsTransferPickerControl struct {
	*BaseRenderer
}

func NewTabsTransferPickerControl() *TabsTransferPickerControl {
    a := &TabsTransferPickerControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "tabs-transfer-picker")
    return a
}


func (a *TabsTransferPickerControl) Set(name string, value interface{}) *TabsTransferPickerControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * hidden
 */
func (a *TabsTransferPickerControl) Hidden(value interface{}) *TabsTransferPickerControl {
    a.Set("hidden", value)
    return a
}

/**
 * hiddenOn
 */
func (a *TabsTransferPickerControl) HiddenOn(value interface{}) *TabsTransferPickerControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * valueTpl
 */
func (a *TabsTransferPickerControl) ValueTpl(value interface{}) *TabsTransferPickerControl {
    a.Set("valueTpl", value)
    return a
}

/**
 * creatable
 */
func (a *TabsTransferPickerControl) Creatable(value interface{}) *TabsTransferPickerControl {
    a.Set("creatable", value)
    return a
}

/**
 * label
 */
func (a *TabsTransferPickerControl) Label(value interface{}) *TabsTransferPickerControl {
    a.Set("label", value)
    return a
}

/**
 * labelOverflow
 */
func (a *TabsTransferPickerControl) LabelOverflow(value interface{}) *TabsTransferPickerControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * extraName
 */
func (a *TabsTransferPickerControl) ExtraName(value interface{}) *TabsTransferPickerControl {
    a.Set("extraName", value)
    return a
}

/**
 * labelRemark
 */
func (a *TabsTransferPickerControl) LabelRemark(value interface{}) *TabsTransferPickerControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *TabsTransferPickerControl) StaticPlaceholder(value interface{}) *TabsTransferPickerControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *TabsTransferPickerControl) StaticInputClassName(value interface{}) *TabsTransferPickerControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * selectTitle
 */
func (a *TabsTransferPickerControl) SelectTitle(value interface{}) *TabsTransferPickerControl {
    a.Set("selectTitle", value)
    return a
}

/**
 * joinValues
 */
func (a *TabsTransferPickerControl) JoinValues(value interface{}) *TabsTransferPickerControl {
    a.Set("joinValues", value)
    return a
}

/**
 * delimiter
 */
func (a *TabsTransferPickerControl) Delimiter(value interface{}) *TabsTransferPickerControl {
    a.Set("delimiter", value)
    return a
}

/**
 * deleteConfirmText
 */
func (a *TabsTransferPickerControl) DeleteConfirmText(value interface{}) *TabsTransferPickerControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * validationErrors
 */
func (a *TabsTransferPickerControl) ValidationErrors(value interface{}) *TabsTransferPickerControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * initAutoFill
 */
func (a *TabsTransferPickerControl) InitAutoFill(value interface{}) *TabsTransferPickerControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * leftMode
 */
func (a *TabsTransferPickerControl) LeftMode(value interface{}) *TabsTransferPickerControl {
    a.Set("leftMode", value)
    return a
}

/**
 * menuTpl
 */
func (a *TabsTransferPickerControl) MenuTpl(value interface{}) *TabsTransferPickerControl {
    a.Set("menuTpl", value)
    return a
}

/**
 * initiallyOpen
 */
func (a *TabsTransferPickerControl) InitiallyOpen(value interface{}) *TabsTransferPickerControl {
    a.Set("initiallyOpen", value)
    return a
}

/**
 * source
 */
func (a *TabsTransferPickerControl) Source(value interface{}) *TabsTransferPickerControl {
    a.Set("source", value)
    return a
}

/**
 * checkAll
 */
func (a *TabsTransferPickerControl) CheckAll(value interface{}) *TabsTransferPickerControl {
    a.Set("checkAll", value)
    return a
}

/**
 * clearable
 */
func (a *TabsTransferPickerControl) Clearable(value interface{}) *TabsTransferPickerControl {
    a.Set("clearable", value)
    return a
}

/**
 * remark
 */
func (a *TabsTransferPickerControl) Remark(value interface{}) *TabsTransferPickerControl {
    a.Set("remark", value)
    return a
}

/**
 * submitOnChange
 */
func (a *TabsTransferPickerControl) SubmitOnChange(value interface{}) *TabsTransferPickerControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * horizontal
 */
func (a *TabsTransferPickerControl) Horizontal(value interface{}) *TabsTransferPickerControl {
    a.Set("horizontal", value)
    return a
}

/**
 * staticClassName
 */
func (a *TabsTransferPickerControl) StaticClassName(value interface{}) *TabsTransferPickerControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *TabsTransferPickerControl) StaticLabelClassName(value interface{}) *TabsTransferPickerControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * initFetch
 */
func (a *TabsTransferPickerControl) InitFetch(value interface{}) *TabsTransferPickerControl {
    a.Set("initFetch", value)
    return a
}

/**
 * multiple
 */
func (a *TabsTransferPickerControl) Multiple(value interface{}) *TabsTransferPickerControl {
    a.Set("multiple", value)
    return a
}

/**
 * valuesNoWrap
 */
func (a *TabsTransferPickerControl) ValuesNoWrap(value interface{}) *TabsTransferPickerControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * labelClassName
 */
func (a *TabsTransferPickerControl) LabelClassName(value interface{}) *TabsTransferPickerControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *TabsTransferPickerControl) ClearValueOnHidden(value interface{}) *TabsTransferPickerControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * disabled
 */
func (a *TabsTransferPickerControl) Disabled(value interface{}) *TabsTransferPickerControl {
    a.Set("disabled", value)
    return a
}

/**
 * id
 */
func (a *TabsTransferPickerControl) Id(value interface{}) *TabsTransferPickerControl {
    a.Set("id", value)
    return a
}

/**
 * columns
 */
func (a *TabsTransferPickerControl) Columns(value interface{}) *TabsTransferPickerControl {
    a.Set("columns", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *TabsTransferPickerControl) DescriptionClassName(value interface{}) *TabsTransferPickerControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * required
 */
func (a *TabsTransferPickerControl) Required(value interface{}) *TabsTransferPickerControl {
    a.Set("required", value)
    return a
}

/**
 * value
 */
func (a *TabsTransferPickerControl) Value(value interface{}) *TabsTransferPickerControl {
    a.Set("value", value)
    return a
}

/**
 * disabledOn
 */
func (a *TabsTransferPickerControl) DisabledOn(value interface{}) *TabsTransferPickerControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * useMobileUI
 */
func (a *TabsTransferPickerControl) UseMobileUI(value interface{}) *TabsTransferPickerControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * searchPlaceholder
 */
func (a *TabsTransferPickerControl) SearchPlaceholder(value interface{}) *TabsTransferPickerControl {
    a.Set("searchPlaceholder", value)
    return a
}

/**
 * statistics
 */
func (a *TabsTransferPickerControl) Statistics(value interface{}) *TabsTransferPickerControl {
    a.Set("statistics", value)
    return a
}

/**
 * onlyChildren
 */
func (a *TabsTransferPickerControl) OnlyChildren(value interface{}) *TabsTransferPickerControl {
    a.Set("onlyChildren", value)
    return a
}

/**
 * mode
 */
func (a *TabsTransferPickerControl) Mode(value interface{}) *TabsTransferPickerControl {
    a.Set("mode", value)
    return a
}

/**
 * inline
 */
func (a *TabsTransferPickerControl) Inline(value interface{}) *TabsTransferPickerControl {
    a.Set("inline", value)
    return a
}

/**
 * visible
 */
func (a *TabsTransferPickerControl) Visible(value interface{}) *TabsTransferPickerControl {
    a.Set("visible", value)
    return a
}

/**
 * staticSchema
 */
func (a *TabsTransferPickerControl) StaticSchema(value interface{}) *TabsTransferPickerControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * style
 */
func (a *TabsTransferPickerControl) Style(value interface{}) *TabsTransferPickerControl {
    a.Set("style", value)
    return a
}

/**
 * selectMode
 */
func (a *TabsTransferPickerControl) SelectMode(value interface{}) *TabsTransferPickerControl {
    a.Set("selectMode", value)
    return a
}

/**
 * searchResultColumns
 */
func (a *TabsTransferPickerControl) SearchResultColumns(value interface{}) *TabsTransferPickerControl {
    a.Set("searchResultColumns", value)
    return a
}

/**
 * createBtnLabel
 */
func (a *TabsTransferPickerControl) CreateBtnLabel(value interface{}) *TabsTransferPickerControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * validateOnChange
 */
func (a *TabsTransferPickerControl) ValidateOnChange(value interface{}) *TabsTransferPickerControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * className
 */
func (a *TabsTransferPickerControl) ClassName(value interface{}) *TabsTransferPickerControl {
    a.Set("className", value)
    return a
}

/**
 */
func (a *TabsTransferPickerControl) Type(value interface{}) *TabsTransferPickerControl {
    a.Set("type", value)
    return a
}

/**
 * showArrow
 */
func (a *TabsTransferPickerControl) ShowArrow(value interface{}) *TabsTransferPickerControl {
    a.Set("showArrow", value)
    return a
}

/**
 * sortable
 */
func (a *TabsTransferPickerControl) Sortable(value interface{}) *TabsTransferPickerControl {
    a.Set("sortable", value)
    return a
}

/**
 * virtualThreshold
 */
func (a *TabsTransferPickerControl) VirtualThreshold(value interface{}) *TabsTransferPickerControl {
    a.Set("virtualThreshold", value)
    return a
}

/**
 * autoCheckChildren
 */
func (a *TabsTransferPickerControl) AutoCheckChildren(value interface{}) *TabsTransferPickerControl {
    a.Set("autoCheckChildren", value)
    return a
}

/**
 * deleteApi
 */
func (a *TabsTransferPickerControl) DeleteApi(value interface{}) *TabsTransferPickerControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * row
 */
func (a *TabsTransferPickerControl) Row(value interface{}) *TabsTransferPickerControl {
    a.Set("row", value)
    return a
}

/**
 * resultListModeFollowSelect
 */
func (a *TabsTransferPickerControl) ResultListModeFollowSelect(value interface{}) *TabsTransferPickerControl {
    a.Set("resultListModeFollowSelect", value)
    return a
}

/**
 * clearValueOnSourceChange
 */
func (a *TabsTransferPickerControl) ClearValueOnSourceChange(value interface{}) *TabsTransferPickerControl {
    a.Set("clearValueOnSourceChange", value)
    return a
}

/**
 * labelWidth
 */
func (a *TabsTransferPickerControl) LabelWidth(value interface{}) *TabsTransferPickerControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * name
 */
func (a *TabsTransferPickerControl) Name(value interface{}) *TabsTransferPickerControl {
    a.Set("name", value)
    return a
}

/**
 * autoFill
 */
func (a *TabsTransferPickerControl) AutoFill(value interface{}) *TabsTransferPickerControl {
    a.Set("autoFill", value)
    return a
}

/**
 * static
 */
func (a *TabsTransferPickerControl) Static(value interface{}) *TabsTransferPickerControl {
    a.Set("static", value)
    return a
}

/**
 * rightMode
 */
func (a *TabsTransferPickerControl) RightMode(value interface{}) *TabsTransferPickerControl {
    a.Set("rightMode", value)
    return a
}

/**
 * itemHeight
 */
func (a *TabsTransferPickerControl) ItemHeight(value interface{}) *TabsTransferPickerControl {
    a.Set("itemHeight", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *TabsTransferPickerControl) Width(value interface{}) *TabsTransferPickerControl {
    a.Set("width", value)
    return a
}

/**
 * editControls
 */
func (a *TabsTransferPickerControl) EditControls(value interface{}) *TabsTransferPickerControl {
    a.Set("editControls", value)
    return a
}

/**
 * readOnly
 */
func (a *TabsTransferPickerControl) ReadOnly(value interface{}) *TabsTransferPickerControl {
    a.Set("readOnly", value)
    return a
}

/**
 * onEvent
 */
func (a *TabsTransferPickerControl) OnEvent(value interface{}) *TabsTransferPickerControl {
    a.Set("onEvent", value)
    return a
}

/**
 * size
 */
func (a *TabsTransferPickerControl) Size(value interface{}) *TabsTransferPickerControl {
    a.Set("size", value)
    return a
}

/**
 * leftOptions
 */
func (a *TabsTransferPickerControl) LeftOptions(value interface{}) *TabsTransferPickerControl {
    a.Set("leftOptions", value)
    return a
}

/**
 * searchApi
 */
func (a *TabsTransferPickerControl) SearchApi(value interface{}) *TabsTransferPickerControl {
    a.Set("searchApi", value)
    return a
}

/**
 * pagination
 */
func (a *TabsTransferPickerControl) Pagination(value interface{}) *TabsTransferPickerControl {
    a.Set("pagination", value)
    return a
}

/**
 * editDialog
 */
func (a *TabsTransferPickerControl) EditDialog(value interface{}) *TabsTransferPickerControl {
    a.Set("editDialog", value)
    return a
}

/**
 * removable
 */
func (a *TabsTransferPickerControl) Removable(value interface{}) *TabsTransferPickerControl {
    a.Set("removable", value)
    return a
}

/**
 * deferField
 */
func (a *TabsTransferPickerControl) DeferField(value interface{}) *TabsTransferPickerControl {
    a.Set("deferField", value)
    return a
}

/**
 * editable
 */
func (a *TabsTransferPickerControl) Editable(value interface{}) *TabsTransferPickerControl {
    a.Set("editable", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *TabsTransferPickerControl) ReadOnlyOn(value interface{}) *TabsTransferPickerControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * desc
 */
func (a *TabsTransferPickerControl) Desc(value interface{}) *TabsTransferPickerControl {
    a.Set("desc", value)
    return a
}

/**
 * validations
 */
func (a *TabsTransferPickerControl) Validations(value interface{}) *TabsTransferPickerControl {
    a.Set("validations", value)
    return a
}

/**
 * selectFirst
 */
func (a *TabsTransferPickerControl) SelectFirst(value interface{}) *TabsTransferPickerControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * initFetchOn
 */
func (a *TabsTransferPickerControl) InitFetchOn(value interface{}) *TabsTransferPickerControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * resetValue
 */
func (a *TabsTransferPickerControl) ResetValue(value interface{}) *TabsTransferPickerControl {
    a.Set("resetValue", value)
    return a
}

/**
 * deferApi
 */
func (a *TabsTransferPickerControl) DeferApi(value interface{}) *TabsTransferPickerControl {
    a.Set("deferApi", value)
    return a
}

/**
 * inputClassName
 */
func (a *TabsTransferPickerControl) InputClassName(value interface{}) *TabsTransferPickerControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * placeholder
 */
func (a *TabsTransferPickerControl) Placeholder(value interface{}) *TabsTransferPickerControl {
    a.Set("placeholder", value)
    return a
}

/**
 * validateApi
 */
func (a *TabsTransferPickerControl) ValidateApi(value interface{}) *TabsTransferPickerControl {
    a.Set("validateApi", value)
    return a
}

/**
 * searchable
 */
func (a *TabsTransferPickerControl) Searchable(value interface{}) *TabsTransferPickerControl {
    a.Set("searchable", value)
    return a
}

/**
 * options
 */
func (a *TabsTransferPickerControl) Options(value interface{}) *TabsTransferPickerControl {
    a.Set("options", value)
    return a
}

/**
 * editApi
 */
func (a *TabsTransferPickerControl) EditApi(value interface{}) *TabsTransferPickerControl {
    a.Set("editApi", value)
    return a
}

/**
 * staticOn
 */
func (a *TabsTransferPickerControl) StaticOn(value interface{}) *TabsTransferPickerControl {
    a.Set("staticOn", value)
    return a
}

/**
 * searchResultMode
 */
func (a *TabsTransferPickerControl) SearchResultMode(value interface{}) *TabsTransferPickerControl {
    a.Set("searchResultMode", value)
    return a
}

/**
 * resultSearchable
 */
func (a *TabsTransferPickerControl) ResultSearchable(value interface{}) *TabsTransferPickerControl {
    a.Set("resultSearchable", value)
    return a
}

/**
 * resultTitle
 */
func (a *TabsTransferPickerControl) ResultTitle(value interface{}) *TabsTransferPickerControl {
    a.Set("resultTitle", value)
    return a
}

/**
 * resultSearchPlaceholder
 */
func (a *TabsTransferPickerControl) ResultSearchPlaceholder(value interface{}) *TabsTransferPickerControl {
    a.Set("resultSearchPlaceholder", value)
    return a
}

/**
 * showInvalidMatch
 */
func (a *TabsTransferPickerControl) ShowInvalidMatch(value interface{}) *TabsTransferPickerControl {
    a.Set("showInvalidMatch", value)
    return a
}

/**
 * extractValue
 */
func (a *TabsTransferPickerControl) ExtractValue(value interface{}) *TabsTransferPickerControl {
    a.Set("extractValue", value)
    return a
}

/**
 * visibleOn
 */
func (a *TabsTransferPickerControl) VisibleOn(value interface{}) *TabsTransferPickerControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * loadingConfig
 */
func (a *TabsTransferPickerControl) LoadingConfig(value interface{}) *TabsTransferPickerControl {
    a.Set("loadingConfig", value)
    return a
}

/**
 * addControls
 */
func (a *TabsTransferPickerControl) AddControls(value interface{}) *TabsTransferPickerControl {
    a.Set("addControls", value)
    return a
}

/**
 * labelAlign
 */
func (a *TabsTransferPickerControl) LabelAlign(value interface{}) *TabsTransferPickerControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * description
 */
func (a *TabsTransferPickerControl) Description(value interface{}) *TabsTransferPickerControl {
    a.Set("description", value)
    return a
}

/**
 * editorSetting
 */
func (a *TabsTransferPickerControl) EditorSetting(value interface{}) *TabsTransferPickerControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * addApi
 */
func (a *TabsTransferPickerControl) AddApi(value interface{}) *TabsTransferPickerControl {
    a.Set("addApi", value)
    return a
}

/**
 * addDialog
 */
func (a *TabsTransferPickerControl) AddDialog(value interface{}) *TabsTransferPickerControl {
    a.Set("addDialog", value)
    return a
}

/**
 * hint
 */
func (a *TabsTransferPickerControl) Hint(value interface{}) *TabsTransferPickerControl {
    a.Set("hint", value)
    return a
}
