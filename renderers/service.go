package renderers


/**
 * Service 服务类控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/service

 * @author wcz0
 * @version 6.2.2
 */
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
/**
 * stopAutoRefreshWhen
 */
func (a *Service) StopAutoRefreshWhen(value interface{}) *Service {
    a.Set("stopAutoRefreshWhen", value)
    return a
}

/**
 * disabled
 */
func (a *Service) Disabled(value interface{}) *Service {
    a.Set("disabled", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Service) StaticPlaceholder(value interface{}) *Service {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * initFetchSchema
 */
func (a *Service) InitFetchSchema(value interface{}) *Service {
    a.Set("initFetchSchema", value)
    return a
}

/**
 * interval
 */
func (a *Service) Interval(value interface{}) *Service {
    a.Set("interval", value)
    return a
}

/**
 * static
 */
func (a *Service) Static(value interface{}) *Service {
    a.Set("static", value)
    return a
}

/**
 * style
 */
func (a *Service) Style(value interface{}) *Service {
    a.Set("style", value)
    return a
}

/**
 * ws
 */
func (a *Service) Ws(value interface{}) *Service {
    a.Set("ws", value)
    return a
}

/**
 * dataProvider
 */
func (a *Service) DataProvider(value interface{}) *Service {
    a.Set("dataProvider", value)
    return a
}

/**
 * initFetchOn
 */
func (a *Service) InitFetchOn(value interface{}) *Service {
    a.Set("initFetchOn", value)
    return a
}

/**
 * showErrorMsg
 */
func (a *Service) ShowErrorMsg(value interface{}) *Service {
    a.Set("showErrorMsg", value)
    return a
}

/**
 * hidden
 */
func (a *Service) Hidden(value interface{}) *Service {
    a.Set("hidden", value)
    return a
}

/**
 * id
 */
func (a *Service) Id(value interface{}) *Service {
    a.Set("id", value)
    return a
}

/**
 * staticOn
 */
func (a *Service) StaticOn(value interface{}) *Service {
    a.Set("staticOn", value)
    return a
}

/**
 * editorSetting
 */
func (a *Service) EditorSetting(value interface{}) *Service {
    a.Set("editorSetting", value)
    return a
}

/**
 * initFetchSchemaOn
 */
func (a *Service) InitFetchSchemaOn(value interface{}) *Service {
    a.Set("initFetchSchemaOn", value)
    return a
}

/**
 * onEvent
 */
func (a *Service) OnEvent(value interface{}) *Service {
    a.Set("onEvent", value)
    return a
}

/**
 * fetchOn
 */
func (a *Service) FetchOn(value interface{}) *Service {
    a.Set("fetchOn", value)
    return a
}

/**
 * schemaApi
 */
func (a *Service) SchemaApi(value interface{}) *Service {
    a.Set("schemaApi", value)
    return a
}

/**
 * messages
 */
func (a *Service) Messages(value interface{}) *Service {
    a.Set("messages", value)
    return a
}

/**
 * staticSchema
 */
func (a *Service) StaticSchema(value interface{}) *Service {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *Service) Type(value interface{}) *Service {
    a.Set("type", value)
    return a
}

/**
 * name
 */
func (a *Service) Name(value interface{}) *Service {
    a.Set("name", value)
    return a
}

/**
 * className
 */
func (a *Service) ClassName(value interface{}) *Service {
    a.Set("className", value)
    return a
}

/**
 * visible
 */
func (a *Service) Visible(value interface{}) *Service {
    a.Set("visible", value)
    return a
}

/**
 * visibleOn
 */
func (a *Service) VisibleOn(value interface{}) *Service {
    a.Set("visibleOn", value)
    return a
}

/**
 * api
 */
func (a *Service) Api(value interface{}) *Service {
    a.Set("api", value)
    return a
}

/**
 * loadingConfig
 */
func (a *Service) LoadingConfig(value interface{}) *Service {
    a.Set("loadingConfig", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Service) UseMobileUI(value interface{}) *Service {
    a.Set("useMobileUI", value)
    return a
}

/**
 * initFetch
 */
func (a *Service) InitFetch(value interface{}) *Service {
    a.Set("initFetch", value)
    return a
}

/**
 * silentPolling
 */
func (a *Service) SilentPolling(value interface{}) *Service {
    a.Set("silentPolling", value)
    return a
}

/**
 * disabledOn
 */
func (a *Service) DisabledOn(value interface{}) *Service {
    a.Set("disabledOn", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Service) HiddenOn(value interface{}) *Service {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *Service) StaticClassName(value interface{}) *Service {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Service) StaticLabelClassName(value interface{}) *Service {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Service) StaticInputClassName(value interface{}) *Service {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * testid
 */
func (a *Service) Testid(value interface{}) *Service {
    a.Set("testid", value)
    return a
}

/**
 * body
 */
func (a *Service) Body(value interface{}) *Service {
    a.Set("body", value)
    return a
}
