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
/**
 * 对应渲染器类型
 */
func (a *Expandable) Type(value string) *Expandable {
    a.Set("type", value)
    return a
}

/**
 * 对应数据源的key值
 */
func (a *Expandable) KeyField(value string) *Expandable {
    a.Set("keyField", value)
    return a
}

/**
 * 行是否可展开表达式
 */
func (a *Expandable) ExpandableOn(value string) *Expandable {
    a.Set("expandableOn", value)
    return a
}

/**
 * 展开行自定义样式表达式
 */
func (a *Expandable) ExpandedRowClassNameExpr(value string) *Expandable {
    a.Set("expandedRowClassNameExpr", value)
    return a
}

/**
 * 已展开的key值
 */
func (a *Expandable) ExpandedRowKeys(value string) *Expandable {
    a.Set("expandedRowKeys", value)
    return a
}

/**
 * 已展开的key值表达式
 */
func (a *Expandable) ExpandedRowKeysExpr(value string) *Expandable {
    a.Set("expandedRowKeysExpr", value)
    return a
}
