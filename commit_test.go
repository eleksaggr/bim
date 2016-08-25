package bim

import (
	"os"
	"testing"
)

const (
	testAuthor    = "Tester McTest"
	testEmail     = "test@bim.sexy"
	testDirectory = "test"
	testMode      = os.ModePerm
	testFilename  = "testfile"
)

func TestNewCommit(t *testing.T) {
	// Create an empty tree.
	tree := NewTree(testDirectory, testMode)
	tree.InsertFile(testFilename, testMode, Blob("This is a test file."))

	commit, err := NewCommit(testAuthor, testEmail, tree, nil)
	if err != nil {
		t.Errorf("Creation of commit failed: %v\n", err)
	}

	if commit.author != testAuthor {
		t.Errorf("Did not set commit author correctly.")
	}
	if commit.email != testEmail {
		t.Errorf("Did not set commit email correctly.")
	}
	if commit.tree != tree {
		t.Errorf("Did not set commit working tree correctly.")
	}
}

func TestNewCommitNilTree(t *testing.T) {
	_, err := NewCommit(testAuthor, testEmail, nil, nil)
	if err == nil {
		t.Errorf("Did not catch nil tree correctly.")
	}
}
