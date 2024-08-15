package renderers


/**
 * InputArray 数组输入框。 combo 的别名。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/array

 * @author wcz0
 * @version 6.2.2
 */
type ArrayControl struct {
	*BaseRenderer
}

func NewArrayControl() *ArrayControl {
    a := &ArrayControl{
        BaseRenderer: NewBaseRenderer(),
    }

func (a *ArrayControl) Set(name string, value interface{}) *ArrayControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}

    a.Set("type", "input-array")
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *ArrayControl) Validateonchange(value interface{}) *ArrayControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 是否为必填
 */
func (a *ArrayControl) Required(value interface{}) *ArrayControl {
    a.Set("required", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *ArrayControl) Id(value interface{}) *ArrayControl {
    a.Set("id", value)
    return a
}

/**
 * 成员渲染器配置
 */
func (a *ArrayControl) Items(value interface{}) *ArrayControl {
    a.Set("items", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *ArrayControl) Width(value interface{}) *ArrayControl {
    a.Set("width", value)
    return a
}

/**
 * 配置同步字段。只有 `strictMode` 为 `false` 时有效。 如果 Combo 层级比较深，底层的获取外层的数据可能不同步。 但是给 combo 配置这个属性就能同步下来。输入格式：`["os"]`
 */
func (a *ArrayControl) Syncfields(value interface{}) *ArrayControl {
    a.Set("syncFields", value)
    return a
}

/**
 * 描述标题
 */
func (a *ArrayControl) Labelalign(value interface{}) *ArrayControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 是否可以访问父级数据，正常 combo 已经关联到数组成员，是不能访问父级数据的。
 */
func (a *ArrayControl) Canaccesssuperdata(value interface{}) *ArrayControl {
    a.Set("canAccessSuperData", value)
    return a
}

/**
 */
func (a *ArrayControl) Testidbuilder(value interface{}) *ArrayControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *ArrayControl) Descriptionclassname(value interface{}) *ArrayControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 是否可切换条件，配合`conditions`使用
 */
func (a *ArrayControl) Typeswitchable(value interface{}) *ArrayControl {
    a.Set("typeSwitchable", value)
    return a
}

/**
 * 新增按钮CSS类名
 */
func (a *ArrayControl) Addbuttonclassname(value interface{}) *ArrayControl {
    a.Set("addButtonClassName", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *ArrayControl) Horizontal(value interface{}) *ArrayControl {
    a.Set("horizontal", value)
    return a
}

/**
 */
func (a *ArrayControl) Initautofill(value interface{}) *ArrayControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 是否禁用
 */
func (a *ArrayControl) Disabled(value interface{}) *ArrayControl {
    a.Set("disabled", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *ArrayControl) Onevent(value interface{}) *ArrayControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *ArrayControl) Editorsetting(value interface{}) *ArrayControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 可拖拽排序的提示信息。
 */
func (a *ArrayControl) Draggabletip(value interface{}) *ArrayControl {
    a.Set("draggableTip", value)
    return a
}

/**
 * 是否可多选
 */
func (a *ArrayControl) Multiple(value interface{}) *ArrayControl {
    a.Set("multiple", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *ArrayControl) Labelwidth(value interface{}) *ArrayControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *ArrayControl) Hint(value interface{}) *ArrayControl {
    a.Set("hint", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *ArrayControl) Validationerrors(value interface{}) *ArrayControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 组件样式
 */
func (a *ArrayControl) Style(value interface{}) *ArrayControl {
    a.Set("style", value)
    return a
}

/**
 * 是否可新增
 */
func (a *ArrayControl) Addable(value interface{}) *ArrayControl {
    a.Set("addable", value)
    return a
}

/**
 * 允许为空，如果子表单项里面配置验证器，且又是单条模式。可以允许用户选择清空（不填）。
 */
func (a *ArrayControl) Nullable(value interface{}) *ArrayControl {
    a.Set("nullable", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *ArrayControl) Clearvalueonhidden(value interface{}) *ArrayControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * Add at top
 */
func (a *ArrayControl) Addattop(value interface{}) *ArrayControl {
    a.Set("addattop", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *ArrayControl) Name(value interface{}) *ArrayControl {
    a.Set("name", value)
    return a
}

/**
 */
func (a *ArrayControl) Updatepristineafterstoredatareinit(value interface{}) *ArrayControl {
    a.Set("updatePristineAfterStoreDataReInit", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *ArrayControl) Description(value interface{}) *ArrayControl {
    a.Set("description", value)
    return a
}

/**
 */
func (a *ArrayControl) Desc(value interface{}) *ArrayControl {
    a.Set("desc", value)
    return a
}

/**
 * 是否显示
 */
func (a *ArrayControl) Visible(value interface{}) *ArrayControl {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *ArrayControl) Staticplaceholder(value interface{}) *ArrayControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *ArrayControl) Staticclassname(value interface{}) *ArrayControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 子表单的模式。
 * 可选值: normal | horizontal | inline
 */
func (a *ArrayControl) Subformmode(value interface{}) *ArrayControl {
    a.Set("subFormMode", value)
    return a
}

/**
 * 提示信息
 */
func (a *ArrayControl) Messages(value interface{}) *ArrayControl {
    a.Set("messages", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *ArrayControl) Staticlabelclassname(value interface{}) *ArrayControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *ArrayControl) Labelremark(value interface{}) *ArrayControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 描述标题
 */
func (a *ArrayControl) Label(value interface{}) *ArrayControl {
    a.Set("label", value)
    return a
}

/**
 * 是否只读
 */
func (a *ArrayControl) Readonly(value interface{}) *ArrayControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *ArrayControl) Inline(value interface{}) *ArrayControl {
    a.Set("inline", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *ArrayControl) Classname(value interface{}) *ArrayControl {
    a.Set("className", value)
    return a
}

/**
 * 是否多行模式，默认一行展示完
 */
func (a *ArrayControl) Multiline(value interface{}) *ArrayControl {
    a.Set("multiLine", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *ArrayControl) Remark(value interface{}) *ArrayControl {
    a.Set("remark", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *ArrayControl) Disabledon(value interface{}) *ArrayControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 配置 input className
 */
func (a *ArrayControl) Inputclassname(value interface{}) *ArrayControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *ArrayControl) Validateapi(value interface{}) *ArrayControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *ArrayControl) Staticinputclassname(value interface{}) *ArrayControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 限制最小个数
 */
func (a *ArrayControl) Minlength(value interface{}) *ArrayControl {
    a.Set("minLength", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *ArrayControl) Submitonchange(value interface{}) *ArrayControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 当扁平化开启并且joinValues为true时，用什么分隔符
 */
func (a *ArrayControl) Delimiter(value interface{}) *ArrayControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 当扁平化开启的时候，是否用分隔符的形式发送给后端，否则采用array的方式
 */
func (a *ArrayControl) Joinvalues(value interface{}) *ArrayControl {
    a.Set("joinValues", value)
    return a
}

/**
 * 是否可删除
 */
func (a *ArrayControl) Removable(value interface{}) *ArrayControl {
    a.Set("removable", value)
    return a
}

/**
 * 采用 Tabs 展示方式？
 */
func (a *ArrayControl) Tabsmode(value interface{}) *ArrayControl {
    a.Set("tabsMode", value)
    return a
}

/**
 * Tabs 的展示模式。
 * 可选值:  | line | card | radio
 */
func (a *ArrayControl) Tabsstyle(value interface{}) *ArrayControl {
    a.Set("tabsStyle", value)
    return a
}

/**
 * 数据比较多，比较卡时，可以试试开启。
 */
func (a *ArrayControl) Lazyload(value interface{}) *ArrayControl {
    a.Set("lazyLoad", value)
    return a
}

/**
 * 是否含有边框
 */
func (a *ArrayControl) Noborder(value interface{}) *ArrayControl {
    a.Set("noBorder", value)
    return a
}

/**
 * 内部单组表单项的类名
 */
func (a *ArrayControl) Formclassname(value interface{}) *ArrayControl {
    a.Set("formClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *ArrayControl) Usemobileui(value interface{}) *ArrayControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 指定为数组输入框类型
 */
func (a *ArrayControl) Type(value interface{}) *ArrayControl {
    a.Set("type", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *ArrayControl) Size(value interface{}) *ArrayControl {
    a.Set("size", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *ArrayControl) Static(value interface{}) *ArrayControl {
    a.Set("static", value)
    return a
}

/**
 * 确认删除时的提示
 */
func (a *ArrayControl) Deleteconfirmtext(value interface{}) *ArrayControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * 选项卡标题的生成模板。
 */
func (a *ArrayControl) Tabslabeltpl(value interface{}) *ArrayControl {
    a.Set("tabsLabelTpl", value)
    return a
}

/**
 * 配置 label className
 */
func (a *ArrayControl) Labelclassname(value interface{}) *ArrayControl {
    a.Set("labelClassName", value)
    return a
}

/**
 */
func (a *ArrayControl) Validations(value interface{}) *ArrayControl {
    a.Set("validations", value)
    return a
}

/**
 * 新增成员时的默认值
 */
func (a *ArrayControl) Scaffold(value interface{}) *ArrayControl {
    a.Set("scaffold", value)
    return a
}

/**
 * 严格模式，为了性能默认不开的。
 */
func (a *ArrayControl) Strictmode(value interface{}) *ArrayControl {
    a.Set("strictMode", value)
    return a
}

/**
 * 只读条件
 */
func (a *ArrayControl) Readonlyon(value interface{}) *ArrayControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *ArrayControl) Value(value interface{}) *ArrayControl {
    a.Set("value", value)
    return a
}

/**
 */
func (a *ArrayControl) Row(value interface{}) *ArrayControl {
    a.Set("row", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *ArrayControl) Visibleon(value interface{}) *ArrayControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 如果是水平排版，这个属性可以细化水平排版的左右宽度占比。
 */
func (a *ArrayControl) Subformhorizontal(value interface{}) *ArrayControl {
    a.Set("subFormHorizontal", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *ArrayControl) Extraname(value interface{}) *ArrayControl {
    a.Set("extraName", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *ArrayControl) Autofill(value interface{}) *ArrayControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *ArrayControl) Hidden(value interface{}) *ArrayControl {
    a.Set("hidden", value)
    return a
}

/**
 * 删除时调用的api
 */
func (a *ArrayControl) Deleteapi(value interface{}) *ArrayControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *ArrayControl) Mode(value interface{}) *ArrayControl {
    a.Set("mode", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *ArrayControl) Staticon(value interface{}) *ArrayControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 是否可拖拽排序
 */
func (a *ArrayControl) Draggable(value interface{}) *ArrayControl {
    a.Set("draggable", value)
    return a
}

/**
 * 限制最大个数
 */
func (a *ArrayControl) Maxlength(value interface{}) *ArrayControl {
    a.Set("maxLength", value)
    return a
}

/**
 * 没有成员时显示。
 */
func (a *ArrayControl) Placeholder(value interface{}) *ArrayControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *ArrayControl) Hiddenon(value interface{}) *ArrayControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 */
func (a *ArrayControl) Staticschema(value interface{}) *ArrayControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 新增按钮文字
 */
func (a *ArrayControl) Addbuttontext(value interface{}) *ArrayControl {
    a.Set("addButtonText", value)
    return a
}

/**
 * 是否将结果扁平化(去掉name),只有当controls的length为1且multiple为true的时候才有效
 */
func (a *ArrayControl) Flat(value interface{}) *ArrayControl {
    a.Set("flat", value)
    return a
}
