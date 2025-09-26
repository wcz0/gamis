package renderers


/**
 * JSON Schema Editor 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/json-schema-editor

 * @author wcz0
 * @version 6.2.2
 */
type JSONSchemaEditorControl struct {
	*BaseRenderer
}

func NewJSONSchemaEditorControl() *JSONSchemaEditorControl {
    a := &JSONSchemaEditorControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "json-schema-editor")
    return a
}


func (a *JSONSchemaEditorControl) Set(name string, value interface{}) *JSONSchemaEditorControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * name
 */
func (a *JSONSchemaEditorControl) Name(value interface{}) *JSONSchemaEditorControl {
    a.Set("name", value)
    return a
}

/**
 * useMobileUI
 */
func (a *JSONSchemaEditorControl) UseMobileUI(value interface{}) *JSONSchemaEditorControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * size
 */
func (a *JSONSchemaEditorControl) Size(value interface{}) *JSONSchemaEditorControl {
    a.Set("size", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *JSONSchemaEditorControl) StaticInputClassName(value interface{}) *JSONSchemaEditorControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * editorSetting
 */
func (a *JSONSchemaEditorControl) EditorSetting(value interface{}) *JSONSchemaEditorControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *JSONSchemaEditorControl) Width(value interface{}) *JSONSchemaEditorControl {
    a.Set("width", value)
    return a
}

/**
 * autoFill
 */
func (a *JSONSchemaEditorControl) AutoFill(value interface{}) *JSONSchemaEditorControl {
    a.Set("autoFill", value)
    return a
}

/**
 * labelAlign
 */
func (a *JSONSchemaEditorControl) LabelAlign(value interface{}) *JSONSchemaEditorControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * extraName
 */
func (a *JSONSchemaEditorControl) ExtraName(value interface{}) *JSONSchemaEditorControl {
    a.Set("extraName", value)
    return a
}

/**
 * desc
 */
func (a *JSONSchemaEditorControl) Desc(value interface{}) *JSONSchemaEditorControl {
    a.Set("desc", value)
    return a
}

/**
 * visibleOn
 */
func (a *JSONSchemaEditorControl) VisibleOn(value interface{}) *JSONSchemaEditorControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *JSONSchemaEditorControl) StaticPlaceholder(value interface{}) *JSONSchemaEditorControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * enableAdvancedSetting
 */
func (a *JSONSchemaEditorControl) EnableAdvancedSetting(value interface{}) *JSONSchemaEditorControl {
    a.Set("enableAdvancedSetting", value)
    return a
}

/**
 * labelClassName
 */
func (a *JSONSchemaEditorControl) LabelClassName(value interface{}) *JSONSchemaEditorControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * remark
 */
