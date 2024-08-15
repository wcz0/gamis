package renderers


/**
 * Nested Select 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/nested-select

 * @author wcz0
 * @version 6.2.2
 */
type NestedSelectControl struct {
	*BaseRenderer
}

func NewNestedSelectControl() *NestedSelectControl {
    a := &NestedSelectControl{
        BaseRenderer: NewBaseRenderer(),
    }

func (a *NestedSelectControl) Set(name string, value interface{}) *NestedSelectControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}

    a.Set("type", "nested-select")
    return a
}

/**
 * 点清除按钮时，将表单项设置成当前配置的值。
 */
func (a *NestedSelectControl) Resetvalue(value interface{}) *NestedSelectControl {
    a.Set("resetValue", value)
    return a
}

/**
 * 是否只读
 */
func (a *NestedSelectControl) Readonly(value interface{}) *NestedSelectControl {
    a.Set("readOnly", value)
    return a
}

/**
 */
func (a *NestedSelectControl) Validations(value interface{}) *NestedSelectControl {
    a.Set("validations", value)
    return a
}

/**
 * 是否可以编辑
 */
func (a *NestedSelectControl) Editable(value interface{}) *NestedSelectControl {
    a.Set("editable", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *NestedSelectControl) Clearvalueonhidden(value interface{}) *NestedSelectControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 */
func (a *NestedSelectControl) Staticschema(value interface{}) *NestedSelectControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *NestedSelectControl) Staticon(value interface{}) *NestedSelectControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 标签的最大展示数量，超出数量后以收纳浮层的方式展示，仅在多选模式开启后生效
 */
func (a *NestedSelectControl) Maxtagcount(value interface{}) *NestedSelectControl {
    a.Set("maxTagCount", value)
    return a
}

/**
 * 可用来通过 API 拉取 options。
 */
func (a *NestedSelectControl) Source(value interface{}) *NestedSelectControl {
    a.Set("source", value)
    return a
}

/**
 * 用表达式来配置 source 接口初始要不要拉取
 */
func (a *NestedSelectControl) Initfetchon(value interface{}) *NestedSelectControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 是否可以新增
 */
