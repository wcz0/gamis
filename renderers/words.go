package renderers


/**
 * Words
 *

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
func (a *Words) OnEvent(value string) *Words {
    a.Set("onEvent", value)
    return a
}

/**
 * 展示文字
 */
func (a *Words) ExpendButtonText(value string) *Words {
    a.Set("expendButtonText", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Words) ClassName(value string) *Words {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Words) Disabled(value string) *Words {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Words) Hidden(value string) *Words {
    a.Set("hidden", value)
    return a
}

/**
 * 收起文字
 */
func (a *Words) CollapseButtonText(value string) *Words {
    a.Set("collapseButtonText", value)
    return a
}

/**
 * 分割符
 */
func (a *Words) Delimiter(value string) *Words {
    a.Set("delimiter", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Words) DisabledOn(value string) *Words {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Words) HiddenOn(value string) *Words {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Words) Static(value string) *Words {
    a.Set("static", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Words) UseMobileUI(value string) *Words {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Words) VisibleOn(value string) *Words {
    a.Set("visibleOn", value)
    return a
}

/**
 * 展示限制, 为0时也无限制
 */
func (a *Words) Limit(value string) *Words {
    a.Set("limit", value)
    return a
}

/**
 * 展示文字
 */
func (a *Words) CollapseButton(value string) *Words {
    a.Set("collapseButton", value)
    return a
}

/**
 * 是否显示
 */
func (a *Words) Visible(value string) *Words {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Words) StaticClassName(value string) *Words {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Words) StaticPlaceholder(value string) *Words {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Words) StaticLabelClassName(value string) *Words {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Words) StaticInputClassName(value string) *Words {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *Words) StaticSchema(value string) *Words {
    a.Set("staticSchema", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Words) EditorSetting(value string) *Words {
    a.Set("editorSetting", value)
    return a
}

/**
 * 展示文字
 */
func (a *Words) ExpendButton(value string) *Words {
    a.Set("expendButton", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Words) Id(value string) *Words {
    a.Set("id", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Words) StaticOn(value string) *Words {
    a.Set("staticOn", value)
    return a
}

/**
 * useTag 当数据是数组时，是否使用tag的方式展示
 */
func (a *Words) InTag(value string) *Words {
    a.Set("inTag", value)
    return a
}

/**
 * tags数据
 */
func (a *Words) Words(value string) *Words {
    a.Set("words", value)
    return a
}

/**
 * 组件样式
 */
func (a *Words) Style(value string) *Words {
    a.Set("style", value)
    return a
}

/**
 */
func (a *Words) Type(value string) *Words {
    a.Set("type", value)
    return a
}
