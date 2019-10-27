package opcodes

type StdinInterface interface {
	Readln() string
}

type StdoutInterface interface {
	Println(text interface{})
}
