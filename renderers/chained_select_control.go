package renderers


/**
 * 链式下拉框 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/chained-select

 * @author wcz0
 * @version 6.2.2
 */
type ChainedSelectControl struct {
	*BaseRenderer
}

func NewChainedSelectControl() *ChainedSelectControl {
    a := &ChainedSelectControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "chained-select")
    return a
}


func (a *ChainedSelectControl) Set(name string, value interface{}) *ChainedSelectControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 是否隐藏表达式
 */
func (a *ChainedSelectControl) HiddenOn(value interface{}) *ChainedSelectControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 选项集合
 */
func (a *ChainedSelectControl) Options(value interface{}) *ChainedSelectControl {
    a.Set("options", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *ChainedSelectControl) Value(value interface{}) *ChainedSelectControl {
    a.Set("value", value)
    return a
}

/**
 */
func (a *ChainedSelectControl) Row(value interface{}) *ChainedSelectControl {
    a.Set("row", value)
    return a
}

/**
 * 描述标题
 */
func (a *ChainedSelectControl) LabelAlign(value interface{}) *ChainedSelectControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *ChainedSelectControl) ValidateApi(value interface{}) *ChainedSelectControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *ChainedSelectControl) AutoFill(value interface{}) *ChainedSelectControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 点清除按钮时，将表单项设置成当前配置的值。
 */
func (a *ChainedSelectControl) ResetValue(value interface{}) *ChainedSelectControl {
    a.Set("resetValue", value)
    return a
}

/**
 * 描述标题
 */
func (a *ChainedSelectControl) Label(value interface{}) *ChainedSelectControl {
    a.Set("label", value)
    return a
}

/**
 */
func (a *ChainedSelectControl) Desc(value interface{}) *ChainedSelectControl {
    a.Set("desc", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *ChainedSelectControl) StaticLabelClassName(value interface{}) *ChainedSelectControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *ChainedSelectControl) Style(value interface{}) *ChainedSelectControl {
    a.Set("style", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *ChainedSelectControl) EditorSetting(value interface{}) *ChainedSelectControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *ChainedSelectControl) Width(value interface{}) *ChainedSelectControl {
    a.Set("width", value)
    return a
}

/**
 * 分割符
 */
