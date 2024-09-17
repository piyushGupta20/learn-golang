package main
import ("fmt")

type Employee struct {
	id int
	name string
	age int
	company string
}

func main(){
	var emp1 Employee
	emp1.id = 123
	emp1.name = "piyush"
	emp1.age = 22
	emp1.company = "star tele logic"
	fmt.Println(emp1)
	fmt.Println("hello world!")
}