// <?php

// namespace Slowlyo\AmisRenderers;

// /**
//  * AnchorNavSection 锚点区域渲染器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/anchor-nav
//  *
//  * @author slowlyo
//  * @version 6.2.2
//  */
// class AnchorNavSection extends BaseRenderer
// {
//     public function __construct()
//     {


//     }

//     /**
//      * 内容
//      */
//     public function body($value = '')
//     {
//         return $this->set('body', $value);
//     }

//     /**
//      * 子节点
//      */
//     public function children($value = '')
//     {
//         return $this->set('children', $value);
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
//      * 锚点链接
//      */
//     public function href($value = '')
//     {
//         return $this->set('href', $value);
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
//      * 导航文字说明
//      */
//     public function title($value = '')
//     {
//         return $this->set('title', $value);
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

type AnchorNavSection struct {
	*BaseRenderer
}

func NewAnchorNavSection() *AnchorNavSection{
	return &AnchorNavSection{
		BaseRenderer: NewBaseRenderer(),
	}
}

/**
 * 内容
 */
func (a *AnchorNavSection) Body(value string) *AnchorNavSection {
	a.Set("body", value)
	return a
}

/**
 * 子节点
 */
func (a *AnchorNavSection) Children(value string) *AnchorNavSection {
	a.Set("children", value)
	return a
}

/**
 * 容器 css 类名
 */
func (a *AnchorNavSection) ClassName(value string) *AnchorNavSection {
	a.Set("className", value)
	return a
}

/**
 * 是否禁用
 */
func (a *AnchorNavSection) Disabled(value bool) *AnchorNavSection {
	a.Set("disabled", value)
	return a
}

/**
 * 是否禁用表达式
 */
func (a *AnchorNavSection) DisabledOn(value string) *AnchorNavSection {
	a.Set("disabledOn", value)
	return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *AnchorNavSection) EditorSetting(value string) *AnchorNavSection {
	a.Set("editorSetting", value)
	return a
}

/**
 * 是否隐藏
 */
func (a *AnchorNavSection) Hidden(value bool) *AnchorNavSection {
	a.Set("hidden", value)
	return a
}

/**
 * 是否隐藏表达式
 */
func (a *AnchorNavSection) HiddenOn(value string) *AnchorNavSection {
	a.Set("hiddenOn", value)
	return a
}

/**
 * 锚点链接
 */
func (a *AnchorNavSection) Href(value string) *AnchorNavSection {
	a.Set("href", value)
	return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *AnchorNavSection) Id(value string) *AnchorNavSection {
	a.Set("id", value)
	return a
}

/**
 * 事件动作配置
 */
func (a *AnchorNavSection) OnEvent(value string) *AnchorNavSection {
	a.Set("onEvent", value)
	return a
}

/**
 * 是否静态展示
 */
func (a *AnchorNavSection) Static(value bool) *AnchorNavSection {
	a.Set("static", value)
	return a
}

/**
 * 静态展示表单项类名
 */
func (a *AnchorNavSection) StaticClassName(value string) *AnchorNavSection {
	a.Set("staticClassName", value)
	return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *AnchorNavSection) StaticInputClassName(value string) *AnchorNavSection {
	a.Set("staticInputClassName", value)
	return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *AnchorNavSection) StaticLabelClassName(value string) *AnchorNavSection {
	a.Set("staticLabelClassName", value)
	return a
}

/**
 * 是否静态展示表达式
 */
func (a *AnchorNavSection) StaticOn(value string) *AnchorNavSection {
	a.Set("staticOn", value)
	return a
}

/**
 * 静态展示空值占位
 */
func (a *AnchorNavSection) StaticPlaceholder(value string) *AnchorNavSection {
	a.Set("staticPlaceholder", value)
	return a
}

/**
 *
 */
func (a *AnchorNavSection) StaticSchema(value string) *AnchorNavSection {
	a.Set("staticSchema", value)
	return a
}

/**
 * 组件样式
 */
func (a *AnchorNavSection) Style(value string) *AnchorNavSection {
	a.Set("style", value)
	return a
}

/**
 * 导航文字说明
 */
func (a *AnchorNavSection) Title(value string) *AnchorNavSection {
	a.Set("title", value)
	return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *AnchorNavSection) UseMobileUI(value bool) *AnchorNavSection {
	a.Set("useMobileUI", value)
	return a
}

/**
 * 是否显示
 */
func (a *AnchorNavSection) Visible(value bool) *AnchorNavSection {
	a.Set("visible", value)
	return a
}

/**
 * 是否显示表达式
 */
func (a *AnchorNavSection) VisibleOn(value string) *AnchorNavSection {
	a.Set("visibleOn", value)
	return a
}
