package main
import ("fmt")

func isPalindrome(str string) bool {
	j:= len(str) - 1
	for i:= 0; i < len(str)/2; i++ {
		if (str[i] != str[j]) {
			return false
		  }
		j = j - 1
	}
	return true
}

func main(){
	if isPalindrome("naman"){
		fmt.Println("It is palindrome")
	} else {
		fmt.Println("It is not")
	}
}