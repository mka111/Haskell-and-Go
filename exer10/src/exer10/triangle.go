package exer10

// TODO: Point (with everything from exercise 8 and) and methods that modify them in-place
//copied functions from exercise 8
import (
  "fmt"
  "math"
)
type Point struct {
    x float64
    y float64
}

type Triangle struct {
	A, B, C Point
}

type Transformable interface{
  Scale(f float64)
  Rotate(a float64)
}

func NewPoint(a float64, b float64) Point {
  p := Point{x: a, y: b}
  return p
}

func (t Point) String() string{
  return fmt.Sprintf("(%v, %v)",t.x, t.y)
}
func (t Point) Norm() float64{
  f := math.Sqrt((t.x * t.x) + (t.y *t.y))
  return f
}

func (t *Point) Scale(f float64){
  t.x = (t.x)*f
  t.y = (t.y)*f
  //p := NewPoint(a, b)
//  return t

}

func (t *Point) Rotate(a float64){
  val1 := math.Cos(a)
  val2 := math.Sin(a)
  temp := t.x
  //fmt.Println("Sin value", val2)
  //fmt.Println("t.x", t.x)
  //fmt.Println("t,y", t.y)
  t.x = ((t.x)*val1) - ((t.y)*val2)
  t.y = ((temp)*val2) + ((t.y)*val1)

}

func (t *Triangle) Scale(f float64){
  t.A.Scale(f)
  t.B.Scale(f)
  t.C.Scale(f)
}

func (t *Triangle) Rotate(a float64){
  t.A.Rotate(a)
  t.B.Rotate(a)
  t.C.Rotate(a)

}

func TurnDouble(t Transformable, angle float64){
  t.Scale(2)
  t.Rotate(angle)
}
