// <?php

// namespace Slowlyo\AmisRenderers;

// /**
//  * Calendar
//  *
//  * @author slowlyo
//  * @version 6.2.2
//  */
// class Calendar extends BaseRenderer
// {
//     public function __construct()
//     {
//         $this->set('type', 'calendar');


//     }

//     /**
//      * 容器 css 类名
//      */
//     public function className($value = '')
//     {
//         return $this->set('className', $value);
//     }

//     /**
//      * 是否禁用
//      */
//     public function disabled($value = true)
//     {
//         return $this->set('disabled', $value);
//     }

//     /**
//      * 是否禁用表达式
//      */
//     public function disabledOn($value = '')
//     {
//         return $this->set('disabledOn', $value);
//     }

//     /**
//      * 编辑器配置，运行时可以忽略
//      */
//     public function editorSetting($value = '')
//     {
//         return $this->set('editorSetting', $value);
//     }

//     /**
//      * 是否隐藏
//      */
//     public function hidden($value = true)
//     {
//         return $this->set('hidden', $value);
//     }

//     /**
//      * 是否隐藏表达式
//      */
//     public function hiddenOn($value = '')
//     {
//         return $this->set('hiddenOn', $value);
//     }

//     /**
//      * 组件唯一 id，主要用于日志采集
//      */
//     public function id($value = '')
//     {
//         return $this->set('id', $value);
//     }

//     /**
//      * 是否开启放大模式
//      */
//     public function largeMode($value = true)
//     {
//         return $this->set('largeMode', $value);
//     }

//     /**
//      * 事件动作配置
//      */
//     public function onEvent($value = '')
//     {
//         return $this->set('onEvent', $value);
//     }

//     /**
//      * 日程点击展示
//      */
//     public function scheduleAction($value = '')
//     {
//         return $this->set('scheduleAction', $value);
//     }

//     /**
//      * 日程显示颜色自定义
//      */
//     public function scheduleClassNames($value = '')
//     {
//         return $this->set('scheduleClassNames', $value);
//     }

//     /**
//      * 日程
//      */
//     public function schedules($value = '')
//     {
//         return $this->set('schedules', $value);
//     }

//     /**
//      * 是否静态展示
//      */
//     public function static($value = true)
//     {
//         return $this->set('static', $value);
//     }

//     /**
//      * 静态展示表单项类名
//      */
//     public function staticClassName($value = '')
//     {
//         return $this->set('staticClassName', $value);
//     }

//     /**
//      * 静态展示表单项Value类名
//      */
//     public function staticInputClassName($value = '')
//     {
//         return $this->set('staticInputClassName', $value);
//     }

//     /**
//      * 静态展示表单项Label类名
//      */
//     public function staticLabelClassName($value = '')
//     {
//         return $this->set('staticLabelClassName', $value);
//     }

//     /**
//      * 是否静态展示表达式
//      */
//     public function staticOn($value = '')
//     {
//         return $this->set('staticOn', $value);
//     }

//     /**
//      * 静态展示空值占位
//      */
//     public function staticPlaceholder($value = '')
//     {
//         return $this->set('staticPlaceholder', $value);
//     }

//     /**
//      *
//      */
//     public function staticSchema($value = '')
//     {
//         return $this->set('staticSchema', $value);
//     }

//     /**
//      * 组件样式
//      */
//     public function style($value = '')
//     {
//         return $this->set('style', $value);
//     }

//     /**
//      * 今日激活时的自定义样式
//      */
//     public function todayActiveStyle($value = '')
//     {
//         return $this->set('todayActiveStyle', $value);
//     }

//     /**
//      * 指定为日历选择控件
//      */
//     public function type($value = 'calendar')
//     {
//         return $this->set('type', $value);
//     }

//     /**
//      * 可以组件级别用来关闭移动端样式
//      */
//     public function useMobileUI($value = true)
//     {
//         return $this->set('useMobileUI', $value);
//     }

//     /**
//      * 是否显示
//      */
//     public function visible($value = true)
//     {
//         return $this->set('visible', $value);
//     }

//     /**
//      * 是否显示表达式
//      */
//     public function visibleOn($value = '')
//     {
//         return $this->set('visibleOn', $value);
//     }


