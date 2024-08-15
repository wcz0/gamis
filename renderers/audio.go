package renderers


/**
 * Audio 音频渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/audio

 * @author wcz0
 * @version 6.2.2
 */
type Audio struct {
	*BaseRenderer
}

func NewAudio() *Audio {
    a := &Audio{
        BaseRenderer: NewBaseRenderer(),
    }
    a.Set("type", "audio")
    return a
}

/**
 * 是否禁用
 */
func (a *Audio) Disabled(value interface{}) *Audio {
    a.Set("disabled", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Audio) OnEvent(value interface{}) *Audio {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否循环播放
 */
func (a *Audio) Loop(value interface{}) *Audio {
    a.Set("loop", value)
    return a
}

/**
 * 可以配置控制器
 */
func (a *Audio) Controls(value interface{}) *Audio {
    a.Set("controls", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Audio) EditorSetting(value interface{}) *Audio {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *Audio) TestIdBuilder(value interface{}) *Audio {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Audio) ClassName(value interface{}) *Audio {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Audio) Hidden(value interface{}) *Audio {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示
 */
func (a *Audio) Visible(value interface{}) *Audio {
    a.Set("visible", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Audio) VisibleOn(value interface{}) *Audio {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Audio) Static(value interface{}) *Audio {
    a.Set("static", value)
    return a
}

/**
 * 组件样式
 */
func (a *Audio) Style(value interface{}) *Audio {
    a.Set("style", value)
    return a
}

/**
 * "视频播放地址, 支持 $ 取变量。
 */
func (a *Audio) Src(value interface{}) *Audio {
    a.Set("src", value)
    return a
}

/**
 * 配置可选播放倍速
 */
func (a *Audio) Rates(value interface{}) *Audio {
    a.Set("rates", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Audio) StaticOn(value interface{}) *Audio {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Audio) StaticInputClassName(value interface{}) *Audio {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *Audio) StaticSchema(value interface{}) *Audio {
    a.Set("staticSchema", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Audio) DisabledOn(value interface{}) *Audio {
    a.Set("disabledOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Audio) Id(value interface{}) *Audio {
    a.Set("id", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Audio) StaticPlaceholder(value interface{}) *Audio {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 是否自动播放
 */
func (a *Audio) AutoPlay(value interface{}) *Audio {
    a.Set("autoPlay", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Audio) HiddenOn(value interface{}) *Audio {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Audio) StaticClassName(value interface{}) *Audio {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *Audio) Testid(value interface{}) *Audio {
    a.Set("testid", value)
    return a
}

/**
 * 是否是内联模式
 */
func (a *Audio) Inline(value interface{}) *Audio {
    a.Set("inline", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Audio) StaticLabelClassName(value interface{}) *Audio {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Audio) UseMobileUI(value interface{}) *Audio {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 指定为音频播放器
 */
func (a *Audio) Type(value interface{}) *Audio {
    a.Set("type", value)
    return a
}
