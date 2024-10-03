package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
  fmt.Println(m)
}

func main() {
  userNames := make([]string, 2, 6) // 2 is length, 5 is capacity
  userNames[0]="Julie";
  userNames = append(userNames,"Max")
  userNames = append(userNames,"Max")
  userNames = append(userNames,"Max")
  userNames = append(userNames,"Max")
  userNames = append(userNames,"Max")
  fmt.Println(userNames)
  println(len(userNames))
  println(cap(userNames))

  /* courseRatings := make(map[string]float64, 3) */
  courseRatings := make(floatMap, 3)

  courseRatings["go"] = 4.7
  courseRatings["react"] = 4.8
  courseRatings["angular"] = 4.7

  courseRatings.output()

  for index, value := range userNames {
    fmt.Println("Index:", index)
    fmt.Println("Value:", value)
  }
}
