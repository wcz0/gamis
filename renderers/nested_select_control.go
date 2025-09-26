package renderers


/**
 * Nested Select 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/nested-select

 * @author wcz0
 * @version 6.2.2
 */
type NestedSelectControl struct {
	*BaseRenderer
}

func NewNestedSelectControl() *NestedSelectControl {
    a := &NestedSelectControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "nested-select")
    return a
}


func (a *NestedSelectControl) Set(name string, value interface{}) *NestedSelectControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * labelOverflow
 */
func (a *NestedSelectControl) LabelOverflow(value interface{}) *NestedSelectControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * inline
 */
func (a *NestedSelectControl) Inline(value interface{}) *NestedSelectControl {
    a.Set("inline", value)
    return a
}

/**
 * validateApi
 */
func (a *NestedSelectControl) ValidateApi(value interface{}) *NestedSelectControl {
    a.Set("validateApi", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *NestedSelectControl) ClearValueOnHidden(value interface{}) *NestedSelectControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * hideNodePathLabel
 */
func (a *NestedSelectControl) HideNodePathLabel(value interface{}) *NestedSelectControl {
    a.Set("hideNodePathLabel", value)
    return a
}

/**
 * joinValues
 */
func (a *NestedSelectControl) JoinValues(value interface{}) *NestedSelectControl {
    a.Set("joinValues", value)
    return a
}

/**
 * labelAlign
 */
func (a *NestedSelectControl) LabelAlign(value interface{}) *NestedSelectControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * readOnly
 */
func (a *NestedSelectControl) ReadOnly(value interface{}) *NestedSelectControl {
    a.Set("readOnly", value)
    return a
}

/**
 * id
 */
func (a *NestedSelectControl) Id(value interface{}) *NestedSelectControl {
    a.Set("id", value)
    return a
}

/**
 * staticOn
 */
func (a *NestedSelectControl) StaticOn(value interface{}) *NestedSelectControl {
    a.Set("staticOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *NestedSelectControl) StaticPlaceholder(value interface{}) *NestedSelectControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * deferField
 */
func (a *NestedSelectControl) DeferField(value interface{}) *NestedSelectControl {
    a.Set("deferField", value)
    return a
}

/**
 * inputClassName
 */
func (a *NestedSelectControl) InputClassName(value interface{}) *NestedSelectControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * visible
 */
func (a *NestedSelectControl) Visible(value interface{}) *NestedSelectControl {
    a.Set("visible", value)
    return a
}

/**
 * editorSetting
 */
func (a *NestedSelectControl) EditorSetting(value interface{}) *NestedSelectControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * borderMode
 */
func (a *NestedSelectControl) BorderMode(value interface{}) *NestedSelectControl {
    a.Set("borderMode", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *NestedSelectControl) DescriptionClassName(value interface{}) *NestedSelectControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * onEvent
 */
func (a *NestedSelectControl) OnEvent(value interface{}) *NestedSelectControl {
    a.Set("onEvent", value)
    return a
}

/**
 * static
 */
func (a *NestedSelectControl) Static(value interface{}) *NestedSelectControl {
    a.Set("static", value)
    return a
}

/**
 * extractValue
 */
func (a *NestedSelectControl) ExtractValue(value interface{}) *NestedSelectControl {
    a.Set("extractValue", value)
    return a
}

/**
 * placeholder
 */
func (a *NestedSelectControl) Placeholder(value interface{}) *NestedSelectControl {
    a.Set("placeholder", value)
    return a
}

/**
 * overflowTagPopover
 */
func (a *NestedSelectControl) OverflowTagPopover(value interface{}) *NestedSelectControl {
    a.Set("overflowTagPopover", value)
    return a
}

/**
 * checkAll
 */
func (a *NestedSelectControl) CheckAll(value interface{}) *NestedSelectControl {
    a.Set("checkAll", value)
    return a
}

/**
 * deferApi
 */
func (a *NestedSelectControl) DeferApi(value interface{}) *NestedSelectControl {
    a.Set("deferApi", value)
    return a
}

/**
 * editApi
 */
func (a *NestedSelectControl) EditApi(value interface{}) *NestedSelectControl {
    a.Set("editApi", value)
    return a
}

/**
 * editDialog
 */
func (a *NestedSelectControl) EditDialog(value interface{}) *NestedSelectControl {
    a.Set("editDialog", value)
    return a
}

/**
 * removable
 */
func (a *NestedSelectControl) Removable(value interface{}) *NestedSelectControl {
    a.Set("removable", value)
    return a
}

/**
 * remark
 */
func (a *NestedSelectControl) Remark(value interface{}) *NestedSelectControl {
    a.Set("remark", value)
    return a
}

/**
 * className
 */
func (a *NestedSelectControl) ClassName(value interface{}) *NestedSelectControl {
    a.Set("className", value)
    return a
}

/**
 * initFetchOn
 */
func (a *NestedSelectControl) InitFetchOn(value interface{}) *NestedSelectControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * delimiter
 */
func (a *NestedSelectControl) Delimiter(value interface{}) *NestedSelectControl {
    a.Set("delimiter", value)
    return a
}

/**
 * creatable
 */
func (a *NestedSelectControl) Creatable(value interface{}) *NestedSelectControl {
    a.Set("creatable", value)
    return a
}

/**
 * editControls
 */
func (a *NestedSelectControl) EditControls(value interface{}) *NestedSelectControl {
    a.Set("editControls", value)
    return a
}

/**
 * validations
 */
func (a *NestedSelectControl) Validations(value interface{}) *NestedSelectControl {
    a.Set("validations", value)
    return a
}

/**
 * disabled
 */
func (a *NestedSelectControl) Disabled(value interface{}) *NestedSelectControl {
    a.Set("disabled", value)
    return a
}

/**
 * hidden
 */
func (a *NestedSelectControl) Hidden(value interface{}) *NestedSelectControl {
    a.Set("hidden", value)
    return a
}

/**
 * hiddenOn
 */
func (a *NestedSelectControl) HiddenOn(value interface{}) *NestedSelectControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * source
 */
func (a *NestedSelectControl) Source(value interface{}) *NestedSelectControl {
    a.Set("source", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *NestedSelectControl) StaticLabelClassName(value interface{}) *NestedSelectControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *NestedSelectControl) StaticInputClassName(value interface{}) *NestedSelectControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * clearable
 */
func (a *NestedSelectControl) Clearable(value interface{}) *NestedSelectControl {
    a.Set("clearable", value)
    return a
}

/**
 * editable
 */
func (a *NestedSelectControl) Editable(value interface{}) *NestedSelectControl {
    a.Set("editable", value)
    return a
}

/**
 * value
 */
func (a *NestedSelectControl) Value(value interface{}) *NestedSelectControl {
    a.Set("value", value)
    return a
}

/**
 * autoFill
 */
func (a *NestedSelectControl) AutoFill(value interface{}) *NestedSelectControl {
    a.Set("autoFill", value)
    return a
}

/**
 * row
 */
func (a *NestedSelectControl) Row(value interface{}) *NestedSelectControl {
    a.Set("row", value)
    return a
}

/**
 * staticClassName
 */
func (a *NestedSelectControl) StaticClassName(value interface{}) *NestedSelectControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * style
 */
func (a *NestedSelectControl) Style(value interface{}) *NestedSelectControl {
    a.Set("style", value)
    return a
}

/**
 * resetValue
 */
func (a *NestedSelectControl) ResetValue(value interface{}) *NestedSelectControl {
    a.Set("resetValue", value)
    return a
}

/**
 * clearValueOnSourceChange
 */
func (a *NestedSelectControl) ClearValueOnSourceChange(value interface{}) *NestedSelectControl {
    a.Set("clearValueOnSourceChange", value)
    return a
}

/**
 * label
 */
func (a *NestedSelectControl) Label(value interface{}) *NestedSelectControl {
    a.Set("label", value)
    return a
}

/**
 * name
 */
func (a *NestedSelectControl) Name(value interface{}) *NestedSelectControl {
    a.Set("name", value)
    return a
}

/**
 * labelRemark
 */
func (a *NestedSelectControl) LabelRemark(value interface{}) *NestedSelectControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * mode
 */
func (a *NestedSelectControl) Mode(value interface{}) *NestedSelectControl {
    a.Set("mode", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *NestedSelectControl) Width(value interface{}) *NestedSelectControl {
    a.Set("width", value)
    return a
}

/**
 * required
 */
func (a *NestedSelectControl) Required(value interface{}) *NestedSelectControl {
    a.Set("required", value)
    return a
}

/**
 * staticSchema
 */
func (a *NestedSelectControl) StaticSchema(value interface{}) *NestedSelectControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * useMobileUI
 */
func (a *NestedSelectControl) UseMobileUI(value interface{}) *NestedSelectControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * multiple
 */
func (a *NestedSelectControl) Multiple(value interface{}) *NestedSelectControl {
    a.Set("multiple", value)
    return a
}

/**
 * addApi
 */
func (a *NestedSelectControl) AddApi(value interface{}) *NestedSelectControl {
    a.Set("addApi", value)
    return a
}

/**
 * createBtnLabel
 */
func (a *NestedSelectControl) CreateBtnLabel(value interface{}) *NestedSelectControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * labelClassName
 */
func (a *NestedSelectControl) LabelClassName(value interface{}) *NestedSelectControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * horizontal
 */
func (a *NestedSelectControl) Horizontal(value interface{}) *NestedSelectControl {
    a.Set("horizontal", value)
    return a
}

/**
 */
func (a *NestedSelectControl) Type(value interface{}) *NestedSelectControl {
    a.Set("type", value)
    return a
}

/**
 * cascade
 */
func (a *NestedSelectControl) Cascade(value interface{}) *NestedSelectControl {
    a.Set("cascade", value)
    return a
}

/**
 * withChildren
 */
func (a *NestedSelectControl) WithChildren(value interface{}) *NestedSelectControl {
    a.Set("withChildren", value)
    return a
}

/**
 * addControls
 */
func (a *NestedSelectControl) AddControls(value interface{}) *NestedSelectControl {
    a.Set("addControls", value)
    return a
}

/**
 * extraName
 */
func (a *NestedSelectControl) ExtraName(value interface{}) *NestedSelectControl {
    a.Set("extraName", value)
    return a
}

/**
 * description
 */
func (a *NestedSelectControl) Description(value interface{}) *NestedSelectControl {
    a.Set("description", value)
    return a
}

/**
 * validationErrors
 */
func (a *NestedSelectControl) ValidationErrors(value interface{}) *NestedSelectControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * selectFirst
 */
func (a *NestedSelectControl) SelectFirst(value interface{}) *NestedSelectControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * labelWidth
 */
func (a *NestedSelectControl) LabelWidth(value interface{}) *NestedSelectControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * hint
 */
func (a *NestedSelectControl) Hint(value interface{}) *NestedSelectControl {
    a.Set("hint", value)
    return a
}

/**
 * validateOnChange
 */
func (a *NestedSelectControl) ValidateOnChange(value interface{}) *NestedSelectControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * desc
 */
func (a *NestedSelectControl) Desc(value interface{}) *NestedSelectControl {
    a.Set("desc", value)
    return a
}

/**
 * initAutoFill
 */
func (a *NestedSelectControl) InitAutoFill(value interface{}) *NestedSelectControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * options
 */
func (a *NestedSelectControl) Options(value interface{}) *NestedSelectControl {
    a.Set("options", value)
    return a
}

/**
 * initFetch
 */
func (a *NestedSelectControl) InitFetch(value interface{}) *NestedSelectControl {
    a.Set("initFetch", value)
    return a
}

/**
 * valuesNoWrap
 */
func (a *NestedSelectControl) ValuesNoWrap(value interface{}) *NestedSelectControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * deleteApi
 */
func (a *NestedSelectControl) DeleteApi(value interface{}) *NestedSelectControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * deleteConfirmText
 */
func (a *NestedSelectControl) DeleteConfirmText(value interface{}) *NestedSelectControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *NestedSelectControl) ReadOnlyOn(value interface{}) *NestedSelectControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * disabledOn
 */
func (a *NestedSelectControl) DisabledOn(value interface{}) *NestedSelectControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * size
 */
func (a *NestedSelectControl) Size(value interface{}) *NestedSelectControl {
    a.Set("size", value)
    return a
}

/**
 * submitOnChange
 */
func (a *NestedSelectControl) SubmitOnChange(value interface{}) *NestedSelectControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * visibleOn
 */
func (a *NestedSelectControl) VisibleOn(value interface{}) *NestedSelectControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * menuClassName
 */
func (a *NestedSelectControl) MenuClassName(value interface{}) *NestedSelectControl {
    a.Set("menuClassName", value)
    return a
}

/**
 * onlyChildren
 */
func (a *NestedSelectControl) OnlyChildren(value interface{}) *NestedSelectControl {
    a.Set("onlyChildren", value)
    return a
}

/**
 * onlyLeaf
 */
func (a *NestedSelectControl) OnlyLeaf(value interface{}) *NestedSelectControl {
    a.Set("onlyLeaf", value)
    return a
}

/**
 * maxTagCount
 */
func (a *NestedSelectControl) MaxTagCount(value interface{}) *NestedSelectControl {
    a.Set("maxTagCount", value)
    return a
}

/**
 * addDialog
 */
func (a *NestedSelectControl) AddDialog(value interface{}) *NestedSelectControl {
    a.Set("addDialog", value)
    return a
}
