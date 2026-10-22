package renderers

type CRUD2 struct {
	*BaseRenderer
}

func NewCRUD2() *CRUD2 {
	a := &CRUD2{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "crud2")

	return a
}

func (a *CRUD2) Set(name string, value interface{}) *CRUD2 {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *CRUD2) ItemCheckableOn(value interface{}) *CRUD2 {
	a.Set("itemCheckableOn", value)
	return a
}
