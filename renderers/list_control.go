package renderers


/**
 * List 复选框 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/list
 *

*/
type ListControl struct {
	*BaseRenderer
}

func NewListControl() *ListControl {
    a := &ListControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "list-select")
    return a
}
/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *ListControl) ExtraName(value string) *ListControl {
    a.Set("extraName", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *ListControl) DescriptionClassName(value string) *ListControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *ListControl) VisibleOn(value string) *ListControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 描述标题
 */
func (a *ListControl) LabelAlign(value string) *ListControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *ListControl) LabelWidth(value string) *ListControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 新增文字
 */
func (a *ListControl) CreateBtnLabel(value string) *ListControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *ListControl) Description(value string) *ListControl {
    a.Set("description", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *ListControl) Static(value string) *ListControl {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *ListControl) StaticLabelClassName(value string) *ListControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *ListControl) StaticInputClassName(value string) *ListControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 图片div类名
 */
func (a *ListControl) ImageClassName(value string) *ListControl {
    a.Set("imageClassName", value)
    return a
}

/**
 * 默认选择选项第一个值。
 */
func (a *ListControl) SelectFirst(value string) *ListControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
 */
func (a *ListControl) DeferApi(value string) *ListControl {
    a.Set("deferApi", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *ListControl) ClearValueOnHidden(value string) *ListControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *ListControl) OnEvent(value string) *ListControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 用表达式来配置 source 接口初始要不要拉取
 */
func (a *ListControl) InitFetchOn(value string) *ListControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *ListControl) ValidationErrors(value string) *ListControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *ListControl) Hint(value string) *ListControl {
    a.Set("hint", value)
    return a
}

/**
 */
func (a *ListControl) StaticSchema(value string) *ListControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 多选模式，值太多时是否避免折行
 */
func (a *ListControl) ValuesNoWrap(value string) *ListControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *ListControl) Name(value string) *ListControl {
    a.Set("name", value)
    return a
}

/**
 * 点清除按钮时，将表单项设置成当前配置的值。
 */
func (a *ListControl) ResetValue(value string) *ListControl {
    a.Set("resetValue", value)
    return a
}

/**
 * 是否可以编辑
 */
