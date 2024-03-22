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
