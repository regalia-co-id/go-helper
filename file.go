package helpers

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func UploadFile(file *multipart.FileHeader, folderName string, idAsFileName string) (string, error) {
	// Source
	src, err := file.Open()
	if err != nil {
		return "", err
	}

	defer src.Close()

	arrayExt := strings.Split(file.Filename, ".")
	fileExt := arrayExt[len(arrayExt)-1]

	var str strings.Builder
	str.WriteString("assets")
	str.WriteString(string(filepath.Separator))
	str.WriteString(folderName)
	str.WriteString(string(filepath.Separator))
	str.WriteString(idAsFileName)
	str.WriteString(".")
	str.WriteString(fileExt)

	fileFullPath := str.String()

	// Destination
	dst, err := os.Create(fileFullPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return fileFullPath, nil
}
