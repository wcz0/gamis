package renderers


/**
 * SubForm 子表单 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/subform
 *

*/
type SubFormControl struct {
	*BaseRenderer
}

func NewSubFormControl() *SubFormControl {
    a := &SubFormControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-sub-form")
    return a
}
/**
 */
func (a *SubFormControl) StaticSchema(value string) *SubFormControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *SubFormControl) DisabledOn(value string) *SubFormControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *SubFormControl) Id(value string) *SubFormControl {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *SubFormControl) StaticLabelClassName(value string) *SubFormControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *SubFormControl) StaticOn(value string) *SubFormControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *SubFormControl) EditorSetting(value string) *SubFormControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *SubFormControl) Size(value string) *SubFormControl {
    a.Set("size", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *SubFormControl) Remark(value string) *SubFormControl {
    a.Set("remark", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *SubFormControl) SubmitOnChange(value string) *SubFormControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *SubFormControl) Mode(value string) *SubFormControl {
    a.Set("mode", value)
    return a
}

/**
 * 值列表元素的类名
 */
func (a *SubFormControl) ItemsClassName(value string) *SubFormControl {
    a.Set("itemsClassName", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *SubFormControl) Static(value string) *SubFormControl {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *SubFormControl) StaticClassName(value string) *SubFormControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *SubFormControl) Style(value string) *SubFormControl {
    a.Set("style", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *SubFormControl) ExtraName(value string) *SubFormControl {
    a.Set("extraName", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *SubFormControl) Hint(value string) *SubFormControl {
    a.Set("hint", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *SubFormControl) Horizontal(value string) *SubFormControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *SubFormControl) Inline(value string) *SubFormControl {
    a.Set("inline", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *SubFormControl) ClearValueOnHidden(value string) *SubFormControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 是否在左下角显示报错信息
 */
func (a *SubFormControl) ShowErrorMsg(value string) *SubFormControl {
    a.Set("showErrorMsg", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *SubFormControl) LabelRemark(value string) *SubFormControl {
    a.Set("labelRemark", value)
    return a
}

/**
 */
func (a *SubFormControl) Scaffold(value string) *SubFormControl {
    a.Set("scaffold", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *SubFormControl) LabelWidth(value string) *SubFormControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 是否只读
 */
func (a *SubFormControl) ReadOnly(value string) *SubFormControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *SubFormControl) Description(value string) *SubFormControl {
    a.Set("description", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *SubFormControl) DescriptionClassName(value string) *SubFormControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 是否可删除
 */
func (a *SubFormControl) Removable(value string) *SubFormControl {
    a.Set("removable", value)
    return a
}

/**
 * 当值中存在这个字段，则按钮名称将使用此字段的值来展示。
 */
func (a *SubFormControl) LabelField(value string) *SubFormControl {
    a.Set("labelField", value)
    return a
}

/**
 * 按钮默认名称
 */
func (a *SubFormControl) BtnLabel(value string) *SubFormControl {
    a.Set("btnLabel", value)
    return a
}

/**
 * 子表单详情
 */
func (a *SubFormControl) Form(value string) *SubFormControl {
    a.Set("form", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *SubFormControl) Hidden(value string) *SubFormControl {
    a.Set("hidden", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *SubFormControl) ValidateOnChange(value string) *SubFormControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 新增按钮 CSS 类名
 */
func (a *SubFormControl) AddButtonClassName(value string) *SubFormControl {
    a.Set("addButtonClassName", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *SubFormControl) HiddenOn(value string) *SubFormControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 只读条件
 */
func (a *SubFormControl) ReadOnlyOn(value string) *SubFormControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 是否多选
 */
func (a *SubFormControl) Multiple(value string) *SubFormControl {
    a.Set("multiple", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *SubFormControl) Width(value string) *SubFormControl {
    a.Set("width", value)
    return a
}

/**
 * 是否显示
 */
func (a *SubFormControl) Visible(value string) *SubFormControl {
    a.Set("visible", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *SubFormControl) Name(value string) *SubFormControl {
    a.Set("name", value)
    return a
}

/**
 * 指定为 SubForm 子表单
 */
func (a *SubFormControl) Type(value string) *SubFormControl {
    a.Set("type", value)
    return a
}

/**
 * 是否可拖拽排序
 */
func (a *SubFormControl) Draggable(value string) *SubFormControl {
    a.Set("draggable", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *SubFormControl) ClassName(value string) *SubFormControl {
    a.Set("className", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *SubFormControl) UseMobileUI(value string) *SubFormControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *SubFormControl) Desc(value string) *SubFormControl {
    a.Set("desc", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *SubFormControl) ValidationErrors(value string) *SubFormControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *SubFormControl) StaticPlaceholder(value string) *SubFormControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 占位符
 */
func (a *SubFormControl) Placeholder(value string) *SubFormControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *SubFormControl) Value(value string) *SubFormControl {
    a.Set("value", value)
    return a
}

/**
 * 是否可新增
 */
func (a *SubFormControl) Addable(value string) *SubFormControl {
    a.Set("addable", value)
    return a
}

/**
 * 最多个数
 */
func (a *SubFormControl) MaxLength(value string) *SubFormControl {
    a.Set("maxLength", value)
    return a
}

/**
 * 值元素的类名
 */
func (a *SubFormControl) ItemClassName(value string) *SubFormControl {
    a.Set("itemClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *SubFormControl) StaticInputClassName(value string) *SubFormControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *SubFormControl) Validations(value string) *SubFormControl {
    a.Set("validations", value)
    return a
}

/**
 * 新增按钮文字
 */
func (a *SubFormControl) AddButtonText(value string) *SubFormControl {
    a.Set("addButtonText", value)
    return a
}

/**
 * 是否禁用
 */
func (a *SubFormControl) Disabled(value string) *SubFormControl {
    a.Set("disabled", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *SubFormControl) OnEvent(value string) *SubFormControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 描述标题
 */
func (a *SubFormControl) Label(value string) *SubFormControl {
    a.Set("label", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *SubFormControl) ValidateApi(value string) *SubFormControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 最少个数
 */
func (a *SubFormControl) MinLength(value string) *SubFormControl {
    a.Set("minLength", value)
    return a
}

/**
 * 描述标题
 */
func (a *SubFormControl) LabelAlign(value string) *SubFormControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *SubFormControl) VisibleOn(value string) *SubFormControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 配置 label className
 */
func (a *SubFormControl) LabelClassName(value string) *SubFormControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 配置 input className
 */
func (a *SubFormControl) InputClassName(value string) *SubFormControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 是否为必填
 */
func (a *SubFormControl) Required(value string) *SubFormControl {
    a.Set("required", value)
    return a
}

/**
 * 拖拽提示信息
 */
func (a *SubFormControl) DraggableTip(value string) *SubFormControl {
    a.Set("draggableTip", value)
    return a
}
