package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	person1 := Person{Age: 25,
		Name: "manidhar",
	}

	fmt.Printf("Person details %+v\n", person1)

	//annonymous struct

	employe := struct {
		name string
		id   int
	}{
		name: "Praveen", id: 420}

	fmt.Printf("Person details %+v\n", employe)

	// nested struct

	type Address struct {
		Street string
		City   string
	}

	type Contact struct {
		Person          // embadded
		address Address //nested
		Number  string
	}

	address := Address{
		"vr colony",
		"kurnool",
	}

	contact1 := Contact{
		person1,
		Address{"Gandi nagar",
			"kurnool"},
		"56732",
	}

	fmt.Printf("Contact details %+v\n", contact1)

	contact2 := Contact{
		Person{
			Name: "Ravi",
		},
		address,
		"56732",
	}

	/*%v prints the struct values, while %+v prints the struct values with their corresponding field names, providing more explicit output about the struct's contents*/

	fmt.Printf("Contact details %+v\n", contact2)

	fmt.Println(contact1.address)
	fmt.Println(contact2.Name)

}
