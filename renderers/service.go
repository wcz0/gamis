package renderers

type Service struct {
	*BaseRenderer
}

func NewService() *Service {
	a := &Service{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "service")
	return a
}

func (a *Service) Set(name string, value interface{}) *Service {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Service) Api(value interface{}) *Service {
	a.Set("api", value)
	return a
}

func (a *Service) Body(value interface{}) *Service {
	a.Set("body", value)
	return a
}

func (a *Service) ClassName(value interface{}) *Service {
	a.Set("className", value)
	return a
}

func (a *Service) DataProvider(value interface{}) *Service {
	a.Set("dataProvider", value)
	return a
}

func (a *Service) Disabled(value interface{}) *Service {
	a.Set("disabled", value)
	return a
}

func (a *Service) DisabledOn(value interface{}) *Service {
	a.Set("disabledOn", value)
	return a
}

func (a *Service) EditorSetting(value interface{}) *Service {
	a.Set("editorSetting", value)
	return a
}

func (a *Service) FetchOn(value interface{}) *Service {
	a.Set("fetchOn", value)
	return a
}

func (a *Service) Hidden(value interface{}) *Service {
	a.Set("hidden", value)
	return a
}

func (a *Service) HiddenOn(value interface{}) *Service {
	a.Set("hiddenOn", value)
	return a
}

func (a *Service) Id(value interface{}) *Service {
	a.Set("id", value)
	return a
}

func (a *Service) InitFetch(value interface{}) *Service {
	a.Set("initFetch", value)
	return a
}

func (a *Service) InitFetchOn(value interface{}) *Service {
	a.Set("initFetchOn", value)
	return a
}

func (a *Service) InitFetchSchema(value interface{}) *Service {
	a.Set("initFetchSchema", value)
	return a
}

func (a *Service) InitFetchSchemaOn(value interface{}) *Service {
	a.Set("initFetchSchemaOn", value)
	return a
}

func (a *Service) Interval(value interface{}) *Service {
	a.Set("interval", value)
	return a
}

func (a *Service) LoadingConfig(value interface{}) *Service {
	a.Set("loadingConfig", value)
	return a
}

func (a *Service) Messages(value interface{}) *Service {
	a.Set("messages", value)
	return a
}

func (a *Service) Name(value interface{}) *Service {
	a.Set("name", value)
	return a
}

func (a *Service) OnEvent(value interface{}) *Service {
	a.Set("onEvent", value)
	return a
}

func (a *Service) SchemaApi(value interface{}) *Service {
	a.Set("schemaApi", value)
	return a
}

func (a *Service) ShowErrorMsg(value interface{}) *Service {
	a.Set("showErrorMsg", value)
	return a
}

func (a *Service) SilentPolling(value interface{}) *Service {
	a.Set("silentPolling", value)
	return a
}

func (a *Service) Static(value interface{}) *Service {
	a.Set("static", value)
	return a
}

func (a *Service) StaticClassName(value interface{}) *Service {
	a.Set("staticClassName", value)
	return a
}

func (a *Service) StaticInputClassName(value interface{}) *Service {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Service) StaticLabelClassName(value interface{}) *Service {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Service) StaticOn(value interface{}) *Service {
	a.Set("staticOn", value)
	return a
}

func (a *Service) StaticPlaceholder(value interface{}) *Service {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Service) StaticSchema(value interface{}) *Service {
	a.Set("staticSchema", value)
	return a
}

func (a *Service) StopAutoRefreshWhen(value interface{}) *Service {
	a.Set("stopAutoRefreshWhen", value)
	return a
}

func (a *Service) Style(value interface{}) *Service {
	a.Set("style", value)
	return a
}

func (a *Service) TestIdBuilder(value interface{}) *Service {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Service) Testid(value interface{}) *Service {
	a.Set("testid", value)
	return a
}

func (a *Service) Type(value interface{}) *Service {
	a.Set("type", value)
	return a
}

func (a *Service) UseMobileUI(value interface{}) *Service {
	a.Set("useMobileUI", value)
	return a
}

func (a *Service) Visible(value interface{}) *Service {
	a.Set("visible", value)
	return a
}

func (a *Service) VisibleOn(value interface{}) *Service {
	a.Set("visibleOn", value)
	return a
}

func (a *Service) Ws(value interface{}) *Service {
	a.Set("ws", value)
	return a
}
