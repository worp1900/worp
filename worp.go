package main

import (
    "flag"
    "fmt"
    "os"
    "github.com/worp1900/worp/cvs"
)

func main() {
    // Establish a worp setup functionality to create config files et al
    // Read a config file - is there an init() function for main()?

    flag.Usage = func() {
        fmt.Printf("\n")
        fmt.Printf("Usage of %s:\n", os.Args[0])
        fmt.Printf("    worp command [arguments]\n")
        fmt.Printf("\n")
        fmt.Printf("Available commands:\n")
        fmt.Printf("    pull\n")
        fmt.Printf("    push\n")
        fmt.Printf("    commit message\n")
        fmt.Printf("    deleteOldBranches\n")
        fmt.Printf("    \n")
        flag.PrintDefaults()
    }

    flag.Parse()

    if flag.NArg() == 0 {
        displayWelcomeMessage()
        displayUsageAndExit()
    }

    command := flag.Arg(0)

    switch command {
        case "commit":
            message := flag.Arg(1)

            fmt.Println(message)
 
            if flag.NArg() == 1 {
                fmt.Printf("\n")
                fmt.Printf("Could not detect necessary message parameter - exiting")
                fmt.Printf("\n")
                displayUsageAndExit()
            }

            cvs.Commit(message)
        case "pull":
            cvs.Pull()
        case "push":
            cvs.Push()
        case "deleteOldBranches":
            cvs.DeleteOldBranches()
        default:
            fmt.Printf("unknown command\n")
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
