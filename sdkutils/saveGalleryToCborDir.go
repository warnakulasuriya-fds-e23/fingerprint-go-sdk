package sdkutils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fxamacker/cbor/v2"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/entities"
)

func saveTemplateToCborDir(searchTemplateRecordptr *entities.SearchTemplateRecord, cborDir string) error {
	data, err := cbor.Marshal(searchTemplateRecordptr.Template)
	if err != nil {
		fmt.Errorf("There was an error with Marshalling search template record with id %s, error: %w", searchTemplateRecordptr.Id, err)
	}
	saveFilePath := filepath.Join(cborDir, searchTemplateRecordptr.Id+".cbor")

	err = os.WriteFile(saveFilePath, data, 0755)
	if err != nil {
		return fmt.Errorf("There was an error writing to %s, error: %w", saveFilePath, err)
	}

	log.Printf("Successfully saved %s \n", searchTemplateRecordptr.Id+".cbor")
	return nil
}

func SaveGalleryToCborDir(gallery *[]*entities.SearchTemplateRecord, cborDir string) error {
	for _, searchTemplateRecordptr := range *gallery {
		err := saveTemplateToCborDir(searchTemplateRecordptr, cborDir)
		if err != nil {
			return err
		}
	}
	return nil
}
