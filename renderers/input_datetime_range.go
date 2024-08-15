package renderers

type InputDatetimeRange struct {
	*BaseRenderer
}

func NewInputDatetimeRange() *InputDatetimeRange {
	a := &InputDatetimeRange{
		BaseRenderer: NewBaseRenderer(),
	}
	a.Set("type", "input-datetime-range")
	return a
}

func (a *InputDatetimeRange) Set(name string, value interface{}) *InputDatetimeRange {
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
func (a *InputDatetimeRange) AutoFill(value interface{}) *InputDatetimeRange {
	a.Set("autoFill", value)
	return a
}

/*
*
  - 表单最外层类名
*/
func (a *InputDatetimeRange) ClassName(value interface{}) *InputDatetimeRange {
	a.Set("className", value)
	return a
}

/**
* - 是否可清除
 */
func (a *InputDatetimeRange) Clearable(value bool) *InputDatetimeRange {
	a.Set("clearable", value)
	return a
}

/**
* - 表单项描述
 */
func (a *InputDatetimeRange) Description(value interface{}) *InputDatetimeRange {
	a.Set("description", value)
	return a
}

/**
* - 是否禁用
 */
func (a *InputDatetimeRange) Disabled(value bool) *InputDatetimeRange {
	a.Set("disabled", value)
	return a
}

/**
* - 当前表单项是否禁用的条件
 */
func (a *InputDatetimeRange) DisabledOn(value interface{}) *InputDatetimeRange {
	a.Set("disabledOn", value)
	return a
}

/**
* - 日期时间选择器值格式
 */
func (a *InputDatetimeRange) Format(value interface{}) *InputDatetimeRange {
	a.Set("format", value)
	return a
}

/**
* - 是否内联
 */
func (a *InputDatetimeRange) Inline(value bool) *InputDatetimeRange {
	a.Set("inline", value)
	return a
}

/**
* - 表单控制器类名
 */
func (a *InputDatetimeRange) InputClassName(value interface{}) *InputDatetimeRange {
	a.Set("inputClassName", value)
	return a
}

/**
* - 日期时间选择器显示格式
 */
func (a *InputDatetimeRange) InputFormat(value interface{}) *InputDatetimeRange {
	a.Set("inputFormat", value)
	return a
}

/**
* - 表单项标签
 */
func (a *InputDatetimeRange) Label(value interface{}) *InputDatetimeRange {
	a.Set("label", value)
	return a
}

/**
* - 表单项标签对齐方式，默认右对齐，仅在 mode为horizontal 时生效
 */
func (a *InputDatetimeRange) LabelAlign(value interface{}) *InputDatetimeRange {
	a.Set("labelAlign", value)
	return a
}

/**
* - label 的类名
 */
func (a *InputDatetimeRange) LabelClassName(value interface{}) *InputDatetimeRange {
	a.Set("labelClassName", value)
	return a
}

/**
* - 表单项标签描述
 */
func (a *InputDatetimeRange) LabelRemark(value interface{}) *InputDatetimeRange {
	a.Set("labelRemark", value)
	return a
}

/**
* - 限制最大日期时间，用法同 限制范围
 */
func (a *InputDatetimeRange) MaxDate(value interface{}) *InputDatetimeRange {
	a.Set("maxDate", value)
	return a
}

/**
* - 限制最小日期时间，用法同 限制范围
 */
func (a *InputDatetimeRange) MinDate(value interface{}) *InputDatetimeRange {
	a.Set("minDate", value)
	return a
}

/**
* - 字段名，指定该表单项提交时的 key
 */
func (a *InputDatetimeRange) Name(value interface{}) *InputDatetimeRange {
	a.Set("name", value)
	return a
}

/**
* - 表单项描述
 */
func (a *InputDatetimeRange) Placeholder(value interface{}) *InputDatetimeRange {
	a.Set("placeholder", value)
	return a
}

/**
* - 日期范围快捷键
 */
func (a *InputDatetimeRange) Ranges(value interface{}) *InputDatetimeRange {
	a.Set("ranges", value)
	return a
}

/**
* - 是否必填
 */
func (a *InputDatetimeRange) Required(value bool) *InputDatetimeRange {
	a.Set("required", value)
	return a
}

/**
* - 通过表达式来配置当前表单项是否为必填。
 */
func (a *InputDatetimeRange) RequiredOn(value interface{}) *InputDatetimeRange {
	a.Set("requiredOn", value)
	return a
}

/**
* - 是否该表单项值发生变化时就提交当前表单。
 */
func (a *InputDatetimeRange) SubmitOnChange(value bool) *InputDatetimeRange {
	a.Set("submitOnChange", value)
	return a
}

/**
* - 指定为 input-datetime-range 渲染器。
 */
func (a *InputDatetimeRange) Type(value interface{}) *InputDatetimeRange {
	a.Set("type", value)
	return a
}

/**
* - 保存 UTC 值
 */
func (a *InputDatetimeRange) Utc(value bool) *InputDatetimeRange {
	a.Set("utc", value)
	return a
}

/**
* - 表单校验接口
 */
func (a *InputDatetimeRange) ValidateApi(value interface{}) *InputDatetimeRange {
	a.Set("validateApi", value)
	return a
}

/**
* - 表单项值格式验证，支持设置多个，多个规则用英文逗号隔开。
 */
func (a *InputDatetimeRange) Validations(value interface{}) *InputDatetimeRange {
	a.Set("validations", value)
	return a
}

/**
* - 表单默认值
 */
func (a *InputDatetimeRange) Value(value interface{}) *InputDatetimeRange {
	a.Set("value", value)
	return a
}

/**
* - 是否显示
 */
func (a *InputDatetimeRange) Visible(value bool) *InputDatetimeRange {
	a.Set("visible", value)
	return a
}

/* - 当前表单项是否禁用的条件
 */
func (a *InputDatetimeRange) VisibleOn(value interface{}) *InputDatetimeRange {
	a.Set("visibleOn", value)
	return a
}
