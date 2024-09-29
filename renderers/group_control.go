package renderers


/**
 * Group 表单集合渲染器，能让多个表单在一行显示 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/group

 * @author wcz0
 * @version 6.2.2
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


func (a *GroupControl) Set(name string, value interface{}) *GroupControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 */
func (a *GroupControl) StaticSchema(value interface{}) *GroupControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *GroupControl) EditorSetting(value interface{}) *GroupControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *GroupControl) ValidateOnChange(value interface{}) *GroupControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 配置 input className
 */
func (a *GroupControl) InputClassName(value interface{}) *GroupControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 是否禁用
 */
func (a *GroupControl) Disabled(value interface{}) *GroupControl {
    a.Set("disabled", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *GroupControl) StaticOn(value interface{}) *GroupControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *GroupControl) StaticLabelClassName(value interface{}) *GroupControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 是否只读
 */
func (a *GroupControl) ReadOnly(value interface{}) *GroupControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *GroupControl) Description(value interface{}) *GroupControl {
    a.Set("description", value)
    return a
}

/**
 * 配置子表单项默认的展示方式。
 * 可选值: normal | inline | horizontal
 */
func (a *GroupControl) SubFormMode(value interface{}) *GroupControl {
    a.Set("subFormMode", value)
    return a
}

/**
 * 描述标题
 */
func (a *GroupControl) LabelAlign(value interface{}) *GroupControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *GroupControl) LabelWidth(value interface{}) *GroupControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *GroupControl) StaticInputClassName(value interface{}) *GroupControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 只读条件
 */
func (a *GroupControl) ReadOnlyOn(value interface{}) *GroupControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *GroupControl) HiddenOn(value interface{}) *GroupControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *GroupControl) SubmitOnChange(value interface{}) *GroupControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *GroupControl) AutoFill(value interface{}) *GroupControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 是否为必填
 */
func (a *GroupControl) Required(value interface{}) *GroupControl {
    a.Set("required", value)
    return a
}

/**
 */
func (a *GroupControl) TestIdBuilder(value interface{}) *GroupControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 配置 label className
 */
func (a *GroupControl) LabelClassName(value interface{}) *GroupControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *GroupControl) Remark(value interface{}) *GroupControl {
    a.Set("remark", value)
    return a
}

/**
 * 配置时垂直摆放还是左右摆放。
 * 可选值: horizontal | vertical
 */
func (a *GroupControl) Direction(value interface{}) *GroupControl {
    a.Set("direction", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *GroupControl) DisabledOn(value interface{}) *GroupControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *GroupControl) Id(value interface{}) *GroupControl {
    a.Set("id", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *GroupControl) Hint(value interface{}) *GroupControl {
    a.Set("hint", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *GroupControl) DescriptionClassName(value interface{}) *GroupControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 占位符
 */
func (a *GroupControl) Placeholder(value interface{}) *GroupControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *GroupControl) ValidationErrors(value interface{}) *GroupControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 如果是水平排版，这个属性可以细化水平排版的左右宽度占比。
 */
func (a *GroupControl) SubFormHorizontal(value interface{}) *GroupControl {
    a.Set("subFormHorizontal", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *GroupControl) Hidden(value interface{}) *GroupControl {
    a.Set("hidden", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *GroupControl) Name(value interface{}) *GroupControl {
    a.Set("name", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *GroupControl) Width(value interface{}) *GroupControl {
    a.Set("width", value)
    return a
}

/**
 * 表单项类型
 */
func (a *GroupControl) Type(value interface{}) *GroupControl {
    a.Set("type", value)
    return a
}

/**
 * 间隔
 * 可选值: xs | sm | normal
 */
func (a *GroupControl) Gap(value interface{}) *GroupControl {
    a.Set("gap", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *GroupControl) Mode(value interface{}) *GroupControl {
    a.Set("mode", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *GroupControl) Value(value interface{}) *GroupControl {
    a.Set("value", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *GroupControl) UseMobileUI(value interface{}) *GroupControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *GroupControl) LabelRemark(value interface{}) *GroupControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *GroupControl) StaticPlaceholder(value interface{}) *GroupControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *GroupControl) StaticClassName(value interface{}) *GroupControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *GroupControl) ValidateApi(value interface{}) *GroupControl {
    a.Set("validateApi", value)
    return a
}

/**
 * FormItem 集合
 */
func (a *GroupControl) Body(value interface{}) *GroupControl {
    a.Set("body", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *GroupControl) ClassName(value interface{}) *GroupControl {
    a.Set("className", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *GroupControl) VisibleOn(value interface{}) *GroupControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *GroupControl) Static(value interface{}) *GroupControl {
    a.Set("static", value)
    return a
}

/**
 * 是否显示
 */
func (a *GroupControl) Visible(value interface{}) *GroupControl {
    a.Set("visible", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *GroupControl) OnEvent(value interface{}) *GroupControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 组件样式
 */
func (a *GroupControl) Style(value interface{}) *GroupControl {
    a.Set("style", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *GroupControl) ClearValueOnHidden(value interface{}) *GroupControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 */
func (a *GroupControl) Row(value interface{}) *GroupControl {
    a.Set("row", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *GroupControl) ExtraName(value interface{}) *GroupControl {
    a.Set("extraName", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *GroupControl) Inline(value interface{}) *GroupControl {
    a.Set("inline", value)
    return a
}

/**
 */
func (a *GroupControl) Desc(value interface{}) *GroupControl {
    a.Set("desc", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *GroupControl) Horizontal(value interface{}) *GroupControl {
    a.Set("horizontal", value)
    return a
}

/**
 */
func (a *GroupControl) Validations(value interface{}) *GroupControl {
    a.Set("validations", value)
    return a
}

/**
 */
func (a *GroupControl) InitAutoFill(value interface{}) *GroupControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *GroupControl) Size(value interface{}) *GroupControl {
    a.Set("size", value)
    return a
}

/**
 * 描述标题
 */
func (a *GroupControl) Label(value interface{}) *GroupControl {
    a.Set("label", value)
    return a
}
