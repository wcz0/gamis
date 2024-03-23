// <?php

// namespace Slowlyo\AmisRenderers;

// /**
//  * Audio 音频渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/audio
//  *
//  * @author slowlyo
//  * @version 6.2.2
//  */
// class Audio extends BaseRenderer
// {
//     public function __construct()
//     {
//         $this->set('type', 'audio');


//     }

//     /**
//      * 是否自动播放
//      */
//     public function autoPlay($value = true)
//     {
//         return $this->set('autoPlay', $value);
//     }

//     /**
//      * 容器 css 类名
//      */
//     public function className($value = '')
//     {
//         return $this->set('className', $value);
//     }

//     /**
//      * 可以配置控制器
//      */
//     public function controls($value = '')
//     {
//         return $this->set('controls', $value);
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
//      * 是否是内联模式
//      */
//     public function inline($value = true)
//     {
//         return $this->set('inline', $value);
//     }

//     /**
//      * 是否循环播放
//      */
//     public function loop($value = true)
//     {
//         return $this->set('loop', $value);
//     }

//     /**
//      * 事件动作配置
//      */
//     public function onEvent($value = '')
//     {
//         return $this->set('onEvent', $value);
//     }

//     /**
//      * 配置可选播放倍速
//      */
//     public function rates($value = '')
//     {
//         return $this->set('rates', $value);
//     }

//     /**
//      * "视频播放地址, 支持 $ 取变量。
//      */
//     public function src($value = '')
//     {
//         return $this->set('src', $value);
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
//      * 指定为音频播放器
//      */
//     public function type($value = 'audio')
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

type Audio struct {
	*BaseRenderer
}

func NewAudio() *Audio{
	a := &Audio{
		BaseRenderer: NewBaseRenderer(),
	}
	a.Set("type", "audio")
	return a
}

/**
 * 是否自动播放
 */
func (a *Audio) AutoPlay(value bool) *Audio {
	a.Set("autoPlay", value)
	return a
}

/**
 * 容器 css 类名
 */
func (a *Audio) ClassName(value string) *Audio {
	a.Set("className", value)
	return a
}

/**
 * 可以配置控制器
 */
func (a *Audio) Controls(value string) *Audio {
	a.Set("controls", value)
	return a
}

/**
 * 是否禁用
 */
func (a *Audio) Disabled(value bool) *Audio {
	a.Set("disabled", value)
	return a
}

/**
 * 是否禁用表达式
 */
func (a *Audio) DisabledOn(value string) *Audio {
	a.Set("disabledOn", value)
	return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Audio) EditorSetting(value string) *Audio {
	a.Set("editorSetting", value)
	return a
}

/**
 * 是否隐藏
 */
func (a *Audio) Hidden(value bool) *Audio {
	a.Set("hidden", value)
	return a
}

/**
 * 是否隐藏表达式
 */
func (a *Audio) HiddenOn(value string) *Audio {
	a.Set("hiddenOn", value)
	return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Audio) Id(value string) *Audio {
	a.Set("id", value)
	return a
}

/**
 * 是否是内联模式
 */
func (a *Audio) Inline(value bool) *Audio {
	a.Set("inline", value)
	return a
}

/**
 * 是否循环播放
 */
func (a *Audio) Loop(value bool) *Audio {
	a.Set("loop", value)
	return a
}

/**
 * 事件动作配置
 */
func (a *Audio) OnEvent(value string) *Audio {
	a.Set("onEvent", value)
	return a
}

/**
 * 配置可选播放倍速
 */
func (a *Audio) Rates(value string) *Audio {
	a.Set("rates", value)
	return a
}

/**
 * "视频播放地址, 支持 $ 取变量。
 */
func (a *Audio) Src(value string) *Audio {
	a.Set("src", value)
	return a
}

/**
 * 是否静态展示
 */
func (a *Audio) Static(value bool) *Audio {
	a.Set("static", value)
	return a
}

/**
 * 静态展示表单项类名
 */
func (a *Audio) StaticClassName(value string) *Audio {
	a.Set("staticClassName", value)
	return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Audio) StaticInputClassName(value string) *Audio {
	a.Set("staticInputClassName", value)
	return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Audio) StaticLabelClassName(value string) *Audio {
	a.Set("staticLabelClassName", value)
	return a
}

/**
 * 是否静态展示表达式
 */
func (a *Audio) StaticOn(value string) *Audio {
	a.Set("staticOn", value)
	return a
}

/**
 * 静态展示空值占位
 */
func (a *Audio) StaticPlaceholder(value string) *Audio {
	a.Set("staticPlaceholder", value)
	return a
}

/**
 *
 */
func (a *Audio) StaticSchema(value string) *Audio {
	a.Set("staticSchema", value)
	return a
}

/**
 * 组件样式
 */
func (a *Audio) Style(value string) *Audio {
	a.Set("style", value)
	return a
}

/**
 * 指定为音频播放器
 */
func (a *Audio) Type(value string) *Audio {
	a.Set("type", value)
	return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Audio) UseMobileUI(value bool) *Audio {
	a.Set("useMobileUI", value)
	return a
}

/**
 * 是否显示
 */
func (a *Audio) Visible(value bool) *Audio {
	a.Set("visible", value)
	return a
}

/**
 * 是否显示表达式
 */
func (a *Audio) VisibleOn(value string) *Audio {
	a.Set("visibleOn", value)
	return a
}