package renderers

type ReloadAction struct {
	*BaseRenderer
}

func NewReloadAction() *ReloadAction {
	a := &ReloadAction{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "button")
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

func (a *ReloadAction) ActionType(value interface{}) *ReloadAction {
	a.Set("actionType", value)
	return a
}

func (a *ReloadAction) ActiveClassName(value interface{}) *ReloadAction {
	a.Set("activeClassName", value)
	return a
}

func (a *ReloadAction) ActiveLevel(value interface{}) *ReloadAction {
	a.Set("activeLevel", value)
	return a
}

func (a *ReloadAction) Badge(value interface{}) *ReloadAction {
	a.Set("badge", value)
	return a
}

func (a *ReloadAction) Block(value interface{}) *ReloadAction {
	a.Set("block", value)
	return a
}

func (a *ReloadAction) Body(value interface{}) *ReloadAction {
	a.Set("body", value)
	return a
}

func (a *ReloadAction) ClassName(value interface{}) *ReloadAction {
	a.Set("className", value)
	return a
}

func (a *ReloadAction) Close(value interface{}) *ReloadAction {
	a.Set("close", value)
	return a
}

func (a *ReloadAction) ConfirmText(value interface{}) *ReloadAction {
	a.Set("confirmText", value)
	return a
}

func (a *ReloadAction) ConfirmTitle(value interface{}) *ReloadAction {
	a.Set("confirmTitle", value)
	return a
}

func (a *ReloadAction) CountDown(value interface{}) *ReloadAction {
	a.Set("countDown", value)
	return a
}

func (a *ReloadAction) CountDownTpl(value interface{}) *ReloadAction {
	a.Set("countDownTpl", value)
	return a
}

func (a *ReloadAction) Disabled(value interface{}) *ReloadAction {
	a.Set("disabled", value)
	return a
}

func (a *ReloadAction) DisabledOn(value interface{}) *ReloadAction {
	a.Set("disabledOn", value)
	return a
}

func (a *ReloadAction) DisabledOnAction(value interface{}) *ReloadAction {
	a.Set("disabledOnAction", value)
	return a
}

func (a *ReloadAction) DisabledTip(value interface{}) *ReloadAction {
	a.Set("disabledTip", value)
	return a
}

func (a *ReloadAction) EditorSetting(value interface{}) *ReloadAction {
	a.Set("editorSetting", value)
	return a
}

func (a *ReloadAction) Hidden(value interface{}) *ReloadAction {
	a.Set("hidden", value)
	return a
}

func (a *ReloadAction) HiddenOn(value interface{}) *ReloadAction {
	a.Set("hiddenOn", value)
	return a
}

func (a *ReloadAction) HotKey(value interface{}) *ReloadAction {
	a.Set("hotKey", value)
	return a
}

func (a *ReloadAction) Icon(value interface{}) *ReloadAction {
	a.Set("icon", value)
	return a
}

func (a *ReloadAction) IconClassName(value interface{}) *ReloadAction {
	a.Set("iconClassName", value)
	return a
}

func (a *ReloadAction) Id(value interface{}) *ReloadAction {
	a.Set("id", value)
	return a
}

func (a *ReloadAction) Label(value interface{}) *ReloadAction {
	a.Set("label", value)
	return a
}

func (a *ReloadAction) Level(value interface{}) *ReloadAction {
	a.Set("level", value)
	return a
}

func (a *ReloadAction) LoadingClassName(value interface{}) *ReloadAction {
	a.Set("loadingClassName", value)
	return a
}

func (a *ReloadAction) LoadingOn(value interface{}) *ReloadAction {
	a.Set("loadingOn", value)
	return a
}

func (a *ReloadAction) MergeData(value interface{}) *ReloadAction {
	a.Set("mergeData", value)
	return a
}

func (a *ReloadAction) OnClick(value interface{}) *ReloadAction {
	a.Set("onClick", value)
	return a
}

func (a *ReloadAction) OnEvent(value interface{}) *ReloadAction {
	a.Set("onEvent", value)
	return a
}

func (a *ReloadAction) Primary(value interface{}) *ReloadAction {
	a.Set("primary", value)
	return a
}

func (a *ReloadAction) RequireSelected(value interface{}) *ReloadAction {
	a.Set("requireSelected", value)
	return a
}

func (a *ReloadAction) Required(value interface{}) *ReloadAction {
	a.Set("required", value)
	return a
}

func (a *ReloadAction) RightIcon(value interface{}) *ReloadAction {
	a.Set("rightIcon", value)
	return a
}

func (a *ReloadAction) RightIconClassName(value interface{}) *ReloadAction {
	a.Set("rightIconClassName", value)
	return a
}

func (a *ReloadAction) Size(value interface{}) *ReloadAction {
	a.Set("size", value)
	return a
}

func (a *ReloadAction) Static(value interface{}) *ReloadAction {
	a.Set("static", value)
	return a
}

func (a *ReloadAction) StaticClassName(value interface{}) *ReloadAction {
	a.Set("staticClassName", value)
	return a
}

func (a *ReloadAction) StaticInputClassName(value interface{}) *ReloadAction {
	a.Set("staticInputClassName", value)
	return a
}

func (a *ReloadAction) StaticLabelClassName(value interface{}) *ReloadAction {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *ReloadAction) StaticOn(value interface{}) *ReloadAction {
	a.Set("staticOn", value)
	return a
}

func (a *ReloadAction) StaticPlaceholder(value interface{}) *ReloadAction {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *ReloadAction) StaticSchema(value interface{}) *ReloadAction {
	a.Set("staticSchema", value)
	return a
}

func (a *ReloadAction) Style(value interface{}) *ReloadAction {
	a.Set("style", value)
	return a
}

func (a *ReloadAction) TabIndex(value interface{}) *ReloadAction {
	a.Set("tabIndex", value)
	return a
}

func (a *ReloadAction) Target(value interface{}) *ReloadAction {
	a.Set("target", value)
	return a
}

func (a *ReloadAction) TestIdBuilder(value interface{}) *ReloadAction {
	a.Set("testIdBuilder", value)
	return a
}

func (a *ReloadAction) Testid(value interface{}) *ReloadAction {
	a.Set("testid", value)
	return a
}

func (a *ReloadAction) Tooltip(value interface{}) *ReloadAction {
	a.Set("tooltip", value)
	return a
}

func (a *ReloadAction) TooltipPlacement(value interface{}) *ReloadAction {
	a.Set("tooltipPlacement", value)
	return a
}

func (a *ReloadAction) Type(value interface{}) *ReloadAction {
	a.Set("type", value)
	return a
}

func (a *ReloadAction) UseMobileUI(value interface{}) *ReloadAction {
	a.Set("useMobileUI", value)
	return a
}

func (a *ReloadAction) Visible(value interface{}) *ReloadAction {
	a.Set("visible", value)
	return a
}

func (a *ReloadAction) VisibleOn(value interface{}) *ReloadAction {
	a.Set("visibleOn", value)
	return a
}
