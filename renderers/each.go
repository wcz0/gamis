package renderers


/**
 * Each 循环功能渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/each

 * @author wcz0
 * @version 6.2.2
 */
type Each struct {
	*BaseRenderer
}

func NewEach() *Each {
    a := &Each{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "each")
    return a
}


func (a *Each) Set(name string, value interface{}) *Each {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 静态展示表单项类名
 */
func (a *Each) StaticClassName(value interface{}) *Each {
    a.Set("staticClassName", value)
    return a
}

/**
 * 用来控制通过什么字段读取序号，考虑到可能多层嵌套 如果名字一样会读取不到上层变量，所以这里可以指定一下
 */
func (a *Each) IndexKeyName(value interface{}) *Each {
    a.Set("indexKeyName", value)
    return a
}

/**
 */
func (a *Each) Items(value interface{}) *Each {
    a.Set("items", value)
    return a
}

/**
 * 用来控制通过什么字段读取成员数据，考虑到可能多层嵌套 如果名字一样会读取不到上层变量，所以这里可以指定一下
 */
func (a *Each) ItemKeyName(value interface{}) *Each {
    a.Set("itemKeyName", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Each) ClassName(value interface{}) *Each {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Each) DisabledOn(value interface{}) *Each {
    a.Set("disabledOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Each) Id(value interface{}) *Each {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Each) StaticLabelClassName(value interface{}) *Each {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *Each) Style(value interface{}) *Each {
    a.Set("style", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Each) EditorSetting(value interface{}) *Each {
    a.Set("editorSetting", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Each) UseMobileUI(value interface{}) *Each {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 是否显示
 */
func (a *Each) Visible(value interface{}) *Each {
    a.Set("visible", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Each) Hidden(value interface{}) *Each {
    a.Set("hidden", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Each) StaticPlaceholder(value interface{}) *Each {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *Each) Testid(value interface{}) *Each {
    a.Set("testid", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Each) Disabled(value interface{}) *Each {
    a.Set("disabled", value)
    return a
}

/**
 */
func (a *Each) StaticSchema(value interface{}) *Each {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *Each) TestIdBuilder(value interface{}) *Each {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 关联字段名 支持数据映射
 */
func (a *Each) Source(value interface{}) *Each {
    a.Set("source", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Each) VisibleOn(value interface{}) *Each {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Each) Static(value interface{}) *Each {
    a.Set("static", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Each) StaticOn(value interface{}) *Each {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *Each) Placeholder(value interface{}) *Each {
    a.Set("placeholder", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Each) OnEvent(value interface{}) *Each {
    a.Set("onEvent", value)
    return a
}

/**
 * 指定为each展示类型
 */
func (a *Each) Type(value interface{}) *Each {
    a.Set("type", value)
    return a
}

/**
 * 关联字段名
 */
func (a *Each) Name(value interface{}) *Each {
    a.Set("name", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Each) HiddenOn(value interface{}) *Each {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Each) StaticInputClassName(value interface{}) *Each {
    a.Set("staticInputClassName", value)
    return a
}
