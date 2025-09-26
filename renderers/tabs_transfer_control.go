package renderers


/**
 * TabsTransfer 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/tabs-transfer

 * @author wcz0
 * @version 6.2.2
 */
type TabsTransferControl struct {
	*BaseRenderer
}

func NewTabsTransferControl() *TabsTransferControl {
    a := &TabsTransferControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "tabs-transfer")
    return a
}


func (a *TabsTransferControl) Set(name string, value interface{}) *TabsTransferControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * itemHeight
 */
func (a *TabsTransferControl) ItemHeight(value interface{}) *TabsTransferControl {
    a.Set("itemHeight", value)
    return a
}

/**
 * creatable
 */
func (a *TabsTransferControl) Creatable(value interface{}) *TabsTransferControl {
    a.Set("creatable", value)
    return a
}

/**
 * removable
 */
func (a *TabsTransferControl) Removable(value interface{}) *TabsTransferControl {
    a.Set("removable", value)
    return a
}

/**
 * required
 */
func (a *TabsTransferControl) Required(value interface{}) *TabsTransferControl {
    a.Set("required", value)
    return a
}

/**
 * searchPlaceholder
 */
func (a *TabsTransferControl) SearchPlaceholder(value interface{}) *TabsTransferControl {
    a.Set("searchPlaceholder", value)
    return a
}

/**
 * showInvalidMatch
 */
func (a *TabsTransferControl) ShowInvalidMatch(value interface{}) *TabsTransferControl {
    a.Set("showInvalidMatch", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *TabsTransferControl) Width(value interface{}) *TabsTransferControl {
    a.Set("width", value)
    return a
}

/**
 * selectFirst
 */
func (a *TabsTransferControl) SelectFirst(value interface{}) *TabsTransferControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * clearable
 */
func (a *TabsTransferControl) Clearable(value interface{}) *TabsTransferControl {
    a.Set("clearable", value)
    return a
}

/**
 * addApi
 */
func (a *TabsTransferControl) AddApi(value interface{}) *TabsTransferControl {
    a.Set("addApi", value)
    return a
}

/**
 * labelWidth
 */
func (a *TabsTransferControl) LabelWidth(value interface{}) *TabsTransferControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * id
 */
func (a *TabsTransferControl) Id(value interface{}) *TabsTransferControl {
    a.Set("id", value)
    return a
}

/**
 * style
 */
func (a *TabsTransferControl) Style(value interface{}) *TabsTransferControl {
    a.Set("style", value)
    return a
}

/**
 */
func (a *TabsTransferControl) Type(value interface{}) *TabsTransferControl {
    a.Set("type", value)
    return a
}

/**
 * resultTitle
 */
func (a *TabsTransferControl) ResultTitle(value interface{}) *TabsTransferControl {
    a.Set("resultTitle", value)
    return a
}

/**
 * options
 */
func (a *TabsTransferControl) Options(value interface{}) *TabsTransferControl {
    a.Set("options", value)
    return a
}

/**
 * extraName
 */
func (a *TabsTransferControl) ExtraName(value interface{}) *TabsTransferControl {
    a.Set("extraName", value)
    return a
}

/**
 * remark
 */
func (a *TabsTransferControl) Remark(value interface{}) *TabsTransferControl {
    a.Set("remark", value)
    return a
}

/**
 * inputClassName
 */
func (a *TabsTransferControl) InputClassName(value interface{}) *TabsTransferControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * resultSearchPlaceholder
 */
func (a *TabsTransferControl) ResultSearchPlaceholder(value interface{}) *TabsTransferControl {
    a.Set("resultSearchPlaceholder", value)
    return a
}

/**
 * autoCheckChildren
 */
func (a *TabsTransferControl) AutoCheckChildren(value interface{}) *TabsTransferControl {
    a.Set("autoCheckChildren", value)
    return a
}

/**
 * pagination
 */
func (a *TabsTransferControl) Pagination(value interface{}) *TabsTransferControl {
    a.Set("pagination", value)
    return a
}

/**
 * delimiter
 */
func (a *TabsTransferControl) Delimiter(value interface{}) *TabsTransferControl {
    a.Set("delimiter", value)
    return a
}

/**
 * labelAlign
 */
func (a *TabsTransferControl) LabelAlign(value interface{}) *TabsTransferControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * labelClassName
 */
func (a *TabsTransferControl) LabelClassName(value interface{}) *TabsTransferControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * staticOn
 */
func (a *TabsTransferControl) StaticOn(value interface{}) *TabsTransferControl {
    a.Set("staticOn", value)
    return a
}

/**
 * selectTitle
 */
func (a *TabsTransferControl) SelectTitle(value interface{}) *TabsTransferControl {
    a.Set("selectTitle", value)
    return a
}

/**
 * onlyChildren
 */
func (a *TabsTransferControl) OnlyChildren(value interface{}) *TabsTransferControl {
    a.Set("onlyChildren", value)
    return a
}

/**
 * joinValues
 */
func (a *TabsTransferControl) JoinValues(value interface{}) *TabsTransferControl {
    a.Set("joinValues", value)
    return a
}

/**
 * deferApi
 */
func (a *TabsTransferControl) DeferApi(value interface{}) *TabsTransferControl {
    a.Set("deferApi", value)
    return a
}

/**
 * validateApi
 */
func (a *TabsTransferControl) ValidateApi(value interface{}) *TabsTransferControl {
    a.Set("validateApi", value)
    return a
}

/**
 * submitOnChange
 */
func (a *TabsTransferControl) SubmitOnChange(value interface{}) *TabsTransferControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *TabsTransferControl) ReadOnlyOn(value interface{}) *TabsTransferControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * validations
 */
