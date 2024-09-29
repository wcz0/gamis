package renderers

type Custom struct {
    *BaseRenderer
}

func NewCustom() *Custom {
    c := &Custom{
        BaseRenderer: NewBaseRenderer(),
    }
    c.Set("type", "custom")
    return c
}

func (c *Custom) Set(name string, value interface{}) *Custom {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    c.AmisSchema[name] = value
    return c
}

func (c *Custom) ClassName(value interface{}) *Custom {
    c.Set("className", value)
    return c
}



func (c *Custom) Html(value interface{}) *Custom {
    c.Set("html", value)
    return c
}

func (c *Custom) Id(value interface{}) *Custom {
    c.Set("id", value)
    return c
}

func (c *Custom) Inline(value interface{}) *Custom {
    c.Set("inline", value)
    return c
}

func (c *Custom) Name(value interface{}) *Custom {
    c.Set("name", value)
    return c
}

func (c *Custom) OnMount(value interface{}) *Custom {
    c.Set("onMount", value)
    return c
}

func (c *Custom) OnUnmount(value interface{}) *Custom {
    c.Set("onUnmount", value)
    return c
}

func (c *Custom) OnUpdate(value interface{}) *Custom {
    c.Set("onUpdate", value)
    return c
}

func (c *Custom) Type(value interface{}) *Custom {
    c.Set("type", value)
    return c
}