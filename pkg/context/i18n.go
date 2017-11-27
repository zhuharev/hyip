package context

func (c *Context) Tr(fmt string, args ...interface{}) string {
	c.Locale.Lang = c.User.LangString()
	return c.Locale.Tr(fmt, args...)
}
