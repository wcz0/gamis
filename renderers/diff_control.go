package renderers


/**
 * Diff 编辑器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/diff

* @author wcz0
* @version 6.2.2
*/
type DiffControl struct {
	*BaseRenderer
}

func NewDiffControl() *DiffControl {
    a := &DiffControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "diff-editor")
    return a
}
/**
 * 是否隐藏
 */
func (a *DiffControl) Hidden(value string) *DiffControl {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *DiffControl) VisibleOn(value string) *DiffControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *DiffControl) Mode(value string) *DiffControl {
    a.Set("mode", value)
    return a
}

/**
 * 占位符
 */
func (a *DiffControl) Placeholder(value string) *DiffControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *DiffControl) ValidateApi(value string) *DiffControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *DiffControl) ValidateOnChange(value string) *DiffControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *DiffControl) StaticInputClassName(value string) *DiffControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 编辑器配置
 */
func (a *DiffControl) Options(value string) *DiffControl {
    a.Set("options", value)
    return a
}

/**
 * 指定为 Diff 编辑器
 */
func (a *DiffControl) Type(value string) *DiffControl {
    a.Set("type", value)
    return a
}

/**
 * 配置 input className
 */
func (a *DiffControl) InputClassName(value string) *DiffControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 左侧面板的值， 支持取变量。
 */
func (a *DiffControl) DiffValue(value string) *DiffControl {
    a.Set("diffValue", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *DiffControl) Static(value string) *DiffControl {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *DiffControl) StaticClassName(value string) *DiffControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *DiffControl) ClassName(value string) *DiffControl {
    a.Set("className", value)
    return a
}

/**
 */
func (a *DiffControl) StaticSchema(value string) *DiffControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *DiffControl) Remark(value string) *DiffControl {
    a.Set("remark", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *DiffControl) LabelRemark(value string) *DiffControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *DiffControl) SubmitOnChange(value string) *DiffControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *DiffControl) Inline(value string) *DiffControl {
    a.Set("inline", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *DiffControl) Value(value string) *DiffControl {
    a.Set("value", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *DiffControl) HiddenOn(value string) *DiffControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *DiffControl) Id(value string) *DiffControl {
    a.Set("id", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *DiffControl) Size(value string) *DiffControl {
    a.Set("size", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *DiffControl) LabelWidth(value string) *DiffControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *DiffControl) Width(value string) *DiffControl {
    a.Set("width", value)
    return a
}

/**
 * 是否为必填
 */
func (a *DiffControl) Required(value string) *DiffControl {
    a.Set("required", value)
    return a
}

/**
 * 描述标题
 */
func (a *DiffControl) Label(value string) *DiffControl {
    a.Set("label", value)
    return a
}

/**
 * 配置 label className
 */
func (a *DiffControl) LabelClassName(value string) *DiffControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 只读条件
 */
func (a *DiffControl) ReadOnlyOn(value string) *DiffControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *DiffControl) DescriptionClassName(value string) *DiffControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *DiffControl) ClearValueOnHidden(value string) *DiffControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *DiffControl) StaticOn(value string) *DiffControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *DiffControl) EditorSetting(value string) *DiffControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *DiffControl) Hint(value string) *DiffControl {
    a.Set("hint", value)
    return a
}

/**
 * 是否只读
 */
func (a *DiffControl) ReadOnly(value string) *DiffControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *DiffControl) ExtraName(value string) *DiffControl {
    a.Set("extraName", value)
    return a
}

/**
 */
func (a *DiffControl) Validations(value string) *DiffControl {
    a.Set("validations", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *DiffControl) ValidationErrors(value string) *DiffControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *DiffControl) OnEvent(value string) *DiffControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *DiffControl) UseMobileUI(value string) *DiffControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 描述标题
 */
func (a *DiffControl) LabelAlign(value string) *DiffControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *DiffControl) Description(value string) *DiffControl {
    a.Set("description", value)
    return a
}

/**
 */
func (a *DiffControl) Desc(value string) *DiffControl {
    a.Set("desc", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *DiffControl) StaticPlaceholder(value string) *DiffControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *DiffControl) StaticLabelClassName(value string) *DiffControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *DiffControl) Name(value string) *DiffControl {
    a.Set("name", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *DiffControl) Horizontal(value string) *DiffControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 是否禁用
 */
func (a *DiffControl) Disabled(value string) *DiffControl {
    a.Set("disabled", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *DiffControl) DisabledOn(value string) *DiffControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *DiffControl) Visible(value string) *DiffControl {
    a.Set("visible", value)
    return a
}

/**
 * 语言，参考 monaco-editor
 */
func (a *DiffControl) Language(value string) *DiffControl {
    a.Set("language", value)
    return a
}

/**
 * 组件样式
 */
func (a *DiffControl) Style(value string) *DiffControl {
    a.Set("style", value)
    return a
}
