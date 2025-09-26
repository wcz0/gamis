package renderers

type Breadcrumb struct {
	*BaseRenderer
}

func NewBreadcrumb() *Breadcrumb {
	a := &Breadcrumb{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "breadcrumb")
	return a
}

func (a *Breadcrumb) Set(name string, value interface{}) *Breadcrumb {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Breadcrumb) ClassName(value interface{}) *Breadcrumb {
	a.Set("className", value)
	return a
}

func (a *Breadcrumb) DropdownClassName(value interface{}) *Breadcrumb {
	a.Set("dropdownClassName", value)
	return a
}

func (a *Breadcrumb) DropdownItemClassName(value interface{}) *Breadcrumb {
	a.Set("dropdownItemClassName", value)
	return a
}

func (a *Breadcrumb) ItemClassName(value interface{}) *Breadcrumb {
	a.Set("itemClassName", value)
	return a
}

func (a *Breadcrumb) Items(value interface{}) *Breadcrumb {
	a.Set("items", value)
	return a
}

func (a *Breadcrumb) LabelMaxLength(value interface{}) *Breadcrumb {
	a.Set("labelMaxLength", value)
	return a
}

func (a *Breadcrumb) Separator(value interface{}) *Breadcrumb {
	a.Set("separator", value)
	return a
}

func (a *Breadcrumb) SeparatorClassName(value interface{}) *Breadcrumb {
	a.Set("separatorClassName", value)
	return a
}

func (a *Breadcrumb) Source(value interface{}) *Breadcrumb {
	a.Set("source", value)
	return a
}

func (a *Breadcrumb) TooltipPosition(value interface{}) *Breadcrumb {
	a.Set("tooltipPosition", value)
	return a
}

func (a *Breadcrumb) Type(value interface{}) *Breadcrumb {
	a.Set("type", value)
	return a
}
