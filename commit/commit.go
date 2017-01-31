
package commit

import (
    "fmt"
    "os"
)

func init() {
    // Detect which CVS is present by its hidden folder
}

// checkout branch and pull from remote
func Checkout() {
    fmt.Printf("Hello World!")
}

// commit with message, appending branch name in the front of the comment
// make appending the commend configurable via a config file

// create branch both locally and remotely and check it out
// completely pull/update current branch to branch off from before creating the new one

// delete remote branch

// delete all local branches but the one you are currently on

// merge with --no-ff flag

func detectCvs (path string) (bool, err) {
    _, err = os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExists(err) { return false, nil }
    return true, err
}
