
package commit

import (
    "fmt"
    "os"
)

type Cvs struct {
    category string
}

var cvs = Cvs{}

func init() {
    // Detect which CVS is present by its hidden folder
    cvs.category = detectCvs()
}

// checkout branch and pull from remote
func Checkout() {
    fmt.Printf("Doing %v checkout", cvs.category)
}

// commit with message, appending branch name in the front of the comment
// make appending the commend configurable via a config file

// create branch both locally and remotely and check it out
// completely pull/update current branch to branch off from before creating the new one

// delete remote branch

// delete all local branches but the one you are currently on

// merge with --no-ff flag

func detectCvs() (string) {
    git, err := isFolderPresent("/home/worp/repositories/github/niftySnippets/.git")
    if err != nil { panic(err) }
    
    var cvs string
    if git == true { cvs = "git" }

    fmt.Printf("Detected cvs in path is: %v\n", cvs)
    return cvs
}

func isFolderPresent(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}
