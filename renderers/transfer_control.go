package renderers


/**
 * Transfer 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/transfer

 * @author wcz0
 * @version 6.2.2
 */
type TransferControl struct {
	*BaseRenderer
}

func NewTransferControl() *TransferControl {
    a := &TransferControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "transfer")
    return a
}


func (a *TransferControl) Set(name string, value interface{}) *TransferControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * label
 */
func (a *TransferControl) Label(value interface{}) *TransferControl {
    a.Set("label", value)
    return a
}

/**
 * autoFill
 */
func (a *TransferControl) AutoFill(value interface{}) *TransferControl {
    a.Set("autoFill", value)
    return a
}

/**
 * addApi
 */
func (a *TransferControl) AddApi(value interface{}) *TransferControl {
    a.Set("addApi", value)
    return a
}

/**
 * creatable
 */
func (a *TransferControl) Creatable(value interface{}) *TransferControl {
    a.Set("creatable", value)
    return a
}

/**
 * removable
 */
func (a *TransferControl) Removable(value interface{}) *TransferControl {
    a.Set("removable", value)
    return a
}

/**
 * labelOverflow
 */
func (a *TransferControl) LabelOverflow(value interface{}) *TransferControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * extraName
 */
func (a *TransferControl) ExtraName(value interface{}) *TransferControl {
    a.Set("extraName", value)
    return a
}

/**
 * submitOnChange
 */
func (a *TransferControl) SubmitOnChange(value interface{}) *TransferControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *TransferControl) ClearValueOnHidden(value interface{}) *TransferControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * row
 */
func (a *TransferControl) Row(value interface{}) *TransferControl {
    a.Set("row", value)
    return a
}

/**
 * delimiter
 */
func (a *TransferControl) Delimiter(value interface{}) *TransferControl {
    a.Set("delimiter", value)
    return a
}

/**
 * clearValueOnSourceChange
 */
func (a *TransferControl) ClearValueOnSourceChange(value interface{}) *TransferControl {
    a.Set("clearValueOnSourceChange", value)
    return a
}

/**
 * labelAlign
 */
func (a *TransferControl) LabelAlign(value interface{}) *TransferControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *TransferControl) ReadOnlyOn(value interface{}) *TransferControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * validateOnChange
 */
func (a *TransferControl) ValidateOnChange(value interface{}) *TransferControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * validateApi
 */
func (a *TransferControl) ValidateApi(value interface{}) *TransferControl {
    a.Set("validateApi", value)
    return a
}

/**
 * visible
 */
func (a *TransferControl) Visible(value interface{}) *TransferControl {
    a.Set("visible", value)
    return a
}

/**
 * resultListModeFollowSelect
 */
func (a *TransferControl) ResultListModeFollowSelect(value interface{}) *TransferControl {
    a.Set("resultListModeFollowSelect", value)
    return a
}

/**
 * extractValue
 */
func (a *TransferControl) ExtractValue(value interface{}) *TransferControl {
    a.Set("extractValue", value)
    return a
}

/**
 * mode
 */
func (a *TransferControl) Mode(value interface{}) *TransferControl {
    a.Set("mode", value)
    return a
}

/**
 * initAutoFill
 */
func (a *TransferControl) InitAutoFill(value interface{}) *TransferControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * disabledOn
 */
func (a *TransferControl) DisabledOn(value interface{}) *TransferControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * hiddenOn
 */
func (a *TransferControl) HiddenOn(value interface{}) *TransferControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *TransferControl) StaticInputClassName(value interface{}) *TransferControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *TransferControl) Type(value interface{}) *TransferControl {
    a.Set("type", value)
    return a
}

/**
 * showArrow
 */
func (a *TransferControl) ShowArrow(value interface{}) *TransferControl {
    a.Set("showArrow", value)
    return a
}

/**
 * labelRemark
 */
