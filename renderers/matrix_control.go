package renderers


/**
 * Matrix 选择控件。适合做权限勾选。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/matrix

 * @author wcz0
 * @version 6.2.2
 */
type MatrixControl struct {
	*BaseRenderer
}

func NewMatrixControl() *MatrixControl {
    a := &MatrixControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "matrix-checkboxes")
    return a
}


func (a *MatrixControl) Set(name string, value interface{}) *MatrixControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * editorSetting
 */
func (a *MatrixControl) EditorSetting(value interface{}) *MatrixControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *MatrixControl) ReadOnlyOn(value interface{}) *MatrixControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *MatrixControl) StaticInputClassName(value interface{}) *MatrixControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * name
 */
func (a *MatrixControl) Name(value interface{}) *MatrixControl {
    a.Set("name", value)
    return a
}

/**
 * labelOverflow
 */
func (a *MatrixControl) LabelOverflow(value interface{}) *MatrixControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *MatrixControl) ClearValueOnHidden(value interface{}) *MatrixControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * validateApi
 */
func (a *MatrixControl) ValidateApi(value interface{}) *MatrixControl {
    a.Set("validateApi", value)
    return a
}

/**
 * labelWidth
 */
func (a *MatrixControl) LabelWidth(value interface{}) *MatrixControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * source
 */
func (a *MatrixControl) Source(value interface{}) *MatrixControl {
    a.Set("source", value)
    return a
}

/**
 * rows
 */
func (a *MatrixControl) Rows(value interface{}) *MatrixControl {
    a.Set("rows", value)
    return a
}

/**
 * label
 */
func (a *MatrixControl) Label(value interface{}) *MatrixControl {
    a.Set("label", value)
    return a
}

/**
 * labelClassName
 */
func (a *MatrixControl) LabelClassName(value interface{}) *MatrixControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * hint
 */
func (a *MatrixControl) Hint(value interface{}) *MatrixControl {
    a.Set("hint", value)
    return a
}

/**
 * validateOnChange
 */
func (a *MatrixControl) ValidateOnChange(value interface{}) *MatrixControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * mode
 */
func (a *MatrixControl) Mode(value interface{}) *MatrixControl {
    a.Set("mode", value)
    return a
}

/**
 * required
 */
func (a *MatrixControl) Required(value interface{}) *MatrixControl {
    a.Set("required", value)
    return a
}

/**
 * row
 */
func (a *MatrixControl) Row(value interface{}) *MatrixControl {
    a.Set("row", value)
    return a
}

/**
 * style
 */
func (a *MatrixControl) Style(value interface{}) *MatrixControl {
    a.Set("style", value)
    return a
}

/**
 * inline
 */
func (a *MatrixControl) Inline(value interface{}) *MatrixControl {
    a.Set("inline", value)
    return a
}

/**
 * autoFill
 */
func (a *MatrixControl) AutoFill(value interface{}) *MatrixControl {
    a.Set("autoFill", value)
    return a
}

/**
 * initAutoFill
 */
func (a *MatrixControl) InitAutoFill(value interface{}) *MatrixControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * useMobileUI
 */
func (a *MatrixControl) UseMobileUI(value interface{}) *MatrixControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *MatrixControl) Width(value interface{}) *MatrixControl {
    a.Set("width", value)
    return a
}

/**
 * validationErrors
 */
func (a *MatrixControl) ValidationErrors(value interface{}) *MatrixControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * onEvent
 */
func (a *MatrixControl) OnEvent(value interface{}) *MatrixControl {
    a.Set("onEvent", value)
    return a
}

/**
 * className
 */
func (a *MatrixControl) ClassName(value interface{}) *MatrixControl {
    a.Set("className", value)
    return a
}

/**
 * disabledOn
 */
func (a *MatrixControl) DisabledOn(value interface{}) *MatrixControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * multiple
 */
func (a *MatrixControl) Multiple(value interface{}) *MatrixControl {
    a.Set("multiple", value)
    return a
}

/**
 * singleSelectMode
 */
