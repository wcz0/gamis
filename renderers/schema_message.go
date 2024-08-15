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

func (a *SchemaMessage) Set(name string, value interface{}) *SchemaMessage {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}

    return a
}

/**
 * 获取成功的提示，默认为空。
 */
func (a *SchemaMessage) Fetchsuccess(value interface{}) *SchemaMessage {
    a.Set("fetchSuccess", value)
    return a
}

/**
 * 保存失败时的提示。
 */
func (a *SchemaMessage) Savefailed(value interface{}) *SchemaMessage {
    a.Set("saveFailed", value)
    return a
}

/**
 * 保存成功时的提示。
 */
func (a *SchemaMessage) Savesuccess(value interface{}) *SchemaMessage {
    a.Set("saveSuccess", value)
    return a
}

/**
 * 获取失败时的提示
 */
func (a *SchemaMessage) Fetchfailed(value interface{}) *SchemaMessage {
    a.Set("fetchFailed", value)
    return a
}
