package renderers


/**

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
/**
 */
func (a *ComboCondition) Test(value string) *ComboCondition {
    a.Set("test", value)
    return a
}

/**
 */
func (a *ComboCondition) Items(value string) *ComboCondition {
    a.Set("items", value)
    return a
}

/**
 */
func (a *ComboCondition) Label(value string) *ComboCondition {
    a.Set("label", value)
    return a
}

/**
 */
func (a *ComboCondition) Scaffold(value string) *ComboCondition {
    a.Set("scaffold", value)
    return a
}

/**
 */
func (a *ComboCondition) Mode(value string) *ComboCondition {
    a.Set("mode", value)
    return a
}
