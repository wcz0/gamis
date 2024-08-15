package renderers

type InputKV struct {
	*BaseRenderer
}

func NewInputKV() *InputKV {
	a := &InputKV{
		BaseRenderer: NewBaseRenderer(),
	}
	a.Set("type", "input-kv")
	return a
}

func (a *InputKV) Set(name string, value interface{}) *InputKV {
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
func (a *InputKV) AutoFill(value interface{}) *InputKV {
	a.Set("autoFill", value)
	return a
}

/**
 * 表单最外层类名
 */
func (a *InputKV) ClassName(value interface{}) *InputKV {
	a.Set("className", value)
	return a
}

/**
 * 默认值
 */
func (a *InputKV) DefaultValue(value interface{}) *InputKV {
	a.Set("defaultValue", value)
	return a
}

/**
 * 表单项描述
 */
func (a *InputKV) Description(value interface{}) *InputKV {
	a.Set("description", value)
	return a
}

/**
 * 是否禁用
 */
func (a *InputKV) Disabled(value bool) *InputKV {
	a.Set("disabled", value)
	return a
}

/**
 * 当前表单项是否禁用的条件
 */
func (a *InputKV) DisabledOn(value interface{}) *InputKV {
	a.Set("disabledOn", value)
	return a
}

/**
 * 是否可以拖拽排序
 */
func (a *InputKV) Draggable(value bool) *InputKV {
	a.Set("draggable", value)
	return a
}

/**
 * 是否内联
 */
func (a *InputKV) Inline(value bool) *InputKV {
	a.Set("inline", value)
	return a
}

/**
 * 表单控制器类名
 */
func (a *InputKV) InputClassName(value interface{}) *InputKV {
	a.Set("inputClassName", value)
	return a
}

/**
 * key 的提示信息的
 */
func (a *InputKV) KeyPlaceholder(value interface{}) *InputKV {
	a.Set("keyPlaceholder", value)
	return a
}

/**
 * 表单项标签
 */
func (a *InputKV) Label(value interface{}) *InputKV {
	a.Set("label", value)
	return a
}

/**
 * 表单项标签对齐方式，默认右对齐，仅在 mode为horizontal 时生效
 */
func (a *InputKV) LabelAlign(value interface{}) *InputKV {
	a.Set("labelAlign", value)
	return a
}

/**
 * label 的类名
 */
func (a *InputKV) LabelClassName(value interface{}) *InputKV {
	a.Set("labelClassName", value)
	return a
}

/**
 * 表单项标签描述
 */
func (a *InputKV) LabelRemark(value interface{}) *InputKV {
	a.Set("labelRemark", value)
	return a
}

/**
 * 字段名，指定该表单项提交时的 key
 */
func (a *InputKV) Name(value interface{}) *InputKV {
	a.Set("name", value)
	return a
}

/**
 * 表单项描述
 */
func (a *InputKV) Placeholder(value interface{}) *InputKV {
	a.Set("placeholder", value)
	return a
}

/**
 * 是否必填
 */
func (a *InputKV) Required(value bool) *InputKV {
	a.Set("required", value)
	return a
}

/**
 * 通过表达式来配置当前表单项是否为必填。
 */
func (a *InputKV) RequiredOn(value interface{}) *InputKV {
	a.Set("requiredOn", value)
	return a
}

/**
 * 是否该表单项值发生变化时就提交当前表单。
 */
func (a *InputKV) SubmitOnChange(value bool) *InputKV {
	a.Set("submitOnChange", value)
	return a
}

/**
 * 指定为 input-kv 渲染器。
 */
func (a *InputKV) Type(value interface{}) *InputKV {
	a.Set("type", value)
	return a
}

/**
 * 表单校验接口
 */
func (a *InputKV) ValidateApi(value interface{}) *InputKV {
	a.Set("validateApi", value)
	return a
}

/**
 * 表单项值格式验证，支持设置多个，多个规则用英文逗号隔开。
 */
func (a *InputKV) Validations(value interface{}) *InputKV {
	a.Set("validations", value)
	return a
}

/**
 * 表单默认值
 */
func (a *InputKV) Value(value interface{}) *InputKV {
	a.Set("value", value)
	return a
}

/**
 * value 的提示信息的
 */
func (a *InputKV) ValuePlaceholder(value interface{}) *InputKV {
	a.Set("valuePlaceholder", value)
	return a
}

/**
 * 值类型
 */
func (a *InputKV) ValueType(value interface{}) *InputKV {
	a.Set("valueType", value)
	return a
}

/**
 * 是否可见
 */
func (a *InputKV) Visible(value bool) *InputKV {
	a.Set("visible", value)
	return a
}

/**
 * 当前表单项是否禁用的条件
 */
func (a *InputKV) VisibleOn(value interface{}) *InputKV {
	a.Set("visibleOn", value)
	return a
}
