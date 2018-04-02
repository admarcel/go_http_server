package version

import (
	"fmt"
	"net/http"
	"encoding/json"
)

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

func Handler(w http.ResponseWriter, r *http.Request){
	current_version := GetVersion()
	payload, err := json.Marshal(current_version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}


func (b Version) String() string {
	return fmt.Sprintf("Branch: %s Commit: %s", b.Branch, b.Hash)
}