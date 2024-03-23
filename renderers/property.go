package renderers

type Property struct {
    *BaseRenderer
}

func NewProperty() *Property {
    p := &Property{
        BaseRenderer: NewBaseRenderer(),
    }
    p.Set("type", "property")
    return p
}

/**
    * 外层 dom 的类名
    */
func (p *Property) ClassName(value string) *Property {
    p.Set("className", value)
    return p
}

/**
    * 每行几列
    */
func (p *Property) Column(value string) *Property {
    p.Set("column", value)
    return p
}

/**
    * 属性值的样式
    */
func (p *Property) ContentStyle(value string) *Property {
    p.Set("contentStyle", value)
    return p
}

/**
    *
    */
func (p *Property) Items(value string) *Property {
    p.Set("items", value)
    return p
}

/**
    * 属性名的样式
    */
func (p *Property) LabelStyle(value string) *Property {
    p.Set("labelStyle", value)
    return p
}

/**
    * 显示模式，目前只有 'table' 和 'simple'
    */
func (p *Property) Mode(value string) *Property {
    p.Set("mode", value)
    return p
}

/**
    * 模式下属性名和值之间的分隔符
    */
func (p *Property) Separator(value string) *Property {
    p.Set("separator", value)
    return p
}

/**
    * 数据源
    */
func (p *Property) Source(value string) *Property {
    p.Set("source", value)
    return p
}

/**
    * 外层 dom 的样式
    */
func (p *Property) Style(value string) *Property {
    p.Set("style", value)
    return p
}

/**
    * 标题
    */
func (p *Property) Title(value string) *Property {
    p.Set("title", value)
    return p
}

/**
    * 指定为 property 渲染器。
    */
func (p *Property) Type(value string) *Property {
    p.Set("type", value)
    return p
}

