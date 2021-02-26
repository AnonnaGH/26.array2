package main

import "fmt"

func main() {

var a [3]float64

a[0] = 20
a[1] = 30
a[2] = 47

var total float64 = 0
   for i := 0 ; i<3 ; i++{
    total += a[i]
   }

fmt.Println(total/2)
}