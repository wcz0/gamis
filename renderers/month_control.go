package renderers


/**
 * Month 月份选择控件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/Month
 *

*/
type MonthControl struct {
	*BaseRenderer
}

func NewMonthControl() *MonthControl {
    a := &MonthControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-month")
    return a
}
/**
 * 静态展示表单项Label类名
 */
func (a *MonthControl) StaticLabelClassName(value string) *MonthControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *MonthControl) Width(value string) *MonthControl {
    a.Set("width", value)
    return a
}

/**
 */
func (a *MonthControl) Desc(value string) *MonthControl {
    a.Set("desc", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *MonthControl) Horizontal(value string) *MonthControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *MonthControl) DisabledOn(value string) *MonthControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *MonthControl) Id(value string) *MonthControl {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *MonthControl) StaticClassName(value string) *MonthControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *MonthControl) LabelWidth(value string) *MonthControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *MonthControl) StaticPlaceholder(value string) *MonthControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 是否为必填
 */
func (a *MonthControl) Required(value string) *MonthControl {
    a.Set("required", value)
    return a
}

/**
 * 配置 label className
 */
func (a *MonthControl) LabelClassName(value string) *MonthControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 只读条件
 */
func (a *MonthControl) ReadOnlyOn(value string) *MonthControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 日期展示格式(新：替代inputFormat)
 */
func (a *MonthControl) DisplayFormat(value string) *MonthControl {
    a.Set("displayFormat", value)
    return a
}

/**
 * 是否为内联模式？
 */
func (a *MonthControl) Emebed(value string) *MonthControl {
    a.Set("emebed", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *MonthControl) Hidden(value string) *MonthControl {
    a.Set("hidden", value)
    return a
}

/**
 * 占位符
 */
func (a *MonthControl) Placeholder(value string) *MonthControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 设定是否存储 utc 时间。
 */
func (a *MonthControl) Utc(value string) *MonthControl {
    a.Set("utc", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *MonthControl) Name(value string) *MonthControl {
    a.Set("name", value)
    return a
}

/**
 */
func (a *MonthControl) Validations(value string) *MonthControl {
    a.Set("validations", value)
    return a
}

/**
 * 指定为月份时间选择控件
 */
func (a *MonthControl) Type(value string) *MonthControl {
    a.Set("type", value)
    return a
}

/**
 * 月份存储格式
 */
func (a *MonthControl) Format(value string) *MonthControl {
    a.Set("format", value)
    return a
}

/**
 * 字符串函数，用来决定是否禁用某个日期。(currentDate: moment.Moment, props: any) => boolean;
 */
func (a *MonthControl) DisabledDate(value string) *MonthControl {
    a.Set("disabledDate", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *MonthControl) Static(value string) *MonthControl {
    a.Set("static", value)
    return a
}

/**
 * 组件样式
 */
func (a *MonthControl) Style(value string) *MonthControl {
    a.Set("style", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *MonthControl) Description(value string) *MonthControl {
    a.Set("description", value)
    return a
}

/**
 * 是否显示
 */
func (a *MonthControl) Visible(value string) *MonthControl {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *MonthControl) StaticInputClassName(value string) *MonthControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 配置 input className
 */
func (a *MonthControl) InputClassName(value string) *MonthControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 边框模式，全边框，还是半边框，或者没边框。
 * 可选值: full | half | none
 */
func (a *MonthControl) BorderMode(value string) *MonthControl {
    a.Set("borderMode", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *MonthControl) Value(value string) *MonthControl {
    a.Set("value", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *MonthControl) HiddenOn(value string) *MonthControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *MonthControl) VisibleOn(value string) *MonthControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *MonthControl) OnEvent(value string) *MonthControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *MonthControl) Remark(value string) *MonthControl {
    a.Set("remark", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *MonthControl) LabelRemark(value string) *MonthControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *MonthControl) ClassName(value string) *MonthControl {
    a.Set("className", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *MonthControl) DescriptionClassName(value string) *MonthControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *MonthControl) Inline(value string) *MonthControl {
    a.Set("inline", value)
    return a
}

/**
 */
func (a *MonthControl) StaticSchema(value string) *MonthControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 描述标题
 */
func (a *MonthControl) Label(value string) *MonthControl {
    a.Set("label", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *MonthControl) ValidationErrors(value string) *MonthControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 替代format
 */
func (a *MonthControl) ValueFormat(value string) *MonthControl {
    a.Set("valueFormat", value)
    return a
}

/**
 * 月份展示格式
 */
func (a *MonthControl) InputFormat(value string) *MonthControl {
    a.Set("inputFormat", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *MonthControl) ExtraName(value string) *MonthControl {
    a.Set("extraName", value)
    return a
}

/**
 * 是否只读
 */
func (a *MonthControl) ReadOnly(value string) *MonthControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 描述标题
 */
func (a *MonthControl) LabelAlign(value string) *MonthControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *MonthControl) Hint(value string) *MonthControl {
    a.Set("hint", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *MonthControl) Mode(value string) *MonthControl {
    a.Set("mode", value)
    return a
}

/**
 * 是否显示清除按钮
 */
func (a *MonthControl) Clearable(value string) *MonthControl {
    a.Set("clearable", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *MonthControl) StaticOn(value string) *MonthControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *MonthControl) ClearValueOnHidden(value string) *MonthControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *MonthControl) EditorSetting(value string) *MonthControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *MonthControl) UseMobileUI(value string) *MonthControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *MonthControl) Size(value string) *MonthControl {
    a.Set("size", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *MonthControl) SubmitOnChange(value string) *MonthControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *MonthControl) ValidateOnChange(value string) *MonthControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 是否禁用
 */
func (a *MonthControl) Disabled(value string) *MonthControl {
    a.Set("disabled", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *MonthControl) ValidateApi(value string) *MonthControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 日期快捷键
 */
func (a *MonthControl) Shortcuts(value string) *MonthControl {
    a.Set("shortcuts", value)
    return a
}
