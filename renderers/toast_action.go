package renderers

type ToastAction struct {
	*BaseRenderer
}

func NewToastAction() *ToastAction {
	a := &ToastAction{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "button")
	return a
}

func (a *ToastAction) Set(name string, value interface{}) *ToastAction {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *ToastAction) ActionType(value interface{}) *ToastAction {
	a.Set("actionType", value)
	return a
}

func (a *ToastAction) ActiveClassName(value interface{}) *ToastAction {
	a.Set("activeClassName", value)
	return a
}

func (a *ToastAction) ActiveLevel(value interface{}) *ToastAction {
	a.Set("activeLevel", value)
	return a
}

func (a *ToastAction) Badge(value interface{}) *ToastAction {
	a.Set("badge", value)
	return a
}

func (a *ToastAction) Block(value interface{}) *ToastAction {
	a.Set("block", value)
	return a
}

func (a *ToastAction) Body(value interface{}) *ToastAction {
	a.Set("body", value)
	return a
}

func (a *ToastAction) ClassName(value interface{}) *ToastAction {
	a.Set("className", value)
	return a
}

func (a *ToastAction) Close(value interface{}) *ToastAction {
	a.Set("close", value)
	return a
}

func (a *ToastAction) ConfirmText(value interface{}) *ToastAction {
	a.Set("confirmText", value)
	return a
}

func (a *ToastAction) ConfirmTitle(value interface{}) *ToastAction {
	a.Set("confirmTitle", value)
	return a
}

func (a *ToastAction) CountDown(value interface{}) *ToastAction {
	a.Set("countDown", value)
	return a
}

func (a *ToastAction) CountDownTpl(value interface{}) *ToastAction {
	a.Set("countDownTpl", value)
	return a
}

func (a *ToastAction) Disabled(value interface{}) *ToastAction {
	a.Set("disabled", value)
	return a
}

func (a *ToastAction) DisabledOn(value interface{}) *ToastAction {
	a.Set("disabledOn", value)
	return a
}

func (a *ToastAction) DisabledOnAction(value interface{}) *ToastAction {
	a.Set("disabledOnAction", value)
	return a
}

func (a *ToastAction) DisabledTip(value interface{}) *ToastAction {
	a.Set("disabledTip", value)
	return a
}

func (a *ToastAction) EditorSetting(value interface{}) *ToastAction {
	a.Set("editorSetting", value)
	return a
}

func (a *ToastAction) Hidden(value interface{}) *ToastAction {
	a.Set("hidden", value)
	return a
}

func (a *ToastAction) HiddenOn(value interface{}) *ToastAction {
	a.Set("hiddenOn", value)
	return a
}

func (a *ToastAction) HotKey(value interface{}) *ToastAction {
	a.Set("hotKey", value)
	return a
}

func (a *ToastAction) Icon(value interface{}) *ToastAction {
	a.Set("icon", value)
	return a
}

func (a *ToastAction) IconClassName(value interface{}) *ToastAction {
	a.Set("iconClassName", value)
	return a
}

func (a *ToastAction) Id(value interface{}) *ToastAction {
	a.Set("id", value)
	return a
}

func (a *ToastAction) Label(value interface{}) *ToastAction {
	a.Set("label", value)
	return a
}

func (a *ToastAction) Level(value interface{}) *ToastAction {
	a.Set("level", value)
	return a
}

func (a *ToastAction) LoadingClassName(value interface{}) *ToastAction {
	a.Set("loadingClassName", value)
	return a
}

func (a *ToastAction) LoadingOn(value interface{}) *ToastAction {
	a.Set("loadingOn", value)
	return a
}

func (a *ToastAction) MergeData(value interface{}) *ToastAction {
	a.Set("mergeData", value)
	return a
}

func (a *ToastAction) OnClick(value interface{}) *ToastAction {
	a.Set("onClick", value)
	return a
}

func (a *ToastAction) OnEvent(value interface{}) *ToastAction {
	a.Set("onEvent", value)
	return a
}

func (a *ToastAction) Primary(value interface{}) *ToastAction {
	a.Set("primary", value)
	return a
}

func (a *ToastAction) RequireSelected(value interface{}) *ToastAction {
	a.Set("requireSelected", value)
	return a
}

func (a *ToastAction) Required(value interface{}) *ToastAction {
	a.Set("required", value)
	return a
}

func (a *ToastAction) RightIcon(value interface{}) *ToastAction {
	a.Set("rightIcon", value)
	return a
}

func (a *ToastAction) RightIconClassName(value interface{}) *ToastAction {
	a.Set("rightIconClassName", value)
	return a
}

func (a *ToastAction) Size(value interface{}) *ToastAction {
	a.Set("size", value)
	return a
}

func (a *ToastAction) Static(value interface{}) *ToastAction {
	a.Set("static", value)
	return a
}

func (a *ToastAction) StaticClassName(value interface{}) *ToastAction {
	a.Set("staticClassName", value)
	return a
}

func (a *ToastAction) StaticInputClassName(value interface{}) *ToastAction {
	a.Set("staticInputClassName", value)
	return a
}

func (a *ToastAction) StaticLabelClassName(value interface{}) *ToastAction {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *ToastAction) StaticOn(value interface{}) *ToastAction {
	a.Set("staticOn", value)
	return a
}

func (a *ToastAction) StaticPlaceholder(value interface{}) *ToastAction {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *ToastAction) StaticSchema(value interface{}) *ToastAction {
	a.Set("staticSchema", value)
	return a
}

func (a *ToastAction) Style(value interface{}) *ToastAction {
	a.Set("style", value)
	return a
}

func (a *ToastAction) TabIndex(value interface{}) *ToastAction {
	a.Set("tabIndex", value)
	return a
}

func (a *ToastAction) Target(value interface{}) *ToastAction {
	a.Set("target", value)
	return a
}

func (a *ToastAction) TestIdBuilder(value interface{}) *ToastAction {
	a.Set("testIdBuilder", value)
	return a
}

func (a *ToastAction) Testid(value interface{}) *ToastAction {
	a.Set("testid", value)
	return a
}

func (a *ToastAction) Toast(value interface{}) *ToastAction {
	a.Set("toast", value)
	return a
}

func (a *ToastAction) Tooltip(value interface{}) *ToastAction {
	a.Set("tooltip", value)
	return a
}

func (a *ToastAction) TooltipPlacement(value interface{}) *ToastAction {
	a.Set("tooltipPlacement", value)
	return a
}

func (a *ToastAction) Type(value interface{}) *ToastAction {
	a.Set("type", value)
	return a
}

func (a *ToastAction) UseMobileUI(value interface{}) *ToastAction {
	a.Set("useMobileUI", value)
	return a
}

func (a *ToastAction) Visible(value interface{}) *ToastAction {
	a.Set("visible", value)
	return a
}

func (a *ToastAction) VisibleOn(value interface{}) *ToastAction {
	a.Set("visibleOn", value)
	return a
}
