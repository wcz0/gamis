package renderers


/**
 * QuarterRange 季度范围控件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/input-quarter-range

 * @author wcz0
 * @version 6.2.2
 */
type QuarterRangeControl struct {
	*BaseRenderer
}

func NewQuarterRangeControl() *QuarterRangeControl {
    a := &QuarterRangeControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-quarter-range")
    return a
}


func (a *QuarterRangeControl) Set(name string, value interface{}) *QuarterRangeControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 占位符
 */
func (a *QuarterRangeControl) Placeholder(value interface{}) *QuarterRangeControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否为必填
 */
func (a *QuarterRangeControl) Required(value interface{}) *QuarterRangeControl {
    a.Set("required", value)
    return a
}

/**
 */
func (a *QuarterRangeControl) Row(value interface{}) *QuarterRangeControl {
    a.Set("row", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *QuarterRangeControl) StaticPlaceholder(value interface{}) *QuarterRangeControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *QuarterRangeControl) StaticClassName(value interface{}) *QuarterRangeControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 最大日期限制，支持变量 $xxx 来取值，或者用相对值如：* `-2mins` 2分钟前\n * `+2days` 2天后\n* `-10week` 十周前\n可用单位： `min`、`hour`、`day`、`week`、`month`、`year`。所有单位支持复数形式。
 */
func (a *QuarterRangeControl) MaxDate(value interface{}) *QuarterRangeControl {
    a.Set("maxDate", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *QuarterRangeControl) SubmitOnChange(value interface{}) *QuarterRangeControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *QuarterRangeControl) DescriptionClassName(value interface{}) *QuarterRangeControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *QuarterRangeControl) Style(value interface{}) *QuarterRangeControl {
    a.Set("style", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *QuarterRangeControl) EditorSetting(value interface{}) *QuarterRangeControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *QuarterRangeControl) Hint(value interface{}) *QuarterRangeControl {
    a.Set("hint", value)
    return a
}

/**
 * 描述标题
 */
func (a *QuarterRangeControl) LabelAlign(value interface{}) *QuarterRangeControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *QuarterRangeControl) ValidateApi(value interface{}) *QuarterRangeControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *QuarterRangeControl) HiddenOn(value interface{}) *QuarterRangeControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *QuarterRangeControl) UseMobileUI(value interface{}) *QuarterRangeControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *QuarterRangeControl) LabelRemark(value interface{}) *QuarterRangeControl {
    a.Set("labelRemark", value)
    return a
}

/**
 */
func (a *QuarterRangeControl) InitAutoFill(value interface{}) *QuarterRangeControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *QuarterRangeControl) Hidden(value interface{}) *QuarterRangeControl {
    a.Set("hidden", value)
    return a
}

/**
 */
func (a *QuarterRangeControl) Type(value interface{}) *QuarterRangeControl {
    a.Set("type", value)
    return a
}

/**
 * 默认 `YYYY-MM-DD` 用来配置显示的时间格式。
 */
func (a *QuarterRangeControl) InputFormat(value interface{}) *QuarterRangeControl {
    a.Set("inputFormat", value)
    return a
}

/**
 * 是否显示
 */
func (a *QuarterRangeControl) Visible(value interface{}) *QuarterRangeControl {
    a.Set("visible", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *QuarterRangeControl) VisibleOn(value interface{}) *QuarterRangeControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *QuarterRangeControl) ClearValueOnHidden(value interface{}) *QuarterRangeControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 用来配置显示的时间格式（新：同inputFormat）
 */
func (a *QuarterRangeControl) DisplayFormat(value interface{}) *QuarterRangeControl {
    a.Set("displayFormat", value)
    return a
}

/**
 * 开启后将选中的选项 value 的值用连接符拼接起来，作为当前表单项的值。如： `value1,value2` 否则为 `[value1, value2]`
 */
func (a *QuarterRangeControl) JoinValues(value interface{}) *QuarterRangeControl {
    a.Set("joinValues", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *QuarterRangeControl) ValidationErrors(value interface{}) *QuarterRangeControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 配置 input className
 */
func (a *QuarterRangeControl) InputClassName(value interface{}) *QuarterRangeControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *QuarterRangeControl) StaticLabelClassName(value interface{}) *QuarterRangeControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *QuarterRangeControl) Width(value interface{}) *QuarterRangeControl {
    a.Set("width", value)
    return a
}

/**
 * 只读条件
 */
