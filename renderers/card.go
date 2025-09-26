package renderers


/**
 * Card 卡片渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/card

 * @author wcz0
 * @version 6.2.2
 */
type Card struct {
	*BaseRenderer
}

func NewCard() *Card {
    a := &Card{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "card")
    return a
}


func (a *Card) Set(name string, value interface{}) *Card {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * disabled
 */
func (a *Card) Disabled(value interface{}) *Card {
    a.Set("disabled", value)
    return a
}

/**
 * style
 */
func (a *Card) Style(value interface{}) *Card {
    a.Set("style", value)
    return a
}

/**
 * toolbar
 */
func (a *Card) Toolbar(value interface{}) *Card {
    a.Set("toolbar", value)
    return a
}

/**
 * hidden
 */
func (a *Card) Hidden(value interface{}) *Card {
    a.Set("hidden", value)
    return a
}

/**
 * staticOn
 */
func (a *Card) StaticOn(value interface{}) *Card {
    a.Set("staticOn", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Card) StaticLabelClassName(value interface{}) *Card {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Card) UseMobileUI(value interface{}) *Card {
    a.Set("useMobileUI", value)
    return a
}

/**
 * actions
 */
func (a *Card) Actions(value interface{}) *Card {
    a.Set("actions", value)
    return a
}

/**
 * visible
 */
func (a *Card) Visible(value interface{}) *Card {
    a.Set("visible", value)
    return a
}

/**
 * id
 */
func (a *Card) Id(value interface{}) *Card {
    a.Set("id", value)
    return a
}

/**
 * static
 */
func (a *Card) Static(value interface{}) *Card {
    a.Set("static", value)
    return a
}

/**
 * body
 */
func (a *Card) Body(value interface{}) *Card {
    a.Set("body", value)
    return a
}

/**
 * secondary
 */
func (a *Card) Secondary(value interface{}) *Card {
    a.Set("secondary", value)
    return a
}

/**
 * useCardLabel
 */
func (a *Card) UseCardLabel(value interface{}) *Card {
    a.Set("useCardLabel", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Card) StaticPlaceholder(value interface{}) *Card {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Card) StaticInputClassName(value interface{}) *Card {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *Card) StaticSchema(value interface{}) *Card {
    a.Set("staticSchema", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Card) HiddenOn(value interface{}) *Card {
    a.Set("hiddenOn", value)
    return a
}

/**
 * testid
 */
func (a *Card) Testid(value interface{}) *Card {
    a.Set("testid", value)
    return a
}

/**
 * media
 */
func (a *Card) Media(value interface{}) *Card {
    a.Set("media", value)
    return a
}

/**
 * disabledOn
 */
func (a *Card) DisabledOn(value interface{}) *Card {
    a.Set("disabledOn", value)
    return a
}

/**
 */
func (a *Card) Type(value interface{}) *Card {
    a.Set("type", value)
    return a
}

/**
 * visibleOn
 */
func (a *Card) VisibleOn(value interface{}) *Card {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *Card) StaticClassName(value interface{}) *Card {
    a.Set("staticClassName", value)
    return a
}

/**
 * header
 */
func (a *Card) Header(value interface{}) *Card {
    a.Set("header", value)
    return a
}

/**
 * className
 */
func (a *Card) ClassName(value interface{}) *Card {
    a.Set("className", value)
    return a
}

/**
 * onEvent
 */
func (a *Card) OnEvent(value interface{}) *Card {
    a.Set("onEvent", value)
    return a
}

/**
 * editorSetting
 */
func (a *Card) EditorSetting(value interface{}) *Card {
    a.Set("editorSetting", value)
    return a
}
