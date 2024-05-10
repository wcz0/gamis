package renderers


/**
 * City 城市选择框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/city

 * @author wcz0
 * @version 6.2.2
 */
type InputCityControl struct {
	*BaseRenderer
}

func NewInputCityControl() *InputCityControl {
    a := &InputCityControl{
        BaseRenderer: NewBaseRenderer(),
    }
    a.Set("type", "input-city")
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *InputCityControl) StaticLabelClassName(value interface{}) *InputCityControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *InputCityControl) VisibleOn(value interface{}) *InputCityControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 拼接的符号是啥？
 */
func (a *InputCityControl) Delimiter(value interface{}) *InputCityControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *InputCityControl) Value(value interface{}) *InputCityControl {
    a.Set("value", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *InputCityControl) AutoFill(value interface{}) *InputCityControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 指定为城市选择框。
 */
func (a *InputCityControl) Type(value interface{}) *InputCityControl {
    a.Set("type", value)
    return a
}

/**
 * 允许选择街道？
 */
func (a *InputCityControl) AllowStreet(value interface{}) *InputCityControl {
    a.Set("allowStreet", value)
    return a
}

/**
 * 是否显示搜索框
 */
func (a *InputCityControl) Searchable(value interface{}) *InputCityControl {
    a.Set("searchable", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *InputCityControl) StaticPlaceholder(value interface{}) *InputCityControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 占位符
 */
func (a *InputCityControl) Placeholder(value interface{}) *InputCityControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *InputCityControl) Static(value interface{}) *InputCityControl {
    a.Set("static", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *InputCityControl) LabelRemark(value interface{}) *InputCityControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 下拉框className
 */
func (a *InputCityControl) ItemClassName(value interface{}) *InputCityControl {
    a.Set("itemClassName", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *InputCityControl) OnEvent(value interface{}) *InputCityControl {
    a.Set("onEvent", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *InputCityControl) LabelWidth(value interface{}) *InputCityControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *InputCityControl) ValidateOnChange(value interface{}) *InputCityControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *InputCityControl) DisabledOn(value interface{}) *InputCityControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 描述标题
 */
func (a *InputCityControl) Label(value interface{}) *InputCityControl {
    a.Set("label", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *InputCityControl) ClearValueOnHidden(value interface{}) *InputCityControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 是否显示
 */
func (a *InputCityControl) Visible(value interface{}) *InputCityControl {
    a.Set("visible", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *InputCityControl) StaticOn(value interface{}) *InputCityControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *InputCityControl) Inline(value interface{}) *InputCityControl {
    a.Set("inline", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *InputCityControl) HiddenOn(value interface{}) *InputCityControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *InputCityControl) Id(value interface{}) *InputCityControl {
    a.Set("id", value)
    return a
}

/**
 * 配置 input className
 */
func (a *InputCityControl) InputClassName(value interface{}) *InputCityControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *InputCityControl) Hidden(value interface{}) *InputCityControl {
    a.Set("hidden", value)
    return a
}

/**
 */
func (a *InputCityControl) TestIdBuilder(value interface{}) *InputCityControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *InputCityControl) DescriptionClassName(value interface{}) *InputCityControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 允许选择地区？
 */
func (a *InputCityControl) AllowDistrict(value interface{}) *InputCityControl {
    a.Set("allowDistrict", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *InputCityControl) Width(value interface{}) *InputCityControl {
    a.Set("width", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *InputCityControl) ExtraName(value interface{}) *InputCityControl {
    a.Set("extraName", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *InputCityControl) Remark(value interface{}) *InputCityControl {
    a.Set("remark", value)
    return a
}

/**
 */
func (a *InputCityControl) StaticSchema(value interface{}) *InputCityControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *InputCityControl) SubmitOnChange(value interface{}) *InputCityControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 是否只读
 */
func (a *InputCityControl) ReadOnly(value interface{}) *InputCityControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *InputCityControl) Description(value interface{}) *InputCityControl {
    a.Set("description", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *InputCityControl) Mode(value interface{}) *InputCityControl {
    a.Set("mode", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *InputCityControl) Horizontal(value interface{}) *InputCityControl {
    a.Set("horizontal", value)
    return a
}

/**
 */
func (a *InputCityControl) Validations(value interface{}) *InputCityControl {
    a.Set("validations", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *InputCityControl) EditorSetting(value interface{}) *InputCityControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *InputCityControl) Name(value interface{}) *InputCityControl {
    a.Set("name", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *InputCityControl) UseMobileUI(value interface{}) *InputCityControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *InputCityControl) Size(value interface{}) *InputCityControl {
    a.Set("size", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *InputCityControl) ValidationErrors(value interface{}) *InputCityControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 是否将各个信息拼接成字符串。
 */
func (a *InputCityControl) JoinValues(value interface{}) *InputCityControl {
    a.Set("joinValues", value)
    return a
}

/**
 * 允许选择城市？
 */
func (a *InputCityControl) AllowCity(value interface{}) *InputCityControl {
    a.Set("allowCity", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *InputCityControl) StaticClassName(value interface{}) *InputCityControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *InputCityControl) Style(value interface{}) *InputCityControl {
    a.Set("style", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *InputCityControl) Hint(value interface{}) *InputCityControl {
    a.Set("hint", value)
    return a
}

/**
 */
func (a *InputCityControl) Desc(value interface{}) *InputCityControl {
    a.Set("desc", value)
    return a
}

/**
 * 是否为必填
 */
func (a *InputCityControl) Required(value interface{}) *InputCityControl {
    a.Set("required", value)
    return a
}

/**
 */
func (a *InputCityControl) LoadingConfig(value interface{}) *InputCityControl {
    a.Set("loadingConfig", value)
    return a
}

/**
 * 配置 label className
 */
func (a *InputCityControl) LabelClassName(value interface{}) *InputCityControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *InputCityControl) ValidateApi(value interface{}) *InputCityControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *InputCityControl) StaticInputClassName(value interface{}) *InputCityControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 只读条件
 */
func (a *InputCityControl) ReadOnlyOn(value interface{}) *InputCityControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 是否禁用
 */
func (a *InputCityControl) Disabled(value interface{}) *InputCityControl {
    a.Set("disabled", value)
    return a
}

/**
 * 描述标题
 */
func (a *InputCityControl) LabelAlign(value interface{}) *InputCityControl {
    a.Set("labelAlign", value)
    return a
}

/**
 */
func (a *InputCityControl) InitAutoFill(value interface{}) *InputCityControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 开启后只会存城市的 code 信息
 */
func (a *InputCityControl) ExtractValue(value interface{}) *InputCityControl {
    a.Set("extractValue", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *InputCityControl) ClassName(value interface{}) *InputCityControl {
    a.Set("className", value)
    return a
}