func (a *TransferControl) LabelRemark(value interface{}) *TransferControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * useMobileUI
 */
func (a *TransferControl) UseMobileUI(value interface{}) *TransferControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * searchResultColumns
 */
func (a *TransferControl) SearchResultColumns(value interface{}) *TransferControl {
    a.Set("searchResultColumns", value)
    return a
}

/**
 * selectTitle
 */
func (a *TransferControl) SelectTitle(value interface{}) *TransferControl {
    a.Set("selectTitle", value)
    return a
}

/**
 * menuTpl
 */
func (a *TransferControl) MenuTpl(value interface{}) *TransferControl {
    a.Set("menuTpl", value)
    return a
}

/**
 * valueTpl
 */
func (a *TransferControl) ValueTpl(value interface{}) *TransferControl {
    a.Set("valueTpl", value)
    return a
}

/**
 * searchPlaceholder
 */
func (a *TransferControl) SearchPlaceholder(value interface{}) *TransferControl {
    a.Set("searchPlaceholder", value)
    return a
}

/**
 * showInvalidMatch
 */
func (a *TransferControl) ShowInvalidMatch(value interface{}) *TransferControl {
    a.Set("showInvalidMatch", value)
    return a
}

/**
 * deleteConfirmText
 */
func (a *TransferControl) DeleteConfirmText(value interface{}) *TransferControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * addControls
 */
func (a *TransferControl) AddControls(value interface{}) *TransferControl {
    a.Set("addControls", value)
    return a
}

/**
 * description
 */
func (a *TransferControl) Description(value interface{}) *TransferControl {
    a.Set("description", value)
    return a
}

/**
 * pagination
 */
func (a *TransferControl) Pagination(value interface{}) *TransferControl {
    a.Set("pagination", value)
    return a
}

/**
 * source
 */
func (a *TransferControl) Source(value interface{}) *TransferControl {
    a.Set("source", value)
    return a
}

/**
 * selectFirst
 */
func (a *TransferControl) SelectFirst(value interface{}) *TransferControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * disabled
 */
func (a *TransferControl) Disabled(value interface{}) *TransferControl {
    a.Set("disabled", value)
    return a
}

/**
 * visibleOn
 */
func (a *TransferControl) VisibleOn(value interface{}) *TransferControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * id
 */
