package console

type Console interface {
	Printf(format string, args ...interface{})
	ReadLine() string
}
