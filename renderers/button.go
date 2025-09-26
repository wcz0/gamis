package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type Button struct {
	*BaseRenderer
}

func NewButton() *Button {
    a := &Button{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "action")
    return a
}


func (a *Button) Set(name string, value interface{}) *Button {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * label
 */
func (a *Button) Label(value interface{}) *Button {
    a.Set("label", value)
    return a
}

/**
 * close
 */
func (a *Button) Close(value interface{}) *Button {
    a.Set("close", value)
    return a
}

/**
 * requireSelected
 */
func (a *Button) RequireSelected(value interface{}) *Button {
    a.Set("requireSelected", value)
    return a
}

/**
 * href
 */
func (a *Button) Href(value interface{}) *Button {
    a.Set("href", value)
    return a
}

/**
 * staticOn
 */
func (a *Button) StaticOn(value interface{}) *Button {
    a.Set("staticOn", value)
    return a
}

/**
 * iconClassName
 */
func (a *Button) IconClassName(value interface{}) *Button {
    a.Set("iconClassName", value)
    return a
}

/**
 * tooltipPlacement
 */
func (a *Button) TooltipPlacement(value interface{}) *Button {
    a.Set("tooltipPlacement", value)
    return a
}

/**
 * confirmText
 */
func (a *Button) ConfirmText(value interface{}) *Button {
    a.Set("confirmText", value)
    return a
}

/**
 * disabled
 */
func (a *Button) Disabled(value interface{}) *Button {
    a.Set("disabled", value)
    return a
}

/**
 * staticClassName
 */
func (a *Button) StaticClassName(value interface{}) *Button {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Button) StaticInputClassName(value interface{}) *Button {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * editorSetting
 */
func (a *Button) EditorSetting(value interface{}) *Button {
    a.Set("editorSetting", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Button) UseMobileUI(value interface{}) *Button {
    a.Set("useMobileUI", value)
    return a
}

/**
 * rightIconClassName
 */
func (a *Button) RightIconClassName(value interface{}) *Button {
    a.Set("rightIconClassName", value)
    return a
}

/**
 * loadingClassName
 */
func (a *Button) LoadingClassName(value interface{}) *Button {
    a.Set("loadingClassName", value)
    return a
}

/**
 * level
 */
func (a *Button) Level(value interface{}) *Button {
    a.Set("level", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Button) StaticPlaceholder(value interface{}) *Button {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 可选值: action | button | submit | reset
 */
func (a *Button) Type(value interface{}) *Button {
    a.Set("type", value)
    return a
}

/**
 * rightIcon
 */
func (a *Button) RightIcon(value interface{}) *Button {
    a.Set("rightIcon", value)
    return a
}

/**
 * primary
 */
func (a *Button) Primary(value interface{}) *Button {
    a.Set("primary", value)
    return a
}

/**
 * required
 */
func (a *Button) Required(value interface{}) *Button {
    a.Set("required", value)
    return a
}

/**
 * hotKey
 */
func (a *Button) HotKey(value interface{}) *Button {
    a.Set("hotKey", value)
    return a
}

/**
 * disabledOnAction
 */
func (a *Button) DisabledOnAction(value interface{}) *Button {
    a.Set("disabledOnAction", value)
    return a
}

/**
 * visibleOn
 */
func (a *Button) VisibleOn(value interface{}) *Button {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticSchema
 */
func (a *Button) StaticSchema(value interface{}) *Button {
    a.Set("staticSchema", value)
    return a
}

/**
 * mergeData
 */
func (a *Button) MergeData(value interface{}) *Button {
    a.Set("mergeData", value)
    return a
}

/**
 * countDown
 */
func (a *Button) CountDown(value interface{}) *Button {
    a.Set("countDown", value)
    return a
}

/**
 * static
 */
func (a *Button) Static(value interface{}) *Button {
    a.Set("static", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Button) StaticLabelClassName(value interface{}) *Button {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * icon
 */
func (a *Button) Icon(value interface{}) *Button {
    a.Set("icon", value)
    return a
}

/**
 * tooltip
 */
func (a *Button) Tooltip(value interface{}) *Button {
    a.Set("tooltip", value)
    return a
}

/**
 * activeClassName
 */
func (a *Button) ActiveClassName(value interface{}) *Button {
    a.Set("activeClassName", value)
    return a
}

/**
 * badge
 */
func (a *Button) Badge(value interface{}) *Button {
    a.Set("badge", value)
    return a
}

/**
 * onClick
 */
func (a *Button) OnClick(value interface{}) *Button {
    a.Set("onClick", value)
    return a
}

/**
 * body
 */
func (a *Button) Body(value interface{}) *Button {
    a.Set("body", value)
    return a
}

/**
 * disabledOn
 */
func (a *Button) DisabledOn(value interface{}) *Button {
    a.Set("disabledOn", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Button) HiddenOn(value interface{}) *Button {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visible
 */
func (a *Button) Visible(value interface{}) *Button {
    a.Set("visible", value)
    return a
}

/**
 * target
 */
func (a *Button) Target(value interface{}) *Button {
    a.Set("target", value)
    return a
}

/**
 * loadingOn
 */
func (a *Button) LoadingOn(value interface{}) *Button {
    a.Set("loadingOn", value)
    return a
}

/**
 * tabIndex
 */
func (a *Button) TabIndex(value interface{}) *Button {
    a.Set("tabIndex", value)
    return a
}

/**
 * hidden
 */
func (a *Button) Hidden(value interface{}) *Button {
    a.Set("hidden", value)
    return a
}

/**
 * onEvent
 */
func (a *Button) OnEvent(value interface{}) *Button {
    a.Set("onEvent", value)
    return a
}

/**
 * testid
 */
func (a *Button) Testid(value interface{}) *Button {
    a.Set("testid", value)
    return a
}

/**
 * disabledTip
 */
func (a *Button) DisabledTip(value interface{}) *Button {
    a.Set("disabledTip", value)
    return a
}

/**
 * size
 */
func (a *Button) Size(value interface{}) *Button {
    a.Set("size", value)
    return a
}

/**
 * activeLevel
 */
func (a *Button) ActiveLevel(value interface{}) *Button {
    a.Set("activeLevel", value)
    return a
}

/**
 * countDownTpl
 */
func (a *Button) CountDownTpl(value interface{}) *Button {
    a.Set("countDownTpl", value)
    return a
}

/**
 * className
 */
func (a *Button) ClassName(value interface{}) *Button {
    a.Set("className", value)
    return a
}

/**
 * id
 */
func (a *Button) Id(value interface{}) *Button {
    a.Set("id", value)
    return a
}

/**
 * style
 */
func (a *Button) Style(value interface{}) *Button {
    a.Set("style", value)
    return a
}

/**
 * block
 */
func (a *Button) Block(value interface{}) *Button {
    a.Set("block", value)
    return a
}
