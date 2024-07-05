package renderers


/**
 * 按钮组控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/button-group

 * @author wcz0
 * @version 6.2.2
 */
type ButtonGroupControl struct {
	*BaseRenderer
}

func NewButtonGroupControl() *ButtonGroupControl {
    a := &ButtonGroupControl{
        BaseRenderer: NewBaseRenderer(),
    }
    a.Set("type", "button-group-select")
    return a
}

/**
 * 按钮大小
 * 可选值: xs | sm | md | lg
 */
func (a *ButtonGroupControl) Size(value interface{}) *ButtonGroupControl {
    a.Set("size", value)
    return a
}

/**
 * 用表达式来配置 source 接口初始要不要拉取
 */
func (a *ButtonGroupControl) InitFetchOn(value interface{}) *ButtonGroupControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 编辑时调用的 API
 */
func (a *ButtonGroupControl) EditApi(value interface{}) *ButtonGroupControl {
    a.Set("editApi", value)
    return a
}

/**
 * 选项删除 API
 */
func (a *ButtonGroupControl) DeleteApi(value interface{}) *ButtonGroupControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * 描述标题
 */
func (a *ButtonGroupControl) Label(value interface{}) *ButtonGroupControl {
    a.Set("label", value)
    return a
}

/**
 * 通过 JS 表达式来配置当前表单项的禁用状态。
 */
func (a *ButtonGroupControl) DisabledOn(value interface{}) *ButtonGroupControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否可清除。
 */
func (a *ButtonGroupControl) Clearable(value interface{}) *ButtonGroupControl {
    a.Set("clearable", value)
    return a
}

/**
 * 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
 */
func (a *ButtonGroupControl) DeferApi(value interface{}) *ButtonGroupControl {
    a.Set("deferApi", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *ButtonGroupControl) ExtraName(value interface{}) *ButtonGroupControl {
    a.Set("extraName", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *ButtonGroupControl) Mode(value interface{}) *ButtonGroupControl {
    a.Set("mode", value)
    return a
}

/**
 */
func (a *ButtonGroupControl) Type(value interface{}) *ButtonGroupControl {
    a.Set("type", value)
    return a
}

/**
 * 默认选择选项第一个值。
 */
func (a *ButtonGroupControl) SelectFirst(value interface{}) *ButtonGroupControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *ButtonGroupControl) Inline(value interface{}) *ButtonGroupControl {
    a.Set("inline", value)
    return a
}

/**
 * 是否为禁用状态。
 */
func (a *ButtonGroupControl) Disabled(value interface{}) *ButtonGroupControl {
    a.Set("disabled", value)
    return a
}

/**
 * 通过 JS 表达式来配置当前表单项是否显示
 */
func (a *ButtonGroupControl) VisibleOn(value interface{}) *ButtonGroupControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *ButtonGroupControl) Static(value interface{}) *ButtonGroupControl {
    a.Set("static", value)
    return a
}

/**
 * 按钮选中的样式级别
 */
func (a *ButtonGroupControl) BtnActiveLevel(value interface{}) *ButtonGroupControl {
    a.Set("btnActiveLevel", value)
    return a
}

/**
 * 配置 source 接口初始拉不拉取。
 */
