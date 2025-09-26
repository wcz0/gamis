package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type EmailAction struct {
	*BaseRenderer
}

func NewEmailAction() *EmailAction {
    a := &EmailAction{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("actionType", "email")
    return a
}


func (a *EmailAction) Set(name string, value interface{}) *EmailAction {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * rightIcon
 */
func (a *EmailAction) RightIcon(value interface{}) *EmailAction {
    a.Set("rightIcon", value)
    return a
}

/**
 * className
 */
func (a *EmailAction) ClassName(value interface{}) *EmailAction {
    a.Set("className", value)
    return a
}

/**
 * static
 */
func (a *EmailAction) Static(value interface{}) *EmailAction {
    a.Set("static", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *EmailAction) StaticLabelClassName(value interface{}) *EmailAction {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * block
 */
func (a *EmailAction) Block(value interface{}) *EmailAction {
    a.Set("block", value)
    return a
}

/**
 * staticClassName
 */
func (a *EmailAction) StaticClassName(value interface{}) *EmailAction {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *EmailAction) StaticSchema(value interface{}) *EmailAction {
    a.Set("staticSchema", value)
    return a
}

/**
 * testid
 */
func (a *EmailAction) Testid(value interface{}) *EmailAction {
    a.Set("testid", value)
    return a
}

/**
 * hotKey
 */
func (a *EmailAction) HotKey(value interface{}) *EmailAction {
    a.Set("hotKey", value)
    return a
}

/**
 * disabledOn
 */
func (a *EmailAction) DisabledOn(value interface{}) *EmailAction {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticOn
 */
func (a *EmailAction) StaticOn(value interface{}) *EmailAction {
    a.Set("staticOn", value)
    return a
}

/**
 * confirmText
 */
func (a *EmailAction) ConfirmText(value interface{}) *EmailAction {
    a.Set("confirmText", value)
    return a
}

/**
 * countDown
 */
func (a *EmailAction) CountDown(value interface{}) *EmailAction {
    a.Set("countDown", value)
    return a
}

/**
 * required
 */
func (a *EmailAction) Required(value interface{}) *EmailAction {
    a.Set("required", value)
    return a
}

/**
 * visible
 */
func (a *EmailAction) Visible(value interface{}) *EmailAction {
    a.Set("visible", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *EmailAction) StaticPlaceholder(value interface{}) *EmailAction {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * editorSetting
 */
func (a *EmailAction) EditorSetting(value interface{}) *EmailAction {
    a.Set("editorSetting", value)
    return a
}

/**
 * label
 */
func (a *EmailAction) Label(value interface{}) *EmailAction {
    a.Set("label", value)
    return a
}

/**
 * hiddenOn
 */
func (a *EmailAction) HiddenOn(value interface{}) *EmailAction {
    a.Set("hiddenOn", value)
    return a
}

/**
 * icon
 */
func (a *EmailAction) Icon(value interface{}) *EmailAction {
    a.Set("icon", value)
    return a
}

/**
 * body
 */
func (a *EmailAction) Body(value interface{}) *EmailAction {
    a.Set("body", value)
    return a
}

/**
 * rightIconClassName
 */
func (a *EmailAction) RightIconClassName(value interface{}) *EmailAction {
    a.Set("rightIconClassName", value)
    return a
}

/**
 * to
 */
func (a *EmailAction) To(value interface{}) *EmailAction {
    a.Set("to", value)
    return a
}

/**
 * disabled
 */
func (a *EmailAction) Disabled(value interface{}) *EmailAction {
    a.Set("disabled", value)
    return a
}

/**
 * close
 */
func (a *EmailAction) Close(value interface{}) *EmailAction {
    a.Set("close", value)
    return a
}

/**
 * useMobileUI
 */
func (a *EmailAction) UseMobileUI(value interface{}) *EmailAction {
    a.Set("useMobileUI", value)
    return a
}

/**
 * primary
 */
func (a *EmailAction) Primary(value interface{}) *EmailAction {
    a.Set("primary", value)
    return a
}

/**
 * activeClassName
 */
func (a *EmailAction) ActiveClassName(value interface{}) *EmailAction {
    a.Set("activeClassName", value)
    return a
}

/**
 * loadingOn
 */
func (a *EmailAction) LoadingOn(value interface{}) *EmailAction {
    a.Set("loadingOn", value)
    return a
}

/**
 * onClick
 */
func (a *EmailAction) OnClick(value interface{}) *EmailAction {
    a.Set("onClick", value)
    return a
}

/**
 * cc
 */
func (a *EmailAction) Cc(value interface{}) *EmailAction {
    a.Set("cc", value)
    return a
}

/**
 * 可选值: action | button | submit | reset
 */
func (a *EmailAction) Type(value interface{}) *EmailAction {
    a.Set("type", value)
    return a
}

/**
 * disabledOnAction
 */
func (a *EmailAction) DisabledOnAction(value interface{}) *EmailAction {
    a.Set("disabledOnAction", value)
    return a
}

/**
 * tabIndex
 */
func (a *EmailAction) TabIndex(value interface{}) *EmailAction {
    a.Set("tabIndex", value)
    return a
}

/**
 * subject
 */
func (a *EmailAction) Subject(value interface{}) *EmailAction {
    a.Set("subject", value)
    return a
}

/**
 * visibleOn
 */
func (a *EmailAction) VisibleOn(value interface{}) *EmailAction {
    a.Set("visibleOn", value)
    return a
}

/**
 * size
 */
func (a *EmailAction) Size(value interface{}) *EmailAction {
    a.Set("size", value)
    return a
}

/**
 * activeLevel
 */
func (a *EmailAction) ActiveLevel(value interface{}) *EmailAction {
    a.Set("activeLevel", value)
    return a
}

/**
 * mergeData
 */
func (a *EmailAction) MergeData(value interface{}) *EmailAction {
    a.Set("mergeData", value)
    return a
}

/**
 * href
 */
func (a *EmailAction) Href(value interface{}) *EmailAction {
    a.Set("href", value)
    return a
}

/**
 * target
 */
func (a *EmailAction) Target(value interface{}) *EmailAction {
    a.Set("target", value)
    return a
}

/**
 * countDownTpl
 */
func (a *EmailAction) CountDownTpl(value interface{}) *EmailAction {
    a.Set("countDownTpl", value)
    return a
}

/**
 */
func (a *EmailAction) ActionType(value interface{}) *EmailAction {
    a.Set("actionType", value)
    return a
}

/**
 * iconClassName
 */
func (a *EmailAction) IconClassName(value interface{}) *EmailAction {
    a.Set("iconClassName", value)
    return a
}

/**
 * onEvent
 */
func (a *EmailAction) OnEvent(value interface{}) *EmailAction {
    a.Set("onEvent", value)
    return a
}

/**
 * tooltip
 */
func (a *EmailAction) Tooltip(value interface{}) *EmailAction {
    a.Set("tooltip", value)
    return a
}

/**
 * tooltipPlacement
 */
func (a *EmailAction) TooltipPlacement(value interface{}) *EmailAction {
    a.Set("tooltipPlacement", value)
    return a
}

/**
 * hidden
 */
func (a *EmailAction) Hidden(value interface{}) *EmailAction {
    a.Set("hidden", value)
    return a
}

/**
 * id
 */
func (a *EmailAction) Id(value interface{}) *EmailAction {
    a.Set("id", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *EmailAction) StaticInputClassName(value interface{}) *EmailAction {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * style
 */
func (a *EmailAction) Style(value interface{}) *EmailAction {
    a.Set("style", value)
    return a
}

/**
 * level
 */
func (a *EmailAction) Level(value interface{}) *EmailAction {
    a.Set("level", value)
    return a
}

/**
 * disabledTip
 */
func (a *EmailAction) DisabledTip(value interface{}) *EmailAction {
    a.Set("disabledTip", value)
    return a
}

/**
 * loadingClassName
 */
func (a *EmailAction) LoadingClassName(value interface{}) *EmailAction {
    a.Set("loadingClassName", value)
    return a
}

/**
 * requireSelected
 */
func (a *EmailAction) RequireSelected(value interface{}) *EmailAction {
    a.Set("requireSelected", value)
    return a
}

/**
 * badge
 */
func (a *EmailAction) Badge(value interface{}) *EmailAction {
    a.Set("badge", value)
    return a
}

/**
 * bcc
 */
func (a *EmailAction) Bcc(value interface{}) *EmailAction {
    a.Set("bcc", value)
    return a
}
