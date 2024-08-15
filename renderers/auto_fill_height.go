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


func (a *AutoFillHeight) Set(name string, value interface{}) *AutoFillHeight {
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