func (a *TabsTransferControl) Validations(value interface{}) *TabsTransferControl {
    a.Set("validations", value)
    return a
}

/**
 * editorSetting
 */
func (a *TabsTransferControl) EditorSetting(value interface{}) *TabsTransferControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * leftOptions
 */
func (a *TabsTransferControl) LeftOptions(value interface{}) *TabsTransferControl {
    a.Set("leftOptions", value)
    return a
}

/**
 * searchResultColumns
 */
func (a *TabsTransferControl) SearchResultColumns(value interface{}) *TabsTransferControl {
    a.Set("searchResultColumns", value)
    return a
}

/**
 * deferField
 */
func (a *TabsTransferControl) DeferField(value interface{}) *TabsTransferControl {
    a.Set("deferField", value)
    return a
}

/**
 * autoFill
 */
func (a *TabsTransferControl) AutoFill(value interface{}) *TabsTransferControl {
    a.Set("autoFill", value)
    return a
}

/**
 * hidden
 */
func (a *TabsTransferControl) Hidden(value interface{}) *TabsTransferControl {
    a.Set("hidden", value)
    return a
}

/**
 * onEvent
 */
func (a *TabsTransferControl) OnEvent(value interface{}) *TabsTransferControl {
    a.Set("onEvent", value)
    return a
}

/**
 * staticSchema
 */
func (a *TabsTransferControl) StaticSchema(value interface{}) *TabsTransferControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * resultListModeFollowSelect
 */
func (a *TabsTransferControl) ResultListModeFollowSelect(value interface{}) *TabsTransferControl {
    a.Set("resultListModeFollowSelect", value)
    return a
}

/**
 * columns
 */
func (a *TabsTransferControl) Columns(value interface{}) *TabsTransferControl {
    a.Set("columns", value)
    return a
}

/**
 * resetValue
 */
func (a *TabsTransferControl) ResetValue(value interface{}) *TabsTransferControl {
    a.Set("resetValue", value)
    return a
}

/**
 * editApi
 */
func (a *TabsTransferControl) EditApi(value interface{}) *TabsTransferControl {
    a.Set("editApi", value)
    return a
}

/**
 * labelRemark
 */
func (a *TabsTransferControl) LabelRemark(value interface{}) *TabsTransferControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * description
 */
