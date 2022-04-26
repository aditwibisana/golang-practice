package main

import (
	"fmt"
)

// Check Point:
// Search Roles
// - Input: Role
// - Output: Rearch Result

// Contoh:
// Data Users: [{Aditira 20 Programmer} {Dimas 18 Designer} {Yuni 21 DevOps} {Dito 22 Programmer} {Juno 25 DevOps}]

// Input:
//   - Masukan Role : Programmer

// Output:
// Programmer Found:
// Name:  Aditira  Age:  20  Role:  Programmer
// Name:  Dito  Age:  22  Role:  Programmer

// Input:
//   - Masukan Role : Secretary

// Output:
// Role: Sercretary Not Found!

type User struct {
	name string
	age  int
	role string
}

func main() {

	users := []User{
		{
			name: "Aditira",
			age:  20,
			role: "Programmer",
		},
		{
			name: "Dimas",
			age:  18,
			role: "Designer",
		},
		{
			name: "Yuni",
			age:  21,
			role: "DevOps",
		},
		{
			name: "Dito",
			age:  22,
			role: "Programmer",
		},
		{
			name: "Juno",
			age:  25,
			role: "DevOps",
		},
	}

<<<<<<< HEAD

	// fmt.Println(len(users))
	// fmt.Println(users[1].role)
	//MYANSWER
	// var searchRole string

	// fmt.Print("Masukkan Role: ")
	// fmt.Scan(&searchRole)
	// for i:=0 ; i<len(users); i++{
	// 	if (searchRole == users[i].role ){
	// 		fmt.Println(searchRole, "Found!")
	// 		fmt.Println("Nama: ",users[i].name,
	// 					" Age: ", users[i].age,
	// 					" Role: ", users[i].role)
	// 	}else{
	// 		fmt.Println(searchRole, "Not Found!")
	// 	}
	// }
	//ENDSHERE
	// TODO: answer here

	//beginanswer
	var searchResult []User
	var role string
	fmt.Printf("Masukan Role : ")
	fmt.Scan(&role)

	for _, user := range users {
		if user.role == role {
			searchResult = append(searchResult, user)
		}
	}

	if len(searchResult) != 0 {
		fmt.Printf("%v Found: ", role)
		for _, u := range searchResult {
			fmt.Println("Name: ", u.name, " Age: ", u.age, " Role: ", u.role)
		}
	} else {
		fmt.Printf("Role: %v Not Found!", role)
	}
	//endanswer

=======
	// TODO: answer here
>>>>>>> 2ae988a7ab8fa2ea4f367e7c567e39c5bce0995c
}