func (a *ButtonGroupControl) InitFetch(value interface{}) *ButtonGroupControl {
    a.Set("initFetch", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *ButtonGroupControl) Value(value interface{}) *ButtonGroupControl {
    a.Set("value", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *ButtonGroupControl) StaticLabelClassName(value interface{}) *ButtonGroupControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *ButtonGroupControl) BtnClassName(value interface{}) *ButtonGroupControl {
    a.Set("btnClassName", value)
    return a
}

/**
 * 新增时的表单项。
 */
func (a *ButtonGroupControl) AddControls(value interface{}) *ButtonGroupControl {
    a.Set("addControls", value)
    return a
}

/**
 * 是否为必填
 */
func (a *ButtonGroupControl) Required(value interface{}) *ButtonGroupControl {
    a.Set("required", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *ButtonGroupControl) StaticInputClassName(value interface{}) *ButtonGroupControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 平铺展示？
 */
func (a *ButtonGroupControl) Tiled(value interface{}) *ButtonGroupControl {
    a.Set("tiled", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *ButtonGroupControl) Width(value interface{}) *ButtonGroupControl {
    a.Set("width", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *ButtonGroupControl) AutoFill(value interface{}) *ButtonGroupControl {
    a.Set("autoFill", value)
    return a
}

/**
 */
func (a *ButtonGroupControl) BtnActiveClassName(value interface{}) *ButtonGroupControl {
    a.Set("btnActiveClassName", value)
    return a
}

/**
 * 是否可以新增
 */
func (a *ButtonGroupControl) Creatable(value interface{}) *ButtonGroupControl {
    a.Set("creatable", value)
    return a
}

/**
 * 是否可以编辑
 */
func (a *ButtonGroupControl) Editable(value interface{}) *ButtonGroupControl {
    a.Set("editable", value)
    return a
}

/**
 * 选项修改的表单项
 */
func (a *ButtonGroupControl) EditControls(value interface{}) *ButtonGroupControl {
    a.Set("editControls", value)
    return a
}

/**
 * 配置 label className
 */
func (a *ButtonGroupControl) LabelClassName(value interface{}) *ButtonGroupControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *ButtonGroupControl) Horizontal(value interface{}) *ButtonGroupControl {
    a.Set("horizontal", value)
    return a
}

/**
 */
func (a *ButtonGroupControl) Testid(value interface{}) *ButtonGroupControl {
    a.Set("testid", value)
    return a
}

/**
 * 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
 */
func (a *ButtonGroupControl) ExtractValue(value interface{}) *ButtonGroupControl {
    a.Set("extractValue", value)
    return a
}

/**
 * 描述标题
 */
func (a *ButtonGroupControl) LabelAlign(value interface{}) *ButtonGroupControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *ButtonGroupControl) SubmitOnChange(value interface{}) *ButtonGroupControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *ButtonGroupControl) ClassName(value interface{}) *ButtonGroupControl {
    a.Set("className", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *ButtonGroupControl) StaticClassName(value interface{}) *ButtonGroupControl {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *ButtonGroupControl) Validations(value interface{}) *ButtonGroupControl {
    a.Set("validations", value)
    return a
}

/**
 * 控制编辑弹框设置项
 */
func (a *ButtonGroupControl) EditDialog(value interface{}) *ButtonGroupControl {
    a.Set("editDialog", value)
    return a
}

/**
 * 按钮集合
 */
func (a *ButtonGroupControl) Buttons(value interface{}) *ButtonGroupControl {
    a.Set("buttons", value)
    return a
}

/**
 * 可用来通过 API 拉取 options。
 */
func (a *ButtonGroupControl) Source(value interface{}) *ButtonGroupControl {
    a.Set("source", value)
    return a
}

/**
 * 控制新增弹框设置项
 */
func (a *ButtonGroupControl) AddDialog(value interface{}) *ButtonGroupControl {
    a.Set("addDialog", value)
    return a
}

/**
 */
func (a *ButtonGroupControl) Row(value interface{}) *ButtonGroupControl {
    a.Set("row", value)
    return a
}

/**
 * 是否显示
 */
func (a *ButtonGroupControl) Visible(value interface{}) *ButtonGroupControl {
    a.Set("visible", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *ButtonGroupControl) StaticOn(value interface{}) *ButtonGroupControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *ButtonGroupControl) LabelRemark(value interface{}) *ButtonGroupControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 只读条件
 */
func (a *ButtonGroupControl) ReadOnlyOn(value interface{}) *ButtonGroupControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *ButtonGroupControl) ValidationErrors(value interface{}) *ButtonGroupControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 选项集合
 */
func (a *ButtonGroupControl) Options(value interface{}) *ButtonGroupControl {
    a.Set("options", value)
    return a
}

/**
 * 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
 */
func (a *ButtonGroupControl) JoinValues(value interface{}) *ButtonGroupControl {
    a.Set("joinValues", value)
    return a
}

/**
 * 多选模式，值太多时是否避免折行
 */
func (a *ButtonGroupControl) ValuesNoWrap(value interface{}) *ButtonGroupControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * 添加时调用的接口
 */
func (a *ButtonGroupControl) AddApi(value interface{}) *ButtonGroupControl {
    a.Set("addApi", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *ButtonGroupControl) Remark(value interface{}) *ButtonGroupControl {
    a.Set("remark", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *ButtonGroupControl) ClearValueOnHidden(value interface{}) *ButtonGroupControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 */
func (a *ButtonGroupControl) StaticSchema(value interface{}) *ButtonGroupControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 组件样式
 */
func (a *ButtonGroupControl) Style(value interface{}) *ButtonGroupControl {
    a.Set("style", value)
    return a
}

/**
 * 垂直展示？
 */
func (a *ButtonGroupControl) Vertical(value interface{}) *ButtonGroupControl {
    a.Set("vertical", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *ButtonGroupControl) HiddenOn(value interface{}) *ButtonGroupControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 */
func (a *ButtonGroupControl) TestIdBuilder(value interface{}) *ButtonGroupControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 点清除按钮时，将表单项设置成当前配置的值。
 */
func (a *ButtonGroupControl) ResetValue(value interface{}) *ButtonGroupControl {
    a.Set("resetValue", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *ButtonGroupControl) Description(value interface{}) *ButtonGroupControl {
    a.Set("description", value)
    return a
}

/**
 */
func (a *ButtonGroupControl) Desc(value interface{}) *ButtonGroupControl {
    a.Set("desc", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *ButtonGroupControl) DescriptionClassName(value interface{}) *ButtonGroupControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 */
func (a *ButtonGroupControl) InitAutoFill(value interface{}) *ButtonGroupControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 分割符
 */
func (a *ButtonGroupControl) Delimiter(value interface{}) *ButtonGroupControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *ButtonGroupControl) Name(value interface{}) *ButtonGroupControl {
    a.Set("name", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *ButtonGroupControl) Id(value interface{}) *ButtonGroupControl {
    a.Set("id", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *ButtonGroupControl) UseMobileUI(value interface{}) *ButtonGroupControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 懒加载字段
 */
func (a *ButtonGroupControl) DeferField(value interface{}) *ButtonGroupControl {
    a.Set("deferField", value)
    return a
}

/**
 * 新增文字
 */
func (a *ButtonGroupControl) CreateBtnLabel(value interface{}) *ButtonGroupControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *ButtonGroupControl) LabelWidth(value interface{}) *ButtonGroupControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *ButtonGroupControl) ValidateOnChange(value interface{}) *ButtonGroupControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *ButtonGroupControl) EditorSetting(value interface{}) *ButtonGroupControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 配置 input className
 */
func (a *ButtonGroupControl) InputClassName(value interface{}) *ButtonGroupControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *ButtonGroupControl) ValidateApi(value interface{}) *ButtonGroupControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *ButtonGroupControl) OnEvent(value interface{}) *ButtonGroupControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 占位符
 */
func (a *ButtonGroupControl) Placeholder(value interface{}) *ButtonGroupControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *ButtonGroupControl) Hidden(value interface{}) *ButtonGroupControl {
    a.Set("hidden", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *ButtonGroupControl) StaticPlaceholder(value interface{}) *ButtonGroupControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 是否为多选模式
 */
func (a *ButtonGroupControl) Multiple(value interface{}) *ButtonGroupControl {
    a.Set("multiple", value)
    return a
}

/**
 * 是否可删除
 */
func (a *ButtonGroupControl) Removable(value interface{}) *ButtonGroupControl {
    a.Set("removable", value)
    return a
}

/**
 * 选项删除提示文字。
 */
func (a *ButtonGroupControl) DeleteConfirmText(value interface{}) *ButtonGroupControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *ButtonGroupControl) Hint(value interface{}) *ButtonGroupControl {
    a.Set("hint", value)
    return a
}

/**
 * 是否只读
 */
func (a *ButtonGroupControl) ReadOnly(value interface{}) *ButtonGroupControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 按钮样式级别
 */
func (a *ButtonGroupControl) BtnLevel(value interface{}) *ButtonGroupControl {
    a.Set("btnLevel", value)
    return a
}
