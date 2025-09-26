package renderers

type Watermark struct {
	*BaseRenderer
}

func NewWatermark() *Watermark {
	a := &Watermark{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "custom-watermark")
	return a
}

func (a *Watermark) Set(name string, value interface{}) *Watermark {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Watermark) Body(value interface{}) *Watermark {
	a.Set("body", value)
	return a
}

func (a *Watermark) ClassName(value interface{}) *Watermark {
	a.Set("className", value)
	return a
}

func (a *Watermark) Content(value interface{}) *Watermark {
	a.Set("content", value)
	return a
}

func (a *Watermark) Font(value interface{}) *Watermark {
	a.Set("font", value)
	return a
}

func (a *Watermark) Gap(value interface{}) *Watermark {
	a.Set("gap", value)
	return a
}

func (a *Watermark) Height(value interface{}) *Watermark {
	a.Set("height", value)
	return a
}

func (a *Watermark) Image(value interface{}) *Watermark {
	a.Set("image", value)
	return a
}

func (a *Watermark) Inherit(value interface{}) *Watermark {
	a.Set("inherit", value)
	return a
}

func (a *Watermark) Offset(value interface{}) *Watermark {
	a.Set("offset", value)
	return a
}

func (a *Watermark) Rotate(value interface{}) *Watermark {
	a.Set("rotate", value)
	return a
}

func (a *Watermark) Type(value interface{}) *Watermark {
	a.Set("type", value)
	return a
}

func (a *Watermark) Width(value interface{}) *Watermark {
	a.Set("width", value)
	return a
}

func (a *Watermark) ZIndex(value interface{}) *Watermark {
	a.Set("zIndex", value)
	return a
}
