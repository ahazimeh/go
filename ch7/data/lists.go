package main

import "fmt"

func main() {
  prices := []float64{10.99, 8.99} // this syntax create a slice and go create a slice bhind the scenes
  fmt.Println(prices[1])

  prices = append(prices, 5.99) // go will create a new array and add this element to the new array
  prices = prices[1:] // removes the first element
  fmt.Println(prices)
}

/* func main() {
  var productNames [4]string = [4]string{"A book"}
  prices := [4]float64{10.99, 9.99, 45.99, 20.0}
  fmt.Println(prices)

  productNames[2] = "A Carpet"

  fmt.Println(productNames)

  fmt.Println(prices[0])

  featuredPrices := prices[1:3]

  fmt.Println(featuredPrices)

} */
