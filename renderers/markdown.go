package renderers

type Markdown struct {
	*BaseRenderer
}

func NewMarkdown() *Markdown {
	m := &Markdown{
		BaseRenderer: NewBaseRenderer(),
	}
	m.Set("type", "markdown")
	return m
}

func (m *Markdown) Set(name string, value interface{}) *Markdown {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	m.AmisSchema[name] = value
	return m
}

/**
 * 类名
 */
func (m *Markdown) ClassName(value interface{}) *Markdown {
	m.Set("className", value)
	return m
}

/**
 * 名称
 */
func (m *Markdown) Name(value interface{}) *Markdown {
	m.Set("name", value)
	return m
}

/**
 * 外部地址
 */
func (m *Markdown) Src(value interface{}) *Markdown {
	m.Set("src", value)
	return m
}

/**
 * 静态值
 */
func (m *Markdown) Value(value interface{}) *Markdown {
	m.Set("value", value)
	return m
}
