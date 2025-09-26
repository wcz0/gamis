package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type Drawer struct {
	*BaseRenderer
}

func NewDrawer() *Drawer {
    a := &Drawer{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "drawer")
    return a
}


func (a *Drawer) Set(name string, value interface{}) *Drawer {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * resizable
 */
func (a *Drawer) Resizable(value interface{}) *Drawer {
    a.Set("resizable", value)
    return a
}

/**
 * visible
 */
func (a *Drawer) Visible(value interface{}) *Drawer {
    a.Set("visible", value)
    return a
}

/**
 * id
 */
func (a *Drawer) Id(value interface{}) *Drawer {
    a.Set("id", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Drawer) StaticLabelClassName(value interface{}) *Drawer {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *Drawer) Type(value interface{}) *Drawer {
    a.Set("type", value)
    return a
}

/**
 * body
 */
func (a *Drawer) Body(value interface{}) *Drawer {
    a.Set("body", value)
    return a
}

/**
 * bodyClassName
 */
func (a *Drawer) BodyClassName(value interface{}) *Drawer {
    a.Set("bodyClassName", value)
    return a
}

/**
 * closeOnEsc
 */
func (a *Drawer) CloseOnEsc(value interface{}) *Drawer {
    a.Set("closeOnEsc", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Drawer) HiddenOn(value interface{}) *Drawer {
    a.Set("hiddenOn", value)
    return a
}

/**
 * editorSetting
 */
func (a *Drawer) EditorSetting(value interface{}) *Drawer {
    a.Set("editorSetting", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Drawer) UseMobileUI(value interface{}) *Drawer {
    a.Set("useMobileUI", value)
    return a
}

/**
 * header
 */
func (a *Drawer) Header(value interface{}) *Drawer {
    a.Set("header", value)
    return a
}

/**
 * testid
 */
func (a *Drawer) Testid(value interface{}) *Drawer {
    a.Set("testid", value)
    return a
}

/**
 * title
 */
func (a *Drawer) Title(value interface{}) *Drawer {
    a.Set("title", value)
    return a
}

/**
 * closeOnOutside
 */
func (a *Drawer) CloseOnOutside(value interface{}) *Drawer {
    a.Set("closeOnOutside", value)
    return a
}

/**
 * staticOn
 */
func (a *Drawer) StaticOn(value interface{}) *Drawer {
    a.Set("staticOn", value)
    return a
}

/**
 * staticSchema
 */
func (a *Drawer) StaticSchema(value interface{}) *Drawer {
    a.Set("staticSchema", value)
    return a
}

/**
 * style
 */
func (a *Drawer) Style(value interface{}) *Drawer {
    a.Set("style", value)
    return a
}

/**
 * footerClassName
 */
func (a *Drawer) FooterClassName(value interface{}) *Drawer {
    a.Set("footerClassName", value)
    return a
}

/**
 * size
 */
func (a *Drawer) Size(value interface{}) *Drawer {
    a.Set("size", value)
    return a
}

/**
 * showCloseButton
 */
func (a *Drawer) ShowCloseButton(value interface{}) *Drawer {
    a.Set("showCloseButton", value)
    return a
}

/**
 * footer
 */
func (a *Drawer) Footer(value interface{}) *Drawer {
    a.Set("footer", value)
    return a
}

/**
 * static
 */
func (a *Drawer) Static(value interface{}) *Drawer {
    a.Set("static", value)
    return a
}

/**
 * actions
 */
func (a *Drawer) Actions(value interface{}) *Drawer {
    a.Set("actions", value)
    return a
}

/**
 * name
 */
func (a *Drawer) Name(value interface{}) *Drawer {
    a.Set("name", value)
    return a
}

/**
 * overlay
 */
func (a *Drawer) Overlay(value interface{}) *Drawer {
    a.Set("overlay", value)
    return a
}

/**
 * showErrorMsg
 */
func (a *Drawer) ShowErrorMsg(value interface{}) *Drawer {
    a.Set("showErrorMsg", value)
    return a
}

/**
 * data
 */
func (a *Drawer) Data(value interface{}) *Drawer {
    a.Set("data", value)
    return a
}

/**
 * hidden
 */
func (a *Drawer) Hidden(value interface{}) *Drawer {
    a.Set("hidden", value)
    return a
}

/**
 * staticClassName
 */
func (a *Drawer) StaticClassName(value interface{}) *Drawer {
    a.Set("staticClassName", value)
    return a
}

/**
 * headerClassName
 */
func (a *Drawer) HeaderClassName(value interface{}) *Drawer {
    a.Set("headerClassName", value)
    return a
}

/**
 * position
 */
func (a *Drawer) Position(value interface{}) *Drawer {
    a.Set("position", value)
    return a
}

/**
 * height
 */
func (a *Drawer) Height(value interface{}) *Drawer {
    a.Set("height", value)
    return a
}

/**
 * className
 */
func (a *Drawer) ClassName(value interface{}) *Drawer {
    a.Set("className", value)
    return a
}

/**
 * disabledOn
 */
func (a *Drawer) DisabledOn(value interface{}) *Drawer {
    a.Set("disabledOn", value)
    return a
}

/**
 * visibleOn
 */
func (a *Drawer) VisibleOn(value interface{}) *Drawer {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Drawer) StaticInputClassName(value interface{}) *Drawer {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * inputParams
 */
func (a *Drawer) InputParams(value interface{}) *Drawer {
    a.Set("inputParams", value)
    return a
}

/**
 * width
 */
func (a *Drawer) Width(value interface{}) *Drawer {
    a.Set("width", value)
    return a
}

/**
 * disabled
 */
func (a *Drawer) Disabled(value interface{}) *Drawer {
    a.Set("disabled", value)
    return a
}

/**
 * onEvent
 */
func (a *Drawer) OnEvent(value interface{}) *Drawer {
    a.Set("onEvent", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Drawer) StaticPlaceholder(value interface{}) *Drawer {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * confirm
 */
func (a *Drawer) Confirm(value interface{}) *Drawer {
    a.Set("confirm", value)
    return a
}
