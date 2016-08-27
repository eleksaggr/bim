package bim

//
// import (
// 	"errors"
// 	"io/ioutil"
// 	"os"
// )
//
// const (
// 	// MaxFileSize is the maximum size a file can have to be added to the working tree.
// 	MaxFileSize = 9223372036854775807 // 64 Megs
// )
//
// var errWorkingTreeNil = errors.New("Can not add file to non-existant working tree.")
//
// // Add adds a file to the working tree. Should the file be a directory, all files in the directory
// // will be added to the tree.
// func Add(file os.File, workingTree *Tree) (tree *Tree, err error) {
// 	if workingTree == nil {
// 		return nil, errWorkingTreeNil
// 	}
// 	stat, err := file.Stat()
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	if stat.IsDir() {
// 		// Insert new tree to workingTree
// 		workingTree = workingTree.Insert(NewEmptyTree(file.Name(), stat.Mode().Perm()))
//
// 		files, err := ioutil.ReadDir(file.Name())
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		tree = workingTree
// 		for _, f := range files {
// 			go func(info os.FileInfo) {
//
// 			}(f)
// 		}
// 	} else {
// 		data := make([]byte, MaxFileSize)
// 		n, err := file.Read(data)
// 		if err != nil {
// 			return nil, err
// 		}
// 		data = data[:n]
// 		tree = workingTree.InsertBlob(file.Name(), stat.Mode().Perm(), Blob(data))
// 	}
// 	return tree, nil
// }
