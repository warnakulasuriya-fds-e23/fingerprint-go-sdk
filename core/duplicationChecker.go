package core

import (
	"fmt"

	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/templates"
)

func (sdk *SDKCore) duplicationChecker(templateToCheck *templates.SearchTemplate, idToCheck string) (err error) {
	for _, record := range *sdk.gallery {
		isMatch, errMatchProcess := sdk.match(templateToCheck, &record.Template)
		if errMatchProcess != nil {
			return fmt.Errorf("within duplication checker there was an error while running the match process between the provided template of id %s and the in memory template of %s, because %w", idToCheck, record.Id, errMatchProcess)
		}
		if isMatch == true {
			return fmt.Errorf("running duplication checker a matching template under the id %s was founded for the provided template under provided id %s", record.Id, idToCheck)
		}
		if record.Id == idToCheck {
			return fmt.Errorf("running duplication checker a duplicate id  %s was found for the provided id %s", record.Id, idToCheck)
		}
	}
	return nil

}
