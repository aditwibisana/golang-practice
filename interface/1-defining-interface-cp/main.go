package main

import "fmt"

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk membuat interface dan mengimplementasikan interface.
// Buatlah interface Employee yang memiliki method signature GetBonus() int
// Buatlah implementasi interface Employee pada objek Manager, SeniorEngineer, dan JuniorEngineer.
// Pada objek Manager, SeniorEngineer, dan JuniorEngineer memiliki satu atribut yaitu BaseSalary.
// Untuk menghitung bonus terdapat tiga aturan sebagai berikut:
// Bonus untuk Manager adalah 3 * BaseSalary
// Bonus untuk SeniorEngineer adalah 2 * BaseSalary
// Bonus untuk JuniorEngineer adalah 1 * BaseSalary

// TODO: answer here

//interface
type Employee interface{
	GetBonus() int
}

//objek
type Manager struct{
	BaseSalary int
}

func(m Manager) GetBonus() int{
	return 3 * m.BaseSalary
}

//objek
type SeniorEngineer struct{
	BaseSalary int
}

func(s SeniorEngineer) GetBonus()int{
	return 2 * s.BaseSalary
}

//objek
type JuniorEngineer struct{
	BaseSalary int
}

func(j JuniorEngineer) GetBonus() int{
	return j.BaseSalary
}

func TotalEmployeeBonus(employees []Employee) int {
	// Hitunglah total bonus yang dikeluarkan dari list of Employee
	// TODO: answer here
	total := 0
	for _, value := range employees{
		total += value.GetBonus()
	}
	return total
}

func main() {
	// Buatlah objek konkret untuk masing-masing objek dan panggil function TotalEmployeeBonus. Print total bonus untuk semua employee.
	// TODO: answer here
	manager := Manager{
		BaseSalary: 20000000,
	}

	seniorEngineer := SeniorEngineer{
		BaseSalary: 15000000,
	}

	juniorEngineer := JuniorEngineer{
		BaseSalary: 10000000,
	}

	totalBonus := TotalEmployeeBonus([]Employee{manager, seniorEngineer, juniorEngineer})

	fmt.Println("total bonus", totalBonus)
}
