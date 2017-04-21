package main

type myError struct {
	Msg string
}

func (m *myError) Error() string {
	return m.Msg
}

func raiseErrorWhenBadThingsHappens(bad bool) error {
	var me *myError
	if bad {
		me = &myError{Msg: "this is bad!"}
	}
	return me
}
