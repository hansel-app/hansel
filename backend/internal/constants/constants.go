package constants

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

const GretelUserId = 1

var GretelGemAttachment []byte

//nolint:gochecknoinits
func init() {
	gretelGemAttachmentPath := filepath.Join("assets", "gretel-gem-attachment.png")
	imageBytes, ioErr := ioutil.ReadFile(filepath.Clean(gretelGemAttachmentPath))
	if ioErr != nil {
		log.Fatalf("failed to read gretel avatar file: %v", ioErr)
	}
	GretelGemAttachment = imageBytes
}
