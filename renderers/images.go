package renderers


/**
 * 图片集展示控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/images
 *

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
/**
 * 是否显示尺寸。
 */
func (a *Images) ShowDimensions(value string) *Images {
    a.Set("showDimensions", value)
    return a
}

/**
 * 工具栏配置
 */
func (a *Images) ToolbarActions(value string) *Images {
    a.Set("toolbarActions", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Images) DisabledOn(value string) *Images {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Images) StaticLabelClassName(value string) *Images {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 指定为图片集渲染器
 * 可选值: images | static-images
 */
func (a *Images) Type(value string) *Images {
    a.Set("type", value)
    return a
}

/**
 * 列表为空时显示
 */
func (a *Images) Placeholder(value string) *Images {
    a.Set("placeholder", value)
    return a
}

/**
 * 关联字段名，也可以直接配置 src
 */
func (a *Images) Name(value string) *Images {
    a.Set("name", value)
    return a
}

/**
 */
func (a *Images) Source(value string) *Images {
    a.Set("source", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Images) StaticInputClassName(value string) *Images {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Images) EditorSetting(value string) *Images {
    a.Set("editorSetting", value)
    return a
}

/**
 * 图片地址，默认读取数据中的 image 属性，如果不是请配置 ,如  ${imageUrl}
 */
func (a *Images) Src(value string) *Images {
    a.Set("src", value)
    return a
}

/**
 * 大图地址，不设置用 src 属性，如果不是请配置，如：${imageOriginUrl}
 */
func (a *Images) OriginalSrc(value string) *Images {
    a.Set("originalSrc", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Images) Disabled(value string) *Images {
    a.Set("disabled", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Images) OnEvent(value string) *Images {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Images) StaticPlaceholder(value string) *Images {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 配置值的连接符
 */
func (a *Images) Delimiter(value string) *Images {
    a.Set("delimiter", value)
    return a
}

/**
 */
func (a *Images) Options(value string) *Images {
    a.Set("options", value)
    return a
}

/**
 * 放大时是否显示图片集
 */
func (a *Images) EnlargetWithImages(value string) *Images {
    a.Set("enlargetWithImages", value)
    return a
}

/**
 * 外层 CSS 类名
 */
func (a *Images) ClassName(value string) *Images {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Images) HiddenOn(value string) *Images {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Images) StaticClassName(value string) *Images {
    a.Set("staticClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *Images) Style(value string) *Images {
    a.Set("style", value)
    return a
}

/**
 */
func (a *Images) StaticSchema(value string) *Images {
    a.Set("staticSchema", value)
    return a
}

/**
 * 预览图模式
 * 可选值: w-full | h-full | contain | cover
 */
func (a *Images) ThumbMode(value string) *Images {
    a.Set("thumbMode", value)
    return a
}

/**
 * 列表 CSS 类名
 */
func (a *Images) ListClassName(value string) *Images {
    a.Set("listClassName", value)
    return a
}

/**
 * 是否展示图片工具栏
 */
func (a *Images) ShowToolbar(value string) *Images {
    a.Set("showToolbar", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Images) Static(value string) *Images {
    a.Set("static", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Images) StaticOn(value string) *Images {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *Images) Value(value string) *Images {
    a.Set("value", value)
    return a
}

/**
 * 是否启动放大功能。
 */
func (a *Images) EnlargeAble(value string) *Images {
    a.Set("enlargeAble", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Images) Hidden(value string) *Images {
    a.Set("hidden", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Images) Id(value string) *Images {
    a.Set("id", value)
    return a
}

/**
 * 默认图片地址
 */
func (a *Images) DefaultImage(value string) *Images {
    a.Set("defaultImage", value)
    return a
}

/**
 * 放大详情图 CSS 类名
 */
func (a *Images) ImageGallaryClassName(value string) *Images {
    a.Set("imageGallaryClassName", value)
    return a
}

/**
 * 是否显示
 */
func (a *Images) Visible(value string) *Images {
    a.Set("visible", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Images) VisibleOn(value string) *Images {
    a.Set("visibleOn", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Images) UseMobileUI(value string) *Images {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 预览图比率
 * 可选值: 1:1 | 4:3 | 16:9
 */
func (a *Images) ThumbRatio(value string) *Images {
    a.Set("thumbRatio", value)
    return a
}
