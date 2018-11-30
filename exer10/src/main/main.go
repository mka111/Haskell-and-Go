package main

import (
  "fmt"
  "exer10"
  "math"
)

func main() {

  pt := exer10.NewPoint(3, 4)
  exer10.TurnDouble(&pt, 3*math.Pi/2)
  fmt.Println(pt)
  tri := exer10.NewTriangle(Point{1, 2}, Point{-3, 4}, Point{5, -6}}
  exer10.TurnDouble(&tri, math.Pi)
  fmt.Println(tri)
  fmt.Println(exer10.Fibonacci(25))




//  fmt.Println(exer8.Hailstone(5))
//  fmt.Println(exer8.HailstoneSequenceAppend(5))
//  fmt.Println(exer8.HailLen(5))
//  fmt.Println(exer8.HailstoneSequenceAllocate(5))
//  pt := exer8.NewPoint(3, 4)
// fmt.Println(pt.Norm() == 5.0)
}
