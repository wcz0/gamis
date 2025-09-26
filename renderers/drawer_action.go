package renderers

type DrawerAction struct {
	*BaseRenderer
}

func NewDrawerAction() *DrawerAction {
	a := &DrawerAction{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "button")
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

func (a *DrawerAction) ActionType(value interface{}) *DrawerAction {
	a.Set("actionType", value)
	return a
}

func (a *DrawerAction) ActiveClassName(value interface{}) *DrawerAction {
	a.Set("activeClassName", value)
	return a
}

func (a *DrawerAction) ActiveLevel(value interface{}) *DrawerAction {
	a.Set("activeLevel", value)
	return a
}

func (a *DrawerAction) Align(value interface{}) *DrawerAction {
	a.Set("align", value)
	return a
}

func (a *DrawerAction) Badge(value interface{}) *DrawerAction {
	a.Set("badge", value)
	return a
}

func (a *DrawerAction) Block(value interface{}) *DrawerAction {
	a.Set("block", value)
	return a
}

func (a *DrawerAction) Body(value interface{}) *DrawerAction {
	a.Set("body", value)
	return a
}

func (a *DrawerAction) ClassName(value interface{}) *DrawerAction {
	a.Set("className", value)
	return a
}

func (a *DrawerAction) Close(value interface{}) *DrawerAction {
	a.Set("close", value)
	return a
}

func (a *DrawerAction) ConfirmText(value interface{}) *DrawerAction {
	a.Set("confirmText", value)
	return a
}

func (a *DrawerAction) ConfirmTitle(value interface{}) *DrawerAction {
	a.Set("confirmTitle", value)
	return a
}

func (a *DrawerAction) CountDown(value interface{}) *DrawerAction {
	a.Set("countDown", value)
	return a
}

func (a *DrawerAction) CountDownTpl(value interface{}) *DrawerAction {
	a.Set("countDownTpl", value)
	return a
}

func (a *DrawerAction) Data(value interface{}) *DrawerAction {
	a.Set("data", value)
	return a
}

func (a *DrawerAction) Disabled(value interface{}) *DrawerAction {
	a.Set("disabled", value)
	return a
}

func (a *DrawerAction) DisabledOn(value interface{}) *DrawerAction {
	a.Set("disabledOn", value)
	return a
}

func (a *DrawerAction) DisabledOnAction(value interface{}) *DrawerAction {
	a.Set("disabledOnAction", value)
	return a
}

func (a *DrawerAction) DisabledTip(value interface{}) *DrawerAction {
	a.Set("disabledTip", value)
	return a
}

func (a *DrawerAction) Drawer(value interface{}) *DrawerAction {
	a.Set("drawer", value)
	return a
}

func (a *DrawerAction) EditorSetting(value interface{}) *DrawerAction {
	a.Set("editorSetting", value)
	return a
}

func (a *DrawerAction) Hidden(value interface{}) *DrawerAction {
	a.Set("hidden", value)
	return a
}

func (a *DrawerAction) HiddenOn(value interface{}) *DrawerAction {
	a.Set("hiddenOn", value)
	return a
}

func (a *DrawerAction) HotKey(value interface{}) *DrawerAction {
	a.Set("hotKey", value)
	return a
}

func (a *DrawerAction) Icon(value interface{}) *DrawerAction {
	a.Set("icon", value)
	return a
}

func (a *DrawerAction) IconClassName(value interface{}) *DrawerAction {
	a.Set("iconClassName", value)
	return a
}

func (a *DrawerAction) Id(value interface{}) *DrawerAction {
	a.Set("id", value)
	return a
}

func (a *DrawerAction) Label(value interface{}) *DrawerAction {
	a.Set("label", value)
	return a
}

func (a *DrawerAction) Level(value interface{}) *DrawerAction {
	a.Set("level", value)
	return a
}

func (a *DrawerAction) LoadingClassName(value interface{}) *DrawerAction {
	a.Set("loadingClassName", value)
	return a
}

func (a *DrawerAction) LoadingOn(value interface{}) *DrawerAction {
	a.Set("loadingOn", value)
	return a
}

func (a *DrawerAction) MergeData(value interface{}) *DrawerAction {
	a.Set("mergeData", value)
	return a
}

func (a *DrawerAction) NextCondition(value interface{}) *DrawerAction {
	a.Set("nextCondition", value)
	return a
}

func (a *DrawerAction) OnClick(value interface{}) *DrawerAction {
	a.Set("onClick", value)
	return a
}

func (a *DrawerAction) OnEvent(value interface{}) *DrawerAction {
	a.Set("onEvent", value)
	return a
}

func (a *DrawerAction) Primary(value interface{}) *DrawerAction {
	a.Set("primary", value)
	return a
}

func (a *DrawerAction) Redirect(value interface{}) *DrawerAction {
	a.Set("redirect", value)
	return a
}

func (a *DrawerAction) Reload(value interface{}) *DrawerAction {
	a.Set("reload", value)
	return a
}

func (a *DrawerAction) RequireSelected(value interface{}) *DrawerAction {
	a.Set("requireSelected", value)
	return a
}

func (a *DrawerAction) Required(value interface{}) *DrawerAction {
	a.Set("required", value)
	return a
}

func (a *DrawerAction) RightIcon(value interface{}) *DrawerAction {
	a.Set("rightIcon", value)
	return a
}

func (a *DrawerAction) RightIconClassName(value interface{}) *DrawerAction {
	a.Set("rightIconClassName", value)
	return a
}

func (a *DrawerAction) Size(value interface{}) *DrawerAction {
	a.Set("size", value)
	return a
}

func (a *DrawerAction) Static(value interface{}) *DrawerAction {
	a.Set("static", value)
	return a
}

func (a *DrawerAction) StaticClassName(value interface{}) *DrawerAction {
	a.Set("staticClassName", value)
	return a
}

func (a *DrawerAction) StaticInputClassName(value interface{}) *DrawerAction {
	a.Set("staticInputClassName", value)
	return a
}

func (a *DrawerAction) StaticLabelClassName(value interface{}) *DrawerAction {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *DrawerAction) StaticOn(value interface{}) *DrawerAction {
	a.Set("staticOn", value)
	return a
}

func (a *DrawerAction) StaticPlaceholder(value interface{}) *DrawerAction {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *DrawerAction) StaticSchema(value interface{}) *DrawerAction {
	a.Set("staticSchema", value)
	return a
}

func (a *DrawerAction) Style(value interface{}) *DrawerAction {
	a.Set("style", value)
	return a
}

func (a *DrawerAction) TabIndex(value interface{}) *DrawerAction {
	a.Set("tabIndex", value)
	return a
}

func (a *DrawerAction) Target(value interface{}) *DrawerAction {
	a.Set("target", value)
	return a
}

func (a *DrawerAction) TestIdBuilder(value interface{}) *DrawerAction {
	a.Set("testIdBuilder", value)
	return a
}

func (a *DrawerAction) Testid(value interface{}) *DrawerAction {
	a.Set("testid", value)
	return a
}

func (a *DrawerAction) Tooltip(value interface{}) *DrawerAction {
	a.Set("tooltip", value)
	return a
}

func (a *DrawerAction) TooltipPlacement(value interface{}) *DrawerAction {
	a.Set("tooltipPlacement", value)
	return a
}

func (a *DrawerAction) Type(value interface{}) *DrawerAction {
	a.Set("type", value)
	return a
}

func (a *DrawerAction) UseMobileUI(value interface{}) *DrawerAction {
	a.Set("useMobileUI", value)
	return a
}

func (a *DrawerAction) Visible(value interface{}) *DrawerAction {
	a.Set("visible", value)
	return a
}

func (a *DrawerAction) VisibleOn(value interface{}) *DrawerAction {
	a.Set("visibleOn", value)
	return a
}
