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
