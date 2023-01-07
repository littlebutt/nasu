package context

type Context struct {
	Password     string
	ResourcesDir string
}

var NasuContext = Context{}
