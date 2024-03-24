package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type IconChecked struct {
	*BaseRenderer
}

func NewIconChecked() *IconChecked {
    a := &IconChecked{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}
/**
 */
func (a *IconChecked) Id(value interface{}) *IconChecked {
    a.Set("id", value)
    return a
}

/**
 */
func (a *IconChecked) Name(value interface{}) *IconChecked {
    a.Set("name", value)
    return a
}

/**
 */
func (a *IconChecked) Svg(value interface{}) *IconChecked {
    a.Set("svg", value)
    return a
}
