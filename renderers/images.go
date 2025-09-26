package renderers

type Images struct {
	*BaseRenderer
}

func NewImages() *Images {
	a := &Images{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "images")
	return a
}

func (a *Images) Set(name string, value interface{}) *Images {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Images) ClassName(value interface{}) *Images {
	a.Set("className", value)
	return a
}

func (a *Images) CurrentIndex(value interface{}) *Images {
	a.Set("currentIndex", value)
	return a
}

func (a *Images) DefaultImage(value interface{}) *Images {
	a.Set("defaultImage", value)
	return a
}

func (a *Images) Delimiter(value interface{}) *Images {
	a.Set("delimiter", value)
	return a
}

func (a *Images) Disabled(value interface{}) *Images {
	a.Set("disabled", value)
	return a
}

func (a *Images) DisabledOn(value interface{}) *Images {
	a.Set("disabledOn", value)
	return a
}

func (a *Images) DisplayMode(value interface{}) *Images {
	a.Set("displayMode", value)
	return a
}

func (a *Images) EditorSetting(value interface{}) *Images {
	a.Set("editorSetting", value)
	return a
}

func (a *Images) EnlargeAble(value interface{}) *Images {
	a.Set("enlargeAble", value)
	return a
}

func (a *Images) EnlargetWithImages(value interface{}) *Images {
	a.Set("enlargetWithImages", value)
	return a
}

func (a *Images) FontStyle(value interface{}) *Images {
	a.Set("fontStyle", value)
	return a
}

func (a *Images) FullThumbMode(value interface{}) *Images {
	a.Set("fullThumbMode", value)
	return a
}

func (a *Images) Height(value interface{}) *Images {
	a.Set("height", value)
	return a
}

func (a *Images) Hidden(value interface{}) *Images {
	a.Set("hidden", value)
	return a
}

func (a *Images) HiddenOn(value interface{}) *Images {
	a.Set("hiddenOn", value)
	return a
}

func (a *Images) HoverMode(value interface{}) *Images {
	a.Set("hoverMode", value)
	return a
}

func (a *Images) Id(value interface{}) *Images {
	a.Set("id", value)
	return a
}

func (a *Images) ImageGallaryClassName(value interface{}) *Images {
	a.Set("imageGallaryClassName", value)
	return a
}

func (a *Images) ListClassName(value interface{}) *Images {
	a.Set("listClassName", value)
	return a
}

func (a *Images) MaskColor(value interface{}) *Images {
	a.Set("maskColor", value)
	return a
}

func (a *Images) Name(value interface{}) *Images {
	a.Set("name", value)
	return a
}

func (a *Images) OnEvent(value interface{}) *Images {
	a.Set("onEvent", value)
	return a
}

func (a *Images) Options(value interface{}) *Images {
	a.Set("options", value)
	return a
}

func (a *Images) OriginalSrc(value interface{}) *Images {
	a.Set("originalSrc", value)
	return a
}

func (a *Images) Placeholder(value interface{}) *Images {
	a.Set("placeholder", value)
	return a
}

func (a *Images) ShowDimensions(value interface{}) *Images {
	a.Set("showDimensions", value)
	return a
}

func (a *Images) ShowToolbar(value interface{}) *Images {
	a.Set("showToolbar", value)
	return a
}

func (a *Images) SortType(value interface{}) *Images {
	a.Set("sortType", value)
	return a
}

func (a *Images) Source(value interface{}) *Images {
	a.Set("source", value)
	return a
}

func (a *Images) Src(value interface{}) *Images {
	a.Set("src", value)
	return a
}

func (a *Images) Static(value interface{}) *Images {
	a.Set("static", value)
	return a
}

func (a *Images) StaticClassName(value interface{}) *Images {
	a.Set("staticClassName", value)
	return a
}

func (a *Images) StaticInputClassName(value interface{}) *Images {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Images) StaticLabelClassName(value interface{}) *Images {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Images) StaticOn(value interface{}) *Images {
	a.Set("staticOn", value)
	return a
}

func (a *Images) StaticPlaceholder(value interface{}) *Images {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Images) StaticSchema(value interface{}) *Images {
	a.Set("staticSchema", value)
	return a
}

func (a *Images) Style(value interface{}) *Images {
	a.Set("style", value)
	return a
}

func (a *Images) TestIdBuilder(value interface{}) *Images {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Images) Testid(value interface{}) *Images {
	a.Set("testid", value)
	return a
}

func (a *Images) ThumbMode(value interface{}) *Images {
	a.Set("thumbMode", value)
	return a
}

func (a *Images) ThumbRatio(value interface{}) *Images {
	a.Set("thumbRatio", value)
	return a
}

func (a *Images) ToolbarActions(value interface{}) *Images {
	a.Set("toolbarActions", value)
	return a
}

func (a *Images) Type(value interface{}) *Images {
	a.Set("type", value)
	return a
}

func (a *Images) UseMobileUI(value interface{}) *Images {
	a.Set("useMobileUI", value)
	return a
}

func (a *Images) Value(value interface{}) *Images {
	a.Set("value", value)
	return a
}

func (a *Images) Visible(value interface{}) *Images {
	a.Set("visible", value)
	return a
}

func (a *Images) VisibleOn(value interface{}) *Images {
	a.Set("visibleOn", value)
	return a
}

func (a *Images) Width(value interface{}) *Images {
	a.Set("width", value)
	return a
}
