package renderers

type FeedbackDialog struct {
	*BaseRenderer
}

func NewFeedbackDialog() *FeedbackDialog {
	a := &FeedbackDialog{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *FeedbackDialog) Set(name string, value interface{}) *FeedbackDialog {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *FeedbackDialog) Actions(value interface{}) *FeedbackDialog {
	a.Set("actions", value)
	return a
}

func (a *FeedbackDialog) Body(value interface{}) *FeedbackDialog {
	a.Set("body", value)
	return a
}

func (a *FeedbackDialog) BodyClassName(value interface{}) *FeedbackDialog {
	a.Set("bodyClassName", value)
	return a
}

func (a *FeedbackDialog) ClassName(value interface{}) *FeedbackDialog {
	a.Set("className", value)
	return a
}

func (a *FeedbackDialog) CloseOnEsc(value interface{}) *FeedbackDialog {
	a.Set("closeOnEsc", value)
	return a
}

func (a *FeedbackDialog) CloseOnOutside(value interface{}) *FeedbackDialog {
	a.Set("closeOnOutside", value)
	return a
}

func (a *FeedbackDialog) Confirm(value interface{}) *FeedbackDialog {
	a.Set("confirm", value)
	return a
}

func (a *FeedbackDialog) Data(value interface{}) *FeedbackDialog {
	a.Set("data", value)
	return a
}

func (a *FeedbackDialog) DialogType(value interface{}) *FeedbackDialog {
	a.Set("dialogType", value)
	return a
}

func (a *FeedbackDialog) Disabled(value interface{}) *FeedbackDialog {
	a.Set("disabled", value)
	return a
}

func (a *FeedbackDialog) DisabledOn(value interface{}) *FeedbackDialog {
	a.Set("disabledOn", value)
	return a
}

func (a *FeedbackDialog) Draggable(value interface{}) *FeedbackDialog {
	a.Set("draggable", value)
	return a
}

func (a *FeedbackDialog) EditorSetting(value interface{}) *FeedbackDialog {
	a.Set("editorSetting", value)
	return a
}

func (a *FeedbackDialog) Footer(value interface{}) *FeedbackDialog {
	a.Set("footer", value)
	return a
}

func (a *FeedbackDialog) Header(value interface{}) *FeedbackDialog {
	a.Set("header", value)
	return a
}

func (a *FeedbackDialog) HeaderClassName(value interface{}) *FeedbackDialog {
	a.Set("headerClassName", value)
	return a
}

func (a *FeedbackDialog) Height(value interface{}) *FeedbackDialog {
	a.Set("height", value)
	return a
}

func (a *FeedbackDialog) Hidden(value interface{}) *FeedbackDialog {
	a.Set("hidden", value)
	return a
}

func (a *FeedbackDialog) HiddenOn(value interface{}) *FeedbackDialog {
	a.Set("hiddenOn", value)
	return a
}

func (a *FeedbackDialog) Id(value interface{}) *FeedbackDialog {
	a.Set("id", value)
	return a
}

func (a *FeedbackDialog) InputParams(value interface{}) *FeedbackDialog {
	a.Set("inputParams", value)
	return a
}

func (a *FeedbackDialog) Name(value interface{}) *FeedbackDialog {
	a.Set("name", value)
	return a
}

func (a *FeedbackDialog) OnEvent(value interface{}) *FeedbackDialog {
	a.Set("onEvent", value)
	return a
}

func (a *FeedbackDialog) Overlay(value interface{}) *FeedbackDialog {
	a.Set("overlay", value)
	return a
}

func (a *FeedbackDialog) ShowCloseButton(value interface{}) *FeedbackDialog {
	a.Set("showCloseButton", value)
	return a
}

func (a *FeedbackDialog) ShowErrorMsg(value interface{}) *FeedbackDialog {
	a.Set("showErrorMsg", value)
	return a
}

func (a *FeedbackDialog) ShowLoading(value interface{}) *FeedbackDialog {
	a.Set("showLoading", value)
	return a
}

func (a *FeedbackDialog) Size(value interface{}) *FeedbackDialog {
	a.Set("size", value)
	return a
}

func (a *FeedbackDialog) SkipRestOnCancel(value interface{}) *FeedbackDialog {
	a.Set("skipRestOnCancel", value)
	return a
}

func (a *FeedbackDialog) SkipRestOnConfirm(value interface{}) *FeedbackDialog {
	a.Set("skipRestOnConfirm", value)
	return a
}

func (a *FeedbackDialog) Static(value interface{}) *FeedbackDialog {
	a.Set("static", value)
	return a
}

func (a *FeedbackDialog) StaticClassName(value interface{}) *FeedbackDialog {
	a.Set("staticClassName", value)
	return a
}

func (a *FeedbackDialog) StaticInputClassName(value interface{}) *FeedbackDialog {
	a.Set("staticInputClassName", value)
	return a
}

func (a *FeedbackDialog) StaticLabelClassName(value interface{}) *FeedbackDialog {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *FeedbackDialog) StaticOn(value interface{}) *FeedbackDialog {
	a.Set("staticOn", value)
	return a
}

func (a *FeedbackDialog) StaticPlaceholder(value interface{}) *FeedbackDialog {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *FeedbackDialog) StaticSchema(value interface{}) *FeedbackDialog {
	a.Set("staticSchema", value)
	return a
}

func (a *FeedbackDialog) Style(value interface{}) *FeedbackDialog {
	a.Set("style", value)
	return a
}

func (a *FeedbackDialog) TestIdBuilder(value interface{}) *FeedbackDialog {
	a.Set("testIdBuilder", value)
	return a
}

func (a *FeedbackDialog) Testid(value interface{}) *FeedbackDialog {
	a.Set("testid", value)
	return a
}

func (a *FeedbackDialog) Title(value interface{}) *FeedbackDialog {
	a.Set("title", value)
	return a
}

func (a *FeedbackDialog) UseMobileUI(value interface{}) *FeedbackDialog {
	a.Set("useMobileUI", value)
	return a
}

func (a *FeedbackDialog) Visible(value interface{}) *FeedbackDialog {
	a.Set("visible", value)
	return a
}

func (a *FeedbackDialog) VisibleOn(value interface{}) *FeedbackDialog {
	a.Set("visibleOn", value)
	return a
}

func (a *FeedbackDialog) Width(value interface{}) *FeedbackDialog {
	a.Set("width", value)
	return a
}
