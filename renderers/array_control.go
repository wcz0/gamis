package renderers


/**
 * InputArray 数组输入框。 combo 的别名。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/array

 * @author wcz0
 * @version 6.2.2
 */
type ArrayControl struct {
	*BaseRenderer
}

func NewArrayControl() *ArrayControl {
    a := &ArrayControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-array")
    return a
}


func (a *ArrayControl) Set(name string, value interface{}) *ArrayControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * validations
 */
func (a *ArrayControl) Validations(value interface{}) *ArrayControl {
    a.Set("validations", value)
    return a
}

/**
 * scaffold
 */
func (a *ArrayControl) Scaffold(value interface{}) *ArrayControl {
    a.Set("scaffold", value)
    return a
}

/**
 * strictMode
 */
func (a *ArrayControl) StrictMode(value interface{}) *ArrayControl {
    a.Set("strictMode", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *ArrayControl) Width(value interface{}) *ArrayControl {
    a.Set("width", value)
    return a
}

/**
 * labelClassName
 */
func (a *ArrayControl) LabelClassName(value interface{}) *ArrayControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * extraName
 */
func (a *ArrayControl) ExtraName(value interface{}) *ArrayControl {
    a.Set("extraName", value)
    return a
}

/**
 * inputClassName
 */
func (a *ArrayControl) InputClassName(value interface{}) *ArrayControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *ArrayControl) DescriptionClassName(value interface{}) *ArrayControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * row
 */
func (a *ArrayControl) Row(value interface{}) *ArrayControl {
    a.Set("row", value)
    return a
}

/**
 * className
 */
func (a *ArrayControl) ClassName(value interface{}) *ArrayControl {
    a.Set("className", value)
    return a
}

/**
 * disabledOn
 */
func (a *ArrayControl) DisabledOn(value interface{}) *ArrayControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticSchema
 */
func (a *ArrayControl) StaticSchema(value interface{}) *ArrayControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * addButtonClassName
 */
func (a *ArrayControl) AddButtonClassName(value interface{}) *ArrayControl {
    a.Set("addButtonClassName", value)
    return a
}

/**
 * minLength
 */
func (a *ArrayControl) MinLength(value interface{}) *ArrayControl {
    a.Set("minLength", value)
    return a
}

/**
 * labelAlign
 */
func (a *ArrayControl) LabelAlign(value interface{}) *ArrayControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * validateOnChange
 */
func (a *ArrayControl) ValidateOnChange(value interface{}) *ArrayControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * staticOn
 */
func (a *ArrayControl) StaticOn(value interface{}) *ArrayControl {
    a.Set("staticOn", value)
    return a
}

/**
 * noBorder
 */
func (a *ArrayControl) NoBorder(value interface{}) *ArrayControl {
    a.Set("noBorder", value)
    return a
}

/**
 * tabsStyle
 */
func (a *ArrayControl) TabsStyle(value interface{}) *ArrayControl {
    a.Set("tabsStyle", value)
    return a
}

/**
 * testIdBuilder
 */
func (a *ArrayControl) TestIdBuilder(value interface{}) *ArrayControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * submitOnChange
 */
func (a *ArrayControl) SubmitOnChange(value interface{}) *ArrayControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * required
 */
func (a *ArrayControl) Required(value interface{}) *ArrayControl {
    a.Set("required", value)
    return a
}

/**
 * validateApi
 */
func (a *ArrayControl) ValidateApi(value interface{}) *ArrayControl {
    a.Set("validateApi", value)
    return a
}

/**
 * initAutoFill
 */
func (a *ArrayControl) InitAutoFill(value interface{}) *ArrayControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * static
 */
