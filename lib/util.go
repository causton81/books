package lib

func Must(err error) {
	if nil != err {
		panic(err)
	}
}
