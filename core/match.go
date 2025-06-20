package core

import (
	"log"

	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/templates"
)

func (sdk *sdkCore) match(probe *templates.SearchTemplate, candidate *templates.SearchTemplate) bool {
	matcher, err := sourceafis.NewMatcher(sdk.transparencyLogger, probe)
	if err != nil {
		log.Fatal(err.Error())
	}

	score := matcher.Match(sdk.cntx, candidate)
	return score >= sdk.matchThreshold
}
