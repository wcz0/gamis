// <?php

// namespace Slowlyo\AmisRenderers;

// /**
//  * Badge 角标。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/badge
//  *
//  * @author slowlyo
//  * @version 6.2.2
//  */
// class Badge extends BaseRenderer
// {
//     public function __construct()
//     {


//     }

//     /**
//      * 是否显示动画
//      */
//     public function animation($value = true)
//     {
//         return $this->set('animation', $value);
//     }

//     /**
//      *
//      */
//     public function className($value = '')
//     {
//         return $this->set('className', $value);
//     }

//     /**
//      * 提示类型
//      */
//     public function level($value = '')
//     {
//         return $this->set('level', $value);
//     }

//     /**
//      * 角标类型 可选值: text | dot | ribbon
//      */
//     public function mode($value = '')
//     {
//         return $this->set('mode', $value);
//     }

//     /**
//      * 角标位置，相对于position的位置进行偏移
//      */
//     public function offset($value = '')
//     {
//         return $this->set('offset', $value);
//     }

//     /**
//      * 封顶的数字值
//      */
//     public function overflowCount($value = '')
//     {
//         return $this->set('overflowCount', $value);
//     }

//     /**
//      * 角标位置 可选值: top-right | top-left | bottom-right | bottom-left
//      */
//     public function position($value = '')
//     {
//         return $this->set('position', $value);
//     }

//     /**
//      * 大小
//      */
//     public function size($value = '')
//     {
//         return $this->set('size', $value);
//     }

//     /**
//      * 角标的自定义样式
//      */
//     public function style($value = '')
//     {
//         return $this->set('style', $value);
//     }

//     /**
//      * 文本内容
//      */
//     public function text($value = '')
//     {
//         return $this->set('text', $value);
//     }

//     /**
//      * 动态控制是否显示
//      */
//     public function visibleOn($value = '')
//     {
//         return $this->set('visibleOn', $value);
//     }


// }

package renderers

type Badge struct {
	*BaseRenderer
}

func NewBadge() *Badge{
	return &Badge{
		BaseRenderer: NewBaseRenderer(),
	}
}

/**
 * 是否显示动画
 */
func (b *Badge) Animation(value bool) *Badge {
	b.Set("animation", value)
	return b
}

/**
 *
 */
func (b *Badge) ClassName(value string) *Badge {
	b.Set("className", value)
	return b
}

/**
 * 提示类型
 */
func (b *Badge) Level(value string) *Badge {
	b.Set("level", value)
	return b
}

/**
 * 角标类型 可选值: text | dot | ribbon
 */
func (b *Badge) Mode(value string) *Badge {
	b.Set("mode", value)
	return b
}

/**
 * 角标位置，相对于position的位置进行偏移
 */
func (b *Badge) Offset(value string) *Badge {
	b.Set("offset", value)
	return b
}

/**
 * 封顶的数字值
 */
func (b *Badge) OverflowCount(value string) *Badge {
	b.Set("overflowCount", value)
	return b
}

/**
 * 角标位置 可选值: top-right | top-left | bottom-right | bottom-left
 */
func (b *Badge) Position(value string) *Badge {
	b.Set("position", value)
	return b
}

/**
 * 大小
 */
func (b *Badge) Size(value string) *Badge {
	b.Set("size", value)
	return b
}

/**
 * 角标的自定义样式
 */
func (b *Badge) Style(value string) *Badge {
	b.Set("style", value)
	return b
}

/**
 * 文本内容
 */
func (b *Badge) Text(value string) *Badge {
	b.Set("text", value)
	return b
}

/**
 * 动态控制是否显示
 */
func (b *Badge) VisibleOn(value string) *Badge {
	b.Set("visibleOn", value)
	return b
}


