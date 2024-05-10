package renderers

type WebComponent struct {
	*BaseRenderer
}

func NewWebComponent() *WebComponent {
	w := &WebComponent{
		BaseRenderer: NewBaseRenderer(),
	}
	w.Set("type", "web_component")
	return w
}

func (w *WebComponent) body(value interface{}) *WebComponent {
	w.Set("body", value)
	return w
}

func (w *WebComponent) Props(value interface{}) *WebComponent {
	w.Set("props", value)
	return w
}

func (w *WebComponent) Tag(value interface{}) *WebComponent {
	w.Set("tag", value)
	return w
}

func (w *WebComponent) Type(value interface{}) *WebComponent {
	w.Set("type", value)
	return w
}
