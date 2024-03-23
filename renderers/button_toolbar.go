// <?php

// namespace Slowlyo\AmisRenderers;

// /**
//  * Button Toolar 渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/button-toolbar
//  *
//  * @author slowlyo
//  * @version 6.2.2
//  */
// class ButtonToolbar extends BaseRenderer
// {
//     public function __construct()
//     {
//         $this->set('type', 'button-toolbar');


//     }

//     /**
//      *
//      */
//     public function buttons($value = '')
//     {
//         return $this->set('buttons', $value);
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
//      * 事件动作配置
//      */
//     public function onEvent($value = '')
//     {
//         return $this->set('onEvent', $value);
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
//      * 指定为按钮工具集合类型
//      */
//     public function type($value = 'button-toolbar')
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

type ButtonToolbar struct {
	*BaseRenderer
}

func NewButtonToolbar() *ButtonToolbar {
	b := &ButtonToolbar{
		BaseRenderer: NewBaseRenderer(),
	}
	b.Set("type", "button-toolbar")
	return b
}

/**
 *
 */
func (b *ButtonToolbar) Buttons(value string) *ButtonToolbar {
	b.Set("buttons", value)
	return b
}

/**
 * 容器 css 类名
 */
func (b *ButtonToolbar) ClassName(value string) *ButtonToolbar {
	b.Set("className", value)
	return b
}

/**
 * 是否禁用
 */
func (b *ButtonToolbar) Disabled(value bool) *ButtonToolbar {
	b.Set("disabled", value)
	return b
}

/**
 * 是否禁用表达式
 */
func (b *ButtonToolbar) DisabledOn(value string) *ButtonToolbar {
	b.Set("disabledOn", value)
	return b
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (b *ButtonToolbar) EditorSetting(value string) *ButtonToolbar {
	b.Set("editorSetting", value)
	return b
}

/**
 * 是否隐藏
 */
func (b *ButtonToolbar) Hidden(value bool) *ButtonToolbar {
	b.Set("hidden", value)
	return b
}

/**
 * 是否隐藏表达式
 */
func (b *ButtonToolbar) HiddenOn(value string) *ButtonToolbar {
	b.Set("hiddenOn", value)
	return b
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (b *ButtonToolbar) Id(value string) *ButtonToolbar {
	b.Set("id", value)
	return b
}

/**
 * 事件动作配置
 */
func (b *ButtonToolbar) OnEvent(value string) *ButtonToolbar {
	b.Set("onEvent", value)
	return b
}

/**
 * 是否静态展示
 */
func (b *ButtonToolbar) Static(value bool) *ButtonToolbar {
	b.Set("static", value)
	return b
}

/**
 * 静态展示表单项类名
 */
func (b *ButtonToolbar) StaticClassName(value string) *ButtonToolbar {
	b.Set("staticClassName", value)
	return b
}

/**
 * 静态展示表单项Value类名
 */
func (b *ButtonToolbar) StaticInputClassName(value string) *ButtonToolbar {
	b.Set("staticInputClassName", value)
	return b
}

/**
 * 静态展示表单项Label类名
 */
func (b *ButtonToolbar) StaticLabelClassName(value string) *ButtonToolbar {
	b.Set("staticLabelClassName", value)
	return b
}

/**
 * 是否静态展示表达式
 */
func (b *ButtonToolbar) StaticOn(value string) *ButtonToolbar {
	b.Set("staticOn", value)
	return b
}

/**
 * 静态展示空值占位
 */
func (b *ButtonToolbar) StaticPlaceholder(value string) *ButtonToolbar {
	b.Set("staticPlaceholder", value)
	return b
}

/**
 *
 */
func (b *ButtonToolbar) StaticSchema(value string) *ButtonToolbar {
	b.Set("staticSchema", value)
	return b
}

/**
 * 组件样式
 */
func (b *ButtonToolbar) Style(value string) *ButtonToolbar {
	b.Set("style", value)
	return b
}

/**
 * 指定为按钮工具集合类型
 */
func (b *ButtonToolbar) Type(value string) *ButtonToolbar {
	b.Set("type", value)
	return b
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (b *ButtonToolbar) UseMobileUI(value bool) *ButtonToolbar {
	b.Set("useMobileUI", value)
	return b
}

/**
 * 是否显示
 */
func (b *ButtonToolbar) Visible(value bool) *ButtonToolbar {
	b.Set("visible", value)
	return b
}

/**
 * 是否显示表达式
 */
func (b *ButtonToolbar) VisibleOn(value string) *ButtonToolbar {
	b.Set("visibleOn", value)
	return b
}