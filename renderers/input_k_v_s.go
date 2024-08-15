package renderers

type InputKVS struct {
	*BaseRenderer
}

func NewInputKVS() *InputKVS {
	a := &InputKVS{
		BaseRenderer: NewBaseRenderer(),
	}
	a.Set("type", "input-kvs")
	return a
}

func (a *InputKVS) Set(name string, value interface{}) *InputKVS {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

/**
 * 数据录入配置，自动填充或者参照录入
 */
func (a *InputKVS) AutoFill(value interface{}) *InputKVS {
	a.Set("autoFill", value)
	return a
}

/**
 * 表单最外层类名
 */
func (a *InputKVS) ClassName(value interface{}) *InputKVS {
	a.Set("className", value)
	return a
}

/**
 * 表单项描述
 */
func (a *InputKVS) Description(value interface{}) *InputKVS {
	a.Set("description", value)
	return a
}

/**
 * 是否禁用
 */
func (a *InputKVS) Disabled(value bool) *InputKVS {
	a.Set("disabled", value)
	return a
}

/**
 * 当前表单项是否禁用的条件
 */
func (a *InputKVS) DisabledOn(value interface{}) *InputKVS {
	a.Set("disabledOn", value)
	return a
}

/**
 * 是否内联
 */
func (a *InputKVS) Inline(value bool) *InputKVS {
	a.Set("inline", value)
	return a
}

/**
 * 表单控制器类名
 */
func (a *InputKVS) InputClassName(value interface{}) *InputKVS {
	a.Set("inputClassName", value)
	return a
}

/**
 * 表单项标签
 */
func (a *InputKVS) Label(value interface{}) *InputKVS {
	a.Set("label", value)
	return a
}

/**
 * 表单项标签对齐方式，默认右对齐，仅在 mode为horizontal 时生效
 */
func (a *InputKVS) LabelAlign(value interface{}) *InputKVS {
	a.Set("labelAlign", value)
	return a
}

/**
 * label 的类名
 */
func (a *InputKVS) LabelClassName(value interface{}) *InputKVS {
	a.Set("labelClassName", value)
	return a
}

/**
 * 表单项标签描述
 */
func (a *InputKVS) LabelRemark(value interface{}) *InputKVS {
	a.Set("labelRemark", value)
	return a
}

/**
 * 字段名，指定该表单项提交时的 key
 */
func (a *InputKVS) Name(value interface{}) *InputKVS {
	a.Set("name", value)
	return a
}

/**
 * 表单项描述
 */
func (a *InputKVS) Placeholder(value interface{}) *InputKVS {
	a.Set("placeholder", value)
	return a
}

/**
 * 是否必填
 */
func (a *InputKVS) Required(value bool) *InputKVS {
	a.Set("required", value)
	return a
}

/**
 * 通过表达式来配置当前表单项是否为必填。
 */
func (a *InputKVS) RequiredOn(value interface{}) *InputKVS {
	a.Set("requiredOn", value)
	return a
}

/**
 * 是否该表单项值发生变化时就提交当前表单。
 */
func (a *InputKVS) SubmitOnChange(value bool) *InputKVS {
	a.Set("submitOnChange", value)
	return a
}

/**
 * 指定为 input-kvs 渲染器。
 */
func (a *InputKVS) Type(value interface{}) *InputKVS {
	a.Set("type", value)
	return a
}

/**
 * 表单校验接口
 */
func (a *InputKVS) ValidateApi(value interface{}) *InputKVS {
	a.Set("validateApi", value)
	return a
}

/**
 * 表单项值格式验证，支持设置多个，多个规则用英文逗号隔开。
 */
func (a *InputKVS) Validations(value interface{}) *InputKVS {
	a.Set("validations", value)
	return a
}

/**
 * 表单默认值
 */
func (a *InputKVS) Value(value interface{}) *InputKVS {
	a.Set("value", value)
	return a
}

/**
 * 是否可见
 */
func (a *InputKVS) Visible(value bool) *InputKVS {
	a.Set("visible", value)
	return a
}

/**
 * 当前表单项是否禁用的条件
 */
func (a *InputKVS) VisibleOn(value interface{}) *InputKVS {
	a.Set("visibleOn", value)
	return a
}
