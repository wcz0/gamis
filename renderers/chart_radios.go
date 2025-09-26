package renderers

type ChartRadios struct {
	*BaseRenderer
}

func NewChartRadios() *ChartRadios {
	a := &ChartRadios{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "chart-radios")
	return a
}

func (a *ChartRadios) Set(name string, value interface{}) *ChartRadios {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *ChartRadios) ChartValueField(value interface{}) *ChartRadios {
	a.Set("chartValueField", value)
	return a
}

func (a *ChartRadios) Config(value interface{}) *ChartRadios {
	a.Set("config", value)
	return a
}

func (a *ChartRadios) ShowTooltipOnHighlight(value interface{}) *ChartRadios {
	a.Set("showTooltipOnHighlight", value)
	return a
}

func (a *ChartRadios) Type(value interface{}) *ChartRadios {
	a.Set("type", value)
	return a
}
