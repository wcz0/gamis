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
/**
 * 可选值: merge | override
 */
func (a *ListenerAction) DataMergeMode(value string) *ListenerAction {
    a.Set("dataMergeMode", value)
    return a
}

/**
 */
func (a *ListenerAction) OutputVar(value string) *ListenerAction {
    a.Set("outputVar", value)
    return a
}

/**
 */
func (a *ListenerAction) ActionType(value string) *ListenerAction {
    a.Set("actionType", value)
    return a
}

/**
 */
func (a *ListenerAction) ComponentId(value string) *ListenerAction {
    a.Set("componentId", value)
    return a
}

/**
 */
func (a *ListenerAction) Data(value string) *ListenerAction {
    a.Set("data", value)
    return a
}

/**
 */
func (a *ListenerAction) Args(value string) *ListenerAction {
    a.Set("args", value)
    return a
}

/**
 */
func (a *ListenerAction) PreventDefault(value string) *ListenerAction {
    a.Set("preventDefault", value)
    return a
}

/**
 */
func (a *ListenerAction) StopPropagation(value string) *ListenerAction {
    a.Set("stopPropagation", value)
    return a
}

/**
 */
func (a *ListenerAction) Expression(value string) *ListenerAction {
    a.Set("expression", value)
    return a
}

/**
 */
func (a *ListenerAction) ExecOn(value string) *ListenerAction {
    a.Set("execOn", value)
    return a
}

/**
 */
func (a *ListenerAction) Description(value string) *ListenerAction {
    a.Set("description", value)
    return a
}

/**
 */
func (a *ListenerAction) ComponentName(value string) *ListenerAction {
    a.Set("componentName", value)
    return a
}

/**
 */
func (a *ListenerAction) IgnoreError(value string) *ListenerAction {
    a.Set("ignoreError", value)
    return a
}
