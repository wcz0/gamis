package renderers

type Barcode struct {
	*BaseRenderer
}

func NewBarcode() *Barcode {
	a := &Barcode{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "barcode")
	return a
}

func (a *Barcode) Set(name string, value interface{}) *Barcode {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Barcode) ClassName(value interface{}) *Barcode {
	a.Set("className", value)
	return a
}

func (a *Barcode) Type(value interface{}) *Barcode {
	a.Set("type", value)
	return a
}
