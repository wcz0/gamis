package renderers


/**
 * 二维码展示控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/qrcode

 * @author wcz0
 * @version 6.2.2
 */
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
/**
 * id
 */
func (a *QRCode) Id(value interface{}) *QRCode {
    a.Set("id", value)
    return a
}

/**
 * name
 */
func (a *QRCode) Name(value interface{}) *QRCode {
    a.Set("name", value)
    return a
}

/**
 * codeSize
 */
func (a *QRCode) CodeSize(value interface{}) *QRCode {
    a.Set("codeSize", value)
    return a
}

/**
 * foregroundColor
 */
func (a *QRCode) ForegroundColor(value interface{}) *QRCode {
    a.Set("foregroundColor", value)
    return a
}

/**
 * eyeType
 */
func (a *QRCode) EyeType(value interface{}) *QRCode {
    a.Set("eyeType", value)
    return a
}

/**
 * visibleOn
 */
func (a *QRCode) VisibleOn(value interface{}) *QRCode {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticSchema
 */
func (a *QRCode) StaticSchema(value interface{}) *QRCode {
    a.Set("staticSchema", value)
    return a
}

/**
 * editorSetting
 */
func (a *QRCode) EditorSetting(value interface{}) *QRCode {
    a.Set("editorSetting", value)
    return a
}

/**
 * qrcodeClassName
 */
func (a *QRCode) QrcodeClassName(value interface{}) *QRCode {
    a.Set("qrcodeClassName", value)
    return a
}

/**
 * placeholder
 */
func (a *QRCode) Placeholder(value interface{}) *QRCode {
    a.Set("placeholder", value)
    return a
}

/**
 * eyeBorderColor
 */
func (a *QRCode) EyeBorderColor(value interface{}) *QRCode {
    a.Set("eyeBorderColor", value)
    return a
}

/**
 * pointSizeRandom
 */
func (a *QRCode) PointSizeRandom(value interface{}) *QRCode {
    a.Set("pointSizeRandom", value)
    return a
}

/**
 * hidden
 */
func (a *QRCode) Hidden(value interface{}) *QRCode {
    a.Set("hidden", value)
    return a
}

/**
 * staticOn
 */
func (a *QRCode) StaticOn(value interface{}) *QRCode {
    a.Set("staticOn", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *QRCode) StaticLabelClassName(value interface{}) *QRCode {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * style
 */
func (a *QRCode) Style(value interface{}) *QRCode {
    a.Set("style", value)
    return a
}

/**
 * useMobileUI
 */
func (a *QRCode) UseMobileUI(value interface{}) *QRCode {
    a.Set("useMobileUI", value)
    return a
}

/**
 * testid
 */
func (a *QRCode) Testid(value interface{}) *QRCode {
    a.Set("testid", value)
    return a
}

/**
 * pointType
 */
func (a *QRCode) PointType(value interface{}) *QRCode {
    a.Set("pointType", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *QRCode) StaticInputClassName(value interface{}) *QRCode {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * mode
 */
func (a *QRCode) Mode(value interface{}) *QRCode {
    a.Set("mode", value)
    return a
}

/**
 * pointSize
 */
func (a *QRCode) PointSize(value interface{}) *QRCode {
    a.Set("pointSize", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *QRCode) StaticPlaceholder(value interface{}) *QRCode {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * backgroundColor
 */
func (a *QRCode) BackgroundColor(value interface{}) *QRCode {
    a.Set("backgroundColor", value)
    return a
}

/**
 * eyeBorderSize
 */
func (a *QRCode) EyeBorderSize(value interface{}) *QRCode {
    a.Set("eyeBorderSize", value)
    return a
}

/**
 * className
 */
func (a *QRCode) ClassName(value interface{}) *QRCode {
    a.Set("className", value)
    return a
}

/**
 * hiddenOn
 */
func (a *QRCode) HiddenOn(value interface{}) *QRCode {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visible
 */
func (a *QRCode) Visible(value interface{}) *QRCode {
    a.Set("visible", value)
    return a
}

/**
 * level
 */
func (a *QRCode) Level(value interface{}) *QRCode {
    a.Set("level", value)
    return a
}

/**
 * eyeInnerColor
 */
func (a *QRCode) EyeInnerColor(value interface{}) *QRCode {
    a.Set("eyeInnerColor", value)
    return a
}

/**
 * disabledOn
 */
func (a *QRCode) DisabledOn(value interface{}) *QRCode {
    a.Set("disabledOn", value)
    return a
}

/**
 * onEvent
 */
func (a *QRCode) OnEvent(value interface{}) *QRCode {
    a.Set("onEvent", value)
    return a
}

/**
 * disabled
 */
func (a *QRCode) Disabled(value interface{}) *QRCode {
    a.Set("disabled", value)
    return a
}

/**
 * static
 */
func (a *QRCode) Static(value interface{}) *QRCode {
    a.Set("static", value)
    return a
}

/**
 * staticClassName
 */
func (a *QRCode) StaticClassName(value interface{}) *QRCode {
    a.Set("staticClassName", value)
    return a
}

/**
 * 可选值: qrcode | qr-code
 */
func (a *QRCode) Type(value interface{}) *QRCode {
    a.Set("type", value)
    return a
}

/**
 * imageSettings
 */
func (a *QRCode) ImageSettings(value interface{}) *QRCode {
    a.Set("imageSettings", value)
    return a
}
