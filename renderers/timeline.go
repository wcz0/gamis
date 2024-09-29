package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type Timeline struct {
	*BaseRenderer
}

func NewTimeline() *Timeline {
    a := &Timeline{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "timeline")
    return a
}


func (a *Timeline) Set(name string, value interface{}) *Timeline {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 节点时间的CSS类名
 */
func (a *Timeline) TimeClassName(value interface{}) *Timeline {
    a.Set("timeClassName", value)
    return a
}

/**
 * 节点标题的CSS类名
 */
func (a *Timeline) TitleClassName(value interface{}) *Timeline {
    a.Set("titleClassName", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Timeline) StaticOn(value interface{}) *Timeline {
    a.Set("staticOn", value)
    return a
}

/**
 * 指定为 Timeline 时间轴渲染器
 */
func (a *Timeline) Type(value interface{}) *Timeline {
    a.Set("type", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Timeline) StaticClassName(value interface{}) *Timeline {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *Timeline) StaticSchema(value interface{}) *Timeline {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *Timeline) TestIdBuilder(value interface{}) *Timeline {
    a.Set("testIdBuilder", value)
    return a
}

/**
 */
func (a *Timeline) Testid(value interface{}) *Timeline {
    a.Set("testid", value)
    return a
}

/**
 * 节点数据
 */
func (a *Timeline) Items(value interface{}) *Timeline {
    a.Set("items", value)
    return a
}

/**
 * 节点title自定一展示模板
 */
func (a *Timeline) ItemTitleSchema(value interface{}) *Timeline {
    a.Set("itemTitleSchema", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Timeline) ClassName(value interface{}) *Timeline {
    a.Set("className", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Timeline) StaticPlaceholder(value interface{}) *Timeline {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Timeline) UseMobileUI(value interface{}) *Timeline {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 展示方向
 * 可选值: horizontal | vertical
 */
func (a *Timeline) Direction(value interface{}) *Timeline {
    a.Set("direction", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Timeline) Disabled(value interface{}) *Timeline {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Timeline) Hidden(value interface{}) *Timeline {
    a.Set("hidden", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Timeline) HiddenOn(value interface{}) *Timeline {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Timeline) Id(value interface{}) *Timeline {
    a.Set("id", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Timeline) Static(value interface{}) *Timeline {
    a.Set("static", value)
    return a
}

/**
 * 文字相对于时间轴展示方向
 * 可选值: left | right | alternate
 */
func (a *Timeline) Mode(value interface{}) *Timeline {
    a.Set("mode", value)
    return a
}

/**
 * 图标的CSS类名
 */
func (a *Timeline) IconClassName(value interface{}) *Timeline {
    a.Set("iconClassName", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Timeline) DisabledOn(value interface{}) *Timeline {
    a.Set("disabledOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Timeline) OnEvent(value interface{}) *Timeline {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Timeline) StaticInputClassName(value interface{}) *Timeline {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Timeline) EditorSetting(value interface{}) *Timeline {
    a.Set("editorSetting", value)
    return a
}

/**
 * 节点倒序
 */
func (a *Timeline) Reverse(value interface{}) *Timeline {
    a.Set("reverse", value)
    return a
}

/**
 * 节点详情的CSS类名
 */
func (a *Timeline) DetailClassName(value interface{}) *Timeline {
    a.Set("detailClassName", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Timeline) VisibleOn(value interface{}) *Timeline {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Timeline) StaticLabelClassName(value interface{}) *Timeline {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * API 或 数据映射
 */
func (a *Timeline) Source(value interface{}) *Timeline {
    a.Set("source", value)
    return a
}

/**
 * 是否显示
 */
func (a *Timeline) Visible(value interface{}) *Timeline {
    a.Set("visible", value)
    return a
}

/**
 * 组件样式
 */
func (a *Timeline) Style(value interface{}) *Timeline {
    a.Set("style", value)
    return a
}
