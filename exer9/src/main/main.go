package main

import (
  "fmt"

  //"math/rand"
  //"exer9"
  //"math"
)
type ParSum struct {
    sum float64
		SumSquare float64
}


func CalculateSum(arr1 []int, sum1 chan ParSum){
  //  defer close(sum1)
		p := new(ParSum)
		var sum float64  = 0
		var square float64 =0
		for _, a := range arr1{
			sum += float64(a)
			calc := float64(a*a)
			square += calc
		}
		p.sum = sum
		p.SumSquare = square
		sum1 <- *p

}
func MeanStddev(arr []int, chunks int) () {
	//chunks :=2
	//arr1 := []int{2,3,4,5,6,7,8,9,10,11}
	if len(arr)%chunks != 0 {
		panic("You promised that chunks would divide slice size!")
	}

  //go CalculateSum(arr1, c1)

	length := len(arr)
	c1 := make(chan ParSum)
	divide := length/chunks
//	i:=1
	start := divide


	for i := 1; i<=chunks; i++{
		if(i==1){

			go CalculateSum(arr[:(divide)],c1)
      close(c1)
			fmt.Println(divide)

		} else {
			go CalculateSum(arr[start:(divide*i)],c1)
      close(c1)
			start = divide*i
			fmt.Println(divide)
		}

		}
		//sum:=0
		for i:=1; i<=chunks; i++ {
			var sum float64 = 0
      for res := range c1 {
        sum = sum+(res.sum)
        fmt.Println(sum)
        fmt.Println(res.sum, res.SumSquare)
      }

		}

	//	SumSquare := float64(sum)*float64(sum)
	//	mean := (float64(sum)/float64(length))
	//	calc := (float64(SumSquare)/float64(length))
	//	stdDev := math.Sqrt(calc - (mean*mean))

		//fmt.Println(sum)
		//fmt.Println(mean)
		//fmt.Println(stdDev)

	}


func main() {
  //p := exer9.NewPoint(1,2)
  //p.Scale(5)
  //fmt.Println(p)
  	arr1 := []int{2,3,4,5,6,7,8,9,10,11,12,13}
  //  arr:= []float64{3,2,6,8,4,7,5,12,9,11}
  //  QuickSort(arr)
  //  for _, a := range arr{
  //    fmt.Println(a)
  //  }
    //length := len(arr1)
  //  fmt.Println(arr1[5:10])
    //c1 := make(chan int)
    //c2 := make(chan float64)
  //  p := exer9.NewPoint()
 MeanStddev(arr1,3)
//    c1 := make(chan ParSum)
//    go CalculateSum(arr1, c1)
//    for res := range c1 {
//      fmt.Println(res.sum, res.SumSquare)
//    }


    //exer9.MeanStddev(arr1,3)


}
  //fmt.Println(exer9.NewPoint(5,6))
  //fmt.Println(exer8.Hailstone(5))
  //fmt.Println(exer8.HailstoneSequenceAppend(5))
  //fmt.Println(exer8.HailLen(5))
  //fmt.Println(exer8.HailstoneSequenceAllocate(5))
  //pt := exer9.NewPoint(3, 4)
//fmt.Println(pt.Norm() == 5.0)
