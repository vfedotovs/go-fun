package main

import (
        "fmt"
)

func main() {
    var n1, n2 int
    var sum, sub, mult  int
    var n3, n4, div float32

    n1 = 9
    n2 = 6
    n3 = 9
    n4 = 6

    sum = n1 + n2
    sub = n1 - n2
    mult = n1 * n2
    div = n3 / n4
    fmt.Println(n1, " + ", n2, " = ", sum)
    fmt.Println(n1, " - ", n2, " = ", sub)
    fmt.Println(n1, " * ", n2, " = ", mult)
    fmt.Println(n3, " / ", n4, " = ", div)
}
