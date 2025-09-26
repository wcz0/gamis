package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type ComboControl struct {
	*BaseRenderer
}

func NewComboControl() *ComboControl {
    a := &ComboControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "combo")
    return a
}


func (a *ComboControl) Set(name string, value interface{}) *ComboControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * tabsStyle
 */
func (a *ComboControl) TabsStyle(value interface{}) *ComboControl {
    a.Set("tabsStyle", value)
    return a
}

/**
 * label
 */
func (a *ComboControl) Label(value interface{}) *ComboControl {
    a.Set("label", value)
    return a
}

/**
 * inline
 */
func (a *ComboControl) Inline(value interface{}) *ComboControl {
    a.Set("inline", value)
    return a
}

/**
 * validationErrors
 */
func (a *ComboControl) ValidationErrors(value interface{}) *ComboControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * autoFill
 */
func (a *ComboControl) AutoFill(value interface{}) *ComboControl {
    a.Set("autoFill", value)
    return a
}

/**
 * subFormMode
 */
func (a *ComboControl) SubFormMode(value interface{}) *ComboControl {
    a.Set("subFormMode", value)
    return a
}

/**
 * lazyLoad
 */
func (a *ComboControl) LazyLoad(value interface{}) *ComboControl {
    a.Set("lazyLoad", value)
    return a
}

/**
 * labelAlign
 */
func (a *ComboControl) LabelAlign(value interface{}) *ComboControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * required
 */
func (a *ComboControl) Required(value interface{}) *ComboControl {
    a.Set("required", value)
    return a
}

/**
 * addable
 */
func (a *ComboControl) Addable(value interface{}) *ComboControl {
    a.Set("addable", value)
    return a
}

/**
 * size
 */
func (a *ComboControl) Size(value interface{}) *ComboControl {
    a.Set("size", value)
    return a
}

/**
 * value
 */
func (a *ComboControl) Value(value interface{}) *ComboControl {
    a.Set("value", value)
    return a
}

/**
 * hidden
 */
func (a *ComboControl) Hidden(value interface{}) *ComboControl {
    a.Set("hidden", value)
    return a
}

/**
 * staticClassName
 */
func (a *ComboControl) StaticClassName(value interface{}) *ComboControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * addattop
 */
func (a *ComboControl) Addattop(value interface{}) *ComboControl {
    a.Set("addattop", value)
    return a
}

/**
 * draggable
 */
func (a *ComboControl) Draggable(value interface{}) *ComboControl {
    a.Set("draggable", value)
    return a
}

/**
 * joinValues
 */
func (a *ComboControl) JoinValues(value interface{}) *ComboControl {
    a.Set("joinValues", value)
    return a
}

/**
 * placeholder
 */
func (a *ComboControl) Placeholder(value interface{}) *ComboControl {
    a.Set("placeholder", value)
    return a
}

/**
 * validateApi
 */
func (a *ComboControl) ValidateApi(value interface{}) *ComboControl {
    a.Set("validateApi", value)
    return a
}

/**
 * conditions
 */
func (a *ComboControl) Conditions(value interface{}) *ComboControl {
    a.Set("conditions", value)
    return a
}

/**
 * submitOnChange
 */
func (a *ComboControl) SubmitOnChange(value interface{}) *ComboControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * desc
 */
func (a *ComboControl) Desc(value interface{}) *ComboControl {
    a.Set("desc", value)
    return a
}

/**
 * initAutoFill
 */
func (a *ComboControl) InitAutoFill(value interface{}) *ComboControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * disabledOn
 */
func (a *ComboControl) DisabledOn(value interface{}) *ComboControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * deleteApi
 */
func (a *ComboControl) DeleteApi(value interface{}) *ComboControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * addButtonText
 */
func (a *ComboControl) AddButtonText(value interface{}) *ComboControl {
    a.Set("addButtonText", value)
    return a
}

/**
 * items
 */
func (a *ComboControl) Items(value interface{}) *ComboControl {
    a.Set("items", value)
    return a
}

/**
 * multiple
 */
