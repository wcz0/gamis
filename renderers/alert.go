// <?php

// namespace Slowlyo\AmisRenderers;

// /**
//  * Alert 提示渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/alert
//  *
//  * @author slowlyo
//  * @version 6.2.2
//  */
// class Alert extends BaseRenderer
// {
//     public function __construct()
//     {
//         $this->set('type', 'alert');


//     }

//     /**
//      * 操作区域
//      */
//     public function actions($value = '')
//     {
//         return $this->set('actions', $value);
//     }

//     /**
//      * 内容区域
//      */
//     public function body($value = '')
//     {
//         return $this->set('body', $value);
//     }

//     /**
//      * 容器 css 类名
//      */
//     public function className($value = '')
//     {
//         return $this->set('className', $value);
//     }

//     /**
//      * 关闭按钮CSS类名
//      */
//     public function closeButtonClassName($value = '')
//     {
//         return $this->set('closeButtonClassName', $value);
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
//      * 左侧图标
//      */
//     public function icon($value = '')
//     {
//         return $this->set('icon', $value);
//     }

//     /**
//      * 图标CSS类名
//      */
//     public function iconClassName($value = '')
//     {
//         return $this->set('iconClassName', $value);
//     }

//     /**
//      * 组件唯一 id，主要用于日志采集
//      */
//     public function id($value = '')
//     {
//         return $this->set('id', $value);
//     }

//     /**
//      * 提示类型 可选值: info | warning | success | danger
//      */
//     public function level($value = '')
//     {
//         return $this->set('level', $value);
//     }

//     /**
//      * 事件动作配置
//      */
//     public function onEvent($value = '')
//     {
//         return $this->set('onEvent', $value);
//     }

//     /**
//      * 是否显示关闭按钮
//      */
//     public function showCloseButton($value = true)
//     {
//         return $this->set('showCloseButton', $value);
//     }

//     /**
//      * 是否显示ICON
//      */
//     public function showIcon($value = true)
//     {
//         return $this->set('showIcon', $value);
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
//      * 提示框标题
//      */
//     public function title($value = '')
//     {
//         return $this->set('title', $value);
//     }

//     /**
//      * 指定为提示框类型
//      */
//     public function type($value = 'alert')
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

type Alert struct {
	*BaseRenderer
}

func NewAlert() *Alert{
	a := &Alert{
		BaseRenderer: NewBaseRenderer(),
	}
	a.Set("type", "alert")
	return a
}

/**
 * 操作区域
 */
func (a *Alert) Actions() *Alert{
	a.Set("actions", "")
	return a
}

/**
 * 内容区域
 */
func (a *Alert) Body() *Alert{
	a.Set("body", "")
	return a
}

/**
 * 容器 css 类名
 */
func (a *Alert) ClassName() *Alert{
	a.Set("className", "")
	return a
}

/**
 * 关闭按钮CSS类名
 */
func (a *Alert) CloseButtonClassName() *Alert{
	a.Set("closeButtonClassName", "")
	return a
}

/**
 * 是否禁用
 */
func (a *Alert) Disabled() *Alert{
	a.Set("disabled", true)
	return a
}

/**
 * 是否禁用表达式
 */
func (a *Alert) DisabledOn() *Alert{
	a.Set("disabledOn", "")
	return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Alert) EditorSetting() *Alert{
	a.Set("editorSetting", "")
	return a
}

/**
 * 是否隐藏
 */
func (a *Alert) Hidden() *Alert{
	a.Set("hidden", true)
	return a
}

/**
 * 是否隐藏表达式
 */
func (a *Alert) HiddenOn() *Alert{
	a.Set("hiddenOn", "")
	return a
}

/**
 * 左侧图标
 */
func (a *Alert) Icon() *Alert{
	a.Set("icon", "")
	return a
}

/**
 * 图标CSS类名
 */
func (a *Alert) IconClassName() *Alert{
	a.Set("iconClassName", "")
	return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Alert) Id() *Alert{
	a.Set("id", "")
	return a
}

/**
 * 提示类型 可选值: info | warning | success | danger
 */
func (a *Alert) Level() *Alert{
	a.Set("level", "")
	return a
}

/**
 * 事件动作配置
 */
func (a *Alert) OnEvent() *Alert{
	a.Set("onEvent", "")
	return a
}

/**
 * 是否显示关闭按钮
 */
func (a *Alert) ShowCloseButton() *Alert{
	a.Set("showCloseButton", true)
	return a
}

/**
 * 是否显示ICON
 */
func (a *Alert) ShowIcon() *Alert{
	a.Set("showIcon", true)
	return a
}

/**
 * 是否静态展示
 */
func (a *Alert) Static() *Alert{
	a.Set("static", true)
	return a
}

/**
 * 静态展示表单项类名
 */
func (a *Alert) StaticClassName() *Alert{
	a.Set("staticClassName", "")
	return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Alert) StaticInputClassName() *Alert{
	a.Set("staticInputClassName", "")
	return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Alert) StaticLabelClassName() *Alert{
	a.Set("staticLabelClassName", "")
	return a
}

/**
 * 是否静态展示表达式
 */
func (a *Alert) StaticOn() *Alert{
	a.Set("staticOn", "")
	return a
}

/**
 * 静态展示空值占位
 */
func (a *Alert) StaticPlaceholder() *Alert{
	a.Set("staticPlaceholder", "")
	return a
}

/**
 *
 */
func (a *Alert) StaticSchema() *Alert{
	a.Set("staticSchema", "")
	return a
}

/**
 * 组件样式
 */
func (a *Alert) Style() *Alert{
	a.Set("style", "")
	return a
}

/**
 * 提示框标题
 */
func (a *Alert) Title() *Alert{
	a.Set("title", "")
	return a
}

/**
 * 指定为提示框类型
 */
func (a *Alert) Type() *Alert{
	a.Set("type", "alert")
	return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Alert) UseMobileUI() *Alert{
	a.Set("useMobileUI", true)
	return a
}

/**
 * 是否显示
 */
func (a *Alert) Visible() *Alert{
	a.Set("visible", true)
	return a
}

/**
 * 是否显示表达式
 */
func (a *Alert) VisibleOn() *Alert{
	a.Set("visibleOn", "")
	return a
}