package main

import "fmt"

func main() {

	name := getName("hzb")
	fmt.Println(name)

	_, age := getNameAndAge("hzb", 24)
	fmt.Println(age)

}

func getName(name string) string {
	return name
}

func getNameAndAge(name string, age int) (string, int) {
	return name, age
}