func (a *ListControl) Editable(value string) *ListControl {
    a.Set("editable", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *ListControl) Size(value string) *ListControl {
    a.Set("size", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *ListControl) SubmitOnChange(value string) *ListControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *ListControl) ClassName(value string) *ListControl {
    a.Set("className", value)
    return a
}

/**
 * 选项集合
 */
func (a *ListControl) Options(value string) *ListControl {
    a.Set("options", value)
    return a
}

/**
 * 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
 */
func (a *ListControl) ExtractValue(value string) *ListControl {
    a.Set("extractValue", value)
    return a
}

/**
 * 激活态自定义展示模板。
 */
func (a *ListControl) ActiveItemSchema(value string) *ListControl {
    a.Set("activeItemSchema", value)
    return a
}

/**
 * 支持配置 list div 的 css 类名。 比如：flex justify-between
 */
func (a *ListControl) ListClassName(value string) *ListControl {
    a.Set("listClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *ListControl) EditorSetting(value string) *ListControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *ListControl) UseMobileUI(value string) *ListControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *ListControl) Mode(value string) *ListControl {
    a.Set("mode", value)
    return a
}

/**
 * 配置 source 接口初始拉不拉取。
 */
func (a *ListControl) InitFetch(value string) *ListControl {
    a.Set("initFetch", value)
    return a
}

/**
 * 配置 label className
 */
func (a *ListControl) LabelClassName(value string) *ListControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *ListControl) Inline(value string) *ListControl {
    a.Set("inline", value)
    return a
}

/**
 * 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
 */
func (a *ListControl) JoinValues(value string) *ListControl {
    a.Set("joinValues", value)
    return a
}

/**
 * 控制新增弹框设置项
 */
func (a *ListControl) AddDialog(value string) *ListControl {
    a.Set("addDialog", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *ListControl) ValidateOnChange(value string) *ListControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 是否为必填
 */
func (a *ListControl) Required(value string) *ListControl {
    a.Set("required", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *ListControl) HiddenOn(value string) *ListControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否可清除。
 */
func (a *ListControl) Clearable(value string) *ListControl {
    a.Set("clearable", value)
    return a
}

/**
 * 是否可以新增
 */
func (a *ListControl) Creatable(value string) *ListControl {
    a.Set("creatable", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *ListControl) DisabledOn(value string) *ListControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 编辑时调用的 API
 */
func (a *ListControl) EditApi(value string) *ListControl {
    a.Set("editApi", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *ListControl) AutoFill(value string) *ListControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 描述标题
 */
func (a *ListControl) Label(value string) *ListControl {
    a.Set("label", value)
    return a
}

/**
 */
func (a *ListControl) Desc(value string) *ListControl {
    a.Set("desc", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *ListControl) Remark(value string) *ListControl {
    a.Set("remark", value)
    return a
}

/**
 * 表单项类型
 */
func (a *ListControl) Type(value string) *ListControl {
    a.Set("type", value)
    return a
}

/**
 * 可以自定义展示模板。
 */
func (a *ListControl) ItemSchema(value string) *ListControl {
    a.Set("itemSchema", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *ListControl) Width(value string) *ListControl {
    a.Set("width", value)
    return a
}

/**
 * 添加时调用的接口
 */
func (a *ListControl) AddApi(value string) *ListControl {
    a.Set("addApi", value)
    return a
}

/**
 * 选项修改的表单项
 */
func (a *ListControl) EditControls(value string) *ListControl {
    a.Set("editControls", value)
    return a
}

/**
 * 是否禁用
 */
func (a *ListControl) Disabled(value string) *ListControl {
    a.Set("disabled", value)
    return a
}

/**
 * 是否可删除
 */
func (a *ListControl) Removable(value string) *ListControl {
    a.Set("removable", value)
    return a
}

/**
 */
func (a *ListControl) InitAutoFill(value string) *ListControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 只读条件
 */
func (a *ListControl) ReadOnlyOn(value string) *ListControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *ListControl) ValidateApi(value string) *ListControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 懒加载字段
 */
func (a *ListControl) DeferField(value string) *ListControl {
    a.Set("deferField", value)
    return a
}

/**
 * 新增时的表单项。
 */
func (a *ListControl) AddControls(value string) *ListControl {
    a.Set("addControls", value)
    return a
}

/**
 * 是否显示
 */
func (a *ListControl) Visible(value string) *ListControl {
    a.Set("visible", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *ListControl) Id(value string) *ListControl {
    a.Set("id", value)
    return a
}

/**
 * 组件样式
 */
func (a *ListControl) Style(value string) *ListControl {
    a.Set("style", value)
    return a
}

/**
 * 开启双击点选并提交。
 */
func (a *ListControl) SubmitOnDBClick(value string) *ListControl {
    a.Set("submitOnDBClick", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *ListControl) Horizontal(value string) *ListControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *ListControl) Value(value string) *ListControl {
    a.Set("value", value)
    return a
}

/**
 * 控制编辑弹框设置项
 */
func (a *ListControl) EditDialog(value string) *ListControl {
    a.Set("editDialog", value)
    return a
}

/**
 */
func (a *ListControl) Validations(value string) *ListControl {
    a.Set("validations", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *ListControl) Hidden(value string) *ListControl {
    a.Set("hidden", value)
    return a
}

/**
 * 是否为多选模式
 */
func (a *ListControl) Multiple(value string) *ListControl {
    a.Set("multiple", value)
    return a
}

/**
 * 分割符
 */
func (a *ListControl) Delimiter(value string) *ListControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 占位符
 */
func (a *ListControl) Placeholder(value string) *ListControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *ListControl) StaticOn(value string) *ListControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 选项删除提示文字。
 */
func (a *ListControl) DeleteConfirmText(value string) *ListControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * 是否只读
 */
func (a *ListControl) ReadOnly(value string) *ListControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 配置 input className
 */
func (a *ListControl) InputClassName(value string) *ListControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *ListControl) StaticPlaceholder(value string) *ListControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *ListControl) StaticClassName(value string) *ListControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *ListControl) LabelRemark(value string) *ListControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 可用来通过 API 拉取 options。
 */
func (a *ListControl) Source(value string) *ListControl {
    a.Set("source", value)
    return a
}

/**
 * 选项删除 API
 */
func (a *ListControl) DeleteApi(value string) *ListControl {
    a.Set("deleteApi", value)
    return a
}