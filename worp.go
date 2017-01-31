package main

import (
    "flag"
    "fmt"
    "os"
    "github.com/worp1900/worp/commit"
)

func main() {
    flag.Usage = func() {
        fmt.Printf("\n")
        fmt.Printf("Usage of %s:\n", os.Args[0])
        fmt.Printf("    worp package command [arguments]\n")
        fmt.Printf("\n")
        fmt.Printf("Available packages:\n")
        fmt.Printf("    git\n")
        flag.PrintDefaults()
    }

    flag.Parse()

    if flag.NArg() == 0 {
        displayWelcomeMessage()
        displayUsageAndExit()
    }

    packageName := flag.Arg(0)

    switch packageName {
        case "commit":
            commit.Checkout()
        default:
            fmt.Printf("no package found\n")
            displayUsageAndExit()
    }
}

func displayWelcomeMessage() {
    fmt.Printf("\n")
    fmt.Printf("Welcome to worp's utility")
    fmt.Printf("\n")
}

func displayUsageAndExit() {
    flag.Usage()
    os.Exit(1)
}
