package renderers

type Toast struct {
	*BaseRenderer
}

func NewToast() *Toast {
	t := &Toast{
		BaseRenderer: NewBaseRenderer(),
	}
	t.Set("type", "toast")
	return t
}

/**
 * 内容
 */
func (t *Toast) Body(value string) *Toast {
	t.Set("body", value)
	return t
}

/**
 * 是否显示关闭按钮
 */
func (t *Toast) CloseButton(value bool) *Toast {
	t.Set("closeButton", value)
	return t
}

/**
 * 轻提示内容
 */
func (t *Toast) Items(value string) *Toast {
	t.Set("items", value)
	return t
}

/**
 * 展示图标，可选'info'/'success'/'error'/'warning'
 */
func (t *Toast) Level(value string) *Toast {
	t.Set("level", value)
	return t
}

/**
 * 提示显示位置，可选值: top-right | top-center | top-left | bottom-center | bottom-left | bottom-right | center
 */
func (t *Toast) Position(value string) *Toast {
	t.Set("position", value)
	return t
}

/**
 * 是否显示图标
 */
func (t *Toast) ShowIcon(value bool) *Toast {
	t.Set("showIcon", value)
	return t
}

/**
 * 持续时间
 */
func (t *Toast) Timeout(value string) *Toast {
	t.Set("timeout", value)
	return t
}

/**
 * 标题
 */
func (t *Toast) Title(value string) *Toast {
	t.Set("title", value)
	return t
}