package renderers


/**
 * Flex 布局 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/flex

 * @author wcz0
 * @version 6.2.2
 */
type Flex struct {
	*BaseRenderer
}

func NewFlex() *Flex {
    a := &Flex{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "flex")
    return a
}


func (a *Flex) Set(name string, value interface{}) *Flex {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * justify
 */
func (a *Flex) Justify(value interface{}) *Flex {
    a.Set("justify", value)
    return a
}

/**
 * id
 */
func (a *Flex) Id(value interface{}) *Flex {
    a.Set("id", value)
    return a
}

/**
 * editorSetting
 */
func (a *Flex) EditorSetting(value interface{}) *Flex {
    a.Set("editorSetting", value)
    return a
}

/**
 * alignItems
 */
func (a *Flex) AlignItems(value interface{}) *Flex {
    a.Set("alignItems", value)
    return a
}

/**
 * alignContent
 */
func (a *Flex) AlignContent(value interface{}) *Flex {
    a.Set("alignContent", value)
    return a
}

/**
 * direction
 */
func (a *Flex) Direction(value interface{}) *Flex {
    a.Set("direction", value)
    return a
}

/**
 * items
 */
func (a *Flex) Items(value interface{}) *Flex {
    a.Set("items", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Flex) StaticPlaceholder(value interface{}) *Flex {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticClassName
 */
func (a *Flex) StaticClassName(value interface{}) *Flex {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Flex) StaticLabelClassName(value interface{}) *Flex {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *Flex) StaticSchema(value interface{}) *Flex {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *Flex) Type(value interface{}) *Flex {
    a.Set("type", value)
    return a
}

/**
 * className
 */
func (a *Flex) ClassName(value interface{}) *Flex {
    a.Set("className", value)
    return a
}

/**
 * disabled
 */
func (a *Flex) Disabled(value interface{}) *Flex {
    a.Set("disabled", value)
    return a
}

/**
 * disabledOn
 */
func (a *Flex) DisabledOn(value interface{}) *Flex {
    a.Set("disabledOn", value)
    return a
}

/**
 * hidden
 */
func (a *Flex) Hidden(value interface{}) *Flex {
    a.Set("hidden", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Flex) HiddenOn(value interface{}) *Flex {
    a.Set("hiddenOn", value)
    return a
}

/**
 * onEvent
 */
func (a *Flex) OnEvent(value interface{}) *Flex {
    a.Set("onEvent", value)
    return a
}

/**
 * static
 */
func (a *Flex) Static(value interface{}) *Flex {
    a.Set("static", value)
    return a
}

/**
 * staticOn
 */
func (a *Flex) StaticOn(value interface{}) *Flex {
    a.Set("staticOn", value)
    return a
}

/**
 * visible
 */
func (a *Flex) Visible(value interface{}) *Flex {
    a.Set("visible", value)
    return a
}

/**
 * visibleOn
 */
func (a *Flex) VisibleOn(value interface{}) *Flex {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Flex) StaticInputClassName(value interface{}) *Flex {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * style
 */
func (a *Flex) Style(value interface{}) *Flex {
    a.Set("style", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Flex) UseMobileUI(value interface{}) *Flex {
    a.Set("useMobileUI", value)
    return a
}

/**
 * testid
 */
func (a *Flex) Testid(value interface{}) *Flex {
    a.Set("testid", value)
    return a
}
