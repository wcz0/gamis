package renderers

type ChartRadios struct {
	*BaseRenderer
}

func NewChartRadios() *ChartRadios {
	c := &ChartRadios{
		BaseRenderer: NewBaseRenderer(),
	}
	c.Set("type", "chart-radios")
	return c
}

/**
 * 图表数值字段名
 */

func (c *ChartRadios) ChartValueField(value interface{}) *ChartRadios {
	c.Set("chartValueField", value)
	return c
}

/**
 * 图表配置
 */
func (c *ChartRadios) Config(value interface{}) *ChartRadios {
	c.Set("config", value)
	return c
}

/**
 * 高亮的时候是否显示 tooltip
 */
func (c *ChartRadios) ShowTooltipOnHighlight(value bool) *ChartRadios {
	c.Set("showTooltipOnHighlight", value)
	return c
}

/**
 * 指定为 chart-radios 渲染器。
 */
func (c *ChartRadios) Type(value interface{}) *ChartRadios {
	c.Set("type", value)
	return c
}
