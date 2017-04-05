
package git

import (
    "fmt"
    "os"
    "os/exec"
    "bytes"
    "strings"
)

type Command struct {
    command string
    parameters []string
    output string
}

var p = fmt.Println
var pf = fmt.Printf

func init() {
    // Do some init magic?
}

// pull from remote
func Pull() {
    fmt.Println("Pulling...")
    params := []string{"push"}
    execute(Command{"git", params, ""})
}

// push to remote
func Push() {
    fmt.Println("Pushing...")
    params := []string{"push"}
    execute(Command{"git", params, ""})
}

// checkout branch and pull from remote

// commit with message, appending branch name in the front of the comment
func Commit(message string) {
    fmt.Println("Commiting")

    currentBranchName := getCurrentBranchName()

    if(currentBranchName != "") {
        params := []string{"commit", "-m " + currentBranchName + ": " + message}
        execute(Command{"git", params, ""})
    } else {
        p("Could not detect branch name. Exiting...")
        os.Exit(1)
    }

}

// make appending the commend configurable via a config file

// create branch both locally and remotely and check it out
// completely pull/update current branch to branch off from before creating the new one

// delete remote branch

// delete all local branches but the one you are currently on
func DeleteOldBranches() {
    // Add user confirmation with a warning that this will FORCE delete!!!
    params := []string{ "branch | grep -v \"master\" | grep -v \\* | xargs git branch -D" }
    execute(Command{"git", params, ""})
}

// merge with --no-ff flag

// ================ HELPERS ================

func execute(command Command) (string){
    cmd := exec.Command(command.command, command.parameters...)

    var out bytes.Buffer
    var stderr bytes.Buffer

    cmd.Stdout = &out
    cmd.Stderr = &stderr
    err := cmd.Run()

    if err != nil {
        fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
        os.Exit(1)
    }
    return out.String()
}

func getCurrentBranchName() (string){
    params := []string{"branch"}
    result := strings.Split(execute(Command{"git", params, ""}), "\n")
    var branchName string

    for _,value := range result {
        position := strings.Index(value, "* ") 
        if position != -1 {
            words := strings.Split(value, " ")
            branchName = strings.Trim(words[1], "\n")
            return branchName
        }
    }

    return ""
}
