package main
import ("fmt")

func print(num int){
	fmt.Println(num)
}

func printfibonacci(num int) {
	var first int = 0
	var second int = 1
	var temp int;

	fmt.Print(first, " ")
	fmt.Print(second, " ")


	for i:= 0; i < num - 2; i++ {
		temp = first + second
		first = second
		second = temp
		fmt.Print(temp, " ")
	}
}

func main(){
	printfibonacci(10)
}