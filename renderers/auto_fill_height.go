// <?php

// namespace Slowlyo\AmisRenderers;

// /**
//  * AutoFillHeight
//  *
//  * @author slowlyo
//  * @version 6.2.2
//  */
// class AutoFillHeight extends BaseRenderer
// {
//     public function __construct()
//     {


//     }

//     /**
//      *
//      */
//     public function height($value = '')
//     {
//         return $this->set('height', $value);
//     }

//     /**
//      *
//      */
//     public function maxHeight($value = '')
//     {
//         return $this->set('maxHeight', $value);
//     }


// }

package renderers

type AutoFillHeight struct {
	*BaseRenderer
}

func NewAutoFillHeight() *AutoFillHeight {
	return &AutoFillHeight{
		BaseRenderer: NewBaseRenderer(),
	}
}

func (a *AutoFillHeight) Height(value string) *AutoFillHeight {
	a.Set("height", value)
	return a
}

func (a *AutoFillHeight) MaxHeight(value string) *AutoFillHeight {
	a.Set("maxHeight", value)
	return a
}

