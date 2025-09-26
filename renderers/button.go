package renderers

type Button struct {
	*BaseRenderer
}

func NewButton() *Button {
	a := &Button{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "button")
	return a
}

func (a *Button) Set(name string, value interface{}) *Button {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Button) ActionType(value interface{}) *Button {
	a.Set("actionType", value)
	return a
}

func (a *Button) ActiveClassName(value interface{}) *Button {
	a.Set("activeClassName", value)
	return a
}

func (a *Button) ActiveLevel(value interface{}) *Button {
	a.Set("activeLevel", value)
	return a
}

func (a *Button) Badge(value interface{}) *Button {
	a.Set("badge", value)
	return a
}

func (a *Button) Block(value interface{}) *Button {
	a.Set("block", value)
	return a
}

func (a *Button) Body(value interface{}) *Button {
	a.Set("body", value)
	return a
}

func (a *Button) ClassName(value interface{}) *Button {
	a.Set("className", value)
	return a
}

func (a *Button) Close(value interface{}) *Button {
	a.Set("close", value)
	return a
}

func (a *Button) ConfirmText(value interface{}) *Button {
	a.Set("confirmText", value)
	return a
}

func (a *Button) ConfirmTitle(value interface{}) *Button {
	a.Set("confirmTitle", value)
	return a
}

func (a *Button) CountDown(value interface{}) *Button {
	a.Set("countDown", value)
	return a
}

func (a *Button) CountDownTpl(value interface{}) *Button {
	a.Set("countDownTpl", value)
	return a
}

func (a *Button) Disabled(value interface{}) *Button {
	a.Set("disabled", value)
	return a
}

func (a *Button) DisabledOn(value interface{}) *Button {
	a.Set("disabledOn", value)
	return a
}

func (a *Button) DisabledOnAction(value interface{}) *Button {
	a.Set("disabledOnAction", value)
	return a
}

func (a *Button) DisabledTip(value interface{}) *Button {
	a.Set("disabledTip", value)
	return a
}

func (a *Button) DownloadFileName(value interface{}) *Button {
	a.Set("downloadFileName", value)
	return a
}

func (a *Button) EditorSetting(value interface{}) *Button {
	a.Set("editorSetting", value)
	return a
}

func (a *Button) Hidden(value interface{}) *Button {
	a.Set("hidden", value)
	return a
}

func (a *Button) HiddenOn(value interface{}) *Button {
	a.Set("hiddenOn", value)
	return a
}

func (a *Button) HotKey(value interface{}) *Button {
	a.Set("hotKey", value)
	return a
}

func (a *Button) Icon(value interface{}) *Button {
	a.Set("icon", value)
	return a
}

func (a *Button) IconClassName(value interface{}) *Button {
	a.Set("iconClassName", value)
	return a
}

func (a *Button) Id(value interface{}) *Button {
	a.Set("id", value)
	return a
}

func (a *Button) Label(value interface{}) *Button {
	a.Set("label", value)
	return a
}

func (a *Button) Level(value interface{}) *Button {
	a.Set("level", value)
	return a
}

func (a *Button) LoadingClassName(value interface{}) *Button {
	a.Set("loadingClassName", value)
	return a
}

func (a *Button) LoadingOn(value interface{}) *Button {
	a.Set("loadingOn", value)
	return a
}

func (a *Button) MergeData(value interface{}) *Button {
	a.Set("mergeData", value)
	return a
}

func (a *Button) OnClick(value interface{}) *Button {
	a.Set("onClick", value)
	return a
}

func (a *Button) OnEvent(value interface{}) *Button {
	a.Set("onEvent", value)
	return a
}

func (a *Button) Primary(value interface{}) *Button {
	a.Set("primary", value)
	return a
}

func (a *Button) RequireSelected(value interface{}) *Button {
	a.Set("requireSelected", value)
	return a
}

func (a *Button) Required(value interface{}) *Button {
	a.Set("required", value)
	return a
}

func (a *Button) RightIcon(value interface{}) *Button {
	a.Set("rightIcon", value)
	return a
}

func (a *Button) RightIconClassName(value interface{}) *Button {
	a.Set("rightIconClassName", value)
	return a
}

func (a *Button) Size(value interface{}) *Button {
	a.Set("size", value)
	return a
}

func (a *Button) Static(value interface{}) *Button {
	a.Set("static", value)
	return a
}

func (a *Button) StaticClassName(value interface{}) *Button {
	a.Set("staticClassName", value)
	return a
}

func (a *Button) StaticInputClassName(value interface{}) *Button {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Button) StaticLabelClassName(value interface{}) *Button {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Button) StaticOn(value interface{}) *Button {
	a.Set("staticOn", value)
	return a
}

func (a *Button) StaticPlaceholder(value interface{}) *Button {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Button) StaticSchema(value interface{}) *Button {
	a.Set("staticSchema", value)
	return a
}

func (a *Button) Style(value interface{}) *Button {
	a.Set("style", value)
	return a
}

func (a *Button) TabIndex(value interface{}) *Button {
	a.Set("tabIndex", value)
	return a
}

func (a *Button) Target(value interface{}) *Button {
	a.Set("target", value)
	return a
}

func (a *Button) TestIdBuilder(value interface{}) *Button {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Button) Testid(value interface{}) *Button {
	a.Set("testid", value)
	return a
}

func (a *Button) Tooltip(value interface{}) *Button {
	a.Set("tooltip", value)
	return a
}

func (a *Button) TooltipPlacement(value interface{}) *Button {
	a.Set("tooltipPlacement", value)
	return a
}

func (a *Button) Type(value interface{}) *Button {
	a.Set("type", value)
	return a
}

func (a *Button) UseMobileUI(value interface{}) *Button {
	a.Set("useMobileUI", value)
	return a
}

func (a *Button) Visible(value interface{}) *Button {
	a.Set("visible", value)
	return a
}

func (a *Button) VisibleOn(value interface{}) *Button {
	a.Set("visibleOn", value)
	return a
}
