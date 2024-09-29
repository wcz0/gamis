package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type InputSignature struct {
	*BaseRenderer
}

func NewInputSignature() *InputSignature {
    a := &InputSignature{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-signature")
    return a
}


func (a *InputSignature) Set(name string, value interface{}) *InputSignature {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 是否禁用表达式
 */
func (a *InputSignature) DisabledOn(value interface{}) *InputSignature {
    a.Set("disabledOn", value)
    return a
}

/**
 * 组件样式
 */
func (a *InputSignature) Style(value interface{}) *InputSignature {
    a.Set("style", value)
    return a
}

/**
 * 配置 input className
 */
func (a *InputSignature) InputClassName(value interface{}) *InputSignature {
    a.Set("inputClassName", value)
    return a
}

/**
 * 组件字段颜色
 */
func (a *InputSignature) Color(value interface{}) *InputSignature {
    a.Set("color", value)
    return a
}

/**
 * 表单项类型
 */
func (a *InputSignature) Type(value interface{}) *InputSignature {
    a.Set("type", value)
    return a
}

/**
 * 清空按钮图标
 */
func (a *InputSignature) UndoBtnIcon(value interface{}) *InputSignature {
    a.Set("undoBtnIcon", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *InputSignature) EditorSetting(value interface{}) *InputSignature {
    a.Set("editorSetting", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *InputSignature) UseMobileUI(value interface{}) *InputSignature {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *InputSignature) Name(value interface{}) *InputSignature {
    a.Set("name", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *InputSignature) Horizontal(value interface{}) *InputSignature {
    a.Set("horizontal", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *InputSignature) ValidationErrors(value interface{}) *InputSignature {
    a.Set("validationErrors", value)
    return a
}

/**
 */
func (a *InputSignature) TestIdBuilder(value interface{}) *InputSignature {
    a.Set("testIdBuilder", value)
    return a
}

/**
 */
func (a *InputSignature) Row(value interface{}) *InputSignature {
    a.Set("row", value)
    return a
}

/**
 * 弹窗按钮图标
 */
func (a *InputSignature) EmbedBtnIcon(value interface{}) *InputSignature {
    a.Set("embedBtnIcon", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *InputSignature) OnEvent(value interface{}) *InputSignature {
    a.Set("onEvent", value)
    return a
}

/**
 * 描述标题
 */
func (a *InputSignature) Label(value interface{}) *InputSignature {
    a.Set("label", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *InputSignature) Description(value interface{}) *InputSignature {
    a.Set("description", value)
    return a
}

/**
 */
func (a *InputSignature) Validations(value interface{}) *InputSignature {
    a.Set("validations", value)
    return a
}

/**
 * 弹窗确认按钮图标
 */
func (a *InputSignature) EmbedConfirmIcon(value interface{}) *InputSignature {
    a.Set("embedConfirmIcon", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *InputSignature) Remark(value interface{}) *InputSignature {
    a.Set("remark", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *InputSignature) SubmitOnChange(value interface{}) *InputSignature {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *InputSignature) Inline(value interface{}) *InputSignature {
    a.Set("inline", value)
    return a
}

/**
 */
func (a *InputSignature) InitAutoFill(value interface{}) *InputSignature {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *InputSignature) ClassName(value interface{}) *InputSignature {
    a.Set("className", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *InputSignature) StaticOn(value interface{}) *InputSignature {
    a.Set("staticOn", value)
    return a
}

/**
 * 配置 label className
 */
func (a *InputSignature) LabelClassName(value interface{}) *InputSignature {
    a.Set("labelClassName", value)
    return a
}

/**
 * 清空按钮名称
 */
func (a *InputSignature) ClearBtnLabel(value interface{}) *InputSignature {
    a.Set("clearBtnLabel", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *InputSignature) VisibleOn(value interface{}) *InputSignature {
    a.Set("visibleOn", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *InputSignature) Hint(value interface{}) *InputSignature {
    a.Set("hint", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *InputSignature) Mode(value interface{}) *InputSignature {
    a.Set("mode", value)
    return a
}

/**
 * 组件宽度，默认占满父容器
 */
func (a *InputSignature) Width(value interface{}) *InputSignature {
    a.Set("width", value)
    return a
}

/**
 * 确认按钮名称
 */
func (a *InputSignature) ConfirmBtnLabel(value interface{}) *InputSignature {
    a.Set("confirmBtnLabel", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *InputSignature) StaticInputClassName(value interface{}) *InputSignature {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *InputSignature) AutoFill(value interface{}) *InputSignature {
    a.Set("autoFill", value)
    return a
}

/**
 * 组件背景颜色
 */
func (a *InputSignature) BgColor(value interface{}) *InputSignature {
    a.Set("bgColor", value)
    return a
}

/**
 * 弹窗取消按钮名称
 */
func (a *InputSignature) EbmedCancelLabel(value interface{}) *InputSignature {
    a.Set("ebmedCancelLabel", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *InputSignature) Hidden(value interface{}) *InputSignature {
    a.Set("hidden", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *InputSignature) Static(value interface{}) *InputSignature {
    a.Set("static", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *InputSignature) StaticPlaceholder(value interface{}) *InputSignature {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 是否只读
 */
func (a *InputSignature) ReadOnly(value interface{}) *InputSignature {
    a.Set("readOnly", value)
    return a
}

/**
 * 组件高度，默认占满父容器
 */
func (a *InputSignature) Height(value interface{}) *InputSignature {
    a.Set("height", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *InputSignature) HiddenOn(value interface{}) *InputSignature {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否为必填
 */
func (a *InputSignature) Required(value interface{}) *InputSignature {
    a.Set("required", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *InputSignature) Size(value interface{}) *InputSignature {
    a.Set("size", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *InputSignature) StaticClassName(value interface{}) *InputSignature {
    a.Set("staticClassName", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *InputSignature) DescriptionClassName(value interface{}) *InputSignature {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 撤销按钮名称
 */
func (a *InputSignature) UndoBtnLabel(value interface{}) *InputSignature {
    a.Set("undoBtnLabel", value)
    return a
}

/**
 * 确认按钮图标
 */
func (a *InputSignature) ConfirmBtnIcon(value interface{}) *InputSignature {
    a.Set("confirmBtnIcon", value)
    return a
}

/**
 * 是否显示
 */
func (a *InputSignature) Visible(value interface{}) *InputSignature {
    a.Set("visible", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *InputSignature) ExtraName(value interface{}) *InputSignature {
    a.Set("extraName", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *InputSignature) ValidateOnChange(value interface{}) *InputSignature {
    a.Set("validateOnChange", value)
    return a
}

/**
 */
func (a *InputSignature) Desc(value interface{}) *InputSignature {
    a.Set("desc", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *InputSignature) ValidateApi(value interface{}) *InputSignature {
    a.Set("validateApi", value)
    return a
}

/**
 * 弹窗按钮文案
 */
func (a *InputSignature) EmbedBtnLabel(value interface{}) *InputSignature {
    a.Set("embedBtnLabel", value)
    return a
}

/**
 * 是否禁用
 */
func (a *InputSignature) Disabled(value interface{}) *InputSignature {
    a.Set("disabled", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *InputSignature) LabelWidth(value interface{}) *InputSignature {
    a.Set("labelWidth", value)
    return a
}

/**
 * 占位符
 */
func (a *InputSignature) Placeholder(value interface{}) *InputSignature {
    a.Set("placeholder", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *InputSignature) Value(value interface{}) *InputSignature {
    a.Set("value", value)
    return a
}

/**
 * 弹窗取消按钮图标
 */
func (a *InputSignature) EbmedCancelIcon(value interface{}) *InputSignature {
    a.Set("ebmedCancelIcon", value)
    return a
}

/**
 * 弹窗确认按钮名称
 */
func (a *InputSignature) EmbedConfirmLabel(value interface{}) *InputSignature {
    a.Set("embedConfirmLabel", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *InputSignature) Id(value interface{}) *InputSignature {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *InputSignature) StaticLabelClassName(value interface{}) *InputSignature {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *InputSignature) StaticSchema(value interface{}) *InputSignature {
    a.Set("staticSchema", value)
    return a
}

/**
 * 描述标题
 */
func (a *InputSignature) LabelAlign(value interface{}) *InputSignature {
    a.Set("labelAlign", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *InputSignature) LabelRemark(value interface{}) *InputSignature {
    a.Set("labelRemark", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *InputSignature) ClearValueOnHidden(value interface{}) *InputSignature {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 只读条件
 */
func (a *InputSignature) ReadOnlyOn(value interface{}) *InputSignature {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 清空按钮图标
 */
func (a *InputSignature) ClearBtnIcon(value interface{}) *InputSignature {
    a.Set("clearBtnIcon", value)
    return a
}

/**
 * 是否内嵌
 */
func (a *InputSignature) Embed(value interface{}) *InputSignature {
    a.Set("embed", value)
    return a
}
