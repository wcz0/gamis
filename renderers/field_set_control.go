package renderers


/**
 * FieldSet 表单项集合 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/fieldset
 *

*/
type FieldSetControl struct {
	*BaseRenderer
}

func NewFieldSetControl() *FieldSetControl {
    a := &FieldSetControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "fieldset")
    a.Set("titlePosition", "top")
    return a
}
/**
 * 标识
 */
func (a *FieldSetControl) Key(value string) *FieldSetControl {
    a.Set("key", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *FieldSetControl) DisabledOn(value string) *FieldSetControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *FieldSetControl) HiddenOn(value string) *FieldSetControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 收起的标题
 */
func (a *FieldSetControl) CollapseTitle(value string) *FieldSetControl {
    a.Set("collapseTitle", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *FieldSetControl) Name(value string) *FieldSetControl {
    a.Set("name", value)
    return a
}

/**
 */
func (a *FieldSetControl) Validations(value string) *FieldSetControl {
    a.Set("validations", value)
    return a
}

/**
 * 是否可折叠
 */
func (a *FieldSetControl) Collapsable(value string) *FieldSetControl {
    a.Set("collapsable", value)
    return a
}

/**
 * 自定义切换图标
 */
func (a *FieldSetControl) ExpandIcon(value string) *FieldSetControl {
    a.Set("expandIcon", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *FieldSetControl) ClassName(value string) *FieldSetControl {
    a.Set("className", value)
    return a
}

/**
 * 描述标题
 */
func (a *FieldSetControl) LabelAlign(value string) *FieldSetControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *FieldSetControl) LabelWidth(value string) *FieldSetControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *FieldSetControl) Width(value string) *FieldSetControl {
    a.Set("width", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *FieldSetControl) StaticLabelClassName(value string) *FieldSetControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 配置 label className
 */
func (a *FieldSetControl) LabelClassName(value string) *FieldSetControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *FieldSetControl) ValidateApi(value string) *FieldSetControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 配置 Body 容器 className
 */
func (a *FieldSetControl) BodyClassName(value string) *FieldSetControl {
    a.Set("bodyClassName", value)
    return a
}

/**
 * 是否显示
 */
func (a *FieldSetControl) Visible(value string) *FieldSetControl {
    a.Set("visible", value)
    return a
}

/**
 * 只读条件
 */
