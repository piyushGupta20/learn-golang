package main
import ("fmt")


func main(){
	// they are decalred with const keyword
	// const PI float32 = 3.14
	
	// declare multiple cosntants
	// const (
	// 	PI = 3.14
	// 	MONGO_DB_PORT = 27017
	// 	MONGO_URL = "mongodb://127.0.0.1:27017"
	// )

	//in single line, you have to use semi-colon
	const ( PI = 3.14; MONGO_DB_PORT = 27017; MONGO_URL = "mongodb://127.0.0.1:27017" )

	fmt.Println(PI, MONGO_DB_PORT, MONGO_URL)

}