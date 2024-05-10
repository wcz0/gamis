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

func (s *SvgIcon) ClassName(value interface{}) *SvgIcon {
	s.Set("className", value)
	return s
}

func (s *SvgIcon) Icon(value interface{}) *SvgIcon {
	s.Set("icon", value)
	return s
}

func (s *SvgIcon) Type(value interface{}) *SvgIcon {
	s.Set("type", value)
	return s
}
