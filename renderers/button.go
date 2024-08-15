package renderers

type Button struct {
	*VanillaAction
}

func NewButton() *Button{
	return &Button{
		VanillaAction: NewVanillaAction(),
	}
}

func (b *Button) Set(name string, value interface{}) *Button {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	b.AmisSchema[name] = value
	return b
}