package renderers

type Breadcrumb struct {
	*BaseRenderer
}

func NewBreadcrumb() *Breadcrumb {
	b := &Breadcrumb{
		BaseRenderer: NewBaseRenderer(),
	}
	b.Set("type", "breadcrumb")
	return b
}

/**
 * 外层类名
 */
func (b *Breadcrumb) ClassName(value interface{}) *Breadcrumb {
	b.Set("className", value)
	return b
}

/**
 * 下拉菜单类名
 */
func (b *Breadcrumb) DropdownClassName(value interface{}) *Breadcrumb {
	b.Set("dropdownClassName", value)
	return b
}

/**
 * 下拉菜单项类名
 */
func (b *Breadcrumb) DropdownItemClassName(value interface{}) *Breadcrumb {
	b.Set("dropdownItemClassName", value)
	return b
}

/**
 * 导航项类名
 */
func (b *Breadcrumb) ItemClassName(value interface{}) *Breadcrumb {
	b.Set("itemClassName", value)
	return b
}

/**
 * 文本
 */
func (b *Breadcrumb) Items(value interface{}) *Breadcrumb {
	b.Set("items", value)
	return b
}

/**
 * 最大展示长度
 */
func (b *Breadcrumb) LabelMaxLength(value interface{}) *Breadcrumb {
	b.Set("labelMaxLength", value)
	return b
}

/**
 * 分隔符
 */
func (b *Breadcrumb) Separator(value interface{}) *Breadcrumb {
	b.Set("separator", value)
	return b
}

/**
 * 分割符类名
 */
func (b *Breadcrumb) SeparatorClassName(value interface{}) *Breadcrumb {
	b.Set("separatorClassName", value)
	return b
}

/**
 * 动态数据
 */
func (b *Breadcrumb) Source(value interface{}) *Breadcrumb {
	b.Set("source", value)
	return b
}

/**
 * 浮窗提示位置
 */
func (b *Breadcrumb) TooltipPosition(value interface{}) *Breadcrumb {
	b.Set("tooltipPosition", value)
	return b
}

/**
 * 指定为 breadcrumb 渲染器。
 */
func (b *Breadcrumb) Type(value interface{}) *Breadcrumb {
	b.Set("type", value)
	return b
}
