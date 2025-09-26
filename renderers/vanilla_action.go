package renderers

type VanillaAction struct {
	*BaseRenderer
}

func NewVanillaAction() *VanillaAction {
	a := &VanillaAction{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "button")
	return a
}

func (a *VanillaAction) Set(name string, value interface{}) *VanillaAction {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *VanillaAction) ActionType(value interface{}) *VanillaAction {
	a.Set("actionType", value)
	return a
}

func (a *VanillaAction) ActiveClassName(value interface{}) *VanillaAction {
	a.Set("activeClassName", value)
	return a
}

func (a *VanillaAction) ActiveLevel(value interface{}) *VanillaAction {
	a.Set("activeLevel", value)
	return a
}

func (a *VanillaAction) Badge(value interface{}) *VanillaAction {
	a.Set("badge", value)
	return a
}

func (a *VanillaAction) Block(value interface{}) *VanillaAction {
	a.Set("block", value)
	return a
}

func (a *VanillaAction) Body(value interface{}) *VanillaAction {
	a.Set("body", value)
	return a
}

func (a *VanillaAction) ClassName(value interface{}) *VanillaAction {
	a.Set("className", value)
	return a
}

func (a *VanillaAction) Close(value interface{}) *VanillaAction {
	a.Set("close", value)
	return a
}

func (a *VanillaAction) ConfirmText(value interface{}) *VanillaAction {
	a.Set("confirmText", value)
	return a
}

func (a *VanillaAction) ConfirmTitle(value interface{}) *VanillaAction {
	a.Set("confirmTitle", value)
	return a
}

func (a *VanillaAction) CountDown(value interface{}) *VanillaAction {
	a.Set("countDown", value)
	return a
}

func (a *VanillaAction) CountDownTpl(value interface{}) *VanillaAction {
	a.Set("countDownTpl", value)
	return a
}

func (a *VanillaAction) Disabled(value interface{}) *VanillaAction {
	a.Set("disabled", value)
	return a
}

func (a *VanillaAction) DisabledOn(value interface{}) *VanillaAction {
	a.Set("disabledOn", value)
	return a
}

func (a *VanillaAction) DisabledOnAction(value interface{}) *VanillaAction {
	a.Set("disabledOnAction", value)
	return a
}

func (a *VanillaAction) DisabledTip(value interface{}) *VanillaAction {
	a.Set("disabledTip", value)
	return a
}

func (a *VanillaAction) DownloadFileName(value interface{}) *VanillaAction {
	a.Set("downloadFileName", value)
	return a
}

func (a *VanillaAction) EditorSetting(value interface{}) *VanillaAction {
	a.Set("editorSetting", value)
	return a
}

func (a *VanillaAction) Hidden(value interface{}) *VanillaAction {
	a.Set("hidden", value)
	return a
}

func (a *VanillaAction) HiddenOn(value interface{}) *VanillaAction {
	a.Set("hiddenOn", value)
	return a
}

func (a *VanillaAction) HotKey(value interface{}) *VanillaAction {
	a.Set("hotKey", value)
	return a
}

func (a *VanillaAction) Icon(value interface{}) *VanillaAction {
	a.Set("icon", value)
	return a
}

func (a *VanillaAction) IconClassName(value interface{}) *VanillaAction {
	a.Set("iconClassName", value)
	return a
}

func (a *VanillaAction) Id(value interface{}) *VanillaAction {
	a.Set("id", value)
	return a
}

func (a *VanillaAction) Label(value interface{}) *VanillaAction {
	a.Set("label", value)
	return a
}

func (a *VanillaAction) Level(value interface{}) *VanillaAction {
	a.Set("level", value)
	return a
}

func (a *VanillaAction) LoadingClassName(value interface{}) *VanillaAction {
	a.Set("loadingClassName", value)
	return a
}

func (a *VanillaAction) LoadingOn(value interface{}) *VanillaAction {
	a.Set("loadingOn", value)
	return a
}

func (a *VanillaAction) MergeData(value interface{}) *VanillaAction {
	a.Set("mergeData", value)
	return a
}

func (a *VanillaAction) OnClick(value interface{}) *VanillaAction {
	a.Set("onClick", value)
	return a
}

func (a *VanillaAction) OnEvent(value interface{}) *VanillaAction {
	a.Set("onEvent", value)
	return a
}

func (a *VanillaAction) Primary(value interface{}) *VanillaAction {
	a.Set("primary", value)
	return a
}

func (a *VanillaAction) RequireSelected(value interface{}) *VanillaAction {
	a.Set("requireSelected", value)
	return a
}

func (a *VanillaAction) Required(value interface{}) *VanillaAction {
	a.Set("required", value)
	return a
}

func (a *VanillaAction) RightIcon(value interface{}) *VanillaAction {
	a.Set("rightIcon", value)
	return a
}

func (a *VanillaAction) RightIconClassName(value interface{}) *VanillaAction {
	a.Set("rightIconClassName", value)
	return a
}

func (a *VanillaAction) Size(value interface{}) *VanillaAction {
	a.Set("size", value)
	return a
}

func (a *VanillaAction) Static(value interface{}) *VanillaAction {
	a.Set("static", value)
	return a
}

func (a *VanillaAction) StaticClassName(value interface{}) *VanillaAction {
	a.Set("staticClassName", value)
	return a
}

func (a *VanillaAction) StaticInputClassName(value interface{}) *VanillaAction {
	a.Set("staticInputClassName", value)
	return a
}

func (a *VanillaAction) StaticLabelClassName(value interface{}) *VanillaAction {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *VanillaAction) StaticOn(value interface{}) *VanillaAction {
	a.Set("staticOn", value)
	return a
}

func (a *VanillaAction) StaticPlaceholder(value interface{}) *VanillaAction {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *VanillaAction) StaticSchema(value interface{}) *VanillaAction {
	a.Set("staticSchema", value)
	return a
}

func (a *VanillaAction) Style(value interface{}) *VanillaAction {
	a.Set("style", value)
	return a
}

func (a *VanillaAction) TabIndex(value interface{}) *VanillaAction {
	a.Set("tabIndex", value)
	return a
}

func (a *VanillaAction) Target(value interface{}) *VanillaAction {
	a.Set("target", value)
	return a
}

func (a *VanillaAction) TestIdBuilder(value interface{}) *VanillaAction {
	a.Set("testIdBuilder", value)
	return a
}

func (a *VanillaAction) Testid(value interface{}) *VanillaAction {
	a.Set("testid", value)
	return a
}

func (a *VanillaAction) Tooltip(value interface{}) *VanillaAction {
	a.Set("tooltip", value)
	return a
}

func (a *VanillaAction) TooltipPlacement(value interface{}) *VanillaAction {
	a.Set("tooltipPlacement", value)
	return a
}

func (a *VanillaAction) Type(value interface{}) *VanillaAction {
	a.Set("type", value)
	return a
}

func (a *VanillaAction) UseMobileUI(value interface{}) *VanillaAction {
	a.Set("useMobileUI", value)
	return a
}

func (a *VanillaAction) Visible(value interface{}) *VanillaAction {
	a.Set("visible", value)
	return a
}

func (a *VanillaAction) VisibleOn(value interface{}) *VanillaAction {
	a.Set("visibleOn", value)
	return a
}
