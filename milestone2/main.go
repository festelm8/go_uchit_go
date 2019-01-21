package main

import "fmt"
import "go_uchit_go/milestone2/bitch"

func main() {
    xs := []float64{1,2,3,4}
    avg := bitch.Average(xs)
    fmt.Println(avg)
}