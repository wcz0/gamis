// <?php

// namespace Slowlyo\AmisRenderers;

// /**
//  * Avatar
//  *
//  * @author slowlyo
//  * @version 6.2.2
//  */
// class Avatar extends BaseRenderer
// {
//     public function __construct()
//     {
//         $this->set('type', 'avatar');
// $this->set('crossOrigin', 'anonymous');


//     }

//     /**
//      * 图片无法显示时的替换文字地址
//      */
//     public function alt($value = '')
//     {
//         return $this->set('alt', $value);
//     }

//     /**
//      * 角标
//      */
//     public function badge($value = '')
//     {
//         return $this->set('badge', $value);
//     }

//     /**
//      * 类名
//      */
//     public function className($value = '')
//     {
//         return $this->set('className', $value);
//     }

//     /**
//      * 图片CORS属性 可选值: anonymous | use-credentials |
//      */
//     public function crossOrigin($value = '')
//     {
//         return $this->set('crossOrigin', $value);
//     }

//     /**
//      * 默认头像
//      */
//     public function defaultAvatar($value = '')
//     {
//         return $this->set('defaultAvatar', $value);
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
//      * 图片是否允许拖动
//      */
//     public function draggable($value = true)
//     {
//         return $this->set('draggable', $value);
//     }

//     /**
//      * 编辑器配置，运行时可以忽略
//      */
//     public function editorSetting($value = '')
//     {
//         return $this->set('editorSetting', $value);
//     }

//     /**
//      * 图片相对于容器的缩放方式 可选值: fill | contain | cover | none | scale-down
//      */
//     public function fit($value = '')
//     {
//         return $this->set('fit', $value);
//     }

//     /**
//      * 字符类型距离左右两侧边界单位像素
//      */
//     public function gap($value = '')
//     {
//         return $this->set('gap', $value);
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
//      * 图标
//      */
//     public function icon($value = '')
//     {
//         return $this->set('icon', $value);
//     }

//     /**
//      * 组件唯一 id，主要用于日志采集
//      */
//     public function id($value = '')
//     {
//         return $this->set('id', $value);
//     }

//     /**
//      * 图片加载失败的是否默认处理，字符串函数
//      */
//     public function onError($value = '')
//     {
//         return $this->set('onError', $value);
//     }

//     /**
//      * 事件动作配置
//      */
//     public function onEvent($value = '')
//     {
//         return $this->set('onEvent', $value);
//     }

//     /**
//      * 形状 可选值: circle | square | rounded
//      */
//     public function shape($value = '')
//     {
//         return $this->set('shape', $value);
//     }

//     /**
//      * 大小
//      */
//     public function size($value = '')
//     {
//         return $this->set('size', $value);
//     }

//     /**
//      * 图片地址
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
//      * 自定义样式
//      */
//     public function style($value = '')
//     {
//         return $this->set('style', $value);
//     }

//     /**
//      * 文本
//      */
//     public function text($value = '')
//     {
//         return $this->set('text', $value);
//     }

//     /**
//      *
//      */
//     public function type($value = 'avatar')
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

type Avatar struct {
	*BaseRenderer
}

func NewAvatar() *Avatar{
	a := &Avatar{
		BaseRenderer: NewBaseRenderer(),
	}
	a.Set("type", "avatar")
	a.Set("crossOrigin", "anonymous")
	return a
}

/**
 * 图片无法显示时的替换文字地址
 */
func (a *Avatar) Alt(value string) *Avatar{
	a.Set("alt", value)
	return a
}

/**
 * 角标
 */
func (a *Avatar) Badge(value string) *Avatar{
	a.Set("badge", value)
	return a
}

/**
 * 类名
 */
func (a *Avatar) ClassName(value string) *Avatar{
	a.Set("className", value)
	return a
}

/**
 * 图片CORS属性 可选值: anonymous | use-credentials |
 */
func (a *Avatar) CrossOrigin(value string) *Avatar{
	a.Set("crossOrigin", value)
	return a
}

/**
 * 默认头像
 */
func (a *Avatar) DefaultAvatar(value string) *Avatar{
	a.Set("defaultAvatar", value)
	return a
}

