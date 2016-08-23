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
	id       [sha1.Size]byte
	children []*Tree
	blob     Blob
	meta     *treeMeta
}

func NewTree(filename string, mode os.FileMode, blob Blob, children []*Tree) *Tree {
	return &Tree{
		meta: &treeMeta{
			Filename: filename,
			Mode:     mode,
		},
		blob:     blob,
		children: children,
	}
}

func NewEmptyTree(filename string, mode os.FileMode) *Tree {
	return NewTree(filename, mode, nil, nil)
}

func (tree Tree) Insert(t *Tree) (newTree *Tree) {
	newTree = NewTree(tree.Filename(), tree.Mode(), tree.blob, tree.children)
	// Append the tree to be inserted to the new tree.
	newTree.children = append(newTree.children, t)
	newTree.updateID()
	return newTree
}

func (tree Tree) InsertBlob(filename string, mode os.FileMode, blob Blob) *Tree {
	return tree.Insert(NewTree(tree.Filename(), tree.Mode(), tree.blob, nil))
}

func (tree Tree) Remove(id [sha1.Size]byte) (t *Tree) {
	t = NewTree(tree.Filename(), tree.Mode(), tree.blob, tree.children)
	for i, child := range tree.children {
		if child.id == id {
			t.children = append(t.children[:i], t.children[:i+1]...)
			break
		}
	}
	t.updateID()
	return t
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
	copy(tree.id[:], h.Sum(nil)[0:sha1.Size])
}

func (tree Tree) isTree() bool {
	if tree.children == nil || len(tree.children) == 0 {
		return false
	}
	return true
}

func (tree Tree) Filename() string {
	return tree.meta.Filename
}

func (tree Tree) Mode() os.FileMode {
	return tree.meta.Mode
}
