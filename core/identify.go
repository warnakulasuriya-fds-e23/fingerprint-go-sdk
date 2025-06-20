package core

import (
	"log"

	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/entities"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/templates"
)

func (sdk *sdkCore) identify(probe *templates.SearchTemplate) (ismatched bool, discoverId string) {
	matcher, err := sourceafis.NewMatcher(sdk.transparencyLogger, probe)
	if err != nil {
		log.Fatal(err.Error())
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
	// automatically returns the values of ismatched and discoverId
	return

}