func (a *TransferControl) Id(value interface{}) *TransferControl {
    a.Set("id", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *TransferControl) StaticLabelClassName(value interface{}) *TransferControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * searchResultMode
 */
func (a *TransferControl) SearchResultMode(value interface{}) *TransferControl {
    a.Set("searchResultMode", value)
    return a
}

/**
 * searchApi
 */
func (a *TransferControl) SearchApi(value interface{}) *TransferControl {
    a.Set("searchApi", value)
    return a
}

/**
 * editDialog
 */
func (a *TransferControl) EditDialog(value interface{}) *TransferControl {
    a.Set("editDialog", value)
    return a
}

/**
 * name
 */
func (a *TransferControl) Name(value interface{}) *TransferControl {
    a.Set("name", value)
    return a
}

/**
 * inline
 */
func (a *TransferControl) Inline(value interface{}) *TransferControl {
    a.Set("inline", value)
    return a
}

/**
 * leftMode
 */
func (a *TransferControl) LeftMode(value interface{}) *TransferControl {
    a.Set("leftMode", value)
    return a
}

/**
 * resultSearchable
 */
func (a *TransferControl) ResultSearchable(value interface{}) *TransferControl {
    a.Set("resultSearchable", value)
    return a
}

/**
 * resultSearchPlaceholder
 */
func (a *TransferControl) ResultSearchPlaceholder(value interface{}) *TransferControl {
    a.Set("resultSearchPlaceholder", value)
    return a
}

/**
 * statistics
 */
func (a *TransferControl) Statistics(value interface{}) *TransferControl {
    a.Set("statistics", value)
    return a
}

/**
 * virtualThreshold
 */
func (a *TransferControl) VirtualThreshold(value interface{}) *TransferControl {
    a.Set("virtualThreshold", value)
    return a
}

/**
 * valuesNoWrap
 */
func (a *TransferControl) ValuesNoWrap(value interface{}) *TransferControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * resetValue
 */
func (a *TransferControl) ResetValue(value interface{}) *TransferControl {
    a.Set("resetValue", value)
    return a
}

/**
 * loadingConfig
 */
func (a *TransferControl) LoadingConfig(value interface{}) *TransferControl {
    a.Set("loadingConfig", value)
    return a
}

/**
 * deferField
 */
func (a *TransferControl) DeferField(value interface{}) *TransferControl {
    a.Set("deferField", value)
    return a
}

/**
 * deferApi
 */
func (a *TransferControl) DeferApi(value interface{}) *TransferControl {
    a.Set("deferApi", value)
    return a
}

/**
 * labelClassName
 */
func (a *TransferControl) LabelClassName(value interface{}) *TransferControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * value
 */
func (a *TransferControl) Value(value interface{}) *TransferControl {
    a.Set("value", value)
    return a
}

/**
 * resultTitle
 */
func (a *TransferControl) ResultTitle(value interface{}) *TransferControl {
    a.Set("resultTitle", value)
    return a
}

/**
 * initFetch
 */
func (a *TransferControl) InitFetch(value interface{}) *TransferControl {
    a.Set("initFetch", value)
    return a
}

/**
 * editControls
 */
func (a *TransferControl) EditControls(value interface{}) *TransferControl {
    a.Set("editControls", value)
    return a
}

/**
 * labelWidth
 */
func (a *TransferControl) LabelWidth(value interface{}) *TransferControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * remark
 */
func (a *TransferControl) Remark(value interface{}) *TransferControl {
    a.Set("remark", value)
    return a
}

/**
 * hint
 */
func (a *TransferControl) Hint(value interface{}) *TransferControl {
    a.Set("hint", value)
    return a
}

/**
 * validations
 */
func (a *TransferControl) Validations(value interface{}) *TransferControl {
    a.Set("validations", value)
    return a
}

/**
 * onEvent
 */
func (a *TransferControl) OnEvent(value interface{}) *TransferControl {
    a.Set("onEvent", value)
    return a
}

/**
 * static
 */
func (a *TransferControl) Static(value interface{}) *TransferControl {
    a.Set("static", value)
    return a
}

/**
 * editable
 */
func (a *TransferControl) Editable(value interface{}) *TransferControl {
    a.Set("editable", value)
    return a
}

/**
 * sortable
 */
func (a *TransferControl) Sortable(value interface{}) *TransferControl {
    a.Set("sortable", value)
    return a
}

/**
 * leftOptions
 */
func (a *TransferControl) LeftOptions(value interface{}) *TransferControl {
    a.Set("leftOptions", value)
    return a
}

/**
 * autoCheckChildren
 */
func (a *TransferControl) AutoCheckChildren(value interface{}) *TransferControl {
    a.Set("autoCheckChildren", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *TransferControl) Width(value interface{}) *TransferControl {
    a.Set("width", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *TransferControl) DescriptionClassName(value interface{}) *TransferControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * onlyChildren
 */
func (a *TransferControl) OnlyChildren(value interface{}) *TransferControl {
    a.Set("onlyChildren", value)
    return a
}

/**
 * initiallyOpen
 */
func (a *TransferControl) InitiallyOpen(value interface{}) *TransferControl {
    a.Set("initiallyOpen", value)
    return a
}

/**
 * createBtnLabel
 */
func (a *TransferControl) CreateBtnLabel(value interface{}) *TransferControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * initFetchOn
 */
func (a *TransferControl) InitFetchOn(value interface{}) *TransferControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * inputClassName
 */
func (a *TransferControl) InputClassName(value interface{}) *TransferControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * staticOn
 */
func (a *TransferControl) StaticOn(value interface{}) *TransferControl {
    a.Set("staticOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *TransferControl) StaticPlaceholder(value interface{}) *TransferControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticClassName
 */
func (a *TransferControl) StaticClassName(value interface{}) *TransferControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * style
 */
func (a *TransferControl) Style(value interface{}) *TransferControl {
    a.Set("style", value)
    return a
}

/**
 * itemHeight
 */
func (a *TransferControl) ItemHeight(value interface{}) *TransferControl {
    a.Set("itemHeight", value)
    return a
}

/**
 * addDialog
 */
func (a *TransferControl) AddDialog(value interface{}) *TransferControl {
    a.Set("addDialog", value)
    return a
}

/**
 * editApi
 */
func (a *TransferControl) EditApi(value interface{}) *TransferControl {
    a.Set("editApi", value)
    return a
}

/**
 * deleteApi
 */
func (a *TransferControl) DeleteApi(value interface{}) *TransferControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * required
 */
func (a *TransferControl) Required(value interface{}) *TransferControl {
    a.Set("required", value)
    return a
}

/**
 * selectMode
 */
func (a *TransferControl) SelectMode(value interface{}) *TransferControl {
    a.Set("selectMode", value)
    return a
}

/**
 * rightMode
 */
func (a *TransferControl) RightMode(value interface{}) *TransferControl {
    a.Set("rightMode", value)
    return a
}

/**
 * searchable
 */
func (a *TransferControl) Searchable(value interface{}) *TransferControl {
    a.Set("searchable", value)
    return a
}

/**
 * readOnly
 */
func (a *TransferControl) ReadOnly(value interface{}) *TransferControl {
    a.Set("readOnly", value)
    return a
}

/**
 * desc
 */
func (a *TransferControl) Desc(value interface{}) *TransferControl {
    a.Set("desc", value)
    return a
}

/**
 * hidden
 */
func (a *TransferControl) Hidden(value interface{}) *TransferControl {
    a.Set("hidden", value)
    return a
}

/**
 * staticSchema
 */
func (a *TransferControl) StaticSchema(value interface{}) *TransferControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * editorSetting
 */
func (a *TransferControl) EditorSetting(value interface{}) *TransferControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * columns
 */
func (a *TransferControl) Columns(value interface{}) *TransferControl {
    a.Set("columns", value)
    return a
}

/**
 * horizontal
 */
func (a *TransferControl) Horizontal(value interface{}) *TransferControl {
    a.Set("horizontal", value)
    return a
}

/**
 * placeholder
 */
func (a *TransferControl) Placeholder(value interface{}) *TransferControl {
    a.Set("placeholder", value)
    return a
}

/**
 * validationErrors
 */
func (a *TransferControl) ValidationErrors(value interface{}) *TransferControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * className
 */
func (a *TransferControl) ClassName(value interface{}) *TransferControl {
    a.Set("className", value)
    return a
}

/**
 * size
 */
func (a *TransferControl) Size(value interface{}) *TransferControl {
    a.Set("size", value)
    return a
}

/**
 * clearable
 */
func (a *TransferControl) Clearable(value interface{}) *TransferControl {
    a.Set("clearable", value)
    return a
}

/**
 * options
 */
func (a *TransferControl) Options(value interface{}) *TransferControl {
    a.Set("options", value)
    return a
}

/**
 * multiple
 */
func (a *TransferControl) Multiple(value interface{}) *TransferControl {
    a.Set("multiple", value)
    return a
}

/**
 * checkAll
 */
func (a *TransferControl) CheckAll(value interface{}) *TransferControl {
    a.Set("checkAll", value)
    return a
}

/**
 * joinValues
 */
func (a *TransferControl) JoinValues(value interface{}) *TransferControl {
    a.Set("joinValues", value)
    return a
}