func (a *JSONSchemaEditorControl) Remark(value interface{}) *JSONSchemaEditorControl {
    a.Set("remark", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *JSONSchemaEditorControl) StaticLabelClassName(value interface{}) *JSONSchemaEditorControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * labelWidth
 */
func (a *JSONSchemaEditorControl) LabelWidth(value interface{}) *JSONSchemaEditorControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * staticSchema
 */
func (a *JSONSchemaEditorControl) StaticSchema(value interface{}) *JSONSchemaEditorControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * style
 */
func (a *JSONSchemaEditorControl) Style(value interface{}) *JSONSchemaEditorControl {
    a.Set("style", value)
    return a
}

/**
 * description
 */
func (a *JSONSchemaEditorControl) Description(value interface{}) *JSONSchemaEditorControl {
    a.Set("description", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *JSONSchemaEditorControl) DescriptionClassName(value interface{}) *JSONSchemaEditorControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * required
 */
func (a *JSONSchemaEditorControl) Required(value interface{}) *JSONSchemaEditorControl {
    a.Set("required", value)
    return a
}

/**
 * row
 */
func (a *JSONSchemaEditorControl) Row(value interface{}) *JSONSchemaEditorControl {
    a.Set("row", value)
    return a
}

/**
 * disabled
 */
func (a *JSONSchemaEditorControl) Disabled(value interface{}) *JSONSchemaEditorControl {
    a.Set("disabled", value)
    return a
}

/**
 * id
 */
func (a *JSONSchemaEditorControl) Id(value interface{}) *JSONSchemaEditorControl {
    a.Set("id", value)
    return a
}

/**
 * definitions
 */
func (a *JSONSchemaEditorControl) Definitions(value interface{}) *JSONSchemaEditorControl {
    a.Set("definitions", value)
    return a
}

/**
 * readOnly
 */
func (a *JSONSchemaEditorControl) ReadOnly(value interface{}) *JSONSchemaEditorControl {
    a.Set("readOnly", value)
    return a
}

/**
 * validateOnChange
 */
func (a *JSONSchemaEditorControl) ValidateOnChange(value interface{}) *JSONSchemaEditorControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * inline
 */
func (a *JSONSchemaEditorControl) Inline(value interface{}) *JSONSchemaEditorControl {
    a.Set("inline", value)
    return a
}

/**
 * placeholder
 */
func (a *JSONSchemaEditorControl) Placeholder(value interface{}) *JSONSchemaEditorControl {
    a.Set("placeholder", value)
    return a
}

/**
 * staticClassName
 */
func (a *JSONSchemaEditorControl) StaticClassName(value interface{}) *JSONSchemaEditorControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * showRootInfo
 */
func (a *JSONSchemaEditorControl) ShowRootInfo(value interface{}) *JSONSchemaEditorControl {
    a.Set("showRootInfo", value)
    return a
}

/**
 * label
 */
func (a *JSONSchemaEditorControl) Label(value interface{}) *JSONSchemaEditorControl {
    a.Set("label", value)
    return a
}

/**
 * disabledOn
 */
func (a *JSONSchemaEditorControl) DisabledOn(value interface{}) *JSONSchemaEditorControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * hiddenOn
 */
func (a *JSONSchemaEditorControl) HiddenOn(value interface{}) *JSONSchemaEditorControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * mini
 */
func (a *JSONSchemaEditorControl) Mini(value interface{}) *JSONSchemaEditorControl {
    a.Set("mini", value)
    return a
}

/**
 * horizontal
 */
func (a *JSONSchemaEditorControl) Horizontal(value interface{}) *JSONSchemaEditorControl {
    a.Set("horizontal", value)
    return a
}

/**
 * inputClassName
 */
func (a *JSONSchemaEditorControl) InputClassName(value interface{}) *JSONSchemaEditorControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * validationErrors
 */
func (a *JSONSchemaEditorControl) ValidationErrors(value interface{}) *JSONSchemaEditorControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * validateApi
 */
func (a *JSONSchemaEditorControl) ValidateApi(value interface{}) *JSONSchemaEditorControl {
    a.Set("validateApi", value)
    return a
}

/**
 * staticOn
 */
func (a *JSONSchemaEditorControl) StaticOn(value interface{}) *JSONSchemaEditorControl {
    a.Set("staticOn", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *JSONSchemaEditorControl) ReadOnlyOn(value interface{}) *JSONSchemaEditorControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *JSONSchemaEditorControl) ClearValueOnHidden(value interface{}) *JSONSchemaEditorControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * className
 */
func (a *JSONSchemaEditorControl) ClassName(value interface{}) *JSONSchemaEditorControl {
    a.Set("className", value)
    return a
}

/**
 * visible
 */
func (a *JSONSchemaEditorControl) Visible(value interface{}) *JSONSchemaEditorControl {
    a.Set("visible", value)
    return a
}

/**
 * disabledTypes
 */
func (a *JSONSchemaEditorControl) DisabledTypes(value interface{}) *JSONSchemaEditorControl {
    a.Set("disabledTypes", value)
    return a
}

/**
 * advancedSettings
 */
func (a *JSONSchemaEditorControl) AdvancedSettings(value interface{}) *JSONSchemaEditorControl {
    a.Set("advancedSettings", value)
    return a
}

/**
 * labelRemark
 */
func (a *JSONSchemaEditorControl) LabelRemark(value interface{}) *JSONSchemaEditorControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * submitOnChange
 */
func (a *JSONSchemaEditorControl) SubmitOnChange(value interface{}) *JSONSchemaEditorControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * mode
 */
func (a *JSONSchemaEditorControl) Mode(value interface{}) *JSONSchemaEditorControl {
    a.Set("mode", value)
    return a
}

/**
 * value
 */
func (a *JSONSchemaEditorControl) Value(value interface{}) *JSONSchemaEditorControl {
    a.Set("value", value)
    return a
}

/**
 * hidden
 */
func (a *JSONSchemaEditorControl) Hidden(value interface{}) *JSONSchemaEditorControl {
    a.Set("hidden", value)
    return a
}

/**
 * onEvent
 */
func (a *JSONSchemaEditorControl) OnEvent(value interface{}) *JSONSchemaEditorControl {
    a.Set("onEvent", value)
    return a
}

/**
 * rootTypeMutable
 */
func (a *JSONSchemaEditorControl) RootTypeMutable(value interface{}) *JSONSchemaEditorControl {
    a.Set("rootTypeMutable", value)
    return a
}

/**
 * labelOverflow
 */
func (a *JSONSchemaEditorControl) LabelOverflow(value interface{}) *JSONSchemaEditorControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * hint
 */
func (a *JSONSchemaEditorControl) Hint(value interface{}) *JSONSchemaEditorControl {
    a.Set("hint", value)
    return a
}

/**
 * validations
 */
func (a *JSONSchemaEditorControl) Validations(value interface{}) *JSONSchemaEditorControl {
    a.Set("validations", value)
    return a
}

/**
 * initAutoFill
 */
func (a *JSONSchemaEditorControl) InitAutoFill(value interface{}) *JSONSchemaEditorControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * static
 */
func (a *JSONSchemaEditorControl) Static(value interface{}) *JSONSchemaEditorControl {
    a.Set("static", value)
    return a
}

/**
 */
func (a *JSONSchemaEditorControl) Type(value interface{}) *JSONSchemaEditorControl {
    a.Set("type", value)
    return a
}
