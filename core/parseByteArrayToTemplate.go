package core

import (
	"fmt"

	"github.com/fxamacker/cbor/v2"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/templates"
)

func (sdk *SDKCore) parseByteArrayToTemplate(data *[]byte) (templateptr *templates.SearchTemplate, err error) {
	var template templates.SearchTemplate
	err = cbor.Unmarshal(*data, &template)
	if err != nil {
		templateptr = nil
		err = fmt.Errorf("error when while executing unmarshal function for data: %w ", err)
	}
	templateptr = &template
	err = nil
	return
}
