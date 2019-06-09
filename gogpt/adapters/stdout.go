package adapters

import "fmt"

type Stdout struct {
}

func NewStdout() *Stdout {
	return &Stdout{}
}

func (s *Stdout) Println(text interface{}) {
	fmt.Printf("%v\n", text)
}
