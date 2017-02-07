
package cvs

import (
    "fmt"
    "os"
    "github.com/worp1900/worp/git"
)

var path string

type Cvs struct {
    category string
}

var cvs = Cvs{}

func init() {
    var err error
    path, err = os.Getwd()
    if err != nil { panic(err) }
}

func prepare() {
    // Detect which CVS is present by its hidden folder
    cvs.category = detectCvs()
}

// pull from remote
func Pull() {
    prepare()
    // delete code duplication
    switch cvs.category {
        case "git":
            git.Pull()
        default:
            fmt.Printf("Unable to detect known cvs system")
    }
}

// push to remote
func Push() {
    prepare()
    switch cvs.category {
        case "git":
            git.Push()
        default:
            fmt.Printf("Unable to detect known cvs system")
    }
}

// checkout branch and pull from remote
func Commit() {
    prepare()
    switch cvs.category {
        case "git":
            git.Commit()
        default:
            fmt.Printf("Unable to detect known cvs system")
    }
}

// delete all local branches but the one you are currently on
func DeleteOldBranches() {
    prepare()
    switch cvs.category {
        case "git":
            git.DeleteOldBranches()
        default:
            fmt.Printf("Unable to detect known cvs system")
    }
}

// commit with message, appending branch name in the front of the comment
// make appending the commend configurable via a config file

// create branch both locally and remotely and check it out
// completely pull/update current branch to branch off from before creating the new one

// delete remote branch

// merge with --no-ff flag

func detectCvs() (string) {
    git, err := isFolderPresent(path + "/.git")
    if err != nil { panic(err) }

    var cvs string
    if git == true {
        fmt.Printf("detected git cvs in %v\n", path)
        cvs = "git"
    }

    return cvs
}

func isFolderPresent(path string) (bool, error) {
    // make this function accept a folder name and return a string
    // so we can have a switch case in detectCvs()
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}
