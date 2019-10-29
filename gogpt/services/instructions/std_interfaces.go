package instructions

// StdinInterface has the interface necessary to a stdin.
type StdinInterface interface {
	Readln() string
}

// StdoutInterface has the interface necessary to a stdout.
type StdoutInterface interface {
	Println(text interface{})
}
