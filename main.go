package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5, 6}
	arr = append(arr[:2], arr[3:]...)
	fmt.Println(arr)

}
