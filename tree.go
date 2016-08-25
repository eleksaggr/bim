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

// Tree is a structure that helps organize the files as blobs.
type Tree struct {
	id       [sha1.Size]byte
	children []*Tree
	blob     Blob
	meta     *treeMeta
}

// NewTree creates a new tree with the given details.
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

// NewEmptyTree creates a new tree with no children.
func NewEmptyTree(filename string, mode os.FileMode) *Tree {
	return NewTree(filename, mode, nil, nil)
}

// Insert inserts a tree as a child.
func (tree Tree) Insert(t *Tree) (newTree *Tree) {
	newTree = NewTree(tree.Filename(), tree.Mode(), tree.blob, tree.children)
	// Append the tree to be inserted to the new tree.
	newTree.children = append(newTree.children, t)
	newTree.updateID()
	return newTree
}

// InsertBlob inserts a blob as a child.
func (tree Tree) InsertBlob(filename string, mode os.FileMode, blob Blob) *Tree {
	return tree.Insert(NewTree(filename, mode, tree.blob, nil))
}

// Remove removes the child with the name given as filename.
func (tree Tree) Remove(filename string) (t *Tree) {
	// Disallow removing the tree itself.
	if filename == tree.Filename() {
		return
	}

	t = NewTree(tree.Filename(), tree.Mode(), tree.blob, tree.children)
	for i, child := range t.children {
		if child.Filename() == filename {
			t.children = append(t.children[:i], t.children[i+1:]...)
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

// Filename returns the filename of this tree.
func (tree Tree) Filename() string {
	return tree.meta.Filename
}

// Mode returns the mode of this tree.
func (tree Tree) Mode() os.FileMode {
	return tree.meta.Mode
}
