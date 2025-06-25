package core

import (
	"context"
	"fmt"
	"log"
	"runtime"

	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/entities"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/sdkutils"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/config"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/templates"
)

// SDKCore will be an opaque type and can only be initialized using NewSDKCore
type SDKCore struct {
	gallery            *[]*entities.SearchTemplateRecord
	transparencyLogger *sourceafis.DefaultTransparencyLogger
	templateCreator    *sourceafis.TemplateCreator
	imagesDir          string
	cborDir            string
	cntx               context.Context
	matchThreshold     float64
	initialized        bool
}

func NewSDKCore(imagesDir string, cborDir string) (*SDKCore, error) {
	config.LoadDefaultConfig()
	config.Config.Workers = runtime.NumCPU()

	l := sourceafis.NewTransparencyLogger(new(customTransparencyContents))
	tc := sourceafis.NewTemplateCreator(l)
	g := make([]*entities.SearchTemplateRecord, 0, 1300)
	cntx := context.Background()
	matchThreshold := 40.00
	sdk := &SDKCore{
		gallery:            &g,
		transparencyLogger: l,
		templateCreator:    tc,
		imagesDir:          imagesDir,
		cborDir:            cborDir,
		cntx:               cntx,
		matchThreshold:     matchThreshold,
		initialized:        true,
	}
	log.Println("SDK core has been initialized")

	return sdk, nil

}
func (sdk *SDKCore) UpdateImageDir(newImagesDir string) (message string, err error) {
	err = sdkutils.ProcessDirPathString(&newImagesDir)
	if err != nil {
		message = "Error Occured"
		//err alredy set
		return
	}
	sdk.imagesDir = newImagesDir
	message = "successfully updated"
	err = nil
	return
}
func (sdk *SDKCore) GetStaus() string {
	return fmt.Sprintf(
		" Sdk intialization: %t, match threshold: %f, number template records stored in gallery: %d, current imagesDir: %s, current cborDir: %s",
		sdk.initialized,
		sdk.matchThreshold,
		len(*sdk.gallery),
		sdk.imagesDir,
		sdk.cborDir,
	)

}
func (sdk *SDKCore) GetImagesDir() string {
	return sdk.imagesDir
}
func (sdk *SDKCore) UpdateCborDir(newCborDir string) (message string, err error) {
	err = sdkutils.ProcessDirPathString(&newCborDir)
	if err != nil {
		message = "Error Occured"
		//err alredy set
		return
	}
	sdk.imagesDir = newCborDir
	message = "successfully updated"
	err = nil
	return
}
func (sdk *SDKCore) GetCborDir() string {
	return sdk.cborDir
}
func (sdk *SDKCore) LoadImages() error {
	return sdkutils.LoadImagesDirToGallery(sdk.gallery, sdk.imagesDir)
}
func (sdk *SDKCore) LoadCborfiles() error {
	return sdkutils.LoadCborDirToGallery(sdk.gallery, sdk.cborDir)
}
func (sdk *SDKCore) SaveGallery() error {
	return sdkutils.SaveGalleryToCborDir(sdk.gallery, sdk.cborDir)
}
func (sdk *SDKCore) Extract(imagePath string) (template *templates.SearchTemplate, err error) {
	template, err = sdk.extract(imagePath)
	return
}
func (sdk *SDKCore) Match(probe *templates.SearchTemplate, candidate *templates.SearchTemplate) (isMatch bool, err error) {
	isMatch, err = sdk.match(probe, candidate)
	return
}
func (sdk *SDKCore) Identify(probe *templates.SearchTemplate) (isMatched bool, discoveredId string, err error) {
	isMatched, discoveredId, err = sdk.identify(probe)
	return
}
func (sdk *SDKCore) Enroll(newEntry *templates.SearchTemplate, id string) (err error) {
	err = sdk.enroll(newEntry, id)
	return
}
func (sdk *SDKCore) GetAsByteArray(probe *templates.SearchTemplate) (data *[]byte, err error) {
	data, err = sdk.getAsByteArray(probe)
	return
}
func (sdk *SDKCore) ParseByteArrayToTemplate(data *[]byte) (template *templates.SearchTemplate, err error) {
	template, err = sdk.parseByteArrayToTemplate(data)
	return
}
