package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type DialogAction struct {
	*BaseRenderer
}

func NewDialogAction() *DialogAction {
    a := &DialogAction{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("actionType", "dialog")
    return a
}


func (a *DialogAction) Set(name string, value interface{}) *DialogAction {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * onEvent
 */
func (a *DialogAction) OnEvent(value interface{}) *DialogAction {
    a.Set("onEvent", value)
    return a
}

/**
 * 可选值: action | button | submit | reset
 */
func (a *DialogAction) Type(value interface{}) *DialogAction {
    a.Set("type", value)
    return a
}

/**
 * redirect
 */
func (a *DialogAction) Redirect(value interface{}) *DialogAction {
    a.Set("redirect", value)
    return a
}

/**
 * countDownTpl
 */
func (a *DialogAction) CountDownTpl(value interface{}) *DialogAction {
    a.Set("countDownTpl", value)
    return a
}

/**
 * disabledOnAction
 */
func (a *DialogAction) DisabledOnAction(value interface{}) *DialogAction {
    a.Set("disabledOnAction", value)
    return a
}

/**
 * hidden
 */
func (a *DialogAction) Hidden(value interface{}) *DialogAction {
    a.Set("hidden", value)
    return a
}

/**
 * staticClassName
 */
func (a *DialogAction) StaticClassName(value interface{}) *DialogAction {
    a.Set("staticClassName", value)
    return a
}

/**
 * testid
 */
func (a *DialogAction) Testid(value interface{}) *DialogAction {
    a.Set("testid", value)
    return a
}

/**
 * level
 */
func (a *DialogAction) Level(value interface{}) *DialogAction {
    a.Set("level", value)
    return a
}

/**
 * tooltip
 */
func (a *DialogAction) Tooltip(value interface{}) *DialogAction {
    a.Set("tooltip", value)
    return a
}

/**
 * close
 */
func (a *DialogAction) Close(value interface{}) *DialogAction {
    a.Set("close", value)
    return a
}

/**
 * visibleOn
 */
func (a *DialogAction) VisibleOn(value interface{}) *DialogAction {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticOn
 */
func (a *DialogAction) StaticOn(value interface{}) *DialogAction {
    a.Set("staticOn", value)
    return a
}

/**
 * staticSchema
 */
func (a *DialogAction) StaticSchema(value interface{}) *DialogAction {
    a.Set("staticSchema", value)
    return a
}

/**
 * reload
 */
func (a *DialogAction) Reload(value interface{}) *DialogAction {
    a.Set("reload", value)
    return a
}

/**
 * disabledOn
 */
func (a *DialogAction) DisabledOn(value interface{}) *DialogAction {
    a.Set("disabledOn", value)
    return a
}

/**
 * icon
 */
func (a *DialogAction) Icon(value interface{}) *DialogAction {
    a.Set("icon", value)
    return a
}

/**
 * rightIconClassName
 */
func (a *DialogAction) RightIconClassName(value interface{}) *DialogAction {
    a.Set("rightIconClassName", value)
    return a
}

/**
 * target
 */
func (a *DialogAction) Target(value interface{}) *DialogAction {
    a.Set("target", value)
    return a
}

/**
 * loadingOn
 */
func (a *DialogAction) LoadingOn(value interface{}) *DialogAction {
    a.Set("loadingOn", value)
    return a
}

/**
 * visible
 */
func (a *DialogAction) Visible(value interface{}) *DialogAction {
    a.Set("visible", value)
    return a
}

/**
 * disabled
 */
func (a *DialogAction) Disabled(value interface{}) *DialogAction {
    a.Set("disabled", value)
    return a
}

/**
 * id
 */
func (a *DialogAction) Id(value interface{}) *DialogAction {
    a.Set("id", value)
    return a
}

/**
 * loadingClassName
 */
func (a *DialogAction) LoadingClassName(value interface{}) *DialogAction {
    a.Set("loadingClassName", value)
    return a
}

/**
 * activeClassName
 */
func (a *DialogAction) ActiveClassName(value interface{}) *DialogAction {
    a.Set("activeClassName", value)
    return a
}

/**
 * tabIndex
 */
func (a *DialogAction) TabIndex(value interface{}) *DialogAction {
    a.Set("tabIndex", value)
    return a
}

/**
 * static
 */
func (a *DialogAction) Static(value interface{}) *DialogAction {
    a.Set("static", value)
    return a
}

/**
 * block
 */
func (a *DialogAction) Block(value interface{}) *DialogAction {
    a.Set("block", value)
    return a
}

/**
 * size
 */
func (a *DialogAction) Size(value interface{}) *DialogAction {
    a.Set("size", value)
    return a
}

/**
 * badge
 */
func (a *DialogAction) Badge(value interface{}) *DialogAction {
    a.Set("badge", value)
    return a
}

/**
 */
func (a *DialogAction) ActionType(value interface{}) *DialogAction {
    a.Set("actionType", value)
    return a
}

/**
 * nextCondition
 */
func (a *DialogAction) NextCondition(value interface{}) *DialogAction {
    a.Set("nextCondition", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *DialogAction) StaticPlaceholder(value interface{}) *DialogAction {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * hiddenOn
 */
func (a *DialogAction) HiddenOn(value interface{}) *DialogAction {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *DialogAction) StaticLabelClassName(value interface{}) *DialogAction {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * iconClassName
 */
func (a *DialogAction) IconClassName(value interface{}) *DialogAction {
    a.Set("iconClassName", value)
    return a
}

/**
 * style
 */
func (a *DialogAction) Style(value interface{}) *DialogAction {
    a.Set("style", value)
    return a
}

/**
 * confirmText
 */
func (a *DialogAction) ConfirmText(value interface{}) *DialogAction {
    a.Set("confirmText", value)
    return a
}

/**
 * disabledTip
 */
func (a *DialogAction) DisabledTip(value interface{}) *DialogAction {
    a.Set("disabledTip", value)
    return a
}

/**
 * primary
 */
func (a *DialogAction) Primary(value interface{}) *DialogAction {
    a.Set("primary", value)
    return a
}

/**
 * mergeData
 */
func (a *DialogAction) MergeData(value interface{}) *DialogAction {
    a.Set("mergeData", value)
    return a
}

/**
 * hotKey
 */
func (a *DialogAction) HotKey(value interface{}) *DialogAction {
    a.Set("hotKey", value)
    return a
}

/**
 * onClick
 */
func (a *DialogAction) OnClick(value interface{}) *DialogAction {
    a.Set("onClick", value)
    return a
}

/**
 * editorSetting
 */
func (a *DialogAction) EditorSetting(value interface{}) *DialogAction {
    a.Set("editorSetting", value)
    return a
}

/**
 * label
 */
func (a *DialogAction) Label(value interface{}) *DialogAction {
    a.Set("label", value)
    return a
}

/**
 * activeLevel
 */
func (a *DialogAction) ActiveLevel(value interface{}) *DialogAction {
    a.Set("activeLevel", value)
    return a
}

/**
 * body
 */
func (a *DialogAction) Body(value interface{}) *DialogAction {
    a.Set("body", value)
    return a
}

/**
 * className
 */
func (a *DialogAction) ClassName(value interface{}) *DialogAction {
    a.Set("className", value)
    return a
}

/**
 * required
 */
func (a *DialogAction) Required(value interface{}) *DialogAction {
    a.Set("required", value)
    return a
}

/**
 * dialog
 */
func (a *DialogAction) Dialog(value interface{}) *DialogAction {
    a.Set("dialog", value)
    return a
}

/**
 * data
 */
func (a *DialogAction) Data(value interface{}) *DialogAction {
    a.Set("data", value)
    return a
}

/**
 * tooltipPlacement
 */
func (a *DialogAction) TooltipPlacement(value interface{}) *DialogAction {
    a.Set("tooltipPlacement", value)
    return a
}

/**
 * requireSelected
 */
func (a *DialogAction) RequireSelected(value interface{}) *DialogAction {
    a.Set("requireSelected", value)
    return a
}

/**
 * countDown
 */
func (a *DialogAction) CountDown(value interface{}) *DialogAction {
    a.Set("countDown", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *DialogAction) StaticInputClassName(value interface{}) *DialogAction {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * useMobileUI
 */
func (a *DialogAction) UseMobileUI(value interface{}) *DialogAction {
    a.Set("useMobileUI", value)
    return a
}

/**
 * rightIcon
 */
func (a *DialogAction) RightIcon(value interface{}) *DialogAction {
    a.Set("rightIcon", value)
    return a
}

/**
 * href
 */
func (a *DialogAction) Href(value interface{}) *DialogAction {
    a.Set("href", value)
    return a
}
