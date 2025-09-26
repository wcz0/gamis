package renderers


/**
 * List 复选框 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/list-select

 * @author wcz0
 * @version 6.2.2
 */
type ListControl struct {
	*BaseRenderer
}

func NewListControl() *ListControl {
    a := &ListControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "list-select")
    return a
}


func (a *ListControl) Set(name string, value interface{}) *ListControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * className
 */
func (a *ListControl) ClassName(value interface{}) *ListControl {
    a.Set("className", value)
    return a
}

/**
 * label
 */
func (a *ListControl) Label(value interface{}) *ListControl {
    a.Set("label", value)
    return a
}

/**
 * value
 */
func (a *ListControl) Value(value interface{}) *ListControl {
    a.Set("value", value)
    return a
}

/**
 * staticClassName
 */
func (a *ListControl) StaticClassName(value interface{}) *ListControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *ListControl) Width(value interface{}) *ListControl {
    a.Set("width", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *ListControl) DescriptionClassName(value interface{}) *ListControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * mode
 */
func (a *ListControl) Mode(value interface{}) *ListControl {
    a.Set("mode", value)
    return a
}

/**
 * hiddenOn
 */
func (a *ListControl) HiddenOn(value interface{}) *ListControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticSchema
 */
func (a *ListControl) StaticSchema(value interface{}) *ListControl {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *ListControl) Type(value interface{}) *ListControl {
    a.Set("type", value)
    return a
}

/**
 * initFetch
 */
func (a *ListControl) InitFetch(value interface{}) *ListControl {
    a.Set("initFetch", value)
    return a
}

/**
 * creatable
 */
func (a *ListControl) Creatable(value interface{}) *ListControl {
    a.Set("creatable", value)
    return a
}

/**
 * extraName
 */
func (a *ListControl) ExtraName(value interface{}) *ListControl {
    a.Set("extraName", value)
    return a
}

/**
 * readOnly
 */
func (a *ListControl) ReadOnly(value interface{}) *ListControl {
    a.Set("readOnly", value)
    return a
}

/**
 * inline
 */
func (a *ListControl) Inline(value interface{}) *ListControl {
    a.Set("inline", value)
    return a
}

/**
 * resetValue
 */
func (a *ListControl) ResetValue(value interface{}) *ListControl {
    a.Set("resetValue", value)
    return a
}

/**
 * editApi
 */
func (a *ListControl) EditApi(value interface{}) *ListControl {
    a.Set("editApi", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *ListControl) ReadOnlyOn(value interface{}) *ListControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * required
 */
func (a *ListControl) Required(value interface{}) *ListControl {
    a.Set("required", value)
    return a
}

/**
 * deferApi
 */
func (a *ListControl) DeferApi(value interface{}) *ListControl {
    a.Set("deferApi", value)
    return a
}

/**
 * clearValueOnSourceChange
 */
func (a *ListControl) ClearValueOnSourceChange(value interface{}) *ListControl {
    a.Set("clearValueOnSourceChange", value)
    return a
}

/**
 * validations
 */
func (a *ListControl) Validations(value interface{}) *ListControl {
    a.Set("validations", value)
    return a
}

/**
 * static
 */
func (a *ListControl) Static(value interface{}) *ListControl {
    a.Set("static", value)
    return a
}

/**
 * staticOn
 */
func (a *ListControl) StaticOn(value interface{}) *ListControl {
    a.Set("staticOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *ListControl) StaticPlaceholder(value interface{}) *ListControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *ListControl) StaticInputClassName(value interface{}) *ListControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * joinValues
 */
func (a *ListControl) JoinValues(value interface{}) *ListControl {
    a.Set("joinValues", value)
    return a
}

/**
 * delimiter
 */
func (a *ListControl) Delimiter(value interface{}) *ListControl {
    a.Set("delimiter", value)
    return a
}

/**
 * itemSchema
 */
func (a *ListControl) ItemSchema(value interface{}) *ListControl {
    a.Set("itemSchema", value)
    return a
}

/**
 * activeItemSchema
 */
func (a *ListControl) ActiveItemSchema(value interface{}) *ListControl {
    a.Set("activeItemSchema", value)
    return a
}

/**
 * listClassName
 */
func (a *ListControl) ListClassName(value interface{}) *ListControl {
    a.Set("listClassName", value)
    return a
}

/**
 * removable
 */
func (a *ListControl) Removable(value interface{}) *ListControl {
    a.Set("removable", value)
    return a
}

/**
 * editable
 */
func (a *ListControl) Editable(value interface{}) *ListControl {
    a.Set("editable", value)
    return a
}

/**
 * editControls
 */
func (a *ListControl) EditControls(value interface{}) *ListControl {
    a.Set("editControls", value)
    return a
}

/**
 * validateApi
 */
func (a *ListControl) ValidateApi(value interface{}) *ListControl {
    a.Set("validateApi", value)
    return a
}

/**
 * hidden
 */
func (a *ListControl) Hidden(value interface{}) *ListControl {
    a.Set("hidden", value)
    return a
}

/**
 * style
 */
func (a *ListControl) Style(value interface{}) *ListControl {
    a.Set("style", value)
    return a
}

/**
 * editorSetting
 */
func (a *ListControl) EditorSetting(value interface{}) *ListControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * row
 */
func (a *ListControl) Row(value interface{}) *ListControl {
    a.Set("row", value)
    return a
}

/**
 * visible
 */
func (a *ListControl) Visible(value interface{}) *ListControl {
    a.Set("visible", value)
    return a
}

/**
 * onEvent
 */
func (a *ListControl) OnEvent(value interface{}) *ListControl {
    a.Set("onEvent", value)
    return a
}

/**
 * deferField
 */
func (a *ListControl) DeferField(value interface{}) *ListControl {
    a.Set("deferField", value)
    return a
}

/**
 * labelClassName
 */
func (a *ListControl) LabelClassName(value interface{}) *ListControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * submitOnDBClick
 */
func (a *ListControl) SubmitOnDBClick(value interface{}) *ListControl {
    a.Set("submitOnDBClick", value)
    return a
}

/**
 * checkAll
 */
func (a *ListControl) CheckAll(value interface{}) *ListControl {
    a.Set("checkAll", value)
    return a
}

/**
 * deleteConfirmText
 */
func (a *ListControl) DeleteConfirmText(value interface{}) *ListControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * labelAlign
 */
func (a *ListControl) LabelAlign(value interface{}) *ListControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * desc
 */
func (a *ListControl) Desc(value interface{}) *ListControl {
    a.Set("desc", value)
    return a
}

/**
 * horizontal
 */
func (a *ListControl) Horizontal(value interface{}) *ListControl {
    a.Set("horizontal", value)
    return a
}

/**
 * validationErrors
 */
func (a *ListControl) ValidationErrors(value interface{}) *ListControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *ListControl) ClearValueOnHidden(value interface{}) *ListControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * disabled
 */
func (a *ListControl) Disabled(value interface{}) *ListControl {
    a.Set("disabled", value)
    return a
}

/**
 * deleteApi
 */
func (a *ListControl) DeleteApi(value interface{}) *ListControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * extractValue
 */
func (a *ListControl) ExtractValue(value interface{}) *ListControl {
    a.Set("extractValue", value)
    return a
}

/**
 * createBtnLabel
 */
func (a *ListControl) CreateBtnLabel(value interface{}) *ListControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * name
 */
func (a *ListControl) Name(value interface{}) *ListControl {
    a.Set("name", value)
    return a
}

/**
 * submitOnChange
 */
func (a *ListControl) SubmitOnChange(value interface{}) *ListControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * disabledOn
 */
func (a *ListControl) DisabledOn(value interface{}) *ListControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * visibleOn
 */
func (a *ListControl) VisibleOn(value interface{}) *ListControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * initFetchOn
 */
func (a *ListControl) InitFetchOn(value interface{}) *ListControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * clearable
 */
func (a *ListControl) Clearable(value interface{}) *ListControl {
    a.Set("clearable", value)
    return a
}

/**
 * addApi
 */
func (a *ListControl) AddApi(value interface{}) *ListControl {
    a.Set("addApi", value)
    return a
}

/**
 * editDialog
 */
func (a *ListControl) EditDialog(value interface{}) *ListControl {
    a.Set("editDialog", value)
    return a
}

/**
 * initAutoFill
 */
func (a *ListControl) InitAutoFill(value interface{}) *ListControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * id
 */
func (a *ListControl) Id(value interface{}) *ListControl {
    a.Set("id", value)
    return a
}

/**
 * source
 */
func (a *ListControl) Source(value interface{}) *ListControl {
    a.Set("source", value)
    return a
}

/**
 * multiple
 */
func (a *ListControl) Multiple(value interface{}) *ListControl {
    a.Set("multiple", value)
    return a
}

/**
 * description
 */
func (a *ListControl) Description(value interface{}) *ListControl {
    a.Set("description", value)
    return a
}

/**
 * useMobileUI
 */
func (a *ListControl) UseMobileUI(value interface{}) *ListControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * size
 */
func (a *ListControl) Size(value interface{}) *ListControl {
    a.Set("size", value)
    return a
}

/**
 * imageClassName
 */
func (a *ListControl) ImageClassName(value interface{}) *ListControl {
    a.Set("imageClassName", value)
    return a
}

/**
 * options
 */
func (a *ListControl) Options(value interface{}) *ListControl {
    a.Set("options", value)
    return a
}

/**
 * labelWidth
 */
func (a *ListControl) LabelWidth(value interface{}) *ListControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * labelRemark
 */
func (a *ListControl) LabelRemark(value interface{}) *ListControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * validateOnChange
 */
func (a *ListControl) ValidateOnChange(value interface{}) *ListControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * inputClassName
 */
func (a *ListControl) InputClassName(value interface{}) *ListControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * placeholder
 */
func (a *ListControl) Placeholder(value interface{}) *ListControl {
    a.Set("placeholder", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *ListControl) StaticLabelClassName(value interface{}) *ListControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * selectFirst
 */
func (a *ListControl) SelectFirst(value interface{}) *ListControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * valuesNoWrap
 */
func (a *ListControl) ValuesNoWrap(value interface{}) *ListControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * addControls
 */
func (a *ListControl) AddControls(value interface{}) *ListControl {
    a.Set("addControls", value)
    return a
}

/**
 * addDialog
 */
func (a *ListControl) AddDialog(value interface{}) *ListControl {
    a.Set("addDialog", value)
    return a
}

/**
 * remark
 */
func (a *ListControl) Remark(value interface{}) *ListControl {
    a.Set("remark", value)
    return a
}

/**
 * hint
 */
func (a *ListControl) Hint(value interface{}) *ListControl {
    a.Set("hint", value)
    return a
}

/**
 * labelOverflow
 */
func (a *ListControl) LabelOverflow(value interface{}) *ListControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * autoFill
 */
func (a *ListControl) AutoFill(value interface{}) *ListControl {
    a.Set("autoFill", value)
    return a
}
