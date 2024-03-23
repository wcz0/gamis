package renderers


/**

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
/**
 * 可选值: ROTATE_RIGHT | ROTATE_LEFT | ZOOM_IN | ZOOM_OUT | SCALE_ORIGIN
 */
func (a *ImageToolbarAction) Key(value string) *ImageToolbarAction {
    a.Set("key", value)
    return a
}

/**
 */
func (a *ImageToolbarAction) Label(value string) *ImageToolbarAction {
    a.Set("label", value)
    return a
}

/**
 */
func (a *ImageToolbarAction) Icon(value string) *ImageToolbarAction {
    a.Set("icon", value)
    return a
}

/**
 */
func (a *ImageToolbarAction) IconClassName(value string) *ImageToolbarAction {
    a.Set("iconClassName", value)
    return a
}

/**
 */
func (a *ImageToolbarAction) Disabled(value string) *ImageToolbarAction {
    a.Set("disabled", value)
    return a
}
