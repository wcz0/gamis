package renderers

type Badge struct {
	*BaseRenderer
}

func NewBadge() *Badge {
	a := &Badge{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *Badge) Set(name string, value interface{}) *Badge {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Badge) Animation(value interface{}) *Badge {
	a.Set("animation", value)
	return a
}

func (a *Badge) ClassName(value interface{}) *Badge {
	a.Set("className", value)
	return a
}

func (a *Badge) Level(value interface{}) *Badge {
	a.Set("level", value)
	return a
}

func (a *Badge) Mode(value interface{}) *Badge {
	a.Set("mode", value)
	return a
}

func (a *Badge) Offset(value interface{}) *Badge {
	a.Set("offset", value)
	return a
}

func (a *Badge) OverflowCount(value interface{}) *Badge {
	a.Set("overflowCount", value)
	return a
}

func (a *Badge) Position(value interface{}) *Badge {
	a.Set("position", value)
	return a
}

func (a *Badge) Size(value interface{}) *Badge {
	a.Set("size", value)
	return a
}

func (a *Badge) Style(value interface{}) *Badge {
	a.Set("style", value)
	return a
}

func (a *Badge) Text(value interface{}) *Badge {
	a.Set("text", value)
	return a
}

func (a *Badge) VisibleOn(value interface{}) *Badge {
	a.Set("visibleOn", value)
	return a
}
