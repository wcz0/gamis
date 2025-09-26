package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type BaseTransferControl struct {
	*BaseRenderer
}

func NewBaseTransferControl() *BaseTransferControl {
    a := &BaseTransferControl{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}


func (a *BaseTransferControl) Set(name string, value interface{}) *BaseTransferControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * options
 */
func (a *BaseTransferControl) Options(value interface{}) *BaseTransferControl {
    a.Set("options", value)
    return a
}

/**
 * delimiter
 */
func (a *BaseTransferControl) Delimiter(value interface{}) *BaseTransferControl {
    a.Set("delimiter", value)
    return a
}

/**
 * editApi
 */
func (a *BaseTransferControl) EditApi(value interface{}) *BaseTransferControl {
    a.Set("editApi", value)
    return a
}

/**
 * deleteConfirmText
 */
func (a *BaseTransferControl) DeleteConfirmText(value interface{}) *BaseTransferControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * searchPlaceholder
 */
func (a *BaseTransferControl) SearchPlaceholder(value interface{}) *BaseTransferControl {
    a.Set("searchPlaceholder", value)
    return a
}

/**
 * creatable
 */
func (a *BaseTransferControl) Creatable(value interface{}) *BaseTransferControl {
    a.Set("creatable", value)
    return a
}

/**
 * labelAlign
 */
func (a *BaseTransferControl) LabelAlign(value interface{}) *BaseTransferControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * extraName
 */
func (a *BaseTransferControl) ExtraName(value interface{}) *BaseTransferControl {
    a.Set("extraName", value)
    return a
}

/**
 * submitOnChange
 */
func (a *BaseTransferControl) SubmitOnChange(value interface{}) *BaseTransferControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * description
 */
func (a *BaseTransferControl) Description(value interface{}) *BaseTransferControl {
    a.Set("description", value)
    return a
}

/**
 * static
 */
func (a *BaseTransferControl) Static(value interface{}) *BaseTransferControl {
    a.Set("static", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *BaseTransferControl) StaticPlaceholder(value interface{}) *BaseTransferControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * searchable
 */
func (a *BaseTransferControl) Searchable(value interface{}) *BaseTransferControl {
    a.Set("searchable", value)
    return a
}

/**
 * loadingConfig
 */
func (a *BaseTransferControl) LoadingConfig(value interface{}) *BaseTransferControl {
    a.Set("loadingConfig", value)
    return a
}

/**
 * label
 */
func (a *BaseTransferControl) Label(value interface{}) *BaseTransferControl {
    a.Set("label", value)
    return a
}

/**
 * labelRemark
 */
func (a *BaseTransferControl) LabelRemark(value interface{}) *BaseTransferControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * horizontal
 */
func (a *BaseTransferControl) Horizontal(value interface{}) *BaseTransferControl {
    a.Set("horizontal", value)
    return a
}

/**
 * resultTitle
 */
func (a *BaseTransferControl) ResultTitle(value interface{}) *BaseTransferControl {
    a.Set("resultTitle", value)
    return a
}

/**
 * valueTpl
 */
func (a *BaseTransferControl) ValueTpl(value interface{}) *BaseTransferControl {
    a.Set("valueTpl", value)
    return a
}

/**
 * virtualThreshold
 */
func (a *BaseTransferControl) VirtualThreshold(value interface{}) *BaseTransferControl {
    a.Set("virtualThreshold", value)
    return a
}

/**
 * onlyChildren
 */
func (a *BaseTransferControl) OnlyChildren(value interface{}) *BaseTransferControl {
    a.Set("onlyChildren", value)
    return a
}

/**
 * value
 */
func (a *BaseTransferControl) Value(value interface{}) *BaseTransferControl {
    a.Set("value", value)
    return a
}

/**
 * row
 */
func (a *BaseTransferControl) Row(value interface{}) *BaseTransferControl {
    a.Set("row", value)
    return a
}

/**
 * style
 */
func (a *BaseTransferControl) Style(value interface{}) *BaseTransferControl {
    a.Set("style", value)
    return a
}

/**
 * sortable
 */
func (a *BaseTransferControl) Sortable(value interface{}) *BaseTransferControl {
    a.Set("sortable", value)
    return a
}

/**
 * searchResultMode
 */
func (a *BaseTransferControl) SearchResultMode(value interface{}) *BaseTransferControl {
    a.Set("searchResultMode", value)
    return a
}

/**
 * menuTpl
 */
func (a *BaseTransferControl) MenuTpl(value interface{}) *BaseTransferControl {
    a.Set("menuTpl", value)
    return a
}

/**
 * initiallyOpen
 */
func (a *BaseTransferControl) InitiallyOpen(value interface{}) *BaseTransferControl {
    a.Set("initiallyOpen", value)
    return a
}

/**
 * source
 */
func (a *BaseTransferControl) Source(value interface{}) *BaseTransferControl {
    a.Set("source", value)
    return a
}

/**
 * valuesNoWrap
 */
func (a *BaseTransferControl) ValuesNoWrap(value interface{}) *BaseTransferControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * deleteApi
 */
func (a *BaseTransferControl) DeleteApi(value interface{}) *BaseTransferControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * id
 */
func (a *BaseTransferControl) Id(value interface{}) *BaseTransferControl {
    a.Set("id", value)
    return a
}

/**
 * staticOn
 */
func (a *BaseTransferControl) StaticOn(value interface{}) *BaseTransferControl {
    a.Set("staticOn", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *BaseTransferControl) StaticLabelClassName(value interface{}) *BaseTransferControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * resultSearchPlaceholder
 */
func (a *BaseTransferControl) ResultSearchPlaceholder(value interface{}) *BaseTransferControl {
    a.Set("resultSearchPlaceholder", value)
    return a
}

/**
 * initFetchOn
 */
func (a *BaseTransferControl) InitFetchOn(value interface{}) *BaseTransferControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * multiple
 */
func (a *BaseTransferControl) Multiple(value interface{}) *BaseTransferControl {
    a.Set("multiple", value)
    return a
}

/**
 * joinValues
 */
func (a *BaseTransferControl) JoinValues(value interface{}) *BaseTransferControl {
    a.Set("joinValues", value)
    return a
}

/**
 * editable
 */
func (a *BaseTransferControl) Editable(value interface{}) *BaseTransferControl {
    a.Set("editable", value)
    return a
}

/**
 * initAutoFill
 */
func (a *BaseTransferControl) InitAutoFill(value interface{}) *BaseTransferControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * hiddenOn
 */
func (a *BaseTransferControl) HiddenOn(value interface{}) *BaseTransferControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *BaseTransferControl) StaticClassName(value interface{}) *BaseTransferControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * deferField
 */
func (a *BaseTransferControl) DeferField(value interface{}) *BaseTransferControl {
    a.Set("deferField", value)
    return a
}

/**
 * addDialog
 */
func (a *BaseTransferControl) AddDialog(value interface{}) *BaseTransferControl {
    a.Set("addDialog", value)
    return a
}

/**
 * removable
 */
func (a *BaseTransferControl) Removable(value interface{}) *BaseTransferControl {
    a.Set("removable", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *BaseTransferControl) DescriptionClassName(value interface{}) *BaseTransferControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * inline
 */
func (a *BaseTransferControl) Inline(value interface{}) *BaseTransferControl {
    a.Set("inline", value)
    return a
}

/**
 * inputClassName
 */
func (a *BaseTransferControl) InputClassName(value interface{}) *BaseTransferControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * useMobileUI
 */
func (a *BaseTransferControl) UseMobileUI(value interface{}) *BaseTransferControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * selectMode
 */
func (a *BaseTransferControl) SelectMode(value interface{}) *BaseTransferControl {
    a.Set("selectMode", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *BaseTransferControl) ClearValueOnHidden(value interface{}) *BaseTransferControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * labelWidth
 */
func (a *BaseTransferControl) LabelWidth(value interface{}) *BaseTransferControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * mode
 */
func (a *BaseTransferControl) Mode(value interface{}) *BaseTransferControl {
    a.Set("mode", value)
    return a
}

/**
 * validations
 */
func (a *BaseTransferControl) Validations(value interface{}) *BaseTransferControl {
    a.Set("validations", value)
    return a
}

/**
 * staticSchema
 */
func (a *BaseTransferControl) StaticSchema(value interface{}) *BaseTransferControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * showArrow
 */
func (a *BaseTransferControl) ShowArrow(value interface{}) *BaseTransferControl {
    a.Set("showArrow", value)
    return a
}

/**
 * resultSearchable
 */
func (a *BaseTransferControl) ResultSearchable(value interface{}) *BaseTransferControl {
    a.Set("resultSearchable", value)
    return a
}

/**
 * searchApi
 */
func (a *BaseTransferControl) SearchApi(value interface{}) *BaseTransferControl {
    a.Set("searchApi", value)
    return a
}

/**
 * extractValue
 */
func (a *BaseTransferControl) ExtractValue(value interface{}) *BaseTransferControl {
    a.Set("extractValue", value)
    return a
}

/**
 * addControls
 */
func (a *BaseTransferControl) AddControls(value interface{}) *BaseTransferControl {
    a.Set("addControls", value)
    return a
}

/**
 * editDialog
 */
func (a *BaseTransferControl) EditDialog(value interface{}) *BaseTransferControl {
    a.Set("editDialog", value)
    return a
}

/**
 * name
 */
func (a *BaseTransferControl) Name(value interface{}) *BaseTransferControl {
    a.Set("name", value)
    return a
}

/**
 * desc
 */
func (a *BaseTransferControl) Desc(value interface{}) *BaseTransferControl {
    a.Set("desc", value)
    return a
}

/**
 * required
 */
func (a *BaseTransferControl) Required(value interface{}) *BaseTransferControl {
    a.Set("required", value)
    return a
}

/**
 * visibleOn
 */
func (a *BaseTransferControl) VisibleOn(value interface{}) *BaseTransferControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * editorSetting
 */
func (a *BaseTransferControl) EditorSetting(value interface{}) *BaseTransferControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * checkAll
 */
func (a *BaseTransferControl) CheckAll(value interface{}) *BaseTransferControl {
    a.Set("checkAll", value)
    return a
}

/**
 * labelOverflow
 */
func (a *BaseTransferControl) LabelOverflow(value interface{}) *BaseTransferControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * hint
 */
func (a *BaseTransferControl) Hint(value interface{}) *BaseTransferControl {
    a.Set("hint", value)
    return a
}

/**
 * validateOnChange
 */
func (a *BaseTransferControl) ValidateOnChange(value interface{}) *BaseTransferControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * resultListModeFollowSelect
 */
func (a *BaseTransferControl) ResultListModeFollowSelect(value interface{}) *BaseTransferControl {
    a.Set("resultListModeFollowSelect", value)
    return a
}

/**
 * leftOptions
 */
func (a *BaseTransferControl) LeftOptions(value interface{}) *BaseTransferControl {
    a.Set("leftOptions", value)
    return a
}

/**
 * statistics
 */
func (a *BaseTransferControl) Statistics(value interface{}) *BaseTransferControl {
    a.Set("statistics", value)
    return a
}

/**
 * showInvalidMatch
 */
func (a *BaseTransferControl) ShowInvalidMatch(value interface{}) *BaseTransferControl {
    a.Set("showInvalidMatch", value)
    return a
}

/**
 * addApi
 */
func (a *BaseTransferControl) AddApi(value interface{}) *BaseTransferControl {
    a.Set("addApi", value)
    return a
}

/**
 * editControls
 */
func (a *BaseTransferControl) EditControls(value interface{}) *BaseTransferControl {
    a.Set("editControls", value)
    return a
}

/**
 * readOnly
 */
func (a *BaseTransferControl) ReadOnly(value interface{}) *BaseTransferControl {
    a.Set("readOnly", value)
    return a
}

/**
 * validationErrors
 */
func (a *BaseTransferControl) ValidationErrors(value interface{}) *BaseTransferControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * validateApi
 */
func (a *BaseTransferControl) ValidateApi(value interface{}) *BaseTransferControl {
    a.Set("validateApi", value)
    return a
}

/**
 * className
 */
func (a *BaseTransferControl) ClassName(value interface{}) *BaseTransferControl {
    a.Set("className", value)
    return a
}

/**
 * disabledOn
 */
func (a *BaseTransferControl) DisabledOn(value interface{}) *BaseTransferControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * onEvent
 */
func (a *BaseTransferControl) OnEvent(value interface{}) *BaseTransferControl {
    a.Set("onEvent", value)
    return a
}

/**
 * deferApi
 */
func (a *BaseTransferControl) DeferApi(value interface{}) *BaseTransferControl {
    a.Set("deferApi", value)
    return a
}

/**
 * clearValueOnSourceChange
 */
func (a *BaseTransferControl) ClearValueOnSourceChange(value interface{}) *BaseTransferControl {
    a.Set("clearValueOnSourceChange", value)
    return a
}

/**
 * disabled
 */
func (a *BaseTransferControl) Disabled(value interface{}) *BaseTransferControl {
    a.Set("disabled", value)
    return a
}

/**
 * hidden
 */
func (a *BaseTransferControl) Hidden(value interface{}) *BaseTransferControl {
    a.Set("hidden", value)
    return a
}

/**
 * size
 */
func (a *BaseTransferControl) Size(value interface{}) *BaseTransferControl {
    a.Set("size", value)
    return a
}

/**
 * type
 */
func (a *BaseTransferControl) Type(value interface{}) *BaseTransferControl {
    a.Set("type", value)
    return a
}

/**
 * leftMode
 */
func (a *BaseTransferControl) LeftMode(value interface{}) *BaseTransferControl {
    a.Set("leftMode", value)
    return a
}

/**
 * itemHeight
 */
func (a *BaseTransferControl) ItemHeight(value interface{}) *BaseTransferControl {
    a.Set("itemHeight", value)
    return a
}

/**
 * initFetch
 */
func (a *BaseTransferControl) InitFetch(value interface{}) *BaseTransferControl {
    a.Set("initFetch", value)
    return a
}

/**
 * autoFill
 */
func (a *BaseTransferControl) AutoFill(value interface{}) *BaseTransferControl {
    a.Set("autoFill", value)
    return a
}

/**
 * selectTitle
 */
func (a *BaseTransferControl) SelectTitle(value interface{}) *BaseTransferControl {
    a.Set("selectTitle", value)
    return a
}

/**
 * autoCheckChildren
 */
func (a *BaseTransferControl) AutoCheckChildren(value interface{}) *BaseTransferControl {
    a.Set("autoCheckChildren", value)
    return a
}

/**
 * pagination
 */
func (a *BaseTransferControl) Pagination(value interface{}) *BaseTransferControl {
    a.Set("pagination", value)
    return a
}

/**
 * createBtnLabel
 */
func (a *BaseTransferControl) CreateBtnLabel(value interface{}) *BaseTransferControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *BaseTransferControl) StaticInputClassName(value interface{}) *BaseTransferControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *BaseTransferControl) Width(value interface{}) *BaseTransferControl {
    a.Set("width", value)
    return a
}

/**
 * selectFirst
 */
func (a *BaseTransferControl) SelectFirst(value interface{}) *BaseTransferControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * clearable
 */
func (a *BaseTransferControl) Clearable(value interface{}) *BaseTransferControl {
    a.Set("clearable", value)
    return a
}

/**
 * labelClassName
 */
func (a *BaseTransferControl) LabelClassName(value interface{}) *BaseTransferControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * remark
 */
func (a *BaseTransferControl) Remark(value interface{}) *BaseTransferControl {
    a.Set("remark", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *BaseTransferControl) ReadOnlyOn(value interface{}) *BaseTransferControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * placeholder
 */
func (a *BaseTransferControl) Placeholder(value interface{}) *BaseTransferControl {
    a.Set("placeholder", value)
    return a
}

/**
 * visible
 */
func (a *BaseTransferControl) Visible(value interface{}) *BaseTransferControl {
    a.Set("visible", value)
    return a
}

/**
 * resetValue
 */
func (a *BaseTransferControl) ResetValue(value interface{}) *BaseTransferControl {
    a.Set("resetValue", value)
    return a
}

/**
 * rightMode
 */
func (a *BaseTransferControl) RightMode(value interface{}) *BaseTransferControl {
    a.Set("rightMode", value)
    return a
}

/**
 * columns
 */
func (a *BaseTransferControl) Columns(value interface{}) *BaseTransferControl {
    a.Set("columns", value)
    return a
}

/**
 * searchResultColumns
 */
func (a *BaseTransferControl) SearchResultColumns(value interface{}) *BaseTransferControl {
    a.Set("searchResultColumns", value)
    return a
}
