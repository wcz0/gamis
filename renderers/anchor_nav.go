// <?php

// namespace Slowlyo\AmisRenderers;

// /**
//  * AnchorNav 锚点导航渲染器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/anchor-nav
//  *
//  * @author slowlyo
//  * @version 6.2.2
//  */
// class AnchorNav extends BaseRenderer
// {
//     public function __construct()
//     {
//         $this->set('type', 'anchor-nav');


//     }

//     /**
//      * 被激活（定位）的楼层
//      */
//     public function active($value = '')
//     {
//         return $this->set('active', $value);
//     }

//     /**
//      * 样式名
//      */
//     public function className($value = '')
//     {
//         return $this->set('className', $value);
//     }

//     /**
//      *  可选值: vertical | horizontal
//      */
//     public function direction($value = '')
//     {
//         return $this->set('direction', $value);
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
//      * 导航样式名
//      */
//     public function linkClassName($value = '')
//     {
//         return $this->set('linkClassName', $value);
//     }

//     /**
//      * 楼层集合
//      */
//     public function links($value = '')
//     {
//         return $this->set('links', $value);
//     }

//     /**
//      * 事件动作配置
//      */
//     public function onEvent($value = '')
//     {
//         return $this->set('onEvent', $value);
//     }

//     /**
//      * 楼层样式名
//      */
//     public function sectionClassName($value = '')
//     {
//         return $this->set('sectionClassName', $value);
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
//      * 指定为 AnchorNav 锚点导航渲染器
//      */
//     public function type($value = 'anchor-nav')
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

type AnchorNav struct {
	*BaseRenderer
}

func NewAnchorNav() *AnchorNav{
	a := &AnchorNav{
		BaseRenderer: NewBaseRenderer(),
	}
	a.Set("type", "anchor-nav")
	return a
}

/**
 * 被激活（定位）的楼层
 */
func (a *AnchorNav) Active(value string) *AnchorNav{
	a.Set("active", value)
	return a
}

/**
 * 样式名
 */
func (a *AnchorNav) ClassName(value string) *AnchorNav{
	a.Set("className", value)
	return a
}

/**
 *  可选值: vertical | horizontal
 */
func (a *AnchorNav) Direction(value string) *AnchorNav{
	a.Set("direction", value)
	return a
}

/**
 * 是否禁用
 */
func (a *AnchorNav) Disabled(value bool) *AnchorNav{
	a.Set("disabled", value)
	return a
}

/**
 * 是否禁用表达式
 */
func (a *AnchorNav) DisabledOn(value string) *AnchorNav{
	a.Set("disabledOn", value)
	return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *AnchorNav) EditorSetting(value string) *AnchorNav{
	a.Set("editorSetting", value)
	return a
}

/**
 * 是否隐藏
 */
func (a *AnchorNav) Hidden(value bool) *AnchorNav{
	a.Set("hidden", value)
	return a
}

/**
 * 是否隐藏表达式
 */
func (a *AnchorNav) HiddenOn(value string) *AnchorNav{
	a.Set("hiddenOn", value)
	return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *AnchorNav) Id(value string) *AnchorNav{
	a.Set("id", value)
	return a
}

/**
 * 导航样式名
 */
func (a *AnchorNav) LinkClassName(value string) *AnchorNav{
	a.Set("linkClassName", value)
	return a
}

/**
 * 楼层集合
 */
func (a *AnchorNav) Links(value string) *AnchorNav{
	a.Set("links", value)
	return a
}

/**
 * 事件动作配置
 */
func (a *AnchorNav) OnEvent(value string) *AnchorNav{
	a.Set("onEvent", value)
	return a
}

/**
 * 楼层样式名
 */
func (a *AnchorNav) SectionClassName(value string) *AnchorNav{
	a.Set("sectionClassName", value)
	return a
}

/**
 * 是否静态展示
 */
func (a *AnchorNav) Static(value bool) *AnchorNav{
	a.Set("static", value)
	return a
}

/**
 * 静态展示表单项类名
 */
func (a *AnchorNav) StaticClassName(value string) *AnchorNav{
	a.Set("staticClassName", value)
	return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *AnchorNav) StaticInputClassName(value string) *AnchorNav{
	a.Set("staticInputClassName", value)
	return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *AnchorNav) StaticLabelClassName(value string) *AnchorNav{
	a.Set("staticLabelClassName", value)
	return a
}

/**
 * 是否静态展示表达式
 */
func (a *AnchorNav) StaticOn(value string) *AnchorNav{
	a.Set("staticOn", value)
	return a
}

/**
 * 静态展示空值占位
 */
func (a *AnchorNav) StaticPlaceholder(value string) *AnchorNav{
	a.Set("staticPlaceholder", value)
	return a
}

/**
 *
 */
func (a *AnchorNav) StaticSchema(value string) *AnchorNav{
	a.Set("staticSchema", value)
	return a
}

/**
 * 组件样式
 */
func (a *AnchorNav) Style(value string) *AnchorNav{
	a.Set("style", value)
	return a
}

/**
 * 指定为 AnchorNav 锚点导航渲染器
 */
func (a *AnchorNav) Type(value string) *AnchorNav{
	a.Set("type", value)
	return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *AnchorNav) UseMobileUI(value bool) *AnchorNav{
	a.Set("useMobileUI", value)
	return a
}

/**
 * 是否显示
 */
func (a *AnchorNav) Visible(value bool) *AnchorNav{
	a.Set("visible", value)
	return a
}

/**
 * 是否显示表达式
 */
func (a *AnchorNav) VisibleOn(value string) *AnchorNav{
	a.Set("visibleOn", value)
	return a
}