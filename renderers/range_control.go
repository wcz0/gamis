package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type RangeControl struct {
	*BaseRenderer
}

func NewRangeControl() *RangeControl {
    a := &RangeControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-range")
    return a
}


func (a *RangeControl) Set(name string, value interface{}) *RangeControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *RangeControl) LabelRemark(value interface{}) *RangeControl {
    a.Set("labelRemark", value)
    return a
}

/**
 */
func (a *RangeControl) Validations(value interface{}) *RangeControl {
    a.Set("validations", value)
    return a
}

/**
 * 是否展示步长
 */
func (a *RangeControl) ShowSteps(value interface{}) *RangeControl {
    a.Set("showSteps", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *RangeControl) StaticLabelClassName(value interface{}) *RangeControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *RangeControl) UseMobileUI(value interface{}) *RangeControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *RangeControl) SubmitOnChange(value interface{}) *RangeControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *RangeControl) DescriptionClassName(value interface{}) *RangeControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 分隔符
 */
func (a *RangeControl) Delimiter(value interface{}) *RangeControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *RangeControl) VisibleOn(value interface{}) *RangeControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *RangeControl) StaticClassName(value interface{}) *RangeControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *RangeControl) EditorSetting(value interface{}) *RangeControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 刻度
 */
func (a *RangeControl) Marks(value interface{}) *RangeControl {
    a.Set("marks", value)
    return a
}

/**
 * 是否禁用
 */
func (a *RangeControl) Disabled(value interface{}) *RangeControl {
    a.Set("disabled", value)
    return a
}

/**
 * 是否只读
 */
func (a *RangeControl) ReadOnly(value interface{}) *RangeControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *RangeControl) AutoFill(value interface{}) *RangeControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 占位符
 */
func (a *RangeControl) Placeholder(value interface{}) *RangeControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *RangeControl) Hidden(value interface{}) *RangeControl {
    a.Set("hidden", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *RangeControl) HiddenOn(value interface{}) *RangeControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *RangeControl) Visible(value interface{}) *RangeControl {
    a.Set("visible", value)
    return a
}

/**
 */
func (a *RangeControl) StaticSchema(value interface{}) *RangeControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *RangeControl) LabelWidth(value interface{}) *RangeControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *RangeControl) StaticOn(value interface{}) *RangeControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *RangeControl) StaticInputClassName(value interface{}) *RangeControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *RangeControl) Inline(value interface{}) *RangeControl {
    a.Set("inline", value)
    return a
}

/**
 */
func (a *RangeControl) InitAutoFill(value interface{}) *RangeControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 配置 label className
 */
func (a *RangeControl) LabelClassName(value interface{}) *RangeControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *RangeControl) Remark(value interface{}) *RangeControl {
    a.Set("remark", value)
    return a
}

/**
 * 分割块数
 */
func (a *RangeControl) Parts(value interface{}) *RangeControl {
    a.Set("parts", value)
    return a
}

/**
 */
func (a *RangeControl) Desc(value interface{}) *RangeControl {
    a.Set("desc", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *RangeControl) Width(value interface{}) *RangeControl {
    a.Set("width", value)
    return a
}

/**
 * 最大值
 */
func (a *RangeControl) Max(value interface{}) *RangeControl {
    a.Set("max", value)
    return a
}

/**
 * 最小值
 */
func (a *RangeControl) Min(value interface{}) *RangeControl {
    a.Set("min", value)
    return a
}

/**
 * 标签方向
 */
func (a *RangeControl) TooltipPlacement(value interface{}) *RangeControl {
    a.Set("tooltipPlacement", value)
    return a
}

/**
 * 是否通过分隔符连接
 */
func (a *RangeControl) JoinValues(value interface{}) *RangeControl {
    a.Set("joinValues", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *RangeControl) Name(value interface{}) *RangeControl {
    a.Set("name", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *RangeControl) ExtraName(value interface{}) *RangeControl {
    a.Set("extraName", value)
    return a
}

/**
 * 只读条件
 */
func (a *RangeControl) ReadOnlyOn(value interface{}) *RangeControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *RangeControl) ValidateOnChange(value interface{}) *RangeControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 表单项类型
 */
func (a *RangeControl) Type(value interface{}) *RangeControl {
    a.Set("type", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *RangeControl) DisabledOn(value interface{}) *RangeControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *RangeControl) Horizontal(value interface{}) *RangeControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *RangeControl) ValidateApi(value interface{}) *RangeControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 步长
 */
func (a *RangeControl) Step(value interface{}) *RangeControl {
    a.Set("step", value)
    return a
}

/**
 * 滑块值
 */
func (a *RangeControl) Value(value interface{}) *RangeControl {
    a.Set("value", value)
    return a
}

/**
 */
func (a *RangeControl) TestIdBuilder(value interface{}) *RangeControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 描述标题
 */
func (a *RangeControl) LabelAlign(value interface{}) *RangeControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *RangeControl) ValidationErrors(value interface{}) *RangeControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *RangeControl) Hint(value interface{}) *RangeControl {
    a.Set("hint", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *RangeControl) Description(value interface{}) *RangeControl {
    a.Set("description", value)
    return a
}

/**
 * 单位
 */
func (a *RangeControl) Unit(value interface{}) *RangeControl {
    a.Set("unit", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *RangeControl) ClassName(value interface{}) *RangeControl {
    a.Set("className", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *RangeControl) StaticPlaceholder(value interface{}) *RangeControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 组件样式
 */
func (a *RangeControl) Style(value interface{}) *RangeControl {
    a.Set("style", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *RangeControl) Size(value interface{}) *RangeControl {
    a.Set("size", value)
    return a
}

/**
 * 描述标题
 */
func (a *RangeControl) Label(value interface{}) *RangeControl {
    a.Set("label", value)
    return a
}

/**
 * 是否为双滑块
 */
func (a *RangeControl) Multiple(value interface{}) *RangeControl {
    a.Set("multiple", value)
    return a
}

/**
 * 输入框是否可清除
 */
func (a *RangeControl) Clearable(value interface{}) *RangeControl {
    a.Set("clearable", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *RangeControl) Id(value interface{}) *RangeControl {
    a.Set("id", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *RangeControl) Mode(value interface{}) *RangeControl {
    a.Set("mode", value)
    return a
}

/**
 * 是否为必填
 */
func (a *RangeControl) Required(value interface{}) *RangeControl {
    a.Set("required", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *RangeControl) ClearValueOnHidden(value interface{}) *RangeControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 */
func (a *RangeControl) Row(value interface{}) *RangeControl {
    a.Set("row", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *RangeControl) OnEvent(value interface{}) *RangeControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *RangeControl) Static(value interface{}) *RangeControl {
    a.Set("static", value)
    return a
}

/**
 * 配置 input className
 */
func (a *RangeControl) InputClassName(value interface{}) *RangeControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 是否展示标签
 */
func (a *RangeControl) TooltipVisible(value interface{}) *RangeControl {
    a.Set("tooltipVisible", value)
    return a
}

/**
 * 是否展示输入框
 */
func (a *RangeControl) ShowInput(value interface{}) *RangeControl {
    a.Set("showInput", value)
    return a
}
