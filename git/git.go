
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
func Commit() {
}

// commit with message, appending branch name in the front of the comment
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

func execute(cmd Command) (bool){
    cmd := exec.Command("find", "/", "-maxdepth", "1", "-exec", "wc", "-c", "{}", "\\")
    var out bytes.Buffer
    var stderr bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &stderr
    err := cmd.Run()
    if err != nil {
    fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
        return
    }
    fmt.Println("Result: " + out.String())
    return true
}





