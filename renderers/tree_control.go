package renderers


/**
 * Tree 下拉选择框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/tree
 *

*/
type TreeControl struct {
	*BaseRenderer
}

func NewTreeControl() *TreeControl {
    a := &TreeControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-tree")
    return a
}
/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *TreeControl) Size(value string) *TreeControl {
    a.Set("size", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *TreeControl) ExtraName(value string) *TreeControl {
    a.Set("extraName", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *TreeControl) Inline(value string) *TreeControl {
    a.Set("inline", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *TreeControl) Value(value string) *TreeControl {
    a.Set("value", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *TreeControl) Id(value string) *TreeControl {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *TreeControl) StaticClassName(value string) *TreeControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 默认选择选项第一个值。
 */
func (a *TreeControl) SelectFirst(value string) *TreeControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * 是否可删除
 */
func (a *TreeControl) Removable(value string) *TreeControl {
    a.Set("removable", value)
    return a
}

/**
 * 高度自动增长？
 */
func (a *TreeControl) HeightAuto(value string) *TreeControl {
    a.Set("heightAuto", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *TreeControl) UseMobileUI(value string) *TreeControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *TreeControl) Remark(value string) *TreeControl {
    a.Set("remark", value)
    return a
}

/**
 */
func (a *TreeControl) Desc(value string) *TreeControl {
    a.Set("desc", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *TreeControl) Mode(value string) *TreeControl {
    a.Set("mode", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *TreeControl) DisabledOn(value string) *TreeControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *TreeControl) StaticLabelClassName(value string) *TreeControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 选项修改的表单项
 */
func (a *TreeControl) EditControls(value string) *TreeControl {
    a.Set("editControls", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *TreeControl) ValidateOnChange(value string) *TreeControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 描述标题
 */
func (a *TreeControl) LabelAlign(value string) *TreeControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 是否禁用
 */
func (a *TreeControl) Disabled(value string) *TreeControl {
    a.Set("disabled", value)
    return a
}

/**
 * 是否显示
 */
func (a *TreeControl) Visible(value string) *TreeControl {
    a.Set("visible", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *TreeControl) EditorSetting(value string) *TreeControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 顶级选项的名称
 */
func (a *TreeControl) RootLabel(value string) *TreeControl {
    a.Set("rootLabel", value)
    return a
}

/**
 * 需要高亮的字符串
 */
func (a *TreeControl) HighlightTxt(value string) *TreeControl {
    a.Set("highlightTxt", value)
    return a
}

/**
 * 是否可清除。
 */
func (a *TreeControl) Clearable(value string) *TreeControl {
    a.Set("clearable", value)
    return a
}

/**
 * 选项删除提示文字。
 */
func (a *TreeControl) DeleteConfirmText(value string) *TreeControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *TreeControl) LabelWidth(value string) *TreeControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 只读条件
 */
func (a *TreeControl) ReadOnlyOn(value string) *TreeControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *TreeControl) Horizontal(value string) *TreeControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 配置 input className
 */
func (a *TreeControl) InputClassName(value string) *TreeControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *TreeControl) ClassName(value string) *TreeControl {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *TreeControl) Hidden(value string) *TreeControl {
    a.Set("hidden", value)
    return a
}

/**
 * 选项集合
 */
func (a *TreeControl) Options(value string) *TreeControl {
    a.Set("options", value)
    return a
}

/**
 * 可用来通过 API 拉取 options。
 */
func (a *TreeControl) Source(value string) *TreeControl {
    a.Set("source", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *TreeControl) LabelRemark(value string) *TreeControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 是否为选项添加默认的Icon，默认值为true
 */
func (a *TreeControl) EnableDefaultIcon(value string) *TreeControl {
    a.Set("enableDefaultIcon", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *TreeControl) HiddenOn(value string) *TreeControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *TreeControl) StaticInputClassName(value string) *TreeControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 点清除按钮时，将表单项设置成当前配置的值。
 */
func (a *TreeControl) ResetValue(value string) *TreeControl {
    a.Set("resetValue", value)
    return a
}

/**
 * 是否为必填
 */
func (a *TreeControl) Required(value string) *TreeControl {
    a.Set("required", value)
    return a
}

/**
 */
func (a *TreeControl) Validations(value string) *TreeControl {
    a.Set("validations", value)
    return a
}

/**
 * 选父级的时候，是否只把子节点的值包含在内
 */
func (a *TreeControl) OnlyChildren(value string) *TreeControl {
    a.Set("onlyChildren", value)
    return a
}

/**
 * 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
 */
func (a *TreeControl) JoinValues(value string) *TreeControl {
    a.Set("joinValues", value)
    return a
}

/**
 * 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
 */
func (a *TreeControl) ExtractValue(value string) *TreeControl {
    a.Set("extractValue", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *TreeControl) ValidationErrors(value string) *TreeControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *TreeControl) Static(value string) *TreeControl {
    a.Set("static", value)
    return a
}

/**
 * 显示图标
 */
func (a *TreeControl) ShowIcon(value string) *TreeControl {
    a.Set("showIcon", value)
    return a
}

/**
 * ui级联关系，true代表级联选中，false代表不级联，默认为true
 */
func (a *TreeControl) AutoCheckChildren(value string) *TreeControl {
    a.Set("autoCheckChildren", value)
    return a
}

/**
 * 顶级节点是否可以创建子节点
 */
func (a *TreeControl) RootCreatable(value string) *TreeControl {
    a.Set("rootCreatable", value)
    return a
}

/**
 * 是否显示展开线
 */
func (a *TreeControl) ShowOutline(value string) *TreeControl {
    a.Set("showOutline", value)
    return a
}

/**
 * 控制编辑弹框设置项
 */
func (a *TreeControl) EditDialog(value string) *TreeControl {
    a.Set("editDialog", value)
    return a
}

/**
 * 选项删除 API
 */
func (a *TreeControl) DeleteApi(value string) *TreeControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *TreeControl) Width(value string) *TreeControl {
    a.Set("width", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *TreeControl) OnEvent(value string) *TreeControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 顶级选项的值
 */
func (a *TreeControl) RootValue(value string) *TreeControl {
    a.Set("rootValue", value)
    return a
}

/**
 * 该属性代表数据级联关系，autoCheckChildren为true时生效，默认为false，具体数据级联关系如下： 1.casacde为false，ui行为为级联选中子节点，子节点禁用；值只包含父节点的值 2.cascade为false，withChildren为true，ui行为为级联选中子节点，子节点禁用；值包含父子节点的值 3.cascade为true，ui行为级联选中子节点，子节点可反选，值包含父子节点的值，此时withChildren属性失效 4.cascade不论为true还是false，onlyChildren为true，ui行为级联选中子节点，子节点可反选，值只包含子节点的值
 */
func (a *TreeControl) Cascade(value string) *TreeControl {
    a.Set("cascade", value)
    return a
}

/**
 * 分割符
 */
func (a *TreeControl) Delimiter(value string) *TreeControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *TreeControl) ValidateApi(value string) *TreeControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *TreeControl) StaticPlaceholder(value string) *TreeControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 新增文字
 */
func (a *TreeControl) CreateBtnLabel(value string) *TreeControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 */
func (a *TreeControl) InitAutoFill(value string) *TreeControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 是否可以编辑
 */
func (a *TreeControl) Editable(value string) *TreeControl {
    a.Set("editable", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *TreeControl) AutoFill(value string) *TreeControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 多选模式，值太多时是否避免折行
 */
func (a *TreeControl) ValuesNoWrap(value string) *TreeControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * 新增时的表单项。
 */
func (a *TreeControl) AddControls(value string) *TreeControl {
    a.Set("addControls", value)
    return a
}

/**
 * 配置 label className
 */
func (a *TreeControl) LabelClassName(value string) *TreeControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *TreeControl) VisibleOn(value string) *TreeControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 用表达式来配置 source 接口初始要不要拉取
 */
func (a *TreeControl) InitFetchOn(value string) *TreeControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 是否为多选模式
 */
func (a *TreeControl) Multiple(value string) *TreeControl {
    a.Set("multiple", value)
    return a
}

/**
 * 是否只读
 */
func (a *TreeControl) ReadOnly(value string) *TreeControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *TreeControl) Description(value string) *TreeControl {
    a.Set("description", value)
    return a
}

/**
 * 单选时，只运行选择叶子节点
 */
func (a *TreeControl) OnlyLeaf(value string) *TreeControl {
    a.Set("onlyLeaf", value)
    return a
}

/**
 * 编辑时调用的 API
 */
func (a *TreeControl) EditApi(value string) *TreeControl {
    a.Set("editApi", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *TreeControl) SubmitOnChange(value string) *TreeControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 搜索 API
 */
func (a *TreeControl) SearchApi(value string) *TreeControl {
    a.Set("searchApi", value)
    return a
}

/**
 * 占位符
 */
func (a *TreeControl) Placeholder(value string) *TreeControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否开启节点路径模式
 */
func (a *TreeControl) EnableNodePath(value string) *TreeControl {
    a.Set("enableNodePath", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *TreeControl) Name(value string) *TreeControl {
    a.Set("name", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *TreeControl) DescriptionClassName(value string) *TreeControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *TreeControl) ClearValueOnHidden(value string) *TreeControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *TreeControl) StaticOn(value string) *TreeControl {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *TreeControl) StaticSchema(value string) *TreeControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 表单项类型
 */
func (a *TreeControl) Type(value string) *TreeControl {
    a.Set("type", value)
    return a
}

/**
 * 懒加载接口
 */
func (a *TreeControl) DeferApi(value string) *TreeControl {
    a.Set("deferApi", value)
    return a
}

/**
 * 描述标题
 */
func (a *TreeControl) Label(value string) *TreeControl {
    a.Set("label", value)
    return a
}

/**
 * 开启节点路径模式后，节点路径的分隔符
 */
func (a *TreeControl) PathSeparator(value string) *TreeControl {
    a.Set("pathSeparator", value)
    return a
}

/**
 * 是否开启搜索
 */
func (a *TreeControl) Searchable(value string) *TreeControl {
    a.Set("searchable", value)
    return a
}

/**
 * 控制新增弹框设置项
 */
func (a *TreeControl) AddDialog(value string) *TreeControl {
    a.Set("addDialog", value)
    return a
}

/**
 * 选父级的时候是否把子节点的值也包含在内。
 */
func (a *TreeControl) WithChildren(value string) *TreeControl {
    a.Set("withChildren", value)
    return a
}

/**
 * 搜索框的配置
 */
func (a *TreeControl) SearchConfig(value string) *TreeControl {
    a.Set("searchConfig", value)
    return a
}

/**
 * 懒加载字段
 */
func (a *TreeControl) DeferField(value string) *TreeControl {
    a.Set("deferField", value)
    return a
}

/**
 * 添加时调用的接口
 */
func (a *TreeControl) AddApi(value string) *TreeControl {
    a.Set("addApi", value)
    return a
}

/**
 * 组件样式
 */
func (a *TreeControl) Style(value string) *TreeControl {
    a.Set("style", value)
    return a
}

/**
 * 是否可以新增
 */
func (a *TreeControl) Creatable(value string) *TreeControl {
    a.Set("creatable", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *TreeControl) Hint(value string) *TreeControl {
    a.Set("hint", value)
    return a
}

/**
 * 配置 source 接口初始拉不拉取。
 */
func (a *TreeControl) InitFetch(value string) *TreeControl {
    a.Set("initFetch", value)
    return a
}

/**
 * 是否隐藏顶级
 */
func (a *TreeControl) HideRoot(value string) *TreeControl {
    a.Set("hideRoot", value)
    return a
}
