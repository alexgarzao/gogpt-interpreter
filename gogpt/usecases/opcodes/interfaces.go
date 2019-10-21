package opcodes

type StdoutInterface interface {
	Println(text interface{})
	Readln() string
}
