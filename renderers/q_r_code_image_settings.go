package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type QRCodeImageSettings struct {
	*BaseRenderer
}

func NewQRCodeImageSettings() *QRCodeImageSettings {
    a := &QRCodeImageSettings{
        BaseRenderer: NewBaseRenderer(),
    }
    return a
}

/**
 */
func (a *QRCodeImageSettings) Excavate(value interface{}) *QRCodeImageSettings {
    a.Set("excavate", value)
    return a
}

/**
 */
func (a *QRCodeImageSettings) X(value interface{}) *QRCodeImageSettings {
    a.Set("x", value)
    return a
}

/**
 */
func (a *QRCodeImageSettings) Y(value interface{}) *QRCodeImageSettings {
    a.Set("y", value)
    return a
}

/**
 */
func (a *QRCodeImageSettings) Src(value interface{}) *QRCodeImageSettings {
    a.Set("src", value)
    return a
}

/**
 */
func (a *QRCodeImageSettings) Height(value interface{}) *QRCodeImageSettings {
    a.Set("height", value)
    return a
}

/**
 */
func (a *QRCodeImageSettings) Width(value interface{}) *QRCodeImageSettings {
    a.Set("width", value)
    return a
}
