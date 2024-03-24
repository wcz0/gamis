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

    a.Set("type", "matrix-checkboxes")
    return a
}
/**
 * 描述标题
 */
func (a *MatrixControl) LabelAlign(value interface{}) *MatrixControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 只读条件
 */
func (a *MatrixControl) ReadOnlyOn(value interface{}) *MatrixControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *MatrixControl) ValidateApi(value interface{}) *MatrixControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 行标题说明
 */
func (a *MatrixControl) RowLabel(value interface{}) *MatrixControl {
    a.Set("rowLabel", value)
    return a
}

/**
 */
func (a *MatrixControl) StaticSchema(value interface{}) *MatrixControl {
    a.Set("staticSchema", value)
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
 * 是否禁用表达式
 */
func (a *MatrixControl) DisabledOn(value interface{}) *MatrixControl {
    a.Set("disabledOn", value)
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
 * label自定义宽度，默认单位为px
 */
func (a *MatrixControl) LabelWidth(value interface{}) *MatrixControl {
    a.Set("labelWidth", value)
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
 * 是否为必填
 */
func (a *MatrixControl) Required(value interface{}) *MatrixControl {
    a.Set("required", value)
    return a
}

/**
 */
func (a *MatrixControl) Desc(value interface{}) *MatrixControl {
    a.Set("desc", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *MatrixControl) ValidationErrors(value interface{}) *MatrixControl {
    a.Set("validationErrors", value)
    return a
}

/**
 */
func (a *MatrixControl) Rows(value interface{}) *MatrixControl {
    a.Set("rows", value)
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
 * 是否隐藏
 */
func (a *MatrixControl) Hidden(value interface{}) *MatrixControl {
    a.Set("hidden", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *MatrixControl) SubmitOnChange(value interface{}) *MatrixControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 */
func (a *MatrixControl) Validations(value interface{}) *MatrixControl {
    a.Set("validations", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *MatrixControl) HiddenOn(value interface{}) *MatrixControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *MatrixControl) OnEvent(value interface{}) *MatrixControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *MatrixControl) StaticClassName(value interface{}) *MatrixControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 配置 label className
 */
func (a *MatrixControl) LabelClassName(value interface{}) *MatrixControl {
    a.Set("labelClassName", value)
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
 * 静态展示空值占位
 */
func (a *MatrixControl) StaticPlaceholder(value interface{}) *MatrixControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *MatrixControl) LabelRemark(value interface{}) *MatrixControl {
    a.Set("labelRemark", value)
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
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *MatrixControl) ClearValueOnHidden(value interface{}) *MatrixControl {
    a.Set("clearValueOnHidden", value)
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
 * 是否只读
 */
func (a *MatrixControl) ReadOnly(value interface{}) *MatrixControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 配置 input className
 */
func (a *MatrixControl) InputClassName(value interface{}) *MatrixControl {
    a.Set("inputClassName", value)
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
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *MatrixControl) Size(value interface{}) *MatrixControl {
    a.Set("size", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *MatrixControl) ValidateOnChange(value interface{}) *MatrixControl {
    a.Set("validateOnChange", value)
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
 * 是否显示
 */
func (a *MatrixControl) Visible(value interface{}) *MatrixControl {
    a.Set("visible", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *MatrixControl) DescriptionClassName(value interface{}) *MatrixControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 设置单选模式，multiple为false时有效
 */
func (a *MatrixControl) SingleSelectMode(value interface{}) *MatrixControl {
    a.Set("singleSelectMode", value)
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
 * 编辑器配置，运行时可以忽略
 */
func (a *MatrixControl) EditorSetting(value interface{}) *MatrixControl {
    a.Set("editorSetting", value)
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
 * 可以组件级别用来关闭移动端样式
 */
func (a *MatrixControl) UseMobileUI(value interface{}) *MatrixControl {
    a.Set("useMobileUI", value)
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
 * 配置singleSelectMode时设置为false
 */
func (a *MatrixControl) Multiple(value interface{}) *MatrixControl {
    a.Set("multiple", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *MatrixControl) StaticOn(value interface{}) *MatrixControl {
    a.Set("staticOn", value)
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
 * 容器 css 类名
 */
func (a *MatrixControl) ClassName(value interface{}) *MatrixControl {
    a.Set("className", value)
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
 * 表单项类型
 */
func (a *MatrixControl) Type(value interface{}) *MatrixControl {
    a.Set("type", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *MatrixControl) VisibleOn(value interface{}) *MatrixControl {
    a.Set("visibleOn", value)
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
 * 静态展示表单项Label类名
 */
func (a *MatrixControl) StaticLabelClassName(value interface{}) *MatrixControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *MatrixControl) StaticInputClassName(value interface{}) *MatrixControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *MatrixControl) ExtraName(value interface{}) *MatrixControl {
    a.Set("extraName", value)
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
 */
func (a *MatrixControl) Columns(value interface{}) *MatrixControl {
    a.Set("columns", value)
    return a
}
