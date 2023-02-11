package main

import "fmt"

type Live interface {
	Eat(name string)
}

type Person struct {
	Name string
}

/*结构体实现指针*/
func (p Person) Eat(name string) {
	fmt.Printf("Eat %s\n", name)
}

func main() {

	/*指针*/
	person := &Person{}
	fmt.Printf(person.Name)

	/*结构体实例*/
	personTwo := Person{}
	fmt.Printf(personTwo.Name)

	/*指针*/
	personThree := new(Person)
	fmt.Printf(personThree.Name)

	/*结构体实例*/
	var personFour Person
	fmt.Printf(personFour.Name)

	/*指针声明，但是没有指向*/
	var personFive *Person
	if personFive == nil {
		fmt.Printf("personFive is nil\n")
	}

	/*结构体实例*/
	personSix := Person{Name: "hzb"}
	fmt.Printf("%s\n", personSix.Name)
	personSix.Name = "hzb again"
	fmt.Printf("%s\n", personSix.Name)

	personSix.Eat("apple")
}
