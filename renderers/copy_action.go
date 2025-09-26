package renderers

type CopyAction struct {
	*BaseRenderer
}

func NewCopyAction() *CopyAction {
	a := &CopyAction{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "button")
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

func (a *CopyAction) ActionType(value interface{}) *CopyAction {
	a.Set("actionType", value)
	return a
}

func (a *CopyAction) ActiveClassName(value interface{}) *CopyAction {
	a.Set("activeClassName", value)
	return a
}

func (a *CopyAction) ActiveLevel(value interface{}) *CopyAction {
	a.Set("activeLevel", value)
	return a
}

func (a *CopyAction) Badge(value interface{}) *CopyAction {
	a.Set("badge", value)
	return a
}

func (a *CopyAction) Block(value interface{}) *CopyAction {
	a.Set("block", value)
	return a
}

func (a *CopyAction) Body(value interface{}) *CopyAction {
	a.Set("body", value)
	return a
}

func (a *CopyAction) ClassName(value interface{}) *CopyAction {
	a.Set("className", value)
	return a
}

func (a *CopyAction) Close(value interface{}) *CopyAction {
	a.Set("close", value)
	return a
}

func (a *CopyAction) ConfirmText(value interface{}) *CopyAction {
	a.Set("confirmText", value)
	return a
}

func (a *CopyAction) ConfirmTitle(value interface{}) *CopyAction {
	a.Set("confirmTitle", value)
	return a
}

func (a *CopyAction) Content(value interface{}) *CopyAction {
	a.Set("content", value)
	return a
}

func (a *CopyAction) Copy(value interface{}) *CopyAction {
	a.Set("copy", value)
	return a
}

func (a *CopyAction) CountDown(value interface{}) *CopyAction {
	a.Set("countDown", value)
	return a
}

func (a *CopyAction) CountDownTpl(value interface{}) *CopyAction {
	a.Set("countDownTpl", value)
	return a
}

func (a *CopyAction) Disabled(value interface{}) *CopyAction {
	a.Set("disabled", value)
	return a
}

func (a *CopyAction) DisabledOn(value interface{}) *CopyAction {
	a.Set("disabledOn", value)
	return a
}

func (a *CopyAction) DisabledOnAction(value interface{}) *CopyAction {
	a.Set("disabledOnAction", value)
	return a
}

func (a *CopyAction) DisabledTip(value interface{}) *CopyAction {
	a.Set("disabledTip", value)
	return a
}

func (a *CopyAction) EditorSetting(value interface{}) *CopyAction {
	a.Set("editorSetting", value)
	return a
}

func (a *CopyAction) Hidden(value interface{}) *CopyAction {
	a.Set("hidden", value)
	return a
}

func (a *CopyAction) HiddenOn(value interface{}) *CopyAction {
	a.Set("hiddenOn", value)
	return a
}

func (a *CopyAction) HotKey(value interface{}) *CopyAction {
	a.Set("hotKey", value)
	return a
}

func (a *CopyAction) Icon(value interface{}) *CopyAction {
	a.Set("icon", value)
	return a
}

func (a *CopyAction) IconClassName(value interface{}) *CopyAction {
	a.Set("iconClassName", value)
	return a
}

func (a *CopyAction) Id(value interface{}) *CopyAction {
	a.Set("id", value)
	return a
}

func (a *CopyAction) Label(value interface{}) *CopyAction {
	a.Set("label", value)
	return a
}

func (a *CopyAction) Level(value interface{}) *CopyAction {
	a.Set("level", value)
	return a
}

func (a *CopyAction) LoadingClassName(value interface{}) *CopyAction {
	a.Set("loadingClassName", value)
	return a
}

func (a *CopyAction) LoadingOn(value interface{}) *CopyAction {
	a.Set("loadingOn", value)
	return a
}

func (a *CopyAction) MergeData(value interface{}) *CopyAction {
	a.Set("mergeData", value)
	return a
}

func (a *CopyAction) OnClick(value interface{}) *CopyAction {
	a.Set("onClick", value)
	return a
}

func (a *CopyAction) OnEvent(value interface{}) *CopyAction {
	a.Set("onEvent", value)
	return a
}

func (a *CopyAction) Primary(value interface{}) *CopyAction {
	a.Set("primary", value)
	return a
}

func (a *CopyAction) RequireSelected(value interface{}) *CopyAction {
	a.Set("requireSelected", value)
	return a
}

func (a *CopyAction) Required(value interface{}) *CopyAction {
	a.Set("required", value)
	return a
}

func (a *CopyAction) RightIcon(value interface{}) *CopyAction {
	a.Set("rightIcon", value)
	return a
}

func (a *CopyAction) RightIconClassName(value interface{}) *CopyAction {
	a.Set("rightIconClassName", value)
	return a
}

func (a *CopyAction) Size(value interface{}) *CopyAction {
	a.Set("size", value)
	return a
}

func (a *CopyAction) Static(value interface{}) *CopyAction {
	a.Set("static", value)
	return a
}

func (a *CopyAction) StaticClassName(value interface{}) *CopyAction {
	a.Set("staticClassName", value)
	return a
}

func (a *CopyAction) StaticInputClassName(value interface{}) *CopyAction {
	a.Set("staticInputClassName", value)
	return a
}

func (a *CopyAction) StaticLabelClassName(value interface{}) *CopyAction {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *CopyAction) StaticOn(value interface{}) *CopyAction {
	a.Set("staticOn", value)
	return a
}

func (a *CopyAction) StaticPlaceholder(value interface{}) *CopyAction {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *CopyAction) StaticSchema(value interface{}) *CopyAction {
	a.Set("staticSchema", value)
	return a
}

func (a *CopyAction) Style(value interface{}) *CopyAction {
	a.Set("style", value)
	return a
}

func (a *CopyAction) TabIndex(value interface{}) *CopyAction {
	a.Set("tabIndex", value)
	return a
}

func (a *CopyAction) Target(value interface{}) *CopyAction {
	a.Set("target", value)
	return a
}

func (a *CopyAction) TestIdBuilder(value interface{}) *CopyAction {
	a.Set("testIdBuilder", value)
	return a
}

func (a *CopyAction) Testid(value interface{}) *CopyAction {
	a.Set("testid", value)
	return a
}

func (a *CopyAction) Tooltip(value interface{}) *CopyAction {
	a.Set("tooltip", value)
	return a
}

func (a *CopyAction) TooltipPlacement(value interface{}) *CopyAction {
	a.Set("tooltipPlacement", value)
	return a
}

func (a *CopyAction) Type(value interface{}) *CopyAction {
	a.Set("type", value)
	return a
}

func (a *CopyAction) UseMobileUI(value interface{}) *CopyAction {
	a.Set("useMobileUI", value)
	return a
}

func (a *CopyAction) Visible(value interface{}) *CopyAction {
	a.Set("visible", value)
	return a
}

func (a *CopyAction) VisibleOn(value interface{}) *CopyAction {
	a.Set("visibleOn", value)
	return a
}
