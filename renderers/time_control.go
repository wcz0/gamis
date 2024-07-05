package renderers


/**
 * Time 时间选择控件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/time

 * @author wcz0
 * @version 6.2.2
 */
type TimeControl struct {
	*BaseRenderer
}

func NewTimeControl() *TimeControl {
    a := &TimeControl{
        BaseRenderer: NewBaseRenderer(),
    }
    a.Set("type", "input-time")
    return a
}

/**
 * 容器 css 类名
 */
func (a *TimeControl) ClassName(value interface{}) *TimeControl {
    a.Set("className", value)
    return a
}

/**
 */
func (a *TimeControl) Validations(value interface{}) *TimeControl {
    a.Set("validations", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *TimeControl) Horizontal(value interface{}) *TimeControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *TimeControl) ClearValueOnHidden(value interface{}) *TimeControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *TimeControl) Width(value interface{}) *TimeControl {
    a.Set("width", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *TimeControl) LabelRemark(value interface{}) *TimeControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 是否只读
 */
func (a *TimeControl) ReadOnly(value interface{}) *TimeControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *TimeControl) Description(value interface{}) *TimeControl {
    a.Set("description", value)
    return a
}

/**
 */
func (a *TimeControl) Desc(value interface{}) *TimeControl {
    a.Set("desc", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *TimeControl) DescriptionClassName(value interface{}) *TimeControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 指定为日期时间选择控件
 */
func (a *TimeControl) Type(value interface{}) *TimeControl {
    a.Set("type", value)
    return a
}

/**
 * 是否禁用
 */
func (a *TimeControl) Disabled(value interface{}) *TimeControl {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *TimeControl) Hidden(value interface{}) *TimeControl {
    a.Set("hidden", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *TimeControl) UseMobileUI(value interface{}) *TimeControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *TimeControl) StaticSchema(value interface{}) *TimeControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 配置 input className
 */
func (a *TimeControl) InputClassName(value interface{}) *TimeControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *TimeControl) Value(value interface{}) *TimeControl {
    a.Set("value", value)
    return a
}

/**
 * 日期快捷键
 */
func (a *TimeControl) Shortcuts(value interface{}) *TimeControl {
    a.Set("shortcuts", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *TimeControl) DisabledOn(value interface{}) *TimeControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *TimeControl) StaticPlaceholder(value interface{}) *TimeControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *TimeControl) ValidateOnChange(value interface{}) *TimeControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *TimeControl) SubmitOnChange(value interface{}) *TimeControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *TimeControl) AutoFill(value interface{}) *TimeControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 日期展示格式(新：替代inputFormat)
 */
func (a *TimeControl) DisplayFormat(value interface{}) *TimeControl {
    a.Set("displayFormat", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *TimeControl) StaticLabelClassName(value interface{}) *TimeControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 描述标题
 */
func (a *TimeControl) Label(value interface{}) *TimeControl {
    a.Set("label", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *TimeControl) LabelWidth(value interface{}) *TimeControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *TimeControl) Static(value interface{}) *TimeControl {
    a.Set("static", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *TimeControl) StaticOn(value interface{}) *TimeControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *TimeControl) StaticClassName(value interface{}) *TimeControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *TimeControl) ValidateApi(value interface{}) *TimeControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 描述标题
 */
func (a *TimeControl) LabelAlign(value interface{}) *TimeControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 配置 label className
 */
func (a *TimeControl) LabelClassName(value interface{}) *TimeControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 占位符
 */
func (a *TimeControl) Placeholder(value interface{}) *TimeControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 边框模式，全边框，还是半边框，或者没边框。
 * 可选值: full | half | none
 */
func (a *TimeControl) BorderMode(value interface{}) *TimeControl {
    a.Set("borderMode", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *TimeControl) Id(value interface{}) *TimeControl {
    a.Set("id", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *TimeControl) ExtraName(value interface{}) *TimeControl {
    a.Set("extraName", value)
    return a
}

/**
 * 日期存储格式
 */
func (a *TimeControl) Format(value interface{}) *TimeControl {
    a.Set("format", value)
    return a
}

/**
 * 是否显示
 */
func (a *TimeControl) Visible(value interface{}) *TimeControl {
    a.Set("visible", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *TimeControl) Size(value interface{}) *TimeControl {
    a.Set("size", value)
    return a
}

/**
 * 设定是否存储 utc 时间。
 */
func (a *TimeControl) Utc(value interface{}) *TimeControl {
    a.Set("utc", value)
    return a
}

/**
 * 替代format
 */
func (a *TimeControl) ValueFormat(value interface{}) *TimeControl {
    a.Set("valueFormat", value)
    return a
}

/**
 * 是否为内联模式？
 */
func (a *TimeControl) Emebed(value interface{}) *TimeControl {
    a.Set("emebed", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *TimeControl) VisibleOn(value interface{}) *TimeControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *TimeControl) Remark(value interface{}) *TimeControl {
    a.Set("remark", value)
    return a
}

/**
 * 组件样式
 */
func (a *TimeControl) Style(value interface{}) *TimeControl {
    a.Set("style", value)
    return a
}

/**
 */
func (a *TimeControl) Row(value interface{}) *TimeControl {
    a.Set("row", value)
    return a
}

/**
 * 是否显示清除按钮
 */
func (a *TimeControl) Clearable(value interface{}) *TimeControl {
    a.Set("clearable", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *TimeControl) HiddenOn(value interface{}) *TimeControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 日期展示格式
 */
func (a *TimeControl) InputFormat(value interface{}) *TimeControl {
    a.Set("inputFormat", value)
    return a
}

/**
 * 只读条件
 */
func (a *TimeControl) ReadOnlyOn(value interface{}) *TimeControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *TimeControl) ValidationErrors(value interface{}) *TimeControl {
    a.Set("validationErrors", value)
    return a
}

/**
 */
func (a *TimeControl) InitAutoFill(value interface{}) *TimeControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 时间输入范围限制
 */
func (a *TimeControl) TimeConstraints(value interface{}) *TimeControl {
    a.Set("timeConstraints", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *TimeControl) OnEvent(value interface{}) *TimeControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *TimeControl) Name(value interface{}) *TimeControl {
    a.Set("name", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *TimeControl) Hint(value interface{}) *TimeControl {
    a.Set("hint", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *TimeControl) Mode(value interface{}) *TimeControl {
    a.Set("mode", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *TimeControl) Inline(value interface{}) *TimeControl {
    a.Set("inline", value)
    return a
}

/**
 * 是否为必填
 */
func (a *TimeControl) Required(value interface{}) *TimeControl {
    a.Set("required", value)
    return a
}

/**
 * 字符串函数，用来决定是否禁用某个日期。(currentDate: moment.Moment, props: any) => boolean;
 */
func (a *TimeControl) DisabledDate(value interface{}) *TimeControl {
    a.Set("disabledDate", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *TimeControl) StaticInputClassName(value interface{}) *TimeControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *TimeControl) EditorSetting(value interface{}) *TimeControl {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *TimeControl) TestIdBuilder(value interface{}) *TimeControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 时间的格式。
 */
func (a *TimeControl) TimeFormat(value interface{}) *TimeControl {
    a.Set("timeFormat", value)
    return a
}
