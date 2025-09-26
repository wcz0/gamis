package renderers

type UrlAction struct {
	*BaseRenderer
}

func NewUrlAction() *UrlAction {
	a := &UrlAction{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "button")
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

func (a *UrlAction) ActionType(value interface{}) *UrlAction {
	a.Set("actionType", value)
	return a
}

func (a *UrlAction) ActiveClassName(value interface{}) *UrlAction {
	a.Set("activeClassName", value)
	return a
}

func (a *UrlAction) ActiveLevel(value interface{}) *UrlAction {
	a.Set("activeLevel", value)
	return a
}

func (a *UrlAction) Badge(value interface{}) *UrlAction {
	a.Set("badge", value)
	return a
}

func (a *UrlAction) Blank(value interface{}) *UrlAction {
	a.Set("blank", value)
	return a
}

func (a *UrlAction) Block(value interface{}) *UrlAction {
	a.Set("block", value)
	return a
}

func (a *UrlAction) Body(value interface{}) *UrlAction {
	a.Set("body", value)
	return a
}

func (a *UrlAction) ClassName(value interface{}) *UrlAction {
	a.Set("className", value)
	return a
}

func (a *UrlAction) Close(value interface{}) *UrlAction {
	a.Set("close", value)
	return a
}

func (a *UrlAction) ConfirmText(value interface{}) *UrlAction {
	a.Set("confirmText", value)
	return a
}

func (a *UrlAction) ConfirmTitle(value interface{}) *UrlAction {
	a.Set("confirmTitle", value)
	return a
}

func (a *UrlAction) CountDown(value interface{}) *UrlAction {
	a.Set("countDown", value)
	return a
}

func (a *UrlAction) CountDownTpl(value interface{}) *UrlAction {
	a.Set("countDownTpl", value)
	return a
}

func (a *UrlAction) Disabled(value interface{}) *UrlAction {
	a.Set("disabled", value)
	return a
}

func (a *UrlAction) DisabledOn(value interface{}) *UrlAction {
	a.Set("disabledOn", value)
	return a
}

func (a *UrlAction) DisabledOnAction(value interface{}) *UrlAction {
	a.Set("disabledOnAction", value)
	return a
}

func (a *UrlAction) DisabledTip(value interface{}) *UrlAction {
	a.Set("disabledTip", value)
	return a
}

func (a *UrlAction) EditorSetting(value interface{}) *UrlAction {
	a.Set("editorSetting", value)
	return a
}

func (a *UrlAction) Hidden(value interface{}) *UrlAction {
	a.Set("hidden", value)
	return a
}

func (a *UrlAction) HiddenOn(value interface{}) *UrlAction {
	a.Set("hiddenOn", value)
	return a
}

func (a *UrlAction) HotKey(value interface{}) *UrlAction {
	a.Set("hotKey", value)
	return a
}

func (a *UrlAction) Icon(value interface{}) *UrlAction {
	a.Set("icon", value)
	return a
}

func (a *UrlAction) IconClassName(value interface{}) *UrlAction {
	a.Set("iconClassName", value)
	return a
}

func (a *UrlAction) Id(value interface{}) *UrlAction {
	a.Set("id", value)
	return a
}

func (a *UrlAction) Label(value interface{}) *UrlAction {
	a.Set("label", value)
	return a
}

func (a *UrlAction) Level(value interface{}) *UrlAction {
	a.Set("level", value)
	return a
}

func (a *UrlAction) Link(value interface{}) *UrlAction {
	a.Set("link", value)
	return a
}

func (a *UrlAction) LoadingClassName(value interface{}) *UrlAction {
	a.Set("loadingClassName", value)
	return a
}

func (a *UrlAction) LoadingOn(value interface{}) *UrlAction {
	a.Set("loadingOn", value)
	return a
}

func (a *UrlAction) MergeData(value interface{}) *UrlAction {
	a.Set("mergeData", value)
	return a
}

func (a *UrlAction) OnClick(value interface{}) *UrlAction {
	a.Set("onClick", value)
	return a
}

func (a *UrlAction) OnEvent(value interface{}) *UrlAction {
	a.Set("onEvent", value)
	return a
}

func (a *UrlAction) Primary(value interface{}) *UrlAction {
	a.Set("primary", value)
	return a
}

func (a *UrlAction) RequireSelected(value interface{}) *UrlAction {
	a.Set("requireSelected", value)
	return a
}

func (a *UrlAction) Required(value interface{}) *UrlAction {
	a.Set("required", value)
	return a
}

func (a *UrlAction) RightIcon(value interface{}) *UrlAction {
	a.Set("rightIcon", value)
	return a
}

func (a *UrlAction) RightIconClassName(value interface{}) *UrlAction {
	a.Set("rightIconClassName", value)
	return a
}

func (a *UrlAction) Size(value interface{}) *UrlAction {
	a.Set("size", value)
	return a
}

func (a *UrlAction) Static(value interface{}) *UrlAction {
	a.Set("static", value)
	return a
}

func (a *UrlAction) StaticClassName(value interface{}) *UrlAction {
	a.Set("staticClassName", value)
	return a
}

func (a *UrlAction) StaticInputClassName(value interface{}) *UrlAction {
	a.Set("staticInputClassName", value)
	return a
}

func (a *UrlAction) StaticLabelClassName(value interface{}) *UrlAction {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *UrlAction) StaticOn(value interface{}) *UrlAction {
	a.Set("staticOn", value)
	return a
}

func (a *UrlAction) StaticPlaceholder(value interface{}) *UrlAction {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *UrlAction) StaticSchema(value interface{}) *UrlAction {
	a.Set("staticSchema", value)
	return a
}

func (a *UrlAction) Style(value interface{}) *UrlAction {
	a.Set("style", value)
	return a
}

func (a *UrlAction) TabIndex(value interface{}) *UrlAction {
	a.Set("tabIndex", value)
	return a
}

func (a *UrlAction) Target(value interface{}) *UrlAction {
	a.Set("target", value)
	return a
}

func (a *UrlAction) TestIdBuilder(value interface{}) *UrlAction {
	a.Set("testIdBuilder", value)
	return a
}

func (a *UrlAction) Testid(value interface{}) *UrlAction {
	a.Set("testid", value)
	return a
}

func (a *UrlAction) Tooltip(value interface{}) *UrlAction {
	a.Set("tooltip", value)
	return a
}

func (a *UrlAction) TooltipPlacement(value interface{}) *UrlAction {
	a.Set("tooltipPlacement", value)
	return a
}

func (a *UrlAction) Type(value interface{}) *UrlAction {
	a.Set("type", value)
	return a
}

func (a *UrlAction) Url(value interface{}) *UrlAction {
	a.Set("url", value)
	return a
}

func (a *UrlAction) UseMobileUI(value interface{}) *UrlAction {
	a.Set("useMobileUI", value)
	return a
}

func (a *UrlAction) Visible(value interface{}) *UrlAction {
	a.Set("visible", value)
	return a
}

func (a *UrlAction) VisibleOn(value interface{}) *UrlAction {
	a.Set("visibleOn", value)
	return a
}