func (a *NestedSelectControl) Creatable(value interface{}) *NestedSelectControl {
    a.Set("creatable", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *NestedSelectControl) Descriptionclassname(value interface{}) *NestedSelectControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *NestedSelectControl) Hiddenon(value interface{}) *NestedSelectControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *NestedSelectControl) Mode(value interface{}) *NestedSelectControl {
    a.Set("mode", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *NestedSelectControl) Validateapi(value interface{}) *NestedSelectControl {
    a.Set("validateApi", value)
    return a
}

/**
 */
func (a *NestedSelectControl) Testidbuilder(value interface{}) *NestedSelectControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 收纳标签的Popover配置
 */
func (a *NestedSelectControl) Overflowtagpopover(value interface{}) *NestedSelectControl {
    a.Set("overflowTagPopover", value)
    return a
}

/**
 * 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
 */
func (a *NestedSelectControl) Joinvalues(value interface{}) *NestedSelectControl {
    a.Set("joinValues", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *NestedSelectControl) Hint(value interface{}) *NestedSelectControl {
    a.Set("hint", value)
    return a
}

/**
 */
func (a *NestedSelectControl) Row(value interface{}) *NestedSelectControl {
    a.Set("row", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *NestedSelectControl) Staticlabelclassname(value interface{}) *NestedSelectControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 选父级的时候，是否只把子节点的值包含在内
 */
func (a *NestedSelectControl) Onlychildren(value interface{}) *NestedSelectControl {
    a.Set("onlyChildren", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *NestedSelectControl) Staticclassname(value interface{}) *NestedSelectControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *NestedSelectControl) Staticinputclassname(value interface{}) *NestedSelectControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 只允许选择叶子节点
 */
func (a *NestedSelectControl) Onlyleaf(value interface{}) *NestedSelectControl {
    a.Set("onlyLeaf", value)
    return a
}

/**
 * 选项集合
 */
func (a *NestedSelectControl) Options(value interface{}) *NestedSelectControl {
    a.Set("options", value)
    return a
}

/**
 * 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
 */
func (a *NestedSelectControl) Extractvalue(value interface{}) *NestedSelectControl {
    a.Set("extractValue", value)
    return a
}

/**
 * 是否可删除
 */
func (a *NestedSelectControl) Removable(value interface{}) *NestedSelectControl {
    a.Set("removable", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *NestedSelectControl) Validateonchange(value interface{}) *NestedSelectControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *NestedSelectControl) Classname(value interface{}) *NestedSelectControl {
    a.Set("className", value)
    return a
}

/**
 * 是否为多选模式
 */
func (a *NestedSelectControl) Multiple(value interface{}) *NestedSelectControl {
    a.Set("multiple", value)
    return a
}

/**
 * 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
 */
func (a *NestedSelectControl) Deferapi(value interface{}) *NestedSelectControl {
    a.Set("deferApi", value)
    return a
}

/**
 * 编辑时调用的 API
 */
func (a *NestedSelectControl) Editapi(value interface{}) *NestedSelectControl {
    a.Set("editApi", value)
    return a
}

/**
 * 是否显示
 */
func (a *NestedSelectControl) Visible(value interface{}) *NestedSelectControl {
    a.Set("visible", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *NestedSelectControl) Width(value interface{}) *NestedSelectControl {
    a.Set("width", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *NestedSelectControl) Labelremark(value interface{}) *NestedSelectControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 控制新增弹框设置项
 */
func (a *NestedSelectControl) Adddialog(value interface{}) *NestedSelectControl {
    a.Set("addDialog", value)
    return a
}

/**
 * 描述标题
 */
func (a *NestedSelectControl) Label(value interface{}) *NestedSelectControl {
    a.Set("label", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *NestedSelectControl) Extraname(value interface{}) *NestedSelectControl {
    a.Set("extraName", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *NestedSelectControl) Autofill(value interface{}) *NestedSelectControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *NestedSelectControl) Id(value interface{}) *NestedSelectControl {
    a.Set("id", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *NestedSelectControl) Inline(value interface{}) *NestedSelectControl {
    a.Set("inline", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *NestedSelectControl) Onevent(value interface{}) *NestedSelectControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *NestedSelectControl) Editorsetting(value interface{}) *NestedSelectControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 懒加载字段
 */
func (a *NestedSelectControl) Deferfield(value interface{}) *NestedSelectControl {
    a.Set("deferField", value)
    return a
}

/**
 * 选项删除提示文字。
 */
func (a *NestedSelectControl) Deleteconfirmtext(value interface{}) *NestedSelectControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * 组件样式
 */
func (a *NestedSelectControl) Style(value interface{}) *NestedSelectControl {
    a.Set("style", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *NestedSelectControl) Usemobileui(value interface{}) *NestedSelectControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 边框模式，全边框，还是半边框，或者没边框。
 * 可选值: full | half | none
 */
func (a *NestedSelectControl) Bordermode(value interface{}) *NestedSelectControl {
    a.Set("borderMode", value)
    return a
}

/**
 * 配置 input className
 */
func (a *NestedSelectControl) Inputclassname(value interface{}) *NestedSelectControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *NestedSelectControl) Value(value interface{}) *NestedSelectControl {
    a.Set("value", value)
    return a
}

/**
 * 是否禁用
 */
func (a *NestedSelectControl) Disabled(value interface{}) *NestedSelectControl {
    a.Set("disabled", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *NestedSelectControl) Visibleon(value interface{}) *NestedSelectControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *NestedSelectControl) Static(value interface{}) *NestedSelectControl {
    a.Set("static", value)
    return a
}

/**
 * 父子之间是否完全独立。
 */
func (a *NestedSelectControl) Cascade(value interface{}) *NestedSelectControl {
    a.Set("cascade", value)
    return a
}

/**
 * 选父级的时候是否把子节点的值也包含在内。
 */
func (a *NestedSelectControl) Withchildren(value interface{}) *NestedSelectControl {
    a.Set("withChildren", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *NestedSelectControl) Name(value interface{}) *NestedSelectControl {
    a.Set("name", value)
    return a
}

/**
 * 是否为必填
 */
func (a *NestedSelectControl) Required(value interface{}) *NestedSelectControl {
    a.Set("required", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *NestedSelectControl) Disabledon(value interface{}) *NestedSelectControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *NestedSelectControl) Staticplaceholder(value interface{}) *NestedSelectControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *NestedSelectControl) Hidden(value interface{}) *NestedSelectControl {
    a.Set("hidden", value)
    return a
}

/**
 * 表单项类型
 */
func (a *NestedSelectControl) Type(value interface{}) *NestedSelectControl {
    a.Set("type", value)
    return a
}

/**
 * 弹框的 css 类
 */
func (a *NestedSelectControl) Menuclassname(value interface{}) *NestedSelectControl {
    a.Set("menuClassName", value)
    return a
}

/**
 * 新增文字
 */
func (a *NestedSelectControl) Createbtnlabel(value interface{}) *NestedSelectControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * 控制编辑弹框设置项
 */
func (a *NestedSelectControl) Editdialog(value interface{}) *NestedSelectControl {
    a.Set("editDialog", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *NestedSelectControl) Horizontal(value interface{}) *NestedSelectControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 占位符
 */
func (a *NestedSelectControl) Placeholder(value interface{}) *NestedSelectControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *NestedSelectControl) Validationerrors(value interface{}) *NestedSelectControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 描述标题
 */
func (a *NestedSelectControl) Labelalign(value interface{}) *NestedSelectControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 配置 label className
 */
func (a *NestedSelectControl) Labelclassname(value interface{}) *NestedSelectControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *NestedSelectControl) Submitonchange(value interface{}) *NestedSelectControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 默认选择选项第一个值。
 */
func (a *NestedSelectControl) Selectfirst(value interface{}) *NestedSelectControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * 添加时调用的接口
 */
func (a *NestedSelectControl) Addapi(value interface{}) *NestedSelectControl {
    a.Set("addApi", value)
    return a
}

/**
 * 选项修改的表单项
 */
func (a *NestedSelectControl) Editcontrols(value interface{}) *NestedSelectControl {
    a.Set("editControls", value)
    return a
}

/**
 * 选项删除 API
 */
func (a *NestedSelectControl) Deleteapi(value interface{}) *NestedSelectControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *NestedSelectControl) Size(value interface{}) *NestedSelectControl {
    a.Set("size", value)
    return a
}

/**
 * 只读条件
 */
func (a *NestedSelectControl) Readonlyon(value interface{}) *NestedSelectControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *NestedSelectControl) Description(value interface{}) *NestedSelectControl {
    a.Set("description", value)
    return a
}

/**
 */
func (a *NestedSelectControl) Desc(value interface{}) *NestedSelectControl {
    a.Set("desc", value)
    return a
}

/**
 * 是否隐藏选择框中已选中节点的祖先节点的文本信息
 */
func (a *NestedSelectControl) Hidenodepathlabel(value interface{}) *NestedSelectControl {
    a.Set("hideNodePathLabel", value)
    return a
}

/**
 * 配置 source 接口初始拉不拉取。
 */
func (a *NestedSelectControl) Initfetch(value interface{}) *NestedSelectControl {
    a.Set("initFetch", value)
    return a
}

/**
 * 分割符
 */
func (a *NestedSelectControl) Delimiter(value interface{}) *NestedSelectControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 是否可清除。
 */
func (a *NestedSelectControl) Clearable(value interface{}) *NestedSelectControl {
    a.Set("clearable", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *NestedSelectControl) Remark(value interface{}) *NestedSelectControl {
    a.Set("remark", value)
    return a
}

/**
 * 多选模式，值太多时是否避免折行
 */
func (a *NestedSelectControl) Valuesnowrap(value interface{}) *NestedSelectControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * 新增时的表单项。
 */
func (a *NestedSelectControl) Addcontrols(value interface{}) *NestedSelectControl {
    a.Set("addControls", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *NestedSelectControl) Labelwidth(value interface{}) *NestedSelectControl {
    a.Set("labelWidth", value)
    return a
}

/**
 */
func (a *NestedSelectControl) Initautofill(value interface{}) *NestedSelectControl {
    a.Set("initAutoFill", value)
    return a
}
