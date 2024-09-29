package renderers


/**
 * MonthRange 月范围控件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/month-range

 * @author wcz0
 * @version 6.2.2
 */
type MonthRangeControl struct {
	*BaseRenderer
}

func NewMonthRangeControl() *MonthRangeControl {
    a := &MonthRangeControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-month-range")
    return a
}


func (a *MonthRangeControl) Set(name string, value interface{}) *MonthRangeControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 配置描述上的 className
 */
func (a *MonthRangeControl) DescriptionClassName(value interface{}) *MonthRangeControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 默认 `X` 即时间戳格式，用来提交的时间格式。更多格式类型请参考 moment.
 */
func (a *MonthRangeControl) Format(value interface{}) *MonthRangeControl {
    a.Set("format", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *MonthRangeControl) ExtraName(value interface{}) *MonthRangeControl {
    a.Set("extraName", value)
    return a
}

/**
 */
func (a *MonthRangeControl) Validations(value interface{}) *MonthRangeControl {
    a.Set("validations", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *MonthRangeControl) VisibleOn(value interface{}) *MonthRangeControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *MonthRangeControl) EditorSetting(value interface{}) *MonthRangeControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 日期范围快捷键
 */
func (a *MonthRangeControl) Ranges(value interface{}) *MonthRangeControl {
    a.Set("ranges", value)
    return a
}

/**
 * 日期数据处理函数，用来处理选择日期之后的的值(value: moment.Moment, config: {type: 'start' | 'end'; originValue: moment.Moment, timeFormat: string}, props: any, data: any, moment: moment) => moment.Moment;
 */
func (a *MonthRangeControl) Transform(value interface{}) *MonthRangeControl {
    a.Set("transform", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *MonthRangeControl) Size(value interface{}) *MonthRangeControl {
    a.Set("size", value)
    return a
}

/**
 * 是否禁用
 */
func (a *MonthRangeControl) Disabled(value interface{}) *MonthRangeControl {
    a.Set("disabled", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *MonthRangeControl) StaticOn(value interface{}) *MonthRangeControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *MonthRangeControl) StaticClassName(value interface{}) *MonthRangeControl {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *MonthRangeControl) TestIdBuilder(value interface{}) *MonthRangeControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 开启后将选中的选项 value 的值用连接符拼接起来，作为当前表单项的值。如： `value1,value2` 否则为 `[value1, value2]`
 */
func (a *MonthRangeControl) JoinValues(value interface{}) *MonthRangeControl {
    a.Set("joinValues", value)
    return a
}

/**
 * 最小跨度，比如 2days
 */
func (a *MonthRangeControl) MinDuration(value interface{}) *MonthRangeControl {
    a.Set("minDuration", value)
    return a
}

/**
 * 日期范围开始时间-占位符
 */
func (a *MonthRangeControl) StartPlaceholder(value interface{}) *MonthRangeControl {
    a.Set("startPlaceholder", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *MonthRangeControl) Width(value interface{}) *MonthRangeControl {
    a.Set("width", value)
    return a
}

/**
 * 最大日期限制，支持变量 $xxx 来取值，或者用相对值如：* `-2mins` 2分钟前\n * `+2days` 2天后\n* `-10week` 十周前\n可用单位： `min`、`hour`、`day`、`week`、`month`、`year`。所有单位支持复数形式。
 */
func (a *MonthRangeControl) MaxDate(value interface{}) *MonthRangeControl {
    a.Set("maxDate", value)
    return a
}

/**
 * 只读条件
 */
func (a *MonthRangeControl) ReadOnlyOn(value interface{}) *MonthRangeControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *MonthRangeControl) Mode(value interface{}) *MonthRangeControl {
    a.Set("mode", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *MonthRangeControl) ClassName(value interface{}) *MonthRangeControl {
    a.Set("className", value)
    return a
}

/**
 * 最小日期限制，支持变量 $xxx 来取值，或者用相对值如：* `-2mins` 2分钟前\n * `+2days` 2天后\n* `-10week` 十周前\n可用单位： `min`、`hour`、`day`、`week`、`month`、`year`。所有单位支持复数形式。
 */
func (a *MonthRangeControl) MinDate(value interface{}) *MonthRangeControl {
    a.Set("minDate", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *MonthRangeControl) LabelRemark(value interface{}) *MonthRangeControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 描述标题
 */
func (a *MonthRangeControl) LabelAlign(value interface{}) *MonthRangeControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *MonthRangeControl) Name(value interface{}) *MonthRangeControl {
    a.Set("name", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *MonthRangeControl) Description(value interface{}) *MonthRangeControl {
    a.Set("description", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *MonthRangeControl) ClearValueOnHidden(value interface{}) *MonthRangeControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 */
func (a *MonthRangeControl) Row(value interface{}) *MonthRangeControl {
    a.Set("row", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *MonthRangeControl) DisabledOn(value interface{}) *MonthRangeControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *MonthRangeControl) Visible(value interface{}) *MonthRangeControl {
    a.Set("visible", value)
    return a
}

/**
 * 日期范围快捷键
 */
func (a *MonthRangeControl) Shortcuts(value interface{}) *MonthRangeControl {
    a.Set("shortcuts", value)
    return a
}

/**
 * 是否启用游标动画，默认开启
 */
func (a *MonthRangeControl) Animation(value interface{}) *MonthRangeControl {
    a.Set("animation", value)
    return a
}

/**
 * 弹窗容器选择器
 */
func (a *MonthRangeControl) PopOverContainerSelector(value interface{}) *MonthRangeControl {
    a.Set("popOverContainerSelector", value)
    return a
}

/**
 */
func (a *MonthRangeControl) Desc(value interface{}) *MonthRangeControl {
    a.Set("desc", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *MonthRangeControl) ValidationErrors(value interface{}) *MonthRangeControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *MonthRangeControl) HiddenOn(value interface{}) *MonthRangeControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 最大跨度，比如 2days
 */
func (a *MonthRangeControl) MaxDuration(value interface{}) *MonthRangeControl {
    a.Set("maxDuration", value)
    return a
}

/**
 * 描述标题
 */
func (a *MonthRangeControl) Label(value interface{}) *MonthRangeControl {
    a.Set("label", value)
    return a
}

/**
 */
func (a *MonthRangeControl) InitAutoFill(value interface{}) *MonthRangeControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *MonthRangeControl) OnEvent(value interface{}) *MonthRangeControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 组件样式
 */
func (a *MonthRangeControl) Style(value interface{}) *MonthRangeControl {
    a.Set("style", value)
    return a
}

/**
 * 分割符, 因为有两个值，开始时间和结束时间，所以要有连接符。默认为英文逗号。
 */
func (a *MonthRangeControl) Delimiter(value interface{}) *MonthRangeControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *MonthRangeControl) Remark(value interface{}) *MonthRangeControl {
    a.Set("remark", value)
    return a
}

/**
 * 配置 label className
 */
func (a *MonthRangeControl) LabelClassName(value interface{}) *MonthRangeControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *MonthRangeControl) StaticInputClassName(value interface{}) *MonthRangeControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *MonthRangeControl) UseMobileUI(value interface{}) *MonthRangeControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *MonthRangeControl) Type(value interface{}) *MonthRangeControl {
    a.Set("type", value)
    return a
}

