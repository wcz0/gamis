package renderers


/**
 * 图片集展示控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/images

 * @author wcz0
 * @version 6.2.2
 */
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
/**
 * listClassName
 */
func (a *Images) ListClassName(value interface{}) *Images {
    a.Set("listClassName", value)
    return a
}

/**
 * width
 */
func (a *Images) Width(value interface{}) *Images {
    a.Set("width", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Images) StaticInputClassName(value interface{}) *Images {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * editorSetting
 */
func (a *Images) EditorSetting(value interface{}) *Images {
    a.Set("editorSetting", value)
    return a
}

/**
 * options
 */
func (a *Images) Options(value interface{}) *Images {
    a.Set("options", value)
    return a
}

/**
 * enlargeAble
 */
func (a *Images) EnlargeAble(value interface{}) *Images {
    a.Set("enlargeAble", value)
    return a
}

/**
 * toolbarActions
 */
func (a *Images) ToolbarActions(value interface{}) *Images {
    a.Set("toolbarActions", value)
    return a
}

/**
 * currentIndex
 */
func (a *Images) CurrentIndex(value interface{}) *Images {
    a.Set("currentIndex", value)
    return a
}

/**
 * thumbMode
 */
func (a *Images) ThumbMode(value interface{}) *Images {
    a.Set("thumbMode", value)
    return a
}

/**
 * height
 */
func (a *Images) Height(value interface{}) *Images {
    a.Set("height", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Images) HiddenOn(value interface{}) *Images {
    a.Set("hiddenOn", value)
    return a
}

/**
 * id
 */
func (a *Images) Id(value interface{}) *Images {
    a.Set("id", value)
    return a
}

/**
 * onEvent
 */
func (a *Images) OnEvent(value interface{}) *Images {
    a.Set("onEvent", value)
    return a
}

/**
 * static
 */
func (a *Images) Static(value interface{}) *Images {
    a.Set("static", value)
    return a
}

/**
 * staticClassName
 */
func (a *Images) StaticClassName(value interface{}) *Images {
    a.Set("staticClassName", value)
    return a
}

/**
 * name
 */
func (a *Images) Name(value interface{}) *Images {
    a.Set("name", value)
    return a
}

/**
 * showDimensions
 */
func (a *Images) ShowDimensions(value interface{}) *Images {
    a.Set("showDimensions", value)
    return a
}

/**
 * disabled
 */
func (a *Images) Disabled(value interface{}) *Images {
    a.Set("disabled", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Images) StaticLabelClassName(value interface{}) *Images {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * placeholder
 */
func (a *Images) Placeholder(value interface{}) *Images {
    a.Set("placeholder", value)
    return a
}

/**
 * source
 */
func (a *Images) Source(value interface{}) *Images {
    a.Set("source", value)
    return a
}

/**
 * fullThumbMode
 */
func (a *Images) FullThumbMode(value interface{}) *Images {
    a.Set("fullThumbMode", value)
    return a
}

/**
 * fontStyle
 */
func (a *Images) FontStyle(value interface{}) *Images {
    a.Set("fontStyle", value)
    return a
}

/**
 * className
 */
func (a *Images) ClassName(value interface{}) *Images {
    a.Set("className", value)
    return a
}

/**
 * visible
 */
func (a *Images) Visible(value interface{}) *Images {
    a.Set("visible", value)
    return a
}

/**
 * thumbRatio
 */
func (a *Images) ThumbRatio(value interface{}) *Images {
    a.Set("thumbRatio", value)
    return a
}

/**
 * originalSrc
 */
func (a *Images) OriginalSrc(value interface{}) *Images {
    a.Set("originalSrc", value)
    return a
}

/**
 * imageGallaryClassName
 */
func (a *Images) ImageGallaryClassName(value interface{}) *Images {
    a.Set("imageGallaryClassName", value)
    return a
}

/**
 * displayMode
 */
func (a *Images) DisplayMode(value interface{}) *Images {
    a.Set("displayMode", value)
    return a
}

/**
 * sortType
 */
func (a *Images) SortType(value interface{}) *Images {
    a.Set("sortType", value)
    return a
}

/**
 * hoverMode
 */
func (a *Images) HoverMode(value interface{}) *Images {
    a.Set("hoverMode", value)
    return a
}

/**
 * delimiter
 */
func (a *Images) Delimiter(value interface{}) *Images {
    a.Set("delimiter", value)
    return a
}

/**
 * disabledOn
 */
func (a *Images) DisabledOn(value interface{}) *Images {
    a.Set("disabledOn", value)
    return a
}

/**
 * visibleOn
 */
func (a *Images) VisibleOn(value interface{}) *Images {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticSchema
 */
func (a *Images) StaticSchema(value interface{}) *Images {
    a.Set("staticSchema", value)
    return a
}

/**
 * 可选值: images | static-images
 */
func (a *Images) Type(value interface{}) *Images {
    a.Set("type", value)
    return a
}

/**
 * testid
 */
func (a *Images) Testid(value interface{}) *Images {
    a.Set("testid", value)
    return a
}

/**
 * defaultImage
 */
func (a *Images) DefaultImage(value interface{}) *Images {
    a.Set("defaultImage", value)
    return a
}

/**
 * showToolbar
 */
func (a *Images) ShowToolbar(value interface{}) *Images {
    a.Set("showToolbar", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Images) UseMobileUI(value interface{}) *Images {
    a.Set("useMobileUI", value)
    return a
}

/**
 * value
 */
func (a *Images) Value(value interface{}) *Images {
    a.Set("value", value)
    return a
}

/**
 * maskColor
 */
func (a *Images) MaskColor(value interface{}) *Images {
    a.Set("maskColor", value)
    return a
}

/**
 * hidden
 */
func (a *Images) Hidden(value interface{}) *Images {
    a.Set("hidden", value)
    return a
}

/**
 * staticOn
 */
func (a *Images) StaticOn(value interface{}) *Images {
    a.Set("staticOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Images) StaticPlaceholder(value interface{}) *Images {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * style
 */
func (a *Images) Style(value interface{}) *Images {
    a.Set("style", value)
    return a
}

/**
 * src
 */
func (a *Images) Src(value interface{}) *Images {
    a.Set("src", value)
    return a
}

/**
 * enlargetWithImages
 */
func (a *Images) EnlargetWithImages(value interface{}) *Images {
    a.Set("enlargetWithImages", value)
    return a
}
