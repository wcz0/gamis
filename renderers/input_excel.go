package renderers

type InputExcel struct {
	*BaseRenderer
}

func NewInputExcel() *InputExcel {
	i := &InputExcel{
		BaseRenderer: NewBaseRenderer(),
	}
	i.Set("type", "input-excel")
	return i
}

func (i *InputExcel) Set(name string, value interface{}) *InputExcel {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	i.AmisSchema[name] = value
	return i
}

/**
 * 是否解析所有 sheet
 */
func (i *InputExcel) AllSheets(value bool) *InputExcel {
	i.Set("allSheets", value)
	return i
}

/**
 * 数据录入配置，自动填充或者参照录入
 */
func (i *InputExcel) AutoFill(value interface{}) *InputExcel {
	i.Set("autoFill", value)
	return i
}

/**
 * 表单最外层类名
 */
func (i *InputExcel) ClassName(value interface{}) *InputExcel {
	i.Set("className", value)
	return i
}

/**
 * 表单项描述
 */
func (i *InputExcel) Description(value interface{}) *InputExcel {
	i.Set("description", value)
	return i
}

/**
 * 是否禁用
 */
func (i *InputExcel) Disabled(value bool) *InputExcel {
	i.Set("disabled", value)
	return i
}

/**
 * 当前表单项是否禁用的条件
 */
func (i *InputExcel) DisabledOn(value interface{}) *InputExcel {
	i.Set("disabledOn", value)
	return i
}

/**
 * 是否解析所有 sheet
 */
func (i *InputExcel) IncludeEmpty(value bool) *InputExcel {
	i.Set("includeEmpty", value)
	return i
}

/**
 * 是否内联
 */
func (i *InputExcel) Inline(value bool) *InputExcel {
	i.Set("inline", value)
	return i
}

/**
 * 表单控制器类名
 */
func (i *InputExcel) InputClassName(value interface{}) *InputExcel {
	i.Set("inputClassName", value)
	return i
}

/**
 * 表单项标签
 */
func (i *InputExcel) Label(value interface{}) *InputExcel {
	i.Set("label", value)
	return i
}

/**
 * 表单项标签对齐方式，默认右对齐，仅在 mode为horizontal 时生效
 */
func (i *InputExcel) LabelAlign(value interface{}) *InputExcel {
	i.Set("labelAlign", value)
	return i
}

/**
 * label 的类名
 */
func (i *InputExcel) LabelClassName(value interface{}) *InputExcel {
	i.Set("labelClassName", value)
	return i
}

/**
 * 表单项标签描述
 */
func (i *InputExcel) LabelRemark(value interface{}) *InputExcel {
	i.Set("labelRemark", value)
	return i
}

/**
 * 字段名，指定该表单项提交时的 key
 */
func (i *InputExcel) Name(value interface{}) *InputExcel {
	i.Set("name", value)
	return i
}

/**
 * 解析模式
 */
func (i *InputExcel) ParseMode(value interface{}) *InputExcel {
	i.Set("parseMode", value)
	return i
}

/**
 * 表单项描述
 */
func (i *InputExcel) Placeholder(value interface{}) *InputExcel {
	i.Set("placeholder", value)
	return i
}

/**
 * 是否解析为纯文本
 */
func (i *InputExcel) PlainText(value bool) *InputExcel {
	i.Set("plainText", value)
	return i
}

/**
 * 是否必填
 */
func (i *InputExcel) Required(value bool) *InputExcel {
	i.Set("required", value)
	return i
}

/**
 * 通过表达式来配置当前表单项是否为必填。
 */
func (i *InputExcel) RequiredOn(value interface{}) *InputExcel {
	i.Set("requiredOn", value)
	return i
}

/**
 * 是否该表单项值发生变化时就提交当前表单。
 */
func (i *InputExcel) SubmitOnChange(value bool) *InputExcel {
	i.Set("submitOnChange", value)
	return i
}

/**
 * 指定为 input-excel 渲染器。
 */
func (i *InputExcel) Type(value interface{}) *InputExcel {
	i.Set("type", value)
	return i
}

/**
 * 表单校验接口
 */
func (i *InputExcel) ValidateApi(value interface{}) *InputExcel {
	i.Set("validateApi", value)
	return i
}

/**
 * 表单项值格式验证，支持设置多个，多个规则用英文逗号隔开。
 */
func (i *InputExcel) Validations(value interface{}) *InputExcel {
	i.Set("validations", value)
	return i
}

/**
 * 表单默认值
 */
func (i *InputExcel) Value(value interface{}) *InputExcel {
	i.Set("value", value)
	return i
}

/**
 * 是否可见
 */
func (i *InputExcel) Visible(value bool) *InputExcel {
	i.Set("visible", value)
	return i
}

/**
 * 当前表单项是否禁用的条件
 */
func (i *InputExcel) VisibleOn(value interface{}) *InputExcel {
	i.Set("visibleOn", value)
	return i
}
