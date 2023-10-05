package main



import (
	"fmt"

)

func setDataByLinks(variable *int) {
	fmt.Println("number in function = ", *variable)
	*variable = 30
	fmt.Println("number in function = ", *variable)

}


func main(){
	number := 0
	fmt.Println("number = ", number)

	setDataByLinks(&number)

	fmt.Println("number = ", number)
//	fmt.Println(string(number))
}


