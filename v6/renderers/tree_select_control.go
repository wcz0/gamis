package renderers


/**
 * Tree 下拉选择框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/tree

 * @author wcz0
 * @version 6.2.2
 */
type TreeSelectControl struct {
	*BaseRenderer
}

func NewTreeSelectControl() *TreeSelectControl {
    a := &TreeSelectControl{
        BaseRenderer: NewBaseRenderer(),
    }
    a.Set("type", "tree-select")
    return a
}

/**
 * 是否为多选模式
 */
func (a *TreeSelectControl) Multiple(value interface{}) *TreeSelectControl {
    a.Set("multiple", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *TreeSelectControl) ValidateOnChange(value interface{}) *TreeSelectControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *TreeSelectControl) StaticInputClassName(value interface{}) *TreeSelectControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 父子之间是否完全独立。
 */
func (a *TreeSelectControl) Cascade(value interface{}) *TreeSelectControl {
    a.Set("cascade", value)
    return a
}

/**
 * 选父级的时候，是否只把子节点的值包含在内
 */
func (a *TreeSelectControl) OnlyChildren(value interface{}) *TreeSelectControl {
    a.Set("onlyChildren", value)
    return a
}

/**
 * 是否自动选中子节点
 */
func (a *TreeSelectControl) AutoCheckChildren(value interface{}) *TreeSelectControl {
    a.Set("autoCheckChildren", value)
    return a
}

/**
 * 是否可清除。
 */
