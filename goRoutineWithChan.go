package main
// m00ky 20 
import (
  "fmt"
)
// Create a custom struct type
type MyType struct {
	Name string
}

// Declare a method on the type to update the Name attr
func (t *MyType) updateName(ch chan string) {
	t.Name = <-ch
}

func main() {
	// Create a channel
	var ch = make(chan string)
	// Create an instance of our type
	a := MyType{Name: "emma"}
	fmt.Println(a.Name)
	// Create a thread with a closure 
	go func() {
		// Call the method of our custom type in the thread
		a.updateName(ch)
	}()
	// Pass in a new value for the Name 
	ch <- "tina"
	/* 
	  Show that the change of name in the thread is 
	  referencing our custom type in the parent 
	*/
	fmt.Println(a.Name)
}
