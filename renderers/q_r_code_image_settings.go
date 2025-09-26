package renderers

type QRCodeImageSettings struct {
	*BaseRenderer
}

func NewQRCodeImageSettings() *QRCodeImageSettings {
	a := &QRCodeImageSettings{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *QRCodeImageSettings) Set(name string, value interface{}) *QRCodeImageSettings {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *QRCodeImageSettings) Excavate(value interface{}) *QRCodeImageSettings {
	a.Set("excavate", value)
	return a
}

func (a *QRCodeImageSettings) Height(value interface{}) *QRCodeImageSettings {
	a.Set("height", value)
	return a
}

func (a *QRCodeImageSettings) Src(value interface{}) *QRCodeImageSettings {
	a.Set("src", value)
	return a
}

func (a *QRCodeImageSettings) Width(value interface{}) *QRCodeImageSettings {
	a.Set("width", value)
	return a
}

func (a *QRCodeImageSettings) X(value interface{}) *QRCodeImageSettings {
	a.Set("x", value)
	return a
}

func (a *QRCodeImageSettings) Y(value interface{}) *QRCodeImageSettings {
	a.Set("y", value)
	return a
}
