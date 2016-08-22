package bim

import "crypto/sha1"

type treeMetadata struct {
	Filename string
	Mode     uint8
	Blob     Blob
}

type Tree struct {
	nodes []*Tree
	Data  *treeMetadata
}

func (tree Tree) Checksum() []byte {
	checksum := make([]byte, sha1.Size)
	h := sha1.New()
	for _, node := range tree.nodes {
		if node.nodes == nil {
			checksum = h.Sum(node.Data.Blob)
		} else {
			checksum = h.Sum(node.Checksum())
		}
	}
	return checksum
}
