package renderers


/**
 * Rating 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/rating

 * @author wcz0
 * @version 6.2.2
 */
type RatingControl struct {
	*BaseRenderer
}

func NewRatingControl() *RatingControl {
    a := &RatingControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-rating")
    return a
}


func (a *RatingControl) Set(name string, value interface{}) *RatingControl {
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
func (a *RatingControl) DisabledOn(value interface{}) *RatingControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *RatingControl) StaticClassName(value interface{}) *RatingControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 占位符
 */
func (a *RatingControl) Placeholder(value interface{}) *RatingControl {
    a.Set("placeholder", value)
    return a
}

/**
 */
func (a *RatingControl) InitAutoFill(value interface{}) *RatingControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 自定义文字类名
 */
func (a *RatingControl) TextClassName(value interface{}) *RatingControl {
    a.Set("textClassName", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *RatingControl) HiddenOn(value interface{}) *RatingControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *RatingControl) VisibleOn(value interface{}) *RatingControl {
    a.Set("visibleOn", value)
    return a
}

/**
 */
func (a *RatingControl) StaticSchema(value interface{}) *RatingControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 表单项类型
 */
func (a *RatingControl) Type(value interface{}) *RatingControl {
    a.Set("type", value)
    return a
}

/**
 * 是否显示
 */
func (a *RatingControl) Visible(value interface{}) *RatingControl {
    a.Set("visible", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *RatingControl) EditorSetting(value interface{}) *RatingControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *RatingControl) DescriptionClassName(value interface{}) *RatingControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *RatingControl) Value(value interface{}) *RatingControl {
    a.Set("value", value)
    return a
}

/**
 * 描述标题
 */
func (a *RatingControl) Label(value interface{}) *RatingControl {
    a.Set("label", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *RatingControl) LabelRemark(value interface{}) *RatingControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *RatingControl) AutoFill(value interface{}) *RatingControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 是否只读
 */
func (a *RatingControl) Readonly(value interface{}) *RatingControl {
    a.Set("readonly", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *RatingControl) StaticInputClassName(value interface{}) *RatingControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *RatingControl) ValidateOnChange(value interface{}) *RatingControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 */
func (a *RatingControl) Validations(value interface{}) *RatingControl {
    a.Set("validations", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *RatingControl) OnEvent(value interface{}) *RatingControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *RatingControl) ClearValueOnHidden(value interface{}) *RatingControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *RatingControl) ValidateApi(value interface{}) *RatingControl {
    a.Set("validateApi", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *RatingControl) LabelWidth(value interface{}) *RatingControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *RatingControl) SubmitOnChange(value interface{}) *RatingControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *RatingControl) Mode(value interface{}) *RatingControl {
    a.Set("mode", value)
    return a
}

/**
 */
func (a *RatingControl) Row(value interface{}) *RatingControl {
    a.Set("row", value)
    return a
}

/**
 * 分数
 */
func (a *RatingControl) Count(value interface{}) *RatingControl {
    a.Set("count", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *RatingControl) StaticPlaceholder(value interface{}) *RatingControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *RatingControl) Name(value interface{}) *RatingControl {
    a.Set("name", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *RatingControl) ExtraName(value interface{}) *RatingControl {
    a.Set("extraName", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *RatingControl) ValidationErrors(value interface{}) *RatingControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 未被选中的星星的颜色
 */
func (a *RatingControl) InactiveColor(value interface{}) *RatingControl {
    a.Set("inactiveColor", value)
    return a
}

/**
 * 文字的位置
 */
func (a *RatingControl) TextPosition(value interface{}) *RatingControl {
    a.Set("textPosition", value)
    return a
}

/**
 * 是否禁用
 */
func (a *RatingControl) Disabled(value interface{}) *RatingControl {
    a.Set("disabled", value)
    return a
}

/**
 * 描述标题
 */
func (a *RatingControl) LabelAlign(value interface{}) *RatingControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *RatingControl) Hint(value interface{}) *RatingControl {
    a.Set("hint", value)
    return a
}

/**
 * 只读条件
 */
func (a *RatingControl) ReadOnlyOn(value interface{}) *RatingControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 是否为必填
 */
func (a *RatingControl) Required(value interface{}) *RatingControl {
    a.Set("required", value)
    return a
}

/**
 * 是否允许再次点击后清除
 */
func (a *RatingControl) AllowClear(value interface{}) *RatingControl {
    a.Set("allowClear", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *RatingControl) Static(value interface{}) *RatingControl {
    a.Set("static", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *RatingControl) Remark(value interface{}) *RatingControl {
    a.Set("remark", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *RatingControl) Inline(value interface{}) *RatingControl {
    a.Set("inline", value)
    return a
}

/**
 * 允许半颗星
 */
func (a *RatingControl) Half(value interface{}) *RatingControl {
    a.Set("half", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *RatingControl) Id(value interface{}) *RatingControl {
    a.Set("id", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *RatingControl) Horizontal(value interface{}) *RatingControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 自定义字符类名
 */
func (a *RatingControl) CharClassName(value interface{}) *RatingControl {
    a.Set("charClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *RatingControl) UseMobileUI(value interface{}) *RatingControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *RatingControl) Size(value interface{}) *RatingControl {
    a.Set("size", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *RatingControl) Description(value interface{}) *RatingControl {
    a.Set("description", value)
    return a
}

/**
 * 配置 input className
 */
func (a *RatingControl) InputClassName(value interface{}) *RatingControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *RatingControl) StaticLabelClassName(value interface{}) *RatingControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *RatingControl) TestIdBuilder(value interface{}) *RatingControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 自定义字符
 */
func (a *RatingControl) Char(value interface{}) *RatingControl {
    a.Set("char", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *RatingControl) Width(value interface{}) *RatingControl {
    a.Set("width", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *RatingControl) Hidden(value interface{}) *RatingControl {
    a.Set("hidden", value)
    return a
}

/**
 * 配置 label className
 */
func (a *RatingControl) LabelClassName(value interface{}) *RatingControl {
    a.Set("labelClassName", value)
    return a
}

/**
 */
func (a *RatingControl) Desc(value interface{}) *RatingControl {
    a.Set("desc", value)
    return a
}

/**
 * 星星被选中的颜色
 */
func (a *RatingControl) Colors(value interface{}) *RatingControl {
    a.Set("colors", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *RatingControl) ClassName(value interface{}) *RatingControl {
    a.Set("className", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *RatingControl) StaticOn(value interface{}) *RatingControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 组件样式
 */
func (a *RatingControl) Style(value interface{}) *RatingControl {
    a.Set("style", value)
    return a
}

/**
 * 是否只读
 */
func (a *RatingControl) ReadOnly(value interface{}) *RatingControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 星星被选中时的提示文字
 */
func (a *RatingControl) Texts(value interface{}) *RatingControl {
    a.Set("texts", value)
    return a
}
