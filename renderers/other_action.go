package renderers

type OtherAction struct {
	*BaseRenderer
}

func NewOtherAction() *OtherAction {
	a := &OtherAction{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "button")
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

func (a *OtherAction) ActionType(value interface{}) *OtherAction {
	a.Set("actionType", value)
	return a
}

func (a *OtherAction) ActiveClassName(value interface{}) *OtherAction {
	a.Set("activeClassName", value)
	return a
}

func (a *OtherAction) ActiveLevel(value interface{}) *OtherAction {
	a.Set("activeLevel", value)
	return a
}

func (a *OtherAction) Badge(value interface{}) *OtherAction {
	a.Set("badge", value)
	return a
}

func (a *OtherAction) Block(value interface{}) *OtherAction {
	a.Set("block", value)
	return a
}

func (a *OtherAction) Body(value interface{}) *OtherAction {
	a.Set("body", value)
	return a
}

func (a *OtherAction) ClassName(value interface{}) *OtherAction {
	a.Set("className", value)
	return a
}

func (a *OtherAction) Close(value interface{}) *OtherAction {
	a.Set("close", value)
	return a
}

func (a *OtherAction) ConfirmText(value interface{}) *OtherAction {
	a.Set("confirmText", value)
	return a
}

func (a *OtherAction) ConfirmTitle(value interface{}) *OtherAction {
	a.Set("confirmTitle", value)
	return a
}

func (a *OtherAction) CountDown(value interface{}) *OtherAction {
	a.Set("countDown", value)
	return a
}

func (a *OtherAction) CountDownTpl(value interface{}) *OtherAction {
	a.Set("countDownTpl", value)
	return a
}

func (a *OtherAction) Disabled(value interface{}) *OtherAction {
	a.Set("disabled", value)
	return a
}

func (a *OtherAction) DisabledOn(value interface{}) *OtherAction {
	a.Set("disabledOn", value)
	return a
}

func (a *OtherAction) DisabledOnAction(value interface{}) *OtherAction {
	a.Set("disabledOnAction", value)
	return a
}

func (a *OtherAction) DisabledTip(value interface{}) *OtherAction {
	a.Set("disabledTip", value)
	return a
}

func (a *OtherAction) EditorSetting(value interface{}) *OtherAction {
	a.Set("editorSetting", value)
	return a
}

func (a *OtherAction) Hidden(value interface{}) *OtherAction {
	a.Set("hidden", value)
	return a
}

func (a *OtherAction) HiddenOn(value interface{}) *OtherAction {
	a.Set("hiddenOn", value)
	return a
}

func (a *OtherAction) HotKey(value interface{}) *OtherAction {
	a.Set("hotKey", value)
	return a
}

func (a *OtherAction) Icon(value interface{}) *OtherAction {
	a.Set("icon", value)
	return a
}

func (a *OtherAction) IconClassName(value interface{}) *OtherAction {
	a.Set("iconClassName", value)
	return a
}

func (a *OtherAction) Id(value interface{}) *OtherAction {
	a.Set("id", value)
	return a
}

func (a *OtherAction) Label(value interface{}) *OtherAction {
	a.Set("label", value)
	return a
}

func (a *OtherAction) Level(value interface{}) *OtherAction {
	a.Set("level", value)
	return a
}

func (a *OtherAction) LoadingClassName(value interface{}) *OtherAction {
	a.Set("loadingClassName", value)
	return a
}

func (a *OtherAction) LoadingOn(value interface{}) *OtherAction {
	a.Set("loadingOn", value)
	return a
}

func (a *OtherAction) MergeData(value interface{}) *OtherAction {
	a.Set("mergeData", value)
	return a
}

func (a *OtherAction) OnClick(value interface{}) *OtherAction {
	a.Set("onClick", value)
	return a
}

func (a *OtherAction) OnEvent(value interface{}) *OtherAction {
	a.Set("onEvent", value)
	return a
}

func (a *OtherAction) Primary(value interface{}) *OtherAction {
	a.Set("primary", value)
	return a
}

func (a *OtherAction) RequireSelected(value interface{}) *OtherAction {
	a.Set("requireSelected", value)
	return a
}

func (a *OtherAction) Required(value interface{}) *OtherAction {
	a.Set("required", value)
	return a
}

func (a *OtherAction) RightIcon(value interface{}) *OtherAction {
	a.Set("rightIcon", value)
	return a
}

func (a *OtherAction) RightIconClassName(value interface{}) *OtherAction {
	a.Set("rightIconClassName", value)
	return a
}

func (a *OtherAction) Size(value interface{}) *OtherAction {
	a.Set("size", value)
	return a
}

func (a *OtherAction) Static(value interface{}) *OtherAction {
	a.Set("static", value)
	return a
}

func (a *OtherAction) StaticClassName(value interface{}) *OtherAction {
	a.Set("staticClassName", value)
	return a
}

func (a *OtherAction) StaticInputClassName(value interface{}) *OtherAction {
	a.Set("staticInputClassName", value)
	return a
}

func (a *OtherAction) StaticLabelClassName(value interface{}) *OtherAction {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *OtherAction) StaticOn(value interface{}) *OtherAction {
	a.Set("staticOn", value)
	return a
}

func (a *OtherAction) StaticPlaceholder(value interface{}) *OtherAction {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *OtherAction) StaticSchema(value interface{}) *OtherAction {
	a.Set("staticSchema", value)
	return a
}

func (a *OtherAction) Style(value interface{}) *OtherAction {
	a.Set("style", value)
	return a
}

func (a *OtherAction) TabIndex(value interface{}) *OtherAction {
	a.Set("tabIndex", value)
	return a
}

func (a *OtherAction) Target(value interface{}) *OtherAction {
	a.Set("target", value)
	return a
}

func (a *OtherAction) TestIdBuilder(value interface{}) *OtherAction {
	a.Set("testIdBuilder", value)
	return a
}

func (a *OtherAction) Testid(value interface{}) *OtherAction {
	a.Set("testid", value)
	return a
}

func (a *OtherAction) Tooltip(value interface{}) *OtherAction {
	a.Set("tooltip", value)
	return a
}

func (a *OtherAction) TooltipPlacement(value interface{}) *OtherAction {
	a.Set("tooltipPlacement", value)
	return a
}

func (a *OtherAction) Type(value interface{}) *OtherAction {
	a.Set("type", value)
	return a
}

func (a *OtherAction) UseMobileUI(value interface{}) *OtherAction {
	a.Set("useMobileUI", value)
	return a
}

func (a *OtherAction) Visible(value interface{}) *OtherAction {
	a.Set("visible", value)
	return a
}

func (a *OtherAction) VisibleOn(value interface{}) *OtherAction {
	a.Set("visibleOn", value)
	return a
}
