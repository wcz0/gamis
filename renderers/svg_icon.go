package renderers

type SvgIcon struct {
    *BaseRenderer
}

func NewSvgIcon() *SvgIcon {
    s := &SvgIcon{
        BaseRenderer: NewBaseRenderer(),
    }
    s.Set("type", "svg-icon")
    return s
}

func (s *SvgIcon) ClassName(value string) *SvgIcon {
    s.Set("className", value)
    return s
}

func (s *SvgIcon) Icon(value string) *SvgIcon {
    s.Set("icon", value)
    return s
}

func (s *SvgIcon) Type(value string) *SvgIcon {
    s.Set("type", value)
    return s
}