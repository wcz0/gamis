package renderers


/**
 * Select 下拉选择框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/select

 * @author wcz0
 * @version 6.2.2
 */
type SelectControl struct {
	*BaseRenderer
}

func NewSelectControl() *SelectControl {
    a := &SelectControl{
        BaseRenderer: NewBaseRenderer(),
    }
    a.Set("type", "select")
    return a
}

/**
 * 配置 source 接口初始拉不拉取。
 */
func (a *SelectControl) InitFetch(value interface{}) *SelectControl {
    a.Set("initFetch", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *SelectControl) StaticClassName(value interface{}) *SelectControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 勾选展示模式
 * 可选值: table | group | tree | chained | associated
 */
func (a *SelectControl) SelectMode(value interface{}) *SelectControl {
    a.Set("selectMode", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *SelectControl) HiddenOn(value interface{}) *SelectControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 */
func (a *SelectControl) TestIdBuilder(value interface{}) *SelectControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 选项删除提示文字。
 */
func (a *SelectControl) DeleteConfirmText(value interface{}) *SelectControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *SelectControl) ClassName(value interface{}) *SelectControl {
    a.Set("className", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *SelectControl) Remark(value interface{}) *SelectControl {
    a.Set("remark", value)
    return a
}

/**
 * 可以自定义菜单展示。
 */
func (a *SelectControl) MenuTpl(value interface{}) *SelectControl {
    a.Set("menuTpl", value)
    return a
}

/**
 * 可用来通过 API 拉取 options。
 */
func (a *SelectControl) Source(value interface{}) *SelectControl {
    a.Set("source", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *SelectControl) LabelWidth(value interface{}) *SelectControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *SelectControl) Horizontal(value interface{}) *SelectControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 表单项类型
 * 可选值: select | multi-select
 */
func (a *SelectControl) Type(value interface{}) *SelectControl {
    a.Set("type", value)
    return a
}

/**
 * 当 selectMode 为 associated 时用来定义右侧的选择模式
 * 可选值: table | list | tree | chained
 */
func (a *SelectControl) RightMode(value interface{}) *SelectControl {
    a.Set("rightMode", value)
    return a
}

/**
 * 是否为多选模式
 */
func (a *SelectControl) Multiple(value interface{}) *SelectControl {
    a.Set("multiple", value)
    return a
}

/**
 * 选项删除 API
 */
func (a *SelectControl) DeleteApi(value interface{}) *SelectControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *SelectControl) LabelRemark(value interface{}) *SelectControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *SelectControl) UseMobileUI(value interface{}) *SelectControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 选项集合
 */
func (a *SelectControl) Options(value interface{}) *SelectControl {
    a.Set("options", value)
    return a
}

/**
 * 是否可以编辑
 */
func (a *SelectControl) Editable(value interface{}) *SelectControl {
    a.Set("editable", value)
    return a
}

/**
 * 选项修改的表单项
 */
