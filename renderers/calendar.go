package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type Calendar struct {
	*BaseRenderer
}

func NewCalendar() *Calendar {
    a := &Calendar{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "calendar")
    return a
}


func (a *Calendar) Set(name string, value interface{}) *Calendar {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 容器 css 类名
 */
func (a *Calendar) ClassName(value interface{}) *Calendar {
    a.Set("className", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Calendar) VisibleOn(value interface{}) *Calendar {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Calendar) StaticOn(value interface{}) *Calendar {
    a.Set("staticOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Calendar) EditorSetting(value interface{}) *Calendar {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *Calendar) TestIdBuilder(value interface{}) *Calendar {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Calendar) Hidden(value interface{}) *Calendar {
    a.Set("hidden", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Calendar) OnEvent(value interface{}) *Calendar {
    a.Set("onEvent", value)
    return a
}

/**
 * 日程
 */
func (a *Calendar) Schedules(value interface{}) *Calendar {
    a.Set("schedules", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Calendar) StaticLabelClassName(value interface{}) *Calendar {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *Calendar) Style(value interface{}) *Calendar {
    a.Set("style", value)
    return a
}

/**
 * 指定为日历选择控件
 */
func (a *Calendar) Type(value interface{}) *Calendar {
    a.Set("type", value)
    return a
}

/**
 * 日程显示颜色自定义
 */
func (a *Calendar) ScheduleClassNames(value interface{}) *Calendar {
    a.Set("scheduleClassNames", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Calendar) DisabledOn(value interface{}) *Calendar {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否开启放大模式
 */
func (a *Calendar) LargeMode(value interface{}) *Calendar {
    a.Set("largeMode", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Calendar) HiddenOn(value interface{}) *Calendar {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Calendar) Id(value interface{}) *Calendar {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Calendar) StaticClassName(value interface{}) *Calendar {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *Calendar) StaticSchema(value interface{}) *Calendar {
    a.Set("staticSchema", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Calendar) UseMobileUI(value interface{}) *Calendar {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Calendar) Static(value interface{}) *Calendar {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Calendar) StaticInputClassName(value interface{}) *Calendar {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *Calendar) Testid(value interface{}) *Calendar {
    a.Set("testid", value)
    return a
}

/**
 * 日程点击展示
 */
func (a *Calendar) ScheduleAction(value interface{}) *Calendar {
    a.Set("scheduleAction", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Calendar) Disabled(value interface{}) *Calendar {
    a.Set("disabled", value)
    return a
}

/**
 * 今日激活时的自定义样式
 */
func (a *Calendar) TodayActiveStyle(value interface{}) *Calendar {
    a.Set("todayActiveStyle", value)
    return a
}

/**
 * 是否显示
 */
func (a *Calendar) Visible(value interface{}) *Calendar {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Calendar) StaticPlaceholder(value interface{}) *Calendar {
    a.Set("staticPlaceholder", value)
    return a
}
