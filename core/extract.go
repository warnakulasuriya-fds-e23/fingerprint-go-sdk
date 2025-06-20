package core

import (
	"log"

	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/templates"
)

func (sdk *sdkCore) extract(imagePath string) *templates.SearchTemplate {
	Img, err := sourceafis.LoadImage(imagePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	template, err := sdk.templateCreator.Template(Img)
	if err != nil {
		log.Fatal(err.Error())
	}
	return template
}
