package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type OtherAction struct {
	*BaseRenderer
}

func NewOtherAction() *OtherAction {
    a := &OtherAction{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("actionType", "prev")
    return a
}


func (a *OtherAction) Set(name string, value interface{}) *OtherAction {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * rightIconClassName
 */
func (a *OtherAction) RightIconClassName(value interface{}) *OtherAction {
    a.Set("rightIconClassName", value)
    return a
}

/**
 * required
 */
func (a *OtherAction) Required(value interface{}) *OtherAction {
    a.Set("required", value)
    return a
}

/**
 * requireSelected
 */
func (a *OtherAction) RequireSelected(value interface{}) *OtherAction {
    a.Set("requireSelected", value)
    return a
}

/**
 * disabledOnAction
 */
func (a *OtherAction) DisabledOnAction(value interface{}) *OtherAction {
    a.Set("disabledOnAction", value)
    return a
}

/**
 * disabled
 */
func (a *OtherAction) Disabled(value interface{}) *OtherAction {
    a.Set("disabled", value)
    return a
}

/**
 * staticClassName
 */
func (a *OtherAction) StaticClassName(value interface{}) *OtherAction {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *OtherAction) StaticLabelClassName(value interface{}) *OtherAction {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *OtherAction) StaticInputClassName(value interface{}) *OtherAction {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * useMobileUI
 */
func (a *OtherAction) UseMobileUI(value interface{}) *OtherAction {
    a.Set("useMobileUI", value)
    return a
}

/**
 * size
 */
func (a *OtherAction) Size(value interface{}) *OtherAction {
    a.Set("size", value)
    return a
}

/**
 * target
 */
func (a *OtherAction) Target(value interface{}) *OtherAction {
    a.Set("target", value)
    return a
}

/**
 * onEvent
 */
func (a *OtherAction) OnEvent(value interface{}) *OtherAction {
    a.Set("onEvent", value)
    return a
}

/**
 * disabledTip
 */
func (a *OtherAction) DisabledTip(value interface{}) *OtherAction {
    a.Set("disabledTip", value)
    return a
}

/**
 * activeClassName
 */
func (a *OtherAction) ActiveClassName(value interface{}) *OtherAction {
    a.Set("activeClassName", value)
    return a
}

/**
 * countDownTpl
 */
func (a *OtherAction) CountDownTpl(value interface{}) *OtherAction {
    a.Set("countDownTpl", value)
    return a
}

/**
 * badge
 */
func (a *OtherAction) Badge(value interface{}) *OtherAction {
    a.Set("badge", value)
    return a
}

/**
 * onClick
 */
func (a *OtherAction) OnClick(value interface{}) *OtherAction {
    a.Set("onClick", value)
    return a
}

/**
 * href
 */
func (a *OtherAction) Href(value interface{}) *OtherAction {
    a.Set("href", value)
    return a
}

/**
 * style
 */
func (a *OtherAction) Style(value interface{}) *OtherAction {
    a.Set("style", value)
    return a
}

/**
 * testid
 */
func (a *OtherAction) Testid(value interface{}) *OtherAction {
    a.Set("testid", value)
    return a
}

/**
 * label
 */
func (a *OtherAction) Label(value interface{}) *OtherAction {
    a.Set("label", value)
    return a
}

/**
 * primary
 */
func (a *OtherAction) Primary(value interface{}) *OtherAction {
    a.Set("primary", value)
    return a
}

/**
 * tooltipPlacement
 */
func (a *OtherAction) TooltipPlacement(value interface{}) *OtherAction {
    a.Set("tooltipPlacement", value)
    return a
}

/**
 * confirmText
 */
func (a *OtherAction) ConfirmText(value interface{}) *OtherAction {
    a.Set("confirmText", value)
    return a
}

/**
 * close
 */
func (a *OtherAction) Close(value interface{}) *OtherAction {
    a.Set("close", value)
    return a
}

/**
 * mergeData
 */
func (a *OtherAction) MergeData(value interface{}) *OtherAction {
    a.Set("mergeData", value)
    return a
}

/**
 * hidden
 */
func (a *OtherAction) Hidden(value interface{}) *OtherAction {
    a.Set("hidden", value)
    return a
}

/**
 * visible
 */
func (a *OtherAction) Visible(value interface{}) *OtherAction {
    a.Set("visible", value)
    return a
}

/**
 * staticOn
 */
func (a *OtherAction) StaticOn(value interface{}) *OtherAction {
    a.Set("staticOn", value)
    return a
}

/**
 * editorSetting
 */
func (a *OtherAction) EditorSetting(value interface{}) *OtherAction {
    a.Set("editorSetting", value)
    return a
}

/**
 * block
 */
func (a *OtherAction) Block(value interface{}) *OtherAction {
    a.Set("block", value)
    return a
}

/**
 * tooltip
 */
func (a *OtherAction) Tooltip(value interface{}) *OtherAction {
    a.Set("tooltip", value)
    return a
}

/**
 * countDown
 */
func (a *OtherAction) CountDown(value interface{}) *OtherAction {
    a.Set("countDown", value)
    return a
}

/**
 * loadingOn
 */
func (a *OtherAction) LoadingOn(value interface{}) *OtherAction {
    a.Set("loadingOn", value)
    return a
}

/**
 * loadingClassName
 */
func (a *OtherAction) LoadingClassName(value interface{}) *OtherAction {
    a.Set("loadingClassName", value)
    return a
}

/**
 * level
 */
func (a *OtherAction) Level(value interface{}) *OtherAction {
    a.Set("level", value)
    return a
}

/**
 * hotKey
 */
func (a *OtherAction) HotKey(value interface{}) *OtherAction {
    a.Set("hotKey", value)
    return a
}

/**
 * body
 */
func (a *OtherAction) Body(value interface{}) *OtherAction {
    a.Set("body", value)
    return a
}

/**
 * 可选值: prev | next | cancel | close | submit | confirm | add | reset | reset-and-submit
 */
func (a *OtherAction) ActionType(value interface{}) *OtherAction {
    a.Set("actionType", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *OtherAction) StaticPlaceholder(value interface{}) *OtherAction {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * className
 */
func (a *OtherAction) ClassName(value interface{}) *OtherAction {
    a.Set("className", value)
    return a
}

/**
 * staticSchema
 */
func (a *OtherAction) StaticSchema(value interface{}) *OtherAction {
    a.Set("staticSchema", value)
    return a
}

/**
 * 可选值: action | button | submit | reset
 */
func (a *OtherAction) Type(value interface{}) *OtherAction {
    a.Set("type", value)
    return a
}

/**
 * activeLevel
 */
func (a *OtherAction) ActiveLevel(value interface{}) *OtherAction {
    a.Set("activeLevel", value)
    return a
}

/**
 * tabIndex
 */
func (a *OtherAction) TabIndex(value interface{}) *OtherAction {
    a.Set("tabIndex", value)
    return a
}

/**
 * disabledOn
 */
func (a *OtherAction) DisabledOn(value interface{}) *OtherAction {
    a.Set("disabledOn", value)
    return a
}

/**
 * id
 */
func (a *OtherAction) Id(value interface{}) *OtherAction {
    a.Set("id", value)
    return a
}

/**
 * icon
 */
func (a *OtherAction) Icon(value interface{}) *OtherAction {
    a.Set("icon", value)
    return a
}

/**
 * rightIcon
 */
func (a *OtherAction) RightIcon(value interface{}) *OtherAction {
    a.Set("rightIcon", value)
    return a
}

/**
 * iconClassName
 */
func (a *OtherAction) IconClassName(value interface{}) *OtherAction {
    a.Set("iconClassName", value)
    return a
}

/**
 * hiddenOn
 */
func (a *OtherAction) HiddenOn(value interface{}) *OtherAction {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visibleOn
 */
func (a *OtherAction) VisibleOn(value interface{}) *OtherAction {
    a.Set("visibleOn", value)
    return a
}

/**
 * static
 */
func (a *OtherAction) Static(value interface{}) *OtherAction {
    a.Set("static", value)
    return a
}
