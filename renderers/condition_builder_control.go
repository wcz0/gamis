package renderers


/**
 * 条件组合控件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/condition-builder
 *

*/
type ConditionBuilderControl struct {
	*BaseRenderer
}

func NewConditionBuilderControl() *ConditionBuilderControl {
    a := &ConditionBuilderControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "condition-builder")
    return a
}
/**
 * 事件动作配置
 */
func (a *ConditionBuilderControl) OnEvent(value string) *ConditionBuilderControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *ConditionBuilderControl) Remark(value string) *ConditionBuilderControl {
    a.Set("remark", value)
    return a
}

/**
 * 只读条件
 */
func (a *ConditionBuilderControl) ReadOnlyOn(value string) *ConditionBuilderControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *ConditionBuilderControl) LabelWidth(value string) *ConditionBuilderControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 是否只读
 */
func (a *ConditionBuilderControl) ReadOnly(value string) *ConditionBuilderControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 是否显示
 */
func (a *ConditionBuilderControl) Visible(value string) *ConditionBuilderControl {
    a.Set("visible", value)
    return a
}

/**
 * 组件样式
 */
func (a *ConditionBuilderControl) Style(value string) *ConditionBuilderControl {
    a.Set("style", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *ConditionBuilderControl) Value(value string) *ConditionBuilderControl {
    a.Set("value", value)
    return a
}

/**
 * 其他配置
 */
func (a *ConditionBuilderControl) Config(value string) *ConditionBuilderControl {
    a.Set("config", value)
    return a
}

/**
 * 是否显示并或切换键按钮，只在简单模式下有用
 */
func (a *ConditionBuilderControl) ShowANDOR(value string) *ConditionBuilderControl {
    a.Set("showANDOR", value)
    return a
}

/**
 * 是否禁用
 */
func (a *ConditionBuilderControl) Disabled(value string) *ConditionBuilderControl {
    a.Set("disabled", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *ConditionBuilderControl) DisabledOn(value string) *ConditionBuilderControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *ConditionBuilderControl) VisibleOn(value string) *ConditionBuilderControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *ConditionBuilderControl) Mode(value string) *ConditionBuilderControl {
    a.Set("mode", value)
    return a
}

/**
 */
func (a *ConditionBuilderControl) Validations(value string) *ConditionBuilderControl {
    a.Set("validations", value)
    return a
}

/**
 * 非内嵌模式时 弹窗触发icon
 */
func (a *ConditionBuilderControl) PickerIcon(value string) *ConditionBuilderControl {
    a.Set("pickerIcon", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *ConditionBuilderControl) StaticClassName(value string) *ConditionBuilderControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 描述标题
 */
func (a *ConditionBuilderControl) Label(value string) *ConditionBuilderControl {
    a.Set("label", value)
    return a
}

/**
 * 配置 label className
 */
func (a *ConditionBuilderControl) LabelClassName(value string) *ConditionBuilderControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *ConditionBuilderControl) Hint(value string) *ConditionBuilderControl {
    a.Set("hint", value)
    return a
}

/**
 * 是否为必填
 */
func (a *ConditionBuilderControl) Required(value string) *ConditionBuilderControl {
    a.Set("required", value)
    return a
}

/**
 * 表达式：控制按钮“添加条件组”的显示
 */
func (a *ConditionBuilderControl) AddGroupBtnVisibleOn(value string) *ConditionBuilderControl {
    a.Set("addGroupBtnVisibleOn", value)
    return a
}

/**
 * 描述标题
 */
func (a *ConditionBuilderControl) LabelAlign(value string) *ConditionBuilderControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *ConditionBuilderControl) Name(value string) *ConditionBuilderControl {
    a.Set("name", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *ConditionBuilderControl) ExtraName(value string) *ConditionBuilderControl {
    a.Set("extraName", value)
    return a
}

/**
 * 配置 input className
 */
func (a *ConditionBuilderControl) InputClassName(value string) *ConditionBuilderControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *ConditionBuilderControl) ClearValueOnHidden(value string) *ConditionBuilderControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 指定为
 */
