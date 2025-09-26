package renderers

type Wizard struct {
	*BaseRenderer
}

func NewWizard() *Wizard {
	a := &Wizard{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "wizard")
	return a
}

func (a *Wizard) Set(name string, value interface{}) *Wizard {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Wizard) ActionClassName(value interface{}) *Wizard {
	a.Set("actionClassName", value)
	return a
}

func (a *Wizard) ActionFinishLabel(value interface{}) *Wizard {
	a.Set("actionFinishLabel", value)
	return a
}

func (a *Wizard) ActionNextLabel(value interface{}) *Wizard {
	a.Set("actionNextLabel", value)
	return a
}

func (a *Wizard) ActionNextSaveLabel(value interface{}) *Wizard {
	a.Set("actionNextSaveLabel", value)
	return a
}

func (a *Wizard) ActionPrevLabel(value interface{}) *Wizard {
	a.Set("actionPrevLabel", value)
	return a
}

func (a *Wizard) AffixFooter(value interface{}) *Wizard {
	a.Set("affixFooter", value)
	return a
}

func (a *Wizard) Api(value interface{}) *Wizard {
	a.Set("api", value)
	return a
}

func (a *Wizard) BodyClassName(value interface{}) *Wizard {
	a.Set("bodyClassName", value)
	return a
}

func (a *Wizard) BulkSubmit(value interface{}) *Wizard {
	a.Set("bulkSubmit", value)
	return a
}

func (a *Wizard) ClassName(value interface{}) *Wizard {
	a.Set("className", value)
	return a
}

func (a *Wizard) Disabled(value interface{}) *Wizard {
	a.Set("disabled", value)
	return a
}

func (a *Wizard) DisabledOn(value interface{}) *Wizard {
	a.Set("disabledOn", value)
	return a
}

func (a *Wizard) EditorSetting(value interface{}) *Wizard {
	a.Set("editorSetting", value)
	return a
}

func (a *Wizard) FooterClassName(value interface{}) *Wizard {
	a.Set("footerClassName", value)
	return a
}

func (a *Wizard) Hidden(value interface{}) *Wizard {
	a.Set("hidden", value)
	return a
}

func (a *Wizard) HiddenOn(value interface{}) *Wizard {
	a.Set("hiddenOn", value)
	return a
}

func (a *Wizard) Id(value interface{}) *Wizard {
	a.Set("id", value)
	return a
}

func (a *Wizard) InitApi(value interface{}) *Wizard {
	a.Set("initApi", value)
	return a
}

func (a *Wizard) LoadingConfig(value interface{}) *Wizard {
	a.Set("loadingConfig", value)
	return a
}

func (a *Wizard) Mode(value interface{}) *Wizard {
	a.Set("mode", value)
	return a
}

func (a *Wizard) Name(value interface{}) *Wizard {
	a.Set("name", value)
	return a
}

func (a *Wizard) OnEvent(value interface{}) *Wizard {
	a.Set("onEvent", value)
	return a
}

func (a *Wizard) ReadOnly(value interface{}) *Wizard {
	a.Set("readOnly", value)
	return a
}

func (a *Wizard) Redirect(value interface{}) *Wizard {
	a.Set("redirect", value)
	return a
}

func (a *Wizard) Reload(value interface{}) *Wizard {
	a.Set("reload", value)
	return a
}

func (a *Wizard) StartStep(value interface{}) *Wizard {
	a.Set("startStep", value)
	return a
}

func (a *Wizard) Static(value interface{}) *Wizard {
	a.Set("static", value)
	return a
}

func (a *Wizard) StaticClassName(value interface{}) *Wizard {
	a.Set("staticClassName", value)
	return a
}

func (a *Wizard) StaticInputClassName(value interface{}) *Wizard {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Wizard) StaticLabelClassName(value interface{}) *Wizard {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Wizard) StaticOn(value interface{}) *Wizard {
	a.Set("staticOn", value)
	return a
}

func (a *Wizard) StaticPlaceholder(value interface{}) *Wizard {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Wizard) StaticSchema(value interface{}) *Wizard {
	a.Set("staticSchema", value)
	return a
}

func (a *Wizard) StepClassName(value interface{}) *Wizard {
	a.Set("stepClassName", value)
	return a
}

func (a *Wizard) Steps(value interface{}) *Wizard {
	a.Set("steps", value)
	return a
}

func (a *Wizard) StepsClassName(value interface{}) *Wizard {
	a.Set("stepsClassName", value)
	return a
}

func (a *Wizard) Style(value interface{}) *Wizard {
	a.Set("style", value)
	return a
}

func (a *Wizard) Target(value interface{}) *Wizard {
	a.Set("target", value)
	return a
}

func (a *Wizard) TestIdBuilder(value interface{}) *Wizard {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Wizard) Testid(value interface{}) *Wizard {
	a.Set("testid", value)
	return a
}

func (a *Wizard) Type(value interface{}) *Wizard {
	a.Set("type", value)
	return a
}

func (a *Wizard) UseMobileUI(value interface{}) *Wizard {
	a.Set("useMobileUI", value)
	return a
}

func (a *Wizard) Visible(value interface{}) *Wizard {
	a.Set("visible", value)
	return a
}

func (a *Wizard) VisibleOn(value interface{}) *Wizard {
	a.Set("visibleOn", value)
	return a
}

func (a *Wizard) WrapWithPanel(value interface{}) *Wizard {
	a.Set("wrapWithPanel", value)
	return a
}
