package s3

import "strings"

var ContentType = struct {
	ApplicationJson   string
	ImagePng          string
	ImageJpeg         string
	BinaryOctetStream string
}{
	ApplicationJson:   "application/json",
	ImagePng:          "image/png",
	ImageJpeg:         "image/jpeg",
	BinaryOctetStream: "binary/octet-stream",
}

func getContentType(fileName string) string {
	// get file extention
	fileParts := strings.Split(fileName, ".")

	switch fileParts[len(fileParts)-1] {
	case "json":
		return ContentType.ApplicationJson
	case "png":
		return ContentType.ImagePng
	case "jpeg", "jpg":
		return ContentType.ImageJpeg
	}

	return ContentType.BinaryOctetStream
}
