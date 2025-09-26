package renderers

type LinkAction struct {
	*BaseRenderer
}

func NewLinkAction() *LinkAction {
	a := &LinkAction{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "button")
	return a
}

func (a *LinkAction) Set(name string, value interface{}) *LinkAction {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *LinkAction) ActionType(value interface{}) *LinkAction {
	a.Set("actionType", value)
	return a
}

func (a *LinkAction) ActiveClassName(value interface{}) *LinkAction {
	a.Set("activeClassName", value)
	return a
}

func (a *LinkAction) ActiveLevel(value interface{}) *LinkAction {
	a.Set("activeLevel", value)
	return a
}

func (a *LinkAction) Badge(value interface{}) *LinkAction {
	a.Set("badge", value)
	return a
}

func (a *LinkAction) Block(value interface{}) *LinkAction {
	a.Set("block", value)
	return a
}

func (a *LinkAction) Body(value interface{}) *LinkAction {
	a.Set("body", value)
	return a
}

func (a *LinkAction) ClassName(value interface{}) *LinkAction {
	a.Set("className", value)
	return a
}

func (a *LinkAction) Close(value interface{}) *LinkAction {
	a.Set("close", value)
	return a
}

func (a *LinkAction) ConfirmText(value interface{}) *LinkAction {
	a.Set("confirmText", value)
	return a
}

func (a *LinkAction) ConfirmTitle(value interface{}) *LinkAction {
	a.Set("confirmTitle", value)
	return a
}

func (a *LinkAction) CountDown(value interface{}) *LinkAction {
	a.Set("countDown", value)
	return a
}

func (a *LinkAction) CountDownTpl(value interface{}) *LinkAction {
	a.Set("countDownTpl", value)
	return a
}

func (a *LinkAction) Disabled(value interface{}) *LinkAction {
	a.Set("disabled", value)
	return a
}

func (a *LinkAction) DisabledOn(value interface{}) *LinkAction {
	a.Set("disabledOn", value)
	return a
}

func (a *LinkAction) DisabledOnAction(value interface{}) *LinkAction {
	a.Set("disabledOnAction", value)
	return a
}

func (a *LinkAction) DisabledTip(value interface{}) *LinkAction {
	a.Set("disabledTip", value)
	return a
}

func (a *LinkAction) EditorSetting(value interface{}) *LinkAction {
	a.Set("editorSetting", value)
	return a
}

func (a *LinkAction) Hidden(value interface{}) *LinkAction {
	a.Set("hidden", value)
	return a
}

func (a *LinkAction) HiddenOn(value interface{}) *LinkAction {
	a.Set("hiddenOn", value)
	return a
}

func (a *LinkAction) HotKey(value interface{}) *LinkAction {
	a.Set("hotKey", value)
	return a
}

func (a *LinkAction) Icon(value interface{}) *LinkAction {
	a.Set("icon", value)
	return a
}

func (a *LinkAction) IconClassName(value interface{}) *LinkAction {
	a.Set("iconClassName", value)
	return a
}

func (a *LinkAction) Id(value interface{}) *LinkAction {
	a.Set("id", value)
	return a
}

func (a *LinkAction) Label(value interface{}) *LinkAction {
	a.Set("label", value)
	return a
}

func (a *LinkAction) Level(value interface{}) *LinkAction {
	a.Set("level", value)
	return a
}

func (a *LinkAction) Link(value interface{}) *LinkAction {
	a.Set("link", value)
	return a
}

func (a *LinkAction) LoadingClassName(value interface{}) *LinkAction {
	a.Set("loadingClassName", value)
	return a
}

func (a *LinkAction) LoadingOn(value interface{}) *LinkAction {
	a.Set("loadingOn", value)
	return a
}

func (a *LinkAction) MergeData(value interface{}) *LinkAction {
	a.Set("mergeData", value)
	return a
}

func (a *LinkAction) OnClick(value interface{}) *LinkAction {
	a.Set("onClick", value)
	return a
}

func (a *LinkAction) OnEvent(value interface{}) *LinkAction {
	a.Set("onEvent", value)
	return a
}

func (a *LinkAction) Primary(value interface{}) *LinkAction {
	a.Set("primary", value)
	return a
}

func (a *LinkAction) RequireSelected(value interface{}) *LinkAction {
	a.Set("requireSelected", value)
	return a
}

func (a *LinkAction) Required(value interface{}) *LinkAction {
	a.Set("required", value)
	return a
}

func (a *LinkAction) RightIcon(value interface{}) *LinkAction {
	a.Set("rightIcon", value)
	return a
}

func (a *LinkAction) RightIconClassName(value interface{}) *LinkAction {
	a.Set("rightIconClassName", value)
	return a
}

func (a *LinkAction) Size(value interface{}) *LinkAction {
	a.Set("size", value)
	return a
}

func (a *LinkAction) Static(value interface{}) *LinkAction {
	a.Set("static", value)
	return a
}

func (a *LinkAction) StaticClassName(value interface{}) *LinkAction {
	a.Set("staticClassName", value)
	return a
}

func (a *LinkAction) StaticInputClassName(value interface{}) *LinkAction {
	a.Set("staticInputClassName", value)
	return a
}

func (a *LinkAction) StaticLabelClassName(value interface{}) *LinkAction {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *LinkAction) StaticOn(value interface{}) *LinkAction {
	a.Set("staticOn", value)
	return a
}

func (a *LinkAction) StaticPlaceholder(value interface{}) *LinkAction {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *LinkAction) StaticSchema(value interface{}) *LinkAction {
	a.Set("staticSchema", value)
	return a
}

func (a *LinkAction) Style(value interface{}) *LinkAction {
	a.Set("style", value)
	return a
}

func (a *LinkAction) TabIndex(value interface{}) *LinkAction {
	a.Set("tabIndex", value)
	return a
}

func (a *LinkAction) Target(value interface{}) *LinkAction {
	a.Set("target", value)
	return a
}

func (a *LinkAction) TestIdBuilder(value interface{}) *LinkAction {
	a.Set("testIdBuilder", value)
	return a
}

func (a *LinkAction) Testid(value interface{}) *LinkAction {
	a.Set("testid", value)
	return a
}

func (a *LinkAction) Tooltip(value interface{}) *LinkAction {
	a.Set("tooltip", value)
	return a
}

func (a *LinkAction) TooltipPlacement(value interface{}) *LinkAction {
	a.Set("tooltipPlacement", value)
	return a
}

func (a *LinkAction) Type(value interface{}) *LinkAction {
	a.Set("type", value)
	return a
}

func (a *LinkAction) UseMobileUI(value interface{}) *LinkAction {
	a.Set("useMobileUI", value)
	return a
}

func (a *LinkAction) Visible(value interface{}) *LinkAction {
	a.Set("visible", value)
	return a
}

func (a *LinkAction) VisibleOn(value interface{}) *LinkAction {
	a.Set("visibleOn", value)
	return a
}
