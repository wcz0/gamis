package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type Dialog struct {
	*BaseRenderer
}

func NewDialog() *Dialog {
    a := &Dialog{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "dialog")
    return a
}


func (a *Dialog) Set(name string, value interface{}) *Dialog {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * showCloseButton
 */
func (a *Dialog) ShowCloseButton(value interface{}) *Dialog {
    a.Set("showCloseButton", value)
    return a
}

/**
 * showLoading
 */
func (a *Dialog) ShowLoading(value interface{}) *Dialog {
    a.Set("showLoading", value)
    return a
}

/**
 * draggable
 */
func (a *Dialog) Draggable(value interface{}) *Dialog {
    a.Set("draggable", value)
    return a
}

/**
 * disabled
 */
func (a *Dialog) Disabled(value interface{}) *Dialog {
    a.Set("disabled", value)
    return a
}

/**
 * editorSetting
 */
func (a *Dialog) EditorSetting(value interface{}) *Dialog {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *Dialog) Type(value interface{}) *Dialog {
    a.Set("type", value)
    return a
}

/**
 * closeOnEsc
 */
func (a *Dialog) CloseOnEsc(value interface{}) *Dialog {
    a.Set("closeOnEsc", value)
    return a
}

/**
 * size
 */
func (a *Dialog) Size(value interface{}) *Dialog {
    a.Set("size", value)
    return a
}

/**
 * height
 */
func (a *Dialog) Height(value interface{}) *Dialog {
    a.Set("height", value)
    return a
}

/**
 * headerClassName
 */
func (a *Dialog) HeaderClassName(value interface{}) *Dialog {
    a.Set("headerClassName", value)
    return a
}

/**
 * allowFullscreen
 */
func (a *Dialog) AllowFullscreen(value interface{}) *Dialog {
    a.Set("allowFullscreen", value)
    return a
}

/**
 * disabledOn
 */
func (a *Dialog) DisabledOn(value interface{}) *Dialog {
    a.Set("disabledOn", value)
    return a
}

/**
 * hidden
 */
func (a *Dialog) Hidden(value interface{}) *Dialog {
    a.Set("hidden", value)
    return a
}

/**
 * id
 */
func (a *Dialog) Id(value interface{}) *Dialog {
    a.Set("id", value)
    return a
}

/**
 * onEvent
 */
func (a *Dialog) OnEvent(value interface{}) *Dialog {
    a.Set("onEvent", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Dialog) StaticLabelClassName(value interface{}) *Dialog {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *Dialog) StaticSchema(value interface{}) *Dialog {
    a.Set("staticSchema", value)
    return a
}

/**
 * actions
 */
func (a *Dialog) Actions(value interface{}) *Dialog {
    a.Set("actions", value)
    return a
}

/**
 * closeOnOutside
 */
func (a *Dialog) CloseOnOutside(value interface{}) *Dialog {
    a.Set("closeOnOutside", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Dialog) StaticPlaceholder(value interface{}) *Dialog {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticClassName
 */
func (a *Dialog) StaticClassName(value interface{}) *Dialog {
    a.Set("staticClassName", value)
    return a
}

/**
 * width
 */
func (a *Dialog) Width(value interface{}) *Dialog {
    a.Set("width", value)
    return a
}

/**
 * showErrorMsg
 */
func (a *Dialog) ShowErrorMsg(value interface{}) *Dialog {
    a.Set("showErrorMsg", value)
    return a
}

/**
 * static
 */
func (a *Dialog) Static(value interface{}) *Dialog {
    a.Set("static", value)
    return a
}

/**
 * staticOn
 */
func (a *Dialog) StaticOn(value interface{}) *Dialog {
    a.Set("staticOn", value)
    return a
}

/**
 * title
 */
func (a *Dialog) Title(value interface{}) *Dialog {
    a.Set("title", value)
    return a
}

/**
 * inputParams
 */
func (a *Dialog) InputParams(value interface{}) *Dialog {
    a.Set("inputParams", value)
    return a
}

/**
 * data
 */
func (a *Dialog) Data(value interface{}) *Dialog {
    a.Set("data", value)
    return a
}

/**
 * bodyClassName
 */
func (a *Dialog) BodyClassName(value interface{}) *Dialog {
    a.Set("bodyClassName", value)
    return a
}

/**
 * visibleOn
 */
func (a *Dialog) VisibleOn(value interface{}) *Dialog {
    a.Set("visibleOn", value)
    return a
}

/**
 * name
 */
func (a *Dialog) Name(value interface{}) *Dialog {
    a.Set("name", value)
    return a
}

/**
 * confirm
 */
func (a *Dialog) Confirm(value interface{}) *Dialog {
    a.Set("confirm", value)
    return a
}

/**
 * overlay
 */
func (a *Dialog) Overlay(value interface{}) *Dialog {
    a.Set("overlay", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Dialog) StaticInputClassName(value interface{}) *Dialog {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Dialog) UseMobileUI(value interface{}) *Dialog {
    a.Set("useMobileUI", value)
    return a
}

/**
 * testid
 */
func (a *Dialog) Testid(value interface{}) *Dialog {
    a.Set("testid", value)
    return a
}

/**
 * body
 */
func (a *Dialog) Body(value interface{}) *Dialog {
    a.Set("body", value)
    return a
}

/**
 * header
 */
func (a *Dialog) Header(value interface{}) *Dialog {
    a.Set("header", value)
    return a
}

/**
 * footer
 */
func (a *Dialog) Footer(value interface{}) *Dialog {
    a.Set("footer", value)
    return a
}

/**
 */
func (a *Dialog) DialogType(value interface{}) *Dialog {
    a.Set("dialogType", value)
    return a
}

/**
 * className
 */
func (a *Dialog) ClassName(value interface{}) *Dialog {
    a.Set("className", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Dialog) HiddenOn(value interface{}) *Dialog {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visible
 */
func (a *Dialog) Visible(value interface{}) *Dialog {
    a.Set("visible", value)
    return a
}

/**
 * style
 */
func (a *Dialog) Style(value interface{}) *Dialog {
    a.Set("style", value)
    return a
}
