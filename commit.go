package bim

import (
	"crypto/sha1"
	"errors"
	"time"
)

var errCommitNoTree = errors.New("Cannot create commit without tree")

// Commit is a structure that holds a tree and meta information.
type Commit struct {
	id      [sha1.Size]byte
	tree    *FSTree
	parents []*Commit

	author string
	email  string
	time   time.Time
}

// NewCommit creates a new commit with the provided meta information, the provided tree, and parents as its parent commits.
func NewCommit(author string, email string, tree *FSTree, parents []*Commit) (commit *Commit, err error) {
	if tree == nil {
		return nil, errCommitNoTree
	}
	commit = &Commit{
		author: author,
		email:  email,
		tree:   tree,
		time:   time.Now(),
	}

	if parents == nil {
		commit.parents = make([]*Commit, 0)
	} else {
		commit.parents = parents
	}

	commit.generateID()
	return commit, nil
}

func (commit *Commit) generateID() {
	h := sha1.New()
	h.Write([]byte(commit.author))
	h.Write([]byte(commit.email))
	treeChecksum := commit.tree.Checksum()
	h.Write(treeChecksum[:])
	temp := h.Sum(nil)

	checksum := [sha1.Size]byte{}
	copy(checksum[:], temp[0:20])
	commit.id = checksum
}
