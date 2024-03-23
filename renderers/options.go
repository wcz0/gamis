package renderers

type Options struct {
	*BaseRenderer
}

func NewOptions() *Options {
	o := &Options{
		BaseRenderer: NewBaseRenderer(),
	}
	o.Set("type", "options")
	return o
}

/**
 * 数据录入配置，自动填充或者参照录入
 */
func (o *Options) AutoFill(value string) *Options {
	o.Set("autoFill", value)
	return o
}

/**
 * 表单最外层类名
 */
func (o *Options) ClassName(value string) *Options {
	o.Set("className", value)
	return o
}

/**
 * 表单项描述
 */
func (o *Options) Description(value string) *Options {
	o.Set("description", value)
	return o
}

/**
 * 是否禁用
 */
func (o *Options) Disabled(value bool) *Options {
	o.Set("disabled", value)
	return o
}

/**
 * 当前表单项是否禁用的条件
 */
func (o *Options) DisabledOn(value string) *Options {
	o.Set("disabledOn", value)
	return o
}

/**
 * 是否把选中的值从数组中提取出来，只有当 joinValues 为 true 时生效
 */
func (o *Options) ExtractValue(value bool) *Options {
	o.Set("extractValue", value)
	return o
}

/**
 * 是否内联
 */
func (o *Options) Inline(value bool) *Options {
	o.Set("inline", value)
	return o
}

/**
 * 表单控制器类名
 */
func (o *Options) InputClassName(value string) *Options {
	o.Set("inputClassName", value)
	return o
}

/**
 * 每个选项的高度，用于虚拟渲染
 */
func (o *Options) ItemHeight(value string) *Options {
	o.Set("itemHeight", value)
	return o
}

/**
 * 多选时，是否把选中的值用逗号拼接成字符串
 */
func (o *Options) JoinValues(value bool) *Options {
	o.Set("joinValues", value)
	return o
}

/**
 * 表单项标签
 */
func (o *Options) Label(value string) *Options {
	o.Set("label", value)
	return o
}

/**
 * 表单项标签对齐方式，默认右对齐，仅在 mode为horizontal 时生效
 */
func (o *Options) LabelAlign(value string) *Options {
	o.Set("labelAlign", value)
	return o
}

/**
 * label 的类名
 */
func (o *Options) LabelClassName(value string) *Options {
	o.Set("labelClassName", value)
	return o
}

/**
 * 标识选项中哪个字段是label值
 */
func (o *Options) LabelField(value string) *Options {
	o.Set("labelField", value)
	return o
}

/**
 * 表单项标签描述
 */
func (o *Options) LabelRemark(value string) *Options {
	o.Set("labelRemark", value)
	return o
}

/**
 * 是否支持多选
 */
func (o *Options) Multiple(value bool) *Options {
	o.Set("multiple", value)
	return o
}

/**
 * 字段名，指定该表单项提交时的 key
 */
func (o *Options) Name(value string) *Options {
	o.Set("name", value)
	return o
}

/**
 * 选项组，供用户选择
 */
func (o *Options) Options(value string) *Options {
	o.Set("options", value)
	return o
}

/**
 * 是否必填
 */
func (o *Options) Required(value bool) *Options {
	o.Set("required", value)
	return o
}

/**
 * 通过表达式来配置当前表单项是否为必填。
 */
func (o *Options) RequiredOn(value string) *Options {
	o.Set("requiredOn", value)
	return o
}

/**
 * 选项组源，可通过数据映射获取当前数据域变量、或者配置 API 对象
 */
func (o *Options) Source(value string) *Options {
	o.Set("source", value)
	return o
}

/**
 * 是否该表单项值发生变化时就提交当前表单。
 */
func (o *Options) SubmitOnChange(value bool) *Options {
	o.Set("submitOnChange", value)
	return o
}

/**
 * 表单校验接口
 */
func (o *Options) ValidateApi(value string) *Options {
	o.Set("validateApi", value)
	return o
}

/**
 * 表单项值格式验证，支持设置多个，多个规则用英文逗号隔开。
 */
func (o *Options) Validations(value string) *Options {
	o.Set("validations", value)
	return o
}

/**
 * 表单默认值
 */
func (o *Options) Value(value string) *Options {
	o.Set("value", value)
	return o
}

/**
 * 标识选项中哪个字段是value值
 */
func (o *Options) ValueField(value string) *Options {
	o.Set("valueField", value)
	return o
}

/**
 * 默认情况下多选所有选项都会显示，通过这个可以最多显示一行，超出的部分变成 ...
 */
func (o *Options) ValuesNoWrap(value bool) *Options {
	o.Set("valuesNoWrap", value)
	return o
}

/**
 * 在选项数量超过多少时开启虚拟渲染
 */
func (o *Options) VirtualThreshold(value string) *Options {
	o.Set("virtualThreshold", value)
	return o
}

/**
 * 是否可见
 */
func (o *Options) Visible(value bool) *Options {
	o.Set("visible", value)
	return o
}

/**
 * 当前表单项是否禁用的条件
 */
func (o *Options) VisibleOn(value string) *Options {
	o.Set("visibleOn", value)
	return o
}