/**
 * 用来提交的时间格式。更多格式类型请参考 moment.（新：同format）
 */
func (a *MonthRangeControl) ValueFormat(value interface{}) *MonthRangeControl {
    a.Set("valueFormat", value)
    return a
}

/**
 * 默认 `YYYY-MM-DD` 用来配置显示的时间格式。
 */
func (a *MonthRangeControl) InputFormat(value interface{}) *MonthRangeControl {
    a.Set("inputFormat", value)
    return a
}

/**
 * 用来配置显示的时间格式（新：同inputFormat）
 */
func (a *MonthRangeControl) DisplayFormat(value interface{}) *MonthRangeControl {
    a.Set("displayFormat", value)
    return a
}

/**
 * 配置 input className
 */
func (a *MonthRangeControl) InputClassName(value interface{}) *MonthRangeControl {
    a.Set("inputClassName", value)
    return a
}

/**
 */
func (a *MonthRangeControl) StaticSchema(value interface{}) *MonthRangeControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 边框模式，全边框，还是半边框，或者没边框。
 * 可选值: full | half | none
 */
func (a *MonthRangeControl) BorderMode(value interface{}) *MonthRangeControl {
    a.Set("borderMode", value)
    return a
}

/**
 * 开启后变成非弹出模式，即内联模式。
 */
func (a *MonthRangeControl) Embed(value interface{}) *MonthRangeControl {
    a.Set("embed", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *MonthRangeControl) LabelWidth(value interface{}) *MonthRangeControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 是否只读
 */
func (a *MonthRangeControl) ReadOnly(value interface{}) *MonthRangeControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 占位符
 */
func (a *MonthRangeControl) Placeholder(value interface{}) *MonthRangeControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *MonthRangeControl) Hint(value interface{}) *MonthRangeControl {
    a.Set("hint", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *MonthRangeControl) Horizontal(value interface{}) *MonthRangeControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *MonthRangeControl) ValidateApi(value interface{}) *MonthRangeControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 日期范围结束时间-占位符
 */
func (a *MonthRangeControl) EndPlaceholder(value interface{}) *MonthRangeControl {
    a.Set("endPlaceholder", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *MonthRangeControl) Hidden(value interface{}) *MonthRangeControl {
    a.Set("hidden", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *MonthRangeControl) Static(value interface{}) *MonthRangeControl {
    a.Set("static", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *MonthRangeControl) StaticPlaceholder(value interface{}) *MonthRangeControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *MonthRangeControl) StaticLabelClassName(value interface{}) *MonthRangeControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 这里面 value 需要特殊说明一下，因为支持相对值。* `-2mins` 2分钟前\n * `+2days` 2天后\n* `-10week` 十周前\n可用单位： `min`、`hour`、`day`、`week`、`month`、`year`。所有单位支持复数形式。
 */
func (a *MonthRangeControl) Value(value interface{}) *MonthRangeControl {
    a.Set("value", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *MonthRangeControl) SubmitOnChange(value interface{}) *MonthRangeControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *MonthRangeControl) ValidateOnChange(value interface{}) *MonthRangeControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *MonthRangeControl) Inline(value interface{}) *MonthRangeControl {
    a.Set("inline", value)
    return a
}

/**
 * 是否为必填
 */
func (a *MonthRangeControl) Required(value interface{}) *MonthRangeControl {
    a.Set("required", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *MonthRangeControl) AutoFill(value interface{}) *MonthRangeControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *MonthRangeControl) Id(value interface{}) *MonthRangeControl {
    a.Set("id", value)
    return a
}
