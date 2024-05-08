package register

import "fmt"

type Handler struct {
}

func (h Handler) Handle() error {
	fmt.Println("register account asf")
	return nil
}
