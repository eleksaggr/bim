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
	if tree.BlobID() != dirPlaceholderChecksum {
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

func TestInsertFile(t *testing.T) {
	testTreeName := "TestTree"
	testTreePerm := os.ModePerm

	testFileName := "TestFile"
	testFilePerm := os.ModePerm
	testFileBlob := Blob("I'm a test file.")

	tree := NewTree(testTreeName, testTreePerm)
	tree.InsertFile(testFileName, testFilePerm, testFileBlob.Checksum())

	if len(tree.children) != 1 {
		t.Error("Did not insert file correctly.")
	}
	if tree.children[0].IsDir() == true {
		t.Error("Inserted directory instead of file.")
	}

	if tree.children[0].parent != tree {
		t.Error("Did not set childs parent correctly.")
	}
}

func TestInsertDir(t *testing.T) {
	testTreeName := "TestTree"
	testTreePerm := os.ModePerm

	testDirName := "TestDirectory"
	testDirPerm := os.ModePerm

	tree := NewTree(testTreeName, testTreePerm)
	dir := tree.InsertDir(testDirName, testDirPerm)

	if len(tree.children) != 1 {
		t.Error("Did not insert file correctly.")
	}
	if dir.IsDir() != true {
		t.Error("Inserted file instead of directory.")
	}

	if dir.parent != tree {
		t.Error("Did not set childs parent correctly.")
	}
}

func TestFind(t *testing.T) {
	testTreeName := "TestTree"
	testTreePerm := os.ModePerm

	tree := NewTree(testTreeName, testTreePerm)
	tree.InsertFile("Layer 1 File", os.ModePerm, Blob("I'm in layer 1!").Checksum())
	dir := tree.InsertDir("Directory1", os.ModePerm)
	findme := dir.InsertFile("FindMe", os.ModePerm, Blob("Plz find me. thx <3").Checksum())
	dir2 := tree.InsertDir("Directory2", os.ModePerm)
	dir3 := dir2.InsertDir("Directory3", os.ModePerm)
	dir3.InsertFile("DontFindMe", os.ModePerm, Blob("Don't find me!").Checksum())

	file, err := tree.Find("FindMe")
	if err != nil {
		t.Errorf("Did not find file: %v\n", err)
	}

	if file != findme {
		t.Error("Found the wrong file.")
	}
}
