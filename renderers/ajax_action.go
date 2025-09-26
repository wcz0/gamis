package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type AjaxAction struct {
	*BaseRenderer
}

func NewAjaxAction() *AjaxAction {
    a := &AjaxAction{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("actionType", "ajax")
    return a
}


func (a *AjaxAction) Set(name string, value interface{}) *AjaxAction {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * block
 */
func (a *AjaxAction) Block(value interface{}) *AjaxAction {
    a.Set("block", value)
    return a
}

/**
 * size
 */
func (a *AjaxAction) Size(value interface{}) *AjaxAction {
    a.Set("size", value)
    return a
}

/**
 * onClick
 */
func (a *AjaxAction) OnClick(value interface{}) *AjaxAction {
    a.Set("onClick", value)
    return a
}

/**
 */
func (a *AjaxAction) ActionType(value interface{}) *AjaxAction {
    a.Set("actionType", value)
    return a
}

/**
 * redirect
 */
func (a *AjaxAction) Redirect(value interface{}) *AjaxAction {
    a.Set("redirect", value)
    return a
}

/**
 * hidden
 */
func (a *AjaxAction) Hidden(value interface{}) *AjaxAction {
    a.Set("hidden", value)
    return a
}

/**
 * requireSelected
 */
func (a *AjaxAction) RequireSelected(value interface{}) *AjaxAction {
    a.Set("requireSelected", value)
    return a
}

/**
 * hiddenOn
 */
func (a *AjaxAction) HiddenOn(value interface{}) *AjaxAction {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visible
 */
func (a *AjaxAction) Visible(value interface{}) *AjaxAction {
    a.Set("visible", value)
    return a
}

/**
 * id
 */
func (a *AjaxAction) Id(value interface{}) *AjaxAction {
    a.Set("id", value)
    return a
}

/**
 * style
 */
func (a *AjaxAction) Style(value interface{}) *AjaxAction {
    a.Set("style", value)
    return a
}

/**
 * required
 */
func (a *AjaxAction) Required(value interface{}) *AjaxAction {
    a.Set("required", value)
    return a
}

/**
 * mergeData
 */
func (a *AjaxAction) MergeData(value interface{}) *AjaxAction {
    a.Set("mergeData", value)
    return a
}

/**
 * disabled
 */
func (a *AjaxAction) Disabled(value interface{}) *AjaxAction {
    a.Set("disabled", value)
    return a
}

/**
 * reload
 */
func (a *AjaxAction) Reload(value interface{}) *AjaxAction {
    a.Set("reload", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *AjaxAction) StaticPlaceholder(value interface{}) *AjaxAction {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * loadingClassName
 */
func (a *AjaxAction) LoadingClassName(value interface{}) *AjaxAction {
    a.Set("loadingClassName", value)
    return a
}

/**
 * primary
 */
func (a *AjaxAction) Primary(value interface{}) *AjaxAction {
    a.Set("primary", value)
    return a
}

/**
 * loadingOn
 */
func (a *AjaxAction) LoadingOn(value interface{}) *AjaxAction {
    a.Set("loadingOn", value)
    return a
}

/**
 * disabledOnAction
 */
func (a *AjaxAction) DisabledOnAction(value interface{}) *AjaxAction {
    a.Set("disabledOnAction", value)
    return a
}

/**
 * visibleOn
 */
func (a *AjaxAction) VisibleOn(value interface{}) *AjaxAction {
    a.Set("visibleOn", value)
    return a
}

/**
 * label
 */
func (a *AjaxAction) Label(value interface{}) *AjaxAction {
    a.Set("label", value)
    return a
}

/**
 * activeClassName
 */
func (a *AjaxAction) ActiveClassName(value interface{}) *AjaxAction {
    a.Set("activeClassName", value)
    return a
}

/**
 * activeLevel
 */
func (a *AjaxAction) ActiveLevel(value interface{}) *AjaxAction {
    a.Set("activeLevel", value)
    return a
}

/**
 * countDownTpl
 */
func (a *AjaxAction) CountDownTpl(value interface{}) *AjaxAction {
    a.Set("countDownTpl", value)
    return a
}

/**
 * isolateScope
 */
func (a *AjaxAction) IsolateScope(value interface{}) *AjaxAction {
    a.Set("isolateScope", value)
    return a
}

/**
 * staticSchema
 */
func (a *AjaxAction) StaticSchema(value interface{}) *AjaxAction {
    a.Set("staticSchema", value)
    return a
}

/**
 * testid
 */
func (a *AjaxAction) Testid(value interface{}) *AjaxAction {
    a.Set("testid", value)
    return a
}

/**
 * close
 */
func (a *AjaxAction) Close(value interface{}) *AjaxAction {
    a.Set("close", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *AjaxAction) StaticLabelClassName(value interface{}) *AjaxAction {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * rightIconClassName
 */
func (a *AjaxAction) RightIconClassName(value interface{}) *AjaxAction {
    a.Set("rightIconClassName", value)
    return a
}

/**
 * body
 */
func (a *AjaxAction) Body(value interface{}) *AjaxAction {
    a.Set("body", value)
    return a
}

/**
 * className
 */
func (a *AjaxAction) ClassName(value interface{}) *AjaxAction {
    a.Set("className", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *AjaxAction) StaticInputClassName(value interface{}) *AjaxAction {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * confirmText
 */
func (a *AjaxAction) ConfirmText(value interface{}) *AjaxAction {
    a.Set("confirmText", value)
    return a
}

/**
 * href
 */
func (a *AjaxAction) Href(value interface{}) *AjaxAction {
    a.Set("href", value)
    return a
}

/**
 * useMobileUI
 */
func (a *AjaxAction) UseMobileUI(value interface{}) *AjaxAction {
    a.Set("useMobileUI", value)
    return a
}

/**
 * staticOn
 */
func (a *AjaxAction) StaticOn(value interface{}) *AjaxAction {
    a.Set("staticOn", value)
    return a
}

/**
 * tooltip
 */
func (a *AjaxAction) Tooltip(value interface{}) *AjaxAction {
    a.Set("tooltip", value)
    return a
}

/**
 * tooltipPlacement
 */
func (a *AjaxAction) TooltipPlacement(value interface{}) *AjaxAction {
    a.Set("tooltipPlacement", value)
    return a
}

/**
 * feedback
 */
func (a *AjaxAction) Feedback(value interface{}) *AjaxAction {
    a.Set("feedback", value)
    return a
}

/**
 * ignoreConfirm
 */
func (a *AjaxAction) IgnoreConfirm(value interface{}) *AjaxAction {
    a.Set("ignoreConfirm", value)
    return a
}

/**
 * disabledTip
 */
func (a *AjaxAction) DisabledTip(value interface{}) *AjaxAction {
    a.Set("disabledTip", value)
    return a
}

/**
 * countDown
 */
func (a *AjaxAction) CountDown(value interface{}) *AjaxAction {
    a.Set("countDown", value)
    return a
}

/**
 * badge
 */
func (a *AjaxAction) Badge(value interface{}) *AjaxAction {
    a.Set("badge", value)
    return a
}

/**
 * hotKey
 */
func (a *AjaxAction) HotKey(value interface{}) *AjaxAction {
    a.Set("hotKey", value)
    return a
}

/**
 * icon
 */
func (a *AjaxAction) Icon(value interface{}) *AjaxAction {
    a.Set("icon", value)
    return a
}

/**
 * rightIcon
 */
func (a *AjaxAction) RightIcon(value interface{}) *AjaxAction {
    a.Set("rightIcon", value)
    return a
}

/**
 * api
 */
func (a *AjaxAction) Api(value interface{}) *AjaxAction {
    a.Set("api", value)
    return a
}

/**
 * 可选值: action | button | submit | reset
 */
func (a *AjaxAction) Type(value interface{}) *AjaxAction {
    a.Set("type", value)
    return a
}

/**
 * iconClassName
 */
func (a *AjaxAction) IconClassName(value interface{}) *AjaxAction {
    a.Set("iconClassName", value)
    return a
}

/**
 * disabledOn
 */
func (a *AjaxAction) DisabledOn(value interface{}) *AjaxAction {
    a.Set("disabledOn", value)
    return a
}

/**
 * static
 */
func (a *AjaxAction) Static(value interface{}) *AjaxAction {
    a.Set("static", value)
    return a
}

/**
 * target
 */
func (a *AjaxAction) Target(value interface{}) *AjaxAction {
    a.Set("target", value)
    return a
}

/**
 * tabIndex
 */
func (a *AjaxAction) TabIndex(value interface{}) *AjaxAction {
    a.Set("tabIndex", value)
    return a
}

/**
 * onEvent
 */
func (a *AjaxAction) OnEvent(value interface{}) *AjaxAction {
    a.Set("onEvent", value)
    return a
}

/**
 * staticClassName
 */
func (a *AjaxAction) StaticClassName(value interface{}) *AjaxAction {
    a.Set("staticClassName", value)
    return a
}

/**
 * editorSetting
 */
func (a *AjaxAction) EditorSetting(value interface{}) *AjaxAction {
    a.Set("editorSetting", value)
    return a
}

/**
 * level
 */
func (a *AjaxAction) Level(value interface{}) *AjaxAction {
    a.Set("level", value)
    return a
}
