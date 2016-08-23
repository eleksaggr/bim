package bim

import (
	"crypto/sha1"
	"errors"
	"os"
)

var errNilTree = errors.New("Cannot operate on nil-tree.")

type treeMeta struct {
	Filename string
	Mode     os.FileMode
}

type Tree struct {
	id       [20]byte
	children []*Tree
	blob     Blob
	meta     *treeMeta
}

func (tree Tree) Insert(t *Tree) (newTree *Tree, err error) {
	if t == nil {
		return nil, errNilTree
	}
	// Copy old tree meta data to new tree.
	newTree = &Tree{
		blob: tree.blob,
		meta: tree.meta,
	}
	// Append the tree to be inserted to the new tree.
	newTree.children = append(newTree.children, t)
	newTree.updateID()
	return newTree, nil
}

func (tree Tree) InsertBlob(filename string, mode os.FileMode, blob Blob) (*Tree, error) {
	return tree.Insert(&Tree{
		children: nil,
		blob:     blob,
		meta: &treeMeta{
			Filename: filename,
			Mode:     mode,
		},
	})
}

func (tree *Tree) updateID() {
	h := sha1.New()
	for _, child := range tree.children {
		if child.isTree() {
			h.Write(child.id[:])
		} else {
			checksum := child.blob.Checksum()
			h.Write(checksum[:])
		}
	}
}

func (tree Tree) isTree() bool {
	if tree.children == nil || len(tree.children) == 0 {
		return false
	}
	return true
}
