package context

type Context struct {
	Password     string
	ResourcesDir string
	TokenTTL     int64 // hour(s)
}

var NasuContext = Context{}
