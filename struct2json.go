package main

import (
  "encoding/json"
  "fmt"
  "log"
_  "reflect"
)

//type Student struct {
//  Name string `json:"name"`
//  Age int `json:"age"`
//}

type Student struct {
  Name string
  Age int
}







func main() {
//  student := Student{Name: "Krunal Lathiya", Age: 30}
//  student := Student{Name: "Oleksii Untila", Age: 16}
//  fmt.Println("name: ", student.name)
//  fmt.Println("age: ", student.age)


        students := []Student{}

  	student := Student{Name: "Oleksii Untila", Age: 16}
        students = append(students, student)

  	student = Student{Name: "Anton Antonov", Age: 16}
        students = append(students, student)


  // Convert struct to JSON
  jsonData, err := json.Marshal(students)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Struct as JSON:", string(jsonData))




}
