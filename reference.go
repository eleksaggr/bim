package bim

// Reference is a named pointer to a specific commit.
type Reference struct {
	Commit *Commit
	Name   string
}
