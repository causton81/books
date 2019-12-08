package books

func Must(err error) {
	if nil != err {
		panic(err)
	}
}