func (a *SelectControl) EditControls(value interface{}) *SelectControl {
    a.Set("editControls", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *SelectControl) Mode(value interface{}) *SelectControl {
    a.Set("mode", value)
    return a
}

/**
 */
func (a *SelectControl) Validations(value interface{}) *SelectControl {
    a.Set("validations", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *SelectControl) StaticOn(value interface{}) *SelectControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 是否可以新增
 */
func (a *SelectControl) Creatable(value interface{}) *SelectControl {
    a.Set("creatable", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *SelectControl) DisabledOn(value interface{}) *SelectControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *SelectControl) Static(value interface{}) *SelectControl {
    a.Set("static", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *SelectControl) Width(value interface{}) *SelectControl {
    a.Set("width", value)
    return a
}

/**
 * 控制新增弹框设置项
 */
func (a *SelectControl) AddDialog(value interface{}) *SelectControl {
    a.Set("addDialog", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *SelectControl) DescriptionClassName(value interface{}) *SelectControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 是否禁用
 */
func (a *SelectControl) Disabled(value interface{}) *SelectControl {
    a.Set("disabled", value)
    return a
}

/**
 * 设置label字段
 */
func (a *SelectControl) LabelField(value interface{}) *SelectControl {
    a.Set("labelField", value)
    return a
}

/**
 * 是否自动选中子节点
 */
func (a *SelectControl) AutoCheckChildren(value interface{}) *SelectControl {
    a.Set("autoCheckChildren", value)
    return a
}

/**
 * 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
 */
func (a *SelectControl) JoinValues(value interface{}) *SelectControl {
    a.Set("joinValues", value)
    return a
}

/**
 * 新增文字
 */
func (a *SelectControl) CreateBtnLabel(value interface{}) *SelectControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * 占位符
 */
func (a *SelectControl) Placeholder(value interface{}) *SelectControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *SelectControl) AutoFill(value interface{}) *SelectControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
 */
func (a *SelectControl) ExtractValue(value interface{}) *SelectControl {
    a.Set("extractValue", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *SelectControl) ValidateOnChange(value interface{}) *SelectControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *SelectControl) EditorSetting(value interface{}) *SelectControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 当 searchResultMode 为 table 时定义表格列信息。
 */
func (a *SelectControl) SearchResultColumns(value interface{}) *SelectControl {
    a.Set("searchResultColumns", value)
    return a
}

/**
 * 在选项数量达到多少时开启虚拟渲染
 */
func (a *SelectControl) VirtualThreshold(value interface{}) *SelectControl {
    a.Set("virtualThreshold", value)
    return a
}

/**
 * 编辑时调用的 API
 */
func (a *SelectControl) EditApi(value interface{}) *SelectControl {
    a.Set("editApi", value)
    return a
}

/**
 * 是否可删除
 */
func (a *SelectControl) Removable(value interface{}) *SelectControl {
    a.Set("removable", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *SelectControl) SubmitOnChange(value interface{}) *SelectControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 */
func (a *SelectControl) InitAutoFill(value interface{}) *SelectControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 控制编辑弹框设置项
 */
func (a *SelectControl) EditDialog(value interface{}) *SelectControl {
    a.Set("editDialog", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *SelectControl) Id(value interface{}) *SelectControl {
    a.Set("id", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *SelectControl) Description(value interface{}) *SelectControl {
    a.Set("description", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *SelectControl) Inline(value interface{}) *SelectControl {
    a.Set("inline", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *SelectControl) Size(value interface{}) *SelectControl {
    a.Set("size", value)
    return a
}

/**
 */
func (a *SelectControl) StaticSchema(value interface{}) *SelectControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 边框模式，全边框，还是半边框，或者没边框。
 * 可选值: full | half | none
 */
func (a *SelectControl) BorderMode(value interface{}) *SelectControl {
    a.Set("borderMode", value)
    return a
}

/**
 * 选项的自定义CSS类名
 */
func (a *SelectControl) OptionClassName(value interface{}) *SelectControl {
    a.Set("optionClassName", value)
    return a
}

/**
 * 用表达式来配置 source 接口初始要不要拉取
 */
func (a *SelectControl) InitFetchOn(value interface{}) *SelectControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 当在value值未匹配到当前options中的选项时，是否value值对应文本飘红显示
 */
func (a *SelectControl) ShowInvalidMatch(value interface{}) *SelectControl {
    a.Set("showInvalidMatch", value)
    return a
}

/**
 * 当 selectMode 为 table 时定义表格列信息。
 */
func (a *SelectControl) Columns(value interface{}) *SelectControl {
    a.Set("columns", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *SelectControl) ValidateApi(value interface{}) *SelectControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *SelectControl) VisibleOn(value interface{}) *SelectControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *SelectControl) StaticInputClassName(value interface{}) *SelectControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 多选模式，值太多时是否避免折行
 */
func (a *SelectControl) ValuesNoWrap(value interface{}) *SelectControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * 配置 label className
 */
func (a *SelectControl) LabelClassName(value interface{}) *SelectControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 标签的最大展示数量，超出数量后以收纳浮层的方式展示，仅在多选模式开启后生效
 */
func (a *SelectControl) MaxTagCount(value interface{}) *SelectControl {
    a.Set("maxTagCount", value)
    return a
}

/**
 * 收纳标签的Popover配置
 */
func (a *SelectControl) OverflowTagPopover(value interface{}) *SelectControl {
    a.Set("overflowTagPopover", value)
    return a
}

/**
 * 只读条件
 */
func (a *SelectControl) ReadOnlyOn(value interface{}) *SelectControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 */
func (a *SelectControl) LoadingConfig(value interface{}) *SelectControl {
    a.Set("loadingConfig", value)
    return a
}

/**
 * 默认选择选项第一个值。
 */
func (a *SelectControl) SelectFirst(value interface{}) *SelectControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * 当 selectMode 为 associated 时用来定义左侧的选项
 */
func (a *SelectControl) LeftOptions(value interface{}) *SelectControl {
    a.Set("leftOptions", value)
    return a
}

/**
 * 新增时的表单项。
 */
func (a *SelectControl) AddControls(value interface{}) *SelectControl {
    a.Set("addControls", value)
    return a
}

/**
 * 是否为必填
 */
func (a *SelectControl) Required(value interface{}) *SelectControl {
    a.Set("required", value)
    return a
}

/**
 * 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
 */
func (a *SelectControl) DeferApi(value interface{}) *SelectControl {
    a.Set("deferApi", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *SelectControl) Name(value interface{}) *SelectControl {
    a.Set("name", value)
    return a
}

/**
 * 可多选条件下，是否可全选
 */
func (a *SelectControl) CheckAll(value interface{}) *SelectControl {
    a.Set("checkAll", value)
    return a
}

/**
 * 可多选条件下，全选项文案，默认 ”全选“
 */
func (a *SelectControl) CheckAllLabel(value interface{}) *SelectControl {
    a.Set("checkAllLabel", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *SelectControl) ExtraName(value interface{}) *SelectControl {
    a.Set("extraName", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *SelectControl) Hidden(value interface{}) *SelectControl {
    a.Set("hidden", value)
    return a
}

/**
 * 是否可清除。
 */
func (a *SelectControl) Clearable(value interface{}) *SelectControl {
    a.Set("clearable", value)
    return a
}

/**
 * 添加时调用的接口
 */
func (a *SelectControl) AddApi(value interface{}) *SelectControl {
    a.Set("addApi", value)
    return a
}

/**
 * 描述标题
 */
func (a *SelectControl) LabelAlign(value interface{}) *SelectControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 搜索结果展示模式
 * 可选值: table | list | tree | chained
 */
func (a *SelectControl) SearchResultMode(value interface{}) *SelectControl {
    a.Set("searchResultMode", value)
    return a
}

/**
 * 组件样式
 */
func (a *SelectControl) Style(value interface{}) *SelectControl {
    a.Set("style", value)
    return a
}

/**
 * 可多选条件下，是否默认全选中所有值
 */
func (a *SelectControl) DefaultCheckAll(value interface{}) *SelectControl {
    a.Set("defaultCheckAll", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *SelectControl) OnEvent(value interface{}) *SelectControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 当 selectMode 为 associated 时用来定义左侧的选择模式
 * 可选值: tree | list
 */
func (a *SelectControl) LeftMode(value interface{}) *SelectControl {
    a.Set("leftMode", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *SelectControl) Value(value interface{}) *SelectControl {
    a.Set("value", value)
    return a
}

/**
 * 分割符
 */
func (a *SelectControl) Delimiter(value interface{}) *SelectControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *SelectControl) Hint(value interface{}) *SelectControl {
    a.Set("hint", value)
    return a
}

/**
 * 是否可搜索
 */
func (a *SelectControl) Searchable(value interface{}) *SelectControl {
    a.Set("searchable", value)
    return a
}

/**
 * 设置value字段
 */
func (a *SelectControl) ValueField(value interface{}) *SelectControl {
    a.Set("valueField", value)
    return a
}

/**
 * 单个选项的高度，主要用于虚拟渲染
 */
func (a *SelectControl) ItemHeight(value interface{}) *SelectControl {
    a.Set("itemHeight", value)
    return a
}

/**
 */
func (a *SelectControl) Desc(value interface{}) *SelectControl {
    a.Set("desc", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *SelectControl) ValidationErrors(value interface{}) *SelectControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 是否显示
 */
func (a *SelectControl) Visible(value interface{}) *SelectControl {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *SelectControl) StaticPlaceholder(value interface{}) *SelectControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 配置 input className
 */
func (a *SelectControl) InputClassName(value interface{}) *SelectControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *SelectControl) ClearValueOnHidden(value interface{}) *SelectControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 */
func (a *SelectControl) Row(value interface{}) *SelectControl {
    a.Set("row", value)
    return a
}

/**
 * 点清除按钮时，将表单项设置成当前配置的值。
 */
func (a *SelectControl) ResetValue(value interface{}) *SelectControl {
    a.Set("resetValue", value)
    return a
}

/**
 * 懒加载字段
 */
func (a *SelectControl) DeferField(value interface{}) *SelectControl {
    a.Set("deferField", value)
    return a
}

/**
 * 是否只读
 */
func (a *SelectControl) ReadOnly(value interface{}) *SelectControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *SelectControl) StaticLabelClassName(value interface{}) *SelectControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 搜索 API
 */
func (a *SelectControl) SearchApi(value interface{}) *SelectControl {
    a.Set("searchApi", value)
    return a
}

/**
 * 下拉框 Popover 设置
 */
func (a *SelectControl) Overlay(value interface{}) *SelectControl {
    a.Set("overlay", value)
    return a
}

/**
 * 描述标题
 */
func (a *SelectControl) Label(value interface{}) *SelectControl {
    a.Set("label", value)
    return a
}

/**
 * 自动完成 API，当输入部分文字的时候，会将这些文字通过 ${term} 可以取到，发送给接口。 接口可以返回匹配到的选项，帮助用户输入。
 */
func (a *SelectControl) AutoComplete(value interface{}) *SelectControl {
    a.Set("autoComplete", value)
    return a
}
