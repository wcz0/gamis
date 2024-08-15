package renderers

type InputYearRange struct {
	*BaseRenderer
}

func NewInputYearRange() *InputYearRange {
	i := &InputYearRange{
		BaseRenderer: NewBaseRenderer(),
	}
	i.Set("type", "input-year-range")
	return i
}

func (i *InputYearRange) Set(name string, value interface{}) *InputYearRange {
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
func (i *InputYearRange) AutoFill(value interface{}) *InputYearRange {
	i.Set("autoFill", value)
	return i
}

/**
 * 表单最外层类名
 */
func (i *InputYearRange) ClassName(value interface{}) *InputYearRange {
	i.Set("className", value)
	return i
}

/**
 * 是否可清除。
 */
func (i *InputYearRange) Clearable(value bool) *InputYearRange {
	i.Set("clearable", value)
	return i
}

/**
 * 表单项描述
 */
func (i *InputYearRange) Description(value interface{}) *InputYearRange {
	i.Set("description", value)
	return i
}

/**
 * 是否禁用
 */
func (i *InputYearRange) Disabled(value bool) *InputYearRange {
	i.Set("disabled", value)
	return i
}

/**
 * 当前表单项是否禁用的条件
 */
func (i *InputYearRange) DisabledOn(value interface{}) *InputYearRange {
	i.Set("disabledOn", value)
	return i
}

/**
 * 是否内联模式。
 */
func (i *InputYearRange) Embed(value bool) *InputYearRange {
	i.Set("embed", value)
	return i
}

/**
 * 年份选择器值格式
 */
func (i *InputYearRange) Format(value interface{}) *InputYearRange {
	i.Set("format", value)
	return i
}

/**
 * 是否内联
 */
func (i *InputYearRange) Inline(value bool) *InputYearRange {
	i.Set("inline", value)
	return i
}

/**
 * 表单控制器类名
 */
func (i *InputYearRange) InputClassName(value interface{}) *InputYearRange {
	i.Set("inputClassName", value)
	return i
}

/**
 * 年份选择器显示格式
 */
func (i *InputYearRange) InputFormat(value interface{}) *InputYearRange {
	i.Set("inputFormat", value)
	return i
}

/**
 * 表单项标签
 */
func (i *InputYearRange) Label(value interface{}) *InputYearRange {
	i.Set("label", value)
	return i
}

/**
 * 表单项标签对齐方式，默认右对齐，仅在 mode为horizontal 时生效
 */
func (i *InputYearRange) LabelAlign(value interface{}) *InputYearRange {
	i.Set("labelAlign", value)
	return i
}

/**
 * label 的类名
 */
func (i *InputYearRange) LabelClassName(value interface{}) *InputYearRange {
	i.Set("labelClassName", value)
	return i
}

/**
 * 表单项标签描述
 */
func (i *InputYearRange) LabelRemark(value interface{}) *InputYearRange {
	i.Set("labelRemark", value)
	return i
}

/**
 * 限制最大日期，用法同 限制范围
 */
func (i *InputYearRange) MaxDate(value interface{}) *InputYearRange {
	i.Set("maxDate", value)
	return i
}

/**
 * 限制最大跨度，如：4year
 */
func (i *InputYearRange) MaxDuration(value interface{}) *InputYearRange {
	i.Set("maxDuration", value)
	return i
}

/**
 * 限制最小日期，用法同 限制范围
 */
func (i *InputYearRange) MinDate(value interface{}) *InputYearRange {
	i.Set("minDate", value)
	return i
}

/**
 * 限制最小跨度，如： 2year
 */
func (i *InputYearRange) MinDuration(value interface{}) *InputYearRange {
	i.Set("minDuration", value)
	return i
}

/**
 * 字段名，指定该表单项提交时的 key
 */
func (i *InputYearRange) Name(value interface{}) *InputYearRange {
	i.Set("name", value)
	return i
}

/**
 * 占位文本
 */
func (i *InputYearRange) Placeholder(value interface{}) *InputYearRange {
	i.Set("placeholder", value)
	return i
}

/**
 * 是否必填
 */
func (i *InputYearRange) Required(value bool) *InputYearRange {
	i.Set("required", value)
	return i
}

/**
 * 通过表达式来配置当前表单项是否为必填。
 */
func (i *InputYearRange) RequiredOn(value interface{}) *InputYearRange {
	i.Set("requiredOn", value)
	return i
}

/**
 * 是否该表单项值发生变化时就提交当前表单。
 */
func (i *InputYearRange) SubmitOnChange(value bool) *InputYearRange {
	i.Set("submitOnChange", value)
	return i
}

/**
 * 指定为 input-year-range 渲染器。
 */
func (i *InputYearRange) Type(value interface{}) *InputYearRange {
	i.Set("type", value)
	return i
}

/**
 * 保存 UTC 值
 */
func (i *InputYearRange) Utc(value bool) *InputYearRange {
	i.Set("utc", value)
	return i
}

/**
 * 表单校验接口
 */
func (i *InputYearRange) ValidateApi(value interface{}) *InputYearRange {
	i.Set("validateApi", value)
	return i
}

/**
 * 表单项值格式验证，支持设置多个，多个规则用英文逗号隔开。
 */
func (i *InputYearRange) Validations(value interface{}) *InputYearRange {
	i.Set("validations", value)
	return i
}

/**
 * 表单默认值
 */
func (i *InputYearRange) Value(value interface{}) *InputYearRange {
	i.Set("value", value)
	return i
}

/**
 * 是否可见
 */
func (i *InputYearRange) Visible(value bool) *InputYearRange {
	i.Set("visible", value)
	return i
}

/**
 * 当前表单项是否禁用的条件
 */
func (i *InputYearRange) VisibleOn(value interface{}) *InputYearRange {
	i.Set("visibleOn", value)
	return i
}
