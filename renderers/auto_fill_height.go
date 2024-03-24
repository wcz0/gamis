package renderers


/**

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
func (a *AutoFillHeight) Height(value string) *AutoFillHeight {
    a.Set("height", value)
    return a
}

/**
 */
func (a *AutoFillHeight) MaxHeight(value string) *AutoFillHeight {
    a.Set("maxHeight", value)
    return a
}
