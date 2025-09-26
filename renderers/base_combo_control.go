package renderers


/**
 * Combo 组合输入框类型 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/combo

 * @author wcz0
 * @version 6.2.2
 */
type BaseComboControl struct {
	*BaseRenderer
}

func NewBaseComboControl() *BaseComboControl {
    a := &BaseComboControl{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}


func (a *BaseComboControl) Set(name string, value interface{}) *BaseComboControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * scaffold
 */
func (a *BaseComboControl) Scaffold(value interface{}) *BaseComboControl {
    a.Set("scaffold", value)
    return a
}

/**
 * formClassName
 */
func (a *BaseComboControl) FormClassName(value interface{}) *BaseComboControl {
    a.Set("formClassName", value)
    return a
}

/**
 * validateOnChange
 */
func (a *BaseComboControl) ValidateOnChange(value interface{}) *BaseComboControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * static
 */
func (a *BaseComboControl) Static(value interface{}) *BaseComboControl {
    a.Set("static", value)
    return a
}

/**
 * maxLength
 */
func (a *BaseComboControl) MaxLength(value interface{}) *BaseComboControl {
    a.Set("maxLength", value)
    return a
}

/**
 * subFormMode
 */
func (a *BaseComboControl) SubFormMode(value interface{}) *BaseComboControl {
    a.Set("subFormMode", value)
    return a
}

/**
 * inline
 */
func (a *BaseComboControl) Inline(value interface{}) *BaseComboControl {
    a.Set("inline", value)
    return a
}

/**
 * inputClassName
 */
func (a *BaseComboControl) InputClassName(value interface{}) *BaseComboControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * hiddenOn
 */
func (a *BaseComboControl) HiddenOn(value interface{}) *BaseComboControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * id
 */
func (a *BaseComboControl) Id(value interface{}) *BaseComboControl {
    a.Set("id", value)
    return a
}

/**
 * editorSetting
 */
func (a *BaseComboControl) EditorSetting(value interface{}) *BaseComboControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * size
 */
func (a *BaseComboControl) Size(value interface{}) *BaseComboControl {
    a.Set("size", value)
    return a
}

/**
 * noBorder
 */
func (a *BaseComboControl) NoBorder(value interface{}) *BaseComboControl {
    a.Set("noBorder", value)
    return a
}

/**
 * conditions
 */
func (a *BaseComboControl) Conditions(value interface{}) *BaseComboControl {
    a.Set("conditions", value)
    return a
}

/**
 * labelWidth
 */
func (a *BaseComboControl) LabelWidth(value interface{}) *BaseComboControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * remark
 */
func (a *BaseComboControl) Remark(value interface{}) *BaseComboControl {
    a.Set("remark", value)
    return a
}

/**
 * horizontal
 */
func (a *BaseComboControl) Horizontal(value interface{}) *BaseComboControl {
    a.Set("horizontal", value)
    return a
}

/**
 * deleteApi
 */
func (a *BaseComboControl) DeleteApi(value interface{}) *BaseComboControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * addButtonText
 */
func (a *BaseComboControl) AddButtonText(value interface{}) *BaseComboControl {
    a.Set("addButtonText", value)
    return a
}

/**
 * tabsMode
 */
func (a *BaseComboControl) TabsMode(value interface{}) *BaseComboControl {
    a.Set("tabsMode", value)
    return a
}

/**
 * tabsStyle
 */
func (a *BaseComboControl) TabsStyle(value interface{}) *BaseComboControl {
    a.Set("tabsStyle", value)
    return a
}

/**
 * hint
 */
func (a *BaseComboControl) Hint(value interface{}) *BaseComboControl {
    a.Set("hint", value)
    return a
}

/**
 * required
 */
func (a *BaseComboControl) Required(value interface{}) *BaseComboControl {
    a.Set("required", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *BaseComboControl) StaticInputClassName(value interface{}) *BaseComboControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * removable
 */
func (a *BaseComboControl) Removable(value interface{}) *BaseComboControl {
    a.Set("removable", value)
    return a
}

/**
 * strictMode
 */
func (a *BaseComboControl) StrictMode(value interface{}) *BaseComboControl {
    a.Set("strictMode", value)
    return a
}

/**
 * nullable
 */
func (a *BaseComboControl) Nullable(value interface{}) *BaseComboControl {
    a.Set("nullable", value)
    return a
}

/**
 * labelClassName
 */
func (a *BaseComboControl) LabelClassName(value interface{}) *BaseComboControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * placeholder
 */
func (a *BaseComboControl) Placeholder(value interface{}) *BaseComboControl {
    a.Set("placeholder", value)
    return a
}

/**
 * autoFill
 */
func (a *BaseComboControl) AutoFill(value interface{}) *BaseComboControl {
    a.Set("autoFill", value)
    return a
}

/**
 * staticOn
 */
func (a *BaseComboControl) StaticOn(value interface{}) *BaseComboControl {
    a.Set("staticOn", value)
    return a
}

/**
 * typeSwitchable
 */
func (a *BaseComboControl) TypeSwitchable(value interface{}) *BaseComboControl {
    a.Set("typeSwitchable", value)
    return a
}

/**
 * draggableTip
 */
func (a *BaseComboControl) DraggableTip(value interface{}) *BaseComboControl {
    a.Set("draggableTip", value)
    return a
}

/**
 * subFormHorizontal
 */
func (a *BaseComboControl) SubFormHorizontal(value interface{}) *BaseComboControl {
    a.Set("subFormHorizontal", value)
    return a
}

/**
 * labelAlign
 */
func (a *BaseComboControl) LabelAlign(value interface{}) *BaseComboControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * labelRemark
 */
func (a *BaseComboControl) LabelRemark(value interface{}) *BaseComboControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * validationErrors
 */
func (a *BaseComboControl) ValidationErrors(value interface{}) *BaseComboControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * deleteConfirmText
 */
func (a *BaseComboControl) DeleteConfirmText(value interface{}) *BaseComboControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * className
 */
func (a *BaseComboControl) ClassName(value interface{}) *BaseComboControl {
    a.Set("className", value)
    return a
}

/**
 * visibleOn
 */
func (a *BaseComboControl) VisibleOn(value interface{}) *BaseComboControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * style
 */
func (a *BaseComboControl) Style(value interface{}) *BaseComboControl {
    a.Set("style", value)
    return a
}

/**
 * addable
 */
func (a *BaseComboControl) Addable(value interface{}) *BaseComboControl {
    a.Set("addable", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *BaseComboControl) ClearValueOnHidden(value interface{}) *BaseComboControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * disabledOn
 */
func (a *BaseComboControl) DisabledOn(value interface{}) *BaseComboControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticSchema
 */
func (a *BaseComboControl) StaticSchema(value interface{}) *BaseComboControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * canAccessSuperData
 */
func (a *BaseComboControl) CanAccessSuperData(value interface{}) *BaseComboControl {
    a.Set("canAccessSuperData", value)
    return a
}

/**
 * readOnly
 */
func (a *BaseComboControl) ReadOnly(value interface{}) *BaseComboControl {
    a.Set("readOnly", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *BaseComboControl) DescriptionClassName(value interface{}) *BaseComboControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * items
 */
func (a *BaseComboControl) Items(value interface{}) *BaseComboControl {
    a.Set("items", value)
    return a
}

/**
 * delimiter
 */
func (a *BaseComboControl) Delimiter(value interface{}) *BaseComboControl {
    a.Set("delimiter", value)
    return a
}

/**
 * multiple
 */
func (a *BaseComboControl) Multiple(value interface{}) *BaseComboControl {
    a.Set("multiple", value)
    return a
}

/**
 * syncFields
 */
func (a *BaseComboControl) SyncFields(value interface{}) *BaseComboControl {
    a.Set("syncFields", value)
    return a
}

/**
 * visible
 */
func (a *BaseComboControl) Visible(value interface{}) *BaseComboControl {
    a.Set("visible", value)
    return a
}

/**
 * addButtonClassName
 */
func (a *BaseComboControl) AddButtonClassName(value interface{}) *BaseComboControl {
    a.Set("addButtonClassName", value)
    return a
}

/**
 * draggable
 */
func (a *BaseComboControl) Draggable(value interface{}) *BaseComboControl {
    a.Set("draggable", value)
    return a
}

/**
 * flat
 */
func (a *BaseComboControl) Flat(value interface{}) *BaseComboControl {
    a.Set("flat", value)
    return a
}

/**
 * minLength
 */
func (a *BaseComboControl) MinLength(value interface{}) *BaseComboControl {
    a.Set("minLength", value)
    return a
}

/**
 * tabsLabelTpl
 */
func (a *BaseComboControl) TabsLabelTpl(value interface{}) *BaseComboControl {
    a.Set("tabsLabelTpl", value)
    return a
}

/**
 * testIdBuilder
 */
func (a *BaseComboControl) TestIdBuilder(value interface{}) *BaseComboControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * useMobileUI
 */
func (a *BaseComboControl) UseMobileUI(value interface{}) *BaseComboControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * labelOverflow
 */
func (a *BaseComboControl) LabelOverflow(value interface{}) *BaseComboControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * description
 */
func (a *BaseComboControl) Description(value interface{}) *BaseComboControl {
    a.Set("description", value)
    return a
}

/**
 * value
 */
func (a *BaseComboControl) Value(value interface{}) *BaseComboControl {
    a.Set("value", value)
    return a
}

/**
 * disabled
 */
func (a *BaseComboControl) Disabled(value interface{}) *BaseComboControl {
    a.Set("disabled", value)
    return a
}

/**
 * joinValues
 */
func (a *BaseComboControl) JoinValues(value interface{}) *BaseComboControl {
    a.Set("joinValues", value)
    return a
}

/**
 * desc
 */
func (a *BaseComboControl) Desc(value interface{}) *BaseComboControl {
    a.Set("desc", value)
    return a
}

/**
 * onEvent
 */
func (a *BaseComboControl) OnEvent(value interface{}) *BaseComboControl {
    a.Set("onEvent", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *BaseComboControl) StaticPlaceholder(value interface{}) *BaseComboControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticClassName
 */
func (a *BaseComboControl) StaticClassName(value interface{}) *BaseComboControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *BaseComboControl) StaticLabelClassName(value interface{}) *BaseComboControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * multiLine
 */
func (a *BaseComboControl) MultiLine(value interface{}) *BaseComboControl {
    a.Set("multiLine", value)
    return a
}

/**
 * updatePristineAfterStoreDataReInit
 */
func (a *BaseComboControl) UpdatePristineAfterStoreDataReInit(value interface{}) *BaseComboControl {
    a.Set("updatePristineAfterStoreDataReInit", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *BaseComboControl) Width(value interface{}) *BaseComboControl {
    a.Set("width", value)
    return a
}

/**
 * name
 */
func (a *BaseComboControl) Name(value interface{}) *BaseComboControl {
    a.Set("name", value)
    return a
}

/**
 * mode
 */
func (a *BaseComboControl) Mode(value interface{}) *BaseComboControl {
    a.Set("mode", value)
    return a
}

/**
 * hidden
 */
func (a *BaseComboControl) Hidden(value interface{}) *BaseComboControl {
    a.Set("hidden", value)
    return a
}

/**
 * messages
 */
func (a *BaseComboControl) Messages(value interface{}) *BaseComboControl {
    a.Set("messages", value)
    return a
}

/**
 * initAutoFill
 */
func (a *BaseComboControl) InitAutoFill(value interface{}) *BaseComboControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * row
 */
func (a *BaseComboControl) Row(value interface{}) *BaseComboControl {
    a.Set("row", value)
    return a
}

/**
 * type
 */
func (a *BaseComboControl) Type(value interface{}) *BaseComboControl {
    a.Set("type", value)
    return a
}

/**
 * addattop
 */
func (a *BaseComboControl) Addattop(value interface{}) *BaseComboControl {
    a.Set("addattop", value)
    return a
}

/**
 * lazyLoad
 */
func (a *BaseComboControl) LazyLoad(value interface{}) *BaseComboControl {
    a.Set("lazyLoad", value)
    return a
}

/**
 * perPage
 */
func (a *BaseComboControl) PerPage(value interface{}) *BaseComboControl {
    a.Set("perPage", value)
    return a
}

/**
 * label
 */
func (a *BaseComboControl) Label(value interface{}) *BaseComboControl {
    a.Set("label", value)
    return a
}

/**
 * extraName
 */
func (a *BaseComboControl) ExtraName(value interface{}) *BaseComboControl {
    a.Set("extraName", value)
    return a
}

/**
 * submitOnChange
 */
func (a *BaseComboControl) SubmitOnChange(value interface{}) *BaseComboControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *BaseComboControl) ReadOnlyOn(value interface{}) *BaseComboControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * validations
 */
func (a *BaseComboControl) Validations(value interface{}) *BaseComboControl {
    a.Set("validations", value)
    return a
}

/**
 * validateApi
 */
func (a *BaseComboControl) ValidateApi(value interface{}) *BaseComboControl {
    a.Set("validateApi", value)
    return a
}
