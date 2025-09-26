package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type CopyAction struct {
	*BaseRenderer
}

func NewCopyAction() *CopyAction {
    a := &CopyAction{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("actionType", "copy")
    return a
}


func (a *CopyAction) Set(name string, value interface{}) *CopyAction {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * iconClassName
 */
func (a *CopyAction) IconClassName(value interface{}) *CopyAction {
    a.Set("iconClassName", value)
    return a
}

/**
 * body
 */
func (a *CopyAction) Body(value interface{}) *CopyAction {
    a.Set("body", value)
    return a
}

/**
 * copy
 */
func (a *CopyAction) Copy(value interface{}) *CopyAction {
    a.Set("copy", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *CopyAction) StaticInputClassName(value interface{}) *CopyAction {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *CopyAction) StaticSchema(value interface{}) *CopyAction {
    a.Set("staticSchema", value)
    return a
}

/**
 * required
 */
func (a *CopyAction) Required(value interface{}) *CopyAction {
    a.Set("required", value)
    return a
}

/**
 * mergeData
 */
func (a *CopyAction) MergeData(value interface{}) *CopyAction {
    a.Set("mergeData", value)
    return a
}

/**
 * disabledOnAction
 */
func (a *CopyAction) DisabledOnAction(value interface{}) *CopyAction {
    a.Set("disabledOnAction", value)
    return a
}

/**
 * onEvent
 */
func (a *CopyAction) OnEvent(value interface{}) *CopyAction {
    a.Set("onEvent", value)
    return a
}

/**
 * useMobileUI
 */
func (a *CopyAction) UseMobileUI(value interface{}) *CopyAction {
    a.Set("useMobileUI", value)
    return a
}

/**
 * testid
 */
func (a *CopyAction) Testid(value interface{}) *CopyAction {
    a.Set("testid", value)
    return a
}

/**
 * rightIconClassName
 */
func (a *CopyAction) RightIconClassName(value interface{}) *CopyAction {
    a.Set("rightIconClassName", value)
    return a
}

/**
 * size
 */
func (a *CopyAction) Size(value interface{}) *CopyAction {
    a.Set("size", value)
    return a
}

/**
 * tooltipPlacement
 */
func (a *CopyAction) TooltipPlacement(value interface{}) *CopyAction {
    a.Set("tooltipPlacement", value)
    return a
}

/**
 * activeLevel
 */
func (a *CopyAction) ActiveLevel(value interface{}) *CopyAction {
    a.Set("activeLevel", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *CopyAction) StaticPlaceholder(value interface{}) *CopyAction {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *CopyAction) StaticLabelClassName(value interface{}) *CopyAction {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * hotKey
 */
func (a *CopyAction) HotKey(value interface{}) *CopyAction {
    a.Set("hotKey", value)
    return a
}

/**
 * tabIndex
 */
func (a *CopyAction) TabIndex(value interface{}) *CopyAction {
    a.Set("tabIndex", value)
    return a
}

/**
 * hidden
 */
func (a *CopyAction) Hidden(value interface{}) *CopyAction {
    a.Set("hidden", value)
    return a
}

/**
 * visible
 */
func (a *CopyAction) Visible(value interface{}) *CopyAction {
    a.Set("visible", value)
    return a
}

/**
 * rightIcon
 */
func (a *CopyAction) RightIcon(value interface{}) *CopyAction {
    a.Set("rightIcon", value)
    return a
}

/**
 * level
 */
func (a *CopyAction) Level(value interface{}) *CopyAction {
    a.Set("level", value)
    return a
}

/**
 * requireSelected
 */
func (a *CopyAction) RequireSelected(value interface{}) *CopyAction {
    a.Set("requireSelected", value)
    return a
}

/**
 * badge
 */
func (a *CopyAction) Badge(value interface{}) *CopyAction {
    a.Set("badge", value)
    return a
}

/**
 * disabledOn
 */
func (a *CopyAction) DisabledOn(value interface{}) *CopyAction {
    a.Set("disabledOn", value)
    return a
}

/**
 * id
 */
func (a *CopyAction) Id(value interface{}) *CopyAction {
    a.Set("id", value)
    return a
}

/**
 * editorSetting
 */
func (a *CopyAction) EditorSetting(value interface{}) *CopyAction {
    a.Set("editorSetting", value)
    return a
}

/**
 * disabledTip
 */
func (a *CopyAction) DisabledTip(value interface{}) *CopyAction {
    a.Set("disabledTip", value)
    return a
}

/**
 * loadingClassName
 */
func (a *CopyAction) LoadingClassName(value interface{}) *CopyAction {
    a.Set("loadingClassName", value)
    return a
}

/**
 * activeClassName
 */
func (a *CopyAction) ActiveClassName(value interface{}) *CopyAction {
    a.Set("activeClassName", value)
    return a
}

/**
 * close
 */
func (a *CopyAction) Close(value interface{}) *CopyAction {
    a.Set("close", value)
    return a
}

/**
 * loadingOn
 */
func (a *CopyAction) LoadingOn(value interface{}) *CopyAction {
    a.Set("loadingOn", value)
    return a
}

/**
 * className
 */
func (a *CopyAction) ClassName(value interface{}) *CopyAction {
    a.Set("className", value)
    return a
}

/**
 * hiddenOn
 */
func (a *CopyAction) HiddenOn(value interface{}) *CopyAction {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visibleOn
 */
func (a *CopyAction) VisibleOn(value interface{}) *CopyAction {
    a.Set("visibleOn", value)
    return a
}

/**
 * static
 */
func (a *CopyAction) Static(value interface{}) *CopyAction {
    a.Set("static", value)
    return a
}

/**
 * staticOn
 */
func (a *CopyAction) StaticOn(value interface{}) *CopyAction {
    a.Set("staticOn", value)
    return a
}

/**
 * primary
 */
func (a *CopyAction) Primary(value interface{}) *CopyAction {
    a.Set("primary", value)
    return a
}

/**
 * target
 */
func (a *CopyAction) Target(value interface{}) *CopyAction {
    a.Set("target", value)
    return a
}

/**
 * countDownTpl
 */
func (a *CopyAction) CountDownTpl(value interface{}) *CopyAction {
    a.Set("countDownTpl", value)
    return a
}

/**
 * style
 */
func (a *CopyAction) Style(value interface{}) *CopyAction {
    a.Set("style", value)
    return a
}

/**
 * label
 */
func (a *CopyAction) Label(value interface{}) *CopyAction {
    a.Set("label", value)
    return a
}

/**
 * tooltip
 */
func (a *CopyAction) Tooltip(value interface{}) *CopyAction {
    a.Set("tooltip", value)
    return a
}

/**
 * confirmText
 */
func (a *CopyAction) ConfirmText(value interface{}) *CopyAction {
    a.Set("confirmText", value)
    return a
}

/**
 * countDown
 */
func (a *CopyAction) CountDown(value interface{}) *CopyAction {
    a.Set("countDown", value)
    return a
}

/**
 * onClick
 */
func (a *CopyAction) OnClick(value interface{}) *CopyAction {
    a.Set("onClick", value)
    return a
}

/**
 * href
 */
func (a *CopyAction) Href(value interface{}) *CopyAction {
    a.Set("href", value)
    return a
}

/**
 */
func (a *CopyAction) ActionType(value interface{}) *CopyAction {
    a.Set("actionType", value)
    return a
}

/**
 * disabled
 */
func (a *CopyAction) Disabled(value interface{}) *CopyAction {
    a.Set("disabled", value)
    return a
}

/**
 * staticClassName
 */
func (a *CopyAction) StaticClassName(value interface{}) *CopyAction {
    a.Set("staticClassName", value)
    return a
}

/**
 * 可选值: action | button | submit | reset
 */
func (a *CopyAction) Type(value interface{}) *CopyAction {
    a.Set("type", value)
    return a
}

/**
 * block
 */
func (a *CopyAction) Block(value interface{}) *CopyAction {
    a.Set("block", value)
    return a
}

/**
 * icon
 */
func (a *CopyAction) Icon(value interface{}) *CopyAction {
    a.Set("icon", value)
    return a
}
