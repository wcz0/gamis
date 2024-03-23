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
