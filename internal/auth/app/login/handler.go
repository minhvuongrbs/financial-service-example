package login

import "fmt"

type Handler struct {
}

func (h Handler) Handle() error {
	fmt.Println("login account asf")
	return nil
}
