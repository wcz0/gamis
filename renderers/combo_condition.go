package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type ComboCondition struct {
	*BaseRenderer
}

func NewComboCondition() *ComboCondition {
    a := &ComboCondition{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}


func (a *ComboCondition) Set(name string, value interface{}) *ComboCondition {
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
func (a *ComboCondition) Scaffold(value interface{}) *ComboCondition {
    a.Set("scaffold", value)
    return a
}

/**
 */
func (a *ComboCondition) Mode(value interface{}) *ComboCondition {
    a.Set("mode", value)
    return a
}

/**
 */
func (a *ComboCondition) Test(value interface{}) *ComboCondition {
    a.Set("test", value)
    return a
}

/**
 */
func (a *ComboCondition) Items(value interface{}) *ComboCondition {
    a.Set("items", value)
    return a
}

/**
 */
func (a *ComboCondition) Label(value interface{}) *ComboCondition {
    a.Set("label", value)
    return a
}
