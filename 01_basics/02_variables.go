package main
import ("fmt")

func main(){

	// // 2 method of variable declaration

	// //------------1st (thorugh var keyword) --------------------------
	// syntax => var variableName type = value


	var str string = "learning golang"
	var intNum int = 45
	var floatNum float32 = 23.33
	var isLearning bool = true

	// // go lang support only declaration, initialization can be later
	
	str = str + " seems easy."
	intNum = intNum + 55
	floatNum  = floatNum * 10
	isLearning = false
	
	fmt.Println(direct_str, direct_intNum, direct_floatNum, direct_isLearning)
	
	var unknownStr string
	var unknownInt int
	var unknownFloat float32
	var unknownBool bool

	// // imp: if you are declaring a var, you have to use it otherwise it will give error
	// // if no initialization then default values will be print
	fmt.Println(unknownStr)
	// // for str = ""
	fmt.Println(unknownInt)
	// //for int = 0
	fmt.Println(unknownFloat)
	// // for float = 0
	fmt.Println(unknownBool)
	// //for bool = false

	// fmt.Println(str, intNum, floatNum, isLearning)

	//-------------- 2nd (with := method) -----------------------
	// // syntax => varaiableName := value

	direct_str := "hello world"
	direct_intNum :=  12
	direct_floatNum := 36.22
	direct_isLearning := true

	// direct_str := 12 
	// // this will give assignment error

	fmt.Println(direct_str, direct_intNum, direct_floatNum, direct_isLearning)

}