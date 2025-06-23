package core

import (
	"fmt"

	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/templates"
)

func (sdk *SDKCore) extract(imagePath string) (templateptr *templates.SearchTemplate, err error) {
	Img, err := sourceafis.LoadImage(imagePath)
	if err != nil {
		templateptr = nil
		err = fmt.Errorf("error while trying to load single image at %s using sourceafis: %w", imagePath, err)
		return
	}
	templateptr, err = sdk.templateCreator.Template(Img)
	if err != nil {
		templateptr = nil
		err = fmt.Errorf("error while trying to get the template using sourceafis templatecreator: %w", err)
		return
	}
	err = nil
	return
}
