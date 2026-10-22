package renderers

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

func (a *CRUDToolbar) Align(value interface{}) *CRUDToolbar {
	a.Set("align", value)
	return a
}

func (a *CRUDToolbar) ItemCheckableOn(value interface{}) *CRUDToolbar {
	a.Set("itemCheckableOn", value)
	return a
}
