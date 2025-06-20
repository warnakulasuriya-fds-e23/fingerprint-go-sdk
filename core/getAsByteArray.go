package core

import (
	"log"

	"github.com/fxamacker/cbor/v2"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/templates"
)

func (sdk *SDKCore) getAsByteArray(template *templates.SearchTemplate) *[]byte {
	data, err := cbor.Marshal(*template)
	if err != nil {
		log.Printf("There was an error with Marshalling search template record \n")
	}
	return &data
}
