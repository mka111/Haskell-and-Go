package exer9

import (
	//"math"
	"math/rand"
	//"fmt"
)

type ParSum struct {
    sum float64
		SumSquare float64
}

func RandomArray(length int, maxInt int) []int {
	// TODO: create a new random generator with a decent seed; create an array with length values from 0 to values-1.
	array := make([]int, length)
	print(array)
	i := 0
	for i < length{
		array[i] = rand.Intn(maxInt)
		i = i+1
	}
	return array




}



func CalculateSum(arr1 []int, sum1 chan ParSum){
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

//	for i<=chunks{
		//go calculateSum(arr1[:divide*i])
//		go calculateSum(arr1[5:10],c)
//		divide = divide*i

//	}








//func MeanStddev(arr []int, chunks int) (mean, stddev float64) {
	//if len(arr)%chunks != 0 {
	//	panic("You promised that chunks would divide slice size!")
//	}

	// TODO: calculate the mean and population standard deviation of the array, breaking the array into chunks segments
	// and calculating on them in parallel.
	//ch := make(chan int)
















//}
