package renderers


/**
 * Static 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/static
 *

*/
type StaticExactControl struct {
	*BaseRenderer
}

func NewStaticExactControl() *StaticExactControl {
    a := &StaticExactControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "static")
    return a
}
/**
 * 输入提示，聚焦的时候显示
 */
func (a *StaticExactControl) Hint(value string) *StaticExactControl {
    a.Set("hint", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *StaticExactControl) SubmitOnChange(value string) *StaticExactControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *StaticExactControl) ClearValueOnHidden(value string) *StaticExactControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *StaticExactControl) EditorSetting(value string) *StaticExactControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *StaticExactControl) Name(value string) *StaticExactControl {
    a.Set("name", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *StaticExactControl) Remark(value string) *StaticExactControl {
    a.Set("remark", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *StaticExactControl) Description(value string) *StaticExactControl {
    a.Set("description", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *StaticExactControl) ValidationErrors(value string) *StaticExactControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 表单项类型
 */
func (a *StaticExactControl) Type(value string) *StaticExactControl {
    a.Set("type", value)
    return a
}

/**
 * 配置快速编辑功能
 */
func (a *StaticExactControl) QuickEdit(value string) *StaticExactControl {
    a.Set("quickEdit", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *StaticExactControl) StaticInputClassName(value string) *StaticExactControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *StaticExactControl) HiddenOn(value string) *StaticExactControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *StaticExactControl) ValidateOnChange(value string) *StaticExactControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 配置 input className
 */
func (a *StaticExactControl) InputClassName(value string) *StaticExactControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *StaticExactControl) DisabledOn(value string) *StaticExactControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 描述标题
 */
func (a *StaticExactControl) Label(value string) *StaticExactControl {
    a.Set("label", value)
    return a
}

/**
 * 是否禁用
 */
func (a *StaticExactControl) Disabled(value string) *StaticExactControl {
    a.Set("disabled", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *StaticExactControl) Size(value string) *StaticExactControl {
    a.Set("size", value)
    return a
}

/**
 */
func (a *StaticExactControl) Desc(value string) *StaticExactControl {
    a.Set("desc", value)
    return a
}

/**
 */
func (a *StaticExactControl) Validations(value string) *StaticExactControl {
    a.Set("validations", value)
    return a
}

/**
 * 组件样式
 */
func (a *StaticExactControl) Style(value string) *StaticExactControl {
    a.Set("style", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *StaticExactControl) ExtraName(value string) *StaticExactControl {
    a.Set("extraName", value)
    return a
}

/**
 * 占位符
 */
func (a *StaticExactControl) Placeholder(value string) *StaticExactControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 内容模板， 支持 HTML
 */
func (a *StaticExactControl) Tpl(value string) *StaticExactControl {
    a.Set("tpl", value)
    return a
}

/**
 * 边框模式，默认是无边框的
 * 可选值: full | half | none
 */
func (a *StaticExactControl) BorderMode(value string) *StaticExactControl {
    a.Set("borderMode", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *StaticExactControl) Inline(value string) *StaticExactControl {
    a.Set("inline", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *StaticExactControl) Static(value string) *StaticExactControl {
    a.Set("static", value)
    return a
}

/**
 * 是否为必填
 */
func (a *StaticExactControl) Required(value string) *StaticExactControl {
    a.Set("required", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *StaticExactControl) Hidden(value string) *StaticExactControl {
    a.Set("hidden", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *StaticExactControl) UseMobileUI(value string) *StaticExactControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *StaticExactControl) DescriptionClassName(value string) *StaticExactControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *StaticExactControl) Mode(value string) *StaticExactControl {
    a.Set("mode", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *StaticExactControl) StaticClassName(value string) *StaticExactControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *StaticExactControl) Id(value string) *StaticExactControl {
    a.Set("id", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *StaticExactControl) OnEvent(value string) *StaticExactControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *StaticExactControl) StaticOn(value string) *StaticExactControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 配置 label className
 */
func (a *StaticExactControl) LabelClassName(value string) *StaticExactControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *StaticExactControl) Value(value string) *StaticExactControl {
    a.Set("value", value)
    return a
}

/**
 * 内容模板，不支持 HTML
 */
func (a *StaticExactControl) Text(value string) *StaticExactControl {
    a.Set("text", value)
    return a
}

/**
 * 配置点击复制功能
 */
func (a *StaticExactControl) Copyable(value string) *StaticExactControl {
    a.Set("copyable", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *StaticExactControl) VisibleOn(value string) *StaticExactControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 只读条件
 */
func (a *StaticExactControl) ReadOnlyOn(value string) *StaticExactControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *StaticExactControl) StaticLabelClassName(value string) *StaticExactControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *StaticExactControl) StaticPlaceholder(value string) *StaticExactControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *StaticExactControl) StaticSchema(value string) *StaticExactControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 描述标题
 */
func (a *StaticExactControl) LabelAlign(value string) *StaticExactControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *StaticExactControl) LabelWidth(value string) *StaticExactControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *StaticExactControl) Horizontal(value string) *StaticExactControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *StaticExactControl) Width(value string) *StaticExactControl {
    a.Set("width", value)
    return a
}

/**
 * 是否显示
 */
func (a *StaticExactControl) Visible(value string) *StaticExactControl {
    a.Set("visible", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *StaticExactControl) ClassName(value string) *StaticExactControl {
    a.Set("className", value)
    return a
}

/**
 * 是否只读
 */
func (a *StaticExactControl) ReadOnly(value string) *StaticExactControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *StaticExactControl) ValidateApi(value string) *StaticExactControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 配置查看详情功能
 */
func (a *StaticExactControl) PopOver(value string) *StaticExactControl {
    a.Set("popOver", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *StaticExactControl) LabelRemark(value string) *StaticExactControl {
    a.Set("labelRemark", value)
    return a
}
