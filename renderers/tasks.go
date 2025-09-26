package renderers


/**
 * Tasks 渲染器，格式说明 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/tasks

 * @author wcz0
 * @version 6.2.2
 */
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
/**
 * style
 */
func (a *Tasks) Style(value interface{}) *Tasks {
    a.Set("style", value)
    return a
}

/**
 */
func (a *Tasks) Type(value interface{}) *Tasks {
    a.Set("type", value)
    return a
}

/**
 * remarkLabel
 */
func (a *Tasks) RemarkLabel(value interface{}) *Tasks {
    a.Set("remarkLabel", value)
    return a
}

/**
 * readyStatusCode
 */
func (a *Tasks) ReadyStatusCode(value interface{}) *Tasks {
    a.Set("readyStatusCode", value)
    return a
}

/**
 * canRetryStatusCode
 */
func (a *Tasks) CanRetryStatusCode(value interface{}) *Tasks {
    a.Set("canRetryStatusCode", value)
    return a
}

/**
 * finishStatusCode
 */
func (a *Tasks) FinishStatusCode(value interface{}) *Tasks {
    a.Set("finishStatusCode", value)
    return a
}

/**
 * errorStatusCode
 */
func (a *Tasks) ErrorStatusCode(value interface{}) *Tasks {
    a.Set("errorStatusCode", value)
    return a
}

/**
 * onEvent
 */
func (a *Tasks) OnEvent(value interface{}) *Tasks {
    a.Set("onEvent", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Tasks) StaticPlaceholder(value interface{}) *Tasks {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * testid
 */
func (a *Tasks) Testid(value interface{}) *Tasks {
    a.Set("testid", value)
    return a
}

/**
 * name
 */
func (a *Tasks) Name(value interface{}) *Tasks {
    a.Set("name", value)
    return a
}

/**
 * taskNameLabel
 */
func (a *Tasks) TaskNameLabel(value interface{}) *Tasks {
    a.Set("taskNameLabel", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Tasks) HiddenOn(value interface{}) *Tasks {
    a.Set("hiddenOn", value)
    return a
}

/**
 * id
 */
func (a *Tasks) Id(value interface{}) *Tasks {
    a.Set("id", value)
    return a
}

/**
 * static
 */
func (a *Tasks) Static(value interface{}) *Tasks {
    a.Set("static", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Tasks) StaticInputClassName(value interface{}) *Tasks {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * btnClassName
 */
func (a *Tasks) BtnClassName(value interface{}) *Tasks {
    a.Set("btnClassName", value)
    return a
}

/**
 * btnText
 */
func (a *Tasks) BtnText(value interface{}) *Tasks {
    a.Set("btnText", value)
    return a
}

/**
 * submitApi
 */
func (a *Tasks) SubmitApi(value interface{}) *Tasks {
    a.Set("submitApi", value)
    return a
}

/**
 * loadingConfig
 */
func (a *Tasks) LoadingConfig(value interface{}) *Tasks {
    a.Set("loadingConfig", value)
    return a
}

/**
 * hidden
 */
func (a *Tasks) Hidden(value interface{}) *Tasks {
    a.Set("hidden", value)
    return a
}

/**
 * items
 */
func (a *Tasks) Items(value interface{}) *Tasks {
    a.Set("items", value)
    return a
}

/**
 * statusLabel
 */
func (a *Tasks) StatusLabel(value interface{}) *Tasks {
    a.Set("statusLabel", value)
    return a
}

/**
 * statusTextMap
 */
func (a *Tasks) StatusTextMap(value interface{}) *Tasks {
    a.Set("statusTextMap", value)
    return a
}

/**
 * tableClassName
 */
func (a *Tasks) TableClassName(value interface{}) *Tasks {
    a.Set("tableClassName", value)
    return a
}

/**
 * loadingStatusCode
 */
func (a *Tasks) LoadingStatusCode(value interface{}) *Tasks {
    a.Set("loadingStatusCode", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Tasks) StaticLabelClassName(value interface{}) *Tasks {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * retryBtnText
 */
func (a *Tasks) RetryBtnText(value interface{}) *Tasks {
    a.Set("retryBtnText", value)
    return a
}

/**
 * className
 */
func (a *Tasks) ClassName(value interface{}) *Tasks {
    a.Set("className", value)
    return a
}

/**
 * staticOn
 */
func (a *Tasks) StaticOn(value interface{}) *Tasks {
    a.Set("staticOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *Tasks) StaticClassName(value interface{}) *Tasks {
    a.Set("staticClassName", value)
    return a
}

/**
 * reSubmitApi
 */
func (a *Tasks) ReSubmitApi(value interface{}) *Tasks {
    a.Set("reSubmitApi", value)
    return a
}

/**
 * statusLabelMap
 */
func (a *Tasks) StatusLabelMap(value interface{}) *Tasks {
    a.Set("statusLabelMap", value)
    return a
}

/**
 * disabled
 */
func (a *Tasks) Disabled(value interface{}) *Tasks {
    a.Set("disabled", value)
    return a
}

/**
 * visibleOn
 */
func (a *Tasks) VisibleOn(value interface{}) *Tasks {
    a.Set("visibleOn", value)
    return a
}

/**
 * interval
 */
func (a *Tasks) Interval(value interface{}) *Tasks {
    a.Set("interval", value)
    return a
}

/**
 * disabledOn
 */
func (a *Tasks) DisabledOn(value interface{}) *Tasks {
    a.Set("disabledOn", value)
    return a
}

/**
 * visible
 */
func (a *Tasks) Visible(value interface{}) *Tasks {
    a.Set("visible", value)
    return a
}

/**
 * editorSetting
 */
func (a *Tasks) EditorSetting(value interface{}) *Tasks {
    a.Set("editorSetting", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Tasks) UseMobileUI(value interface{}) *Tasks {
    a.Set("useMobileUI", value)
    return a
}

/**
 * checkApi
 */
func (a *Tasks) CheckApi(value interface{}) *Tasks {
    a.Set("checkApi", value)
    return a
}

/**
 * operationLabel
 */
func (a *Tasks) OperationLabel(value interface{}) *Tasks {
    a.Set("operationLabel", value)
    return a
}

/**
 * retryBtnClassName
 */
func (a *Tasks) RetryBtnClassName(value interface{}) *Tasks {
    a.Set("retryBtnClassName", value)
    return a
}

/**
 * initialStatusCode
 */
func (a *Tasks) InitialStatusCode(value interface{}) *Tasks {
    a.Set("initialStatusCode", value)
    return a
}

/**
 * staticSchema
 */
func (a *Tasks) StaticSchema(value interface{}) *Tasks {
    a.Set("staticSchema", value)
    return a
}