func (a *FieldSetControl) ReadOnlyOn(value string) *FieldSetControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *FieldSetControl) Mode(value string) *FieldSetControl {
    a.Set("mode", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *FieldSetControl) Inline(value string) *FieldSetControl {
    a.Set("inline", value)
    return a
}

/**
 * 配置 input className
 */
func (a *FieldSetControl) InputClassName(value string) *FieldSetControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 点开时才加载内容
 */
func (a *FieldSetControl) MountOnEnter(value string) *FieldSetControl {
    a.Set("mountOnEnter", value)
    return a
}

/**
 */
func (a *FieldSetControl) Desc(value string) *FieldSetControl {
    a.Set("desc", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *FieldSetControl) StaticClassName(value string) *FieldSetControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *FieldSetControl) SubmitOnChange(value string) *FieldSetControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 标题
 */
func (a *FieldSetControl) Title(value string) *FieldSetControl {
    a.Set("title", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *FieldSetControl) ClearValueOnHidden(value string) *FieldSetControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 指定为表单项集合
 * 可选值: fieldset | fieldSet
 */
func (a *FieldSetControl) Type(value string) *FieldSetControl {
    a.Set("type", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *FieldSetControl) OnEvent(value string) *FieldSetControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *FieldSetControl) Static(value string) *FieldSetControl {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *FieldSetControl) StaticInputClassName(value string) *FieldSetControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *FieldSetControl) Style(value string) *FieldSetControl {
    a.Set("style", value)
    return a
}

/**
 * 描述标题
 */
func (a *FieldSetControl) Label(value string) *FieldSetControl {
    a.Set("label", value)
    return a
}

/**
 * 是否只读
 */
func (a *FieldSetControl) ReadOnly(value string) *FieldSetControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 图标是否展示
 */
func (a *FieldSetControl) ShowArrow(value string) *FieldSetControl {
    a.Set("showArrow", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *FieldSetControl) VisibleOn(value string) *FieldSetControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *FieldSetControl) EditorSetting(value string) *FieldSetControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 标题展示位置
 * 可选值: top | bottom
 */
func (a *FieldSetControl) TitlePosition(value string) *FieldSetControl {
    a.Set("titlePosition", value)
    return a
}

/**
 * 标题内容分割线
 */
func (a *FieldSetControl) DivideLine(value string) *FieldSetControl {
    a.Set("divideLine", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *FieldSetControl) Id(value string) *FieldSetControl {
    a.Set("id", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *FieldSetControl) StaticPlaceholder(value string) *FieldSetControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 内容区域
 */
func (a *FieldSetControl) Body(value string) *FieldSetControl {
    a.Set("body", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *FieldSetControl) LabelRemark(value string) *FieldSetControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *FieldSetControl) Description(value string) *FieldSetControl {
    a.Set("description", value)
    return a
}

/**
 * 是否为必填
 */
func (a *FieldSetControl) Required(value string) *FieldSetControl {
    a.Set("required", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *FieldSetControl) ValidationErrors(value string) *FieldSetControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *FieldSetControl) ExtraName(value string) *FieldSetControl {
    a.Set("extraName", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *FieldSetControl) Hint(value string) *FieldSetControl {
    a.Set("hint", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *FieldSetControl) ValidateOnChange(value string) *FieldSetControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *FieldSetControl) Value(value string) *FieldSetControl {
    a.Set("value", value)
    return a
}

/**
 * 是否禁用
 */
func (a *FieldSetControl) Disabled(value string) *FieldSetControl {
    a.Set("disabled", value)
    return a
}

/**
 * 控件大小
 * 可选值: xs | sm | md | lg | base
 */
func (a *FieldSetControl) Size(value string) *FieldSetControl {
    a.Set("size", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *FieldSetControl) Hidden(value string) *FieldSetControl {
    a.Set("hidden", value)
    return a
}

/**
 * 如果是水平排版，这个属性可以细化水平排版的左右宽度占比。
 */
func (a *FieldSetControl) SubFormHorizontal(value string) *FieldSetControl {
    a.Set("subFormHorizontal", value)
    return a
}

/**
 * 标题
 */
func (a *FieldSetControl) Header(value string) *FieldSetControl {
    a.Set("header", value)
    return a
}

/**
 * 默认是否折叠
 */
func (a *FieldSetControl) Collapsed(value string) *FieldSetControl {
    a.Set("collapsed", value)
    return a
}

/**
 * 卡片隐藏就销毁内容。
 */
func (a *FieldSetControl) UnmountOnExit(value string) *FieldSetControl {
    a.Set("unmountOnExit", value)
    return a
}

/**
 */
func (a *FieldSetControl) StaticSchema(value string) *FieldSetControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *FieldSetControl) Remark(value string) *FieldSetControl {
    a.Set("remark", value)
    return a
}

/**
 * 标题展示位置
 * 可选值: top | bottom
 */
func (a *FieldSetControl) HeaderPosition(value string) *FieldSetControl {
    a.Set("headerPosition", value)
    return a
}

/**
 * 标题 CSS 类名
 */
func (a *FieldSetControl) HeadingClassName(value string) *FieldSetControl {
    a.Set("headingClassName", value)
    return a
}

/**
 * 收起的标题
 */
func (a *FieldSetControl) CollapseHeader(value string) *FieldSetControl {
    a.Set("collapseHeader", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *FieldSetControl) StaticOn(value string) *FieldSetControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 配置子表单项默认的展示方式。
 * 可选值: normal | inline | horizontal
 */
func (a *FieldSetControl) SubFormMode(value string) *FieldSetControl {
    a.Set("subFormMode", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *FieldSetControl) UseMobileUI(value string) *FieldSetControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *FieldSetControl) DescriptionClassName(value string) *FieldSetControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *FieldSetControl) Horizontal(value string) *FieldSetControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 占位符
 */
func (a *FieldSetControl) Placeholder(value string) *FieldSetControl {
    a.Set("placeholder", value)
    return a
}