func (a *ConditionBuilderControl) Type(value string) *ConditionBuilderControl {
    a.Set("type", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *ConditionBuilderControl) Hidden(value string) *ConditionBuilderControl {
    a.Set("hidden", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *ConditionBuilderControl) StaticPlaceholder(value string) *ConditionBuilderControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *ConditionBuilderControl) StaticSchema(value string) *ConditionBuilderControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *ConditionBuilderControl) EditorSetting(value string) *ConditionBuilderControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 通过远程拉取配置项
 */
func (a *ConditionBuilderControl) Source(value string) *ConditionBuilderControl {
    a.Set("source", value)
    return a
}

/**
 * 展现模式
 * 可选值: simple | full
 */
func (a *ConditionBuilderControl) BuilderMode(value string) *ConditionBuilderControl {
    a.Set("builderMode", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *ConditionBuilderControl) StaticOn(value string) *ConditionBuilderControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *ConditionBuilderControl) Size(value string) *ConditionBuilderControl {
    a.Set("size", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *ConditionBuilderControl) UseMobileUI(value string) *ConditionBuilderControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *ConditionBuilderControl) Horizontal(value string) *ConditionBuilderControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 内嵌模式，默认为 true
 */
func (a *ConditionBuilderControl) Embed(value string) *ConditionBuilderControl {
    a.Set("embed", value)
    return a
}

/**
 */
func (a *ConditionBuilderControl) AddBtnVisibleOn(value string) *ConditionBuilderControl {
    a.Set("addBtnVisibleOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *ConditionBuilderControl) Static(value string) *ConditionBuilderControl {
    a.Set("static", value)
    return a
}

/**
 */
func (a *ConditionBuilderControl) Desc(value string) *ConditionBuilderControl {
    a.Set("desc", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *ConditionBuilderControl) Inline(value string) *ConditionBuilderControl {
    a.Set("inline", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *ConditionBuilderControl) ValidateApi(value string) *ConditionBuilderControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 函数集合
 */
func (a *ConditionBuilderControl) Funcs(value string) *ConditionBuilderControl {
    a.Set("funcs", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *ConditionBuilderControl) ClassName(value string) *ConditionBuilderControl {
    a.Set("className", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *ConditionBuilderControl) LabelRemark(value string) *ConditionBuilderControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *ConditionBuilderControl) ValidateOnChange(value string) *ConditionBuilderControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *ConditionBuilderControl) Description(value string) *ConditionBuilderControl {
    a.Set("description", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *ConditionBuilderControl) DescriptionClassName(value string) *ConditionBuilderControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *ConditionBuilderControl) Id(value string) *ConditionBuilderControl {
    a.Set("id", value)
    return a
}

/**
 * 字段集合
 */
func (a *ConditionBuilderControl) Fields(value string) *ConditionBuilderControl {
    a.Set("fields", value)
    return a
}

/**
 * 将字段输入控件变成公式编辑器。
 */
func (a *ConditionBuilderControl) Formula(value string) *ConditionBuilderControl {
    a.Set("formula", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *ConditionBuilderControl) StaticLabelClassName(value string) *ConditionBuilderControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *ConditionBuilderControl) StaticInputClassName(value string) *ConditionBuilderControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *ConditionBuilderControl) SubmitOnChange(value string) *ConditionBuilderControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 占位符
 */
func (a *ConditionBuilderControl) Placeholder(value string) *ConditionBuilderControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *ConditionBuilderControl) ValidationErrors(value string) *ConditionBuilderControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *ConditionBuilderControl) Width(value string) *ConditionBuilderControl {
    a.Set("width", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *ConditionBuilderControl) HiddenOn(value string) *ConditionBuilderControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否可拖拽，默认为 true
 */
func (a *ConditionBuilderControl) Draggable(value string) *ConditionBuilderControl {
    a.Set("draggable", value)
    return a
}