func (a *MatrixControl) SingleSelectMode(value interface{}) *MatrixControl {
    a.Set("singleSelectMode", value)
    return a
}

/**
 * labelAlign
 */
func (a *MatrixControl) LabelAlign(value interface{}) *MatrixControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * submitOnChange
 */
func (a *MatrixControl) SubmitOnChange(value interface{}) *MatrixControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * placeholder
 */
func (a *MatrixControl) Placeholder(value interface{}) *MatrixControl {
    a.Set("placeholder", value)
    return a
}

/**
 * value
 */
func (a *MatrixControl) Value(value interface{}) *MatrixControl {
    a.Set("value", value)
    return a
}

/**
 * hiddenOn
 */
func (a *MatrixControl) HiddenOn(value interface{}) *MatrixControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * columns
 */
func (a *MatrixControl) Columns(value interface{}) *MatrixControl {
    a.Set("columns", value)
    return a
}

/**
 * readOnly
 */
func (a *MatrixControl) ReadOnly(value interface{}) *MatrixControl {
    a.Set("readOnly", value)
    return a
}

/**
 * horizontal
 */
func (a *MatrixControl) Horizontal(value interface{}) *MatrixControl {
    a.Set("horizontal", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *MatrixControl) StaticLabelClassName(value interface{}) *MatrixControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * labelRemark
 */
func (a *MatrixControl) LabelRemark(value interface{}) *MatrixControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * inputClassName
 */
func (a *MatrixControl) InputClassName(value interface{}) *MatrixControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * rowLabel
 */
func (a *MatrixControl) RowLabel(value interface{}) *MatrixControl {
    a.Set("rowLabel", value)
    return a
}

/**
 * remark
 */
func (a *MatrixControl) Remark(value interface{}) *MatrixControl {
    a.Set("remark", value)
    return a
}

/**
 * desc
 */
func (a *MatrixControl) Desc(value interface{}) *MatrixControl {
    a.Set("desc", value)
    return a
}

/**
 * visibleOn
 */
func (a *MatrixControl) VisibleOn(value interface{}) *MatrixControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticOn
 */
func (a *MatrixControl) StaticOn(value interface{}) *MatrixControl {
    a.Set("staticOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *MatrixControl) StaticClassName(value interface{}) *MatrixControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * extraName
 */
func (a *MatrixControl) ExtraName(value interface{}) *MatrixControl {
    a.Set("extraName", value)
    return a
}

/**
 * validations
 */
func (a *MatrixControl) Validations(value interface{}) *MatrixControl {
    a.Set("validations", value)
    return a
}

/**
 * disabled
 */
func (a *MatrixControl) Disabled(value interface{}) *MatrixControl {
    a.Set("disabled", value)
    return a
}

/**
 * size
 */
func (a *MatrixControl) Size(value interface{}) *MatrixControl {
    a.Set("size", value)
    return a
}

/**
 * id
 */
func (a *MatrixControl) Id(value interface{}) *MatrixControl {
    a.Set("id", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *MatrixControl) StaticPlaceholder(value interface{}) *MatrixControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *MatrixControl) Type(value interface{}) *MatrixControl {
    a.Set("type", value)
    return a
}

/**
 * description
 */
func (a *MatrixControl) Description(value interface{}) *MatrixControl {
    a.Set("description", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *MatrixControl) DescriptionClassName(value interface{}) *MatrixControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * hidden
 */
func (a *MatrixControl) Hidden(value interface{}) *MatrixControl {
    a.Set("hidden", value)
    return a
}

/**
 * visible
 */
func (a *MatrixControl) Visible(value interface{}) *MatrixControl {
    a.Set("visible", value)
    return a
}

/**
 * static
 */
func (a *MatrixControl) Static(value interface{}) *MatrixControl {
    a.Set("static", value)
    return a
}

/**
 * staticSchema
 */
func (a *MatrixControl) StaticSchema(value interface{}) *MatrixControl {
    a.Set("staticSchema", value)
    return a
}
