package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Printf("hello, world\n")

    cmd := os.Args[0]
    fmt.Printf("Program Name: %s\n", cmd)
}
