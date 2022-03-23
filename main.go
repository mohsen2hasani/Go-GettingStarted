package main

import (
	"fmt"
	"github.com/mohsen2hasani/Go-GettingStarted/models"
)

func main() {
	fmt.Println("Hi !")

	varTest()

	arrayTest()

	sliceTest()

	mapTest()

	structTest()

	u := models.User{
		Id:        5,
		FirstName: "Mohsen",
		LastName:  "Hasani",
	}

	fmt.Println(u)
}

func varTest() {
	var intVar int
	intVar = 200
	fmt.Println("var: ", intVar)

	//var 2
	intVar2 := 200
	fmt.Println("var 2: ", intVar2)
}

func arrayTest() {
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println("array: ", arr)

	//array 2
	arr2 := [3]int{1, 2, 3}
	fmt.Println("array2: ", arr2)
}

func sliceTest() {
	arr := [3]int{1, 2, 3}

	slice := arr[:]
	fmt.Println("slice: ", slice)

	arr[0] = 52
	slice[1] = 20
	fmt.Println("slice: ", slice)

	//slice 2
	slice2 := []int{1, 2, 3}
	slice2 = append(slice2, 4)
	fmt.Println("slice2: ", slice2)

	s1 := slice2[1:]
	s2 := slice2[:2]
	s3 := slice2[1:2]

	fmt.Println("slice3: ", s1, s2, s3)
}

func mapTest() {
	m := map[string]int{"foo": 42, "foo2": 22}
	fmt.Println("map: ", m)

	m["foo"] = 200
	fmt.Println("map: ", m)

	delete(m, "foo")
	fmt.Println("map: ", m)
}

func structTest() {
	type user struct {
		Id        int
		FirstName string
		LastName  string
	}

	var u user
	u.Id = 1
	u.FirstName = "Mohsen"
	u.LastName = "Hasani"

	fmt.Println(u)

	u2 := user{Id: 1,
		FirstName: "Mohsen",
		LastName:  "Hasani",
	}

	fmt.Println(u2)
}
