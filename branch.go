package bim

// Branch is a history of commits.
type Branch struct {
	head Reference
	name string
}

// NewBranch creates a new branch from the specified commit.
func NewBranch(name string, commit *Commit) *Branch {
	return &Branch{
		name: name,
		head: Reference{
			Name:   "Head",
			Commit: commit,
		},
	}
}
