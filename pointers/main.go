package main

type thing struct {
	value string
}

func (t thing) SetValue(val string) {
	t.value = val
}
