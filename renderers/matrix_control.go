package renderers


/**
 * Matrix 选择控件。适合做权限勾选。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/matrix
 *

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
 * 配置 input className
 */
func (a *MatrixControl) InputClassName(value string) *MatrixControl {
    a.Set("inputClassName", value)
    return a
}

/**
 */
func (a *MatrixControl) Validations(value string) *MatrixControl {
    a.Set("validations", value)
    return a
}

/**
 * 可用来通过 API 拉取 options。
 */
func (a *MatrixControl) Source(value string) *MatrixControl {
    a.Set("source", value)
    return a
}

/**
 */
func (a *MatrixControl) Columns(value string) *MatrixControl {
    a.Set("columns", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *MatrixControl) HiddenOn(value string) *MatrixControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *MatrixControl) StaticLabelClassName(value string) *MatrixControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *MatrixControl) Style(value string) *MatrixControl {
    a.Set("style", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *MatrixControl) UseMobileUI(value string) *MatrixControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *MatrixControl) StaticPlaceholder(value string) *MatrixControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 是否为必填
 */
func (a *MatrixControl) Required(value string) *MatrixControl {
    a.Set("required", value)
    return a
}

/**
 * 表单项类型
 */
func (a *MatrixControl) Type(value string) *MatrixControl {
    a.Set("type", value)
    return a
}

/**
 * 描述标题
 */
func (a *MatrixControl) LabelAlign(value string) *MatrixControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *MatrixControl) Hint(value string) *MatrixControl {
    a.Set("hint", value)
    return a
}

/**
 * 只读条件
 */
func (a *MatrixControl) ReadOnlyOn(value string) *MatrixControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *MatrixControl) Description(value string) *MatrixControl {
    a.Set("description", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *MatrixControl) ClassName(value string) *MatrixControl {
    a.Set("className", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *MatrixControl) Name(value string) *MatrixControl {
    a.Set("name", value)
    return a
}

/**
 */
func (a *MatrixControl) Desc(value string) *MatrixControl {
    a.Set("desc", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *MatrixControl) DescriptionClassName(value string) *MatrixControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *MatrixControl) Mode(value string) *MatrixControl {
    a.Set("mode", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *MatrixControl) DisabledOn(value string) *MatrixControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *MatrixControl) Static(value string) *MatrixControl {
    a.Set("static", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *MatrixControl) EditorSetting(value string) *MatrixControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *MatrixControl) LabelRemark(value string) *MatrixControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *MatrixControl) Hidden(value string) *MatrixControl {
    a.Set("hidden", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *MatrixControl) StaticClassName(value string) *MatrixControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 描述标题
 */
func (a *MatrixControl) Label(value string) *MatrixControl {
    a.Set("label", value)
    return a
}

/**
 * 配置 label className
 */
func (a *MatrixControl) LabelClassName(value string) *MatrixControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 占位符
 */
func (a *MatrixControl) Placeholder(value string) *MatrixControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *MatrixControl) ValidateApi(value string) *MatrixControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *MatrixControl) Size(value string) *MatrixControl {
    a.Set("size", value)
    return a
}

/**
 * 是否禁用
 */
func (a *MatrixControl) Disabled(value string) *MatrixControl {
    a.Set("disabled", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *MatrixControl) VisibleOn(value string) *MatrixControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *MatrixControl) SubmitOnChange(value string) *MatrixControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 行标题说明
 */
func (a *MatrixControl) RowLabel(value string) *MatrixControl {
    a.Set("rowLabel", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *MatrixControl) LabelWidth(value string) *MatrixControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 配置singleSelectMode时设置为false
 */
func (a *MatrixControl) Multiple(value string) *MatrixControl {
    a.Set("multiple", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *MatrixControl) Inline(value string) *MatrixControl {
    a.Set("inline", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *MatrixControl) ClearValueOnHidden(value string) *MatrixControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *MatrixControl) Id(value string) *MatrixControl {
    a.Set("id", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *MatrixControl) OnEvent(value string) *MatrixControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *MatrixControl) ExtraName(value string) *MatrixControl {
    a.Set("extraName", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *MatrixControl) Remark(value string) *MatrixControl {
    a.Set("remark", value)
    return a
}

/**
 * 是否显示
 */
func (a *MatrixControl) Visible(value string) *MatrixControl {
    a.Set("visible", value)
    return a
}

/**
 * 是否只读
 */
func (a *MatrixControl) ReadOnly(value string) *MatrixControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *MatrixControl) Value(value string) *MatrixControl {
    a.Set("value", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *MatrixControl) StaticOn(value string) *MatrixControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *MatrixControl) Horizontal(value string) *MatrixControl {
    a.Set("horizontal", value)
    return a
}

/**
 */
func (a *MatrixControl) Rows(value string) *MatrixControl {
    a.Set("rows", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *MatrixControl) Width(value string) *MatrixControl {
    a.Set("width", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *MatrixControl) StaticInputClassName(value string) *MatrixControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *MatrixControl) StaticSchema(value string) *MatrixControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *MatrixControl) ValidateOnChange(value string) *MatrixControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 设置单选模式，multiple为false时有效
 */
func (a *MatrixControl) SingleSelectMode(value string) *MatrixControl {
    a.Set("singleSelectMode", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *MatrixControl) ValidationErrors(value string) *MatrixControl {
    a.Set("validationErrors", value)
    return a
}
