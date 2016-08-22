package bim

type treeData struct {
	Filename string
	Mode     uint8
	Blob     []byte
}

type Tree struct {
	nodes []*Tree
	Data  *treeData
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
