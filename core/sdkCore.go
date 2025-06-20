package core

import (
	"context"
	"log"
	"runtime"

	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/entities"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/sdkutils"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/config"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/templates"
)

// sdkCore will be an opaque type and can only be initialized using NewsdkCore
type sdkCore struct {
	gallery            *[]*entities.SearchTemplateRecord
	transparencyLogger *sourceafis.DefaultTransparencyLogger
	templateCreator    *sourceafis.TemplateCreator
	imagesDir          string
	cborDir            string
	cntx               context.Context
	matchThreshold     float64
}

func NewsdkCore(imagesDir string, cborDir string) (*sdkCore, error) {
	config.LoadDefaultConfig()
	config.Config.Workers = runtime.NumCPU()

	l := sourceafis.NewTransparencyLogger(new(CustomTransparencyContents))
	tc := sourceafis.NewTemplateCreator(l)
	g := make([]*entities.SearchTemplateRecord, 0, 1300)
	cntx := context.Background()
	matchThreshold := 40.00
	sdk := &sdkCore{
		gallery:            &g,
		transparencyLogger: l,
		templateCreator:    tc,
		imagesDir:          imagesDir,
		cborDir:            cborDir,
		cntx:               cntx,
		matchThreshold:     matchThreshold,
	}
	log.Println("SDK core has been initialized")

	return sdk, nil

}
func (sdk *sdkCore) LoadImages() {
	sdkutils.LoadImagesDirToGallery(sdk.gallery, sdk.imagesDir)
}
func (sdk *sdkCore) LoadCborfiles() {
	sdkutils.LoadCborDirToGallery(sdk.gallery, sdk.cborDir)
}
func (sdk *sdkCore) Extract(imagePath string) (template *templates.SearchTemplate) {
	template = sdk.extract(imagePath)
	return
}
func (sdk *sdkCore) Match(probe *templates.SearchTemplate, candidate *templates.SearchTemplate) (isMatched bool) {
	isMatched = sdk.match(probe, candidate)
	return
}
func (sdk *sdkCore) Identify(probe *templates.SearchTemplate) (isMatched bool, discoveredId string) {
	isMatched, discoveredId = sdk.identify(probe)
	return
}
