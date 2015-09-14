package main

import "math/big"
import "fmt"

func main() {
  p1, _ := big.NewInt(0).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
  fmt.Printf("%v\n", p1)

  two := big.NewInt(2)
  p2 := big.NewInt(0).Exp(two, big.NewInt(256),nil)
  p2 = p2.Sub(p2, big.NewInt(0).Exp(two, big.NewInt(32),nil))
  p2 = p2.Sub(p2, big.NewInt(0).Exp(two, big.NewInt(9),nil))
  p2 = p2.Sub(p2, big.NewInt(0).Exp(two, big.NewInt(8),nil))
  p2 = p2.Sub(p2, big.NewInt(0).Exp(two, big.NewInt(7),nil))
  p2 = p2.Sub(p2, big.NewInt(0).Exp(two, big.NewInt(6),nil))
  p2 = p2.Sub(p2, big.NewInt(0).Exp(two, big.NewInt(4),nil))
  p2 = p2.Sub(p2, big.NewInt(1))

  fmt.Printf("%v\n", p2)
}
