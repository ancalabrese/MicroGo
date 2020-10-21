package file

import (
	"io"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

//Local is an implementation of @storage interface which works with the local storage
type Local struct {
	maxFileSize int    //max file size allowed
	basePath    string //Base directory
}

//NewLocalStorage create a new @Local filesystem
func NewLocalStorage(basePath string, maxFileSize int) (*Local, error) {
	p, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}
	return &Local{maxFileSize, p}, nil
}

//Save the content of the reader to the path location
//path is relative, will be appended to Local.basePath
func (l *Local) Save(id string, content io.ReadCloser) error {
	p := l.fullPath(id)
	d := filepath.Dir(p)
	err := os.MkdirAll(d, os.ModePerm)
	if err != nil {
		return xerrors.Errorf("Unable to create directory: %w", err)
	}
	_, err = os.Stat(id)
	//File already exstis remove it to update it
	if err == nil {
		err = os.Remove(id)
		if err != nil {
			return xerrors.Errorf("Unable to delete existing file: %w", err)
		}
	} else if !os.IsNotExist(err) {
		return xerrors.Errorf("Unable to create file: %w", err)
	}
	f, err := os.Create(p)
	defer f.Close()
	if err != nil {
		return xerrors.Errorf("Unable to create file: %w", err)
	}
	_, err = io.Copy(f, content)
	if err != nil {
		return xerrors.Errorf("Unable to write to file: %w", err)
	}

	return nil
}


//Get the file for the specified loaction
func (l *Local) Get(id string) (*os.File, error) {
	fp := l.fullPath(id)
	file, err := os.Open(fp)
	if err != nil {
		return nil, xerrors.Errorf("Unable to open file at specified location: %w", err)
	}
	return file, nil
}

//Returns base + relative path
func (l *Local) fullPath(path string) string {
	return filepath.Join(l.basePath, path)
}
