package bim

import (
	"os"
	"testing"
)

const (
	testTreeName = "parentTree"
	testBlobName = "blob"
)

func TestTreeInsert(t *testing.T) {
	tree := NewEmptyTree(testTreeName, os.ModePerm)
	tree = tree.InsertBlob(testBlobName, os.ModePerm, Blob("I am a blob."))
	if tree.id == Blob("I am a blob.").Checksum() {
		t.Errorf("Checksum generated the same hash for blob and tree.")
	}
	if len(tree.children) != 1 {
		t.Errorf("Did not insert blob correctly.")
	}
}

func TestTreeRemove(t *testing.T) {
	tree := NewEmptyTree(testTreeName, os.ModePerm)
	tree = tree.InsertBlob(testBlobName, os.ModePerm, Blob("I am a blob."))
	tree = tree.Remove(testBlobName)

	if len(tree.children) != 0 {
		t.Errorf("Did not remove blob correctly.")
	}
}
