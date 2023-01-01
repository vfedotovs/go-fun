package main

import "fmt"

func main() {
    var n1, n2 int
    var sum, sub, mult  int
    var n3, n4, div float32

    fmt.Print("Enter num1:")
    fmt.Scan(&n1)
    fmt.Print("Enter num2:")
    fmt.Scan(&n2)
    fmt.Print("Enter num3:")
    fmt.Scan(&n3)
    fmt.Print("Enter num4:")
    fmt.Scan(&n4)

    sum = n1 + n2
    sub = n1 - n2
    mult = n1 * n2
    div = n3 / n4
    fmt.Println(n1, " + ", n2, " = ", sum)
    fmt.Println(n1, " - ", n2, " = ", sub)
    fmt.Println(n1, " * ", n2, " = ", mult)
    fmt.Println(n3, " / ", n4, " = ", div)
}
