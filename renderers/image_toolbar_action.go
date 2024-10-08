package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type ImageToolbarAction struct {
	*BaseRenderer
}

func NewImageToolbarAction() *ImageToolbarAction {
    a := &ImageToolbarAction{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("key", "ROTATE_RIGHT")
    return a
}


func (a *ImageToolbarAction) Set(name string, value interface{}) *ImageToolbarAction {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 可选值: ROTATE_RIGHT | ROTATE_LEFT | ZOOM_IN | ZOOM_OUT | SCALE_ORIGIN
 */
func (a *ImageToolbarAction) Key(value interface{}) *ImageToolbarAction {
    a.Set("key", value)
    return a
}

/**
 */
func (a *ImageToolbarAction) Label(value interface{}) *ImageToolbarAction {
    a.Set("label", value)
    return a
}

/**
 */
func (a *ImageToolbarAction) Icon(value interface{}) *ImageToolbarAction {
    a.Set("icon", value)
    return a
}

/**
 */
func (a *ImageToolbarAction) IconClassName(value interface{}) *ImageToolbarAction {
    a.Set("iconClassName", value)
    return a
}

/**
 */
func (a *ImageToolbarAction) Disabled(value interface{}) *ImageToolbarAction {
    a.Set("disabled", value)
    return a
}