func (a *ChainedSelectControl) Delimiter(value interface{}) *ChainedSelectControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *ChainedSelectControl) Hint(value interface{}) *ChainedSelectControl {
    a.Set("hint", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *ChainedSelectControl) StaticClassName(value interface{}) *ChainedSelectControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
 */
func (a *ChainedSelectControl) ExtractValue(value interface{}) *ChainedSelectControl {
    a.Set("extractValue", value)
    return a
}

/**
 * 是否可清除。
 */
func (a *ChainedSelectControl) Clearable(value interface{}) *ChainedSelectControl {
    a.Set("clearable", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *ChainedSelectControl) Inline(value interface{}) *ChainedSelectControl {
    a.Set("inline", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *ChainedSelectControl) StaticPlaceholder(value interface{}) *ChainedSelectControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 表单项类型
 */
func (a *ChainedSelectControl) Type(value interface{}) *ChainedSelectControl {
    a.Set("type", value)
    return a
}

/**
 * 默认选择选项第一个值。
 */
func (a *ChainedSelectControl) SelectFirst(value interface{}) *ChainedSelectControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * 选项修改的表单项
 */
func (a *ChainedSelectControl) EditControls(value interface{}) *ChainedSelectControl {
    a.Set("editControls", value)
    return a
}

/**
 * 配置 label className
 */
func (a *ChainedSelectControl) LabelClassName(value interface{}) *ChainedSelectControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 只读条件
 */
func (a *ChainedSelectControl) ReadOnlyOn(value interface{}) *ChainedSelectControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 是否为必填
 */
func (a *ChainedSelectControl) Required(value interface{}) *ChainedSelectControl {
    a.Set("required", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *ChainedSelectControl) OnEvent(value interface{}) *ChainedSelectControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 添加时调用的接口
 */
func (a *ChainedSelectControl) AddApi(value interface{}) *ChainedSelectControl {
    a.Set("addApi", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *ChainedSelectControl) Size(value interface{}) *ChainedSelectControl {
    a.Set("size", value)
    return a
}

/**
 * 是否可删除
 */
func (a *ChainedSelectControl) Removable(value interface{}) *ChainedSelectControl {
    a.Set("removable", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *ChainedSelectControl) LabelWidth(value interface{}) *ChainedSelectControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *ChainedSelectControl) Horizontal(value interface{}) *ChainedSelectControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 懒加载字段
 */
func (a *ChainedSelectControl) DeferField(value interface{}) *ChainedSelectControl {
    a.Set("deferField", value)
    return a
}

/**
 * 控制新增弹框设置项
 */
func (a *ChainedSelectControl) AddDialog(value interface{}) *ChainedSelectControl {
    a.Set("addDialog", value)
    return a
}

/**
 * 新增文字
 */
func (a *ChainedSelectControl) CreateBtnLabel(value interface{}) *ChainedSelectControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * 控制编辑弹框设置项
 */
func (a *ChainedSelectControl) EditDialog(value interface{}) *ChainedSelectControl {
    a.Set("editDialog", value)
    return a
}

/**
 * 选项删除提示文字。
 */
func (a *ChainedSelectControl) DeleteConfirmText(value interface{}) *ChainedSelectControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *ChainedSelectControl) Mode(value interface{}) *ChainedSelectControl {
    a.Set("mode", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *ChainedSelectControl) Static(value interface{}) *ChainedSelectControl {
    a.Set("static", value)
    return a
}

/**
 */
func (a *ChainedSelectControl) StaticSchema(value interface{}) *ChainedSelectControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 配置 source 接口初始拉不拉取。
 */
func (a *ChainedSelectControl) InitFetch(value interface{}) *ChainedSelectControl {
    a.Set("initFetch", value)
    return a
}

/**
 * 是否可以新增
 */
func (a *ChainedSelectControl) Creatable(value interface{}) *ChainedSelectControl {
    a.Set("creatable", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *ChainedSelectControl) Remark(value interface{}) *ChainedSelectControl {
    a.Set("remark", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *ChainedSelectControl) SubmitOnChange(value interface{}) *ChainedSelectControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *ChainedSelectControl) Description(value interface{}) *ChainedSelectControl {
    a.Set("description", value)
    return a
}

/**
 * 可用来通过 API 拉取 options。
 */
func (a *ChainedSelectControl) Source(value interface{}) *ChainedSelectControl {
    a.Set("source", value)
    return a
}

/**
 * 多选模式，值太多时是否避免折行
 */
func (a *ChainedSelectControl) ValuesNoWrap(value interface{}) *ChainedSelectControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * 占位符
 */
func (a *ChainedSelectControl) Placeholder(value interface{}) *ChainedSelectControl {
    a.Set("placeholder", value)
    return a
}

/**
 */
func (a *ChainedSelectControl) Validations(value interface{}) *ChainedSelectControl {
    a.Set("validations", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *ChainedSelectControl) StaticInputClassName(value interface{}) *ChainedSelectControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 新增时的表单项。
 */
func (a *ChainedSelectControl) AddControls(value interface{}) *ChainedSelectControl {
    a.Set("addControls", value)
    return a
}

/**
 * 选项删除 API
 */
func (a *ChainedSelectControl) DeleteApi(value interface{}) *ChainedSelectControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *ChainedSelectControl) ExtraName(value interface{}) *ChainedSelectControl {
    a.Set("extraName", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *ChainedSelectControl) ValidateOnChange(value interface{}) *ChainedSelectControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 配置 input className
 */
func (a *ChainedSelectControl) InputClassName(value interface{}) *ChainedSelectControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 是否禁用
 */
func (a *ChainedSelectControl) Disabled(value interface{}) *ChainedSelectControl {
    a.Set("disabled", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *ChainedSelectControl) DisabledOn(value interface{}) *ChainedSelectControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *ChainedSelectControl) Id(value interface{}) *ChainedSelectControl {
    a.Set("id", value)
    return a
}

/**
 * 是否为多选模式
 */
func (a *ChainedSelectControl) Multiple(value interface{}) *ChainedSelectControl {
    a.Set("multiple", value)
    return a
}

/**
 * 编辑时调用的 API
 */
func (a *ChainedSelectControl) EditApi(value interface{}) *ChainedSelectControl {
    a.Set("editApi", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *ChainedSelectControl) Hidden(value interface{}) *ChainedSelectControl {
    a.Set("hidden", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *ChainedSelectControl) LabelRemark(value interface{}) *ChainedSelectControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 用表达式来配置 source 接口初始要不要拉取
 */
func (a *ChainedSelectControl) InitFetchOn(value interface{}) *ChainedSelectControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 是否只读
 */
func (a *ChainedSelectControl) ReadOnly(value interface{}) *ChainedSelectControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 是否显示
 */
func (a *ChainedSelectControl) Visible(value interface{}) *ChainedSelectControl {
    a.Set("visible", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *ChainedSelectControl) StaticOn(value interface{}) *ChainedSelectControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *ChainedSelectControl) UseMobileUI(value interface{}) *ChainedSelectControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *ChainedSelectControl) Name(value interface{}) *ChainedSelectControl {
    a.Set("name", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *ChainedSelectControl) ClearValueOnHidden(value interface{}) *ChainedSelectControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *ChainedSelectControl) ClassName(value interface{}) *ChainedSelectControl {
    a.Set("className", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *ChainedSelectControl) VisibleOn(value interface{}) *ChainedSelectControl {
    a.Set("visibleOn", value)
    return a
}

/**
 */
func (a *ChainedSelectControl) TestIdBuilder(value interface{}) *ChainedSelectControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
 */
func (a *ChainedSelectControl) DeferApi(value interface{}) *ChainedSelectControl {
    a.Set("deferApi", value)
    return a
}

/**
 */
func (a *ChainedSelectControl) InitAutoFill(value interface{}) *ChainedSelectControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *ChainedSelectControl) DescriptionClassName(value interface{}) *ChainedSelectControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *ChainedSelectControl) ValidationErrors(value interface{}) *ChainedSelectControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
 */
func (a *ChainedSelectControl) JoinValues(value interface{}) *ChainedSelectControl {
    a.Set("joinValues", value)
    return a
}

/**
 * 是否可以编辑
 */
func (a *ChainedSelectControl) Editable(value interface{}) *ChainedSelectControl {
    a.Set("editable", value)
    return a
}