func (a *TreeSelectControl) Clearable(value interface{}) *TreeSelectControl {
    a.Set("clearable", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *TreeSelectControl) VisibleOn(value interface{}) *TreeSelectControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 显示图标
 */
func (a *TreeSelectControl) ShowIcon(value interface{}) *TreeSelectControl {
    a.Set("showIcon", value)
    return a
}

/**
 * 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
 */
func (a *TreeSelectControl) ExtractValue(value interface{}) *TreeSelectControl {
    a.Set("extractValue", value)
    return a
}

/**
 * 懒加载字段
 */
func (a *TreeSelectControl) DeferField(value interface{}) *TreeSelectControl {
    a.Set("deferField", value)
    return a
}

/**
 * 新增时的表单项。
 */
func (a *TreeSelectControl) AddControls(value interface{}) *TreeSelectControl {
    a.Set("addControls", value)
    return a
}

/**
 */
func (a *TreeSelectControl) Validations(value interface{}) *TreeSelectControl {
    a.Set("validations", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *TreeSelectControl) LabelRemark(value interface{}) *TreeSelectControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 选父级的时候是否把子节点的值也包含在内。
 */
func (a *TreeSelectControl) WithChildren(value interface{}) *TreeSelectControl {
    a.Set("withChildren", value)
    return a
}

/**
 * 自定义选项
 */
func (a *TreeSelectControl) MenuTpl(value interface{}) *TreeSelectControl {
    a.Set("menuTpl", value)
    return a
}

/**
 * 默认选择选项第一个值。
 */
func (a *TreeSelectControl) SelectFirst(value interface{}) *TreeSelectControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * 选项删除 API
 */
func (a *TreeSelectControl) DeleteApi(value interface{}) *TreeSelectControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * 配置 label className
 */
func (a *TreeSelectControl) LabelClassName(value interface{}) *TreeSelectControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 只读条件
 */
func (a *TreeSelectControl) ReadOnlyOn(value interface{}) *TreeSelectControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *TreeSelectControl) UseMobileUI(value interface{}) *TreeSelectControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 是否显示展开线
 */
func (a *TreeSelectControl) ShowOutline(value interface{}) *TreeSelectControl {
    a.Set("showOutline", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *TreeSelectControl) Width(value interface{}) *TreeSelectControl {
    a.Set("width", value)
    return a
}

/**
 * 用表达式来配置 source 接口初始要不要拉取
 */
func (a *TreeSelectControl) InitFetchOn(value interface{}) *TreeSelectControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 配置 source 接口初始拉不拉取。
 */
func (a *TreeSelectControl) InitFetch(value interface{}) *TreeSelectControl {
    a.Set("initFetch", value)
    return a
}

/**
 * 控制新增弹框设置项
 */
func (a *TreeSelectControl) AddDialog(value interface{}) *TreeSelectControl {
    a.Set("addDialog", value)
    return a
}

/**
 * 新增文字
 */
func (a *TreeSelectControl) CreateBtnLabel(value interface{}) *TreeSelectControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * 是否可以编辑
 */
func (a *TreeSelectControl) Editable(value interface{}) *TreeSelectControl {
    a.Set("editable", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *TreeSelectControl) ExtraName(value interface{}) *TreeSelectControl {
    a.Set("extraName", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *TreeSelectControl) SubmitOnChange(value interface{}) *TreeSelectControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *TreeSelectControl) Description(value interface{}) *TreeSelectControl {
    a.Set("description", value)
    return a
}

/**
 */
func (a *TreeSelectControl) StaticSchema(value interface{}) *TreeSelectControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *TreeSelectControl) EditorSetting(value interface{}) *TreeSelectControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 多选模式，值太多时是否避免折行
 */
func (a *TreeSelectControl) ValuesNoWrap(value interface{}) *TreeSelectControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * 选项修改的表单项
 */
func (a *TreeSelectControl) EditControls(value interface{}) *TreeSelectControl {
    a.Set("editControls", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *TreeSelectControl) Mode(value interface{}) *TreeSelectControl {
    a.Set("mode", value)
    return a
}

/**
 */
func (a *TreeSelectControl) InitAutoFill(value interface{}) *TreeSelectControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *TreeSelectControl) Hidden(value interface{}) *TreeSelectControl {
    a.Set("hidden", value)
    return a
}

/**
 * 是否隐藏选择框中已选中节点的祖先节点的文本信息
 */
func (a *TreeSelectControl) HideNodePathLabel(value interface{}) *TreeSelectControl {
    a.Set("hideNodePathLabel", value)
    return a
}

/**
 * 设置label字段
 */
func (a *TreeSelectControl) LabelField(value interface{}) *TreeSelectControl {
    a.Set("labelField", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *TreeSelectControl) Hint(value interface{}) *TreeSelectControl {
    a.Set("hint", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *TreeSelectControl) Horizontal(value interface{}) *TreeSelectControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *TreeSelectControl) Inline(value interface{}) *TreeSelectControl {
    a.Set("inline", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *TreeSelectControl) ClearValueOnHidden(value interface{}) *TreeSelectControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *TreeSelectControl) Static(value interface{}) *TreeSelectControl {
    a.Set("static", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *TreeSelectControl) StaticOn(value interface{}) *TreeSelectControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 收纳标签的Popover配置
 */
func (a *TreeSelectControl) OverflowTagPopover(value interface{}) *TreeSelectControl {
    a.Set("overflowTagPopover", value)
    return a
}

/**
 * 设置value字段
 */
func (a *TreeSelectControl) ValueField(value interface{}) *TreeSelectControl {
    a.Set("valueField", value)
    return a
}

/**
 * 选项集合
 */
func (a *TreeSelectControl) Options(value interface{}) *TreeSelectControl {
    a.Set("options", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *TreeSelectControl) DescriptionClassName(value interface{}) *TreeSelectControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *TreeSelectControl) StaticLabelClassName(value interface{}) *TreeSelectControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *TreeSelectControl) Remark(value interface{}) *TreeSelectControl {
    a.Set("remark", value)
    return a
}

/**
 * 顶级选项的值
 */
func (a *TreeSelectControl) RootValue(value interface{}) *TreeSelectControl {
    a.Set("rootValue", value)
    return a
}

/**
 * 是否为选项添加默认的Icon，默认值为true
 */
func (a *TreeSelectControl) EnableDefaultIcon(value interface{}) *TreeSelectControl {
    a.Set("enableDefaultIcon", value)
    return a
}

/**
 * 点清除按钮时，将表单项设置成当前配置的值。
 */
func (a *TreeSelectControl) ResetValue(value interface{}) *TreeSelectControl {
    a.Set("resetValue", value)
    return a
}

/**
 * 懒加载接口
 */
func (a *TreeSelectControl) DeferApi(value interface{}) *TreeSelectControl {
    a.Set("deferApi", value)
    return a
}

/**
 * 是否可以新增
 */
func (a *TreeSelectControl) Creatable(value interface{}) *TreeSelectControl {
    a.Set("creatable", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *TreeSelectControl) Value(value interface{}) *TreeSelectControl {
    a.Set("value", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *TreeSelectControl) AutoFill(value interface{}) *TreeSelectControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 是否可搜索
 */
func (a *TreeSelectControl) Searchable(value interface{}) *TreeSelectControl {
    a.Set("searchable", value)
    return a
}

/**
 */
func (a *TreeSelectControl) Desc(value interface{}) *TreeSelectControl {
    a.Set("desc", value)
    return a
}

/**
 * 顶级选项的名称
 */
func (a *TreeSelectControl) RootLabel(value interface{}) *TreeSelectControl {
    a.Set("rootLabel", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *TreeSelectControl) HiddenOn(value interface{}) *TreeSelectControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *TreeSelectControl) Id(value interface{}) *TreeSelectControl {
    a.Set("id", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *TreeSelectControl) OnEvent(value interface{}) *TreeSelectControl {
    a.Set("onEvent", value)
    return a
}

/**
 */
func (a *TreeSelectControl) TestIdBuilder(value interface{}) *TreeSelectControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 顶级节点是否可以创建子节点
 */
func (a *TreeSelectControl) RootCreatable(value interface{}) *TreeSelectControl {
    a.Set("rootCreatable", value)
    return a
}

/**
 * 编辑时调用的 API
 */
func (a *TreeSelectControl) EditApi(value interface{}) *TreeSelectControl {
    a.Set("editApi", value)
    return a
}

/**
 * 是否可删除
 */
func (a *TreeSelectControl) Removable(value interface{}) *TreeSelectControl {
    a.Set("removable", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *TreeSelectControl) LabelWidth(value interface{}) *TreeSelectControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *TreeSelectControl) Name(value interface{}) *TreeSelectControl {
    a.Set("name", value)
    return a
}

/**
 * 是否显示
 */
func (a *TreeSelectControl) Visible(value interface{}) *TreeSelectControl {
    a.Set("visible", value)
    return a
}

/**
 * 开启节点路径模式后，节点路径的分隔符
 */
func (a *TreeSelectControl) PathSeparator(value interface{}) *TreeSelectControl {
    a.Set("pathSeparator", value)
    return a
}

/**
 * 标签的最大展示数量，超出数量后以收纳浮层的方式展示，仅在多选模式开启后生效
 */
func (a *TreeSelectControl) MaxTagCount(value interface{}) *TreeSelectControl {
    a.Set("maxTagCount", value)
    return a
}

/**
 * 控制编辑弹框设置项
 */
func (a *TreeSelectControl) EditDialog(value interface{}) *TreeSelectControl {
    a.Set("editDialog", value)
    return a
}

/**
 * 描述标题
 */
func (a *TreeSelectControl) LabelAlign(value interface{}) *TreeSelectControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 占位符
 */
func (a *TreeSelectControl) Placeholder(value interface{}) *TreeSelectControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *TreeSelectControl) ValidationErrors(value interface{}) *TreeSelectControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *TreeSelectControl) StaticPlaceholder(value interface{}) *TreeSelectControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 单选时，只运行选择叶子节点
 */
func (a *TreeSelectControl) OnlyLeaf(value interface{}) *TreeSelectControl {
    a.Set("onlyLeaf", value)
    return a
}

/**
 * 可用来通过 API 拉取 options。
 */
func (a *TreeSelectControl) Source(value interface{}) *TreeSelectControl {
    a.Set("source", value)
    return a
}

/**
 * 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
 */
func (a *TreeSelectControl) JoinValues(value interface{}) *TreeSelectControl {
    a.Set("joinValues", value)
    return a
}

/**
 * 配置 input className
 */
func (a *TreeSelectControl) InputClassName(value interface{}) *TreeSelectControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *TreeSelectControl) ClassName(value interface{}) *TreeSelectControl {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *TreeSelectControl) DisabledOn(value interface{}) *TreeSelectControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 组件样式
 */
func (a *TreeSelectControl) Style(value interface{}) *TreeSelectControl {
    a.Set("style", value)
    return a
}

/**
 * 选项删除提示文字。
 */
func (a *TreeSelectControl) DeleteConfirmText(value interface{}) *TreeSelectControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * 是否只读
 */
func (a *TreeSelectControl) ReadOnly(value interface{}) *TreeSelectControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 是否为必填
 */
func (a *TreeSelectControl) Required(value interface{}) *TreeSelectControl {
    a.Set("required", value)
    return a
}

/**
 * 是否禁用
 */
func (a *TreeSelectControl) Disabled(value interface{}) *TreeSelectControl {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏顶级
 */
func (a *TreeSelectControl) HideRoot(value interface{}) *TreeSelectControl {
    a.Set("hideRoot", value)
    return a
}

/**
 * 分割符
 */
func (a *TreeSelectControl) Delimiter(value interface{}) *TreeSelectControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 添加时调用的接口
 */
func (a *TreeSelectControl) AddApi(value interface{}) *TreeSelectControl {
    a.Set("addApi", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *TreeSelectControl) Size(value interface{}) *TreeSelectControl {
    a.Set("size", value)
    return a
}

/**
 * 描述标题
 */
func (a *TreeSelectControl) Label(value interface{}) *TreeSelectControl {
    a.Set("label", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *TreeSelectControl) ValidateApi(value interface{}) *TreeSelectControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *TreeSelectControl) StaticClassName(value interface{}) *TreeSelectControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 表单项类型
 */
func (a *TreeSelectControl) Type(value interface{}) *TreeSelectControl {
    a.Set("type", value)
    return a
}

/**
 * 是否开启节点路径模式
 */
func (a *TreeSelectControl) EnableNodePath(value interface{}) *TreeSelectControl {
    a.Set("enableNodePath", value)
    return a
}
