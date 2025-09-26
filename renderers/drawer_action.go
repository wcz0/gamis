package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type DrawerAction struct {
	*BaseRenderer
}

func NewDrawerAction() *DrawerAction {
    a := &DrawerAction{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("actionType", "drawer")
    return a
}


func (a *DrawerAction) Set(name string, value interface{}) *DrawerAction {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * id
 */
func (a *DrawerAction) Id(value interface{}) *DrawerAction {
    a.Set("id", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *DrawerAction) StaticLabelClassName(value interface{}) *DrawerAction {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * level
 */
func (a *DrawerAction) Level(value interface{}) *DrawerAction {
    a.Set("level", value)
    return a
}

/**
 * close
 */
func (a *DrawerAction) Close(value interface{}) *DrawerAction {
    a.Set("close", value)
    return a
}

/**
 * data
 */
func (a *DrawerAction) Data(value interface{}) *DrawerAction {
    a.Set("data", value)
    return a
}

/**
 * target
 */
func (a *DrawerAction) Target(value interface{}) *DrawerAction {
    a.Set("target", value)
    return a
}

/**
 * onClick
 */
func (a *DrawerAction) OnClick(value interface{}) *DrawerAction {
    a.Set("onClick", value)
    return a
}

/**
 * testid
 */
func (a *DrawerAction) Testid(value interface{}) *DrawerAction {
    a.Set("testid", value)
    return a
}

/**
 * block
 */
func (a *DrawerAction) Block(value interface{}) *DrawerAction {
    a.Set("block", value)
    return a
}

/**
 * rightIcon
 */
func (a *DrawerAction) RightIcon(value interface{}) *DrawerAction {
    a.Set("rightIcon", value)
    return a
}

/**
 * disabled
 */
func (a *DrawerAction) Disabled(value interface{}) *DrawerAction {
    a.Set("disabled", value)
    return a
}

/**
 * hiddenOn
 */
func (a *DrawerAction) HiddenOn(value interface{}) *DrawerAction {
    a.Set("hiddenOn", value)
    return a
}

/**
 * useMobileUI
 */
func (a *DrawerAction) UseMobileUI(value interface{}) *DrawerAction {
    a.Set("useMobileUI", value)
    return a
}

/**
 * staticClassName
 */
func (a *DrawerAction) StaticClassName(value interface{}) *DrawerAction {
    a.Set("staticClassName", value)
    return a
}

/**
 * loadingClassName
 */
func (a *DrawerAction) LoadingClassName(value interface{}) *DrawerAction {
    a.Set("loadingClassName", value)
    return a
}

/**
 * drawer
 */
func (a *DrawerAction) Drawer(value interface{}) *DrawerAction {
    a.Set("drawer", value)
    return a
}

/**
 * nextCondition
 */
func (a *DrawerAction) NextCondition(value interface{}) *DrawerAction {
    a.Set("nextCondition", value)
    return a
}

/**
 * onEvent
 */
func (a *DrawerAction) OnEvent(value interface{}) *DrawerAction {
    a.Set("onEvent", value)
    return a
}

/**
 * editorSetting
 */
func (a *DrawerAction) EditorSetting(value interface{}) *DrawerAction {
    a.Set("editorSetting", value)
    return a
}

/**
 * disabledOnAction
 */
func (a *DrawerAction) DisabledOnAction(value interface{}) *DrawerAction {
    a.Set("disabledOnAction", value)
    return a
}

/**
 * disabledOn
 */
func (a *DrawerAction) DisabledOn(value interface{}) *DrawerAction {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticSchema
 */
func (a *DrawerAction) StaticSchema(value interface{}) *DrawerAction {
    a.Set("staticSchema", value)
    return a
}

/**
 * loadingOn
 */
func (a *DrawerAction) LoadingOn(value interface{}) *DrawerAction {
    a.Set("loadingOn", value)
    return a
}

/**
 * body
 */
func (a *DrawerAction) Body(value interface{}) *DrawerAction {
    a.Set("body", value)
    return a
}

/**
 * staticOn
 */
func (a *DrawerAction) StaticOn(value interface{}) *DrawerAction {
    a.Set("staticOn", value)
    return a
}

/**
 * icon
 */
func (a *DrawerAction) Icon(value interface{}) *DrawerAction {
    a.Set("icon", value)
    return a
}

/**
 * tabIndex
 */
func (a *DrawerAction) TabIndex(value interface{}) *DrawerAction {
    a.Set("tabIndex", value)
    return a
}

/**
 * static
 */
func (a *DrawerAction) Static(value interface{}) *DrawerAction {
    a.Set("static", value)
    return a
}

/**
 * label
 */
func (a *DrawerAction) Label(value interface{}) *DrawerAction {
    a.Set("label", value)
    return a
}

/**
 * badge
 */
func (a *DrawerAction) Badge(value interface{}) *DrawerAction {
    a.Set("badge", value)
    return a
}

/**
 * disabledTip
 */
func (a *DrawerAction) DisabledTip(value interface{}) *DrawerAction {
    a.Set("disabledTip", value)
    return a
}

/**
 * tooltipPlacement
 */
func (a *DrawerAction) TooltipPlacement(value interface{}) *DrawerAction {
    a.Set("tooltipPlacement", value)
    return a
}

/**
 * primary
 */
func (a *DrawerAction) Primary(value interface{}) *DrawerAction {
    a.Set("primary", value)
    return a
}

/**
 * tooltip
 */
func (a *DrawerAction) Tooltip(value interface{}) *DrawerAction {
    a.Set("tooltip", value)
    return a
}

/**
 * activeLevel
 */
func (a *DrawerAction) ActiveLevel(value interface{}) *DrawerAction {
    a.Set("activeLevel", value)
    return a
}

/**
 * activeClassName
 */
func (a *DrawerAction) ActiveClassName(value interface{}) *DrawerAction {
    a.Set("activeClassName", value)
    return a
}

/**
 * visible
 */
func (a *DrawerAction) Visible(value interface{}) *DrawerAction {
    a.Set("visible", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *DrawerAction) StaticPlaceholder(value interface{}) *DrawerAction {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *DrawerAction) StaticInputClassName(value interface{}) *DrawerAction {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * className
 */
func (a *DrawerAction) ClassName(value interface{}) *DrawerAction {
    a.Set("className", value)
    return a
}

/**
 * 可选值: action | button | submit | reset
 */
func (a *DrawerAction) Type(value interface{}) *DrawerAction {
    a.Set("type", value)
    return a
}

/**
 * rightIconClassName
 */
func (a *DrawerAction) RightIconClassName(value interface{}) *DrawerAction {
    a.Set("rightIconClassName", value)
    return a
}

/**
 * countDownTpl
 */
func (a *DrawerAction) CountDownTpl(value interface{}) *DrawerAction {
    a.Set("countDownTpl", value)
    return a
}

/**
 * redirect
 */
func (a *DrawerAction) Redirect(value interface{}) *DrawerAction {
    a.Set("redirect", value)
    return a
}

/**
 * style
 */
func (a *DrawerAction) Style(value interface{}) *DrawerAction {
    a.Set("style", value)
    return a
}

/**
 * iconClassName
 */
func (a *DrawerAction) IconClassName(value interface{}) *DrawerAction {
    a.Set("iconClassName", value)
    return a
}

/**
 * confirmText
 */
func (a *DrawerAction) ConfirmText(value interface{}) *DrawerAction {
    a.Set("confirmText", value)
    return a
}

/**
 */
func (a *DrawerAction) ActionType(value interface{}) *DrawerAction {
    a.Set("actionType", value)
    return a
}

/**
 * reload
 */
func (a *DrawerAction) Reload(value interface{}) *DrawerAction {
    a.Set("reload", value)
    return a
}

/**
 * mergeData
 */
func (a *DrawerAction) MergeData(value interface{}) *DrawerAction {
    a.Set("mergeData", value)
    return a
}

/**
 * countDown
 */
func (a *DrawerAction) CountDown(value interface{}) *DrawerAction {
    a.Set("countDown", value)
    return a
}

/**
 * href
 */
func (a *DrawerAction) Href(value interface{}) *DrawerAction {
    a.Set("href", value)
    return a
}

/**
 * 设置对齐方式
 */
func (a *DrawerAction) Align(value interface{}) *DrawerAction {
    a.Set("align", value)
    return a
}

/**
 * hidden
 */
func (a *DrawerAction) Hidden(value interface{}) *DrawerAction {
    a.Set("hidden", value)
    return a
}

/**
 * visibleOn
 */
func (a *DrawerAction) VisibleOn(value interface{}) *DrawerAction {
    a.Set("visibleOn", value)
    return a
}

/**
 * size
 */
func (a *DrawerAction) Size(value interface{}) *DrawerAction {
    a.Set("size", value)
    return a
}

/**
 * requireSelected
 */
func (a *DrawerAction) RequireSelected(value interface{}) *DrawerAction {
    a.Set("requireSelected", value)
    return a
}

/**
 * required
 */
func (a *DrawerAction) Required(value interface{}) *DrawerAction {
    a.Set("required", value)
    return a
}

/**
 * hotKey
 */
func (a *DrawerAction) HotKey(value interface{}) *DrawerAction {
    a.Set("hotKey", value)
    return a
}