func (a *TabsTransferControl) Description(value interface{}) *TabsTransferControl {
    a.Set("description", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *TabsTransferControl) ClearValueOnHidden(value interface{}) *TabsTransferControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * visibleOn
 */
func (a *TabsTransferControl) VisibleOn(value interface{}) *TabsTransferControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * useMobileUI
 */
func (a *TabsTransferControl) UseMobileUI(value interface{}) *TabsTransferControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * selectMode
 */
func (a *TabsTransferControl) SelectMode(value interface{}) *TabsTransferControl {
    a.Set("selectMode", value)
    return a
}

/**
 * addDialog
 */
func (a *TabsTransferControl) AddDialog(value interface{}) *TabsTransferControl {
    a.Set("addDialog", value)
    return a
}

/**
 * editDialog
 */
func (a *TabsTransferControl) EditDialog(value interface{}) *TabsTransferControl {
    a.Set("editDialog", value)
    return a
}

/**
 * mode
 */
func (a *TabsTransferControl) Mode(value interface{}) *TabsTransferControl {
    a.Set("mode", value)
    return a
}

/**
 * row
 */
func (a *TabsTransferControl) Row(value interface{}) *TabsTransferControl {
    a.Set("row", value)
    return a
}

/**
 * className
 */
func (a *TabsTransferControl) ClassName(value interface{}) *TabsTransferControl {
    a.Set("className", value)
    return a
}

/**
 * hiddenOn
 */
func (a *TabsTransferControl) HiddenOn(value interface{}) *TabsTransferControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * static
 */
func (a *TabsTransferControl) Static(value interface{}) *TabsTransferControl {
    a.Set("static", value)
    return a
}

/**
 * rightMode
 */
func (a *TabsTransferControl) RightMode(value interface{}) *TabsTransferControl {
    a.Set("rightMode", value)
    return a
}

/**
 * multiple
 */
func (a *TabsTransferControl) Multiple(value interface{}) *TabsTransferControl {
    a.Set("multiple", value)
    return a
}

/**
 * labelOverflow
 */
func (a *TabsTransferControl) LabelOverflow(value interface{}) *TabsTransferControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * disabled
 */
func (a *TabsTransferControl) Disabled(value interface{}) *TabsTransferControl {
    a.Set("disabled", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *TabsTransferControl) StaticInputClassName(value interface{}) *TabsTransferControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * resultSearchable
 */
func (a *TabsTransferControl) ResultSearchable(value interface{}) *TabsTransferControl {
    a.Set("resultSearchable", value)
    return a
}

/**
 * virtualThreshold
 */
func (a *TabsTransferControl) VirtualThreshold(value interface{}) *TabsTransferControl {
    a.Set("virtualThreshold", value)
    return a
}

/**
 * initiallyOpen
 */
func (a *TabsTransferControl) InitiallyOpen(value interface{}) *TabsTransferControl {
    a.Set("initiallyOpen", value)
    return a
}

/**
 * initFetchOn
 */
func (a *TabsTransferControl) InitFetchOn(value interface{}) *TabsTransferControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * inline
 */
func (a *TabsTransferControl) Inline(value interface{}) *TabsTransferControl {
    a.Set("inline", value)
    return a
}

/**
 * disabledOn
 */
func (a *TabsTransferControl) DisabledOn(value interface{}) *TabsTransferControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *TabsTransferControl) StaticPlaceholder(value interface{}) *TabsTransferControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *TabsTransferControl) StaticLabelClassName(value interface{}) *TabsTransferControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * size
 */
func (a *TabsTransferControl) Size(value interface{}) *TabsTransferControl {
    a.Set("size", value)
    return a
}

/**
 * sortable
 */
func (a *TabsTransferControl) Sortable(value interface{}) *TabsTransferControl {
    a.Set("sortable", value)
    return a
}

/**
 * searchApi
 */
func (a *TabsTransferControl) SearchApi(value interface{}) *TabsTransferControl {
    a.Set("searchApi", value)
    return a
}

/**
 * valuesNoWrap
 */
func (a *TabsTransferControl) ValuesNoWrap(value interface{}) *TabsTransferControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * extractValue
 */
func (a *TabsTransferControl) ExtractValue(value interface{}) *TabsTransferControl {
    a.Set("extractValue", value)
    return a
}

/**
 * editable
 */
func (a *TabsTransferControl) Editable(value interface{}) *TabsTransferControl {
    a.Set("editable", value)
    return a
}

/**
 * validateOnChange
 */
func (a *TabsTransferControl) ValidateOnChange(value interface{}) *TabsTransferControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * desc
 */
func (a *TabsTransferControl) Desc(value interface{}) *TabsTransferControl {
    a.Set("desc", value)
    return a
}

/**
 * valueTpl
 */
func (a *TabsTransferControl) ValueTpl(value interface{}) *TabsTransferControl {
    a.Set("valueTpl", value)
    return a
}

/**
 * statistics
 */
