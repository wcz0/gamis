package renderers


/**
 * UserSelect 移动端人员选择。

 * @author wcz0
 * @version 6.2.2
 */
type UserSelectControl struct {
	*BaseRenderer
}

func NewUserSelectControl() *UserSelectControl {
    a := &UserSelectControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "users-select")
    return a
}


func (a *UserSelectControl) Set(name string, value interface{}) *UserSelectControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 多选模式，值太多时是否避免折行
 */
func (a *UserSelectControl) ValuesNoWrap(value interface{}) *UserSelectControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * 选项修改的表单项
 */
func (a *UserSelectControl) EditControls(value interface{}) *UserSelectControl {
    a.Set("editControls", value)
    return a
}

/**
 * 选项删除 API
 */
func (a *UserSelectControl) DeleteApi(value interface{}) *UserSelectControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *UserSelectControl) DescriptionClassName(value interface{}) *UserSelectControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *UserSelectControl) StaticInputClassName(value interface{}) *UserSelectControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *UserSelectControl) Width(value interface{}) *UserSelectControl {
    a.Set("width", value)
    return a
}

/**
 * 是否为多选模式
 */
func (a *UserSelectControl) Multiple(value interface{}) *UserSelectControl {
    a.Set("multiple", value)
    return a
}

/**
 * 点清除按钮时，将表单项设置成当前配置的值。
 */
func (a *UserSelectControl) ResetValue(value interface{}) *UserSelectControl {
    a.Set("resetValue", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *UserSelectControl) ValidateApi(value interface{}) *UserSelectControl {
    a.Set("validateApi", value)
    return a
}

/**
 */
func (a *UserSelectControl) StaticSchema(value interface{}) *UserSelectControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *UserSelectControl) StaticClassName(value interface{}) *UserSelectControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *UserSelectControl) EditorSetting(value interface{}) *UserSelectControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *UserSelectControl) Size(value interface{}) *UserSelectControl {
    a.Set("size", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *UserSelectControl) Mode(value interface{}) *UserSelectControl {
    a.Set("mode", value)
    return a
}

/**
 * 占位符
 */
func (a *UserSelectControl) Placeholder(value interface{}) *UserSelectControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *UserSelectControl) StaticOn(value interface{}) *UserSelectControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 默认选择选项第一个值。
 */
func (a *UserSelectControl) SelectFirst(value interface{}) *UserSelectControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *UserSelectControl) LabelWidth(value interface{}) *UserSelectControl {
    a.Set("labelWidth", value)
    return a
}

/**
 */
func (a *UserSelectControl) Desc(value interface{}) *UserSelectControl {
    a.Set("desc", value)
    return a
}

/**
 * 是否可清除。
 */
func (a *UserSelectControl) Clearable(value interface{}) *UserSelectControl {
    a.Set("clearable", value)
    return a
}

/**
 * 懒加载字段
 */
func (a *UserSelectControl) DeferField(value interface{}) *UserSelectControl {
    a.Set("deferField", value)
    return a
}

/**
 * 配置 label className
 */
func (a *UserSelectControl) LabelClassName(value interface{}) *UserSelectControl {
    a.Set("labelClassName", value)
    return a
}

/**
 */
func (a *UserSelectControl) Row(value interface{}) *UserSelectControl {
    a.Set("row", value)
    return a
}

/**
 * 是否禁用
 */
