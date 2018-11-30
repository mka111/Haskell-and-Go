package exer10


// TODO: your Hailstone, HailstoneSequenceAppend, HailstoneSequenceAllocate functions


func Fib(n uint, cutoff uint) uint {
  if ((n==0) || (n==1)){
    return 1
  }
  if(n > cutoff){
    c := make(chan uint)
    go CalcFib((n-1), cutoff, c)
    go CalcFib((n-2), cutoff, c)
    return <-c + <-c
  }else {
    return Fib((n-1),cutoff) + Fib((n-2), cutoff)
  }

}

func CalcFib(n uint, cutoff uint, c chan uint) {
  if (n==0 || n==1){
    c <- 1
  } else {
    c <- Fib((n-1),cutoff) + Fib((n-2), cutoff)
  //  c <- c1+c2
  }
}

func Fibonacci(n uint) uint {
	return Fib(n, 20)
}