func (a *TabsTransferControl) Statistics(value interface{}) *TabsTransferControl {
    a.Set("statistics", value)
    return a
}

/**
 * loadingConfig
 */
func (a *TabsTransferControl) LoadingConfig(value interface{}) *TabsTransferControl {
    a.Set("loadingConfig", value)
    return a
}

/**
 * deleteApi
 */
func (a *TabsTransferControl) DeleteApi(value interface{}) *TabsTransferControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * clearValueOnSourceChange
 */
func (a *TabsTransferControl) ClearValueOnSourceChange(value interface{}) *TabsTransferControl {
    a.Set("clearValueOnSourceChange", value)
    return a
}

/**
 * placeholder
 */
func (a *TabsTransferControl) Placeholder(value interface{}) *TabsTransferControl {
    a.Set("placeholder", value)
    return a
}

/**
 * value
 */
func (a *TabsTransferControl) Value(value interface{}) *TabsTransferControl {
    a.Set("value", value)
    return a
}

/**
 * visible
 */
func (a *TabsTransferControl) Visible(value interface{}) *TabsTransferControl {
    a.Set("visible", value)
    return a
}

/**
 * showArrow
 */
func (a *TabsTransferControl) ShowArrow(value interface{}) *TabsTransferControl {
    a.Set("showArrow", value)
    return a
}

/**
 * searchable
 */
func (a *TabsTransferControl) Searchable(value interface{}) *TabsTransferControl {
    a.Set("searchable", value)
    return a
}

/**
 * source
 */
func (a *TabsTransferControl) Source(value interface{}) *TabsTransferControl {
    a.Set("source", value)
    return a
}

/**
 * initFetch
 */
func (a *TabsTransferControl) InitFetch(value interface{}) *TabsTransferControl {
    a.Set("initFetch", value)
    return a
}

/**
 * editControls
 */
func (a *TabsTransferControl) EditControls(value interface{}) *TabsTransferControl {
    a.Set("editControls", value)
    return a
}

/**
 * label
 */
func (a *TabsTransferControl) Label(value interface{}) *TabsTransferControl {
    a.Set("label", value)
    return a
}

/**
 * readOnly
 */
func (a *TabsTransferControl) ReadOnly(value interface{}) *TabsTransferControl {
    a.Set("readOnly", value)
    return a
}

/**
 * initAutoFill
 */
func (a *TabsTransferControl) InitAutoFill(value interface{}) *TabsTransferControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * staticClassName
 */
func (a *TabsTransferControl) StaticClassName(value interface{}) *TabsTransferControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * checkAll
 */
func (a *TabsTransferControl) CheckAll(value interface{}) *TabsTransferControl {
    a.Set("checkAll", value)
    return a
}

/**
 * addControls
 */
func (a *TabsTransferControl) AddControls(value interface{}) *TabsTransferControl {
    a.Set("addControls", value)
    return a
}

/**
 * deleteConfirmText
 */
func (a *TabsTransferControl) DeleteConfirmText(value interface{}) *TabsTransferControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *TabsTransferControl) DescriptionClassName(value interface{}) *TabsTransferControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * leftMode
 */
func (a *TabsTransferControl) LeftMode(value interface{}) *TabsTransferControl {
    a.Set("leftMode", value)
    return a
}

/**
 * searchResultMode
 */
func (a *TabsTransferControl) SearchResultMode(value interface{}) *TabsTransferControl {
    a.Set("searchResultMode", value)
    return a
}

/**
 * createBtnLabel
 */
func (a *TabsTransferControl) CreateBtnLabel(value interface{}) *TabsTransferControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * name
 */
func (a *TabsTransferControl) Name(value interface{}) *TabsTransferControl {
    a.Set("name", value)
    return a
}

/**
 * hint
 */
func (a *TabsTransferControl) Hint(value interface{}) *TabsTransferControl {
    a.Set("hint", value)
    return a
}

/**
 * horizontal
 */
func (a *TabsTransferControl) Horizontal(value interface{}) *TabsTransferControl {
    a.Set("horizontal", value)
    return a
}

/**
 * validationErrors
 */
func (a *TabsTransferControl) ValidationErrors(value interface{}) *TabsTransferControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * menuTpl
 */
func (a *TabsTransferControl) MenuTpl(value interface{}) *TabsTransferControl {
    a.Set("menuTpl", value)
    return a
}
