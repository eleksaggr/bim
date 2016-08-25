package bim

import (
	"crypto/sha1"
	"errors"
	"os"
	"sync"
)

// FSTree is a structure that manages the file system as a tree.
type FSTree struct {
	parent   *FSTree
	children []*FSTree
	mutex    sync.RWMutex

	name  string
	perm  os.FileMode
	blob  Blob
	isDir bool
}

var errFileNotFound = errors.New("File not found")

// NewTree creates a new tree. An empty tree must always be a directory.
func NewTree(name string, perm os.FileMode) (tree *FSTree) {
	// A new tree must always be a directory, so isDir is permanently true.
	return &FSTree{
		name:     name,
		perm:     perm,
		blob:     nil,
		isDir:    true,
		children: nil,
	}
}

func newFile(name string, perm os.FileMode, blob Blob) (file *FSTree) {
	return &FSTree{
		name:     name,
		perm:     perm,
		blob:     blob,
		isDir:    false,
		children: nil,
	}
}

func (tree *FSTree) clone() *FSTree {
	return &FSTree{
		name:     tree.Name(),
		perm:     tree.Perm(),
		blob:     tree.Blob(),
		isDir:    tree.IsDir(),
		children: tree.children,
		parent:   tree.parent,
	}
}

// InsertFile inserts a file into the tree. If a file with the same name exists,
// this will replace it.
func (tree *FSTree) InsertFile(name string, perm os.FileMode, blob Blob) {
	// Try to remove the file from the tree.
	// If the file does not exist, nothing will happen.
	tree.Remove(name)

	tree.mutex.Lock()
	defer tree.mutex.Unlock()

	file := newFile(name, perm, blob)
	tree.children = append(tree.children, file)
	file.parent = tree
}

// InsertDir inserts a directory into the tree. If a directory with the same name exists,
// this will replace it. InsertDir returns a reference to the newly inserted directory.
func (tree *FSTree) InsertDir(name string, perm os.FileMode) (dir *FSTree) {
	// Try to remove the directory from the tree.
	// If the directory does not exist, nothing will happen.
	tree.Remove(name)

	tree.mutex.Lock()
	defer tree.mutex.Unlock()

	dir = NewTree(name, perm)
	tree.children = append(tree.children, dir)
	dir.parent = tree
	return dir
}

// Remove tries to removes a file/directory from the tree. If the file can not be found,
// nothing will happen.
func (tree *FSTree) Remove(name string) {
	if tree.Name() == name {
		// Disallow removing the tree itself.
		// Why are you doing this?
		return
	}

	item, err := tree.Find(name)
	if err != nil {
		return
	}

	parent := item.parent
	parent.mutex.Lock()
	defer parent.mutex.Unlock()
	for i, child := range parent.children {
		if child == item {
			parent.children = append(parent.children[:i], parent.children[i+1:]...)
		}
	}
}

// Find finds a file or directory in the tree and returns a handle to it, if possible.
// If the file can not be found an errFileNotFound will be returned.
func (tree *FSTree) Find(name string) (*FSTree, error) {
	if tree.Name() == name {
		return tree, nil
	}

	if tree.IsDir() {
		tree.mutex.Lock()
		for _, child := range tree.children {
			if item, err := child.Find(name); err == nil {
				return item, nil
			}
		}
		tree.mutex.Unlock()
	}
	return nil, errFileNotFound
}

// Name returns the name of the FSTree.
func (tree *FSTree) Name() string {
	return tree.name
}

// Perm returns the file permission of the FSTree.
func (tree *FSTree) Perm() os.FileMode {
	return tree.perm
}

// IsDir returns whether the FSTree represents a directory.
func (tree *FSTree) IsDir() bool {
	return tree.isDir
}

// Blob returns the blob held by this FSTree.
func (tree *FSTree) Blob() Blob {
	return tree.blob
}

// Checksum returns the unique checksum for this tree.
func (tree *FSTree) Checksum() Checksum {
	if tree.isDir == false {
		return tree.blob.Checksum()
	}

	h := sha1.New()
	for _, child := range tree.children {
		checksum := child.Checksum()
		h.Write(checksum[:])
	}
	p := h.Sum(nil)
	return HashSumToChecksum(p)
}
