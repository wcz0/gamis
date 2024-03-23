// <?php

// namespace Slowlyo\AmisRenderers;

// /**
//  * AutoGenerateFilter
//  *
//  * @author slowlyo
//  * @version 6.2.2
//  */
// class AutoGenerateFilter extends BaseRenderer
// {
//     public function __construct()
//     {


//     }

//     /**
//      * 过滤条件单行列数
//      */
//     public function columnsNum($value = '')
//     {
//         return $this->set('columnsNum', $value);
//     }

//     /**
//      * 是否默认收起
//      */
//     public function defaultCollapsed($value = true)
//     {
//         return $this->set('defaultCollapsed', $value);
//     }

//     /**
//      * 是否显示设置查询字段
//      */
//     public function showBtnToolbar($value = true)
//     {
//         return $this->set('showBtnToolbar', $value);
//     }


// }

package renderers

type AutoGenerateFilter struct {
	*BaseRenderer
}

func NewAutoGenerateFilter() *AutoGenerateFilter {
	return &AutoGenerateFilter{
		BaseRenderer: NewBaseRenderer(),
	}
}

/**
 * 过滤条件单行列数
 */
func (a *AutoGenerateFilter) ColumnsNum(value string) *AutoGenerateFilter {
	a.Set("columnsNum", value)
	return a
}

/**
 * 是否默认收起
 */
func (a *AutoGenerateFilter) DefaultCollapsed(value string) *AutoGenerateFilter {
	a.Set("defaultCollapsed", value)
	return a
}

/**
 * 是否显示设置查询字段
 */
func (a *AutoGenerateFilter) ShowBtnToolbar(value string) *AutoGenerateFilter {
	a.Set("showBtnToolbar", value)
	return a
}