func (a *ComboControl) Multiple(value interface{}) *ComboControl {
    a.Set("multiple", value)
    return a
}

/**
 * remark
 */
func (a *ComboControl) Remark(value interface{}) *ComboControl {
    a.Set("remark", value)
    return a
}

/**
 * staticOn
 */
func (a *ComboControl) StaticOn(value interface{}) *ComboControl {
    a.Set("staticOn", value)
    return a
}

/**
 * deleteConfirmText
 */
func (a *ComboControl) DeleteConfirmText(value interface{}) *ComboControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * minLength
 */
func (a *ComboControl) MinLength(value interface{}) *ComboControl {
    a.Set("minLength", value)
    return a
}

/**
 * nullable
 */
func (a *ComboControl) Nullable(value interface{}) *ComboControl {
    a.Set("nullable", value)
    return a
}

/**
 * inputClassName
 */
func (a *ComboControl) InputClassName(value interface{}) *ComboControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * addButtonClassName
 */
func (a *ComboControl) AddButtonClassName(value interface{}) *ComboControl {
    a.Set("addButtonClassName", value)
    return a
}

/**
 * maxLength
 */
func (a *ComboControl) MaxLength(value interface{}) *ComboControl {
    a.Set("maxLength", value)
    return a
}

/**
 * syncFields
 */
func (a *ComboControl) SyncFields(value interface{}) *ComboControl {
    a.Set("syncFields", value)
    return a
}

/**
 * messages
 */
func (a *ComboControl) Messages(value interface{}) *ComboControl {
    a.Set("messages", value)
    return a
}

/**
 * labelClassName
 */
func (a *ComboControl) LabelClassName(value interface{}) *ComboControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * hint
 */
func (a *ComboControl) Hint(value interface{}) *ComboControl {
    a.Set("hint", value)
    return a
}

/**
 * visibleOn
 */
func (a *ComboControl) VisibleOn(value interface{}) *ComboControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * style
 */
func (a *ComboControl) Style(value interface{}) *ComboControl {
    a.Set("style", value)
    return a
}

/**
 * scaffold
 */
func (a *ComboControl) Scaffold(value interface{}) *ComboControl {
    a.Set("scaffold", value)
    return a
}

/**
 * updatePristineAfterStoreDataReInit
 */
func (a *ComboControl) UpdatePristineAfterStoreDataReInit(value interface{}) *ComboControl {
    a.Set("updatePristineAfterStoreDataReInit", value)
    return a
}

/**
 * readOnly
 */
func (a *ComboControl) ReadOnly(value interface{}) *ComboControl {
    a.Set("readOnly", value)
    return a
}

/**
 * mode
 */
func (a *ComboControl) Mode(value interface{}) *ComboControl {
    a.Set("mode", value)
    return a
}

/**
 * subFormHorizontal
 */
func (a *ComboControl) SubFormHorizontal(value interface{}) *ComboControl {
    a.Set("subFormHorizontal", value)
    return a
}

/**
 * tabsLabelTpl
 */
func (a *ComboControl) TabsLabelTpl(value interface{}) *ComboControl {
    a.Set("tabsLabelTpl", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *ComboControl) Width(value interface{}) *ComboControl {
    a.Set("width", value)
    return a
}

/**
 * row
 */
func (a *ComboControl) Row(value interface{}) *ComboControl {
    a.Set("row", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *ComboControl) StaticLabelClassName(value interface{}) *ComboControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * editorSetting
 */
func (a *ComboControl) EditorSetting(value interface{}) *ComboControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * flat
 */
func (a *ComboControl) Flat(value interface{}) *ComboControl {
    a.Set("flat", value)
    return a
}

/**
 * testIdBuilder
 */
func (a *ComboControl) TestIdBuilder(value interface{}) *ComboControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *ComboControl) StaticInputClassName(value interface{}) *ComboControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *ComboControl) StaticSchema(value interface{}) *ComboControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * noBorder
 */
func (a *ComboControl) NoBorder(value interface{}) *ComboControl {
    a.Set("noBorder", value)
    return a
}

/**
 * labelOverflow
 */
