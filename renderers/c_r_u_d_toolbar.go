package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type CRUDToolbar struct {
	*BaseRenderer
}

func NewCRUDToolbar() *CRUDToolbar {
    a := &CRUDToolbar{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}


func (a *CRUDToolbar) Set(name string, value interface{}) *CRUDToolbar {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 对齐方式
 * 可选值: left | right
 */
func (a *CRUDToolbar) Align(value interface{}) *CRUDToolbar {
    a.Set("align", value)
    return a
}
