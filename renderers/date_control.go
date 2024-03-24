package renderers


/**
 * Date日期选择控件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/date

 * @author wcz0
 * @version 6.2.2
 */
type DateControl struct {
	*BaseRenderer
}

func NewDateControl() *DateControl {
    a := &DateControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-date")
    return a
}
/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *DateControl) ClearValueOnHidden(value interface{}) *DateControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 日期展示格式
 */
func (a *DateControl) InputFormat(value interface{}) *DateControl {
    a.Set("inputFormat", value)
    return a
}

/**
 * 边框模式，全边框，还是半边框，或者没边框。
 * 可选值: full | half | none
 */
func (a *DateControl) BorderMode(value interface{}) *DateControl {
    a.Set("borderMode", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *DateControl) ClassName(value interface{}) *DateControl {
    a.Set("className", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *DateControl) StaticOn(value interface{}) *DateControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *DateControl) StaticLabelClassName(value interface{}) *DateControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 描述标题
 */
func (a *DateControl) Label(value interface{}) *DateControl {
    a.Set("label", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *DateControl) Mode(value interface{}) *DateControl {
    a.Set("mode", value)
    return a
}

/**
 * 是否为内联模式？
 */
func (a *DateControl) Emebed(value interface{}) *DateControl {
    a.Set("emebed", value)
    return a
}

/**
 * 限制最小日期
 */
func (a *DateControl) MinDate(value interface{}) *DateControl {
    a.Set("minDate", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *DateControl) OnEvent(value interface{}) *DateControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *DateControl) StaticInputClassName(value interface{}) *DateControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *DateControl) Style(value interface{}) *DateControl {
    a.Set("style", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *DateControl) DescriptionClassName(value interface{}) *DateControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 指定为日期选择控件
 */
func (a *DateControl) Type(value interface{}) *DateControl {
    a.Set("type", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *DateControl) Static(value interface{}) *DateControl {
    a.Set("static", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *DateControl) Name(value interface{}) *DateControl {
    a.Set("name", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *DateControl) Horizontal(value interface{}) *DateControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 是否显示清除按钮
 */
func (a *DateControl) Clearable(value interface{}) *DateControl {
    a.Set("clearable", value)
    return a
}

/**
 * 字符串函数，用来决定是否禁用某个日期。(currentDate: moment.Moment, props: any) => boolean;
 */
func (a *DateControl) DisabledDate(value interface{}) *DateControl {
    a.Set("disabledDate", value)
    return a
}

/**
 * 限制最大日期
 */
func (a *DateControl) MaxDate(value interface{}) *DateControl {
    a.Set("maxDate", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *DateControl) Hint(value interface{}) *DateControl {
    a.Set("hint", value)
    return a
}

/**
 * 只读条件
 */
func (a *DateControl) ReadOnlyOn(value interface{}) *DateControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 配置 input className
 */
func (a *DateControl) InputClassName(value interface{}) *DateControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 点选日期后是否关闭弹窗
 */
func (a *DateControl) CloseOnSelect(value interface{}) *DateControl {
    a.Set("closeOnSelect", value)
    return a
}

/**
 * 是否显示
 */
func (a *DateControl) Visible(value interface{}) *DateControl {
    a.Set("visible", value)
    return a
}

/**
 */
func (a *DateControl) StaticSchema(value interface{}) *DateControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *DateControl) ValidateApi(value interface{}) *DateControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 日期存储格式
 */
func (a *DateControl) Format(value interface{}) *DateControl {
    a.Set("format", value)
    return a
}

/**
 * 设定是否存储 utc 时间。
 */
func (a *DateControl) Utc(value interface{}) *DateControl {
    a.Set("utc", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *DateControl) Hidden(value interface{}) *DateControl {
    a.Set("hidden", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *DateControl) StaticPlaceholder(value interface{}) *DateControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *DateControl) StaticClassName(value interface{}) *DateControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 是否只读
 */
func (a *DateControl) ReadOnly(value interface{}) *DateControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *DateControl) ValidateOnChange(value interface{}) *DateControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *DateControl) Value(value interface{}) *DateControl {
    a.Set("value", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *DateControl) VisibleOn(value interface{}) *DateControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *DateControl) Description(value interface{}) *DateControl {
    a.Set("description", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *DateControl) ValidationErrors(value interface{}) *DateControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 是否禁用
 */
func (a *DateControl) Disabled(value interface{}) *DateControl {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *DateControl) HiddenOn(value interface{}) *DateControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *DateControl) LabelWidth(value interface{}) *DateControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *DateControl) LabelRemark(value interface{}) *DateControl {
    a.Set("labelRemark", value)
    return a
}

/**
 */
func (a *DateControl) Validations(value interface{}) *DateControl {
    a.Set("validations", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *DateControl) DisabledOn(value interface{}) *DateControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *DateControl) Remark(value interface{}) *DateControl {
    a.Set("remark", value)
    return a
}

/**
 * 占位符
 */
func (a *DateControl) Placeholder(value interface{}) *DateControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 替代format
 */
func (a *DateControl) ValueFormat(value interface{}) *DateControl {
    a.Set("valueFormat", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *DateControl) UseMobileUI(value interface{}) *DateControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *DateControl) Size(value interface{}) *DateControl {
    a.Set("size", value)
    return a
}

/**
 * 配置 label className
 */
func (a *DateControl) LabelClassName(value interface{}) *DateControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *DateControl) ExtraName(value interface{}) *DateControl {
    a.Set("extraName", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *DateControl) Inline(value interface{}) *DateControl {
    a.Set("inline", value)
    return a
}

/**
 * 是否为必填
 */
func (a *DateControl) Required(value interface{}) *DateControl {
    a.Set("required", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *DateControl) Id(value interface{}) *DateControl {
    a.Set("id", value)
    return a
}

/**
 * 描述标题
 */
func (a *DateControl) LabelAlign(value interface{}) *DateControl {
    a.Set("labelAlign", value)
    return a
}

/**
 */
func (a *DateControl) Desc(value interface{}) *DateControl {
    a.Set("desc", value)
    return a
}

/**
 * 日期展示格式(新：替代inputFormat)
 */
func (a *DateControl) DisplayFormat(value interface{}) *DateControl {
    a.Set("displayFormat", value)
    return a
}

/**
 * 日期快捷键
 */
func (a *DateControl) Shortcuts(value interface{}) *DateControl {
    a.Set("shortcuts", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *DateControl) Width(value interface{}) *DateControl {
    a.Set("width", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *DateControl) EditorSetting(value interface{}) *DateControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *DateControl) SubmitOnChange(value interface{}) *DateControl {
    a.Set("submitOnChange", value)
    return a
}
