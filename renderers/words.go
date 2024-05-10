package renderers


/**
 * Words

 * @author wcz0
 * @version 6.2.2
 */
type Words struct {
	*BaseRenderer
}

func NewWords() *Words {
    a := &Words{
        BaseRenderer: NewBaseRenderer(),
    }
    a.Set("type", "words")
    return a
}
/**
 * 事件动作配置
 */
func (a *Words) OnEvent(value interface{}) *Words {
    a.Set("onEvent", value)
    return a
}

/**
 */
func (a *Words) StaticSchema(value interface{}) *Words {
    a.Set("staticSchema", value)
    return a
}

/**
 * 分割符
 */
func (a *Words) Delimiter(value interface{}) *Words {
    a.Set("delimiter", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Words) StaticPlaceholder(value interface{}) *Words {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Words) StaticClassName(value interface{}) *Words {
    a.Set("staticClassName", value)
    return a
}

/**
 * 展示文字
 */
func (a *Words) CollapseButton(value interface{}) *Words {
    a.Set("collapseButton", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Words) StaticLabelClassName(value interface{}) *Words {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *Words) Style(value interface{}) *Words {
    a.Set("style", value)
    return a
}

/**
 * 展示文字
 */
func (a *Words) ExpendButton(value interface{}) *Words {
    a.Set("expendButton", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Words) ClassName(value interface{}) *Words {
    a.Set("className", value)
    return a
}

/**
 * 展示文字
 */
func (a *Words) ExpendButtonText(value interface{}) *Words {
    a.Set("expendButtonText", value)
    return a
}

/**
 * 收起文字
 */
func (a *Words) CollapseButtonText(value interface{}) *Words {
    a.Set("collapseButtonText", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Words) DisabledOn(value interface{}) *Words {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *Words) Visible(value interface{}) *Words {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Words) StaticInputClassName(value interface{}) *Words {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Words) Disabled(value interface{}) *Words {
    a.Set("disabled", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Words) VisibleOn(value interface{}) *Words {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Words) StaticOn(value interface{}) *Words {
    a.Set("staticOn", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Words) UseMobileUI(value interface{}) *Words {
    a.Set("useMobileUI", value)
    return a
}

/**
 * useTag 当数据是数组时，是否使用tag的方式展示
 */
func (a *Words) InTag(value interface{}) *Words {
    a.Set("inTag", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Words) Id(value interface{}) *Words {
    a.Set("id", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Words) EditorSetting(value interface{}) *Words {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *Words) Type(value interface{}) *Words {
    a.Set("type", value)
    return a
}

/**
 * 展示限制, 为0时也无限制
 */
func (a *Words) Limit(value interface{}) *Words {
    a.Set("limit", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Words) Hidden(value interface{}) *Words {
    a.Set("hidden", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Words) HiddenOn(value interface{}) *Words {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Words) Static(value interface{}) *Words {
    a.Set("static", value)
    return a
}

/**
 * tags数据
 */
func (a *Words) Words(value interface{}) *Words {
    a.Set("words", value)
    return a
}
