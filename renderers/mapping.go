package renderers


/**
 * Mapping 映射展示控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/mapping

 * @author wcz0
 * @version 6.2.2
 */
type Mapping struct {
	*BaseRenderer
}

func NewMapping() *Mapping {
    a := &Mapping{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "map")
    return a
}


func (a *Mapping) Set(name string, value interface{}) *Mapping {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 静态展示表单项Label类名
 */
func (a *Mapping) StaticLabelClassName(value interface{}) *Mapping {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *Mapping) TestIdBuilder(value interface{}) *Mapping {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 是否显示
 */
func (a *Mapping) Visible(value interface{}) *Mapping {
    a.Set("visible", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Mapping) Id(value interface{}) *Mapping {
    a.Set("id", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Mapping) Static(value interface{}) *Mapping {
    a.Set("static", value)
    return a
}

/**
 * map或source为对象数组时，作为value值的字段名
 */
func (a *Mapping) ValueField(value interface{}) *Mapping {
    a.Set("valueField", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Mapping) OnEvent(value interface{}) *Mapping {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Mapping) StaticClassName(value interface{}) *Mapping {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Mapping) StaticInputClassName(value interface{}) *Mapping {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *Mapping) Style(value interface{}) *Mapping {
    a.Set("style", value)
    return a
}

/**
 * 关联字段名。
 */
func (a *Mapping) Name(value interface{}) *Mapping {
    a.Set("name", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Mapping) Hidden(value interface{}) *Mapping {
    a.Set("hidden", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Mapping) HiddenOn(value interface{}) *Mapping {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Mapping) VisibleOn(value interface{}) *Mapping {
    a.Set("visibleOn", value)
    return a
}

/**
 * map或source为对象数组时，作为label值的字段名
 */
func (a *Mapping) LabelField(value interface{}) *Mapping {
    a.Set("labelField", value)
    return a
}

/**
 * 自定义渲染映射值，支持html或schema
 */
func (a *Mapping) ItemSchema(value interface{}) *Mapping {
    a.Set("itemSchema", value)
    return a
}

/**
 * 如果想远程拉取字典，请配置 source 为接口。
 */
func (a *Mapping) Source(value interface{}) *Mapping {
    a.Set("source", value)
    return a
}

/**
 * 指定为映射展示控件
 * 可选值: map | mapping
 */
func (a *Mapping) Type(value interface{}) *Mapping {
    a.Set("type", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Mapping) ClassName(value interface{}) *Mapping {
    a.Set("className", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Mapping) UseMobileUI(value interface{}) *Mapping {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Mapping) EditorSetting(value interface{}) *Mapping {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *Mapping) Testid(value interface{}) *Mapping {
    a.Set("testid", value)
    return a
}

/**
 * 配置映射规则，值可以使用模板语法。当 key 为 * 时表示 else，也就是说值没有映射到任何规则时用 * 对应的值展示。
 */
func (a *Mapping) Map(value interface{}) *Mapping {
    a.Set("map", value)
    return a
}

/**
 * 占位符
 */
func (a *Mapping) Placeholder(value interface{}) *Mapping {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Mapping) Disabled(value interface{}) *Mapping {
    a.Set("disabled", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Mapping) StaticOn(value interface{}) *Mapping {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *Mapping) StaticSchema(value interface{}) *Mapping {
    a.Set("staticSchema", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Mapping) DisabledOn(value interface{}) *Mapping {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Mapping) StaticPlaceholder(value interface{}) *Mapping {
    a.Set("staticPlaceholder", value)
    return a
}
