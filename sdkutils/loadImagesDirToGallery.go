package sdkutils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/entities"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork"
)

type TransparencyContents struct {
}

func (c *TransparencyContents) Accepts(key string) bool {
	return true
}

func (c *TransparencyContents) Accept(key, mime string, data []byte) error {
	//fmt.Printf("%d B  %s %s \n", len(data), mime, key)
	return nil
}

func loadImageToGallery(galleryptr *[]*entities.SearchTemplateRecord, imagePath string, filename string) error {
	image, err := sourceafis.LoadImage(imagePath)
	if err != nil {
		return fmt.Errorf("couldn't load image at path: %s, error: %w", imagePath, err)
	}
	l := sourceafis.NewTransparencyLogger(new(TransparencyContents))
	tc := sourceafis.NewTemplateCreator(l)

	template, err := tc.Template(image)
	if err != nil {
		return fmt.Errorf("couldn't extract template from image at path: %s, error: %w ", imagePath, err)
	}
	fileNameWithoutExtenstion := strings.TrimSuffix(filename, filepath.Ext(filename))
	*galleryptr = append(*galleryptr, &entities.SearchTemplateRecord{Id: fileNameWithoutExtenstion, Template: *template})

	log.Printf("appended template of %s to gallery\n", filename)
	return nil

}

func LoadImagesDirToGallery(galleryptr *[]*entities.SearchTemplateRecord, imagesDir string) error {
	// iterate over the files in the imagesDir get their absolute path and
	// for every iteration execute the loadImageToGallery function by passing that path
	// as an argument
	files, err := os.ReadDir(imagesDir)
	if err != nil {
		return fmt.Errorf("error reading directory: %w", err)
	}

	for _, file := range files {
		filename := file.Name()
		err = loadImageToGallery(galleryptr, filepath.Join(imagesDir, filename), filename)
		if err != nil {
			return err
		}
	}
	return nil
}
