// <?php

// namespace Slowlyo\AmisRenderers;

// /**
//  * Barcode 条形码 https://aisuda.bce.baidu.com/amis/zh-CN/components/barcode
//  *
//  * @author slowlyo
//  * @version 6.2.2
//  */
// class Barcode extends BaseRenderer
// {
//     public function __construct()
//     {
//         $this->set('type', 'barcode');


//     }

//     /**
//      * 外层类名
//      */
//     public function className($value = '')
//     {
//         return $this->set('className', $value);
//     }

//     /**
//      * 指定为 barcode 渲染器。
//      */
//     public function type($value = 'barcode')
//     {
//         return $this->set('type', $value);
//     }


// }

package renderers

type Barcode struct {
	*BaseRenderer
}

func NewBarcode() *Barcode {
	b := &Barcode{
		BaseRenderer: NewBaseRenderer(),
	}
	b.Set("type", "barcode")
	return b
}

/**
 * 外层类名
 */
func (b *Barcode) ClassName(value string) *Barcode {
	b.Set("className", value)
	return b
}

/**
 * 指定为 barcode 渲染器。
 */
func (b *Barcode) Type(value string) *Barcode {
	b.Set("type", value)
	return b
}
