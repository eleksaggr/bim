package bim

import (
	"crypto/sha1"
	"time"
)

type Commit struct {
	tree    *Tree
	parents []*Commit

	author string
	email  string
	time   time.Time
}

func (commit Commit) Checksum() []byte {
	var data []byte
	data = append(data, commit.tree.Checksum()...)
	data = append(data, []byte(commit.author)...)
	data = append(data, []byte(commit.email)...)

	return sha1.New().Sum(data)
}
