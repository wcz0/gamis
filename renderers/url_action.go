package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type UrlAction struct {
	*BaseRenderer
}

func NewUrlAction() *UrlAction {
    a := &UrlAction{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("actionType", "url")
    return a
}


func (a *UrlAction) Set(name string, value interface{}) *UrlAction {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * level
 */
func (a *UrlAction) Level(value interface{}) *UrlAction {
    a.Set("level", value)
    return a
}

/**
 * primary
 */
func (a *UrlAction) Primary(value interface{}) *UrlAction {
    a.Set("primary", value)
    return a
}

/**
 * tooltip
 */
func (a *UrlAction) Tooltip(value interface{}) *UrlAction {
    a.Set("tooltip", value)
    return a
}

/**
 * close
 */
func (a *UrlAction) Close(value interface{}) *UrlAction {
    a.Set("close", value)
    return a
}

/**
 * href
 */
func (a *UrlAction) Href(value interface{}) *UrlAction {
    a.Set("href", value)
    return a
}

/**
 * className
 */
func (a *UrlAction) ClassName(value interface{}) *UrlAction {
    a.Set("className", value)
    return a
}

/**
 * badge
 */
func (a *UrlAction) Badge(value interface{}) *UrlAction {
    a.Set("badge", value)
    return a
}

/**
 * disabledOnAction
 */
func (a *UrlAction) DisabledOnAction(value interface{}) *UrlAction {
    a.Set("disabledOnAction", value)
    return a
}

/**
 * staticClassName
 */
func (a *UrlAction) StaticClassName(value interface{}) *UrlAction {
    a.Set("staticClassName", value)
    return a
}

/**
 * disabledTip
 */
func (a *UrlAction) DisabledTip(value interface{}) *UrlAction {
    a.Set("disabledTip", value)
    return a
}

/**
 * id
 */
func (a *UrlAction) Id(value interface{}) *UrlAction {
    a.Set("id", value)
    return a
}

/**
 */
func (a *UrlAction) ActionType(value interface{}) *UrlAction {
    a.Set("actionType", value)
    return a
}

/**
 * url
 */
func (a *UrlAction) Url(value interface{}) *UrlAction {
    a.Set("url", value)
    return a
}

/**
 * disabledOn
 */
func (a *UrlAction) DisabledOn(value interface{}) *UrlAction {
    a.Set("disabledOn", value)
    return a
}

/**
 * hidden
 */
func (a *UrlAction) Hidden(value interface{}) *UrlAction {
    a.Set("hidden", value)
    return a
}

/**
 * hiddenOn
 */
func (a *UrlAction) HiddenOn(value interface{}) *UrlAction {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *UrlAction) StaticLabelClassName(value interface{}) *UrlAction {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * iconClassName
 */
func (a *UrlAction) IconClassName(value interface{}) *UrlAction {
    a.Set("iconClassName", value)
    return a
}

/**
 * loadingClassName
 */
func (a *UrlAction) LoadingClassName(value interface{}) *UrlAction {
    a.Set("loadingClassName", value)
    return a
}

/**
 * body
 */
func (a *UrlAction) Body(value interface{}) *UrlAction {
    a.Set("body", value)
    return a
}

/**
 * tabIndex
 */
func (a *UrlAction) TabIndex(value interface{}) *UrlAction {
    a.Set("tabIndex", value)
    return a
}

/**
 * block
 */
func (a *UrlAction) Block(value interface{}) *UrlAction {
    a.Set("block", value)
    return a
}

/**
 * confirmText
 */
func (a *UrlAction) ConfirmText(value interface{}) *UrlAction {
    a.Set("confirmText", value)
    return a
}

/**
 * countDown
 */
func (a *UrlAction) CountDown(value interface{}) *UrlAction {
    a.Set("countDown", value)
    return a
}

/**
 * rightIconClassName
 */
func (a *UrlAction) RightIconClassName(value interface{}) *UrlAction {
    a.Set("rightIconClassName", value)
    return a
}

/**
 * visible
 */
func (a *UrlAction) Visible(value interface{}) *UrlAction {
    a.Set("visible", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *UrlAction) StaticInputClassName(value interface{}) *UrlAction {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * style
 */
func (a *UrlAction) Style(value interface{}) *UrlAction {
    a.Set("style", value)
    return a
}

/**
 * activeClassName
 */
func (a *UrlAction) ActiveClassName(value interface{}) *UrlAction {
    a.Set("activeClassName", value)
    return a
}

/**
 * 设置链接
 */
func (a *UrlAction) Link(value interface{}) *UrlAction {
    a.Set("link", value)
    return a
}

/**
 * onEvent
 */
func (a *UrlAction) OnEvent(value interface{}) *UrlAction {
    a.Set("onEvent", value)
    return a
}

/**
 * staticOn
 */
func (a *UrlAction) StaticOn(value interface{}) *UrlAction {
    a.Set("staticOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *UrlAction) StaticPlaceholder(value interface{}) *UrlAction {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * editorSetting
 */
func (a *UrlAction) EditorSetting(value interface{}) *UrlAction {
    a.Set("editorSetting", value)
    return a
}

/**
 * requireSelected
 */
func (a *UrlAction) RequireSelected(value interface{}) *UrlAction {
    a.Set("requireSelected", value)
    return a
}

/**
 * countDownTpl
 */
func (a *UrlAction) CountDownTpl(value interface{}) *UrlAction {
    a.Set("countDownTpl", value)
    return a
}

/**
 * activeLevel
 */
func (a *UrlAction) ActiveLevel(value interface{}) *UrlAction {
    a.Set("activeLevel", value)
    return a
}

/**
 * mergeData
 */
func (a *UrlAction) MergeData(value interface{}) *UrlAction {
    a.Set("mergeData", value)
    return a
}

/**
 * target
 */
func (a *UrlAction) Target(value interface{}) *UrlAction {
    a.Set("target", value)
    return a
}

/**
 * onClick
 */
func (a *UrlAction) OnClick(value interface{}) *UrlAction {
    a.Set("onClick", value)
    return a
}

/**
 * useMobileUI
 */
func (a *UrlAction) UseMobileUI(value interface{}) *UrlAction {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 可选值: action | button | submit | reset
 */
func (a *UrlAction) Type(value interface{}) *UrlAction {
    a.Set("type", value)
    return a
}

/**
 * visibleOn
 */
func (a *UrlAction) VisibleOn(value interface{}) *UrlAction {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticSchema
 */
func (a *UrlAction) StaticSchema(value interface{}) *UrlAction {
    a.Set("staticSchema", value)
    return a
}

/**
 * icon
 */
func (a *UrlAction) Icon(value interface{}) *UrlAction {
    a.Set("icon", value)
    return a
}

/**
 * size
 */
func (a *UrlAction) Size(value interface{}) *UrlAction {
    a.Set("size", value)
    return a
}

/**
 * hotKey
 */
func (a *UrlAction) HotKey(value interface{}) *UrlAction {
    a.Set("hotKey", value)
    return a
}

/**
 * tooltipPlacement
 */
func (a *UrlAction) TooltipPlacement(value interface{}) *UrlAction {
    a.Set("tooltipPlacement", value)
    return a
}

/**
 * required
 */
func (a *UrlAction) Required(value interface{}) *UrlAction {
    a.Set("required", value)
    return a
}

/**
 * loadingOn
 */
func (a *UrlAction) LoadingOn(value interface{}) *UrlAction {
    a.Set("loadingOn", value)
    return a
}

/**
 * blank
 */
func (a *UrlAction) Blank(value interface{}) *UrlAction {
    a.Set("blank", value)
    return a
}

/**
 * static
 */
func (a *UrlAction) Static(value interface{}) *UrlAction {
    a.Set("static", value)
    return a
}

/**
 * testid
 */
func (a *UrlAction) Testid(value interface{}) *UrlAction {
    a.Set("testid", value)
    return a
}

/**
 * rightIcon
 */
func (a *UrlAction) RightIcon(value interface{}) *UrlAction {
    a.Set("rightIcon", value)
    return a
}

/**
 * label
 */
func (a *UrlAction) Label(value interface{}) *UrlAction {
    a.Set("label", value)
    return a
}

/**
 * disabled
 */
func (a *UrlAction) Disabled(value interface{}) *UrlAction {
    a.Set("disabled", value)
    return a
}
