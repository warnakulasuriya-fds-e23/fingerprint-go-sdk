package core

import (
	"fmt"

	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/templates"
)

func (sdk *SDKCore) match(probe *templates.SearchTemplate, candidate *templates.SearchTemplate) (isMatch bool, err error) {
	matcher, err := sourceafis.NewMatcher(sdk.transparencyLogger, probe)
	if err != nil {
		isMatch = false
		err = fmt.Errorf("error when creating a new matcher using sourceafis: %w ", err)
		return
	}

	score := matcher.Match(sdk.cntx, candidate)
	isMatch = score >= sdk.matchThreshold
	err = nil
	return
}
