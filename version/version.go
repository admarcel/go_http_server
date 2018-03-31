package version

import "fmt"

var GitBranch = ""
var GitHash = ""

type Version struct {
	Branch string `json:"git_branch"`
	Hash   string `json:"git_hash"`
}

func GetVersion() (Version) {
	return Version{
		Branch: GitBranch,
		Hash:   GitHash,
	}
}


func (b Version) String() string {
	return fmt.Sprintf("Branch: %s Commit: %s", b.Branch, b.Hash)
}