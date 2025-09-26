package renderers

type AjaxAction struct {
	*BaseRenderer
}

func NewAjaxAction() *AjaxAction {
	a := &AjaxAction{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "button")
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

func (a *AjaxAction) ActionType(value interface{}) *AjaxAction {
	a.Set("actionType", value)
	return a
}

func (a *AjaxAction) ActiveClassName(value interface{}) *AjaxAction {
	a.Set("activeClassName", value)
	return a
}

func (a *AjaxAction) ActiveLevel(value interface{}) *AjaxAction {
	a.Set("activeLevel", value)
	return a
}

func (a *AjaxAction) Api(value interface{}) *AjaxAction {
	a.Set("api", value)
	return a
}

func (a *AjaxAction) Badge(value interface{}) *AjaxAction {
	a.Set("badge", value)
	return a
}

func (a *AjaxAction) Block(value interface{}) *AjaxAction {
	a.Set("block", value)
	return a
}

func (a *AjaxAction) Body(value interface{}) *AjaxAction {
	a.Set("body", value)
	return a
}

func (a *AjaxAction) ClassName(value interface{}) *AjaxAction {
	a.Set("className", value)
	return a
}

func (a *AjaxAction) Close(value interface{}) *AjaxAction {
	a.Set("close", value)
	return a
}

func (a *AjaxAction) ConfirmText(value interface{}) *AjaxAction {
	a.Set("confirmText", value)
	return a
}

func (a *AjaxAction) ConfirmTitle(value interface{}) *AjaxAction {
	a.Set("confirmTitle", value)
	return a
}

func (a *AjaxAction) CountDown(value interface{}) *AjaxAction {
	a.Set("countDown", value)
	return a
}

func (a *AjaxAction) CountDownTpl(value interface{}) *AjaxAction {
	a.Set("countDownTpl", value)
	return a
}

func (a *AjaxAction) Disabled(value interface{}) *AjaxAction {
	a.Set("disabled", value)
	return a
}

func (a *AjaxAction) DisabledOn(value interface{}) *AjaxAction {
	a.Set("disabledOn", value)
	return a
}

func (a *AjaxAction) DisabledOnAction(value interface{}) *AjaxAction {
	a.Set("disabledOnAction", value)
	return a
}

func (a *AjaxAction) DisabledTip(value interface{}) *AjaxAction {
	a.Set("disabledTip", value)
	return a
}

func (a *AjaxAction) EditorSetting(value interface{}) *AjaxAction {
	a.Set("editorSetting", value)
	return a
}

func (a *AjaxAction) Feedback(value interface{}) *AjaxAction {
	a.Set("feedback", value)
	return a
}

func (a *AjaxAction) Hidden(value interface{}) *AjaxAction {
	a.Set("hidden", value)
	return a
}

func (a *AjaxAction) HiddenOn(value interface{}) *AjaxAction {
	a.Set("hiddenOn", value)
	return a
}

func (a *AjaxAction) HotKey(value interface{}) *AjaxAction {
	a.Set("hotKey", value)
	return a
}

func (a *AjaxAction) Icon(value interface{}) *AjaxAction {
	a.Set("icon", value)
	return a
}

func (a *AjaxAction) IconClassName(value interface{}) *AjaxAction {
	a.Set("iconClassName", value)
	return a
}

func (a *AjaxAction) Id(value interface{}) *AjaxAction {
	a.Set("id", value)
	return a
}

func (a *AjaxAction) IgnoreConfirm(value interface{}) *AjaxAction {
	a.Set("ignoreConfirm", value)
	return a
}

func (a *AjaxAction) IsolateScope(value interface{}) *AjaxAction {
	a.Set("isolateScope", value)
	return a
}

func (a *AjaxAction) Label(value interface{}) *AjaxAction {
	a.Set("label", value)
	return a
}

func (a *AjaxAction) Level(value interface{}) *AjaxAction {
	a.Set("level", value)
	return a
}

func (a *AjaxAction) LoadingClassName(value interface{}) *AjaxAction {
	a.Set("loadingClassName", value)
	return a
}

func (a *AjaxAction) LoadingOn(value interface{}) *AjaxAction {
	a.Set("loadingOn", value)
	return a
}

func (a *AjaxAction) MergeData(value interface{}) *AjaxAction {
	a.Set("mergeData", value)
	return a
}

func (a *AjaxAction) OnClick(value interface{}) *AjaxAction {
	a.Set("onClick", value)
	return a
}

func (a *AjaxAction) OnEvent(value interface{}) *AjaxAction {
	a.Set("onEvent", value)
	return a
}

func (a *AjaxAction) Primary(value interface{}) *AjaxAction {
	a.Set("primary", value)
	return a
}

func (a *AjaxAction) Redirect(value interface{}) *AjaxAction {
	a.Set("redirect", value)
	return a
}

func (a *AjaxAction) Reload(value interface{}) *AjaxAction {
	a.Set("reload", value)
	return a
}

func (a *AjaxAction) RequireSelected(value interface{}) *AjaxAction {
	a.Set("requireSelected", value)
	return a
}

func (a *AjaxAction) Required(value interface{}) *AjaxAction {
	a.Set("required", value)
	return a
}

func (a *AjaxAction) RightIcon(value interface{}) *AjaxAction {
	a.Set("rightIcon", value)
	return a
}

func (a *AjaxAction) RightIconClassName(value interface{}) *AjaxAction {
	a.Set("rightIconClassName", value)
	return a
}

func (a *AjaxAction) Size(value interface{}) *AjaxAction {
	a.Set("size", value)
	return a
}

func (a *AjaxAction) Static(value interface{}) *AjaxAction {
	a.Set("static", value)
	return a
}

func (a *AjaxAction) StaticClassName(value interface{}) *AjaxAction {
	a.Set("staticClassName", value)
	return a
}

func (a *AjaxAction) StaticInputClassName(value interface{}) *AjaxAction {
	a.Set("staticInputClassName", value)
	return a
}

func (a *AjaxAction) StaticLabelClassName(value interface{}) *AjaxAction {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *AjaxAction) StaticOn(value interface{}) *AjaxAction {
	a.Set("staticOn", value)
	return a
}

func (a *AjaxAction) StaticPlaceholder(value interface{}) *AjaxAction {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *AjaxAction) StaticSchema(value interface{}) *AjaxAction {
	a.Set("staticSchema", value)
	return a
}

func (a *AjaxAction) Style(value interface{}) *AjaxAction {
	a.Set("style", value)
	return a
}

func (a *AjaxAction) TabIndex(value interface{}) *AjaxAction {
	a.Set("tabIndex", value)
	return a
}

func (a *AjaxAction) Target(value interface{}) *AjaxAction {
	a.Set("target", value)
	return a
}

func (a *AjaxAction) TestIdBuilder(value interface{}) *AjaxAction {
	a.Set("testIdBuilder", value)
	return a
}

func (a *AjaxAction) Testid(value interface{}) *AjaxAction {
	a.Set("testid", value)
	return a
}

func (a *AjaxAction) Tooltip(value interface{}) *AjaxAction {
	a.Set("tooltip", value)
	return a
}

func (a *AjaxAction) TooltipPlacement(value interface{}) *AjaxAction {
	a.Set("tooltipPlacement", value)
	return a
}

func (a *AjaxAction) Type(value interface{}) *AjaxAction {
	a.Set("type", value)
	return a
}

func (a *AjaxAction) UseMobileUI(value interface{}) *AjaxAction {
	a.Set("useMobileUI", value)
	return a
}

func (a *AjaxAction) Visible(value interface{}) *AjaxAction {
	a.Set("visible", value)
	return a
}

func (a *AjaxAction) VisibleOn(value interface{}) *AjaxAction {
	a.Set("visibleOn", value)
	return a
}
