package renderers


/**
 * 二维码展示控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/qrcode
 *

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
/**
 * 占位符
 */
func (a *QRCode) Placeholder(value string) *QRCode {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *QRCode) DisabledOn(value string) *QRCode {
    a.Set("disabledOn", value)
    return a
}

/**
 */
func (a *QRCode) StaticSchema(value string) *QRCode {
    a.Set("staticSchema", value)
    return a
}

/**
 * css 类名
 */
func (a *QRCode) QrcodeClassName(value string) *QRCode {
    a.Set("qrcodeClassName", value)
    return a
}

/**
 * 二维码的宽高大小，默认 128
 */
func (a *QRCode) CodeSize(value string) *QRCode {
    a.Set("codeSize", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *QRCode) UseMobileUI(value string) *QRCode {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 背景色
 */
func (a *QRCode) BackgroundColor(value string) *QRCode {
    a.Set("backgroundColor", value)
    return a
}

/**
 * 关联字段名。
 */
func (a *QRCode) Name(value string) *QRCode {
    a.Set("name", value)
    return a
}

/**
 * 前景色
 */
func (a *QRCode) ForegroundColor(value string) *QRCode {
    a.Set("foregroundColor", value)
    return a
}

/**
 * 是否禁用
 */
func (a *QRCode) Disabled(value string) *QRCode {
    a.Set("disabled", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *QRCode) VisibleOn(value string) *QRCode {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *QRCode) StaticLabelClassName(value string) *QRCode {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *QRCode) EditorSetting(value string) *QRCode {
    a.Set("editorSetting", value)
    return a
}

/**
 * 是否显示
 */
func (a *QRCode) Visible(value string) *QRCode {
    a.Set("visible", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *QRCode) Id(value string) *QRCode {
    a.Set("id", value)
    return a
}

/**
 * 可选值: qrcode | qr-code
 */
func (a *QRCode) Type(value string) *QRCode {
    a.Set("type", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *QRCode) HiddenOn(value string) *QRCode {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *QRCode) StaticPlaceholder(value string) *QRCode {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 组件样式
 */
func (a *QRCode) Style(value string) *QRCode {
    a.Set("style", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *QRCode) Hidden(value string) *QRCode {
    a.Set("hidden", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *QRCode) OnEvent(value string) *QRCode {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *QRCode) StaticInputClassName(value string) *QRCode {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 二维码复杂级别
 * 可选值: L | M | Q | H
 */
func (a *QRCode) Level(value string) *QRCode {
    a.Set("level", value)
    return a
}

/**
 * 图片配置
 */
func (a *QRCode) ImageSettings(value string) *QRCode {
    a.Set("imageSettings", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *QRCode) StaticClassName(value string) *QRCode {
    a.Set("staticClassName", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *QRCode) ClassName(value string) *QRCode {
    a.Set("className", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *QRCode) Static(value string) *QRCode {
    a.Set("static", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *QRCode) StaticOn(value string) *QRCode {
    a.Set("staticOn", value)
    return a
}
