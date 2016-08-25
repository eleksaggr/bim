package bim

import (
	"os"
	"testing"
)

func TestTreeInsert(t *testing.T) {
	tree := NewEmptyTree("parent", os.ModePerm)
	tree = tree.InsertBlob("child", os.ModePerm, Blob("Child"))
	if tree.id == Blob("Child").Checksum() {
		t.Errorf("Checksum generated the same hash for blob and tree.")
	}
	if len(tree.children) != 1 {
		t.Errorf("Did not insert blob correctly.")
	}
}

func TestTreeRemove(t *testing.T) {
	tree := NewEmptyTree("parent", os.ModePerm)
	tree = tree.InsertBlob("child", os.ModePerm, Blob("Child"))
	tree = tree.Remove(tree.children[0].id)

	if len(tree.children) != 0 {
		t.Errorf("Did not remove blob correctly.")
	}
}
