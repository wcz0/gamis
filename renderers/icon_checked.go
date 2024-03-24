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
func (a *IconChecked) Id(value string) *IconChecked {
    a.Set("id", value)
    return a
}

/**
 */
func (a *IconChecked) Name(value string) *IconChecked {
    a.Set("name", value)
    return a
}

/**
 */
func (a *IconChecked) Svg(value string) *IconChecked {
    a.Set("svg", value)
    return a
}
