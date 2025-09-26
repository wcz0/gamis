package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type ReloadAction struct {
	*BaseRenderer
}

func NewReloadAction() *ReloadAction {
    a := &ReloadAction{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("actionType", "reload")
    return a
}


func (a *ReloadAction) Set(name string, value interface{}) *ReloadAction {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * confirmText
 */
func (a *ReloadAction) ConfirmText(value interface{}) *ReloadAction {
    a.Set("confirmText", value)
    return a
}

/**
 * mergeData
 */
func (a *ReloadAction) MergeData(value interface{}) *ReloadAction {
    a.Set("mergeData", value)
    return a
}

/**
 * loadingOn
 */
func (a *ReloadAction) LoadingOn(value interface{}) *ReloadAction {
    a.Set("loadingOn", value)
    return a
}

/**
 * hidden
 */
func (a *ReloadAction) Hidden(value interface{}) *ReloadAction {
    a.Set("hidden", value)
    return a
}

/**
 * onEvent
 */
func (a *ReloadAction) OnEvent(value interface{}) *ReloadAction {
    a.Set("onEvent", value)
    return a
}

/**
 * staticClassName
 */
func (a *ReloadAction) StaticClassName(value interface{}) *ReloadAction {
    a.Set("staticClassName", value)
    return a
}

/**
 * useMobileUI
 */
func (a *ReloadAction) UseMobileUI(value interface{}) *ReloadAction {
    a.Set("useMobileUI", value)
    return a
}

/**
 * testid
 */
func (a *ReloadAction) Testid(value interface{}) *ReloadAction {
    a.Set("testid", value)
    return a
}

/**
 * primary
 */
func (a *ReloadAction) Primary(value interface{}) *ReloadAction {
    a.Set("primary", value)
    return a
}

/**
 * requireSelected
 */
func (a *ReloadAction) RequireSelected(value interface{}) *ReloadAction {
    a.Set("requireSelected", value)
    return a
}

/**
 * hotKey
 */
func (a *ReloadAction) HotKey(value interface{}) *ReloadAction {
    a.Set("hotKey", value)
    return a
}

/**
 * id
 */
func (a *ReloadAction) Id(value interface{}) *ReloadAction {
    a.Set("id", value)
    return a
}

/**
 * static
 */
func (a *ReloadAction) Static(value interface{}) *ReloadAction {
    a.Set("static", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *ReloadAction) StaticLabelClassName(value interface{}) *ReloadAction {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *ReloadAction) StaticSchema(value interface{}) *ReloadAction {
    a.Set("staticSchema", value)
    return a
}

/**
 * editorSetting
 */
func (a *ReloadAction) EditorSetting(value interface{}) *ReloadAction {
    a.Set("editorSetting", value)
    return a
}

/**
 * block
 */
func (a *ReloadAction) Block(value interface{}) *ReloadAction {
    a.Set("block", value)
    return a
}

/**
 * loadingClassName
 */
func (a *ReloadAction) LoadingClassName(value interface{}) *ReloadAction {
    a.Set("loadingClassName", value)
    return a
}

/**
 * level
 */
func (a *ReloadAction) Level(value interface{}) *ReloadAction {
    a.Set("level", value)
    return a
}

/**
 * staticOn
 */
func (a *ReloadAction) StaticOn(value interface{}) *ReloadAction {
    a.Set("staticOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *ReloadAction) StaticPlaceholder(value interface{}) *ReloadAction {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * iconClassName
 */
func (a *ReloadAction) IconClassName(value interface{}) *ReloadAction {
    a.Set("iconClassName", value)
    return a
}

/**
 * size
 */
func (a *ReloadAction) Size(value interface{}) *ReloadAction {
    a.Set("size", value)
    return a
}

/**
 * disabledOnAction
 */
func (a *ReloadAction) DisabledOnAction(value interface{}) *ReloadAction {
    a.Set("disabledOnAction", value)
    return a
}

/**
 * visible
 */
func (a *ReloadAction) Visible(value interface{}) *ReloadAction {
    a.Set("visible", value)
    return a
}

/**
 * visibleOn
 */
func (a *ReloadAction) VisibleOn(value interface{}) *ReloadAction {
    a.Set("visibleOn", value)
    return a
}

/**
 * disabledTip
 */
func (a *ReloadAction) DisabledTip(value interface{}) *ReloadAction {
    a.Set("disabledTip", value)
    return a
}

/**
 * close
 */
func (a *ReloadAction) Close(value interface{}) *ReloadAction {
    a.Set("close", value)
    return a
}

/**
 * target
 */
func (a *ReloadAction) Target(value interface{}) *ReloadAction {
    a.Set("target", value)
    return a
}

/**
 */
func (a *ReloadAction) ActionType(value interface{}) *ReloadAction {
    a.Set("actionType", value)
    return a
}

/**
 * disabled
 */
func (a *ReloadAction) Disabled(value interface{}) *ReloadAction {
    a.Set("disabled", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *ReloadAction) StaticInputClassName(value interface{}) *ReloadAction {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * icon
 */
func (a *ReloadAction) Icon(value interface{}) *ReloadAction {
    a.Set("icon", value)
    return a
}

/**
 * tooltip
 */
func (a *ReloadAction) Tooltip(value interface{}) *ReloadAction {
    a.Set("tooltip", value)
    return a
}

/**
 * tooltipPlacement
 */
func (a *ReloadAction) TooltipPlacement(value interface{}) *ReloadAction {
    a.Set("tooltipPlacement", value)
    return a
}

/**
 * countDown
 */
func (a *ReloadAction) CountDown(value interface{}) *ReloadAction {
    a.Set("countDown", value)
    return a
}

/**
 * tabIndex
 */
func (a *ReloadAction) TabIndex(value interface{}) *ReloadAction {
    a.Set("tabIndex", value)
    return a
}

/**
 * rightIcon
 */
func (a *ReloadAction) RightIcon(value interface{}) *ReloadAction {
    a.Set("rightIcon", value)
    return a
}

/**
 * label
 */
func (a *ReloadAction) Label(value interface{}) *ReloadAction {
    a.Set("label", value)
    return a
}

/**
 * activeLevel
 */
func (a *ReloadAction) ActiveLevel(value interface{}) *ReloadAction {
    a.Set("activeLevel", value)
    return a
}

/**
 * onClick
 */
func (a *ReloadAction) OnClick(value interface{}) *ReloadAction {
    a.Set("onClick", value)
    return a
}

/**
 * body
 */
func (a *ReloadAction) Body(value interface{}) *ReloadAction {
    a.Set("body", value)
    return a
}

/**
 * style
 */
func (a *ReloadAction) Style(value interface{}) *ReloadAction {
    a.Set("style", value)
    return a
}

/**
 * required
 */
func (a *ReloadAction) Required(value interface{}) *ReloadAction {
    a.Set("required", value)
    return a
}

/**
 * activeClassName
 */
func (a *ReloadAction) ActiveClassName(value interface{}) *ReloadAction {
    a.Set("activeClassName", value)
    return a
}

/**
 * countDownTpl
 */
func (a *ReloadAction) CountDownTpl(value interface{}) *ReloadAction {
    a.Set("countDownTpl", value)
    return a
}

/**
 * badge
 */
func (a *ReloadAction) Badge(value interface{}) *ReloadAction {
    a.Set("badge", value)
    return a
}

/**
 * href
 */
func (a *ReloadAction) Href(value interface{}) *ReloadAction {
    a.Set("href", value)
    return a
}

/**
 * className
 */
func (a *ReloadAction) ClassName(value interface{}) *ReloadAction {
    a.Set("className", value)
    return a
}

/**
 * disabledOn
 */
func (a *ReloadAction) DisabledOn(value interface{}) *ReloadAction {
    a.Set("disabledOn", value)
    return a
}

/**
 * hiddenOn
 */
func (a *ReloadAction) HiddenOn(value interface{}) *ReloadAction {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 可选值: action | button | submit | reset
 */
func (a *ReloadAction) Type(value interface{}) *ReloadAction {
    a.Set("type", value)
    return a
}

/**
 * rightIconClassName
 */
func (a *ReloadAction) RightIconClassName(value interface{}) *ReloadAction {
    a.Set("rightIconClassName", value)
    return a
}