func (a *QuarterRangeControl) ReadOnlyOn(value interface{}) *QuarterRangeControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 */
func (a *QuarterRangeControl) Desc(value interface{}) *QuarterRangeControl {
    a.Set("desc", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *QuarterRangeControl) Mode(value interface{}) *QuarterRangeControl {
    a.Set("mode", value)
    return a
}

/**
 * 日期数据处理函数，用来处理选择日期之后的的值(value: moment.Moment, config: {type: 'start' | 'end'; originValue: moment.Moment, timeFormat: string}, props: any, data: any, moment: moment) => moment.Moment;
 */
func (a *QuarterRangeControl) Transform(value interface{}) *QuarterRangeControl {
    a.Set("transform", value)
    return a
}

/**
 * 描述标题
 */
func (a *QuarterRangeControl) Label(value interface{}) *QuarterRangeControl {
    a.Set("label", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *QuarterRangeControl) Name(value interface{}) *QuarterRangeControl {
    a.Set("name", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *QuarterRangeControl) OnEvent(value interface{}) *QuarterRangeControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *QuarterRangeControl) StaticOn(value interface{}) *QuarterRangeControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 分割符, 因为有两个值，开始时间和结束时间，所以要有连接符。默认为英文逗号。
 */
func (a *QuarterRangeControl) Delimiter(value interface{}) *QuarterRangeControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 日期范围快捷键
 */
func (a *QuarterRangeControl) Shortcuts(value interface{}) *QuarterRangeControl {
    a.Set("shortcuts", value)
    return a
}

/**
 * 日期范围结束时间-占位符
 */
func (a *QuarterRangeControl) EndPlaceholder(value interface{}) *QuarterRangeControl {
    a.Set("endPlaceholder", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *QuarterRangeControl) Inline(value interface{}) *QuarterRangeControl {
    a.Set("inline", value)
    return a
}

/**
 */
func (a *QuarterRangeControl) Validations(value interface{}) *QuarterRangeControl {
    a.Set("validations", value)
    return a
}

/**
 * 日期范围快捷键
 */
func (a *QuarterRangeControl) Ranges(value interface{}) *QuarterRangeControl {
    a.Set("ranges", value)
    return a
}

/**
 * 弹窗容器选择器
 */
func (a *QuarterRangeControl) PopOverContainerSelector(value interface{}) *QuarterRangeControl {
    a.Set("popOverContainerSelector", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *QuarterRangeControl) Size(value interface{}) *QuarterRangeControl {
    a.Set("size", value)
    return a
}

/**
 * 日期范围开始时间-占位符
 */
func (a *QuarterRangeControl) StartPlaceholder(value interface{}) *QuarterRangeControl {
    a.Set("startPlaceholder", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *QuarterRangeControl) AutoFill(value interface{}) *QuarterRangeControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *QuarterRangeControl) DisabledOn(value interface{}) *QuarterRangeControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 用来提交的时间格式。更多格式类型请参考 moment.（新：同format）
 */
func (a *QuarterRangeControl) ValueFormat(value interface{}) *QuarterRangeControl {
    a.Set("valueFormat", value)
    return a
}

/**
 * 最小跨度，比如 2days
 */
func (a *QuarterRangeControl) MinDuration(value interface{}) *QuarterRangeControl {
    a.Set("minDuration", value)
    return a
}

/**
 * 这里面 value 需要特殊说明一下，因为支持相对值。* `-2mins` 2分钟前\n * `+2days` 2天后\n* `-10week` 十周前\n可用单位： `min`、`hour`、`day`、`week`、`month`、`year`。所有单位支持复数形式。
 */
func (a *QuarterRangeControl) Value(value interface{}) *QuarterRangeControl {
    a.Set("value", value)
    return a
}

/**
 * 是否禁用
 */
func (a *QuarterRangeControl) Disabled(value interface{}) *QuarterRangeControl {
    a.Set("disabled", value)
    return a
}

/**
 * 开启后变成非弹出模式，即内联模式。
 */
func (a *QuarterRangeControl) Embed(value interface{}) *QuarterRangeControl {
    a.Set("embed", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *QuarterRangeControl) Description(value interface{}) *QuarterRangeControl {
    a.Set("description", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *QuarterRangeControl) ClassName(value interface{}) *QuarterRangeControl {
    a.Set("className", value)
    return a
}

/**
 * 默认 `X` 即时间戳格式，用来提交的时间格式。更多格式类型请参考 moment.
 */
func (a *QuarterRangeControl) Format(value interface{}) *QuarterRangeControl {
    a.Set("format", value)
    return a
}

/**
 * 最小日期限制，支持变量 $xxx 来取值，或者用相对值如：* `-2mins` 2分钟前\n * `+2days` 2天后\n* `-10week` 十周前\n可用单位： `min`、`hour`、`day`、`week`、`month`、`year`。所有单位支持复数形式。
 */
func (a *QuarterRangeControl) MinDate(value interface{}) *QuarterRangeControl {
    a.Set("minDate", value)
    return a
}

/**
 * 是否启用游标动画，默认开启
 */
func (a *QuarterRangeControl) Animation(value interface{}) *QuarterRangeControl {
    a.Set("animation", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *QuarterRangeControl) ValidateOnChange(value interface{}) *QuarterRangeControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *QuarterRangeControl) StaticInputClassName(value interface{}) *QuarterRangeControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 最大跨度，比如 2days
 */
func (a *QuarterRangeControl) MaxDuration(value interface{}) *QuarterRangeControl {
    a.Set("maxDuration", value)
    return a
}

/**
 * 配置 label className
 */
func (a *QuarterRangeControl) LabelClassName(value interface{}) *QuarterRangeControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 是否只读
 */
func (a *QuarterRangeControl) ReadOnly(value interface{}) *QuarterRangeControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *QuarterRangeControl) Remark(value interface{}) *QuarterRangeControl {
    a.Set("remark", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *QuarterRangeControl) Horizontal(value interface{}) *QuarterRangeControl {
    a.Set("horizontal", value)
    return a
}

/**
 */
func (a *QuarterRangeControl) StaticSchema(value interface{}) *QuarterRangeControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *QuarterRangeControl) LabelWidth(value interface{}) *QuarterRangeControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *QuarterRangeControl) Id(value interface{}) *QuarterRangeControl {
    a.Set("id", value)
    return a
}

/**
 */
func (a *QuarterRangeControl) TestIdBuilder(value interface{}) *QuarterRangeControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 边框模式，全边框，还是半边框，或者没边框。
 * 可选值: full | half | none
 */
func (a *QuarterRangeControl) BorderMode(value interface{}) *QuarterRangeControl {
    a.Set("borderMode", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *QuarterRangeControl) ExtraName(value interface{}) *QuarterRangeControl {
    a.Set("extraName", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *QuarterRangeControl) Static(value interface{}) *QuarterRangeControl {
    a.Set("static", value)
    return a
}
