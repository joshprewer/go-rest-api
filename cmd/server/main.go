package main

import "fmt"

// Setup app layers
func Run() error {
	fmt.Println("starting app")
	return nil
}

func main() {
	fmt.Println("Go Rest API Course")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
