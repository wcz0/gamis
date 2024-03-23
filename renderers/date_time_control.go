package renderers


/**
 * Datetime日期时间选择控件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/datetime
 *

*/
type DateTimeControl struct {
	*BaseRenderer
}

func NewDateTimeControl() *DateTimeControl {
    a := &DateTimeControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-datetime")
    return a
}
/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *DateTimeControl) ValidateOnChange(value string) *DateTimeControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 配置 input className
 */
func (a *DateTimeControl) InputClassName(value string) *DateTimeControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 占位符
 */
func (a *DateTimeControl) Placeholder(value string) *DateTimeControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 日期存储格式
 */
func (a *DateTimeControl) Format(value string) *DateTimeControl {
    a.Set("format", value)
    return a
}

/**
 * 时间的格式。
 */
func (a *DateTimeControl) TimeFormat(value string) *DateTimeControl {
    a.Set("timeFormat", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *DateTimeControl) Size(value string) *DateTimeControl {
    a.Set("size", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *DateTimeControl) Name(value string) *DateTimeControl {
    a.Set("name", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *DateTimeControl) ExtraName(value string) *DateTimeControl {
    a.Set("extraName", value)
    return a
}

/**
 * 限制最小日期
 */
func (a *DateTimeControl) MinDate(value string) *DateTimeControl {
    a.Set("minDate", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *DateTimeControl) DescriptionClassName(value string) *DateTimeControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 */
func (a *DateTimeControl) StaticSchema(value string) *DateTimeControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 描述标题
 */
func (a *DateTimeControl) Label(value string) *DateTimeControl {
    a.Set("label", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *DateTimeControl) LabelWidth(value string) *DateTimeControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *DateTimeControl) Inline(value string) *DateTimeControl {
    a.Set("inline", value)
    return a
}

/**
 * 日期展示格式(新：替代inputFormat)
 */
func (a *DateTimeControl) DisplayFormat(value string) *DateTimeControl {
    a.Set("displayFormat", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *DateTimeControl) ClassName(value string) *DateTimeControl {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用
 */
func (a *DateTimeControl) Disabled(value string) *DateTimeControl {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *DateTimeControl) Hidden(value string) *DateTimeControl {
    a.Set("hidden", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *DateTimeControl) Static(value string) *DateTimeControl {
    a.Set("static", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *DateTimeControl) Hint(value string) *DateTimeControl {
    a.Set("hint", value)
    return a
}

/**
 * 是否为必填
 */
func (a *DateTimeControl) Required(value string) *DateTimeControl {
    a.Set("required", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *DateTimeControl) HiddenOn(value string) *DateTimeControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *DateTimeControl) StaticClassName(value string) *DateTimeControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 日期展示格式
 */
func (a *DateTimeControl) InputFormat(value string) *DateTimeControl {
    a.Set("inputFormat", value)
    return a
}

/**
 * 字符串函数，用来决定是否禁用某个日期。(currentDate: moment.Moment, props: any) => boolean;
 */
func (a *DateTimeControl) DisabledDate(value string) *DateTimeControl {
    a.Set("disabledDate", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *DateTimeControl) EditorSetting(value string) *DateTimeControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *DateTimeControl) SubmitOnChange(value string) *DateTimeControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 日期快捷键
 */
func (a *DateTimeControl) Shortcuts(value string) *DateTimeControl {
    a.Set("shortcuts", value)
    return a
}

/**
 * 配置 label className
 */
func (a *DateTimeControl) LabelClassName(value string) *DateTimeControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 只读条件
 */
func (a *DateTimeControl) ReadOnlyOn(value string) *DateTimeControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 是否为内联模式？
 */
func (a *DateTimeControl) Emebed(value string) *DateTimeControl {
    a.Set("emebed", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *DateTimeControl) OnEvent(value string) *DateTimeControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 描述标题
 */
func (a *DateTimeControl) LabelAlign(value string) *DateTimeControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *DateTimeControl) Remark(value string) *DateTimeControl {
    a.Set("remark", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *DateTimeControl) LabelRemark(value string) *DateTimeControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *DateTimeControl) ClearValueOnHidden(value string) *DateTimeControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 设定是否存储 utc 时间。
 */
func (a *DateTimeControl) Utc(value string) *DateTimeControl {
    a.Set("utc", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *DateTimeControl) Width(value string) *DateTimeControl {
    a.Set("width", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *DateTimeControl) Id(value string) *DateTimeControl {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *DateTimeControl) StaticInputClassName(value string) *DateTimeControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否只读
 */
func (a *DateTimeControl) ReadOnly(value string) *DateTimeControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *DateTimeControl) Horizontal(value string) *DateTimeControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 时间输入范围限制
 */
func (a *DateTimeControl) TimeConstraints(value string) *DateTimeControl {
    a.Set("timeConstraints", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *DateTimeControl) StaticPlaceholder(value string) *DateTimeControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *DateTimeControl) StaticLabelClassName(value string) *DateTimeControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *DateTimeControl) UseMobileUI(value string) *DateTimeControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *DateTimeControl) Validations(value string) *DateTimeControl {
    a.Set("validations", value)
    return a
}

/**
 * 是否显示清除按钮
 */
func (a *DateTimeControl) Clearable(value string) *DateTimeControl {
    a.Set("clearable", value)
    return a
}

/**
 * 限制最大日期
 */
func (a *DateTimeControl) MaxDate(value string) *DateTimeControl {
    a.Set("maxDate", value)
    return a
}

/**
 * 是否为结束时间，如果是，那么会自动加上 23:59:59
 */
func (a *DateTimeControl) IsEndDate(value string) *DateTimeControl {
    a.Set("isEndDate", value)
    return a
}

/**
 * 是否显示
 */
func (a *DateTimeControl) Visible(value string) *DateTimeControl {
    a.Set("visible", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *DateTimeControl) StaticOn(value string) *DateTimeControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 组件样式
 */
func (a *DateTimeControl) Style(value string) *DateTimeControl {
    a.Set("style", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *DateTimeControl) Description(value string) *DateTimeControl {
    a.Set("description", value)
    return a
}

/**
 */
func (a *DateTimeControl) Desc(value string) *DateTimeControl {
    a.Set("desc", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *DateTimeControl) ValidateApi(value string) *DateTimeControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *DateTimeControl) DisabledOn(value string) *DateTimeControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *DateTimeControl) VisibleOn(value string) *DateTimeControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *DateTimeControl) Value(value string) *DateTimeControl {
    a.Set("value", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *DateTimeControl) ValidationErrors(value string) *DateTimeControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 指定为日期时间选择控件
 */
func (a *DateTimeControl) Type(value string) *DateTimeControl {
    a.Set("type", value)
    return a
}

/**
 * 替代format
 */
func (a *DateTimeControl) ValueFormat(value string) *DateTimeControl {
    a.Set("valueFormat", value)
    return a
}

/**
 * 边框模式，全边框，还是半边框，或者没边框。
 * 可选值: full | half | none
 */
func (a *DateTimeControl) BorderMode(value string) *DateTimeControl {
    a.Set("borderMode", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *DateTimeControl) Mode(value string) *DateTimeControl {
    a.Set("mode", value)
    return a
}
