package bim

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Repository asdas FIXME
type Repository struct {
	name    string
	path    string
	staging *os.File
}

// LoadRepo loads a repository from its filesystem representation.
func LoadRepo(absoluteDir string) (repo *Repository, err error) {
	repo = &Repository{
		name: filepath.Base(absoluteDir),
		path: absoluteDir + "/" + repoDir,
	}

	// Change directory into repository directory.
	if err = os.Chdir(repo.path); err != nil {
		return nil, err
	}

	f, err := os.OpenFile(stagingArea, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	repo.staging = f

	return repo, nil
}

func (repo *Repository) resetWorkingDir() error {
	return os.Chdir(repo.path)
}

// SaveBlob saves a blob into the repository and onto the file system.
func (repo *Repository) SaveBlob(blob Blob) (err error) {
	if blob == nil {
		return errors.New("Cannot save empty blob.")
	}

	err = repo.resetWorkingDir()

	// Change into blob directory.
	err = os.Chdir(blobDir)
	if err != nil {
		return err
	}

	name := fmt.Sprintf("%x", blob.Checksum())
	err = ioutil.WriteFile(name, blob, 0644)
	if err != nil {
		return err
	}
	return nil
}

// LoadBlob loads the blob with the specified id from the repository.
func (repo *Repository) LoadBlob(id Checksum) (blob Blob, err error) {
	err = repo.resetWorkingDir()
	if err != nil {
		return nil, err
	}

	err = os.Chdir(blobDir)
	if err != nil {
		return nil, err
	}

	name := fmt.Sprintf("%x", id)
	return ioutil.ReadFile(name)
}
