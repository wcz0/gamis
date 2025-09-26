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


func (a *Video) Set(name string, value interface{}) *Video {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * visibleOn
 */
func (a *Video) VisibleOn(value interface{}) *Video {
    a.Set("visibleOn", value)
    return a
}

/**
 * testid
 */
func (a *Video) Testid(value interface{}) *Video {
    a.Set("testid", value)
    return a
}

/**
 * muted
 */
func (a *Video) Muted(value interface{}) *Video {
    a.Set("muted", value)
    return a
}

/**
 * visible
 */
func (a *Video) Visible(value interface{}) *Video {
    a.Set("visible", value)
    return a
}

/**
 */
func (a *Video) Type(value interface{}) *Video {
    a.Set("type", value)
    return a
}

/**
 * loop
 */
func (a *Video) Loop(value interface{}) *Video {
    a.Set("loop", value)
    return a
}

/**
 * videoType
 */
func (a *Video) VideoType(value interface{}) *Video {
    a.Set("videoType", value)
    return a
}

/**
 * rates
 */
func (a *Video) Rates(value interface{}) *Video {
    a.Set("rates", value)
    return a
}

/**
 * stopOnNextFrame
 */
func (a *Video) StopOnNextFrame(value interface{}) *Video {
    a.Set("stopOnNextFrame", value)
    return a
}

/**
 * isLive
 */
func (a *Video) IsLive(value interface{}) *Video {
    a.Set("isLive", value)
    return a
}

/**
 * static
 */
func (a *Video) Static(value interface{}) *Video {
    a.Set("static", value)
    return a
}

/**
 * style
 */
func (a *Video) Style(value interface{}) *Video {
    a.Set("style", value)
    return a
}

/**
 * playerClassName
 */
func (a *Video) PlayerClassName(value interface{}) *Video {
    a.Set("playerClassName", value)
    return a
}

/**
 * hidden
 */
func (a *Video) Hidden(value interface{}) *Video {
    a.Set("hidden", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Video) HiddenOn(value interface{}) *Video {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticOn
 */
func (a *Video) StaticOn(value interface{}) *Video {
    a.Set("staticOn", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Video) StaticLabelClassName(value interface{}) *Video {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Video) StaticInputClassName(value interface{}) *Video {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Video) UseMobileUI(value interface{}) *Video {
    a.Set("useMobileUI", value)
    return a
}

/**
 * columnsCount
 */
func (a *Video) ColumnsCount(value interface{}) *Video {
    a.Set("columnsCount", value)
    return a
}

/**
 * frames
 */
func (a *Video) Frames(value interface{}) *Video {
    a.Set("frames", value)
    return a
}

/**
 * id
 */
func (a *Video) Id(value interface{}) *Video {
    a.Set("id", value)
    return a
}

/**
 * framesClassName
 */
func (a *Video) FramesClassName(value interface{}) *Video {
    a.Set("framesClassName", value)
    return a
}

/**
 * jumpFrame
 */
func (a *Video) JumpFrame(value interface{}) *Video {
    a.Set("jumpFrame", value)
    return a
}

/**
 * poster
 */
func (a *Video) Poster(value interface{}) *Video {
    a.Set("poster", value)
    return a
}

/**
 * splitPoster
 */
func (a *Video) SplitPoster(value interface{}) *Video {
    a.Set("splitPoster", value)
    return a
}

/**
 * src
 */
func (a *Video) Src(value interface{}) *Video {
    a.Set("src", value)
    return a
}

/**
 * aspectRatio
 */
func (a *Video) AspectRatio(value interface{}) *Video {
    a.Set("aspectRatio", value)
    return a
}

/**
 * jumpBufferDuration
 */
func (a *Video) JumpBufferDuration(value interface{}) *Video {
    a.Set("jumpBufferDuration", value)
    return a
}

/**
 * className
 */
func (a *Video) ClassName(value interface{}) *Video {
    a.Set("className", value)
    return a
}

/**
 * editorSetting
 */
func (a *Video) EditorSetting(value interface{}) *Video {
    a.Set("editorSetting", value)
    return a
}

/**
 * disabled
 */
func (a *Video) Disabled(value interface{}) *Video {
    a.Set("disabled", value)
    return a
}

/**
 * disabledOn
 */
func (a *Video) DisabledOn(value interface{}) *Video {
    a.Set("disabledOn", value)
    return a
}

/**
 * onEvent
 */
func (a *Video) OnEvent(value interface{}) *Video {
    a.Set("onEvent", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Video) StaticPlaceholder(value interface{}) *Video {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticClassName
 */
func (a *Video) StaticClassName(value interface{}) *Video {
    a.Set("staticClassName", value)
    return a
}

/**
 * autoPlay
 */
func (a *Video) AutoPlay(value interface{}) *Video {
    a.Set("autoPlay", value)
    return a
}

/**
 * staticSchema
 */
func (a *Video) StaticSchema(value interface{}) *Video {
    a.Set("staticSchema", value)
    return a
}
