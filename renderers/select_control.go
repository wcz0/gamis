package renderers


/**
 * Select 下拉选择框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/select

 * @author wcz0
 * @version 6.2.2
 */
type SelectControl struct {
	*BaseRenderer
}

func NewSelectControl() *SelectControl {
    a := &SelectControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "select")
    return a
}


func (a *SelectControl) Set(name string, value interface{}) *SelectControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * labelClassName
 */
func (a *SelectControl) LabelClassName(value interface{}) *SelectControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * initAutoFill
 */
func (a *SelectControl) InitAutoFill(value interface{}) *SelectControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * row
 */
func (a *SelectControl) Row(value interface{}) *SelectControl {
    a.Set("row", value)
    return a
}

/**
 * searchApi
 */
func (a *SelectControl) SearchApi(value interface{}) *SelectControl {
    a.Set("searchApi", value)
    return a
}

/**
 * loadingConfig
 */
func (a *SelectControl) LoadingConfig(value interface{}) *SelectControl {
    a.Set("loadingConfig", value)
    return a
}

/**
 * options
 */
func (a *SelectControl) Options(value interface{}) *SelectControl {
    a.Set("options", value)
    return a
}

/**
 * joinValues
 */
func (a *SelectControl) JoinValues(value interface{}) *SelectControl {
    a.Set("joinValues", value)
    return a
}

/**
 * useMobileUI
 */
func (a *SelectControl) UseMobileUI(value interface{}) *SelectControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * source
 */
func (a *SelectControl) Source(value interface{}) *SelectControl {
    a.Set("source", value)
    return a
}

/**
 * delimiter
 */
func (a *SelectControl) Delimiter(value interface{}) *SelectControl {
    a.Set("delimiter", value)
    return a
}

/**
 * creatable
 */
func (a *SelectControl) Creatable(value interface{}) *SelectControl {
    a.Set("creatable", value)
    return a
}

/**
 * label
 */
func (a *SelectControl) Label(value interface{}) *SelectControl {
    a.Set("label", value)
    return a
}

/**
 * labelAlign
 */
func (a *SelectControl) LabelAlign(value interface{}) *SelectControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * remark
 */
func (a *SelectControl) Remark(value interface{}) *SelectControl {
    a.Set("remark", value)
    return a
}

/**
 * labelRemark
 */
func (a *SelectControl) LabelRemark(value interface{}) *SelectControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * visibleOn
 */
func (a *SelectControl) VisibleOn(value interface{}) *SelectControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * addDialog
 */
func (a *SelectControl) AddDialog(value interface{}) *SelectControl {
    a.Set("addDialog", value)
    return a
}

/**
 * validationErrors
 */
func (a *SelectControl) ValidationErrors(value interface{}) *SelectControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * value
 */
func (a *SelectControl) Value(value interface{}) *SelectControl {
    a.Set("value", value)
    return a
}

/**
 * static
 */
func (a *SelectControl) Static(value interface{}) *SelectControl {
    a.Set("static", value)
    return a
}

/**
 * editorSetting
 */
func (a *SelectControl) EditorSetting(value interface{}) *SelectControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * menuTpl
 */
func (a *SelectControl) MenuTpl(value interface{}) *SelectControl {
    a.Set("menuTpl", value)
    return a
}

/**
 * searchResultMode
 */
func (a *SelectControl) SearchResultMode(value interface{}) *SelectControl {
    a.Set("searchResultMode", value)
    return a
}

/**
 * checkAllLabel
 */
func (a *SelectControl) CheckAllLabel(value interface{}) *SelectControl {
    a.Set("checkAllLabel", value)
    return a
}

/**
 * editable
 */
func (a *SelectControl) Editable(value interface{}) *SelectControl {
    a.Set("editable", value)
    return a
}

/**
 * required
 */
func (a *SelectControl) Required(value interface{}) *SelectControl {
    a.Set("required", value)
    return a
}

/**
 * staticSchema
 */
func (a *SelectControl) StaticSchema(value interface{}) *SelectControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * testIdBuilder
 */
func (a *SelectControl) TestIdBuilder(value interface{}) *SelectControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * removable
 */
func (a *SelectControl) Removable(value interface{}) *SelectControl {
    a.Set("removable", value)
    return a
}

/**
 * clearable
 */
func (a *SelectControl) Clearable(value interface{}) *SelectControl {
    a.Set("clearable", value)
    return a
}

/**
 * deferField
 */
func (a *SelectControl) DeferField(value interface{}) *SelectControl {
    a.Set("deferField", value)
    return a
}

/**
 * submitOnChange
 */
func (a *SelectControl) SubmitOnChange(value interface{}) *SelectControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * disabled
 */
func (a *SelectControl) Disabled(value interface{}) *SelectControl {
    a.Set("disabled", value)
    return a
}

/**
 * id
 */
func (a *SelectControl) Id(value interface{}) *SelectControl {
    a.Set("id", value)
    return a
}

/**
 * borderMode
 */
func (a *SelectControl) BorderMode(value interface{}) *SelectControl {
    a.Set("borderMode", value)
    return a
}

