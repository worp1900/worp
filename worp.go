package main

import (
    "fmt"
    "os"
    "flag"
)

func main() {
    fmt.Printf("hello, world\n")

    cmd := os.Args[0]
    fmt.Printf("Program Name: %s\n", cmd)

    argCount := len(os.Args[1:])
    fmt.Printf("Total arguments: %d\n", argCount)

    for i, a := range os.Args[1:] {
    	fmt.Printf("Argument %d is %s\n", i+1, a)
    }
}
