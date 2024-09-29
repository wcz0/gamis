package renderers


/**
 * Text 文本输入框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/text

 * @author wcz0
 * @version 6.2.2
 */
type TextControl struct {
	*BaseRenderer
}

func NewTextControl() *TextControl {
    a := &TextControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-text")
    return a
}


func (a *TextControl) Set(name string, value interface{}) *TextControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 是否为多选模式
 */
func (a *TextControl) Multiple(value interface{}) *TextControl {
    a.Set("multiple", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *TextControl) LabelWidth(value interface{}) *TextControl {
    a.Set("labelWidth", value)
    return a
}

/**
 */
func (a *TextControl) StaticSchema(value interface{}) *TextControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *TextControl) EditorSetting(value interface{}) *TextControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *TextControl) LabelRemark(value interface{}) *TextControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 边框模式，全边框，还是半边框，或者没边框。
 * 可选值: full | half | none
 */
func (a *TextControl) BorderMode(value interface{}) *TextControl {
    a.Set("borderMode", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *TextControl) Width(value interface{}) *TextControl {
    a.Set("width", value)
    return a
}

/**
 * 是否可以编辑
 */
func (a *TextControl) Editable(value interface{}) *TextControl {
    a.Set("editable", value)
    return a
}

/**
 */
func (a *TextControl) AddOn(value interface{}) *TextControl {
    a.Set("addOn", value)
    return a
}

/**
 * 自动完成 API，当输入部分文字的时候，会将这些文字通过 ${term} 可以取到，发送给接口。 接口可以返回匹配到的选项，帮助用户输入。
 */
func (a *TextControl) AutoComplete(value interface{}) *TextControl {
    a.Set("autoComplete", value)
    return a
}

/**
 * 多选模式，值太多时是否避免折行
 */
func (a *TextControl) ValuesNoWrap(value interface{}) *TextControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * 选项删除 API
 */
func (a *TextControl) DeleteApi(value interface{}) *TextControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * 选项删除提示文字。
 */
func (a *TextControl) DeleteConfirmText(value interface{}) *TextControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * 配置 label className
 */
func (a *TextControl) LabelClassName(value interface{}) *TextControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *TextControl) ExtraName(value interface{}) *TextControl {
    a.Set("extraName", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *TextControl) Hint(value interface{}) *TextControl {
    a.Set("hint", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *TextControl) Hidden(value interface{}) *TextControl {
    a.Set("hidden", value)
    return a
}

/**
 * 描述标题
 */
func (a *TextControl) LabelAlign(value interface{}) *TextControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 是否禁用
 */
func (a *TextControl) Disabled(value interface{}) *TextControl {
    a.Set("disabled", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *TextControl) VisibleOn(value interface{}) *TextControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 原生input标签的CSS类名
 */
func (a *TextControl) NativeInputClassName(value interface{}) *TextControl {
    a.Set("nativeInputClassName", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *TextControl) Inline(value interface{}) *TextControl {
    a.Set("inline", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *TextControl) ClassName(value interface{}) *TextControl {
    a.Set("className", value)
    return a
}

/**
 * 在内容为空的时候清除值
 */
func (a *TextControl) ClearValueOnEmpty(value interface{}) *TextControl {
    a.Set("clearValueOnEmpty", value)
    return a
}

/**
 * 点清除按钮时，将表单项设置成当前配置的值。
 */
func (a *TextControl) ResetValue(value interface{}) *TextControl {
    a.Set("resetValue", value)
    return a
}

/**
 * 编辑时调用的 API
 */
func (a *TextControl) EditApi(value interface{}) *TextControl {
    a.Set("editApi", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *TextControl) DescriptionClassName(value interface{}) *TextControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *TextControl) Horizontal(value interface{}) *TextControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 是否显示
 */
func (a *TextControl) Visible(value interface{}) *TextControl {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *TextControl) StaticClassName(value interface{}) *TextControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 后缀
 */
func (a *TextControl) Suffix(value interface{}) *TextControl {
    a.Set("suffix", value)
    return a
}

/**
 * 是否可清除。
 */
func (a *TextControl) Clearable(value interface{}) *TextControl {
    a.Set("clearable", value)
    return a
}

/**
 * 懒加载字段
 */
func (a *TextControl) DeferField(value interface{}) *TextControl {
    a.Set("deferField", value)
    return a
}

/**
 * 添加时调用的接口
 */
func (a *TextControl) AddApi(value interface{}) *TextControl {
    a.Set("addApi", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *TextControl) Size(value interface{}) *TextControl {
    a.Set("size", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *TextControl) Name(value interface{}) *TextControl {
    a.Set("name", value)
    return a
}

/**
 */
func (a *TextControl) Validations(value interface{}) *TextControl {
    a.Set("validations", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *TextControl) AutoFill(value interface{}) *TextControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 可用来通过 API 拉取 options。
 */
func (a *TextControl) Source(value interface{}) *TextControl {
    a.Set("source", value)
    return a
}

/**
 * 只读条件
 */
func (a *TextControl) ReadOnlyOn(value interface{}) *TextControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 */
func (a *TextControl) InitAutoFill(value interface{}) *TextControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *TextControl) Static(value interface{}) *TextControl {
    a.Set("static", value)
    return a
}

/**
 * 是否去除首尾空白文本。
 */
func (a *TextControl) TrimContents(value interface{}) *TextControl {
    a.Set("trimContents", value)
    return a
}

/**
 * 限制文字最大输入个数
 */
func (a *TextControl) MaxLength(value interface{}) *TextControl {
    a.Set("maxLength", value)
    return a
}

/**
 * 控制新增弹框设置项
 */
func (a *TextControl) AddDialog(value interface{}) *TextControl {
    a.Set("addDialog", value)
    return a
}

/**
 * 占位符
 */
func (a *TextControl) Placeholder(value interface{}) *TextControl {
    a.Set("placeholder", value)
    return a
}

/**
 */
func (a *TextControl) Row(value interface{}) *TextControl {
    a.Set("row", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *TextControl) Id(value interface{}) *TextControl {
    a.Set("id", value)
    return a
}

/**
 * 表单项类型
 * 可选值: input-text | input-email | input-url | input-password | native-date | native-time | native-number
 */
func (a *TextControl) Type(value interface{}) *TextControl {
    a.Set("type", value)
    return a
}

/**
 * 前缀
 */
func (a *TextControl) Prefix(value interface{}) *TextControl {
    a.Set("prefix", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *TextControl) HiddenOn(value interface{}) *TextControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *TextControl) StaticPlaceholder(value interface{}) *TextControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 限制文字最小输入个数
 */
func (a *TextControl) MinLength(value interface{}) *TextControl {
    a.Set("minLength", value)
    return a
}

/**
 * 默认选择选项第一个值。
 */
func (a *TextControl) SelectFirst(value interface{}) *TextControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
 */
func (a *TextControl) ExtractValue(value interface{}) *TextControl {
    a.Set("extractValue", value)
    return a
}

/**
 * 是否可删除
 */
func (a *TextControl) Removable(value interface{}) *TextControl {
    a.Set("removable", value)
    return a
}

/**
 * 是否只读
 */
func (a *TextControl) ReadOnly(value interface{}) *TextControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *TextControl) ValidationErrors(value interface{}) *TextControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *TextControl) DisabledOn(value interface{}) *TextControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *TextControl) StaticInputClassName(value interface{}) *TextControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 配置原生 input 的 autoComplete 属性
 */
func (a *TextControl) NativeAutoComplete(value interface{}) *TextControl {
    a.Set("nativeAutoComplete", value)
    return a
}

/**
 * 是否可以新增
 */
func (a *TextControl) Creatable(value interface{}) *TextControl {
    a.Set("creatable", value)
    return a
}

/**
 * 新增文字
 */
func (a *TextControl) CreateBtnLabel(value interface{}) *TextControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 */
func (a *TextControl) Desc(value interface{}) *TextControl {
    a.Set("desc", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *TextControl) Mode(value interface{}) *TextControl {
    a.Set("mode", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *TextControl) OnEvent(value interface{}) *TextControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *TextControl) StaticOn(value interface{}) *TextControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 自动转换值
 */
func (a *TextControl) Transform(value interface{}) *TextControl {
    a.Set("transform", value)
    return a
}

/**
 * 配置 source 接口初始拉不拉取。
 */
func (a *TextControl) InitFetch(value interface{}) *TextControl {
    a.Set("initFetch", value)
    return a
}

/**
 * 选项修改的表单项
 */
func (a *TextControl) EditControls(value interface{}) *TextControl {
    a.Set("editControls", value)
    return a
}

/**
 * 控制编辑弹框设置项
 */
func (a *TextControl) EditDialog(value interface{}) *TextControl {
    a.Set("editDialog", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *TextControl) Description(value interface{}) *TextControl {
    a.Set("description", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *TextControl) Value(value interface{}) *TextControl {
    a.Set("value", value)
    return a
}

/**
 */
func (a *TextControl) TestIdBuilder(value interface{}) *TextControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *TextControl) Remark(value interface{}) *TextControl {
    a.Set("remark", value)
    return a
}

/**
 * 选项集合
 */
func (a *TextControl) Options(value interface{}) *TextControl {
    a.Set("options", value)
    return a
}

/**
 * 用表达式来配置 source 接口初始要不要拉取
 */
func (a *TextControl) InitFetchOn(value interface{}) *TextControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
 */
func (a *TextControl) DeferApi(value interface{}) *TextControl {
    a.Set("deferApi", value)
    return a
}

/**
 * 新增时的表单项。
 */
func (a *TextControl) AddControls(value interface{}) *TextControl {
    a.Set("addControls", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *TextControl) SubmitOnChange(value interface{}) *TextControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *TextControl) ValidateApi(value interface{}) *TextControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 组件样式
 */
func (a *TextControl) Style(value interface{}) *TextControl {
    a.Set("style", value)
    return a
}

/**
 * 分割符
 */
func (a *TextControl) Delimiter(value interface{}) *TextControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *TextControl) ValidateOnChange(value interface{}) *TextControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 配置 input className
 */
func (a *TextControl) InputClassName(value interface{}) *TextControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 是否显示计数
 */
func (a *TextControl) ShowCounter(value interface{}) *TextControl {
    a.Set("showCounter", value)
    return a
}

/**
 * control节点的CSS类名
 */
func (a *TextControl) InputControlClassName(value interface{}) *TextControl {
    a.Set("inputControlClassName", value)
    return a
}

/**
 * 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
 */
func (a *TextControl) JoinValues(value interface{}) *TextControl {
    a.Set("joinValues", value)
    return a
}

/**
 * 描述标题
 */
func (a *TextControl) Label(value interface{}) *TextControl {
    a.Set("label", value)
    return a
}

/**
 * 是否为必填
 */
func (a *TextControl) Required(value interface{}) *TextControl {
    a.Set("required", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *TextControl) ClearValueOnHidden(value interface{}) *TextControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *TextControl) StaticLabelClassName(value interface{}) *TextControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *TextControl) UseMobileUI(value interface{}) *TextControl {
    a.Set("useMobileUI", value)
    return a
}