func (a *ArrayControl) Static(value interface{}) *ArrayControl {
    a.Set("static", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *ArrayControl) StaticLabelClassName(value interface{}) *ArrayControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * maxLength
 */
func (a *ArrayControl) MaxLength(value interface{}) *ArrayControl {
    a.Set("maxLength", value)
    return a
}

/**
 * labelWidth
 */
func (a *ArrayControl) LabelWidth(value interface{}) *ArrayControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * hiddenOn
 */
func (a *ArrayControl) HiddenOn(value interface{}) *ArrayControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * deleteConfirmText
 */
func (a *ArrayControl) DeleteConfirmText(value interface{}) *ArrayControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * deleteApi
 */
func (a *ArrayControl) DeleteApi(value interface{}) *ArrayControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * addButtonText
 */
func (a *ArrayControl) AddButtonText(value interface{}) *ArrayControl {
    a.Set("addButtonText", value)
    return a
}

/**
 * joinValues
 */
func (a *ArrayControl) JoinValues(value interface{}) *ArrayControl {
    a.Set("joinValues", value)
    return a
}

/**
 * perPage
 */
func (a *ArrayControl) PerPage(value interface{}) *ArrayControl {
    a.Set("perPage", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *ArrayControl) ClearValueOnHidden(value interface{}) *ArrayControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * validationErrors
 */
func (a *ArrayControl) ValidationErrors(value interface{}) *ArrayControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *ArrayControl) StaticPlaceholder(value interface{}) *ArrayControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *ArrayControl) StaticInputClassName(value interface{}) *ArrayControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * flat
 */
func (a *ArrayControl) Flat(value interface{}) *ArrayControl {
    a.Set("flat", value)
    return a
}

/**
 * subFormMode
 */
func (a *ArrayControl) SubFormMode(value interface{}) *ArrayControl {
    a.Set("subFormMode", value)
    return a
}

/**
 * inline
 */
func (a *ArrayControl) Inline(value interface{}) *ArrayControl {
    a.Set("inline", value)
    return a
}

/**
 * hidden
 */
func (a *ArrayControl) Hidden(value interface{}) *ArrayControl {
    a.Set("hidden", value)
    return a
}

/**
 * visibleOn
 */
func (a *ArrayControl) VisibleOn(value interface{}) *ArrayControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * style
 */
func (a *ArrayControl) Style(value interface{}) *ArrayControl {
    a.Set("style", value)
    return a
}

/**
 * draggable
 */
func (a *ArrayControl) Draggable(value interface{}) *ArrayControl {
    a.Set("draggable", value)
    return a
}

/**
 * draggableTip
 */
func (a *ArrayControl) DraggableTip(value interface{}) *ArrayControl {
    a.Set("draggableTip", value)
    return a
}

/**
 * multiLine
 */
func (a *ArrayControl) MultiLine(value interface{}) *ArrayControl {
    a.Set("multiLine", value)
    return a
}

/**
 * messages
 */
func (a *ArrayControl) Messages(value interface{}) *ArrayControl {
    a.Set("messages", value)
    return a
}

/**
 * labelOverflow
 */
func (a *ArrayControl) LabelOverflow(value interface{}) *ArrayControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * size
 */
func (a *ArrayControl) Size(value interface{}) *ArrayControl {
    a.Set("size", value)
    return a
}

/**
 * typeSwitchable
 */
func (a *ArrayControl) TypeSwitchable(value interface{}) *ArrayControl {
    a.Set("typeSwitchable", value)
    return a
}

/**
 * delimiter
 */
func (a *ArrayControl) Delimiter(value interface{}) *ArrayControl {
    a.Set("delimiter", value)
    return a
}

/**
 * label
 */
func (a *ArrayControl) Label(value interface{}) *ArrayControl {
    a.Set("label", value)
    return a
}

/**
 * horizontal
 */
func (a *ArrayControl) Horizontal(value interface{}) *ArrayControl {
    a.Set("horizontal", value)
    return a
}

/**
 * placeholder
 */
func (a *ArrayControl) Placeholder(value interface{}) *ArrayControl {
    a.Set("placeholder", value)
    return a
}

/**
 * onEvent
 */
func (a *ArrayControl) OnEvent(value interface{}) *ArrayControl {
    a.Set("onEvent", value)
    return a
}

/**
 * staticClassName
 */
func (a *ArrayControl) StaticClassName(value interface{}) *ArrayControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * formClassName
 */
func (a *ArrayControl) FormClassName(value interface{}) *ArrayControl {
    a.Set("formClassName", value)
    return a
}

/**
 * addable
 */
func (a *ArrayControl) Addable(value interface{}) *ArrayControl {
    a.Set("addable", value)
    return a
}

/**
 * useMobileUI
 */
func (a *ArrayControl) UseMobileUI(value interface{}) *ArrayControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * value
 */
