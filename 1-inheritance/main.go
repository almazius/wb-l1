package main

import "fmt"

// Human is parent struct
type Human struct {
	Name string
	age  uint
}

// GetAge is his method
func (h *Human) GetAge() uint {
	return h.age
}

// Action is inheritor Human
type Action struct {
	Human
}

func main() {
	// create inheritor
	action := new(Action)

	// use fields of the parent
	action.Name = "Almaz"
	action.age = 10

	// use method of the parents
	fmt.Println(action.GetAge())

}
