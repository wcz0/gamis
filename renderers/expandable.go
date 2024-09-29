package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type Expandable struct {
	*BaseRenderer
}

func NewExpandable() *Expandable {
    a := &Expandable{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}


func (a *Expandable) Set(name string, value interface{}) *Expandable {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 对应渲染器类型
 */
func (a *Expandable) Type(value interface{}) *Expandable {
    a.Set("type", value)
    return a
}

/**
 * 对应数据源的key值
 */
func (a *Expandable) KeyField(value interface{}) *Expandable {
    a.Set("keyField", value)
    return a
}

/**
 * 行是否可展开表达式
 */
func (a *Expandable) ExpandableOn(value interface{}) *Expandable {
    a.Set("expandableOn", value)
    return a
}

/**
 * 展开行自定义样式表达式
 */
func (a *Expandable) ExpandedRowClassNameExpr(value interface{}) *Expandable {
    a.Set("expandedRowClassNameExpr", value)
    return a
}

/**
 * 已展开的key值
 */
func (a *Expandable) ExpandedRowKeys(value interface{}) *Expandable {
    a.Set("expandedRowKeys", value)
    return a
}

/**
 * 已展开的key值表达式
 */
func (a *Expandable) ExpandedRowKeysExpr(value interface{}) *Expandable {
    a.Set("expandedRowKeysExpr", value)
    return a
}
