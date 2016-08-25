package bim

import (
	"os"
	"testing"
)

func TestNewTree(t *testing.T) {
	testTreeName := "TestTree"
	testTreePerm := os.ModePerm

	tree := NewTree(testTreeName, testTreePerm)

	if tree.Name() != testTreeName {
		t.Error("Did not set tree name correctly.")
	}
	if tree.Perm() != testTreePerm {
		t.Error("Did not set tree permissions correctly.")
	}
	if tree.Blob() != nil {
		t.Error("Directory tree may not contain blob.")
	}
	if tree.IsDir() == false {
		t.Error("NewTree must always return directory tree.")
	}
	if len(tree.children) != 0 {
		t.Error("Tree has too many children.")
	}
	if tree.parent != nil {
		t.Error("Tree has a parent, when it should be root.")
	}
}
