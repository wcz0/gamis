package renderers

type SvgIcon struct {
	*BaseRenderer
}

func NewSvgIcon() *SvgIcon {
	a := &SvgIcon{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "custom-svg-icon")
	return a
}

func (a *SvgIcon) Set(name string, value interface{}) *SvgIcon {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *SvgIcon) ClassName(value interface{}) *SvgIcon {
	a.Set("className", value)
	return a
}

func (a *SvgIcon) Icon(value interface{}) *SvgIcon {
	a.Set("icon", value)
	return a
}

func (a *SvgIcon) Type(value interface{}) *SvgIcon {
	a.Set("type", value)
	return a
}
