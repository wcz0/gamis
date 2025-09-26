package renderers


/**
 * 图片展示控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/image

 * @author wcz0
 * @version 6.2.2
 */
type Image struct {
	*BaseRenderer
}

func NewImage() *Image {
    a := &Image{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "image")
    return a
}


func (a *Image) Set(name string, value interface{}) *Image {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * title
 */
func (a *Image) Title(value interface{}) *Image {
    a.Set("title", value)
    return a
}

/**
 * staticOn
 */
func (a *Image) StaticOn(value interface{}) *Image {
    a.Set("staticOn", value)
    return a
}

/**
 * hidden
 */
func (a *Image) Hidden(value interface{}) *Image {
    a.Set("hidden", value)
    return a
}

/**
 * id
 */
func (a *Image) Id(value interface{}) *Image {
    a.Set("id", value)
    return a
}

/**
 * blank
 */
func (a *Image) Blank(value interface{}) *Image {
    a.Set("blank", value)
    return a
}

/**
 * toolbarActions
 */
func (a *Image) ToolbarActions(value interface{}) *Image {
    a.Set("toolbarActions", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Image) HiddenOn(value interface{}) *Image {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visible
 */
func (a *Image) Visible(value interface{}) *Image {
    a.Set("visible", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Image) StaticLabelClassName(value interface{}) *Image {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *Image) StaticSchema(value interface{}) *Image {
    a.Set("staticSchema", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Image) UseMobileUI(value interface{}) *Image {
    a.Set("useMobileUI", value)
    return a
}

/**
 * originalSrc
 */
func (a *Image) OriginalSrc(value interface{}) *Image {
    a.Set("originalSrc", value)
    return a
}

/**
 * enlargeAble
 */
func (a *Image) EnlargeAble(value interface{}) *Image {
    a.Set("enlargeAble", value)
    return a
}

/**
 * onEvent
 */
func (a *Image) OnEvent(value interface{}) *Image {
    a.Set("onEvent", value)
    return a
}

/**
 * editorSetting
 */
func (a *Image) EditorSetting(value interface{}) *Image {
    a.Set("editorSetting", value)
    return a
}

/**
 * width
 */
func (a *Image) Width(value interface{}) *Image {
    a.Set("width", value)
    return a
}

/**
 * sortType
 */
func (a *Image) SortType(value interface{}) *Image {
    a.Set("sortType", value)
    return a
}

/**
 * fontStyle
 */
func (a *Image) FontStyle(value interface{}) *Image {
    a.Set("fontStyle", value)
    return a
}

/**
 * className
 */
func (a *Image) ClassName(value interface{}) *Image {
    a.Set("className", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Image) StaticPlaceholder(value interface{}) *Image {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticClassName
 */
func (a *Image) StaticClassName(value interface{}) *Image {
    a.Set("staticClassName", value)
    return a
}

/**
 * style
 */
func (a *Image) Style(value interface{}) *Image {
    a.Set("style", value)
    return a
}

/**
 * testid
 */
func (a *Image) Testid(value interface{}) *Image {
    a.Set("testid", value)
    return a
}

/**
 * showToolbar
 */
func (a *Image) ShowToolbar(value interface{}) *Image {
    a.Set("showToolbar", value)
    return a
}

/**
 * name
 */
func (a *Image) Name(value interface{}) *Image {
    a.Set("name", value)
    return a
}

/**
 * src
 */
func (a *Image) Src(value interface{}) *Image {
    a.Set("src", value)
    return a
}

/**
 * disabledOn
 */
func (a *Image) DisabledOn(value interface{}) *Image {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Image) StaticInputClassName(value interface{}) *Image {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 可选值: image | static-image
 */
func (a *Image) Type(value interface{}) *Image {
    a.Set("type", value)
    return a
}

/**
 * height
 */
func (a *Image) Height(value interface{}) *Image {
    a.Set("height", value)
    return a
}

/**
 * innerClassName
 */
func (a *Image) InnerClassName(value interface{}) *Image {
    a.Set("innerClassName", value)
    return a
}

/**
 * imageClassName
 */
func (a *Image) ImageClassName(value interface{}) *Image {
    a.Set("imageClassName", value)
    return a
}

/**
 * static
 */
func (a *Image) Static(value interface{}) *Image {
    a.Set("static", value)
    return a
}

/**
 * imageCaption
 */
func (a *Image) ImageCaption(value interface{}) *Image {
    a.Set("imageCaption", value)
    return a
}

/**
 * enlargeWithGallary
 */
func (a *Image) EnlargeWithGallary(value interface{}) *Image {
    a.Set("enlargeWithGallary", value)
    return a
}

/**
 * alt
 */
func (a *Image) Alt(value interface{}) *Image {
    a.Set("alt", value)
    return a
}

/**
 * thumbClassName
 */
func (a *Image) ThumbClassName(value interface{}) *Image {
    a.Set("thumbClassName", value)
    return a
}

/**
 * imageGallaryClassName
 */
func (a *Image) ImageGallaryClassName(value interface{}) *Image {
    a.Set("imageGallaryClassName", value)
    return a
}

/**
 * imageMode
 */
func (a *Image) ImageMode(value interface{}) *Image {
    a.Set("imageMode", value)
    return a
}

/**
 * href
 */
func (a *Image) Href(value interface{}) *Image {
    a.Set("href", value)
    return a
}

/**
 * disabled
 */
func (a *Image) Disabled(value interface{}) *Image {
    a.Set("disabled", value)
    return a
}

/**
 * visibleOn
 */
func (a *Image) VisibleOn(value interface{}) *Image {
    a.Set("visibleOn", value)
    return a
}

/**
 * defaultImage
 */
func (a *Image) DefaultImage(value interface{}) *Image {
    a.Set("defaultImage", value)
    return a
}

/**
 * thumbMode
 */
func (a *Image) ThumbMode(value interface{}) *Image {
    a.Set("thumbMode", value)
    return a
}

/**
 * htmlTarget
 */
func (a *Image) HtmlTarget(value interface{}) *Image {
    a.Set("htmlTarget", value)
    return a
}

/**
 * maskColor
 */
func (a *Image) MaskColor(value interface{}) *Image {
    a.Set("maskColor", value)
    return a
}

/**
 * caption
 */
func (a *Image) Caption(value interface{}) *Image {
    a.Set("caption", value)
    return a
}

/**
 * thumbRatio
 */
func (a *Image) ThumbRatio(value interface{}) *Image {
    a.Set("thumbRatio", value)
    return a
}

/**
 * hoverMode
 */
func (a *Image) HoverMode(value interface{}) *Image {
    a.Set("hoverMode", value)
    return a
}
