package renderers


/**
 * 年份选择控件
 *

*/
type YearControl struct {
	*BaseRenderer
}

func NewYearControl() *YearControl {
    a := &YearControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-year")
    return a
}
/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *YearControl) Id(value string) *YearControl {
    a.Set("id", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *YearControl) ExtraName(value string) *YearControl {
    a.Set("extraName", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *YearControl) Hidden(value string) *YearControl {
    a.Set("hidden", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *YearControl) HiddenOn(value string) *YearControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 月份展示格式
 */
func (a *YearControl) InputFormat(value string) *YearControl {
    a.Set("inputFormat", value)
    return a
}

/**
 * 边框模式，全边框，还是半边框，或者没边框。
 * 可选值: full | half | none
 */
func (a *YearControl) BorderMode(value string) *YearControl {
    a.Set("borderMode", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *YearControl) StaticPlaceholder(value string) *YearControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *YearControl) EditorSetting(value string) *YearControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 只读条件
 */
func (a *YearControl) ReadOnlyOn(value string) *YearControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 */
func (a *YearControl) StaticSchema(value string) *YearControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 是否只读
 */
func (a *YearControl) ReadOnly(value string) *YearControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 指定为月份时间选择控件
 */
func (a *YearControl) Type(value string) *YearControl {
    a.Set("type", value)
    return a
}

/**
 * 是否为内联模式？
 */
func (a *YearControl) Emebed(value string) *YearControl {
    a.Set("emebed", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *YearControl) Width(value string) *YearControl {
    a.Set("width", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *YearControl) DisabledOn(value string) *YearControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *YearControl) StaticClassName(value string) *YearControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *YearControl) StaticInputClassName(value string) *YearControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否为必填
 */
func (a *YearControl) Required(value string) *YearControl {
    a.Set("required", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *YearControl) ValidationErrors(value string) *YearControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *YearControl) Size(value string) *YearControl {
    a.Set("size", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *YearControl) LabelRemark(value string) *YearControl {
    a.Set("labelRemark", value)
    return a
}

/**
 */
func (a *YearControl) Desc(value string) *YearControl {
    a.Set("desc", value)
    return a
}

/**
 * 配置 label className
 */
func (a *YearControl) LabelClassName(value string) *YearControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *YearControl) ValidateOnChange(value string) *YearControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 */
func (a *YearControl) Validations(value string) *YearControl {
    a.Set("validations", value)
    return a
}

/**
 * 日期快捷键
 */
func (a *YearControl) Shortcuts(value string) *YearControl {
    a.Set("shortcuts", value)
    return a
}

/**
 * 描述标题
 */
func (a *YearControl) LabelAlign(value string) *YearControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 日期展示格式(新：替代inputFormat)
 */
func (a *YearControl) DisplayFormat(value string) *YearControl {
    a.Set("displayFormat", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *YearControl) ClassName(value string) *YearControl {
    a.Set("className", value)
    return a
}

/**
 * 设定是否存储 utc 时间。
 */
func (a *YearControl) Utc(value string) *YearControl {
    a.Set("utc", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *YearControl) UseMobileUI(value string) *YearControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *YearControl) LabelWidth(value string) *YearControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *YearControl) DescriptionClassName(value string) *YearControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *YearControl) Inline(value string) *YearControl {
    a.Set("inline", value)
    return a
}

/**
 * 月份存储格式
 */
func (a *YearControl) Format(value string) *YearControl {
    a.Set("format", value)
    return a
}

/**
 * 是否禁用
 */
func (a *YearControl) Disabled(value string) *YearControl {
    a.Set("disabled", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *YearControl) OnEvent(value string) *YearControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 组件样式
 */
func (a *YearControl) Style(value string) *YearControl {
    a.Set("style", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *YearControl) Mode(value string) *YearControl {
    a.Set("mode", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *YearControl) Horizontal(value string) *YearControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 配置 input className
 */
func (a *YearControl) InputClassName(value string) *YearControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *YearControl) ValidateApi(value string) *YearControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 替代format
 */
func (a *YearControl) ValueFormat(value string) *YearControl {
    a.Set("valueFormat", value)
    return a
}

/**
 * 是否显示
 */
func (a *YearControl) Visible(value string) *YearControl {
    a.Set("visible", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *YearControl) Static(value string) *YearControl {
    a.Set("static", value)
    return a
}

/**
 * 描述标题
 */
func (a *YearControl) Label(value string) *YearControl {
    a.Set("label", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *YearControl) Description(value string) *YearControl {
    a.Set("description", value)
    return a
}

/**
 * 字符串函数，用来决定是否禁用某个日期。(currentDate: moment.Moment, props: any) => boolean;
 */
func (a *YearControl) DisabledDate(value string) *YearControl {
    a.Set("disabledDate", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *YearControl) VisibleOn(value string) *YearControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *YearControl) Remark(value string) *YearControl {
    a.Set("remark", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *YearControl) SubmitOnChange(value string) *YearControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *YearControl) StaticOn(value string) *YearControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *YearControl) StaticLabelClassName(value string) *YearControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *YearControl) Name(value string) *YearControl {
    a.Set("name", value)
    return a
}

/**
 * 占位符
 */
func (a *YearControl) Placeholder(value string) *YearControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *YearControl) Value(value string) *YearControl {
    a.Set("value", value)
    return a
}

/**
 * 是否显示清除按钮
 */
func (a *YearControl) Clearable(value string) *YearControl {
    a.Set("clearable", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *YearControl) Hint(value string) *YearControl {
    a.Set("hint", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *YearControl) ClearValueOnHidden(value string) *YearControl {
    a.Set("clearValueOnHidden", value)
    return a
}
