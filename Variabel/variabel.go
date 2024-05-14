package main

import "fmt"

func main() {
	var firstName string = "Arif"
	var centerName string = "Putera"
	var lastName string = "Wijaya"

	fmt.Println(firstName, centerName, lastName)
	fmt.Printf("Halo, Arif Putera Wijaya\n")
	fmt.Printf("Halo, %s %s %s \n", firstName, centerName, lastName)

	var (
		age     int
		address string
	)

	age = 22
	address = "Kota Singkawang, Kalimantan Barat"

	fmt.Printf("Halo,saya %s %s %s berumur %d dan saya tinggal di %s \n", firstName, centerName, lastName, age, address)

}
