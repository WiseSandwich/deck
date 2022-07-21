package subpackage

import "fmt"

func privateFunction(name string) {
	fmt.Println("This is the name you receive")
}

func PublicFunction(name string) {
	privateFunction(name)
}
