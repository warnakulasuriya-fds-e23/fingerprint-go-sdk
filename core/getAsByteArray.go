package core

import (
	"fmt"

	"github.com/fxamacker/cbor/v2"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/templates"
)

func (sdk *SDKCore) getAsByteArray(template *templates.SearchTemplate) (bytesArrayptr *[]byte, err error) {
	data, err := cbor.Marshal(*template)
	if err != nil {
		err = fmt.Errorf("There was an error with Marshalling search template record : %w", err)
		d := make([]byte, 0, 1)
		bytesArrayptr = &d
	}
	bytesArrayptr = &data
	return bytesArrayptr, nil
}
