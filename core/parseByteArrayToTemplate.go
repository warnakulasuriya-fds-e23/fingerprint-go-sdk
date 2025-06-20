package core

import (
	"log"

	"github.com/fxamacker/cbor/v2"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/templates"
)

func (sdk *SDKCore) parseByteArrayToTemplate(data *[]byte) *templates.SearchTemplate {
	var template templates.SearchTemplate
	err := cbor.Unmarshal(*data, &template)
	if err != nil {
		log.Println("Error when while executing unmarshal function for data from ")
	}
	return &template
}
