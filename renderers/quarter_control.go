package renderers


/**
 * 季度选择控件
 *

*/
type QuarterControl struct {
	*BaseRenderer
}

func NewQuarterControl() *QuarterControl {
    a := &QuarterControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-quarter")
    return a
}
/**
 * 描述标题
 */
func (a *QuarterControl) Label(value string) *QuarterControl {
    a.Set("label", value)
    return a
}

/**
 */
func (a *QuarterControl) Desc(value string) *QuarterControl {
    a.Set("desc", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *QuarterControl) DescriptionClassName(value string) *QuarterControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *QuarterControl) HiddenOn(value string) *QuarterControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示清除按钮
 */
func (a *QuarterControl) Clearable(value string) *QuarterControl {
    a.Set("clearable", value)
    return a
}

/**
 * 边框模式，全边框，还是半边框，或者没边框。
 * 可选值: full | half | none
 */
func (a *QuarterControl) BorderMode(value string) *QuarterControl {
    a.Set("borderMode", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *QuarterControl) Width(value string) *QuarterControl {
    a.Set("width", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *QuarterControl) SubmitOnChange(value string) *QuarterControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *QuarterControl) Mode(value string) *QuarterControl {
    a.Set("mode", value)
    return a
}

/**
 * 配置 input className
 */
func (a *QuarterControl) InputClassName(value string) *QuarterControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 是否为必填
 */
func (a *QuarterControl) Required(value string) *QuarterControl {
    a.Set("required", value)
    return a
}

/**
 * 是否显示
 */
func (a *QuarterControl) Visible(value string) *QuarterControl {
    a.Set("visible", value)
    return a
}

/**
 */
func (a *QuarterControl) Validations(value string) *QuarterControl {
    a.Set("validations", value)
    return a
}

/**
 * 替代format
 */
func (a *QuarterControl) ValueFormat(value string) *QuarterControl {
    a.Set("valueFormat", value)
    return a
}

/**
 * 是否为内联模式？
 */
func (a *QuarterControl) Emebed(value string) *QuarterControl {
    a.Set("emebed", value)
    return a
}

/**
 * 字符串函数，用来决定是否禁用某个日期。(currentDate: moment.Moment, props: any) => boolean;
 */
func (a *QuarterControl) DisabledDate(value string) *QuarterControl {
    a.Set("disabledDate", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *QuarterControl) StaticLabelClassName(value string) *QuarterControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *QuarterControl) Hint(value string) *QuarterControl {
    a.Set("hint", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *QuarterControl) ClearValueOnHidden(value string) *QuarterControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 配置 label className
 */
func (a *QuarterControl) LabelClassName(value string) *QuarterControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *QuarterControl) StaticClassName(value string) *QuarterControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *QuarterControl) Style(value string) *QuarterControl {
    a.Set("style", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *QuarterControl) EditorSetting(value string) *QuarterControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *QuarterControl) ExtraName(value string) *QuarterControl {
    a.Set("extraName", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *QuarterControl) Remark(value string) *QuarterControl {
    a.Set("remark", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *QuarterControl) Inline(value string) *QuarterControl {
    a.Set("inline", value)
    return a
}

/**
 * 日期快捷键
 */
func (a *QuarterControl) Shortcuts(value string) *QuarterControl {
    a.Set("shortcuts", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *QuarterControl) StaticPlaceholder(value string) *QuarterControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 是否只读
 */
func (a *QuarterControl) ReadOnly(value string) *QuarterControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 只读条件
 */
func (a *QuarterControl) ReadOnlyOn(value string) *QuarterControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *QuarterControl) Horizontal(value string) *QuarterControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 设定是否存储 utc 时间。
 */
func (a *QuarterControl) Utc(value string) *QuarterControl {
    a.Set("utc", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *QuarterControl) StaticOn(value string) *QuarterControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *QuarterControl) ValidationErrors(value string) *QuarterControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *QuarterControl) LabelRemark(value string) *QuarterControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *QuarterControl) Size(value string) *QuarterControl {
    a.Set("size", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *QuarterControl) ValidateOnChange(value string) *QuarterControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *QuarterControl) Description(value string) *QuarterControl {
    a.Set("description", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *QuarterControl) DisabledOn(value string) *QuarterControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *QuarterControl) Name(value string) *QuarterControl {
    a.Set("name", value)
    return a
}

/**
 * 占位符
 */
func (a *QuarterControl) Placeholder(value string) *QuarterControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 月份展示格式
 */
func (a *QuarterControl) InputFormat(value string) *QuarterControl {
    a.Set("inputFormat", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *QuarterControl) LabelWidth(value string) *QuarterControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *QuarterControl) Id(value string) *QuarterControl {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *QuarterControl) StaticInputClassName(value string) *QuarterControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *QuarterControl) StaticSchema(value string) *QuarterControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *QuarterControl) Value(value string) *QuarterControl {
    a.Set("value", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *QuarterControl) ValidateApi(value string) *QuarterControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *QuarterControl) Hidden(value string) *QuarterControl {
    a.Set("hidden", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *QuarterControl) Static(value string) *QuarterControl {
    a.Set("static", value)
    return a
}

/**
 * 月份存储格式
 */
func (a *QuarterControl) Format(value string) *QuarterControl {
    a.Set("format", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *QuarterControl) OnEvent(value string) *QuarterControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否禁用
 */
func (a *QuarterControl) Disabled(value string) *QuarterControl {
    a.Set("disabled", value)
    return a
}

/**
 * 日期展示格式(新：替代inputFormat)
 */
func (a *QuarterControl) DisplayFormat(value string) *QuarterControl {
    a.Set("displayFormat", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *QuarterControl) VisibleOn(value string) *QuarterControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *QuarterControl) UseMobileUI(value string) *QuarterControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 描述标题
 */
func (a *QuarterControl) LabelAlign(value string) *QuarterControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 指定为月份时间选择控件
 */
func (a *QuarterControl) Type(value string) *QuarterControl {
    a.Set("type", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *QuarterControl) ClassName(value string) *QuarterControl {
    a.Set("className", value)
    return a
}
