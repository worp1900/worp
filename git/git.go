
package git

import (
    "fmt"
    "os"
    "os/exec"
)

type Command struct {
    command string
    parameters []string
    output string
}

func init() {
    // Do some init magic
}

// pull from remote
func Pull() {
    fmt.Println("Pulling...")

    var cmdOut []byte
    var err error

    if cmdOut, err = exec.Command("git", "pull").Output(); err != nil {
        fmt.Fprintln(os.Stderr, "There was an error running git pull: ", err)
        os.Exit(1)
    }

    fmt.Println(string(cmdOut))
}

// push to remote
func Push() {
    fmt.Println("Pushing...")

    var cmdOut []byte
    var err error

    if cmdOut, err = exec.Command("git", "push").Output(); err != nil {
        fmt.Fprintln(os.Stderr, "There was an error running git push: ", err)
        os.Exit(1)
    }

    fmt.Println(string(cmdOut))
}

// checkout branch and pull from remote
func Commit() {
}

// commit with message, appending branch name in the front of the comment
// make appending the commend configurable via a config file

// create branch both locally and remotely and check it out
// completely pull/update current branch to branch off from before creating the new one

// delete remote branch

// delete all local branches but the one you are currently on

// merge with --no-ff flag

func execute(cmd Command) (bool){
    fmt.Println(cmd)
    return true
}





