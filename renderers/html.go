package renderers

type Html struct {
	*BaseRenderer
}

func NewHtml() *Html {
	a := &Html{
		BaseRenderer: NewBaseRenderer(),
	}
	a.Set("type", "html")
	return a
}

func (h *Html) Set(name string, value interface{}) *Html {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	h.AmisSchema[name] = value
	return h
}