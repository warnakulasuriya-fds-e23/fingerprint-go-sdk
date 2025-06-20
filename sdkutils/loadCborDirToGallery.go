package sdkutils

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fxamacker/cbor/v2"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/entities"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/templates"
)

func loadCborfileToGallery(galleryptr *[]*entities.SearchTemplateRecord, cborpath string, filename string) {
	readData, err := os.ReadFile(cborpath)
	if err != nil {
		log.Println("Error in reading the file at", cborpath)
		return
	}
	var template templates.SearchTemplate
	err = cbor.Unmarshal(readData, &template)
	if err != nil {
		log.Println("Error when while executing unmarshal function for data from ", cborpath)
	}
	fileNameWithoutExtenstion := strings.TrimSuffix(filename, filepath.Ext(cborpath))
	*galleryptr = append(*galleryptr, &entities.SearchTemplateRecord{Id: fileNameWithoutExtenstion, Template: template})
	log.Printf("appended template from %s to gallery\n", filename)
}

func LoadCborDirToGallery(galleryptr *[]*entities.SearchTemplateRecord, cborDir string) {
	files, err := os.ReadDir(cborDir)
	if err != nil {
		log.Println("Error! Unable to read data")
	}

	for _, file := range files {
		loadCborfileToGallery(galleryptr, filepath.Join(cborDir, file.Name()), file.Name())
	}

}