/**
 * 是否禁用
 */
func (a *Avatar) Disabled(value bool) *Avatar{
	a.Set("disabled", value)
	return a
}

/**
 * 是否禁用表达式
 */
func (a *Avatar) DisabledOn(value string) *Avatar{
	a.Set("disabledOn", value)
	return a
}

/**
 * 图片是否允许拖动
 */
func (a *Avatar) Draggable(value bool) *Avatar{
	a.Set("draggable", value)
	return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Avatar) EditorSetting(value string) *Avatar{
	a.Set("editorSetting", value)
	return a
}

/**
 * 图片相对于容器的缩放方式 可选值: fill | contain | cover | none | scale-down
 */
func (a *Avatar) Fit(value string) *Avatar{
	a.Set("fit", value)
	return a
}

/**
 * 字符类型距离左右两侧边界单位像素
 */
func (a *Avatar) Gap(value string) *Avatar{
	a.Set("gap", value)
	return a
}

/**
 * 是否隐藏
 */
func (a *Avatar) Hidden(value bool) *Avatar{
	a.Set("hidden", value)
	return a
}

/**
 * 是否隐藏表达式
 */
func (a *Avatar) HiddenOn(value string) *Avatar{
	a.Set("hiddenOn", value)
	return a
}

/**
 * 图标
 */
func (a *Avatar) Icon(value string) *Avatar{
	a.Set("icon", value)
	return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Avatar) Id(value string) *Avatar{
	a.Set("id", value)
	return a
}

/**
 * 图片加载失败的是否默认处理，字符串函数
 */
func (a *Avatar) OnError(value string) *Avatar{
	a.Set("onError", value)
	return a
}

/**
 * 事件动作配置
 */
func (a *Avatar) OnEvent(value string) *Avatar{
	a.Set("onEvent", value)
	return a
}

/**
 * 形状 可选值: circle | square | rounded
 */
func (a *Avatar) Shape(value string) *Avatar{
	a.Set("shape", value)
	return a
}

/**
 * 大小
 */
func (a *Avatar) Size(value string) *Avatar{
	a.Set("size", value)
	return a
}

/**
 * 图片地址
 */
func (a *Avatar) Src(value string) *Avatar{
	a.Set("src", value)
	return a
}

/**
 * 是否静态展示
 */
func (a *Avatar) Static(value bool) *Avatar{
	a.Set("static", value)
	return a
}

/**
 * 静态展示表单项类名
 */
func (a *Avatar) StaticClassName(value string) *Avatar{
	a.Set("staticClassName", value)
	return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Avatar) StaticInputClassName(value string) *Avatar{
	a.Set("staticInputClassName", value)
	return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Avatar) StaticLabelClassName(value string) *Avatar{
	a.Set("staticLabelClassName", value)
	return a
}

/**
 * 是否静态展示表达式
 */
func (a *Avatar) StaticOn(value string) *Avatar{
	a.Set("staticOn", value)
	return a
}

/**
 * 静态展示空值占位
 */
func (a *Avatar) StaticPlaceholder(value string) *Avatar{
	a.Set("staticPlaceholder", value)
	return a
}

/**
 *
 */
func (a *Avatar) StaticSchema(value string) *Avatar{
	a.Set("staticSchema", value)
	return a
}

/**
 * 自定义样式
 */
func (a *Avatar) Style(value string) *Avatar{
	a.Set("style", value)
	return a
}

/**
 * 文本
 */
func (a *Avatar) Text(value string) *Avatar{
	a.Set("text", value)
	return a
}

/**
 *
 */
func (a *Avatar) Type(value string) *Avatar{
	a.Set("type", value)
	return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Avatar) UseMobileUI(value bool) *Avatar{
	a.Set("useMobileUI", value)
	return a
}

/**
 * 是否显示
 */
func (a *Avatar) Visible(value bool) *Avatar{
	a.Set("visible", value)
	return a
}

/**
 * 是否显示表达式
 */
func (a *Avatar) VisibleOn(value string) *Avatar{
	a.Set("visibleOn", value)
	return a
}