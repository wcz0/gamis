package renderers

type InputTimeRange struct {
	*BaseRenderer
}

func NewInputTimeRange() *InputTimeRange {
	i := &InputTimeRange{
		BaseRenderer: NewBaseRenderer(),
	}
	i.Set("type", "input-time-range")
	return i
}

func (i *InputTimeRange) Set(name string, value interface{}) *InputTimeRange {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	i.AmisSchema[name] = value
	return i
}

/**
 * 数据录入配置，自动填充或者参照录入
 */
func (i *InputTimeRange) AutoFill(value interface{}) *InputTimeRange {
	i.Set("autoFill", value)
	return i
}

/**
 * 表单最外层类名
 */
func (i *InputTimeRange) ClassName(value interface{}) *InputTimeRange {
	i.Set("className", value)
	return i
}

/**
 * 是否可清除。
 */
func (i *InputTimeRange) Clearable(value bool) *InputTimeRange {
	i.Set("clearable", value)
	return i
}

/**
 * 表单项描述
 */
func (i *InputTimeRange) Description(value interface{}) *InputTimeRange {
	i.Set("description", value)
	return i
}

/**
 * 是否禁用
 */
func (i *InputTimeRange) Disabled(value bool) *InputTimeRange {
	i.Set("disabled", value)
	return i
}

/**
 * 当前表单项是否禁用的条件
 */
func (i *InputTimeRange) DisabledOn(value interface{}) *InputTimeRange {
	i.Set("disabledOn", value)
	return i
}

/**
 * 是否内联模式。
 */
func (i *InputTimeRange) Embed(value bool) *InputTimeRange {
	i.Set("embed", value)
	return i
}

/**
 * 时间范围选择器值格式
 */
func (i *InputTimeRange) Format(value interface{}) *InputTimeRange {
	i.Set("format", value)
	return i
}

/**
 * 是否内联
 */
func (i *InputTimeRange) Inline(value bool) *InputTimeRange {
	i.Set("inline", value)
	return i
}

/**
 * 表单控制器类名
 */
func (i *InputTimeRange) InputClassName(value interface{}) *InputTimeRange {
	i.Set("inputClassName", value)
	return i
}

/**
 * 时间范围选择器显示格式
 */
func (i *InputTimeRange) InputFormat(value interface{}) *InputTimeRange {
	i.Set("inputFormat", value)
	return i
}

/**
 * 表单项标签
 */
func (i *InputTimeRange) Label(value interface{}) *InputTimeRange {
	i.Set("label", value)
	return i
}

/**
 * 表单项标签对齐方式，默认右对齐，仅在 mode为horizontal 时生效
 */
func (i *InputTimeRange) LabelAlign(value interface{}) *InputTimeRange {
	i.Set("labelAlign", value)
	return i
}

/**
 * label 的类名
 */
func (i *InputTimeRange) LabelClassName(value interface{}) *InputTimeRange {
	i.Set("labelClassName", value)
	return i
}

/**
 * 表单项标签描述
 */
func (i *InputTimeRange) LabelRemark(value interface{}) *InputTimeRange {
	i.Set("labelRemark", value)
	return i
}

/**
 * 字段名，指定该表单项提交时的 key
 */
func (i *InputTimeRange) Name(value interface{}) *InputTimeRange {
	i.Set("name", value)
	return i
}

/**
 * 占位文本
 */
func (i *InputTimeRange) Placeholder(value interface{}) *InputTimeRange {
	i.Set("placeholder", value)
	return i
}

/**
 * 是否必填
 */
func (i *InputTimeRange) Required(value bool) *InputTimeRange {
	i.Set("required", value)
	return i
}

/**
 * 通过表达式来配置当前表单项是否为必填。
 */
func (i *InputTimeRange) RequiredOn(value interface{}) *InputTimeRange {
	i.Set("requiredOn", value)
	return i
}

/**
 * 是否该表单项值发生变化时就提交当前表单。
 */
func (i *InputTimeRange) SubmitOnChange(value bool) *InputTimeRange {
	i.Set("submitOnChange", value)
	return i
}

/**
 * 时间范围选择器值格式
 */
func (i *InputTimeRange) TimeFormat(value interface{}) *InputTimeRange {
	i.Set("timeFormat", value)
	return i
}

/**
 * 指定为 input-time-range 渲染器。
 */
func (i *InputTimeRange) Type(value interface{}) *InputTimeRange {
	i.Set("type", value)
	return i
}

/**
 * 表单校验接口
 */
func (i *InputTimeRange) ValidateApi(value interface{}) *InputTimeRange {
	i.Set("validateApi", value)
	return i
}

/**
 * 表单项值格式验证，支持设置多个，多个规则用英文逗号隔开。
 */
func (i *InputTimeRange) Validations(value interface{}) *InputTimeRange {
	i.Set("validations", value)
	return i
}

/**
 * 表单默认值
 */
func (i *InputTimeRange) Value(value interface{}) *InputTimeRange {
	i.Set("value", value)
	return i
}

/**
 * 是否可见
 */
func (i *InputTimeRange) Visible(value bool) *InputTimeRange {
	i.Set("visible", value)
	return i
}

/**
 * 当前表单项是否禁用的条件
 */
func (i *InputTimeRange) VisibleOn(value interface{}) *InputTimeRange {
	i.Set("visibleOn", value)
	return i
}
