package file

import (
	"io"
	"os"
)

//Storage defines the behavior for files
//Implementations may choose to save the file locally, in a db or upload it in the cloud
type Storage interface {
	Save(id string, file io.ReadCloser) error
	Get(id string) (*os.File, error)
}
