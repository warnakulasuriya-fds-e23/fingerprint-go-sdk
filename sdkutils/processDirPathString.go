package sdkutils

import (
	"fmt"
	"os"
	"path/filepath"
)

type ErrInvalidPath struct {
	Path  string
	Issue string
}

func (e *ErrInvalidPath) Error() string {
	return fmt.Sprintf("Invalid Path %s because %s", e.Path, e.Issue)
}

type ErrPathAccess struct {
	Path       string
	Permission string
	Cause      error
}

func (e *ErrPathAccess) Error() string {
	return fmt.Sprintf("path %s is inaccessible under current permission : %s, because %v", e.Path, e.Permission, e.Cause)
}

func ProcessDirPathString(pathString *string) error {
	if pathString == nil || *pathString == "" {
		return &ErrInvalidPath{Path: "nil or empty", Issue: "Path provided cannot be nil or empty"}
	}

	originalPath := *pathString

	cleanedPath := filepath.Clean(originalPath)

	if !filepath.IsAbs(cleanedPath) {
		return &ErrInvalidPath{Path: originalPath, Issue: "path must be absolute"}
	}

	fileinfor, err := os.Stat(cleanedPath)
	if err != nil {
		switch {
		case os.IsNotExist(err):
			errMkdir := os.MkdirAll(cleanedPath, os.ModePerm)
			if errMkdir != nil {
				return fmt.Errorf("the directory %s doesnt exist. Attempt to create that directory has failed because %w ", originalPath, errMkdir)
			}

		case os.IsPermission(err):
			return &ErrPathAccess{Path: originalPath, Permission: "Read/Write", Cause: err}
		default:
			return fmt.Errorf("Error accessing Path %s because %w", originalPath, err)
		}

	}

	if !fileinfor.IsDir() {
		return &ErrInvalidPath{Path: originalPath, Issue: "Path exists but not a directory"}
	}

	testfileName := fmt.Sprintf(".%d_write_test", os.Getpid())
	testfilePath := filepath.Join(cleanedPath, testfileName)
	testfile, writeErr := os.Create(testfilePath)
	if writeErr != nil {
		return &ErrPathAccess{Path: originalPath, Permission: "write", Cause: writeErr}
	}
	testfile.Close()
	os.Remove(testfilePath)

	*pathString = cleanedPath
	return nil

}
