package main
import ("fmt")

func main() {
	new_slice := []int{4, 8, 10, 20}
	new_slice = append(new_slice, 30, 40, 50)
	fmt.Println(new_slice)

	//find length of slice
	fmt.Println(len(new_slice))

	//find capacity of slice
	fmt.Println(cap(new_slice))
}
