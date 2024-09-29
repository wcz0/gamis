package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type ListenerAction struct {
	*BaseRenderer
}

func NewListenerAction() *ListenerAction {
    a := &ListenerAction{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}


func (a *ListenerAction) Set(name string, value interface{}) *ListenerAction {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 */
func (a *ListenerAction) Args(value interface{}) *ListenerAction {
    a.Set("args", value)
    return a
}

/**
 */
func (a *ListenerAction) Data(value interface{}) *ListenerAction {
    a.Set("data", value)
    return a
}

/**
 * 可选值: merge | override
 */
func (a *ListenerAction) DataMergeMode(value interface{}) *ListenerAction {
    a.Set("dataMergeMode", value)
    return a
}

/**
 */
func (a *ListenerAction) Expression(value interface{}) *ListenerAction {
    a.Set("expression", value)
    return a
}

/**
 */
func (a *ListenerAction) Description(value interface{}) *ListenerAction {
    a.Set("description", value)
    return a
}

/**
 */
func (a *ListenerAction) ComponentId(value interface{}) *ListenerAction {
    a.Set("componentId", value)
    return a
}

/**
 */
func (a *ListenerAction) ComponentName(value interface{}) *ListenerAction {
    a.Set("componentName", value)
    return a
}

/**
 */
func (a *ListenerAction) IgnoreError(value interface{}) *ListenerAction {
    a.Set("ignoreError", value)
    return a
}

/**
 */
func (a *ListenerAction) ExecOn(value interface{}) *ListenerAction {
    a.Set("execOn", value)
    return a
}

/**
 */
func (a *ListenerAction) ActionType(value interface{}) *ListenerAction {
    a.Set("actionType", value)
    return a
}

/**
 */
func (a *ListenerAction) OutputVar(value interface{}) *ListenerAction {
    a.Set("outputVar", value)
    return a
}

/**
 */
func (a *ListenerAction) PreventDefault(value interface{}) *ListenerAction {
    a.Set("preventDefault", value)
    return a
}

/**
 */
func (a *ListenerAction) StopPropagation(value interface{}) *ListenerAction {
    a.Set("stopPropagation", value)
    return a
}