// }

package renderers

type Calendar struct {
	*BaseRenderer
}

func NewCalendar() *Calendar {
	c := &Calendar{
		BaseRenderer: NewBaseRenderer(),
	}
	c.Set("type", "calendar")
	return c
}

/**
 * 容器 css 类名
 */
func (c *Calendar) ClassName(value string) *Calendar {
	c.Set("className", value)
	return c
}

/**
 * 是否禁用
 */
func (c *Calendar) Disabled(value bool) *Calendar {
	c.Set("disabled", value)
	return c
}

/**
 * 是否禁用表达式
 */
func (c *Calendar) DisabledOn(value string) *Calendar {
	c.Set("disabledOn", value)
	return c
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (c *Calendar) EditorSetting(value string) *Calendar {
	c.Set("editorSetting", value)
	return c
}

/**
 * 是否隐藏
 */
func (c *Calendar) Hidden(value bool) *Calendar {
	c.Set("hidden", value)
	return c
}

/**
 * 是否隐藏表达式
 */
func (c *Calendar) HiddenOn(value string) *Calendar {
	c.Set("hiddenOn", value)
	return c
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (c *Calendar) Id(value string) *Calendar {
	c.Set("id", value)
	return c
}

/**
 * 是否开启放大模式
 */
func (c *Calendar) LargeMode(value bool) *Calendar {
	c.Set("largeMode", value)
	return c
}

/**
 * 事件动作配置
 */
func (c *Calendar) OnEvent(value string) *Calendar {
	c.Set("onEvent", value)
	return c
}

/**
 * 日程点击展示
 */
func (c *Calendar) ScheduleAction(value string) *Calendar {
	c.Set("scheduleAction", value)
	return c
}

/**
 * 日程显示颜色自定义
 */
func (c *Calendar) ScheduleClassNames(value string) *Calendar {
	c.Set("scheduleClassNames", value)
	return c
}

/**
 * 日程
 */
func (c *Calendar) Schedules(value string) *Calendar {
	c.Set("schedules", value)
	return c
}

/**
 * 是否静态展示
 */
func (c *Calendar) Static(value bool) *Calendar {
	c.Set("static", value)
	return c
}

/**
 * 静态展示表单项类名
 */
func (c *Calendar) StaticClassName(value string) *Calendar {
	c.Set("staticClassName", value)
	return c
}

/**
 * 静态展示表单项Value类名
 */
func (c *Calendar) StaticInputClassName(value string) *Calendar {
	c.Set("staticInputClassName", value)
	return c
}

/**
 * 静态展示表单项Label类名
 */
func (c *Calendar) StaticLabelClassName(value string) *Calendar {
	c.Set("staticLabelClassName", value)
	return c
}

/**
 * 是否静态展示表达式
 */
func (c *Calendar) StaticOn(value string) *Calendar {
	c.Set("staticOn", value)
	return c
}

/**
 * 静态展示空值占位
 */
func (c *Calendar) StaticPlaceholder(value string) *Calendar {
	c.Set("staticPlaceholder", value)
	return c
}

/**
 *
 */
func (c *Calendar) StaticSchema(value string) *Calendar {
	c.Set("staticSchema", value)
	return c
}

/**
 * 组件样式
 */
func (c *Calendar) Style(value string) *Calendar {
	c.Set("style", value)
	return c
}

/**
 * 今日激活时的自定义样式
 */
func (c *Calendar) TodayActiveStyle(value string) *Calendar {
	c.Set("todayActiveStyle", value)
	return c
}

/**
 * 指定为日历选择控件
 */
func (c *Calendar) Type(value string) *Calendar {
	c.Set("type", value)
	return c
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (c *Calendar) UseMobileUI(value bool) *Calendar {
	c.Set("useMobileUI", value)
	return c
}

/**
 * 是否显示
 */
func (c *Calendar) Visible(value bool) *Calendar {
	c.Set("visible", value)
	return c
}

/**
 * 是否显示表达式
 */
func (c *Calendar) VisibleOn(value string) *Calendar {
	c.Set("visibleOn", value)
	return c
}
