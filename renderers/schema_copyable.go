package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type SchemaCopyable struct {
	*BaseRenderer
}

func NewSchemaCopyable() *SchemaCopyable {
    a := &SchemaCopyable{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}


func (a *SchemaCopyable) Set(name string, value interface{}) *SchemaCopyable {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 配置复制时的内容模板。
 */
func (a *SchemaCopyable) Content(value interface{}) *SchemaCopyable {
    a.Set("content", value)
    return a
}

/**
 * 提示文字内容
 */
func (a *SchemaCopyable) Tooltip(value interface{}) *SchemaCopyable {
    a.Set("tooltip", value)
    return a
}

/**
 * 可以配置图标
 */
func (a *SchemaCopyable) Icon(value interface{}) *SchemaCopyable {
    a.Set("icon", value)
    return a
}
