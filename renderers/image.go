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
/**
 * 图片标题
 */
func (a *Image) Title(value string) *Image {
    a.Set("title", value)
    return a
}

/**
 * 宽度
 */
func (a *Image) Width(value string) *Image {
    a.Set("width", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Image) Disabled(value string) *Image {
    a.Set("disabled", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Image) DisabledOn(value string) *Image {
    a.Set("disabledOn", value)
    return a
}

/**
 */
func (a *Image) StaticSchema(value string) *Image {
    a.Set("staticSchema", value)
    return a
}

/**
 * 高度
 */
func (a *Image) Height(value string) *Image {
    a.Set("height", value)
    return a
}

/**
 * 链接的 target
 */
func (a *Image) HtmlTarget(value string) *Image {
    a.Set("htmlTarget", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Image) StaticOn(value string) *Image {
    a.Set("staticOn", value)
    return a
}

/**
 * 组件样式
 */
func (a *Image) Style(value string) *Image {
    a.Set("style", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Image) EditorSetting(value string) *Image {
    a.Set("editorSetting", value)
    return a
}

/**
 * 大图地址，不设置用 src
 */
func (a *Image) OriginalSrc(value string) *Image {
    a.Set("originalSrc", value)
    return a
}

/**
 * 是否新窗口打开
 */
func (a *Image) Blank(value string) *Image {
    a.Set("blank", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Image) VisibleOn(value string) *Image {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Image) StaticLabelClassName(value string) *Image {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 图片说明文字
 */
func (a *Image) Caption(value string) *Image {
    a.Set("caption", value)
    return a
}

/**
 * 预览图模式
 * 可选值: w-full | h-full | contain | cover
 */
func (a *Image) ThumbMode(value string) *Image {
    a.Set("thumbMode", value)
    return a
}

/**
 * 图片描述信息
 */
func (a *Image) ImageCaption(value string) *Image {
    a.Set("imageCaption", value)
    return a
}

/**
 * 组件内层 css 类名
 */
func (a *Image) InnerClassName(value string) *Image {
    a.Set("innerClassName", value)
    return a
}

/**
 * 预览图比率
 * 可选值: 1:1 | 4:3 | 16:9
 */
func (a *Image) ThumbRatio(value string) *Image {
    a.Set("thumbRatio", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Image) Hidden(value string) *Image {
    a.Set("hidden", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Image) OnEvent(value string) *Image {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Image) StaticClassName(value string) *Image {
    a.Set("staticClassName", value)
    return a
}

/**
 * 关联字段名，也可以直接配置 src
 */
func (a *Image) Name(value string) *Image {
    a.Set("name", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Image) StaticInputClassName(value string) *Image {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 链接地址
 */
func (a *Image) Href(value string) *Image {
    a.Set("href", value)
    return a
}

/**
 * 指定为图片展示类型
 * 可选值: image | static-image
 */
func (a *Image) Type(value string) *Image {
    a.Set("type", value)
    return a
}

/**
 * 图片地址，如果配置了 name，这个属性不用配置。
 */
func (a *Image) Src(value string) *Image {
    a.Set("src", value)
    return a
}

/**
 * 图片缩略图外层 css 类名
 */
func (a *Image) ThumbClassName(value string) *Image {
    a.Set("thumbClassName", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Image) StaticPlaceholder(value string) *Image {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Image) UseMobileUI(value string) *Image {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 默认图片地址
 */
func (a *Image) DefaultImage(value string) *Image {
    a.Set("defaultImage", value)
    return a
}

/**
 * 图片无法显示时的替换文本
 */
func (a *Image) Alt(value string) *Image {
    a.Set("alt", value)
    return a
}

/**
 * 外层 css 类名
 */
func (a *Image) ClassName(value string) *Image {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Image) HiddenOn(value string) *Image {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *Image) Visible(value string) *Image {
    a.Set("visible", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Image) Id(value string) *Image {
    a.Set("id", value)
    return a
}

/**
 * 图片展示模式，默认为缩略图模式、可以配置成原图模式
 * 可选值: thumb | original
 */
func (a *Image) ImageMode(value string) *Image {
    a.Set("imageMode", value)
    return a
}

/**
 * 是否展示图片工具栏
 */
func (a *Image) ShowToolbar(value string) *Image {
    a.Set("showToolbar", value)
    return a
}

/**
 * 工具栏配置
 */
func (a *Image) ToolbarActions(value string) *Image {
    a.Set("toolbarActions", value)
    return a
}

/**
 * 放大详情图 CSS 类名
 */
func (a *Image) ImageGallaryClassName(value string) *Image {
    a.Set("imageGallaryClassName", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Image) Static(value string) *Image {
    a.Set("static", value)
    return a
}

/**
 * 是否启动放大功能。
 */
func (a *Image) EnlargeAble(value string) *Image {
    a.Set("enlargeAble", value)
    return a
}

/**
 * 放大时是否显示图片集
 */
func (a *Image) EnlargeWithGallary(value string) *Image {
    a.Set("enlargeWithGallary", value)
    return a
}

/**
 * 图片 css 类名
 */
func (a *Image) ImageClassName(value string) *Image {
    a.Set("imageClassName", value)
    return a
}
