package core

import (
	"fmt"

	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/entities"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/templates"
)

func (sdk *SDKCore) identify(probe *templates.SearchTemplate) (ismatched bool, discoverId string, err error) {
	matcher, err := sourceafis.NewMatcher(sdk.transparencyLogger, probe)
	if err != nil {
		ismatched = false
		discoverId = "error"
		err = fmt.Errorf("error when creating a new matcher using sourceafis: %w", err)
	}
	max := -10000.0
	var matchingRecord *entities.SearchTemplateRecord
	for _, record := range *sdk.gallery {
		score := matcher.Match(sdk.cntx, &record.Template)
		if score >= max {
			max = score
			matchingRecord = record
		}
	}

	if max > sdk.matchThreshold {
		ismatched = true
		discoverId = matchingRecord.Id
	} else {
		ismatched = false
		discoverId = "none"
	}
	err = nil
	// automatically returns the values of ismatched, discoveredId and err
	return

}
