package renderers

type BaseForm struct {
	*BaseRenderer
}

func NewBaseForm() *BaseForm {
	a := &BaseForm{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *BaseForm) Set(name string, value interface{}) *BaseForm {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *BaseForm) Actions(value interface{}) *BaseForm {
	a.Set("actions", value)
	return a
}

func (a *BaseForm) AffixFooter(value interface{}) *BaseForm {
	a.Set("affixFooter", value)
	return a
}

func (a *BaseForm) Api(value interface{}) *BaseForm {
	a.Set("api", value)
	return a
}

func (a *BaseForm) AsyncApi(value interface{}) *BaseForm {
	a.Set("asyncApi", value)
	return a
}

func (a *BaseForm) AutoFocus(value interface{}) *BaseForm {
	a.Set("autoFocus", value)
	return a
}

func (a *BaseForm) Body(value interface{}) *BaseForm {
	a.Set("body", value)
	return a
}

func (a *BaseForm) CheckInterval(value interface{}) *BaseForm {
	a.Set("checkInterval", value)
	return a
}

func (a *BaseForm) ClassName(value interface{}) *BaseForm {
	a.Set("className", value)
	return a
}

func (a *BaseForm) ClearAfterSubmit(value interface{}) *BaseForm {
	a.Set("clearAfterSubmit", value)
	return a
}

func (a *BaseForm) ClearPersistDataAfterSubmit(value interface{}) *BaseForm {
	a.Set("clearPersistDataAfterSubmit", value)
	return a
}

func (a *BaseForm) ColumnCount(value interface{}) *BaseForm {
	a.Set("columnCount", value)
	return a
}

func (a *BaseForm) Data(value interface{}) *BaseForm {
	a.Set("data", value)
	return a
}

func (a *BaseForm) Debug(value interface{}) *BaseForm {
	a.Set("debug", value)
	return a
}

func (a *BaseForm) DebugConfig(value interface{}) *BaseForm {
	a.Set("debugConfig", value)
	return a
}

func (a *BaseForm) Disabled(value interface{}) *BaseForm {
	a.Set("disabled", value)
	return a
}

func (a *BaseForm) DisabledOn(value interface{}) *BaseForm {
	a.Set("disabledOn", value)
	return a
}

func (a *BaseForm) EditorSetting(value interface{}) *BaseForm {
	a.Set("editorSetting", value)
	return a
}

func (a *BaseForm) Feedback(value interface{}) *BaseForm {
	a.Set("feedback", value)
	return a
}

func (a *BaseForm) FieldSet(value interface{}) *BaseForm {
	a.Set("fieldSet", value)
	return a
}

func (a *BaseForm) FinishedField(value interface{}) *BaseForm {
	a.Set("finishedField", value)
	return a
}

func (a *BaseForm) Hidden(value interface{}) *BaseForm {
	a.Set("hidden", value)
	return a
}

func (a *BaseForm) HiddenOn(value interface{}) *BaseForm {
	a.Set("hiddenOn", value)
	return a
}

func (a *BaseForm) Horizontal(value interface{}) *BaseForm {
	a.Set("horizontal", value)
	return a
}

func (a *BaseForm) Id(value interface{}) *BaseForm {
	a.Set("id", value)
	return a
}

func (a *BaseForm) InitApi(value interface{}) *BaseForm {
	a.Set("initApi", value)
	return a
}

func (a *BaseForm) InitAsyncApi(value interface{}) *BaseForm {
	a.Set("initAsyncApi", value)
	return a
}

func (a *BaseForm) InitCheckInterval(value interface{}) *BaseForm {
	a.Set("initCheckInterval", value)
	return a
}

func (a *BaseForm) InitFetch(value interface{}) *BaseForm {
	a.Set("initFetch", value)
	return a
}

func (a *BaseForm) InitFetchOn(value interface{}) *BaseForm {
	a.Set("initFetchOn", value)
	return a
}

func (a *BaseForm) InitFinishedField(value interface{}) *BaseForm {
	a.Set("initFinishedField", value)
	return a
}

func (a *BaseForm) Interval(value interface{}) *BaseForm {
	a.Set("interval", value)
	return a
}

func (a *BaseForm) LabelAlign(value interface{}) *BaseForm {
	a.Set("labelAlign", value)
	return a
}

func (a *BaseForm) LabelWidth(value interface{}) *BaseForm {
	a.Set("labelWidth", value)
	return a
}

func (a *BaseForm) Messages(value interface{}) *BaseForm {
	a.Set("messages", value)
	return a
}

func (a *BaseForm) Mode(value interface{}) *BaseForm {
	a.Set("mode", value)
	return a
}

func (a *BaseForm) Name(value interface{}) *BaseForm {
	a.Set("name", value)
	return a
}

func (a *BaseForm) OnEvent(value interface{}) *BaseForm {
	a.Set("onEvent", value)
	return a
}

func (a *BaseForm) PanelClassName(value interface{}) *BaseForm {
	a.Set("panelClassName", value)
	return a
}

func (a *BaseForm) PersistData(value interface{}) *BaseForm {
	a.Set("persistData", value)
	return a
}

func (a *BaseForm) PersistDataKeys(value interface{}) *BaseForm {
	a.Set("persistDataKeys", value)
	return a
}

func (a *BaseForm) PreventEnterSubmit(value interface{}) *BaseForm {
	a.Set("preventEnterSubmit", value)
	return a
}

func (a *BaseForm) PrimaryField(value interface{}) *BaseForm {
	a.Set("primaryField", value)
	return a
}

func (a *BaseForm) PromptPageLeave(value interface{}) *BaseForm {
	a.Set("promptPageLeave", value)
	return a
}

func (a *BaseForm) PromptPageLeaveMessage(value interface{}) *BaseForm {
	a.Set("promptPageLeaveMessage", value)
	return a
}

func (a *BaseForm) Redirect(value interface{}) *BaseForm {
	a.Set("redirect", value)
	return a
}

func (a *BaseForm) Reload(value interface{}) *BaseForm {
	a.Set("reload", value)
	return a
}

func (a *BaseForm) ResetAfterSubmit(value interface{}) *BaseForm {
	a.Set("resetAfterSubmit", value)
	return a
}

func (a *BaseForm) Rules(value interface{}) *BaseForm {
	a.Set("rules", value)
	return a
}

func (a *BaseForm) SilentPolling(value interface{}) *BaseForm {
	a.Set("silentPolling", value)
	return a
}

func (a *BaseForm) Static(value interface{}) *BaseForm {
	a.Set("static", value)
	return a
}

func (a *BaseForm) StaticClassName(value interface{}) *BaseForm {
	a.Set("staticClassName", value)
	return a
}

func (a *BaseForm) StaticInputClassName(value interface{}) *BaseForm {
	a.Set("staticInputClassName", value)
	return a
}

func (a *BaseForm) StaticLabelClassName(value interface{}) *BaseForm {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *BaseForm) StaticOn(value interface{}) *BaseForm {
	a.Set("staticOn", value)
	return a
}

func (a *BaseForm) StaticPlaceholder(value interface{}) *BaseForm {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *BaseForm) StaticSchema(value interface{}) *BaseForm {
	a.Set("staticSchema", value)
	return a
}

func (a *BaseForm) StopAutoRefreshWhen(value interface{}) *BaseForm {
	a.Set("stopAutoRefreshWhen", value)
	return a
}

func (a *BaseForm) Style(value interface{}) *BaseForm {
	a.Set("style", value)
	return a
}

func (a *BaseForm) SubmitOnChange(value interface{}) *BaseForm {
	a.Set("submitOnChange", value)
	return a
}

func (a *BaseForm) SubmitOnInit(value interface{}) *BaseForm {
	a.Set("submitOnInit", value)
	return a
}

func (a *BaseForm) SubmitText(value interface{}) *BaseForm {
	a.Set("submitText", value)
	return a
}

func (a *BaseForm) Tabs(value interface{}) *BaseForm {
	a.Set("tabs", value)
	return a
}

func (a *BaseForm) Target(value interface{}) *BaseForm {
	a.Set("target", value)
	return a
}

func (a *BaseForm) Title(value interface{}) *BaseForm {
	a.Set("title", value)
	return a
}

func (a *BaseForm) UseMobileUI(value interface{}) *BaseForm {
	a.Set("useMobileUI", value)
	return a
}

func (a *BaseForm) Visible(value interface{}) *BaseForm {
	a.Set("visible", value)
	return a
}

func (a *BaseForm) VisibleOn(value interface{}) *BaseForm {
	a.Set("visibleOn", value)
	return a
}

func (a *BaseForm) WrapWithPanel(value interface{}) *BaseForm {
	a.Set("wrapWithPanel", value)
	return a
}
