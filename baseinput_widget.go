package gforms

func BaseTextWidget(typ string, attrs map[string]string) Widget {
	w := new(textInputWidget)
	w.Type = typ
	if attrs == nil {
		attrs = map[string]string{}
	}
	w.Attrs = attrs
	return w
}
