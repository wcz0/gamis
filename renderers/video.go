package renderers

type Video struct {
	*BaseRenderer
}

func NewVideo() *Video {
	a := &Video{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "video")
	return a
}

func (a *Video) Set(name string, value interface{}) *Video {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Video) AspectRatio(value interface{}) *Video {
	a.Set("aspectRatio", value)
	return a
}

func (a *Video) AutoPlay(value interface{}) *Video {
	a.Set("autoPlay", value)
	return a
}

func (a *Video) ClassName(value interface{}) *Video {
	a.Set("className", value)
	return a
}

func (a *Video) ColumnsCount(value interface{}) *Video {
	a.Set("columnsCount", value)
	return a
}

func (a *Video) Disabled(value interface{}) *Video {
	a.Set("disabled", value)
	return a
}

func (a *Video) DisabledOn(value interface{}) *Video {
	a.Set("disabledOn", value)
	return a
}

func (a *Video) EditorSetting(value interface{}) *Video {
	a.Set("editorSetting", value)
	return a
}

func (a *Video) Frames(value interface{}) *Video {
	a.Set("frames", value)
	return a
}

func (a *Video) FramesClassName(value interface{}) *Video {
	a.Set("framesClassName", value)
	return a
}

func (a *Video) Hidden(value interface{}) *Video {
	a.Set("hidden", value)
	return a
}

func (a *Video) HiddenOn(value interface{}) *Video {
	a.Set("hiddenOn", value)
	return a
}

func (a *Video) Id(value interface{}) *Video {
	a.Set("id", value)
	return a
}

func (a *Video) IsLive(value interface{}) *Video {
	a.Set("isLive", value)
	return a
}

func (a *Video) JumpBufferDuration(value interface{}) *Video {
	a.Set("jumpBufferDuration", value)
	return a
}

func (a *Video) JumpFrame(value interface{}) *Video {
	a.Set("jumpFrame", value)
	return a
}

func (a *Video) Loop(value interface{}) *Video {
	a.Set("loop", value)
	return a
}

func (a *Video) Muted(value interface{}) *Video {
	a.Set("muted", value)
	return a
}

func (a *Video) OnEvent(value interface{}) *Video {
	a.Set("onEvent", value)
	return a
}

func (a *Video) PlayerClassName(value interface{}) *Video {
	a.Set("playerClassName", value)
	return a
}

func (a *Video) Poster(value interface{}) *Video {
	a.Set("poster", value)
	return a
}

func (a *Video) Rates(value interface{}) *Video {
	a.Set("rates", value)
	return a
}

func (a *Video) SplitPoster(value interface{}) *Video {
	a.Set("splitPoster", value)
	return a
}

func (a *Video) Src(value interface{}) *Video {
	a.Set("src", value)
	return a
}

func (a *Video) Static(value interface{}) *Video {
	a.Set("static", value)
	return a
}

func (a *Video) StaticClassName(value interface{}) *Video {
	a.Set("staticClassName", value)
	return a
}

func (a *Video) StaticInputClassName(value interface{}) *Video {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Video) StaticLabelClassName(value interface{}) *Video {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Video) StaticOn(value interface{}) *Video {
	a.Set("staticOn", value)
	return a
}

func (a *Video) StaticPlaceholder(value interface{}) *Video {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Video) StaticSchema(value interface{}) *Video {
	a.Set("staticSchema", value)
	return a
}

func (a *Video) StopOnNextFrame(value interface{}) *Video {
	a.Set("stopOnNextFrame", value)
	return a
}

func (a *Video) Style(value interface{}) *Video {
	a.Set("style", value)
	return a
}

func (a *Video) TestIdBuilder(value interface{}) *Video {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Video) Testid(value interface{}) *Video {
	a.Set("testid", value)
	return a
}

func (a *Video) Type(value interface{}) *Video {
	a.Set("type", value)
	return a
}

func (a *Video) UseMobileUI(value interface{}) *Video {
	a.Set("useMobileUI", value)
	return a
}

func (a *Video) VideoType(value interface{}) *Video {
	a.Set("videoType", value)
	return a
}

func (a *Video) Visible(value interface{}) *Video {
	a.Set("visible", value)
	return a
}

func (a *Video) VisibleOn(value interface{}) *Video {
	a.Set("visibleOn", value)
	return a
}
