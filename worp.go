package main

import (
    "flag"
    "fmt"
    "os"
    "github.com/worp1900/worp/commit"
)

func main() {
    
    // Establish a worp setup functionality to create config files et al
    // Read a config file - is there an init() function for main()?

    flag.Usage = func() {
        fmt.Printf("\n")
        fmt.Printf("Usage of %s:\n", os.Args[0])
        fmt.Printf("    worp package command [arguments]\n")
        fmt.Printf("\n")
        fmt.Printf("Available packages:\n")
        fmt.Printf("    commit\n")
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
