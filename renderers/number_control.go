package renderers


/**
 * 数字输入框 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/input-number

 * @author wcz0
 * @version 6.2.2
 */
type NumberControl struct {
	*BaseRenderer
}

func NewNumberControl() *NumberControl {
    a := &NumberControl{
        BaseRenderer: NewBaseRenderer(),
    }

func (a *NumberControl) Set(name string, value interface{}) *NumberControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}

    a.Set("type", "input-number")
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *NumberControl) Size(value interface{}) *NumberControl {
    a.Set("size", value)
    return a
}

/**
 * 只读
 */
func (a *NumberControl) Readonly(value interface{}) *NumberControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *NumberControl) Validationerrors(value interface{}) *NumberControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *NumberControl) Usemobileui(value interface{}) *NumberControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *NumberControl) Testidbuilder(value interface{}) *NumberControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *NumberControl) Labelwidth(value interface{}) *NumberControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 是否为必填
 */
func (a *NumberControl) Required(value interface{}) *NumberControl {
    a.Set("required", value)
    return a
}

/**
 */
func (a *NumberControl) Row(value interface{}) *NumberControl {
    a.Set("row", value)
    return a
}

/**
 * 是否显示上下点击按钮
 */
func (a *NumberControl) Showsteps(value interface{}) *NumberControl {
    a.Set("showSteps", value)
    return a
}

/**
 * 最小值
 */
func (a *NumberControl) Min(value interface{}) *NumberControl {
    a.Set("min", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *NumberControl) Classname(value interface{}) *NumberControl {
    a.Set("className", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *NumberControl) Static(value interface{}) *NumberControl {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *NumberControl) Staticclassname(value interface{}) *NumberControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *NumberControl) Style(value interface{}) *NumberControl {
    a.Set("style", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *NumberControl) Remark(value interface{}) *NumberControl {
    a.Set("remark", value)
    return a
}

/**
 * 只读条件
 */
func (a *NumberControl) Readonlyon(value interface{}) *NumberControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *NumberControl) Descriptionclassname(value interface{}) *NumberControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *NumberControl) Width(value interface{}) *NumberControl {
    a.Set("width", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *NumberControl) Staticlabelclassname(value interface{}) *NumberControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *NumberControl) Editorsetting(value interface{}) *NumberControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *NumberControl) Extraname(value interface{}) *NumberControl {
    a.Set("extraName", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *NumberControl) Validateapi(value interface{}) *NumberControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 步长
 */
func (a *NumberControl) Step(value interface{}) *NumberControl {
    a.Set("step", value)
    return a
}

/**
 * 单位列表
 */
func (a *NumberControl) Unitoptions(value interface{}) *NumberControl {
    a.Set("unitOptions", value)
    return a
}

/**
 * 是否千分分隔
 */
func (a *NumberControl) Kilobitseparator(value interface{}) *NumberControl {
    a.Set("kilobitSeparator", value)
    return a
}

/**
 * 是否禁用
 */
func (a *NumberControl) Disabled(value interface{}) *NumberControl {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *NumberControl) Hidden(value interface{}) *NumberControl {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示
 */
func (a *NumberControl) Visible(value interface{}) *NumberControl {
    a.Set("visible", value)
    return a
}

/**
 */
func (a *NumberControl) Desc(value interface{}) *NumberControl {
    a.Set("desc", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *NumberControl) Horizontal(value interface{}) *NumberControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 精度
 */
func (a *NumberControl) Precision(value interface{}) *NumberControl {
    a.Set("precision", value)
    return a
}

/**
 * 是否是大数，如果是的话输入输出都将是字符串
 */
func (a *NumberControl) Big(value interface{}) *NumberControl {
    a.Set("big", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *NumberControl) Name(value interface{}) *NumberControl {
    a.Set("name", value)
    return a
}

/**
 * 边框模式，全边框，还是半边框，或者没边框。
 * 可选值: full | half | none
 */
func (a *NumberControl) Bordermode(value interface{}) *NumberControl {
    a.Set("borderMode", value)
    return a
}

/**
 * 描述标题
 */
func (a *NumberControl) Label(value interface{}) *NumberControl {
    a.Set("label", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *NumberControl) Hint(value interface{}) *NumberControl {
    a.Set("hint", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *NumberControl) Clearvalueonhidden(value interface{}) *NumberControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 最大值
 */
func (a *NumberControl) Max(value interface{}) *NumberControl {
    a.Set("max", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *NumberControl) Hiddenon(value interface{}) *NumberControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 表单项类型
 */
func (a *NumberControl) Type(value interface{}) *NumberControl {
    a.Set("type", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *NumberControl) Visibleon(value interface{}) *NumberControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 输入框为基础输入框还是加强输入框
 * 可选值: base | enhance
 */
func (a *NumberControl) Displaymode(value interface{}) *NumberControl {
    a.Set("displayMode", value)
    return a
}

/**
 * 描述标题
 */
func (a *NumberControl) Labelalign(value interface{}) *NumberControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *NumberControl) Submitonchange(value interface{}) *NumberControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 */
func (a *NumberControl) Validations(value interface{}) *NumberControl {
    a.Set("validations", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *NumberControl) Autofill(value interface{}) *NumberControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *NumberControl) Disabledon(value interface{}) *NumberControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *NumberControl) Staticinputclassname(value interface{}) *NumberControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 配置 label className
 */
func (a *NumberControl) Labelclassname(value interface{}) *NumberControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *NumberControl) Inline(value interface{}) *NumberControl {
    a.Set("inline", value)
    return a
}

/**
 * 后缀
 */
func (a *NumberControl) Suffix(value interface{}) *NumberControl {
    a.Set("suffix", value)
    return a
}

/**
 * 用来开启百分号的展示形式
 */
func (a *NumberControl) Showaspercent(value interface{}) *NumberControl {
    a.Set("showAsPercent", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *NumberControl) Validateonchange(value interface{}) *NumberControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *NumberControl) Labelremark(value interface{}) *NumberControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *NumberControl) Description(value interface{}) *NumberControl {
    a.Set("description", value)
    return a
}

/**
 */
func (a *NumberControl) Initautofill(value interface{}) *NumberControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *NumberControl) Onevent(value interface{}) *NumberControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *NumberControl) Staticplaceholder(value interface{}) *NumberControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *NumberControl) Staticschema(value interface{}) *NumberControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *NumberControl) Mode(value interface{}) *NumberControl {
    a.Set("mode", value)
    return a
}

/**
 * 配置 input className
 */
func (a *NumberControl) Inputclassname(value interface{}) *NumberControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 前缀
 */
func (a *NumberControl) Prefix(value interface{}) *NumberControl {
    a.Set("prefix", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *NumberControl) Id(value interface{}) *NumberControl {
    a.Set("id", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *NumberControl) Staticon(value interface{}) *NumberControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 占位符
 */
func (a *NumberControl) Placeholder(value interface{}) *NumberControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *NumberControl) Value(value interface{}) *NumberControl {
    a.Set("value", value)
    return a
}

/**
 * 是否启用键盘行为
 */
func (a *NumberControl) Keyboard(value interface{}) *NumberControl {
    a.Set("keyboard", value)
    return a
}
