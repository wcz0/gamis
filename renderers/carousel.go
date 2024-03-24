package renderers


/**
 * Carousel 轮播图渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/carousel
 *

*/
type Carousel struct {
	*BaseRenderer
}

func NewCarousel() *Carousel {
    a := &Carousel{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "carousel")
    return a
}
/**
 * 静态展示空值占位
 */
func (a *Carousel) StaticPlaceholder(value string) *Carousel {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 配置单条呈现模板
 */
func (a *Carousel) ItemSchema(value string) *Carousel {
    a.Set("itemSchema", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Carousel) ClassName(value string) *Carousel {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Carousel) DisabledOn(value string) *Carousel {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Carousel) HiddenOn(value string) *Carousel {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 设置宽度
 */
func (a *Carousel) Width(value string) *Carousel {
    a.Set("width", value)
    return a
}

/**
 * 占位
 */
func (a *Carousel) Placeholder(value string) *Carousel {
    a.Set("placeholder", value)
    return a
}

/**
 * 动画类型
 * 可选值: fade | slide
 */
func (a *Carousel) Animation(value string) *Carousel {
    a.Set("animation", value)
    return a
}

/**
 * 是否显示
 */
func (a *Carousel) Visible(value string) *Carousel {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Carousel) StaticInputClassName(value string) *Carousel {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 自定义箭头图标
 */
func (a *Carousel) Icons(value string) *Carousel {
    a.Set("icons", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Carousel) VisibleOn(value string) *Carousel {
    a.Set("visibleOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Carousel) OnEvent(value string) *Carousel {
    a.Set("onEvent", value)
    return a
}

/**
 * 配置控件内容
 */
func (a *Carousel) Controls(value string) *Carousel {
    a.Set("controls", value)
    return a
}

/**
 * 是否自动播放
 */
func (a *Carousel) Auto(value string) *Carousel {
    a.Set("auto", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Carousel) StaticOn(value string) *Carousel {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *Carousel) StaticSchema(value string) *Carousel {
    a.Set("staticSchema", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Carousel) EditorSetting(value string) *Carousel {
    a.Set("editorSetting", value)
    return a
}

/**
 * 预览图模式
 * 可选值: contain | cover
 */
func (a *Carousel) ThumbMode(value string) *Carousel {
    a.Set("thumbMode", value)
    return a
}

/**
 * 多图模式配置项
 */
func (a *Carousel) Multiple(value string) *Carousel {
    a.Set("multiple", value)
    return a
}

/**
 * 指定为轮播图类型
 */
func (a *Carousel) Type(value string) *Carousel {
    a.Set("type", value)
    return a
}

/**
 * 轮播间隔时间
 */
func (a *Carousel) Interval(value string) *Carousel {
    a.Set("interval", value)
    return a
}

/**
 */
func (a *Carousel) Name(value string) *Carousel {
    a.Set("name", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Carousel) Disabled(value string) *Carousel {
    a.Set("disabled", value)
    return a
}

/**
 * 设置高度
 */
func (a *Carousel) Height(value string) *Carousel {
    a.Set("height", value)
    return a
}

/**
 * 可选值: light | dark
 */
func (a *Carousel) ControlsTheme(value string) *Carousel {
    a.Set("controlsTheme", value)
    return a
}

/**
 * 配置固定值
 */
func (a *Carousel) Options(value string) *Carousel {
    a.Set("options", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Carousel) Static(value string) *Carousel {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Carousel) StaticLabelClassName(value string) *Carousel {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *Carousel) Style(value string) *Carousel {
    a.Set("style", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Carousel) UseMobileUI(value string) *Carousel {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 动画时长
 */
func (a *Carousel) Duration(value string) *Carousel {
    a.Set("duration", value)
    return a
}

/**
 * 是否一直显示箭头
 */
func (a *Carousel) AlwaysShowArrow(value string) *Carousel {
    a.Set("alwaysShowArrow", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Carousel) Hidden(value string) *Carousel {
    a.Set("hidden", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Carousel) Id(value string) *Carousel {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Carousel) StaticClassName(value string) *Carousel {
    a.Set("staticClassName", value)
    return a
}
