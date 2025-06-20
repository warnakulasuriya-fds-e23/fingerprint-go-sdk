package sdkutils

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fxamacker/cbor/v2"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/entities"
)

func saveTemplateToCborDir(searchTemplateRecordptr *entities.SearchTemplateRecord, cborDir string) {
	data, err := cbor.Marshal(searchTemplateRecordptr.Template)
	if err != nil {
		log.Printf("There was an error with Marshalling search template record with id %s\n", searchTemplateRecordptr.Id)
		return
	}
	saveFilePath := filepath.Join(cborDir, searchTemplateRecordptr.Id+".cbor")

	err = os.WriteFile(saveFilePath, data, 0755)
	if err != nil {
		log.Printf("There was an error writing to %s\n", saveFilePath)
		return
	}

	log.Printf("Successfully saved %s \n", searchTemplateRecordptr.Id+".cbor")
}

func SaveGalleryToCborDir(gallery *[]*entities.SearchTemplateRecord, cborDir string) {
	for _, searchTemplateRecordptr := range *gallery {
		saveTemplateToCborDir(searchTemplateRecordptr, cborDir)
	}
}
