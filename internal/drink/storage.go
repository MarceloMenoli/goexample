package drink

import (
	"io"
)

type Storage interface {
	UploadFile(key string, body io.ReadSeeker, contentType string) (string, error)
}
