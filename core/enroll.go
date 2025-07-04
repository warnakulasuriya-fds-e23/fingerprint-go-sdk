package core

import (
	"log"

	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/entities"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/templates"
)

func (sdk *SDKCore) enroll(newEntry *templates.SearchTemplate, id string) (err error) {
	err = sdk.duplicationChecker(newEntry, id)
	if err != nil {
		return
	}
	*sdk.gallery = append(*sdk.gallery, &entities.SearchTemplateRecord{Id: id, Template: *newEntry})
	log.Printf("appended template of %s to gallery\n", id)
	return nil
}
