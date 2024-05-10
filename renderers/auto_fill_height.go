package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type AutoFillHeight struct {
	*BaseRenderer
}

func NewAutoFillHeight() *AutoFillHeight {
    a := &AutoFillHeight{
        BaseRenderer: NewBaseRenderer(),
    }
    return a
}
/**
 */
func (a *AutoFillHeight) Height(value interface{}) *AutoFillHeight {
    a.Set("height", value)
    return a
}

/**
 */
func (a *AutoFillHeight) MaxHeight(value interface{}) *AutoFillHeight {
    a.Set("maxHeight", value)
    return a
}
