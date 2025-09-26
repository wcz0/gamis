package renderers

type WizardStep struct {
	*BaseRenderer
}

func NewWizardStep() *WizardStep {
	a := &WizardStep{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *WizardStep) Set(name string, value interface{}) *WizardStep {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *WizardStep) Actions(value interface{}) *WizardStep {
	a.Set("actions", value)
	return a
}

func (a *WizardStep) AffixFooter(value interface{}) *WizardStep {
	a.Set("affixFooter", value)
	return a
}

func (a *WizardStep) Api(value interface{}) *WizardStep {
	a.Set("api", value)
	return a
}

func (a *WizardStep) AsyncApi(value interface{}) *WizardStep {
	a.Set("asyncApi", value)
	return a
}

func (a *WizardStep) AutoFocus(value interface{}) *WizardStep {
	a.Set("autoFocus", value)
	return a
}

func (a *WizardStep) Body(value interface{}) *WizardStep {
	a.Set("body", value)
	return a
}

func (a *WizardStep) CheckInterval(value interface{}) *WizardStep {
	a.Set("checkInterval", value)
	return a
}

func (a *WizardStep) ClassName(value interface{}) *WizardStep {
	a.Set("className", value)
	return a
}

func (a *WizardStep) ClearAfterSubmit(value interface{}) *WizardStep {
	a.Set("clearAfterSubmit", value)
	return a
}

func (a *WizardStep) ClearPersistDataAfterSubmit(value interface{}) *WizardStep {
	a.Set("clearPersistDataAfterSubmit", value)
	return a
}

func (a *WizardStep) ColumnCount(value interface{}) *WizardStep {
	a.Set("columnCount", value)
	return a
}

func (a *WizardStep) Data(value interface{}) *WizardStep {
	a.Set("data", value)
	return a
}

func (a *WizardStep) Debug(value interface{}) *WizardStep {
	a.Set("debug", value)
	return a
}

func (a *WizardStep) DebugConfig(value interface{}) *WizardStep {
	a.Set("debugConfig", value)
	return a
}

func (a *WizardStep) Description(value interface{}) *WizardStep {
	a.Set("description", value)
	return a
}

func (a *WizardStep) Disabled(value interface{}) *WizardStep {
	a.Set("disabled", value)
	return a
}

func (a *WizardStep) DisabledOn(value interface{}) *WizardStep {
	a.Set("disabledOn", value)
	return a
}

func (a *WizardStep) EditorSetting(value interface{}) *WizardStep {
	a.Set("editorSetting", value)
	return a
}

func (a *WizardStep) Feedback(value interface{}) *WizardStep {
	a.Set("feedback", value)
	return a
}

func (a *WizardStep) FieldSet(value interface{}) *WizardStep {
	a.Set("fieldSet", value)
	return a
}

func (a *WizardStep) FinishedField(value interface{}) *WizardStep {
	a.Set("finishedField", value)
	return a
}

func (a *WizardStep) Hidden(value interface{}) *WizardStep {
	a.Set("hidden", value)
	return a
}

func (a *WizardStep) HiddenOn(value interface{}) *WizardStep {
	a.Set("hiddenOn", value)
	return a
}

func (a *WizardStep) Horizontal(value interface{}) *WizardStep {
	a.Set("horizontal", value)
	return a
}

func (a *WizardStep) Icon(value interface{}) *WizardStep {
	a.Set("icon", value)
	return a
}

func (a *WizardStep) Id(value interface{}) *WizardStep {
	a.Set("id", value)
	return a
}

func (a *WizardStep) InitApi(value interface{}) *WizardStep {
	a.Set("initApi", value)
	return a
}

func (a *WizardStep) InitAsyncApi(value interface{}) *WizardStep {
	a.Set("initAsyncApi", value)
	return a
}

func (a *WizardStep) InitCheckInterval(value interface{}) *WizardStep {
	a.Set("initCheckInterval", value)
	return a
}

func (a *WizardStep) InitFetch(value interface{}) *WizardStep {
	a.Set("initFetch", value)
	return a
}

func (a *WizardStep) InitFetchOn(value interface{}) *WizardStep {
	a.Set("initFetchOn", value)
	return a
}

func (a *WizardStep) InitFinishedField(value interface{}) *WizardStep {
	a.Set("initFinishedField", value)
	return a
}

func (a *WizardStep) Interval(value interface{}) *WizardStep {
	a.Set("interval", value)
	return a
}

func (a *WizardStep) Jumpable(value interface{}) *WizardStep {
	a.Set("jumpable", value)
	return a
}

func (a *WizardStep) JumpableOn(value interface{}) *WizardStep {
	a.Set("jumpableOn", value)
	return a
}

func (a *WizardStep) Label(value interface{}) *WizardStep {
	a.Set("label", value)
	return a
}

func (a *WizardStep) LabelAlign(value interface{}) *WizardStep {
	a.Set("labelAlign", value)
	return a
}

func (a *WizardStep) LabelWidth(value interface{}) *WizardStep {
	a.Set("labelWidth", value)
	return a
}

func (a *WizardStep) Messages(value interface{}) *WizardStep {
	a.Set("messages", value)
	return a
}

func (a *WizardStep) Mode(value interface{}) *WizardStep {
	a.Set("mode", value)
	return a
}

func (a *WizardStep) Name(value interface{}) *WizardStep {
	a.Set("name", value)
	return a
}

func (a *WizardStep) OnEvent(value interface{}) *WizardStep {
	a.Set("onEvent", value)
	return a
}

func (a *WizardStep) PanelClassName(value interface{}) *WizardStep {
	a.Set("panelClassName", value)
	return a
}

func (a *WizardStep) PersistData(value interface{}) *WizardStep {
	a.Set("persistData", value)
	return a
}

func (a *WizardStep) PersistDataKeys(value interface{}) *WizardStep {
	a.Set("persistDataKeys", value)
	return a
}

func (a *WizardStep) PreventEnterSubmit(value interface{}) *WizardStep {
	a.Set("preventEnterSubmit", value)
	return a
}

func (a *WizardStep) PrimaryField(value interface{}) *WizardStep {
	a.Set("primaryField", value)
	return a
}

func (a *WizardStep) PromptPageLeave(value interface{}) *WizardStep {
	a.Set("promptPageLeave", value)
	return a
}

func (a *WizardStep) PromptPageLeaveMessage(value interface{}) *WizardStep {
	a.Set("promptPageLeaveMessage", value)
	return a
}

func (a *WizardStep) Redirect(value interface{}) *WizardStep {
	a.Set("redirect", value)
	return a
}

func (a *WizardStep) Reload(value interface{}) *WizardStep {
	a.Set("reload", value)
	return a
}

func (a *WizardStep) ResetAfterSubmit(value interface{}) *WizardStep {
	a.Set("resetAfterSubmit", value)
	return a
}

func (a *WizardStep) Rules(value interface{}) *WizardStep {
	a.Set("rules", value)
	return a
}

func (a *WizardStep) SilentPolling(value interface{}) *WizardStep {
	a.Set("silentPolling", value)
	return a
}

func (a *WizardStep) Static(value interface{}) *WizardStep {
	a.Set("static", value)
	return a
}

func (a *WizardStep) StaticClassName(value interface{}) *WizardStep {
	a.Set("staticClassName", value)
	return a
}

func (a *WizardStep) StaticInputClassName(value interface{}) *WizardStep {
	a.Set("staticInputClassName", value)
	return a
}

func (a *WizardStep) StaticLabelClassName(value interface{}) *WizardStep {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *WizardStep) StaticOn(value interface{}) *WizardStep {
	a.Set("staticOn", value)
	return a
}

func (a *WizardStep) StaticPlaceholder(value interface{}) *WizardStep {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *WizardStep) StaticSchema(value interface{}) *WizardStep {
	a.Set("staticSchema", value)
	return a
}

func (a *WizardStep) StopAutoRefreshWhen(value interface{}) *WizardStep {
	a.Set("stopAutoRefreshWhen", value)
	return a
}

func (a *WizardStep) Style(value interface{}) *WizardStep {
	a.Set("style", value)
	return a
}

func (a *WizardStep) SubTitle(value interface{}) *WizardStep {
	a.Set("subTitle", value)
	return a
}

func (a *WizardStep) SubmitOnChange(value interface{}) *WizardStep {
	a.Set("submitOnChange", value)
	return a
}

func (a *WizardStep) SubmitOnInit(value interface{}) *WizardStep {
	a.Set("submitOnInit", value)
	return a
}

func (a *WizardStep) SubmitText(value interface{}) *WizardStep {
	a.Set("submitText", value)
	return a
}

func (a *WizardStep) Tabs(value interface{}) *WizardStep {
	a.Set("tabs", value)
	return a
}

func (a *WizardStep) Target(value interface{}) *WizardStep {
	a.Set("target", value)
	return a
}

func (a *WizardStep) TestIdBuilder(value interface{}) *WizardStep {
	a.Set("testIdBuilder", value)
	return a
}

func (a *WizardStep) Testid(value interface{}) *WizardStep {
	a.Set("testid", value)
	return a
}

func (a *WizardStep) Title(value interface{}) *WizardStep {
	a.Set("title", value)
	return a
}

func (a *WizardStep) UseMobileUI(value interface{}) *WizardStep {
	a.Set("useMobileUI", value)
	return a
}

func (a *WizardStep) Value(value interface{}) *WizardStep {
	a.Set("value", value)
	return a
}

func (a *WizardStep) Visible(value interface{}) *WizardStep {
	a.Set("visible", value)
	return a
}

func (a *WizardStep) VisibleOn(value interface{}) *WizardStep {
	a.Set("visibleOn", value)
	return a
}

func (a *WizardStep) WrapWithPanel(value interface{}) *WizardStep {
	a.Set("wrapWithPanel", value)
	return a
}
