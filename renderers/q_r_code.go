package renderers

type QRCode struct {
	*BaseRenderer
}

func NewQRCode() *QRCode {
	a := &QRCode{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "qrcode")
	return a
}

func (a *QRCode) Set(name string, value interface{}) *QRCode {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *QRCode) BackgroundColor(value interface{}) *QRCode {
	a.Set("backgroundColor", value)
	return a
}

func (a *QRCode) ClassName(value interface{}) *QRCode {
	a.Set("className", value)
	return a
}

func (a *QRCode) CodeSize(value interface{}) *QRCode {
	a.Set("codeSize", value)
	return a
}

func (a *QRCode) Disabled(value interface{}) *QRCode {
	a.Set("disabled", value)
	return a
}

func (a *QRCode) DisabledOn(value interface{}) *QRCode {
	a.Set("disabledOn", value)
	return a
}

func (a *QRCode) EditorSetting(value interface{}) *QRCode {
	a.Set("editorSetting", value)
	return a
}

func (a *QRCode) EyeBorderColor(value interface{}) *QRCode {
	a.Set("eyeBorderColor", value)
	return a
}

func (a *QRCode) EyeBorderSize(value interface{}) *QRCode {
	a.Set("eyeBorderSize", value)
	return a
}

func (a *QRCode) EyeInnerColor(value interface{}) *QRCode {
	a.Set("eyeInnerColor", value)
	return a
}

func (a *QRCode) EyeType(value interface{}) *QRCode {
	a.Set("eyeType", value)
	return a
}

func (a *QRCode) ForegroundColor(value interface{}) *QRCode {
	a.Set("foregroundColor", value)
	return a
}

func (a *QRCode) Hidden(value interface{}) *QRCode {
	a.Set("hidden", value)
	return a
}

func (a *QRCode) HiddenOn(value interface{}) *QRCode {
	a.Set("hiddenOn", value)
	return a
}

func (a *QRCode) Id(value interface{}) *QRCode {
	a.Set("id", value)
	return a
}

func (a *QRCode) ImageSettings(value interface{}) *QRCode {
	a.Set("imageSettings", value)
	return a
}

func (a *QRCode) Level(value interface{}) *QRCode {
	a.Set("level", value)
	return a
}

func (a *QRCode) Mode(value interface{}) *QRCode {
	a.Set("mode", value)
	return a
}

func (a *QRCode) Name(value interface{}) *QRCode {
	a.Set("name", value)
	return a
}

func (a *QRCode) OnEvent(value interface{}) *QRCode {
	a.Set("onEvent", value)
	return a
}

func (a *QRCode) Placeholder(value interface{}) *QRCode {
	a.Set("placeholder", value)
	return a
}

func (a *QRCode) PointSize(value interface{}) *QRCode {
	a.Set("pointSize", value)
	return a
}

func (a *QRCode) PointSizeRandom(value interface{}) *QRCode {
	a.Set("pointSizeRandom", value)
	return a
}

func (a *QRCode) PointType(value interface{}) *QRCode {
	a.Set("pointType", value)
	return a
}

func (a *QRCode) QrcodeClassName(value interface{}) *QRCode {
	a.Set("qrcodeClassName", value)
	return a
}

func (a *QRCode) Static(value interface{}) *QRCode {
	a.Set("static", value)
	return a
}

func (a *QRCode) StaticClassName(value interface{}) *QRCode {
	a.Set("staticClassName", value)
	return a
}

func (a *QRCode) StaticInputClassName(value interface{}) *QRCode {
	a.Set("staticInputClassName", value)
	return a
}

func (a *QRCode) StaticLabelClassName(value interface{}) *QRCode {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *QRCode) StaticOn(value interface{}) *QRCode {
	a.Set("staticOn", value)
	return a
}

func (a *QRCode) StaticPlaceholder(value interface{}) *QRCode {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *QRCode) StaticSchema(value interface{}) *QRCode {
	a.Set("staticSchema", value)
	return a
}

func (a *QRCode) Style(value interface{}) *QRCode {
	a.Set("style", value)
	return a
}

func (a *QRCode) TestIdBuilder(value interface{}) *QRCode {
	a.Set("testIdBuilder", value)
	return a
}

func (a *QRCode) Testid(value interface{}) *QRCode {
	a.Set("testid", value)
	return a
}

func (a *QRCode) Type(value interface{}) *QRCode {
	a.Set("type", value)
	return a
}

func (a *QRCode) UseMobileUI(value interface{}) *QRCode {
	a.Set("useMobileUI", value)
	return a
}

func (a *QRCode) Visible(value interface{}) *QRCode {
	a.Set("visible", value)
	return a
}

func (a *QRCode) VisibleOn(value interface{}) *QRCode {
	a.Set("visibleOn", value)
	return a
}