func (a *ArrayControl) Value(value interface{}) *ArrayControl {
    a.Set("value", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *ArrayControl) ReadOnlyOn(value interface{}) *ArrayControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * mode
 */
func (a *ArrayControl) Mode(value interface{}) *ArrayControl {
    a.Set("mode", value)
    return a
}

/**
 * removable
 */
func (a *ArrayControl) Removable(value interface{}) *ArrayControl {
    a.Set("removable", value)
    return a
}

/**
 * tabsLabelTpl
 */
func (a *ArrayControl) TabsLabelTpl(value interface{}) *ArrayControl {
    a.Set("tabsLabelTpl", value)
    return a
}

/**
 * name
 */
func (a *ArrayControl) Name(value interface{}) *ArrayControl {
    a.Set("name", value)
    return a
}

/**
 * description
 */
func (a *ArrayControl) Description(value interface{}) *ArrayControl {
    a.Set("description", value)
    return a
}

/**
 * visible
 */
func (a *ArrayControl) Visible(value interface{}) *ArrayControl {
    a.Set("visible", value)
    return a
}

/**
 * conditions
 */
func (a *ArrayControl) Conditions(value interface{}) *ArrayControl {
    a.Set("conditions", value)
    return a
}

/**
 * items
 */
func (a *ArrayControl) Items(value interface{}) *ArrayControl {
    a.Set("items", value)
    return a
}

/**
 * editorSetting
 */
func (a *ArrayControl) EditorSetting(value interface{}) *ArrayControl {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *ArrayControl) Type(value interface{}) *ArrayControl {
    a.Set("type", value)
    return a
}

/**
 * multiple
 */
func (a *ArrayControl) Multiple(value interface{}) *ArrayControl {
    a.Set("multiple", value)
    return a
}

/**
 * tabsMode
 */
func (a *ArrayControl) TabsMode(value interface{}) *ArrayControl {
    a.Set("tabsMode", value)
    return a
}

/**
 * syncFields
 */
func (a *ArrayControl) SyncFields(value interface{}) *ArrayControl {
    a.Set("syncFields", value)
    return a
}

/**
 * remark
 */
func (a *ArrayControl) Remark(value interface{}) *ArrayControl {
    a.Set("remark", value)
    return a
}

/**
 * labelRemark
 */
func (a *ArrayControl) LabelRemark(value interface{}) *ArrayControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * subFormHorizontal
 */
func (a *ArrayControl) SubFormHorizontal(value interface{}) *ArrayControl {
    a.Set("subFormHorizontal", value)
    return a
}

/**
 * canAccessSuperData
 */
func (a *ArrayControl) CanAccessSuperData(value interface{}) *ArrayControl {
    a.Set("canAccessSuperData", value)
    return a
}

/**
 * nullable
 */
func (a *ArrayControl) Nullable(value interface{}) *ArrayControl {
    a.Set("nullable", value)
    return a
}

/**
 * hint
 */
func (a *ArrayControl) Hint(value interface{}) *ArrayControl {
    a.Set("hint", value)
    return a
}

/**
 * readOnly
 */
func (a *ArrayControl) ReadOnly(value interface{}) *ArrayControl {
    a.Set("readOnly", value)
    return a
}

/**
 * desc
 */
func (a *ArrayControl) Desc(value interface{}) *ArrayControl {
    a.Set("desc", value)
    return a
}

/**
 * autoFill
 */
func (a *ArrayControl) AutoFill(value interface{}) *ArrayControl {
    a.Set("autoFill", value)
    return a
}

/**
 * addattop
 */
func (a *ArrayControl) Addattop(value interface{}) *ArrayControl {
    a.Set("addattop", value)
    return a
}

/**
 * disabled
 */
func (a *ArrayControl) Disabled(value interface{}) *ArrayControl {
    a.Set("disabled", value)
    return a
}

/**
 * id
 */
func (a *ArrayControl) Id(value interface{}) *ArrayControl {
    a.Set("id", value)
    return a
}

/**
 * lazyLoad
 */
func (a *ArrayControl) LazyLoad(value interface{}) *ArrayControl {
    a.Set("lazyLoad", value)
    return a
}

/**
 * updatePristineAfterStoreDataReInit
 */
func (a *ArrayControl) UpdatePristineAfterStoreDataReInit(value interface{}) *ArrayControl {
    a.Set("updatePristineAfterStoreDataReInit", value)
    return a
}
