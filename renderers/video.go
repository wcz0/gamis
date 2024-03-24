package renderers


/**
 * 视频播放器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/video

 * @author wcz0
 * @version 6.2.2
 */
type Video struct {
	*BaseRenderer
}

func NewVideo() *Video {
    a := &Video{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "video")
    return a
}
/**
 * 指定为视频类型
 */
func (a *Video) Type(value interface{}) *Video {
    a.Set("type", value)
    return a
}

/**
 * 设置后，可以显示切帧.点击帧的时候会将视频跳到对应时间。frames: {  '01:22': 'http://domain/xxx.jpg' }
 */
func (a *Video) Frames(value interface{}) *Video {
    a.Set("frames", value)
    return a
}

/**
 * 是否将视频和封面分开显示
 */
func (a *Video) SplitPoster(value interface{}) *Video {
    a.Set("splitPoster", value)
    return a
}

/**
 * 默认播放的时候到了下一帧会继续播放，同时高亮下一帧。 如果配置这个则会停止播放，等待用户选择新的区间再播放。
 */
func (a *Video) StopOnNextFrame(value interface{}) *Video {
    a.Set("stopOnNextFrame", value)
    return a
}

/**
 */
func (a *Video) StaticSchema(value interface{}) *Video {
    a.Set("staticSchema", value)
    return a
}

/**
 * 组件样式
 */
func (a *Video) Style(value interface{}) *Video {
    a.Set("style", value)
    return a
}

/**
 * 配置帧列表容器className
 */
func (a *Video) FramesClassName(value interface{}) *Video {
    a.Set("framesClassName", value)
    return a
}

/**
 * 是否循环播放
 */
func (a *Video) Loop(value interface{}) *Video {
    a.Set("loop", value)
    return a
}

/**
 * 视频比率
 * 可选值: auto | 4:3 | 16:9
 */
func (a *Video) AspectRatio(value interface{}) *Video {
    a.Set("aspectRatio", value)
    return a
}

/**
 * 是否显示
 */
func (a *Video) Visible(value interface{}) *Video {
    a.Set("visible", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Video) StaticOn(value interface{}) *Video {
    a.Set("staticOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Video) EditorSetting(value interface{}) *Video {
    a.Set("editorSetting", value)
    return a
}

/**
 * 点击帧画面时是否跳转视频对应的点
 */
func (a *Video) JumpFrame(value interface{}) *Video {
    a.Set("jumpFrame", value)
    return a
}

/**
 * 配置播放器 className
 */
func (a *Video) PlayerClassName(value interface{}) *Video {
    a.Set("playerClassName", value)
    return a
}

/**
 * 视频速率
 */
func (a *Video) Rates(value interface{}) *Video {
    a.Set("rates", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Video) ClassName(value interface{}) *Video {
    a.Set("className", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Video) OnEvent(value interface{}) *Video {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Video) StaticClassName(value interface{}) *Video {
    a.Set("staticClassName", value)
    return a
}

/**
 * 视频封面地址
 */
func (a *Video) Poster(value interface{}) *Video {
    a.Set("poster", value)
    return a
}

/**
 * 跳转到帧时，往前多少秒。
 */
func (a *Video) JumpBufferDuration(value interface{}) *Video {
    a.Set("jumpBufferDuration", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Video) Disabled(value interface{}) *Video {
    a.Set("disabled", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Video) StaticLabelClassName(value interface{}) *Video {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 是否自动播放
 */
func (a *Video) AutoPlay(value interface{}) *Video {
    a.Set("autoPlay", value)
    return a
}

/**
 * 是否初始静音
 */
func (a *Video) Muted(value interface{}) *Video {
    a.Set("muted", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Video) StaticInputClassName(value interface{}) *Video {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 如果显示切帧，通过此配置项可以控制每行显示多少帧
 */
func (a *Video) ColumnsCount(value interface{}) *Video {
    a.Set("columnsCount", value)
    return a
}

/**
 * 视频播放地址
 */
func (a *Video) Src(value interface{}) *Video {
    a.Set("src", value)
    return a
}

/**
 * 视频类型如： video/x-flv
 */
func (a *Video) VideoType(value interface{}) *Video {
    a.Set("videoType", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Video) DisabledOn(value interface{}) *Video {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Video) Hidden(value interface{}) *Video {
    a.Set("hidden", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Video) HiddenOn(value interface{}) *Video {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Video) VisibleOn(value interface{}) *Video {
    a.Set("visibleOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Video) Id(value interface{}) *Video {
    a.Set("id", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Video) Static(value interface{}) *Video {
    a.Set("static", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Video) StaticPlaceholder(value interface{}) *Video {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Video) UseMobileUI(value interface{}) *Video {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 如果是实时的，请标记一下
 */
func (a *Video) IsLive(value interface{}) *Video {
    a.Set("isLive", value)
    return a
}
