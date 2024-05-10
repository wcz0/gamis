package renderers


/**
 * 消息文案配置，记住这个优先级是最低的，如果你的接口返回了 msg，接口返回的优先。

 * @author wcz0
 * @version 6.2.2
 */
type SchemaMessage struct {
	*BaseRenderer
}

func NewSchemaMessage() *SchemaMessage {
    a := &SchemaMessage{
        BaseRenderer: NewBaseRenderer(),
    }
    return a
}
/**
 * 获取失败时的提示
 */
func (a *SchemaMessage) FetchFailed(value interface{}) *SchemaMessage {
    a.Set("fetchFailed", value)
    return a
}

/**
 * 获取成功的提示，默认为空。
 */
func (a *SchemaMessage) FetchSuccess(value interface{}) *SchemaMessage {
    a.Set("fetchSuccess", value)
    return a
}

/**
 * 保存失败时的提示。
 */
func (a *SchemaMessage) SaveFailed(value interface{}) *SchemaMessage {
    a.Set("saveFailed", value)
    return a
}

/**
 * 保存成功时的提示。
 */
func (a *SchemaMessage) SaveSuccess(value interface{}) *SchemaMessage {
    a.Set("saveSuccess", value)
    return a
}
