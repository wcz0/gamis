package renderers


/**

* @author wcz0
* @version 6.2.2
*/
type IconItem struct {
	*BaseRenderer
}

func NewIconItem() *IconItem {
    a := &IconItem{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}
/**
 */
func (a *IconItem) Icon(value string) *IconItem {
    a.Set("icon", value)
    return a
}

/**
 */
func (a *IconItem) Position(value string) *IconItem {
    a.Set("position", value)
    return a
}