/**
 * selectMode
 */
func (a *SelectControl) SelectMode(value interface{}) *SelectControl {
    a.Set("selectMode", value)
    return a
}

/**
 * multiple
 */
func (a *SelectControl) Multiple(value interface{}) *SelectControl {
    a.Set("multiple", value)
    return a
}

/**
 * readOnly
 */
func (a *SelectControl) ReadOnly(value interface{}) *SelectControl {
    a.Set("readOnly", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *SelectControl) ReadOnlyOn(value interface{}) *SelectControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * visible
 */
func (a *SelectControl) Visible(value interface{}) *SelectControl {
    a.Set("visible", value)
    return a
}

/**
 * leftOptions
 */
func (a *SelectControl) LeftOptions(value interface{}) *SelectControl {
    a.Set("leftOptions", value)
    return a
}

/**
 * overflowTagPopover
 */
func (a *SelectControl) OverflowTagPopover(value interface{}) *SelectControl {
    a.Set("overflowTagPopover", value)
    return a
}

/**
 * checkAll
 */
func (a *SelectControl) CheckAll(value interface{}) *SelectControl {
    a.Set("checkAll", value)
    return a
}

/**
 * editApi
 */
func (a *SelectControl) EditApi(value interface{}) *SelectControl {
    a.Set("editApi", value)
    return a
}

/**
 * labelWidth
 */
func (a *SelectControl) LabelWidth(value interface{}) *SelectControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * name
 */
func (a *SelectControl) Name(value interface{}) *SelectControl {
    a.Set("name", value)
    return a
}

/**
 * description
 */
func (a *SelectControl) Description(value interface{}) *SelectControl {
    a.Set("description", value)
    return a
}

/**
 * mode
 */
func (a *SelectControl) Mode(value interface{}) *SelectControl {
    a.Set("mode", value)
    return a
}

/**
 * horizontal
 */
func (a *SelectControl) Horizontal(value interface{}) *SelectControl {
    a.Set("horizontal", value)
    return a
}

/**
 * showInvalidMatch
 */
func (a *SelectControl) ShowInvalidMatch(value interface{}) *SelectControl {
    a.Set("showInvalidMatch", value)
    return a
}

/**
 * inline
 */
func (a *SelectControl) Inline(value interface{}) *SelectControl {
    a.Set("inline", value)
    return a
}

/**
 * size
 */
func (a *SelectControl) Size(value interface{}) *SelectControl {
    a.Set("size", value)
    return a
}

/**
 * columns
 */
func (a *SelectControl) Columns(value interface{}) *SelectControl {
    a.Set("columns", value)
    return a
}

/**
 * searchResultColumns
 */
func (a *SelectControl) SearchResultColumns(value interface{}) *SelectControl {
    a.Set("searchResultColumns", value)
    return a
}

/**
 * defaultCheckAll
 */
func (a *SelectControl) DefaultCheckAll(value interface{}) *SelectControl {
    a.Set("defaultCheckAll", value)
    return a
}

/**
 * maxTagCount
 */
func (a *SelectControl) MaxTagCount(value interface{}) *SelectControl {
    a.Set("maxTagCount", value)
    return a
}

/**
 * optionClassName
 */
func (a *SelectControl) OptionClassName(value interface{}) *SelectControl {
    a.Set("optionClassName", value)
    return a
}

/**
 * deleteApi
 */
func (a *SelectControl) DeleteApi(value interface{}) *SelectControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * initFetch
 */
func (a *SelectControl) InitFetch(value interface{}) *SelectControl {
    a.Set("initFetch", value)
    return a
}

/**
 * valuesNoWrap
 */
func (a *SelectControl) ValuesNoWrap(value interface{}) *SelectControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * clearValueOnSourceChange
 */
func (a *SelectControl) ClearValueOnSourceChange(value interface{}) *SelectControl {
    a.Set("clearValueOnSourceChange", value)
    return a
}

/**
 * desc
 */
func (a *SelectControl) Desc(value interface{}) *SelectControl {
    a.Set("desc", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *SelectControl) StaticPlaceholder(value interface{}) *SelectControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 可选值: select | multi-select
 */
func (a *SelectControl) Type(value interface{}) *SelectControl {
    a.Set("type", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *SelectControl) Width(value interface{}) *SelectControl {
    a.Set("width", value)
    return a
}

/**
 * selectFirst
 */
func (a *SelectControl) SelectFirst(value interface{}) *SelectControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * deferApi
 */
func (a *SelectControl) DeferApi(value interface{}) *SelectControl {
    a.Set("deferApi", value)
    return a
}

/**
 * addControls
 */
func (a *SelectControl) AddControls(value interface{}) *SelectControl {
    a.Set("addControls", value)
    return a
}

/**
 * hint
 */
func (a *SelectControl) Hint(value interface{}) *SelectControl {
    a.Set("hint", value)
    return a
}

/**
 * inputClassName
 */
func (a *SelectControl) InputClassName(value interface{}) *SelectControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *SelectControl) ClearValueOnHidden(value interface{}) *SelectControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * hiddenOn
 */
func (a *SelectControl) HiddenOn(value interface{}) *SelectControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticOn
 */
func (a *SelectControl) StaticOn(value interface{}) *SelectControl {
    a.Set("staticOn", value)
    return a
}

/**
 * addApi
 */
func (a *SelectControl) AddApi(value interface{}) *SelectControl {
    a.Set("addApi", value)
    return a
}

/**
 * extraName
 */
func (a *SelectControl) ExtraName(value interface{}) *SelectControl {
    a.Set("extraName", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *SelectControl) DescriptionClassName(value interface{}) *SelectControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * onEvent
 */
func (a *SelectControl) OnEvent(value interface{}) *SelectControl {
    a.Set("onEvent", value)
    return a
}

/**
 * rightMode
 */
func (a *SelectControl) RightMode(value interface{}) *SelectControl {
    a.Set("rightMode", value)
    return a
}

/**
 * overlay
 */
func (a *SelectControl) Overlay(value interface{}) *SelectControl {
    a.Set("overlay", value)
    return a
}

/**
 * 是否自动选中子节点
 */
func (a *SelectControl) AutoCheckChildren(value interface{}) *SelectControl {
    a.Set("autoCheckChildren", value)
    return a
}

/**
 * resetValue
 */
func (a *SelectControl) ResetValue(value interface{}) *SelectControl {
    a.Set("resetValue", value)
    return a
}

/**
 * deleteConfirmText
 */
func (a *SelectControl) DeleteConfirmText(value interface{}) *SelectControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * placeholder
 */
func (a *SelectControl) Placeholder(value interface{}) *SelectControl {
    a.Set("placeholder", value)
    return a
}

/**
 * autoFill
 */
func (a *SelectControl) AutoFill(value interface{}) *SelectControl {
    a.Set("autoFill", value)
    return a
}

/**
 * hidden
 */
func (a *SelectControl) Hidden(value interface{}) *SelectControl {
    a.Set("hidden", value)
    return a
}

/**
 * staticClassName
 */
func (a *SelectControl) StaticClassName(value interface{}) *SelectControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *SelectControl) StaticLabelClassName(value interface{}) *SelectControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * itemHeight
 */
func (a *SelectControl) ItemHeight(value interface{}) *SelectControl {
    a.Set("itemHeight", value)
    return a
}

/**
 * editControls
 */
func (a *SelectControl) EditControls(value interface{}) *SelectControl {
    a.Set("editControls", value)
    return a
}

/**
 * editDialog
 */
func (a *SelectControl) EditDialog(value interface{}) *SelectControl {
    a.Set("editDialog", value)
    return a
}

/**
 * createBtnLabel
 */
func (a *SelectControl) CreateBtnLabel(value interface{}) *SelectControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * validations
 */
func (a *SelectControl) Validations(value interface{}) *SelectControl {
    a.Set("validations", value)
    return a
}

/**
 * validateApi
 */
func (a *SelectControl) ValidateApi(value interface{}) *SelectControl {
    a.Set("validateApi", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *SelectControl) StaticInputClassName(value interface{}) *SelectControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 设置label字段
 */
func (a *SelectControl) LabelField(value interface{}) *SelectControl {
    a.Set("labelField", value)
    return a
}

/**
 * className
 */
func (a *SelectControl) ClassName(value interface{}) *SelectControl {
    a.Set("className", value)
    return a
}

/**
 * style
 */
func (a *SelectControl) Style(value interface{}) *SelectControl {
    a.Set("style", value)
    return a
}

/**
 * 是否可搜索
 */
func (a *SelectControl) Searchable(value interface{}) *SelectControl {
    a.Set("searchable", value)
    return a
}

/**
 * virtualThreshold
 */
func (a *SelectControl) VirtualThreshold(value interface{}) *SelectControl {
    a.Set("virtualThreshold", value)
    return a
}

/**
 * 设置value字段
 */
func (a *SelectControl) ValueField(value interface{}) *SelectControl {
    a.Set("valueField", value)
    return a
}

/**
 * extractValue
 */
func (a *SelectControl) ExtractValue(value interface{}) *SelectControl {
    a.Set("extractValue", value)
    return a
}

/**
 * validateOnChange
 */
func (a *SelectControl) ValidateOnChange(value interface{}) *SelectControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * disabledOn
 */
func (a *SelectControl) DisabledOn(value interface{}) *SelectControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * autoComplete
 */
func (a *SelectControl) AutoComplete(value interface{}) *SelectControl {
    a.Set("autoComplete", value)
    return a
}

/**
 * leftMode
 */
func (a *SelectControl) LeftMode(value interface{}) *SelectControl {
    a.Set("leftMode", value)
    return a
}

/**
 * initFetchOn
 */
func (a *SelectControl) InitFetchOn(value interface{}) *SelectControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * labelOverflow
 */
func (a *SelectControl) LabelOverflow(value interface{}) *SelectControl {
    a.Set("labelOverflow", value)
    return a
}
