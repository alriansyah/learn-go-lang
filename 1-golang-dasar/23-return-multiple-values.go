package main

import "fmt"

// Function Return Multiple Values
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g74a60ad919_0_1098#slide=id.g74a60ad919_0_1098

func getFullName() (string, string) {
	return "Eko", "Khannedy"
}

func main() {
	//firstName, lastName := getFullName()
	//fmt.Println(firstName, lastName)

	// menghiraukan nilai return dengan _, bisa dibilang meng ignore nilai return
	firstName, _ := getFullName()
	fmt.Println(firstName)
}
