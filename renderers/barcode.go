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

func (b *Barcode) Set(name string, value interface{}) *Barcode {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	b.AmisSchema[name] = value
	return b
}

/**
 * 外层类名
 */
func (b *Barcode) ClassName(value interface{}) *Barcode {
	b.Set("className", value)
	return b
}

/**
 * 指定为 barcode 渲染器。
 */
func (b *Barcode) Type(value interface{}) *Barcode {
	b.Set("type", value)
	return b
}
