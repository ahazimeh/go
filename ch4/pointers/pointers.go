package main

import "fmt"

func main() {
  age := 32 // regular variable

  var agePointer *int
  agePointer = &age

  fmt.Println("Age:", *agePointer)

  editAgeToAdultYears(agePointer)
  fmt.Println(age)

}

func editAgeToAdultYears(age *int) {
  *age -= 18
}

