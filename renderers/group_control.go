package renderers


/**
 * Group 表单集合渲染器，能让多个表单在一行显示 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/group
 *

*/
type GroupControl struct {
	*BaseRenderer
}

func NewGroupControl() *GroupControl {
    a := &GroupControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "group")
    return a
}
/**
 * 是否隐藏
 */
func (a *GroupControl) Hidden(value string) *GroupControl {
    a.Set("hidden", value)
    return a
}

/**
 * 配置 label className
 */
func (a *GroupControl) LabelClassName(value string) *GroupControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *GroupControl) Hint(value string) *GroupControl {
    a.Set("hint", value)
    return a
}

/**
 * 配置子表单项默认的展示方式。
 * 可选值: normal | inline | horizontal
 */
func (a *GroupControl) SubFormMode(value string) *GroupControl {
    a.Set("subFormMode", value)
    return a
}

/**
 * 是否显示
 */
func (a *GroupControl) Visible(value string) *GroupControl {
    a.Set("visible", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *GroupControl) Name(value string) *GroupControl {
    a.Set("name", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *GroupControl) StaticClassName(value string) *GroupControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *GroupControl) UseMobileUI(value string) *GroupControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *GroupControl) Size(value string) *GroupControl {
    a.Set("size", value)
    return a
}

/**
 * 是否为必填
 */
func (a *GroupControl) Required(value string) *GroupControl {
    a.Set("required", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *GroupControl) ValidateApi(value string) *GroupControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 间隔
 * 可选值: xs | sm | normal
 */
func (a *GroupControl) Gap(value string) *GroupControl {
    a.Set("gap", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *GroupControl) LabelWidth(value string) *GroupControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 配置 input className
 */
func (a *GroupControl) InputClassName(value string) *GroupControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *GroupControl) ClearValueOnHidden(value string) *GroupControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *GroupControl) ClassName(value string) *GroupControl {
    a.Set("className", value)
    return a
}

/**
 * 组件样式
 */
func (a *GroupControl) Style(value string) *GroupControl {
    a.Set("style", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *GroupControl) Inline(value string) *GroupControl {
    a.Set("inline", value)
    return a
}

/**
 */
func (a *GroupControl) Validations(value string) *GroupControl {
    a.Set("validations", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *GroupControl) StaticLabelClassName(value string) *GroupControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *GroupControl) StaticInputClassName(value string) *GroupControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *GroupControl) Width(value string) *GroupControl {
    a.Set("width", value)
    return a
}

/**
 * 是否禁用
 */
func (a *GroupControl) Disabled(value string) *GroupControl {
    a.Set("disabled", value)
    return a
}

/**
 * 描述标题
 */
func (a *GroupControl) Label(value string) *GroupControl {
    a.Set("label", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *GroupControl) Mode(value string) *GroupControl {
    a.Set("mode", value)
    return a
}

/**
 * 表单项类型
 */
func (a *GroupControl) Type(value string) *GroupControl {
    a.Set("type", value)
    return a
}

/**
 * 如果是水平排版，这个属性可以细化水平排版的左右宽度占比。
 */
func (a *GroupControl) SubFormHorizontal(value string) *GroupControl {
    a.Set("subFormHorizontal", value)
    return a
}

/**
 */
func (a *GroupControl) Desc(value string) *GroupControl {
    a.Set("desc", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *GroupControl) ValidationErrors(value string) *GroupControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 只读条件
 */
func (a *GroupControl) ReadOnlyOn(value string) *GroupControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *GroupControl) OnEvent(value string) *GroupControl {
    a.Set("onEvent", value)
    return a
}

/**
 */
func (a *GroupControl) StaticSchema(value string) *GroupControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *GroupControl) SubmitOnChange(value string) *GroupControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *GroupControl) DescriptionClassName(value string) *GroupControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *GroupControl) Value(value string) *GroupControl {
    a.Set("value", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *GroupControl) StaticPlaceholder(value string) *GroupControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *GroupControl) EditorSetting(value string) *GroupControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 描述标题
 */
func (a *GroupControl) LabelAlign(value string) *GroupControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *GroupControl) Remark(value string) *GroupControl {
    a.Set("remark", value)
    return a
}

/**
 * 是否只读
 */
func (a *GroupControl) ReadOnly(value string) *GroupControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *GroupControl) StaticOn(value string) *GroupControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *GroupControl) Horizontal(value string) *GroupControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 占位符
 */
func (a *GroupControl) Placeholder(value string) *GroupControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *GroupControl) DisabledOn(value string) *GroupControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *GroupControl) HiddenOn(value string) *GroupControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *GroupControl) LabelRemark(value string) *GroupControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 配置时垂直摆放还是左右摆放。
 * 可选值: horizontal | vertical
 */
func (a *GroupControl) Direction(value string) *GroupControl {
    a.Set("direction", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *GroupControl) VisibleOn(value string) *GroupControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *GroupControl) Id(value string) *GroupControl {
    a.Set("id", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *GroupControl) Static(value string) *GroupControl {
    a.Set("static", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *GroupControl) ExtraName(value string) *GroupControl {
    a.Set("extraName", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *GroupControl) ValidateOnChange(value string) *GroupControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *GroupControl) Description(value string) *GroupControl {
    a.Set("description", value)
    return a
}

/**
 * FormItem 集合
 */
func (a *GroupControl) Body(value string) *GroupControl {
    a.Set("body", value)
    return a
}
