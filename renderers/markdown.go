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

/**
 * 类名
 */
func (m *Markdown) ClassName(value string) *Markdown {
    m.Set("className", value)
    return m
}

/**
 * 名称
 */
func (m *Markdown) Name(value string) *Markdown {
    m.Set("name", value)
    return m
}

/**
 * 外部地址
 */
func (m *Markdown) Src(value string) *Markdown {
    m.Set("src", value)
    return m
}

/**
 * 静态值
 */
func (m *Markdown) Value(value string) *Markdown {
    m.Set("value", value)
    return m
}