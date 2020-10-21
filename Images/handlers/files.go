package handlers

import (
	"io"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"

	"github.com/ancalabrese/MicService/Images/file"

	"github.com/hashicorp/go-hclog"
)

//File is a handler fore reading and writing files
type File struct {
	logger hclog.Logger
	store  file.Storage
}

//NewFile creates an instace of @Image handler
func NewFile(l hclog.Logger, s file.Storage) *File {
	return &File{logger: l, store: s}
}

//UploadREST implements the http.handler interface for REST uploads
func (image *File) UploadREST(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]
	err := image.saveFile(id, fn, r.Body)
	if err != nil {
		http.Error(rw, "Internal Server Error, couldn't save image", http.StatusInternalServerError)
		return
	}
}

//UploadMultiPart implements the http.handler interface for Multi Part uploads
func (image *File) UploadMultiPart(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(128 * 1024)
	if err != nil {
		image.logger.Error("Bad request", "error", err)
		http.Error(rw, "Expected MultiPart form data", http.StatusBadRequest)
		return
	}
	id := r.FormValue("id")
	if id == "" {
		image.logger.Error("Bad request, product id", "error")
		http.Error(rw, "Bad request, expected product id", http.StatusBadRequest)
		return
	}

	image.logger.Info("Processing form for product id", "id", id)
	f, fh, err := r.FormFile("img")
	if err != nil {
		image.logger.Error("Bad request, expected img file", "error", err)
		http.Error(rw, "Bad request, expected img file", http.StatusBadRequest)
		return
	}
	fileError := image.saveFile(id, fh.Filename, f)
	if fileError != nil {
		http.Error(rw, "Internal Server Error, couldn't save image", http.StatusInternalServerError)
		return
	}

}

func (image *File) saveFile(id string, filename string, r io.ReadCloser) error {
	image.logger.Info("Saving file for product", id, "file name", filename)
	fp := filepath.Join(id, filename)
	err := image.store.Save(fp, r)
	if err != nil {
		image.logger.Error("Unable to save file:", err)
		return err
	}
	return nil
}
