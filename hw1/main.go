package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := float64(1)
    for i :=0; i<10; i++{
        z -=(z*z-x)/(2*z)
    }
    return z
}

func SqrtMoreThan10(x float64) (float64,int) {
    z := float64(1)
    last := float64(0)
    cnt := 0
    for{
            z -=(z*z-x)/(2*z)
        if (float32(last) == float32(z)){
            break
            }
        last = z
        cnt++
    }
    return z,cnt
}

func main() {
    fmt.Println(Sqrt(19))
    fmt.Println(SqrtMoreThan10(19))
    fmt.Println(math.Sqrt(19))
}

