package main

import "fmt"

func main()  {
	var slice []int
	var cantidad, num int
	total := 0

	fmt.Scanln(&cantidad)

	for i := 0; i < cantidad; i++ {
		fmt.Scanln(&num)
		slice = append(slice, num)
	}

	for _, dato := range slice {
		total += dato
	}

	fmt.Println(total)

}