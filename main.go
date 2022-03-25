package main

import (
	"fmt"
	"github.com/mohsen2hasani/Go-GettingStarted/models"
)

func main() {
	fmt.Println("Hi !")
	fmt.Println("---------------------------------------------------------")

	fmt.Println("varTest")
	varTest()
	fmt.Println("---------------------------------------------------------")

	fmt.Println("arrayTest")
	arrayTest()
	fmt.Println("---------------------------------------------------------")

	fmt.Println("sliceTest")
	sliceTest()
	fmt.Println("---------------------------------------------------------")

	fmt.Println("mapTest")
	mapTest()
	fmt.Println("---------------------------------------------------------")

	fmt.Println("structTest")
	structTest()
	fmt.Println("---------------------------------------------------------")

	fmt.Println("packageTest")
	packageTest()
	fmt.Println("---------------------------------------------------------")

	fmt.Println("loopTest1")
	loopTest1()
	fmt.Println("---------------------------------------------------------")

	fmt.Println("loopTest2")
	loopTest2()
	fmt.Println("---------------------------------------------------------")

	fmt.Println("loopTest3")
	loopTest3()
	fmt.Println("---------------------------------------------------------")

	fmt.Println("loopTest4")
	loopTest4()
	fmt.Println("---------------------------------------------------------")

	fmt.Println("loopTest5")
	loopTest5()
	fmt.Println("---------------------------------------------------------")

	fmt.Println("loopTest6")
	loopTest6()
	fmt.Println("---------------------------------------------------------")

	fmt.Println("loopTest7")
	loopTest7()
	fmt.Println("---------------------------------------------------------")

	fmt.Println("loopTest8")
	loopTest8()
	fmt.Println("---------------------------------------------------------")

	fmt.Println("switchTest")
	switchTest()
	fmt.Println("---------------------------------------------------------")
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

func packageTest() {
	u := models.User{
		Id:        5,
		FirstName: "Mohsen",
		LastName:  "Hasani",
	}

	fmt.Println(u)
}

func loopTest1() {
	var i int
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

func loopTest2() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func loopTest3() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)

		if i == 3 {
			break
		}
	}
}

func loopTest4() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)

		if i == 3 {
			continue
		}
		fmt.Println("continuing...")
	}
}

func loopTest5() {
	var i int
	for {
		if i == 5 {
			break
		}
		fmt.Println(i)
		i++
	}
}

func loopTest6() {
	slice := []int{5, 8, 15}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}

func loopTest7() {
	slice := []int{5, 8, 15}
	for i, v := range slice {
		fmt.Println(i, v)
	}
}

func loopTest8() {
	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for k, v := range wellKnownPorts {
		fmt.Println(k, v)
	}
}

func switchTest() {
	type HTTPRequest struct {
		Method string
	}

	r := HTTPRequest{Method: "GET"}

	switch r.Method {
	case "GET":
		fmt.Println("Get request")
		//fallthrough => continue to next case
	case "DELETE":
		fmt.Println("Delete request")
	case "POST":
		fmt.Println("Post request")
	case "PUT":
		fmt.Println("Put request")
	default:
		fmt.Println("Default case")
	}
}
