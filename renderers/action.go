package renderers

type Action struct {
	*BaseRenderer
}

func NewAction() *Action {
	a := &Action{
		BaseRenderer: NewBaseRenderer(),
	}
	a.Set("type", "action")
	return a
}

/**
 * 【必填】这是 action 最核心的配置，来指定该 action 的作用类型，支持：ajax、link、url、drawer、dialog、confirm、cancel、prev、next、copy、close。
 */
func (a *Action) ActionType(value string) *Action {
	a.Set("actionType", value)
	return a
}

/**
 * 给按钮高亮添加类名。
 */
func (a *Action) ActiveClassName(value string) *Action {
	a.Set("activeClassName", value)
	return a
}

/**
 * 按钮高亮时的样式，配置支持同level。
 */
func (a *Action) ActiveLevel(value string) *Action {
	a.Set("activeLevel", value)
	return a
}

func (a *Action) Api(value string) *Action {
	a.Set("api", value)
	return a
}

/**
 * 添加类名。
 */
func (a *Action) ClassName(value string) *Action {
	a.Set("className", value)
	return a
}

/**
 * 当action配置在dialog或drawer的actions中时，配置为true指定此次操作完后关闭当前dialog或drawer。当值为字符串，并且是祖先层弹框的名字的时候，会把祖先弹框关闭掉。
 */
func (a *Action) Close(value string) *Action {
	a.Set("close", value)
	return a
}

/**
 * 当设置后，操作在开始前会询问用户。可用 '$[xxx]' 取值。
 */
func (a *Action) ConfirmText(value string) *Action {
	a.Set("confirmText", value)
	return a
}

/**
 * 被禁用后鼠标停留时弹出该段文字，也可以配置对象类型：字段为title和content。可用 '$[xxx]' 取值。
 */
func (a *Action) DisabledTip(value string) *Action {
	a.Set("disabledTip", value)
	return a
}

/**
 * 设置图标，例如fa fa-plus。
 */
func (a *Action) Icon(value string) *Action {
	a.Set("icon", value)
	return a
}

/**
 * 给图标上添加类名。
 */
func (a *Action) IconClassName(value string) *Action {
	a.Set("iconClassName", value)
	return a
}

/**
 * 按钮文本。可用 '$[xxx]' 取值。
 */
func (a *Action) Label(value string) *Action {
	a.Set("label", value)
	return a
}

/**
 * 按钮样式，支持：link /primary/secondary/info/success/warning/danger/light/dark/default。
 */
func (a *Action) Level(value string) *Action {
	a.Set("level", value)
	return a
}

func (a *Action) Link(value string) *Action {
	a.Set("link", value)
	return a
}

/**
 * 指定此次操作完后，需要刷新的目标组件名字（组件的name值，自己配置的），多个请用, 号隔开。
 */
func (a *Action) Reload(value string) *Action {
	a.Set("reload", value)
	return a
}

/**
 * 配置字符串数组，指定在form中进行操作之前，需要指定的字段名的表单项通过验证
 */
func (a *Action) Required(value string) *Action {
	a.Set("required", value)
	return a
}

/**
 * 在按钮文本右侧设置图标，例如fa fa-plus。
 */
func (a *Action) RightIcon(value string) *Action {
	a.Set("rightIcon", value)
	return a
}

/**
 * 给右侧图标上添加类名。
 */
func (a *Action) RightIconClassName(value string) *Action {
	a.Set("rightIconClassName", value)
	return a
}

/**
 * 按钮大小，支持：xs、sm、md、lg。
 */
func (a *Action) Size(value string) *Action {
	a.Set("size", value)
	return a
}

/**
 * 鼠标停留时弹出该段文字，也可以配置对象类型：字段为title和content。可用 '$[xxx]' 取值。
 */
func (a *Action) Tooltip(value string) *Action {
	a.Set("tooltip", value)
	return a
}

/**
 * 如果配置了tooltip或者disabledTip，指定提示信息位置，可配置top、bottom、left、right。
 */
func (a *Action) TooltipPlacement(value string) *Action {
	a.Set("tooltipPlacement", value)
	return a
}

/**
 * 指定为 action 渲染器。
 */
func (a *Action) Type(value string) *Action {
	a.Set("type", value)
	return a
}
