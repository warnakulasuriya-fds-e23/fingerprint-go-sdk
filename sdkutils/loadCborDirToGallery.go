package sdkutils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fxamacker/cbor/v2"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/entities"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/templates"
)

func loadCborfileToGallery(galleryptr *[]*entities.SearchTemplateRecord, cborpath string, filename string) error {
	readData, err := os.ReadFile(cborpath)
	if err != nil {
		return fmt.Errorf("Error in reading the file at: %s, error: %w", cborpath, err)
	}
	var template templates.SearchTemplate
	err = cbor.Unmarshal(readData, &template)
	if err != nil {
		return fmt.Errorf("Error when while executing unmarshal function for data from : %s, error: %w ", cborpath, err)
	}
	fileNameWithoutExtenstion := strings.TrimSuffix(filename, filepath.Ext(cborpath))
	*galleryptr = append(*galleryptr, &entities.SearchTemplateRecord{Id: fileNameWithoutExtenstion, Template: template})
	log.Printf("appended template from %s to gallery\n", filename)
	return nil
}

func LoadCborDirToGallery(galleryptr *[]*entities.SearchTemplateRecord, cborDir string) error {
	files, err := os.ReadDir(cborDir)
	if err != nil {
		return fmt.Errorf("Error! Unable to read data from %s, error: %w", cborDir, err)
	}

	for _, file := range files {
		err = loadCborfileToGallery(galleryptr, filepath.Join(cborDir, file.Name()), file.Name())
		if err != nil {
			return err
		}
	}
	return nil

}