func (a *UserSelectControl) Disabled(value interface{}) *UserSelectControl {
    a.Set("disabled", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *UserSelectControl) AutoFill(value interface{}) *UserSelectControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 组件样式
 */
func (a *UserSelectControl) Style(value interface{}) *UserSelectControl {
    a.Set("style", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *UserSelectControl) ExtraName(value interface{}) *UserSelectControl {
    a.Set("extraName", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *UserSelectControl) Hint(value interface{}) *UserSelectControl {
    a.Set("hint", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *UserSelectControl) ValidateOnChange(value interface{}) *UserSelectControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 配置 input className
 */
func (a *UserSelectControl) InputClassName(value interface{}) *UserSelectControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *UserSelectControl) ValidationErrors(value interface{}) *UserSelectControl {
    a.Set("validationErrors", value)
    return a
}

/**
 */
func (a *UserSelectControl) TestIdBuilder(value interface{}) *UserSelectControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 控制编辑弹框设置项
 */
func (a *UserSelectControl) EditDialog(value interface{}) *UserSelectControl {
    a.Set("editDialog", value)
    return a
}

/**
 */
func (a *UserSelectControl) InitAutoFill(value interface{}) *UserSelectControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 是否显示
 */
func (a *UserSelectControl) Visible(value interface{}) *UserSelectControl {
    a.Set("visible", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *UserSelectControl) Static(value interface{}) *UserSelectControl {
    a.Set("static", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *UserSelectControl) UseMobileUI(value interface{}) *UserSelectControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 描述标题
 */
func (a *UserSelectControl) LabelAlign(value interface{}) *UserSelectControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 是否只读
 */
func (a *UserSelectControl) ReadOnly(value interface{}) *UserSelectControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *UserSelectControl) Description(value interface{}) *UserSelectControl {
    a.Set("description", value)
    return a
}

/**
 * 用表达式来配置 source 接口初始要不要拉取
 */
func (a *UserSelectControl) InitFetchOn(value interface{}) *UserSelectControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 分割符
 */
func (a *UserSelectControl) Delimiter(value interface{}) *UserSelectControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
 */
func (a *UserSelectControl) ExtractValue(value interface{}) *UserSelectControl {
    a.Set("extractValue", value)
    return a
}

/**
 * 新增文字
 */
func (a *UserSelectControl) CreateBtnLabel(value interface{}) *UserSelectControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * 选项删除提示文字。
 */
func (a *UserSelectControl) DeleteConfirmText(value interface{}) *UserSelectControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *UserSelectControl) ClassName(value interface{}) *UserSelectControl {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *UserSelectControl) Hidden(value interface{}) *UserSelectControl {
    a.Set("hidden", value)
    return a
}

/**
 * 编辑时调用的 API
 */
func (a *UserSelectControl) EditApi(value interface{}) *UserSelectControl {
    a.Set("editApi", value)
    return a
}

/**
 * 是否可删除
 */
func (a *UserSelectControl) Removable(value interface{}) *UserSelectControl {
    a.Set("removable", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *UserSelectControl) SubmitOnChange(value interface{}) *UserSelectControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *UserSelectControl) HiddenOn(value interface{}) *UserSelectControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *UserSelectControl) StaticLabelClassName(value interface{}) *UserSelectControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 添加时调用的接口
 */
func (a *UserSelectControl) AddApi(value interface{}) *UserSelectControl {
    a.Set("addApi", value)
    return a
}

/**
 * 是否可以新增
 */
func (a *UserSelectControl) Creatable(value interface{}) *UserSelectControl {
    a.Set("creatable", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *UserSelectControl) OnEvent(value interface{}) *UserSelectControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 表单项类型
 */
func (a *UserSelectControl) Type(value interface{}) *UserSelectControl {
    a.Set("type", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *UserSelectControl) Name(value interface{}) *UserSelectControl {
    a.Set("name", value)
    return a
}

/**
 */
func (a *UserSelectControl) Validations(value interface{}) *UserSelectControl {
    a.Set("validations", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *UserSelectControl) ClearValueOnHidden(value interface{}) *UserSelectControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *UserSelectControl) VisibleOn(value interface{}) *UserSelectControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *UserSelectControl) Remark(value interface{}) *UserSelectControl {
    a.Set("remark", value)
    return a
}

/**
 * 可用来通过 API 拉取 options。
 */
func (a *UserSelectControl) Source(value interface{}) *UserSelectControl {
    a.Set("source", value)
    return a
}

/**
 * 配置 source 接口初始拉不拉取。
 */
func (a *UserSelectControl) InitFetch(value interface{}) *UserSelectControl {
    a.Set("initFetch", value)
    return a
}

/**
 * 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
 */
func (a *UserSelectControl) JoinValues(value interface{}) *UserSelectControl {
    a.Set("joinValues", value)
    return a
}

/**
 * 是否可以编辑
 */
func (a *UserSelectControl) Editable(value interface{}) *UserSelectControl {
    a.Set("editable", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *UserSelectControl) DisabledOn(value interface{}) *UserSelectControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *UserSelectControl) Id(value interface{}) *UserSelectControl {
    a.Set("id", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *UserSelectControl) LabelRemark(value interface{}) *UserSelectControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
 */
func (a *UserSelectControl) DeferApi(value interface{}) *UserSelectControl {
    a.Set("deferApi", value)
    return a
}

/**
 * 新增时的表单项。
 */
func (a *UserSelectControl) AddControls(value interface{}) *UserSelectControl {
    a.Set("addControls", value)
    return a
}

/**
 * 控制新增弹框设置项
 */
func (a *UserSelectControl) AddDialog(value interface{}) *UserSelectControl {
    a.Set("addDialog", value)
    return a
}

/**
 * 只读条件
 */
func (a *UserSelectControl) ReadOnlyOn(value interface{}) *UserSelectControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *UserSelectControl) Value(value interface{}) *UserSelectControl {
    a.Set("value", value)
    return a
}

/**
 * 选项集合
 */
func (a *UserSelectControl) Options(value interface{}) *UserSelectControl {
    a.Set("options", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *UserSelectControl) Inline(value interface{}) *UserSelectControl {
    a.Set("inline", value)
    return a
}

/**
 * 是否为必填
 */
func (a *UserSelectControl) Required(value interface{}) *UserSelectControl {
    a.Set("required", value)
    return a
}

/**
 * 描述标题
 */
func (a *UserSelectControl) Label(value interface{}) *UserSelectControl {
    a.Set("label", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *UserSelectControl) Horizontal(value interface{}) *UserSelectControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *UserSelectControl) StaticPlaceholder(value interface{}) *UserSelectControl {
    a.Set("staticPlaceholder", value)
    return a
}
