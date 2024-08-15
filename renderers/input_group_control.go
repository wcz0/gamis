package renderers


/**
 * InputGroup 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/input-group

 * @author wcz0
 * @version 6.2.2
 */
type InputGroupControl struct {
	*BaseRenderer
}

func NewInputGroupControl() *InputGroupControl {
    a := &InputGroupControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-group")
    return a
}


func (a *InputGroupControl) Set(name string, value interface{}) *InputGroupControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 是否隐藏
 */
func (a *InputGroupControl) Hidden(value interface{}) *InputGroupControl {
    a.Set("hidden", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *InputGroupControl) Staticclassname(value interface{}) *InputGroupControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *InputGroupControl) Validateapi(value interface{}) *InputGroupControl {
    a.Set("validateApi", value)
    return a
}

/**
 */
func (a *InputGroupControl) Staticschema(value interface{}) *InputGroupControl {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *InputGroupControl) Testidbuilder(value interface{}) *InputGroupControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *InputGroupControl) Labelremark(value interface{}) *InputGroupControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *InputGroupControl) Autofill(value interface{}) *InputGroupControl {
    a.Set("autoFill", value)
    return a
}

/**
 * FormItem 集合
 */
func (a *InputGroupControl) Body(value interface{}) *InputGroupControl {
    a.Set("body", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *InputGroupControl) Validationerrors(value interface{}) *InputGroupControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 是否禁用
 */
func (a *InputGroupControl) Disabled(value interface{}) *InputGroupControl {
    a.Set("disabled", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *InputGroupControl) Visibleon(value interface{}) *InputGroupControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 配置 label className
 */
func (a *InputGroupControl) Labelclassname(value interface{}) *InputGroupControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 只读条件
 */
func (a *InputGroupControl) Readonlyon(value interface{}) *InputGroupControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 */
func (a *InputGroupControl) Desc(value interface{}) *InputGroupControl {
    a.Set("desc", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *InputGroupControl) Inline(value interface{}) *InputGroupControl {
    a.Set("inline", value)
    return a
}

/**
 * 配置 input className
 */
func (a *InputGroupControl) Inputclassname(value interface{}) *InputGroupControl {
    a.Set("inputClassName", value)
    return a
}

/**
 */
func (a *InputGroupControl) Initautofill(value interface{}) *InputGroupControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 表单项类型
 */
func (a *InputGroupControl) Type(value interface{}) *InputGroupControl {
    a.Set("type", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *InputGroupControl) Horizontal(value interface{}) *InputGroupControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *InputGroupControl) Clearvalueonhidden(value interface{}) *InputGroupControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *InputGroupControl) Onevent(value interface{}) *InputGroupControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *InputGroupControl) Classname(value interface{}) *InputGroupControl {
    a.Set("className", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *InputGroupControl) Hint(value interface{}) *InputGroupControl {
    a.Set("hint", value)
    return a
}

/**
 */
func (a *InputGroupControl) Validations(value interface{}) *InputGroupControl {
    a.Set("validations", value)
    return a
}

/**
 * 是否显示
 */
func (a *InputGroupControl) Visible(value interface{}) *InputGroupControl {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *InputGroupControl) Staticlabelclassname(value interface{}) *InputGroupControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 描述标题
 */
func (a *InputGroupControl) Label(value interface{}) *InputGroupControl {
    a.Set("label", value)
    return a
}

/**
 * 是否为必填
 */
func (a *InputGroupControl) Required(value interface{}) *InputGroupControl {
    a.Set("required", value)
    return a
}

/**
 * 描述标题
 */
func (a *InputGroupControl) Labelalign(value interface{}) *InputGroupControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *InputGroupControl) Remark(value interface{}) *InputGroupControl {
    a.Set("remark", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *InputGroupControl) Staticinputclassname(value interface{}) *InputGroupControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *InputGroupControl) Editorsetting(value interface{}) *InputGroupControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *InputGroupControl) Labelwidth(value interface{}) *InputGroupControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *InputGroupControl) Name(value interface{}) *InputGroupControl {
    a.Set("name", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *InputGroupControl) Descriptionclassname(value interface{}) *InputGroupControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 占位符
 */
func (a *InputGroupControl) Placeholder(value interface{}) *InputGroupControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *InputGroupControl) Width(value interface{}) *InputGroupControl {
    a.Set("width", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *InputGroupControl) Usemobileui(value interface{}) *InputGroupControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 是否只读
 */
func (a *InputGroupControl) Readonly(value interface{}) *InputGroupControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *InputGroupControl) Description(value interface{}) *InputGroupControl {
    a.Set("description", value)
    return a
}

/**
 * 校验提示信息配置
 */
func (a *InputGroupControl) Validationconfig(value interface{}) *InputGroupControl {
    a.Set("validationConfig", value)
    return a
}

/**
 * 组件样式
 */
func (a *InputGroupControl) Style(value interface{}) *InputGroupControl {
    a.Set("style", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *InputGroupControl) Mode(value interface{}) *InputGroupControl {
    a.Set("mode", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *InputGroupControl) Value(value interface{}) *InputGroupControl {
    a.Set("value", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *InputGroupControl) Static(value interface{}) *InputGroupControl {
    a.Set("static", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *InputGroupControl) Hiddenon(value interface{}) *InputGroupControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *InputGroupControl) Id(value interface{}) *InputGroupControl {
    a.Set("id", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *InputGroupControl) Staticon(value interface{}) *InputGroupControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *InputGroupControl) Submitonchange(value interface{}) *InputGroupControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *InputGroupControl) Validateonchange(value interface{}) *InputGroupControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *InputGroupControl) Disabledon(value interface{}) *InputGroupControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *InputGroupControl) Staticplaceholder(value interface{}) *InputGroupControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *InputGroupControl) Size(value interface{}) *InputGroupControl {
    a.Set("size", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *InputGroupControl) Extraname(value interface{}) *InputGroupControl {
    a.Set("extraName", value)
    return a
}

/**
 */
func (a *InputGroupControl) Row(value interface{}) *InputGroupControl {
    a.Set("row", value)
    return a
}