func (a *ComboControl) LabelOverflow(value interface{}) *ComboControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * extraName
 */
func (a *ComboControl) ExtraName(value interface{}) *ComboControl {
    a.Set("extraName", value)
    return a
}

/**
 * validateOnChange
 */
func (a *ComboControl) ValidateOnChange(value interface{}) *ComboControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * validations
 */
func (a *ComboControl) Validations(value interface{}) *ComboControl {
    a.Set("validations", value)
    return a
}

/**
 * onEvent
 */
func (a *ComboControl) OnEvent(value interface{}) *ComboControl {
    a.Set("onEvent", value)
    return a
}

/**
 * hiddenOn
 */
func (a *ComboControl) HiddenOn(value interface{}) *ComboControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * static
 */
func (a *ComboControl) Static(value interface{}) *ComboControl {
    a.Set("static", value)
    return a
}

/**
 */
func (a *ComboControl) Type(value interface{}) *ComboControl {
    a.Set("type", value)
    return a
}

/**
 * removable
 */
func (a *ComboControl) Removable(value interface{}) *ComboControl {
    a.Set("removable", value)
    return a
}

/**
 * perPage
 */
func (a *ComboControl) PerPage(value interface{}) *ComboControl {
    a.Set("perPage", value)
    return a
}

/**
 * labelWidth
 */
func (a *ComboControl) LabelWidth(value interface{}) *ComboControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *ComboControl) DescriptionClassName(value interface{}) *ComboControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * visible
 */
func (a *ComboControl) Visible(value interface{}) *ComboControl {
    a.Set("visible", value)
    return a
}

/**
 * useMobileUI
 */
func (a *ComboControl) UseMobileUI(value interface{}) *ComboControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * delimiter
 */
func (a *ComboControl) Delimiter(value interface{}) *ComboControl {
    a.Set("delimiter", value)
    return a
}

/**
 * multiLine
 */
func (a *ComboControl) MultiLine(value interface{}) *ComboControl {
    a.Set("multiLine", value)
    return a
}

/**
 * name
 */
func (a *ComboControl) Name(value interface{}) *ComboControl {
    a.Set("name", value)
    return a
}

/**
 * description
 */
func (a *ComboControl) Description(value interface{}) *ComboControl {
    a.Set("description", value)
    return a
}

/**
 * horizontal
 */
func (a *ComboControl) Horizontal(value interface{}) *ComboControl {
    a.Set("horizontal", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *ComboControl) ClearValueOnHidden(value interface{}) *ComboControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * disabled
 */
func (a *ComboControl) Disabled(value interface{}) *ComboControl {
    a.Set("disabled", value)
    return a
}

/**
 * id
 */
func (a *ComboControl) Id(value interface{}) *ComboControl {
    a.Set("id", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *ComboControl) StaticPlaceholder(value interface{}) *ComboControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * labelRemark
 */
func (a *ComboControl) LabelRemark(value interface{}) *ComboControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * formClassName
 */
func (a *ComboControl) FormClassName(value interface{}) *ComboControl {
    a.Set("formClassName", value)
    return a
}

/**
 * draggableTip
 */
func (a *ComboControl) DraggableTip(value interface{}) *ComboControl {
    a.Set("draggableTip", value)
    return a
}

/**
 * strictMode
 */
func (a *ComboControl) StrictMode(value interface{}) *ComboControl {
    a.Set("strictMode", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *ComboControl) ReadOnlyOn(value interface{}) *ComboControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * className
 */
func (a *ComboControl) ClassName(value interface{}) *ComboControl {
    a.Set("className", value)
    return a
}

/**
 * typeSwitchable
 */
func (a *ComboControl) TypeSwitchable(value interface{}) *ComboControl {
    a.Set("typeSwitchable", value)
    return a
}

/**
 * canAccessSuperData
 */
func (a *ComboControl) CanAccessSuperData(value interface{}) *ComboControl {
    a.Set("canAccessSuperData", value)
    return a
}

/**
 * tabsMode
 */
func (a *ComboControl) TabsMode(value interface{}) *ComboControl {
    a.Set("tabsMode", value)
    return a
}
