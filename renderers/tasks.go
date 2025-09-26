package renderers

type Tasks struct {
	*BaseRenderer
}

func NewTasks() *Tasks {
	a := &Tasks{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "tasks")
	return a
}

func (a *Tasks) Set(name string, value interface{}) *Tasks {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Tasks) BtnClassName(value interface{}) *Tasks {
	a.Set("btnClassName", value)
	return a
}

func (a *Tasks) BtnText(value interface{}) *Tasks {
	a.Set("btnText", value)
	return a
}

func (a *Tasks) CanRetryStatusCode(value interface{}) *Tasks {
	a.Set("canRetryStatusCode", value)
	return a
}

func (a *Tasks) CheckApi(value interface{}) *Tasks {
	a.Set("checkApi", value)
	return a
}

func (a *Tasks) ClassName(value interface{}) *Tasks {
	a.Set("className", value)
	return a
}

func (a *Tasks) Disabled(value interface{}) *Tasks {
	a.Set("disabled", value)
	return a
}

func (a *Tasks) DisabledOn(value interface{}) *Tasks {
	a.Set("disabledOn", value)
	return a
}

func (a *Tasks) EditorSetting(value interface{}) *Tasks {
	a.Set("editorSetting", value)
	return a
}

func (a *Tasks) ErrorStatusCode(value interface{}) *Tasks {
	a.Set("errorStatusCode", value)
	return a
}

func (a *Tasks) FinishStatusCode(value interface{}) *Tasks {
	a.Set("finishStatusCode", value)
	return a
}

func (a *Tasks) Hidden(value interface{}) *Tasks {
	a.Set("hidden", value)
	return a
}

func (a *Tasks) HiddenOn(value interface{}) *Tasks {
	a.Set("hiddenOn", value)
	return a
}

func (a *Tasks) Id(value interface{}) *Tasks {
	a.Set("id", value)
	return a
}

func (a *Tasks) InitialStatusCode(value interface{}) *Tasks {
	a.Set("initialStatusCode", value)
	return a
}

func (a *Tasks) Interval(value interface{}) *Tasks {
	a.Set("interval", value)
	return a
}

func (a *Tasks) Items(value interface{}) *Tasks {
	a.Set("items", value)
	return a
}

func (a *Tasks) LoadingConfig(value interface{}) *Tasks {
	a.Set("loadingConfig", value)
	return a
}

func (a *Tasks) LoadingStatusCode(value interface{}) *Tasks {
	a.Set("loadingStatusCode", value)
	return a
}

func (a *Tasks) Name(value interface{}) *Tasks {
	a.Set("name", value)
	return a
}

func (a *Tasks) OnEvent(value interface{}) *Tasks {
	a.Set("onEvent", value)
	return a
}

func (a *Tasks) OperationLabel(value interface{}) *Tasks {
	a.Set("operationLabel", value)
	return a
}

func (a *Tasks) ReSubmitApi(value interface{}) *Tasks {
	a.Set("reSubmitApi", value)
	return a
}

func (a *Tasks) ReadyStatusCode(value interface{}) *Tasks {
	a.Set("readyStatusCode", value)
	return a
}

func (a *Tasks) RemarkLabel(value interface{}) *Tasks {
	a.Set("remarkLabel", value)
	return a
}

func (a *Tasks) RetryBtnClassName(value interface{}) *Tasks {
	a.Set("retryBtnClassName", value)
	return a
}

func (a *Tasks) RetryBtnText(value interface{}) *Tasks {
	a.Set("retryBtnText", value)
	return a
}

func (a *Tasks) Static(value interface{}) *Tasks {
	a.Set("static", value)
	return a
}

func (a *Tasks) StaticClassName(value interface{}) *Tasks {
	a.Set("staticClassName", value)
	return a
}

func (a *Tasks) StaticInputClassName(value interface{}) *Tasks {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Tasks) StaticLabelClassName(value interface{}) *Tasks {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Tasks) StaticOn(value interface{}) *Tasks {
	a.Set("staticOn", value)
	return a
}

func (a *Tasks) StaticPlaceholder(value interface{}) *Tasks {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Tasks) StaticSchema(value interface{}) *Tasks {
	a.Set("staticSchema", value)
	return a
}

func (a *Tasks) StatusLabel(value interface{}) *Tasks {
	a.Set("statusLabel", value)
	return a
}

func (a *Tasks) StatusLabelMap(value interface{}) *Tasks {
	a.Set("statusLabelMap", value)
	return a
}

func (a *Tasks) StatusTextMap(value interface{}) *Tasks {
	a.Set("statusTextMap", value)
	return a
}

func (a *Tasks) Style(value interface{}) *Tasks {
	a.Set("style", value)
	return a
}

func (a *Tasks) SubmitApi(value interface{}) *Tasks {
	a.Set("submitApi", value)
	return a
}

func (a *Tasks) TableClassName(value interface{}) *Tasks {
	a.Set("tableClassName", value)
	return a
}

func (a *Tasks) TaskNameLabel(value interface{}) *Tasks {
	a.Set("taskNameLabel", value)
	return a
}

func (a *Tasks) TestIdBuilder(value interface{}) *Tasks {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Tasks) Testid(value interface{}) *Tasks {
	a.Set("testid", value)
	return a
}

func (a *Tasks) Type(value interface{}) *Tasks {
	a.Set("type", value)
	return a
}

func (a *Tasks) UseMobileUI(value interface{}) *Tasks {
	a.Set("useMobileUI", value)
	return a
}

func (a *Tasks) Visible(value interface{}) *Tasks {
	a.Set("visible", value)
	return a
}

func (a *Tasks) VisibleOn(value interface{}) *Tasks {
	a.Set("visibleOn", value)
	return a
}
