package renderers


/**

* @author wcz0
* @version 6.2.2
*/
type Avatar struct {
	*BaseRenderer
}

func NewAvatar() *Avatar {
    a := &Avatar{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("crossOrigin", "anonymous")
    a.Set("type", "avatar")
    return a
}
/**
 * 图片无法显示时的替换文字地址
 */
func (a *Avatar) Alt(value string) *Avatar {
    a.Set("alt", value)
    return a
}

/**
 * 图片加载失败的是否默认处理，字符串函数
 */
func (a *Avatar) OnError(value string) *Avatar {
    a.Set("onError", value)
    return a
}

/**
 * 类名
 */
func (a *Avatar) ClassName(value string) *Avatar {
    a.Set("className", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Avatar) UseMobileUI(value string) *Avatar {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 图标
 */
func (a *Avatar) Icon(value string) *Avatar {
    a.Set("icon", value)
    return a
}

/**
 * 形状
 * 可选值: circle | square | rounded
 */
func (a *Avatar) Shape(value string) *Avatar {
    a.Set("shape", value)
    return a
}

/**
 * 图片CORS属性
 * 可选值: anonymous | use-credentials | 
 */
func (a *Avatar) CrossOrigin(value string) *Avatar {
    a.Set("crossOrigin", value)
    return a
}

/**
 * 文本
 */
func (a *Avatar) Text(value string) *Avatar {
    a.Set("text", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Avatar) Hidden(value string) *Avatar {
    a.Set("hidden", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Avatar) StaticLabelClassName(value string) *Avatar {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *Avatar) StaticSchema(value string) *Avatar {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *Avatar) Type(value string) *Avatar {
    a.Set("type", value)
    return a
}

/**
 * 图片地址
 */
func (a *Avatar) Src(value string) *Avatar {
    a.Set("src", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Avatar) HiddenOn(value string) *Avatar {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Avatar) StaticPlaceholder(value string) *Avatar {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 自定义样式
 */
func (a *Avatar) Style(value string) *Avatar {
    a.Set("style", value)
    return a
}

/**
 * 是否显示
 */
func (a *Avatar) Visible(value string) *Avatar {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Avatar) StaticInputClassName(value string) *Avatar {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 大小
 */
func (a *Avatar) Size(value string) *Avatar {
    a.Set("size", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Avatar) VisibleOn(value string) *Avatar {
    a.Set("visibleOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Avatar) EditorSetting(value string) *Avatar {
    a.Set("editorSetting", value)
    return a
}

/**
 * 字符类型距离左右两侧边界单位像素
 */
func (a *Avatar) Gap(value string) *Avatar {
    a.Set("gap", value)
    return a
}

/**
 * 图片是否允许拖动
 */
func (a *Avatar) Draggable(value string) *Avatar {
    a.Set("draggable", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Avatar) Disabled(value string) *Avatar {
    a.Set("disabled", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Avatar) StaticOn(value string) *Avatar {
    a.Set("staticOn", value)
    return a
}

/**
 * 角标
 */
func (a *Avatar) Badge(value string) *Avatar {
    a.Set("badge", value)
    return a
}

/**
 * 默认头像
 */
func (a *Avatar) DefaultAvatar(value string) *Avatar {
    a.Set("defaultAvatar", value)
    return a
}

/**
 * 图片相对于容器的缩放方式
 * 可选值: fill | contain | cover | none | scale-down
 */
func (a *Avatar) Fit(value string) *Avatar {
    a.Set("fit", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Avatar) DisabledOn(value string) *Avatar {
    a.Set("disabledOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Avatar) Id(value string) *Avatar {
    a.Set("id", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Avatar) OnEvent(value string) *Avatar {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Avatar) Static(value string) *Avatar {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Avatar) StaticClassName(value string) *Avatar {
    a.Set("staticClassName", value)
    return a
}
