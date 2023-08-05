// package main
// import "fmt"

// type Client struct {
// 	Nome string
// 	Email string
// }

// func main()  {
// 	client := Client{
// 		Nome: "Thiago",
// 		Email: "thiago123@gmail.com",
// 	}

// 	fmt.Println(client.Nome)
// 	fmt.Println(client.Email)
// }

package main
import "fmt"

type Animal interface{
	Sound() string
	Move() string
}

type Dog struct{
	Name string
}

func (d Dog) Sound() string {
	return d.Name + " Say, woof"
}

func (d Dog) Move() string {
	return d.Name + " Running"
}

func PerformActions(a Animal) {
	fmt.Println(a.Sound())
	fmt.Println(a.Move())
}

func main() {
	dog := Dog{
		Name: "Spike",
	}
	PerformActions(dog)
}

