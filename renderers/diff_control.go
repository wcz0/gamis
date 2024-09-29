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


func (a *DiffControl) Set(name string, value interface{}) *DiffControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *DiffControl) Id(value interface{}) *DiffControl {
    a.Set("id", value)
    return a
}

/**
 */
func (a *DiffControl) StaticSchema(value interface{}) *DiffControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *DiffControl) Name(value interface{}) *DiffControl {
    a.Set("name", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *DiffControl) Width(value interface{}) *DiffControl {
    a.Set("width", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *DiffControl) DisabledOn(value interface{}) *DiffControl {
    a.Set("disabledOn", value)
    return a
}

/**
 */
func (a *DiffControl) TestIdBuilder(value interface{}) *DiffControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 配置 label className
 */
func (a *DiffControl) LabelClassName(value interface{}) *DiffControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *DiffControl) Value(value interface{}) *DiffControl {
    a.Set("value", value)
    return a
}

/**
 * 是否显示
 */
func (a *DiffControl) Visible(value interface{}) *DiffControl {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *DiffControl) StaticClassName(value interface{}) *DiffControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *DiffControl) LabelRemark(value interface{}) *DiffControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *DiffControl) SubmitOnChange(value interface{}) *DiffControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *DiffControl) Horizontal(value interface{}) *DiffControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *DiffControl) Inline(value interface{}) *DiffControl {
    a.Set("inline", value)
    return a
}

/**
 * 指定为 Diff 编辑器
 */
func (a *DiffControl) Type(value interface{}) *DiffControl {
    a.Set("type", value)
    return a
}

/**
 * 语言，参考 monaco-editor
 */
func (a *DiffControl) Language(value interface{}) *DiffControl {
    a.Set("language", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *DiffControl) ClassName(value interface{}) *DiffControl {
    a.Set("className", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *DiffControl) StaticOn(value interface{}) *DiffControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 组件样式
 */
func (a *DiffControl) Style(value interface{}) *DiffControl {
    a.Set("style", value)
    return a
}

/**
 * 占位符
 */
func (a *DiffControl) Placeholder(value interface{}) *DiffControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否为必填
 */
func (a *DiffControl) Required(value interface{}) *DiffControl {
    a.Set("required", value)
    return a
}

/**
 */
func (a *DiffControl) Validations(value interface{}) *DiffControl {
    a.Set("validations", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *DiffControl) ValidateApi(value interface{}) *DiffControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 是否禁用
 */
func (a *DiffControl) Disabled(value interface{}) *DiffControl {
    a.Set("disabled", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *DiffControl) StaticPlaceholder(value interface{}) *DiffControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *DiffControl) Size(value interface{}) *DiffControl {
    a.Set("size", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *DiffControl) AutoFill(value interface{}) *DiffControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *DiffControl) OnEvent(value interface{}) *DiffControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *DiffControl) Hint(value interface{}) *DiffControl {
    a.Set("hint", value)
    return a
}

/**
 * 是否只读
 */
func (a *DiffControl) ReadOnly(value interface{}) *DiffControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *DiffControl) ValidationErrors(value interface{}) *DiffControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *DiffControl) StaticInputClassName(value interface{}) *DiffControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *DiffControl) ValidateOnChange(value interface{}) *DiffControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *DiffControl) ClearValueOnHidden(value interface{}) *DiffControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *DiffControl) Hidden(value interface{}) *DiffControl {
    a.Set("hidden", value)
    return a
}

/**
 * 描述标题
 */
func (a *DiffControl) LabelAlign(value interface{}) *DiffControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 编辑器配置
 */
func (a *DiffControl) Options(value interface{}) *DiffControl {
    a.Set("options", value)
    return a
}

/**
 */
func (a *DiffControl) Desc(value interface{}) *DiffControl {
    a.Set("desc", value)
    return a
}

/**
 */
func (a *DiffControl) InitAutoFill(value interface{}) *DiffControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *DiffControl) UseMobileUI(value interface{}) *DiffControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *DiffControl) StaticLabelClassName(value interface{}) *DiffControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *DiffControl) EditorSetting(value interface{}) *DiffControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *DiffControl) ExtraName(value interface{}) *DiffControl {
    a.Set("extraName", value)
    return a
}

/**
 * 只读条件
 */
func (a *DiffControl) ReadOnlyOn(value interface{}) *DiffControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *DiffControl) Mode(value interface{}) *DiffControl {
    a.Set("mode", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *DiffControl) VisibleOn(value interface{}) *DiffControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 描述标题
 */
func (a *DiffControl) Label(value interface{}) *DiffControl {
    a.Set("label", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *DiffControl) LabelWidth(value interface{}) *DiffControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *DiffControl) DescriptionClassName(value interface{}) *DiffControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *DiffControl) HiddenOn(value interface{}) *DiffControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *DiffControl) Static(value interface{}) *DiffControl {
    a.Set("static", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *DiffControl) Description(value interface{}) *DiffControl {
    a.Set("description", value)
    return a
}

/**
 * 配置 input className
 */
func (a *DiffControl) InputClassName(value interface{}) *DiffControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *DiffControl) Remark(value interface{}) *DiffControl {
    a.Set("remark", value)
    return a
}

/**
 */
func (a *DiffControl) Row(value interface{}) *DiffControl {
    a.Set("row", value)
    return a
}

/**
 * 左侧面板的值， 支持取变量。
 */
func (a *DiffControl) DiffValue(value interface{}) *DiffControl {
    a.Set("diffValue", value)
    return a
}
