package renderers


/**
 * Matrix 选择控件。适合做权限勾选。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/matrix

 * @author wcz0
 * @version 6.2.2
 */
type MatrixControl struct {
	*BaseRenderer
}

func NewMatrixControl() *MatrixControl {
    a := &MatrixControl{
        BaseRenderer: NewBaseRenderer(),
    }

func (a *MatrixControl) Set(name string, value interface{}) *MatrixControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}

    a.Set("type", "matrix-checkboxes")
    return a
}

/**
 */
func (a *MatrixControl) Desc(value interface{}) *MatrixControl {
    a.Set("desc", value)
    return a
}

/**
 * 是否为必填
 */
func (a *MatrixControl) Required(value interface{}) *MatrixControl {
    a.Set("required", value)
    return a
}

/**
 * 可用来通过 API 拉取 options。
 */
func (a *MatrixControl) Source(value interface{}) *MatrixControl {
    a.Set("source", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *MatrixControl) Classname(value interface{}) *MatrixControl {
    a.Set("className", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *MatrixControl) Staticon(value interface{}) *MatrixControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *MatrixControl) Staticlabelclassname(value interface{}) *MatrixControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *MatrixControl) Clearvalueonhidden(value interface{}) *MatrixControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *MatrixControl) Width(value interface{}) *MatrixControl {
    a.Set("width", value)
    return a
}

/**
 * 是否只读
 */
func (a *MatrixControl) Readonly(value interface{}) *MatrixControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *MatrixControl) Validationerrors(value interface{}) *MatrixControl {
    a.Set("validationErrors", value)
    return a
}

/**
 */
func (a *MatrixControl) Columns(value interface{}) *MatrixControl {
    a.Set("columns", value)
    return a
}

/**
 * 是否显示
 */
func (a *MatrixControl) Visible(value interface{}) *MatrixControl {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *MatrixControl) Staticplaceholder(value interface{}) *MatrixControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *MatrixControl) Hint(value interface{}) *MatrixControl {
    a.Set("hint", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *MatrixControl) Mode(value interface{}) *MatrixControl {
    a.Set("mode", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *MatrixControl) Visibleon(value interface{}) *MatrixControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 组件样式
 */
func (a *MatrixControl) Style(value interface{}) *MatrixControl {
    a.Set("style", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *MatrixControl) Labelwidth(value interface{}) *MatrixControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 占位符
 */
func (a *MatrixControl) Placeholder(value interface{}) *MatrixControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *MatrixControl) Staticclassname(value interface{}) *MatrixControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *MatrixControl) Size(value interface{}) *MatrixControl {
    a.Set("size", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *MatrixControl) Horizontal(value interface{}) *MatrixControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *MatrixControl) Name(value interface{}) *MatrixControl {
    a.Set("name", value)
    return a
}

/**
 * 只读条件
 */
func (a *MatrixControl) Readonlyon(value interface{}) *MatrixControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *MatrixControl) Descriptionclassname(value interface{}) *MatrixControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 配置 input className
 */
func (a *MatrixControl) Inputclassname(value interface{}) *MatrixControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *MatrixControl) Autofill(value interface{}) *MatrixControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *MatrixControl) Staticinputclassname(value interface{}) *MatrixControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *MatrixControl) Editorsetting(value interface{}) *MatrixControl {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *MatrixControl) Initautofill(value interface{}) *MatrixControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 */
func (a *MatrixControl) Rows(value interface{}) *MatrixControl {
    a.Set("rows", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *MatrixControl) Value(value interface{}) *MatrixControl {
    a.Set("value", value)
    return a
}

/**
 * 表单项类型
 */
func (a *MatrixControl) Type(value interface{}) *MatrixControl {
    a.Set("type", value)
    return a
}

/**
 * 行标题说明
 */
func (a *MatrixControl) Rowlabel(value interface{}) *MatrixControl {
    a.Set("rowLabel", value)
    return a
}

/**
 * 描述标题
 */
func (a *MatrixControl) Labelalign(value interface{}) *MatrixControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 配置 label className
 */
func (a *MatrixControl) Labelclassname(value interface{}) *MatrixControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *MatrixControl) Remark(value interface{}) *MatrixControl {
    a.Set("remark", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *MatrixControl) Validateonchange(value interface{}) *MatrixControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *MatrixControl) Description(value interface{}) *MatrixControl {
    a.Set("description", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *MatrixControl) Hidden(value interface{}) *MatrixControl {
    a.Set("hidden", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *MatrixControl) Hiddenon(value interface{}) *MatrixControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 描述标题
 */
func (a *MatrixControl) Label(value interface{}) *MatrixControl {
    a.Set("label", value)
    return a
}

/**
 */
func (a *MatrixControl) Row(value interface{}) *MatrixControl {
    a.Set("row", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *MatrixControl) Id(value interface{}) *MatrixControl {
    a.Set("id", value)
    return a
}

/**
 */
func (a *MatrixControl) Testidbuilder(value interface{}) *MatrixControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *MatrixControl) Extraname(value interface{}) *MatrixControl {
    a.Set("extraName", value)
    return a
}

/**
 */
func (a *MatrixControl) Validations(value interface{}) *MatrixControl {
    a.Set("validations", value)
    return a
}

/**
 * 是否禁用
 */
func (a *MatrixControl) Disabled(value interface{}) *MatrixControl {
    a.Set("disabled", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *MatrixControl) Onevent(value interface{}) *MatrixControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *MatrixControl) Labelremark(value interface{}) *MatrixControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *MatrixControl) Disabledon(value interface{}) *MatrixControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *MatrixControl) Static(value interface{}) *MatrixControl {
    a.Set("static", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *MatrixControl) Submitonchange(value interface{}) *MatrixControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *MatrixControl) Inline(value interface{}) *MatrixControl {
    a.Set("inline", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *MatrixControl) Validateapi(value interface{}) *MatrixControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 设置单选模式，multiple为false时有效
 */
func (a *MatrixControl) Singleselectmode(value interface{}) *MatrixControl {
    a.Set("singleSelectMode", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *MatrixControl) Usemobileui(value interface{}) *MatrixControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 配置singleSelectMode时设置为false
 */
func (a *MatrixControl) Multiple(value interface{}) *MatrixControl {
    a.Set("multiple", value)
    return a
}

/**
 */
func (a *MatrixControl) Staticschema(value interface{}) *MatrixControl {
    a.Set("staticSchema", value)
    return a
}
