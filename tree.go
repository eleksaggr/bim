package bim

import "crypto/sha1"

type treeData struct {
	Filename string
	Mode     uint8
	Blob     []byte
}

type Tree struct {
	id    []byte
	nodes []*Tree
	Data  *treeData
}

func NewTree() (tree *Tree) {
	tree.updateID()
	return tree
}

func (tree *Tree) Insert(filename string, mode uint8, blob []byte) {
	tree.InsertTree(&Tree{
		nodes: nil,
		Data: &treeData{
			Filename: filename,
			Mode:     mode,
			Blob:     blob,
		},
	})
}

func (tree *Tree) InsertTree(t *Tree) {
	tree.nodes = append(tree.nodes, t)
}

func (tree *Tree) Remove(t *Tree) {
	for i, node := range tree.nodes {
		if node == t {
			tree.nodes = append(tree.nodes[:i], tree.nodes[i+1:]...)
		}
	}
}

func (tree Tree) Hash() []byte {
	hash := make([]byte, sha1.Size)
	hasher := sha1.New()
	for _, node := range tree.nodes {
		if node.nodes == nil {
			// Node is a blob.
			hash = hasher.Sum(node.Data.Blob)
		} else {
			// Node is a tree.
			hash = hasher.Sum(node.Hash())
		}
	}
	return hash
}

func (tree *Tree) updateID() {
	tree.id = tree.Hash()
}

func (tree Tree) ID() []byte {
	return tree.id
}
