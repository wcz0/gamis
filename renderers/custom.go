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