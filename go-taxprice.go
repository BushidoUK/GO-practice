package main

import "fmt"

func main() {
  coolSneakers := 65.99
  niceNecklace := 45.50
  
  // Define taxCalculation here
  var taxCalculation float64
  
  // Add coolSneakers to taxCalculation
  taxCalculation += coolSneakers
  
  // Add niceNecklace to taxCalculation
  taxCalculation += niceNecklace
  
  // Compute the NYC sales tax
  // 8.875% of the purchase here:
  taxCalculation *= .08875
  
  //Receipt
  fmt.Println("Purchase of", coolSneakers + niceNecklace, "with 8.875% sales tax", taxCalculation, "equal to", coolSneakers + niceNecklace + taxCalculation)
}
